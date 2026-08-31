package outcome

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/connector/authorize"
	"github.com/The-Frank-Organization/frank/internal/connector/request"
)

const (
	testB = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	testE = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
)

func TestTotalAttemptResultDispositionMapping(t *testing.T) {
	tests := []struct {
		outcome Outcome
		want    string
	}{
		{outcome: Outcome{Kind: RejectedLocal, RejectReason: request.MalformedRequest, RefusalStage: PreFreeze}, want: "rejected_local(malformed_request)"},
		{outcome: Outcome{Kind: Denied, DenyReason: authorize.EndpointMismatch, Digests: testDigests()}, want: "denied(endpoint-mismatch)"},
		{outcome: Outcome{Kind: SentCompleted, Digests: testDigests()}, want: "sent_completed"},
		{outcome: Outcome{Kind: TransportFailed, Digests: testDigests()}, want: "transport_failed"},
		{outcome: Outcome{Kind: Unknown, Digests: testDigests()}, want: "unknown"},
		{outcome: Outcome{Kind: Cancelled, CancelPoint: PostInvocation, Digests: testDigests()}, want: "cancelled(post_invocation)"},
	}
	for _, test := range tests {
		result, err := AttemptResult("attempt-1", 7, test.outcome)
		if err != nil {
			t.Fatalf("AttemptResult(%s): %v", test.outcome.Kind, err)
		}
		if result.Schema != AttemptResultSchemaV2 || result.Disposition != test.want {
			t.Fatalf("AttemptResult(%s) = %+v", test.outcome.Kind, result)
		}
	}
}

func TestRejectStageClassifiesWithoutDigestAndValidatesPresence(t *testing.T) {
	for _, test := range []struct {
		reason request.RejectReason
		stage  RefusalStage
		want   RejectCut
	}{
		{request.MalformedRequest, PreFreeze, RejectPreFreeze},
		{request.LaneCapabilityMismatch, PreFreeze, RejectPreFreeze},
		{request.ReplayScopeViolation, PreFreeze, RejectPreFreeze},
		{request.InternalIntegrityFault, PreFreeze, RejectPreFreeze},
		{request.InternalIntegrityFault, PostFreeze, RejectPostFreeze},
	} {
		got, err := ClassifyReject(test.reason, test.stage)
		if err != nil || got != test.want {
			t.Fatalf("ClassifyReject(%s,%s) = %s, %v", test.reason, test.stage, got, err)
		}
	}
	if _, err := ClassifyReject(request.MalformedRequest, PostFreeze); err == nil {
		t.Fatal("post-freeze malformed_request classified")
	}

	invalid := []Outcome{
		{Kind: RejectedLocal, RejectReason: request.InternalIntegrityFault, RefusalStage: PostFreeze},
		{Kind: RejectedLocal, RejectReason: request.InternalIntegrityFault, RefusalStage: PreFreeze, Digests: testDigests()},
		{Kind: RejectedLocal, RejectReason: request.InternalIntegrityFault, RefusalStage: PostFreeze, Digests: &Digests{FrozenCore: testB}},
	}
	for _, value := range invalid {
		if _, err := AttemptResult("attempt-1", 7, value); err == nil {
			t.Fatalf("invalid digest/stage accepted: %+v", value)
		}
	}
}

func TestReplyMatrixAndEpochNoResult(t *testing.T) {
	pre := Outcome{Kind: RejectedLocal, RejectReason: request.MalformedRequest, RefusalStage: PreFreeze}
	preReply, err := DataReply("attempt-1", 7, pre)
	if err != nil || preReply.Kind != "rejected_local" || preReply.RefusalStage != PreFreeze || preReply.FrozenCoreDigest != "" || preReply.ProviderLoweredToolsDigest != "" {
		t.Fatalf("pre-freeze reply = %+v, %v", preReply, err)
	}
	post := Outcome{Kind: RejectedLocal, RejectReason: request.InternalIntegrityFault, RefusalStage: PostFreeze, Digests: testDigests()}
	postReply, err := DataReply("attempt-1", 7, post)
	if err != nil || postReply.Kind != "internal_integrity_reject" || postReply.FrozenCoreDigest != testB || postReply.ProviderLoweredToolsDigest != testE {
		t.Fatalf("post-freeze reply = %+v, %v", postReply, err)
	}
	deny := Outcome{Kind: Denied, DenyReason: authorize.PolicyUnavailable, Digests: testDigests()}
	denyReply, err := DataReply("attempt-1", 7, deny)
	if err != nil || denyReply.Kind != "egress_denied" || denyReply.DenyReason != authorize.PolicyUnavailable || denyReply.FrozenCoreDigest != testB || denyReply.ProviderLoweredToolsDigest != testE {
		t.Fatalf("deny reply = %+v, %v", denyReply, err)
	}
	for _, fence := range []string{"STALE_EPOCH", "EPOCH_AHEAD"} {
		reply, result := EpochDataReply("attempt-1", 7, fence)
		if reply.Kind != fence || reply.FrozenCoreDigest != "" || result != nil {
			t.Fatalf("epoch reply/result = %+v/%+v", reply, result)
		}
	}
}

func TestLocalRejectWritesControlBeforeData(t *testing.T) {
	writer := &orderingWriter{}
	outcome := Outcome{Kind: RejectedLocal, RejectReason: request.InternalIntegrityFault, RefusalStage: PostFreeze, Digests: testDigests()}
	if err := EmitLocalReject(writer, "attempt-1", 7, outcome); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(writer.order, []string{"control", "data"}) {
		t.Fatalf("write order = %v", writer.order)
	}

	mutation := &orderingWriter{}
	reply, _ := DataReply("attempt-1", 7, outcome)
	result, _ := AttemptResult("attempt-1", 7, outcome)
	if err := mutation.SendDataReply(reply); err == nil {
		t.Fatal("reversed-order mutation did not fail at the DATA-P barrier")
	}
	if len(mutation.order) != 0 {
		t.Fatalf("reversed mutation completed a write: %v", mutation.order)
	}
	_ = result
}

func TestCarriersContainOnlyPayloadFreeIdentityFields(t *testing.T) {
	result, err := AttemptResult("attempt-secret-free", 7, Outcome{Kind: Denied, DenyReason: authorize.MethodMismatch, Digests: testDigests()})
	if err != nil {
		t.Fatal(err)
	}
	raw := result.CanonicalBytes()
	for _, forbidden := range []string{"prompt", "SENTINEL_SECRET", "request_write_completed"} {
		if strings.Contains(string(raw), forbidden) {
			t.Fatalf("result contains forbidden payload fact %q: %s", forbidden, raw)
		}
	}
}

func testDigests() *Digests {
	return &Digests{FrozenCore: testB, ProviderLoweredTools: testE}
}

type orderingWriter struct {
	order []string
}

func (writer *orderingWriter) SendAttemptResult(AttemptResultV2) error {
	writer.order = append(writer.order, "control")
	return nil
}

func (writer *orderingWriter) SendDataReply(DataReplyV2) error {
	if len(writer.order) == 0 || writer.order[len(writer.order)-1] != "control" {
		return errors.New("DATA-P reply released before CTRL-C result")
	}
	writer.order = append(writer.order, "data")
	return nil
}
