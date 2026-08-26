package supervisor

import (
	"context"
	"fmt"
	"testing"
)

func TestBrokerFailureClassesIncrementExactlyOnceWithoutEpochCascade(t *testing.T) {
	fixture := newSupervisorFixture(t, "broker-failures")
	classes := []BrokerFailureClass{BrokerSpawnFail, BrokerNoReady, BrokerMalformedReady, BrokerReadyCrash, BrokerReattachDeadline}
	for index, class := range classes {
		request := BrokerFailureRequest{RunID: fixture.runID, InstanceID: fmt.Sprintf("instance-%d", index), Class: class, At: int64(index + 10)}
		result, err := fixture.controller.RecordBrokerFailure(fixture.ctx, request)
		if err != nil || result.Failures != uint64(index+1) || result.Idempotent {
			t.Fatalf("%s result=%#v err=%v", class, result, err)
		}
		replay, err := fixture.controller.RecordBrokerFailure(fixture.ctx, request)
		if err != nil || !replay.Idempotent || replay.Failures != uint64(index+1) {
			t.Fatalf("%s replay=%#v err=%v", class, replay, err)
		}
	}
	fixture.assertScalar(`SELECT consecutive_failures FROM runs WHERE run_id=?`, fixture.runID, storeCounter(5))
	fixture.assertScalar(`SELECT turn_epoch FROM epochs WHERE run_id=?`, fixture.runID, storeCounter(1))
	fixture.assertScalar(`SELECT state FROM workers WHERE generation_id=?`, fixture.generationID, "LEASED")
}

func TestTenthBrokerFailureIsSharedTerminalAndResetDoesNotForgetIdentity(t *testing.T) {
	fixture := newSupervisorFixture(t, "broker-terminal")
	for index := 0; index < MaxConsecutiveFailures; index++ {
		result, err := fixture.controller.RecordBrokerFailure(fixture.ctx, BrokerFailureRequest{RunID: fixture.runID, InstanceID: fmt.Sprintf("terminal-%d", index), Class: BrokerSpawnFail, At: int64(index + 1)})
		if err != nil || result.Terminal != (index == MaxConsecutiveFailures-1) {
			t.Fatalf("failure %d result=%#v err=%v", index+1, result, err)
		}
	}
	fixture.assertScalar(`SELECT state FROM runs WHERE run_id=?`, fixture.runID, "FAILED")
	fixture.assertScalar(`SELECT turn_epoch FROM epochs WHERE run_id=?`, fixture.runID, storeCounter(1))

	reset := newSupervisorFixture(t, "broker-reset")
	request := BrokerFailureRequest{RunID: reset.runID, InstanceID: "same", Class: BrokerReadyCrash, At: 1}
	if _, err := reset.controller.RecordBrokerFailure(reset.ctx, request); err != nil {
		t.Fatal(err)
	}
	if err := reset.controller.ResetFailures(context.Background(), reset.runID, 2); err != nil {
		t.Fatal(err)
	}
	replay, err := reset.controller.RecordBrokerFailure(reset.ctx, request)
	if err != nil || !replay.Idempotent || replay.Failures != 0 {
		t.Fatalf("post-reset replay=%#v err=%v", replay, err)
	}
}
