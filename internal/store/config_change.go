package store

import (
	"fmt"

	"github.com/jackli/frank/internal/record"
)

const IntentConfig = "config"

func ConfigChangeIntents(rec record.Record) []Intent {
	intents, _ := ConfigChangeIntentsStrict(rec)
	return intents
}

func ConfigChangeIntentsStrict(rec record.Record) ([]Intent, error) {
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
	recipients, err := DeliveryRecipients(rec)
	if err != nil {
		return nil, err
	}
	for _, recipient := range recipients {
		intents = append(intents, Intent{Kind: IntentMailbox, Path: safeMailbox(recipient) + ".jsonl", Payload: []byte(relayID + "\n")})
	}
	return intents, nil
}
