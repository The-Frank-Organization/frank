//go:build frank_test_reduced_limits

package scheduler

import (
	"context"
	"testing"

	"github.com/jackli/frank/internal/appctl/store"
	"github.com/jackli/frank/internal/appipc"
)

func TestReducedRuntimeOverflowBoundaries(t *testing.T) {
	exact := newFixture(t, "reduced-exact")
	exact.mutate(func(ctx context.Context, tx *store.Tx) error {
		if _, err := tx.ExecContext(ctx, `DELETE FROM turns WHERE turn_id=?`, exact.turnID); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `UPDATE runs SET run_phase='created' WHERE run_id=?`, exact.runID)
		return err
	})
	request := exact.request("exact", appipc.FrameMax, false)
	if result, err := exact.scheduler.Admit(exact.ctx, request); err != nil || result.Decision != AdmissionCommitted {
		t.Fatalf("exact fit = %#v err=%v", result, err)
	}

	encodedReserve := newFixture(t, "reduced-reserve")
	encodedReserve.mutate(func(ctx context.Context, tx *store.Tx) error {
		if _, err := tx.ExecContext(ctx, `DELETE FROM turns WHERE turn_id=?`, encodedReserve.turnID); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `UPDATE runs SET run_phase='created' WHERE run_id=?`, encodedReserve.runID)
		return err
	})
	request = encodedReserve.request("reserve", 100, false)
	request.AdmissionRefEncodedSize = appipc.AdmissionRefEncMax + 1
	if result, err := encodedReserve.scheduler.Admit(encodedReserve.ctx, request); err != nil || result.Decision != AdmissionTaskOverflow || encodedReserve.count(`SELECT COUNT(*) FROM turns WHERE turn_id='reserve'`, nil) != 0 {
		t.Fatalf("encoded reserve = %#v err=%v", result, err)
	}
}
