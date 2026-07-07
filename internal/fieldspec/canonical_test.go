package fieldspec_test

import (
	"reflect"
	"testing"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
)

func TestAddressListCodecRoundTripAndRejectsNonCanonical(t *testing.T) {
	decoded, err := fieldspec.DecodeAddressList(`["seat-b","seat-d","seat-c"]`)
	if err != nil {
		t.Fatalf("DecodeAddressList canonical: %v", err)
	}
	if want := []string{"seat-b", "seat-d", "seat-c"}; !reflect.DeepEqual(decoded, want) {
		t.Fatalf("decoded = %v, want %v", decoded, want)
	}

	encoded, err := fieldspec.EncodeAddressList([]string{"seat-b", "seat-d", "seat-c"})
	if err != nil {
		t.Fatalf("EncodeAddressList: %v", err)
	}
	if encoded != `["seat-b","seat-d","seat-c"]` {
		t.Fatalf("encoded = %q, want canonical address list", encoded)
	}

	empty, err := fieldspec.DecodeAddressList("")
	if err != nil {
		t.Fatalf("DecodeAddressList empty: %v", err)
	}
	if empty != nil {
		t.Fatalf("empty decoded = %v, want nil", empty)
	}

	if _, err := fieldspec.DecodeAddressList(`["seat-b", "seat-c"]`); err == nil {
		t.Fatalf("DecodeAddressList accepted non-canonical JSON")
	}
}

func TestCanonicalMarshalDeterministicNoWhitespace(t *testing.T) {
	got, err := fieldspec.CanonicalMarshal([]map[string]string{
		{"status": "in", "path": "README.md"},
	})
	if err != nil {
		t.Fatalf("CanonicalMarshal: %v", err)
	}
	if got != `[{"path":"README.md","status":"in"}]` {
		t.Fatalf("canonical = %q", got)
	}
}

func TestParseTypedStructuredRejectsEquivalentNonCanonicalEncoding(t *testing.T) {
	spec := &fieldspec.FieldSpec{ID: "SCOPE_DIFF", Type: "row_array"}
	raw := `[{ "status" : "in", "path" : "README.md" }]`
	if _, err := fieldspec.ParseTyped(spec, raw); err == nil {
		t.Fatalf("ParseTyped accepted non-canonical structured header")
	}

	canonical := `[{"path":"README.md","status":"in"}]`
	value, err := fieldspec.ParseTyped(spec, canonical)
	if err != nil {
		t.Fatalf("ParseTyped canonical row_array: %v", err)
	}
	rows, ok := value.([]map[string]string)
	if !ok || len(rows) != 1 || rows[0]["path"] != "README.md" {
		t.Fatalf("row_array value = %#v", value)
	}
}

func TestParseTypedPerTypeValidation(t *testing.T) {
	cases := []struct {
		name string
		spec fieldspec.FieldSpec
		raw  string
	}{
		{"object", fieldspec.FieldSpec{ID: "OBJ", Type: "object"}, `{"a":"1","b":"2"}`},
		{"address list", fieldspec.FieldSpec{ID: "TO", Type: "address_list"}, `["a","b"]`},
		{"bool", fieldspec.FieldSpec{ID: "HUMAN_GATE_REQUIRED", Type: "bool"}, "yes"},
		{"int", fieldspec.FieldSpec{ID: "COUNT", Type: "int"}, "42"},
		{"enum", fieldspec.FieldSpec{ID: "PHASE", Type: "enum"}, "IMPL"},
		{"id ref", fieldspec.FieldSpec{ID: "PARENT_DISPATCH_ID", Type: "id_ref"}, "dispatch-1"},
		{"evidence ref", fieldspec.FieldSpec{ID: "EVIDENCE", Type: "evidence_ref"}, "E1"},
		{"text", fieldspec.FieldSpec{ID: "SUBJECT", Type: "text"}, "subject"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := fieldspec.ParseTyped(&tc.spec, tc.raw); err != nil {
				t.Fatalf("ParseTyped(%s): %v", tc.name, err)
			}
		})
	}

	for _, tc := range []struct {
		name string
		spec fieldspec.FieldSpec
		raw  string
	}{
		{"bad row array", fieldspec.FieldSpec{ID: "SCOPE_DIFF", Type: "row_array"}, `{"path":"README.md"}`},
		{"bad object", fieldspec.FieldSpec{ID: "OBJ", Type: "object"}, `["a"]`},
		{"bad address list", fieldspec.FieldSpec{ID: "TO", Type: "address_list"}, `["a",1]`},
		{"bad bool", fieldspec.FieldSpec{ID: "HUMAN_GATE_REQUIRED", Type: "bool"}, "maybe"},
		{"bad int", fieldspec.FieldSpec{ID: "COUNT", Type: "int"}, "NaN"},
		{"bad scalar", fieldspec.FieldSpec{ID: "SUBJECT", Type: "text"}, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := fieldspec.ParseTyped(&tc.spec, tc.raw); err == nil {
				t.Fatalf("ParseTyped(%s) succeeded", tc.name)
			}
		})
	}
}

func TestRecordSealVerifyKeepsCanonicalStructuredHeaderStable(t *testing.T) {
	raw := `[{"path":"README.md","status":"in"}]`
	rec := record.Record{
		Envelope: record.Envelope{RelayID: "r1", From: "seat", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "IMPL", "SCOPE_DIFF": raw},
		Body:     "body",
	}
	data, err := record.Seal(rec)
	if err != nil {
		t.Fatalf("Seal: %v", err)
	}
	got, err := record.Verify(data)
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}
	if got.Headers["SCOPE_DIFF"] != raw {
		t.Fatalf("SCOPE_DIFF = %q, want %q", got.Headers["SCOPE_DIFF"], raw)
	}
}
