package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	frankconfig "github.com/The-Frank-Organization/frank/internal/config"
	"github.com/The-Frank-Organization/frank/internal/crashpoint"
	"github.com/The-Frank-Organization/frank/internal/derived"
	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/record"
	frankrecover "github.com/The-Frank-Organization/frank/internal/recover"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
	"github.com/The-Frank-Organization/frank/internal/tables"
)

var (
	errRecoverySelectionRequired = errors.New("recovery-selection-required")
	errRecoverySelectionInvalid  = errors.New("recovery-selection-invalid")
	errRecoveryTargetUnknown     = errors.New("recovery-target-unknown")
	errRecoveryTargetHealthy     = errors.New("recovery-target-not-quarantined")
	errAnchorTargetResolved      = errors.New("anchor-target-resolved")
	errDeliveryRetryUnauthorized = errors.New("delivery-retry-not-authorized")
	errRecoveryAuthorityFields   = errors.New("recovery-authority-fields-forbidden")
	errRecoverySelectorWithRetry = errors.New("recovery-selector-forbidden-on-retry")
)

// runRecoveryCeremony is deliberately a one-shot, process-local verb. The
// root lock is phase -1: no store, binding, recovery, or socket operation may
// occur before AcquireRoot succeeds.
func runRecoveryCeremony(ctx context.Context, cfg config) error {
	rootLock, err := store.AcquireRoot(cfg.Root)
	if err != nil {
		return err
	}
	defer func() { _ = rootLock.Release() }()
	if cfg.RecoverHasAuth {
		return errRecoveryAuthorityFields
	}
	if cfg.RetryReason != "" && cfg.RecoverSelect != "" {
		return errRecoverySelectorWithRetry
	}

	socket := cfg.Socket
	if socket == "" {
		socket = filepath.Join(cfg.Root, "frank.sock")
	}
	if socketIsLive(ctx, socket) {
		fmt.Fprintln(os.Stderr, "warning: conductor socket appears live; root lock remains authoritative")
	}

	st, err := store.Open(cfg.Root)
	if err != nil {
		return err
	}
	if err := st.CompleteAdoptionConfig(); err != nil {
		return err
	}
	if err := store.RequireAdoptedConfig(cfg.Root); err != nil {
		return err
	}
	pinned, err := frankconfig.Load(store.StoreRootConfigPaths(cfg.Root))
	if err != nil {
		return err
	}
	if _, err := frankrecover.Run(cfg.Root, pinned); err != nil {
		return err
	}
	if err := runMintUpgradeAnchors(st); err != nil {
		return err
	}

	mgr, err := seat.Open(cfg.Root)
	if err != nil {
		return err
	}
	records, err := st.Records()
	if err != nil {
		return err
	}
	chains, err := engine.BuildMintChains(records)
	if err != nil {
		return err
	}
	chain, exists := chains[cfg.RecoverSeat]
	if !exists {
		return errRecoveryTargetUnknown
	}
	anchoredByCeremony := false
	retryRotation := false

	if cfg.RecoverSelect != "" && !chain.Conflicted {
		if err := commitResolvedAnchorAnomaly(st, cfg.RecoverSeat, cfg.RecoverSelect); err != nil {
			return err
		}
		return errAnchorTargetResolved
	}

	if chain.Conflicted {
		if cfg.RecoverSelect == "" {
			return errRecoverySelectionRequired
		}
		if !validLegacySelection(records, cfg.RecoverSeat, cfg.RecoverSelect) {
			return errRecoverySelectionInvalid
		}
		anchor := engine.MintChainAnchorRecord(cfg.RecoverSeat, cfg.RecoverSelect)
		anchor.Headers["admin_provenance"] = "ceremony"
		if _, err := st.Commit(anchor, nil); err != nil {
			return err
		}
		anchoredByCeremony = true
		crashpoint.Hit("ceremony_post_anchor")

		records, err = st.Records()
		if err != nil {
			return err
		}
		chains, err = engine.BuildMintChains(records)
		if err != nil {
			return err
		}
		chain = chains[cfg.RecoverSeat]
	}

	if chain.Conflicted || chain.Tip.Envelope.RelayID == "" {
		return errRecoverySelectionInvalid
	}
	if err := terminalizeRealizedCeremonyAttempt(st, mgr, chain.Tip); err != nil {
		return err
	}
	records, err = st.Records()
	if err != nil {
		return err
	}
	if realized, ok := mgr.RealizedMintRef(cfg.RecoverSeat); !anchoredByCeremony && ok && realized == chain.Tip.Envelope.RelayID {
		if cfg.RetryReason == "" {
			return errRecoveryTargetHealthy
		}
		if !ceremonyRetryAuthorized(records, cfg.RecoverSeat, chain.Tip) {
			return errDeliveryRetryUnauthorized
		}
		rotation, err := ceremonyRetryRecord(chain.Tip, cfg.RetryReason)
		if err != nil {
			return err
		}
		if violation := engine.ValidateCeremonyRetryAuthority(records, rotation); violation != nil {
			rotation.Envelope.DeliveryState = record.Rejected
			rotation.Headers["failing_edge"] = "retry-authority-delta"
			rotation.Body = violation.Class
			if _, err := st.Commit(rotation, nil); err != nil {
				return err
			}
			return errors.New("retry-authority-delta")
		}
		derived.Stamp(&rotation)
		relayID, err := st.Commit(rotation, nil)
		if err != nil {
			return err
		}
		rotation.Envelope.RelayID = relayID
		chain.Tip = rotation
		retryRotation = true
		crashpoint.Hit("ceremony_retry_post_pivot")
	}
	if chain.Tip.Headers["admin_provenance"] == "ceremony" {
		retryRotation = true
	}
	credential, err := completeCeremonyMint(st, mgr, chain.Tip, cfg.RecoverSeat)
	if err != nil {
		return err
	}
	if retryRotation {
		crashpoint.Hit("ceremony_retry_pre_reply")
		fmt.Print("credential=")
		crashpoint.Hit("ceremony_retry_partial_reply")
		fmt.Printf("%s\n", credential.Value)
		crashpoint.Hit("ceremony_retry_post_reply")
	} else {
		crashpoint.Hit("ceremony_pre_reply")
		fmt.Printf("credential=%s\n", credential.Value)
	}
	return nil
}

func terminalizeRealizedCeremonyAttempt(st *store.Store, mgr *seat.Manager, tip record.Record) error {
	if tip.Headers["hook_contract"] != derived.HookContractV1 {
		return nil
	}
	tab, err := tables.Build(st)
	if err != nil {
		return err
	}
	status, ok := tab.DerivedWork[tip.Envelope.RelayID]
	if !ok || status.Status != "unknown" {
		return nil
	}
	open, _, valid := derived.AttemptState(tab.Records, tip.Envelope.RelayID, "mint")
	if !valid || !open {
		return nil
	}
	req, violation := engine.ParseSeatMintBody(tip.Body, tip.Envelope.From)
	if violation != nil {
		return errRecoverySelectionInvalid
	}
	realized, exists := mgr.RealizedMintRef(req.Seat)
	if !exists || realized != tip.Envelope.RelayID {
		return nil
	}
	_, err = engine.CommitAutomaticDerived(st, derived.RealizedUndeliveredRecord(tip.Envelope.RelayID))
	return err
}

func completeCeremonyMint(st *store.Store, mgr *seat.Manager, tip record.Record, seatName string) (seat.Cred, error) {
	req, violation := engine.ParseSeatMintBody(tip.Body, tip.Envelope.From)
	if violation != nil || req.Seat != seatName {
		return seat.Cred{}, errRecoverySelectionInvalid
	}
	if tip.Headers["hook_contract"] == derived.HookContractV1 {
		tab, err := tables.Build(st)
		if err != nil {
			return seat.Cred{}, err
		}
		status, ok := tab.DerivedWork[tip.Envelope.RelayID]
		if !ok || status.Status == "failed" || status.Status == "" {
			return seat.Cred{}, errDeliveryRetryUnauthorized
		}
		open, predecessor, valid := derived.AttemptState(tab.Records, tip.Envelope.RelayID, "mint")
		if !valid {
			return seat.Cred{}, errDeliveryRetryUnauthorized
		}
		if !open {
			if _, err := engine.CommitAutomaticDerived(st, derived.AttemptRecord(tip.Envelope.RelayID, "mint", predecessor)); err != nil {
				return seat.Cred{}, err
			}
			crashpoint.Hit("ceremony_retry_post_marker")
		}
	}
	credential, err := mgr.MintOrReplace(req.Seat, req.Role, req.IsOperator, tip.Envelope.RelayID)
	if err != nil {
		return seat.Cred{}, err
	}
	crashpoint.Hit("ceremony_post_binding")
	if tip.Headers["hook_contract"] == derived.HookContractV1 {
		if _, err := engine.CommitAutomaticDerived(st, derived.CursorAdvanceRecord(tip.Envelope.RelayID, []string{"mint"})); err != nil {
			return seat.Cred{}, err
		}
	}
	return credential, nil
}

func ceremonyRetryRecord(tip record.Record, reason string) (record.Record, error) {
	req, violation := engine.ParseSeatMintBody(tip.Body, tip.Envelope.From)
	if violation != nil {
		return record.Record{}, errRecoverySelectionInvalid
	}
	body, err := json.Marshal(struct {
		Seat        string `json:"seat"`
		Role        string `json:"role"`
		IsOperator  bool   `json:"is_operator"`
		RetryReason string `json:"retry_reason"`
	}{Seat: req.Seat, Role: req.Role, IsOperator: req.IsOperator, RetryReason: reason})
	if err != nil {
		return record.Record{}, err
	}
	return record.Record{
		Envelope: record.Envelope{From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers: map[string]string{
			"PHASE": "SITREP", "SUBJECT": "offline ceremony delivery retry", "record_kind": "seat_mint",
			"mint_predecessor": tip.Envelope.RelayID, "admin_provenance": "ceremony",
		},
		Body: string(body),
	}, nil
}

func ceremonyRetryAuthorized(records []record.Record, seatName string, tip record.Record) bool {
	if tip.Envelope.DeliveryState != record.Accepted || tip.Headers["record_kind"] != "seat_mint" {
		return false
	}
	if tip.Headers["admin_provenance"] == "ceremony" {
		return true
	}
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Accepted {
			continue
		}
		switch rec.Headers["record_kind"] {
		case "mint-chain-anchor":
			if rec.Headers["admin_provenance"] != "ceremony" {
				continue
			}
			var body struct {
				Seat    string `json:"seat"`
				Selects string `json:"selects"`
			}
			if json.Unmarshal([]byte(rec.Body), &body) == nil && body.Seat == seatName && body.Selects == tip.Envelope.RelayID {
				return true
			}
		case "derived-work-transition":
			var body struct {
				SourceRelayID string `json:"source_relay_id"`
				Kind          string `json:"kind"`
			}
			if json.Unmarshal([]byte(rec.Body), &body) == nil && body.SourceRelayID == tip.Envelope.RelayID && body.Kind == "realized-undelivered" {
				return true
			}
		}
	}
	return false
}

func validLegacySelection(records []record.Record, seatName, relayID string) bool {
	for _, rec := range records {
		if rec.Envelope.RelayID != relayID || rec.Envelope.DeliveryState != record.Accepted || rec.Headers["record_kind"] != "seat_mint" || rec.Headers["mint_predecessor"] != "" {
			continue
		}
		req, violation := engine.ParseSeatMintBody(rec.Body, rec.Envelope.From)
		return violation == nil && req.Seat == seatName
	}
	return false
}

func commitResolvedAnchorAnomaly(st *store.Store, seatName, selects string) error {
	anomaly := engine.MintChainAnchorRecord(seatName, selects)
	anomaly.Envelope.DeliveryState = record.Rejected
	anomaly.Headers["admin_provenance"] = "ceremony"
	anomaly.Headers["failing_edge"] = "anchor-target-resolved"
	_, err := st.Commit(anomaly, nil)
	return err
}
