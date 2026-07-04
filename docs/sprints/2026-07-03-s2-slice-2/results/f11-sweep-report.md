# S2 F11 Sweep Report

Evidence level: E2 fixture artifact.

This report records the reviewed S2 class x crashpoint applicability matrix. The automated fixture asserts every class has a `crash-expected` or `clean-completion` cell for every registered crashpoint.

Executed crash-expected cells: the child-process F11 fixture drives the S1 mutation boundaries plus the new S2 crashpoints for quarantine disposition, segmented rotation, recovery phase boundaries, GC marker/unlink, owed item, and owed disposition. Each crash leg uses `FRANK_TEST_CRASHPOINT`, expects SIGKILL, reruns recovery/convergence, and verifies committed records.

Hit-trace row equality: `TestS2ApplicabilityMapRowsMatchHitTrace` runs one child process per class with `FRANK_TEST_HIT_TRACE=<file>`. `crashpoint.Hit` appends every fired name before crashpoint target filtering; the test asserts the trace-observed fired set equals the map row's `crash-expected` set exactly. A fired point labeled `clean-completion` fails; a `crash-expected` point that does not fire fails.

Clean-completion evidence: every point not in the trace-observed fired set is a clean-completion cell for that class. The legacy class execution smoke still runs submit-accept, submit-reject, held, operator-verdict, park, outbox-enqueue, genesis, quarantine-disposition, gc-marker, owed-item, and owed-disposition to verify those classes exit normally and converge with no crashpoint armed.

Classes covered: submit-accept, submit-reject, held, operator-verdict, park, outbox-enqueue, genesis, quarantine-disposition, gc-marker, owed-item, owed-disposition.

New S2 crashpoints executed: pre_quarantine_evict, post_quarantine_evict, pre_segment_rotate, post_segment_rotate, recovery_post_phase0, recovery_post_phase1, recovery_post_phase2, recovery_post_phase3, recovery_post_phase3_5, recovery_post_phase3_6, recovery_post_phase4, pre_gc_marker, post_gc_marker, pre_gc_unlink, post_gc_unlink.

O3 fixture proof: `TestFrankBinaryOperatorChannelO3OwedSweepOpenAndDisposition` submits the real `OI-S1-F11-SWEEP` owed-item payload through the authenticated operator channel in a fixture store (`owner: s1 (S2 slice)`, `source` citing `docs/sprints/2026-07-03-s1-slice-1/RECONCILE.md:160-161` and `master/relays/s1-exit-gate/SITREP-planner-20260703-200827.md`, `target_surface: F11 full class×point sweep`, `disposition_path: S2 exit gate`). It verifies the open item through channel `project`, channel `read`, and `projections/owed/OPEN.md`; then it submits an `owed_disposition` citing this report and verifies the open set empties. The real Task-13.5 operator submit remains operator-executed; this fixture proves the semantics only.
