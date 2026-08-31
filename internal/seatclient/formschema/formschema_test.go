package formschema

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/fieldspec"
)

func TestSchemaStableVolatilePartitionAndOpenAdvertisedHeaders(t *testing.T) {
	form := referenceForm(t)
	schema := SchemaFromForm(form, "ref-digest-1")
	properties := schema["properties"].(map[string]any)
	headers := properties["headers"].(map[string]any)
	if headers["additionalProperties"] != true {
		t.Fatalf("advertised headers are not open: %#v", headers)
	}
	headerProperties := headers["properties"].(map[string]any)
	if !reflect.DeepEqual(headerProperties["AUTHORITY"].(map[string]any)["enum"], []string{"read-only", "report-only"}) {
		t.Fatalf("stable enum missing: %#v", headerProperties["AUTHORITY"])
	}
	for _, name := range []string{"HUMAN_GATE_REQUIRED", "PARENT_DISPATCH_ID"} {
		property := headerProperties[name].(map[string]any)
		if _, enforcing := property["enum"]; enforcing {
			t.Fatalf("volatile %s carried enforcing enum: %#v", name, property)
		}
		if _, advisory := property["description"]; !advisory {
			t.Fatalf("volatile %s lacks advisory description: %#v", name, property)
		}
	}
	for _, skipped := range []string{"From", "TO", "CC", "DISPATCH_ID"} {
		if _, present := headerProperties[skipped]; present {
			t.Fatalf("schema retained skipped field %s", skipped)
		}
	}
}

func TestStrictParseTypedErrorIdentities(t *testing.T) {
	tests := []struct {
		name string
		args string
		id   ErrorID
	}{
		{"unknown", `{"headres":{},"form_digest":"d"}`, ErrorUnknownMember},
		{"case variant", `{"Body":"b","headers":{},"form_digest":"d"}`, ErrorCaseVariantMember},
		{"type", `{"headers":{"SUBJECT":123},"form_digest":"d"}`, ErrorMemberType},
		{"reserved system", `{"headers":{"relay_id":"r"},"form_digest":"d"}`, ErrorReservedSystemHeader},
		{"reserved envelope", `{"headers":{"TO":"x"},"form_digest":"d"}`, ErrorReservedEnvelopeHeader},
		{"cc conflict", `{"cc":"a","headers":{"CC":"b"},"form_digest":"d"}`, ErrorCCConflict},
		{"duplicate top", `{"body":"a","body":"b","headers":{},"form_digest":"d"}`, ErrorDuplicateTopMember},
		{"duplicate nested", `{"headers":{"SUBJECT":"a","SUBJECT":"b"},"form_digest":"d"}`, ErrorDuplicateNestedMember},
		{"trailing", `{"headers":{},"form_digest":"d"}{}`, ErrorTrailingData},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := ParseSubmitArguments(json.RawMessage(test.args))
			var mappingError *MappingError
			if !errors.As(err, &mappingError) || mappingError.ID != test.id {
				t.Fatalf("error = %v, want %s", err, test.id)
			}
		})
	}
}

func TestCCFoldAndPayloadShape(t *testing.T) {
	for _, arguments := range []string{
		`{"cc":"m-7.planner","headers":{"SUBJECT":"s"},"form_digest":"d"}`,
		`{"cc":"m-7.planner","headers":{"CC":"m-7.planner","SUBJECT":"s"},"form_digest":"d"}`,
	} {
		payload, err := SubmitPayloadFromArguments(json.RawMessage(arguments))
		if err != nil {
			t.Fatal(err)
		}
		var decoded struct {
			Headers map[string]string `json:"headers"`
		}
		if err := json.Unmarshal(payload, &decoded); err != nil {
			t.Fatal(err)
		}
		if decoded.Headers["CC"] != "m-7.planner" {
			t.Fatalf("CC fold = %#v", decoded.Headers)
		}
	}
}

func TestValidationProjection(t *testing.T) {
	form := referenceForm(t)
	tests := []struct {
		name string
		args string
		want []Disposition
	}{
		{"valid", `{"headers":{"AUTHORITY":"report-only","SUBJECT":"s"},"form_digest":"ref-digest-1"}`, nil},
		{"stable enum", `{"headers":{"AUTHORITY":"design-only","SUBJECT":"s"},"form_digest":"ref-digest-1"}`, []Disposition{{Class: "enum", Member: "AUTHORITY"}}},
		{"unknown", `{"headers":{"NOPE":"x","SUBJECT":"s"},"form_digest":"ref-digest-1"}`, []Disposition{{Class: "unknown-header", Member: "NOPE"}}},
		{"missing digest", `{"headers":{"SUBJECT":"s"}}`, []Disposition{{Class: "required", Member: "form_digest"}}},
		{"wrong digest", `{"headers":{"SUBJECT":"s"},"form_digest":"other"}`, []Disposition{{Class: "digest-mismatch", Member: "form_digest"}}},
		{"volatile enum", `{"headers":{"PARENT_DISPATCH_ID":"new","SUBJECT":"s"},"form_digest":"ref-digest-1"}`, nil},
		{"missing headers", `{"form_digest":"ref-digest-1"}`, []Disposition{{Class: "required", Member: "headers"}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := ValidateSubmitArguments(form, "ref-digest-1", json.RawMessage(test.args)); !reflect.DeepEqual(got, test.want) {
				t.Fatalf("dispositions = %#v, want %#v", got, test.want)
			}
		})
	}
}

func TestLockedReferenceFingerprintAndBranchCoverage(t *testing.T) {
	result, err := RunLockedReference()
	if err != nil {
		t.Fatal(err)
	}
	if result.Fingerprint != ExpectedFingerprint {
		t.Fatalf("fingerprint = %s, want %s", result.Fingerprint, ExpectedFingerprint)
	}
	if result.Fingerprint != "306b3149a002f0648995f659926ba0f985ee771d95bb899c7ecc81adadab28ac" {
		t.Fatal("expected fingerprint constant drifted")
	}
	if err := VerifyBranchCoverage(result.Exercises); err != nil {
		t.Fatal(err)
	}
	missing := make(map[string][]string, len(result.Exercises))
	for vector, branches := range result.Exercises {
		missing[vector] = append([]string(nil), branches...)
	}
	delete(missing, "S1")
	if err := VerifyBranchCoverage(missing); err == nil {
		t.Fatal("missing vector did not break branch coverage")
	}
	unknown := make(map[string][]string, len(result.Exercises)+1)
	for vector, branches := range result.Exercises {
		unknown[vector] = append([]string(nil), branches...)
	}
	unknown["extra"] = []string{"R-unknown"}
	if err := VerifyBranchCoverage(unknown); err == nil {
		t.Fatal("unknown branch did not break branch coverage")
	}
}

func TestStaticSchemaDigestsMatchLockedF58Producers(t *testing.T) {
	got, err := StaticSchemaDigests()
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"relay.submit":  "6bb7bbf46d8bf5d210cee410fbd0fa59106145425878c065adf0d54b05ace08e",
		"relay.project": "be5c41ec848bd7f6a7afd16af5acc56c65cf39bc113041941bb6747153bd582a",
		"relay.read":    "a84645cb3f57ea1172661ddcc42e8a710f5a320ee3ed6c944f5e469026b3036e",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("static schema digests = %#v, want %#v", got, want)
	}
}

func referenceForm(t *testing.T) fieldspec.Form {
	t.Helper()
	const encoded = `{"fields":{"AUTHORITY":{"default":"report-only","options":["read-only","report-only"],"type":"enum"},"CC":{"type":"address_list"},"CEREMONY_TIER":{"options":["small","medium","large"],"type":"enum"},"CONTEXT":{"type":"object"},"DISPATCH_ID":{"type":"id_ref"},"From":{"type":"text"},"HUMAN_GATE_REQUIRED":{"conductor_volatile":true,"digest_exempt":true,"options":["no"],"type":"bool"},"NOTIFY":{"type":"address_list"},"PARENT_DISPATCH_ID":{"conductor_volatile":true,"default":"parent-a","digest_exempt":true,"options":["parent-a"],"type":"id_ref"},"PHASE":{"options":["SITREP","PLAN"],"type":"enum"},"SCOPE_DIFF":{"type":"row_array"},"SUBJECT":{"type":"text"},"TITLE":{"type":"text"},"TO":{"type":"address_list"}}}`
	var form fieldspec.Form
	if err := json.Unmarshal([]byte(encoded), &form); err != nil {
		t.Fatal(err)
	}
	return form
}
