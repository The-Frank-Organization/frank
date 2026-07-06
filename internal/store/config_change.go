package store

import (
	"fmt"

	"github.com/jackli/frank/internal/record"
)

const IntentConfig = "config"

func ConfigChangeIntents(rec record.Record) []Intent {
	target, err := configTarget(rec.Headers["member"])
	if err != nil {
		return nil
	}
	relayID := rec.Envelope.RelayID
	intents := []Intent{
		{Kind: IntentIndex, Path: "INDEX.md", Payload: []byte(fmt.Sprintf("| %s | %s | %s | %s | %s |\n", relayID, rec.Headers["PHASE"], rec.Envelope.From, rec.Envelope.To, rec.Envelope.DeliveryState))},
		{Kind: IntentConfig, Path: target, Payload: []byte(rec.Body)},
	}
	for _, recipient := range DeliveryRecipients(rec) {
		intents = append(intents, Intent{Kind: IntentMailbox, Path: safeMailbox(recipient) + ".jsonl", Payload: []byte(relayID + "\n")})
	}
	return intents
}
