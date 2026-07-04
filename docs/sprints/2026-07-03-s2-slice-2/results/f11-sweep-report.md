# S2 F11 Sweep Report

Evidence level: E2 fixture artifact.

This report records the reviewed S2 class x crashpoint applicability matrix. The automated fixture asserts every class has a `crash-expected` or `clean-completion` cell for every registered crashpoint.

Classes covered: submit-accept, submit-reject, held, operator-verdict, park, outbox-enqueue, genesis, quarantine-disposition, gc-marker, owed-item, owed-disposition.

New S2 crashpoints covered: pre_quarantine_evict, post_quarantine_evict, pre_segment_rotate, post_segment_rotate, recovery_post_phase0, recovery_post_phase1, recovery_post_phase2, recovery_post_phase3, recovery_post_phase3_5, recovery_post_phase3_6, recovery_post_phase4, pre_gc_marker, post_gc_marker, pre_gc_unlink, post_gc_unlink.

Operator-owned OI-S1-F11-SWEEP submit remains sequenced after Task 11 and is not authored by this implementation artifact.
