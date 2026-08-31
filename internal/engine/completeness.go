package engine

import (
	"sort"

	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/lineage"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
)

var observeProducerManifest = map[string]bool{
	"achieved_evidence":        true,
	"evidence_integrity":       true,
	"executable_claim_results": true,
	"egress_scan_result":       true,
	"degradation_notes":        true,
	"attestation_source":       true,
	"target_gap_result":        true,
	"record_integrity":         true,
	"deviated_observed":        true,
	"bucket_binding_observed":  true,
}

func AuthorityClass(reg *fieldspec.Registry, cand record.Record, meta seat.SeatMeta) string {
	if (&lineage.Engine{Reg: reg}).AuthorityBearing(cand, meta) {
		return "yes"
	}
	return "no"
}

func CompleteObserved(cand record.Record, writes map[string]string) (record.Record, *fieldspec.Violation) {
	for _, field := range sortedKeys(writes) {
		if !observeProducerManifest[field] {
			return cand, &fieldspec.Violation{Field: field, Class: "producer-manifest"}
		}
	}
	for _, field := range sortedManifestFields() {
		if _, ok := writes[field]; !ok {
			return cand, &fieldspec.Violation{Field: field, Class: "producer-incomplete"}
		}
	}
	if cand.Headers == nil {
		cand.Headers = map[string]string{}
	}
	if cand.Headers["authority_class"] == "" {
		return cand, &fieldspec.Violation{Field: "authority_class", Class: "producer-incomplete"}
	}
	for field, value := range writes {
		cand.Headers[field] = value
	}
	if isGateCandidate(cand) {
		delete(cand.Headers, "surface_intent")
	} else {
		cand.Headers["surface_intent"] = "progress"
	}
	return cand, nil
}

func sortedManifestFields() []string {
	fields := make([]string, 0, len(observeProducerManifest))
	for field := range observeProducerManifest {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	return fields
}

func sortedKeys(values map[string]string) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
