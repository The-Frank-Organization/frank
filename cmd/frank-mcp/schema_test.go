package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/fieldspec"
)

func TestSchemaStringCarrierAllTypes(t *testing.T) {
	form := fieldspec.Form{Fields: map[string]fieldspec.Field{
		"TITLE":               {Type: "text"},
		"AUTHORITY":           {Type: "enum", Options: []string{"read-only", "report-only"}, Default: "report-only"},
		"HUMAN_GATE_REQUIRED": {Type: "bool", Options: []string{"no", "yes"}},
		"PARENT_DISPATCH_ID":  {Type: "id_ref", Options: []string{"dispatch-a"}, Default: "dispatch-a"},
		"SCOPE_DIFF":          {Type: "row_array"},
		"TO":                  {Type: "address_list"},
		"CONTEXT":             {Type: "object"},
	}}
	schema := SchemaFromForm(form, "digest-1")
	headers := schema["properties"].(map[string]any)["headers"].(map[string]any)["properties"].(map[string]any)
	for _, field := range []string{"TITLE", "AUTHORITY", "HUMAN_GATE_REQUIRED", "PARENT_DISPATCH_ID", "SCOPE_DIFF", "CONTEXT"} {
		prop := headers[field].(map[string]any)
		if prop["type"] != "string" {
			t.Fatalf("%s type = %v, want string", field, prop["type"])
		}
	}
	if _, ok := headers["TO"]; ok {
		t.Fatalf("TO should be carried by top-level to, not headers: %#v", headers["TO"])
	}
	auth := headers["AUTHORITY"].(map[string]any)
	if !reflect.DeepEqual(auth["enum"], []string{"read-only", "report-only"}) || auth["default"] != "report-only" {
		t.Fatalf("AUTHORITY schema = %#v", auth)
	}
	parent := headers["PARENT_DISPATCH_ID"].(map[string]any)
	if !reflect.DeepEqual(parent["enum"], []string{"dispatch-a"}) || parent["default"] != "dispatch-a" {
		t.Fatalf("PARENT_DISPATCH_ID schema = %#v", parent)
	}
	for _, field := range []string{"SCOPE_DIFF", "CONTEXT"} {
		description := headers[field].(map[string]any)["description"].(string)
		if !strings.Contains(description, "canonical JSON string") {
			t.Fatalf("%s description = %q", field, description)
		}
	}
	top := schema["properties"].(map[string]any)
	for _, field := range []string{"to", "cc", "dispatch_id", "body"} {
		prop := top[field].(map[string]any)
		if prop["type"] != "string" {
			t.Fatalf("%s type = %v, want string", field, prop["type"])
		}
	}
}

func TestSchemaSystemFieldAbsence(t *testing.T) {
	form := fieldspec.Form{Fields: map[string]fieldspec.Field{
		"from":           {Type: "text"},
		"role":           {Type: "text"},
		"relay_id":       {Type: "id_ref"},
		"delivery_state": {Type: "enum", Options: []string{"accepted"}},
		"SUBJECT":        {Type: "text"},
	}}
	schema := SchemaFromForm(form, "digest-2")
	data, err := json.Marshal(schema)
	if err != nil {
		t.Fatalf("marshal schema: %v", err)
	}
	for _, forbidden := range []string{`"from"`, `"role"`, `"relay_id"`, `"delivery_state"`} {
		if strings.Contains(string(data), forbidden) {
			t.Fatalf("schema leaked %s: %s", forbidden, data)
		}
	}
	if !strings.Contains(string(data), `"SUBJECT"`) {
		t.Fatalf("schema dropped normal header: %s", data)
	}
}

func TestSchemaConstDigest(t *testing.T) {
	schema := SchemaFromForm(fieldspec.Form{Fields: map[string]fieldspec.Field{}}, "digest-3")
	top := schema["properties"].(map[string]any)
	formDigest := top["form_digest"].(map[string]any)
	if formDigest["type"] != "string" || formDigest["const"] != "digest-3" {
		t.Fatalf("form_digest schema = %#v", formDigest)
	}
	if !reflect.DeepEqual(schema["required"], []string{"headers", "form_digest"}) {
		t.Fatalf("required = %#v", schema["required"])
	}
}

func TestSchemaEnumByteExact(t *testing.T) {
	form := fieldspec.Form{Fields: map[string]fieldspec.Field{
		"AUTHORITY": {Type: "enum", Options: []string{"read-only", "design-only", "report-only"}},
	}}
	schema := SchemaFromForm(form, "digest-4")
	headers := schema["properties"].(map[string]any)["headers"].(map[string]any)["properties"].(map[string]any)
	got := headers["AUTHORITY"].(map[string]any)["enum"]
	if !reflect.DeepEqual(got, []string{"read-only", "design-only", "report-only"}) {
		t.Fatalf("enum = %#v", got)
	}
}
