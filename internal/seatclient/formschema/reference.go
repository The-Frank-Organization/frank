package formschema

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/The-Frank-Organization/frank/internal/fieldspec"
)

const ExpectedFingerprint = "306b3149a002f0648995f659926ba0f985ee771d95bb899c7ecc81adadab28ac"

var BranchInventory = []string{
	"R-1.a", "R-1.b", "R-2.a", "R-2.b", "R-2.c", "R-2.d", "R-3.a", "R-3.b", "R-3.c", "R-3.d", "R-3.e", "R-4.a", "R-4.b", "R-5.a",
	"P-0.a", "P-0.b", "P-1.a", "P-1.b", "P-1.c", "P-2.a", "P-2.b", "P-3.a", "P-3.b", "P-3.c", "P-6.a", "P-6.b", "P-6.c",
	"V-0.a", "V-1.a", "V-1.b", "V-1.c", "V-1.d", "V-1.e", "V-2.a", "V-2.b",
	"RR-1.a", "RR-1.b", "RR-2.a", "RR-2.b", "RR-3.a", "RR-3.b",
}

type ReferenceResult struct {
	Fingerprint string
	Exercises   map[string][]string
	Records     [][2]any
}

type referenceVector struct {
	id        string
	op        string
	exercises []string
	args      string
	outcome   string
}

func RunLockedReference() (ReferenceResult, error) {
	form, err := lockedReferenceForm()
	if err != nil {
		return ReferenceResult{}, err
	}
	vectors := lockedVectors()
	records := make([][2]any, 0, len(vectors))
	exercises := make(map[string][]string, len(vectors))
	for _, vector := range vectors {
		result, runErr := runReferenceVector(form, vector)
		if runErr != nil {
			return ReferenceResult{}, fmt.Errorf("reference vector %s: %w", vector.id, runErr)
		}
		records = append(records, [2]any{vector.id, result})
		exercises[vector.id] = append([]string(nil), vector.exercises...)
	}
	encoded, err := json.Marshal(records)
	if err != nil {
		return ReferenceResult{}, err
	}
	digest := sha256.Sum256(encoded)
	return ReferenceResult{Fingerprint: hex.EncodeToString(digest[:]), Exercises: exercises, Records: records}, nil
}

func VerifyLockedReference() error {
	result, err := RunLockedReference()
	if err != nil {
		return err
	}
	if result.Fingerprint != ExpectedFingerprint {
		return fmt.Errorf("formschema: locked fingerprint mismatch: %s != %s", result.Fingerprint, ExpectedFingerprint)
	}
	return VerifyBranchCoverage(result.Exercises)
}

func VerifyBranchCoverage(exercises map[string][]string) error {
	want := make(map[string]struct{}, len(BranchInventory))
	for _, branch := range BranchInventory {
		if _, duplicate := want[branch]; duplicate {
			return fmt.Errorf("formschema: duplicate inventory branch %s", branch)
		}
		want[branch] = struct{}{}
	}
	got := make(map[string]struct{})
	for vector, branches := range exercises {
		if vector == "" {
			return errors.New("formschema: empty vector id")
		}
		for _, branch := range branches {
			if _, known := want[branch]; !known {
				return fmt.Errorf("formschema: vector %s cites unknown branch %s", vector, branch)
			}
			got[branch] = struct{}{}
		}
	}
	var missing []string
	for _, branch := range BranchInventory {
		if _, covered := got[branch]; !covered {
			missing = append(missing, branch)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("formschema: uncovered branches %v", missing)
	}
	return nil
}

func runReferenceVector(form fieldspec.Form, vector referenceVector) (map[string]any, error) {
	switch vector.op {
	case "schema":
		encoded, err := json.Marshal(SchemaFromForm(form, "ref-digest-1"))
		return map[string]any{"bytes": string(encoded), "outcome": "schema"}, err
	case "payload":
		payload, err := SubmitPayloadFromArguments(json.RawMessage(vector.args))
		if err != nil {
			var mappingError *MappingError
			if !AsMappingError(err, &mappingError) {
				return nil, err
			}
			return map[string]any{"error": string(mappingError.ID), "outcome": "mapping_error"}, nil
		}
		return map[string]any{"outcome": "mapped", "payload": string(payload)}, nil
	case "validate":
		dispositions := ValidateSubmitArguments(form, "ref-digest-1", json.RawMessage(vector.args))
		if len(dispositions) == 0 {
			return map[string]any{"outcome": "valid"}, nil
		}
		return map[string]any{"dispositions": dispositions, "outcome": "schema_invalid"}, nil
	case "rr_detect":
		return map[string]any{"detected": SubmitNeedsReRender(json.RawMessage(vector.outcome)), "outcome": "re_render_detect"}, nil
	case "rr_result":
		return map[string]any{"bytes": string(ReRenderResult(json.RawMessage(vector.outcome))), "outcome": "re_render_result"}, nil
	case "rr_key":
		arguments, err := ParseSubmitArguments(json.RawMessage(vector.args))
		if err != nil {
			return nil, err
		}
		phase, tier := DeclaredPhaseTier(arguments)
		return map[string]any{"outcome": "refresh_key", "phase": phase, "tier": tier}, nil
	default:
		return nil, fmt.Errorf("unknown operation %s", vector.op)
	}
}

func lockedReferenceForm() (fieldspec.Form, error) {
	const encoded = `{"fields":{"AUTHORITY":{"default":"report-only","options":["read-only","report-only"],"type":"enum"},"CC":{"type":"address_list"},"CEREMONY_TIER":{"options":["small","medium","large"],"type":"enum"},"CONTEXT":{"type":"object"},"DISPATCH_ID":{"type":"id_ref"},"From":{"type":"text"},"HUMAN_GATE_REQUIRED":{"conductor_volatile":true,"digest_exempt":true,"options":["no"],"type":"bool"},"NOTIFY":{"type":"address_list"},"PARENT_DISPATCH_ID":{"conductor_volatile":true,"default":"parent-a","digest_exempt":true,"options":["parent-a"],"type":"id_ref"},"PHASE":{"options":["SITREP","PLAN"],"type":"enum"},"SCOPE_DIFF":{"type":"row_array"},"SUBJECT":{"type":"text"},"TITLE":{"type":"text"},"TO":{"type":"address_list"}}}`
	var form fieldspec.Form
	err := json.Unmarshal([]byte(encoded), &form)
	return form, err
}

func lockedVectors() []referenceVector {
	return []referenceVector{
		{"S1", "schema", []string{"R-1.a", "R-1.b", "R-2.a", "R-2.b", "R-2.c", "R-2.d", "R-3.a", "R-3.b", "R-3.c", "R-3.d", "R-3.e", "R-4.a", "R-4.b", "R-5.a"}, "", ""},
		{"P1", "payload", []string{"P-0.a"}, `{"headers":{"AUTHORITY":"report-only","PHASE":"SITREP","SUBJECT":"s"},"to":"m-2.implementer","dispatch_id":"d-1","body":"b","form_digest":"ref-digest-1"}`, ""},
		{"P2", "payload", []string{"P-0.b"}, `{"headers":{"SUBJECT":"s"},"form_digest":"ref-digest-1"}`, ""},
		{"P3", "payload", []string{"P-1.a"}, `{"headres":{"SUBJECT":"s"},"form_digest":"ref-digest-1"}`, ""},
		{"P4", "payload", []string{"P-1.b"}, `{"Body":"b","headers":{"SUBJECT":"s"},"form_digest":"ref-digest-1"}`, ""},
		{"P5", "payload", []string{"P-1.c"}, `{"headers":{"SUBJECT":123},"form_digest":"ref-digest-1"}`, ""},
		{"P6", "payload", []string{"P-2.a"}, `{"headers":{"relay_id":"r-9"},"form_digest":"ref-digest-1"}`, ""},
		{"P7", "payload", []string{"P-2.a"}, `{"headers":{"From":"x"},"form_digest":"ref-digest-1"}`, ""},
		{"P8", "payload", []string{"P-2.b"}, `{"headers":{"TO":"x"},"form_digest":"ref-digest-1"}`, ""},
		{"P9", "payload", []string{"P-3.a"}, `{"cc":"m-7.planner","headers":{"SUBJECT":"s"},"form_digest":"ref-digest-1"}`, ""},
		{"P10", "payload", []string{"P-3.b"}, `{"cc":"m-7.planner","headers":{"CC":"m-7.planner","SUBJECT":"s"},"form_digest":"ref-digest-1"}`, ""},
		{"P11", "payload", []string{"P-3.c"}, `{"cc":"m-7.planner","headers":{"CC":"m-9.planner","SUBJECT":"s"},"form_digest":"ref-digest-1"}`, ""},
		{"P12", "payload", []string{"P-6.a"}, `{"body":"b1","body":"b2","headers":{"SUBJECT":"s"},"form_digest":"ref-digest-1"}`, ""},
		{"P13", "payload", []string{"P-6.b"}, `{"headers":{"SUBJECT":"s1","SUBJECT":"s2"},"form_digest":"ref-digest-1"}`, ""},
		{"P14", "payload", []string{"P-6.c"}, `{"headers":{"SUBJECT":"s"},"form_digest":"ref-digest-1"}{}`, ""},
		{"V1", "validate", []string{"V-0.a"}, `{"headers":{"AUTHORITY":"report-only","PHASE":"SITREP","SUBJECT":"s"},"to":"m-2.implementer","dispatch_id":"d-1","body":"b","form_digest":"ref-digest-1"}`, ""},
		{"V2", "validate", []string{"V-1.a"}, `{"headers":{"AUTHORITY":"design-only","SUBJECT":"s"},"form_digest":"ref-digest-1"}`, ""},
		{"V3", "validate", []string{"V-1.b"}, `{"headers":{"NOPE":"x","SUBJECT":"s"},"form_digest":"ref-digest-1"}`, ""},
		{"V4", "validate", []string{"V-1.c"}, `{"headers":{"SUBJECT":"s"}}`, ""},
		{"V5", "validate", []string{"V-1.d"}, `{"headers":{"SUBJECT":"s"},"form_digest":"other-digest"}`, ""},
		{"V6", "validate", []string{"V-2.a"}, `{"headers":{"PARENT_DISPATCH_ID":"parent-zz","SUBJECT":"s"},"form_digest":"ref-digest-1"}`, ""},
		{"V7", "validate", []string{"V-2.b"}, `{"headers":{"HUMAN_GATE_REQUIRED":"yes","SUBJECT":"s"},"form_digest":"ref-digest-1"}`, ""},
		{"V8", "validate", []string{"V-1.e"}, `{"form_digest":"ref-digest-1"}`, ""},
		{"R1", "rr_detect", []string{"RR-1.a"}, "", `{"state":"rejected","detail":"violation form_digest: re-render - stale form digest"}`},
		{"R2", "rr_detect", []string{"RR-1.b"}, "", `{"state":"rejected","detail":"violation AUTHORITY: enum - unknown AUTHORITY"}`},
		{"R3", "rr_result", []string{"RR-2.a", "RR-2.b"}, "", `{"state":"rejected","relay_id":"r-1","intake_id":"i-1","detail":"violation form_digest: re-render - stale form digest"}`},
		{"R4", "rr_key", []string{"RR-3.a"}, `{"headers":{"CEREMONY_TIER":"large","PHASE":"PLAN","SUBJECT":"s"},"form_digest":"ref-digest-1"}`, ""},
		{"R5", "rr_key", []string{"RR-3.b"}, `{"headers":{"SUBJECT":"s"},"form_digest":"ref-digest-1"}`, ""},
	}
}
