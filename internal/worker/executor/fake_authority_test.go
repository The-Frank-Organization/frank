package executor

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type fakeTicket struct {
	identity FullIdentity
	state    TicketState
	outcome  OutcomeRecord
}

type fakeAuthority struct {
	mu                sync.Mutex
	nextTicket        int
	tickets           map[string]*fakeTicket
	outcomes          []OutcomeRecord
	authorizeCalls    int
	consumeCalls      int
	authorizeOverride *AuthorizeReply
	consumeErr        error
	recordErrBefore   error
	recordErrAfter    error
	afterAuthorize    func()
	afterConsume      func()
}

func newFakeAuthority() *fakeAuthority {
	return &fakeAuthority{tickets: make(map[string]*fakeTicket)}
}

func (authority *fakeAuthority) Authorize(_ context.Context, request AuthorizeRequest) (AuthorizeReply, error) {
	authority.mu.Lock()
	authority.authorizeCalls++
	if authority.authorizeOverride != nil {
		reply := *authority.authorizeOverride
		authority.mu.Unlock()
		return reply, nil
	}
	authority.nextTicket++
	ticketID := fmt.Sprintf("ticket-%d", authority.nextTicket)
	authority.tickets[ticketID] = &fakeTicket{identity: request.Identity, state: TicketIssued}
	hook := authority.afterAuthorize
	authority.mu.Unlock()
	if hook != nil {
		hook()
	}
	return AuthorizeReply{Code: AuthorizeGranted, TicketID: ticketID}, nil
}

func (authority *fakeAuthority) Consume(_ context.Context, request ConsumeRequest) (ConsumeReply, error) {
	authority.mu.Lock()
	authority.consumeCalls++
	if authority.consumeErr != nil {
		err := authority.consumeErr
		authority.mu.Unlock()
		return ConsumeReply{}, err
	}
	ticket := authority.tickets[request.TicketID]
	if ticket == nil {
		authority.mu.Unlock()
		return ConsumeReply{}, errors.New("unknown ticket: channel fault")
	}
	var reply ConsumeReply
	switch {
	case request.TurnEpoch < ticket.identity.TurnEpoch:
		reply.Code = ConsumeStaleEpoch
	case request.TurnEpoch > ticket.identity.TurnEpoch:
		authority.mu.Unlock()
		return ConsumeReply{}, errors.New("future epoch: channel fault")
	case ticket.state != TicketIssued:
		reply.Code = ConsumeDuplicate
	case request.TurnEpoch != ticket.identity.TurnEpoch || request.CanonicalToolName != ticket.identity.CanonicalToolName || request.CanonicalArgsDigest != ticket.identity.CanonicalArgsDigest:
		reply.Code = ConsumeIdentityMismatch
	default:
		ticket.state = TicketConsumed
		reply.Code = ConsumeOK
	}
	hook := authority.afterConsume
	authority.mu.Unlock()
	if reply.Code == ConsumeOK && hook != nil {
		hook()
	}
	return reply, nil
}

func (authority *fakeAuthority) RecordOutcome(_ context.Context, record OutcomeRecord) error {
	if err := record.Validate(); err != nil {
		return err
	}
	authority.mu.Lock()
	defer authority.mu.Unlock()
	if authority.recordErrBefore != nil {
		return authority.recordErrBefore
	}
	ticket := authority.tickets[record.TicketID]
	if ticket == nil {
		return errors.New("unknown ticket")
	}
	if ticket.state != TicketConsumed {
		return fmt.Errorf("ticket state %s cannot record outcome", ticket.state)
	}
	if record.Outcome == OutcomeExecuted && *record.InvocationIdentity != ticket.identity.Identity {
		return errors.New("invocation identity does not match ticket")
	}
	if record.Outcome == OutcomeNotInvokedIntegrityFault && record.IntegrityEvidence.Expected != ticket.identity.Identity {
		return errors.New("expected identity does not match ticket")
	}
	ticket.state = TicketOutcomeRecorded
	ticket.outcome = record
	authority.outcomes = append(authority.outcomes, record)
	return authority.recordErrAfter
}

func (authority *fakeAuthority) state(ticketID string) TicketState {
	authority.mu.Lock()
	defer authority.mu.Unlock()
	if ticket := authority.tickets[ticketID]; ticket != nil {
		return ticket.state
	}
	return ""
}

func (authority *fakeAuthority) outcome(ticketID string) OutcomeRecord {
	authority.mu.Lock()
	defer authority.mu.Unlock()
	return authority.tickets[ticketID].outcome
}

func (authority *fakeAuthority) retireEpoch(epoch uint64) {
	authority.mu.Lock()
	defer authority.mu.Unlock()
	for _, ticket := range authority.tickets {
		if ticket.identity.TurnEpoch != epoch {
			continue
		}
		switch ticket.state {
		case TicketIssued:
			ticket.state = TicketVoid
		case TicketConsumed:
			ticket.state = TicketUnknownToolOutcome
		}
	}
}
