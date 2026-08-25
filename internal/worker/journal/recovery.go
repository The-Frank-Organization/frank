package journal

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/jackli/frank/internal/worker/wire"
)

const (
	DispositionResumable = "resumable"
	DispositionDegraded  = "degraded"
	ResumeActionReDerive = "re_derive"
)

// BoundaryKind names the two trusted resume-boundary kinds plus the empty
// prefix used for a genesis fault.
type BoundaryKind string

const (
	BoundaryNone    BoundaryKind = "none"
	BoundaryGenesis BoundaryKind = "genesis"
	BoundaryMarker  BoundaryKind = "round_marker"
)

const (
	FaultNone         = ""
	FaultGenesis      = "genesis_fault"
	FaultStructural   = "structural_fault"
	FaultSequence     = "sequence_fault"
	FaultContentTrust = "content_trust_fault"
	FaultMarker       = "round_marker_fault"
)

// Identity is the trusted comparand for the one-file genesis.
type Identity struct {
	RunID             string
	RunManifestDigest string
	CreateAuthID      string
}

// Boundary identifies the byte offset immediately after the boundary's line.
type Boundary struct {
	Kind   BoundaryKind
	Seq    uint64
	Offset int64
}

// Recovery is the deterministic result of one physical-file-order scan.
type Recovery struct {
	Disposition  string
	ResumeAction string
	Boundary     Boundary
	NextSeq      uint64
	FaultClass   string
	GenesisFault bool
	Records      [][]byte
}

// Recover scans one complete session.log byte string in physical file order.
// It never modifies the supplied bytes.
func Recover(data []byte, identity Identity) Recovery {
	result := Recovery{
		Disposition: DispositionResumable,
		Boundary:    Boundary{Kind: BoundaryNone},
		FaultClass:  FaultNone,
	}
	if err := validateIdentity(identity); err != nil {
		return genesisFailure(FaultGenesis)
	}

	var accepted []Record
	var acceptedRaw [][]byte
	var lastRaw []byte
	var lastSeq uint64
	var haveLast bool
	closedRounds := make(map[string]Boundary)
	position := 0
	for position < len(data) {
		relativeNewline := bytes.IndexByte(data[position:], '\n')
		if relativeNewline < 0 {
			return degrade(result, faultForPosition(result), result.Boundary)
		}
		lineEnd := position + relativeNewline
		raw := data[position:lineEnd]
		nextPosition := lineEnd + 1
		if len(raw) == 0 {
			return degrade(result, faultForPosition(result), result.Boundary)
		}
		record, err := DecodeRecord(raw)
		if err != nil {
			if !haveLast {
				return genesisFailure(FaultGenesis)
			}
			return degrade(result, FaultStructural, result.Boundary)
		}
		seq, err := wire.ParseCounter(record.Seq)
		if err != nil {
			if !haveLast {
				return genesisFailure(FaultGenesis)
			}
			return degrade(result, FaultSequence, result.Boundary)
		}
		if haveLast {
			if seq == lastSeq {
				if bytes.Equal(raw, lastRaw) {
					position = nextPosition
					continue
				}
				return degrade(result, FaultSequence, result.Boundary)
			}
			if seq != lastSeq+1 {
				return degrade(result, FaultSequence, result.Boundary)
			}
		} else {
			if record.Kind != KindRunOpen || seq != 0 {
				return genesisFailure(FaultGenesis)
			}
			if !matchesGenesis(record, identity) {
				return genesisFailure(FaultGenesis)
			}
		}

		if admitEligible(record.Kind) {
			key := record.TurnID + "\x00" + record.RoundIndex
			if previous, exists := closedRounds[key]; exists {
				return degrade(result, FaultMarker, previous)
			}
		}

		digestMatches, err := VerifyRecordDigest(record)
		if err != nil {
			if !haveLast {
				return genesisFailure(FaultGenesis)
			}
			return degrade(result, FaultStructural, result.Boundary)
		}
		if !digestMatches {
			subject, subjectErr := transitionSubject(record)
			if subjectErr != nil {
				if !haveLast {
					return genesisFailure(FaultGenesis)
				}
				return degrade(result, FaultStructural, result.Boundary)
			}
			action := transitionAction(subject, TrustDigestMismatch)
			if action == ActionResolveObjective {
				// No independently trusted objective resolver is supplied to this
				// pure reader; the safe result is the table's degrade fallback.
				action = ActionDegrade
			}
			if action == ActionDegrade {
				if !haveLast {
					return genesisFailure(FaultGenesis)
				}
				return degrade(result, FaultContentTrust, result.Boundary)
			}
		}

		if record.Kind == KindRoundMarker {
			members, outside := markerInputs(record, accepted)
			priorBoundary := result.Boundary
			if err := HonourRoundMarker(record, members, outside); err != nil {
				return degrade(result, FaultMarker, priorBoundary)
			}
			markerSeq, _ := wire.ParseCounter(record.Seq)
			result.Boundary = Boundary{Kind: BoundaryMarker, Seq: markerSeq, Offset: int64(nextPosition)}
			roundIndex, _ := rawStringValue(record.Fields["round_index"], "round_index")
			closedRounds[record.TurnID+"\x00"+roundIndex] = priorBoundary
		}

		accepted = append(accepted, record)
		acceptedRaw = append(acceptedRaw, append([]byte(nil), raw...))
		lastRaw = append(lastRaw[:0], raw...)
		lastSeq = seq
		haveLast = true
		position = nextPosition
		if record.Kind == KindRunOpen {
			result.Boundary = Boundary{Kind: BoundaryGenesis, Seq: 0, Offset: int64(nextPosition)}
		}
	}

	if !haveLast {
		return genesisFailure(FaultGenesis)
	}
	result.Records = acceptedRaw
	result.NextSeq = result.Boundary.Seq + 1
	return result
}

func markerInputs(marker Record, accepted []Record) ([]Record, []Record) {
	firstText, _ := rawStringValue(marker.Fields["first_seq"], "first_seq")
	lastText, _ := rawStringValue(marker.Fields["last_seq"], "last_seq")
	first, firstErr := wire.ParseCounter(firstText)
	last, lastErr := wire.ParseCounter(lastText)
	if firstErr != nil || lastErr != nil || last < first {
		return nil, append([]Record(nil), accepted...)
	}
	members := make([]Record, 0, last-first+1)
	outside := make([]Record, 0, len(accepted))
	for _, record := range accepted {
		seq, err := wire.ParseCounter(record.Seq)
		if err == nil && seq >= first && seq <= last {
			members = append(members, record)
		} else {
			outside = append(outside, record)
		}
	}
	return members, outside
}

func matchesGenesis(record Record, identity Identity) bool {
	runID, err := rawStringValue(record.Fields["run_id"], "run_id")
	if err != nil || runID != identity.RunID {
		return false
	}
	manifest, err := rawStringValue(record.Fields["run_manifest_digest"], "run_manifest_digest")
	if err != nil || manifest != identity.RunManifestDigest {
		return false
	}
	auth, err := rawStringValue(record.Fields["create_auth_id"], "create_auth_id")
	return err == nil && auth == identity.CreateAuthID
}

func validateIdentity(identity Identity) error {
	if identity.RunID == "" || !lowerHex(identity.RunManifestDigest, 64) || !lowerHex(identity.CreateAuthID, 32) {
		return errors.New("invalid journal identity")
	}
	return nil
}

func genesisFailure(faultClass string) Recovery {
	return Recovery{
		Disposition:  DispositionDegraded,
		ResumeAction: ResumeActionReDerive,
		Boundary:     Boundary{Kind: BoundaryNone},
		NextSeq:      0,
		FaultClass:   faultClass,
		GenesisFault: true,
	}
}

func degrade(result Recovery, faultClass string, boundary Boundary) Recovery {
	result.Disposition = DispositionDegraded
	result.ResumeAction = ResumeActionReDerive
	result.Boundary = boundary
	result.NextSeq = boundary.Seq + 1
	result.FaultClass = faultClass
	return result
}

func faultForPosition(result Recovery) string {
	if result.Boundary.Kind == BoundaryNone {
		return FaultGenesis
	}
	return FaultStructural
}

func (recovery Recovery) validateForTruncation() error {
	if recovery.GenesisFault || recovery.Boundary.Kind == BoundaryNone {
		return errors.New("genesis fault has no truncation boundary")
	}
	if recovery.Disposition != DispositionResumable && !(recovery.Disposition == DispositionDegraded && recovery.ResumeAction == ResumeActionReDerive) {
		return fmt.Errorf("unsupported recovery disposition %q", recovery.Disposition)
	}
	return nil
}
