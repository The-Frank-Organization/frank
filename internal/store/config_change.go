package store

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/The-Frank-Organization/frank/internal/record"
)

const IntentConfig = "config"

func ConfigChangeIntents(rec record.Record) []Intent {
	intents, _ := ConfigChangeIntentsStrict(rec)
	return intents
}

func ConfigChangeIntentsStrict(rec record.Record) ([]Intent, error) {
	if rec.Headers["member"] == "adoption" {
		return adoptionConfigChangeIntents(rec)
	}
	target, err := configTarget(rec.Headers["member"])
	if err != nil {
		return nil, err
	}
	relayID := rec.Envelope.RelayID
	toList, ccList, err := decodedHeaderRecipients(rec)
	if err != nil {
		return nil, err
	}
	toDisplay := joinRecipients(toList)
	ccDisplay := joinRecipients(ccList)
	if toDisplay == "" && !hasAddressListHeader(rec, "TO") {
		toDisplay = rec.Envelope.To
	}
	intents := []Intent{
		{Kind: IntentIndex, Path: "INDEX.md", Payload: []byte(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n", relayID, rec.Headers["PHASE"], rec.Envelope.From, toDisplay, ccDisplay, rec.Envelope.DeliveryState))},
		{Kind: IntentConfig, Path: target, Payload: []byte(rec.Body)},
	}
	if rec.Envelope.DeliveryState == record.Accepted {
		recipients, err := DeliveryRecipients(rec)
		if err != nil {
			return nil, err
		}
		for _, recipient := range recipients {
			intents = append(intents, Intent{Kind: IntentMailbox, Path: safeMailbox(recipient) + ".jsonl", Payload: []byte(relayID + "\n")})
		}
	}
	return intents, nil
}

type adoptionBody struct {
	Members []adoptionMember `json:"members"`
}

type adoptionMember struct {
	Name     string `json:"name"`
	BytesB64 string `json:"bytes_b64"`
}

func adoptionConfigChangeIntents(rec record.Record) ([]Intent, error) {
	if err := rejectDuplicateJSONKeys([]byte(rec.Body)); err != nil {
		return nil, fmt.Errorf("invalid adoption body: %w", err)
	}
	dec := json.NewDecoder(bytes.NewBufferString(rec.Body))
	dec.DisallowUnknownFields()
	var body adoptionBody
	if err := dec.Decode(&body); err != nil {
		return nil, fmt.Errorf("invalid adoption body: %w", err)
	}
	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return nil, fmt.Errorf("invalid adoption body: trailing JSON")
	}
	if len(body.Members) != 2 || body.Members[0].Name != "catalog" || body.Members[1].Name != "engine" {
		return nil, fmt.Errorf("invalid adoption members: want catalog,engine")
	}

	intents := []Intent{{
		Kind: IntentIndex,
		Path: "INDEX.md",
		Payload: []byte(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
			rec.Envelope.RelayID, rec.Headers["PHASE"], rec.Envelope.From, rec.Envelope.To, "", rec.Envelope.DeliveryState)),
	}}
	for _, member := range body.Members {
		decoded, err := base64.StdEncoding.Strict().DecodeString(member.BytesB64)
		if err != nil || base64.StdEncoding.EncodeToString(decoded) != member.BytesB64 {
			return nil, fmt.Errorf("invalid adoption bytes for %s: canonical base64 required", member.Name)
		}
		target, err := configTarget(member.Name)
		if err != nil {
			return nil, err
		}
		intents = append(intents, Intent{Kind: IntentConfig, Path: target, Payload: decoded})
	}
	return intents, nil
}

func rejectDuplicateJSONKeys(data []byte) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	var walk func(json.Token) error
	walk = func(token json.Token) error {
		delim, ok := token.(json.Delim)
		if !ok {
			return nil
		}
		switch delim {
		case '{':
			seen := map[string]bool{}
			for dec.More() {
				keyToken, err := dec.Token()
				if err != nil {
					return err
				}
				key, ok := keyToken.(string)
				if !ok {
					return errors.New("object key must be a string")
				}
				if seen[key] {
					return fmt.Errorf("duplicate JSON key %q", key)
				}
				seen[key] = true
				value, err := dec.Token()
				if err != nil {
					return err
				}
				if err := walk(value); err != nil {
					return err
				}
			}
			_, err := dec.Token()
			return err
		case '[':
			for dec.More() {
				value, err := dec.Token()
				if err != nil {
					return err
				}
				if err := walk(value); err != nil {
					return err
				}
			}
			_, err := dec.Token()
			return err
		default:
			return fmt.Errorf("unexpected JSON delimiter %q", delim)
		}
	}
	first, err := dec.Token()
	if err != nil {
		return err
	}
	if err := walk(first); err != nil {
		return err
	}
	if _, err := dec.Token(); err != io.EOF {
		if err == nil {
			return errors.New("trailing JSON")
		}
		return err
	}
	return nil
}
