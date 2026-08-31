package main

import (
	"encoding/json"

	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/seatclient/formschema"
)

// These compatibility wrappers keep the MCP package a thin consumer while the
// m-2 mapping implementation lives in its shared module.
func SchemaFromForm(form fieldspec.Form, digest string) map[string]any {
	return formschema.SchemaFromForm(form, digest)
}

func SubmitPayloadFromArguments(arguments json.RawMessage) (json.RawMessage, error) {
	return formschema.SubmitPayloadFromArguments(arguments)
}

type submitArguments = formschema.SubmitArguments
