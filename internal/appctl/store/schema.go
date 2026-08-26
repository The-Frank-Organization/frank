package store

var schemaV1 = []string{
	`CREATE TABLE IF NOT EXISTS runs (
		run_id TEXT PRIMARY KEY,
		manifest_bytes BLOB NOT NULL,
		run_manifest_digest TEXT NOT NULL CHECK(length(run_manifest_digest)=64 AND run_manifest_digest NOT GLOB '*[^0-9a-f]*'),
		state TEXT NOT NULL CHECK(state IN ('ADMITTED','ACTIVE','COMPLETED','FAILED','CANCELLED','INTERRUPTED')),
		stop_reason TEXT CHECK(stop_reason IN ('resume_frame_overflow','parked_unknown_capacity_exceeded')),
		resume_action TEXT CHECK(resume_action='operator_new_run'),
		run_phase TEXT NOT NULL CHECK(run_phase IN ('created','create_authorized','established')),
		session_log_path TEXT NOT NULL DEFAULT '',
		consecutive_failures TEXT NOT NULL CHECK(length(consecutive_failures)=20 AND consecutive_failures NOT GLOB '*[^0-9]*' AND consecutive_failures <= '18446744073709551615'),
		backoff_until INTEGER,
		created_at INTEGER NOT NULL,
		updated_at INTEGER,
		CHECK((stop_reason IS NULL AND resume_action IS NULL) OR
			(state='FAILED' AND stop_reason='resume_frame_overflow' AND resume_action='operator_new_run') OR
			(state='FAILED' AND stop_reason='parked_unknown_capacity_exceeded' AND resume_action IS NULL))
	)`,
	`CREATE TABLE IF NOT EXISTS workers (
		generation_id TEXT PRIMARY KEY,
		run_id TEXT NOT NULL REFERENCES runs(run_id),
		turn_epoch TEXT NOT NULL CHECK(length(turn_epoch)=20 AND turn_epoch NOT GLOB '*[^0-9]*' AND turn_epoch <= '18446744073709551615'),
		pid INTEGER,
		state TEXT NOT NULL CHECK(state IN ('ALLOCATED','SPAWNING','READY','LEASED','RETIRING','FAILED','TERMINATED')),
		attach_result TEXT CHECK(attach_result IN ('attach-ok','broker:attach-suspended','broker:attach-tuple-mismatch')),
		spawn_evidence TEXT,
		reap_evidence TEXT,
		created_at INTEGER NOT NULL,
		updated_at INTEGER
	)`,
	`CREATE TABLE IF NOT EXISTS turns (
		turn_id TEXT PRIMARY KEY,
		run_id TEXT NOT NULL REFERENCES runs(run_id),
		turn_epoch TEXT NOT NULL CHECK(length(turn_epoch)=20 AND turn_epoch NOT GLOB '*[^0-9]*' AND turn_epoch <= '18446744073709551615'),
		state TEXT NOT NULL CHECK(state IN ('ADMITTED','ACTIVE','COMPLETED','FAILED','CANCELLED','INTERRUPTED')),
		admission_ref BLOB NOT NULL,
		run_disposition TEXT NOT NULL CHECK(run_disposition IN ('fresh','resume')),
		create_auth_id TEXT NOT NULL CHECK(length(create_auth_id)=32 AND create_auth_id NOT GLOB '*[^0-9a-f]*'),
		predecessor_turn_id TEXT REFERENCES turns(turn_id),
		resume_snapshot BLOB,
		resume_disposition TEXT NOT NULL CHECK(resume_disposition IN ('PENDING','RESUMABLE','DEGRADED')),
		resume_action TEXT CHECK(resume_action='re_derive'),
		created_at INTEGER NOT NULL,
		updated_at INTEGER,
		UNIQUE(run_id, predecessor_turn_id),
		CHECK((predecessor_turn_id IS NULL AND run_disposition='fresh') OR (predecessor_turn_id IS NOT NULL AND run_disposition='resume')),
		CHECK((resume_disposition='DEGRADED' AND resume_action='re_derive') OR (resume_disposition IN ('PENDING','RESUMABLE') AND resume_action IS NULL))
	)`,
	`CREATE UNIQUE INDEX IF NOT EXISTS one_initial_turn_per_run ON turns(run_id) WHERE predecessor_turn_id IS NULL`,
	`CREATE TABLE IF NOT EXISTS provider_attempts (
		attempt_id TEXT PRIMARY KEY,
		run_id TEXT NOT NULL REFERENCES runs(run_id),
		turn_id TEXT NOT NULL REFERENCES turns(turn_id),
		turn_epoch TEXT NOT NULL CHECK(length(turn_epoch)=20 AND turn_epoch NOT GLOB '*[^0-9]*' AND turn_epoch <= '18446744073709551615'),
		state TEXT NOT NULL CHECK(state IN ('OPEN','STREAMING','COMPLETED','REJECTED_LOCAL','CANCELLED','UNKNOWN_PROVIDER_OUTCOME','PARTIAL_STREAM')),
		logical_surface_digest TEXT NOT NULL CHECK(length(logical_surface_digest)=64 AND logical_surface_digest NOT GLOB '*[^0-9a-f]*'),
		frozen_core_digest TEXT CHECK(frozen_core_digest IS NULL OR (length(frozen_core_digest)=64 AND frozen_core_digest NOT GLOB '*[^0-9a-f]*')),
		provider_lowered_tools_digest TEXT CHECK(provider_lowered_tools_digest IS NULL OR (length(provider_lowered_tools_digest)=64 AND provider_lowered_tools_digest NOT GLOB '*[^0-9a-f]*')),
		cancel_point TEXT CHECK(cancel_point IN ('pre_transport','post_invocation')),
		refusal_stage TEXT CHECK(refusal_stage IN ('pre_freeze','post_freeze')),
		cancellation_id TEXT REFERENCES cancellations(cancellation_id),
		created_at INTEGER NOT NULL,
		updated_at INTEGER,
		CHECK((state='CANCELLED' AND cancel_point IS NOT NULL AND cancellation_id IS NOT NULL) OR (state<>'CANCELLED' AND cancel_point IS NULL AND cancellation_id IS NULL))
	)`,
	`CREATE TABLE IF NOT EXISTS tool_calls (
		tool_call_id TEXT PRIMARY KEY,
		run_id TEXT NOT NULL REFERENCES runs(run_id),
		turn_id TEXT NOT NULL REFERENCES turns(turn_id),
		turn_epoch TEXT NOT NULL CHECK(length(turn_epoch)=20 AND turn_epoch NOT GLOB '*[^0-9]*' AND turn_epoch <= '18446744073709551615'),
		state TEXT NOT NULL CHECK(state IN ('REQUESTED','AUTHORIZED','CONSUMED','EXECUTED','NOT_INVOKED_INTEGRITY_FAULT','UNKNOWN_TOOL_OUTCOME','PARTIAL_TOOL_EFFECT')),
		canonical_tool_name TEXT NOT NULL,
		canonical_args_digest TEXT NOT NULL CHECK(length(canonical_args_digest)=64 AND canonical_args_digest NOT GLOB '*[^0-9a-f]*'),
		invocation_tool_name TEXT,
		invocation_args_digest TEXT CHECK(invocation_args_digest IS NULL OR (length(invocation_args_digest)=64 AND invocation_args_digest NOT GLOB '*[^0-9a-f]*')),
		expected_tool_name TEXT,
		expected_args_digest TEXT CHECK(expected_args_digest IS NULL OR (length(expected_args_digest)=64 AND expected_args_digest NOT GLOB '*[^0-9a-f]*')),
		observed_tool_name TEXT,
		observed_args_digest TEXT CHECK(observed_args_digest IS NULL OR (length(observed_args_digest)=64 AND observed_args_digest NOT GLOB '*[^0-9a-f]*')),
		expected_turn_epoch TEXT CHECK(expected_turn_epoch IS NULL OR (length(expected_turn_epoch)=20 AND expected_turn_epoch NOT GLOB '*[^0-9]*' AND expected_turn_epoch <= '18446744073709551615')),
		observed_turn_epoch TEXT CHECK(observed_turn_epoch IS NULL OR (length(observed_turn_epoch)=20 AND observed_turn_epoch NOT GLOB '*[^0-9]*' AND observed_turn_epoch <= '18446744073709551615')),
		created_at INTEGER NOT NULL,
		updated_at INTEGER,
		CHECK((state='EXECUTED' AND invocation_tool_name IS NOT NULL AND invocation_args_digest IS NOT NULL AND
			expected_tool_name IS NULL AND expected_args_digest IS NULL AND expected_turn_epoch IS NULL AND
			observed_tool_name IS NULL AND observed_args_digest IS NULL AND observed_turn_epoch IS NULL) OR
			(state='NOT_INVOKED_INTEGRITY_FAULT' AND invocation_tool_name IS NULL AND invocation_args_digest IS NULL AND
			expected_tool_name IS NOT NULL AND expected_args_digest IS NOT NULL AND expected_turn_epoch IS NOT NULL AND
			observed_tool_name IS NOT NULL AND observed_args_digest IS NOT NULL AND observed_turn_epoch IS NOT NULL) OR
			(state NOT IN ('EXECUTED','NOT_INVOKED_INTEGRITY_FAULT') AND invocation_tool_name IS NULL AND invocation_args_digest IS NULL AND
			expected_tool_name IS NULL AND expected_args_digest IS NULL AND expected_turn_epoch IS NULL AND
			observed_tool_name IS NULL AND observed_args_digest IS NULL AND observed_turn_epoch IS NULL))
	)`,
	`CREATE TABLE IF NOT EXISTS tool_authorizations (
		ticket_id TEXT PRIMARY KEY,
		run_id TEXT NOT NULL REFERENCES runs(run_id),
		turn_id TEXT NOT NULL REFERENCES turns(turn_id),
		tool_call_id TEXT NOT NULL,
		turn_epoch TEXT NOT NULL CHECK(length(turn_epoch)=20 AND turn_epoch NOT GLOB '*[^0-9]*' AND turn_epoch <= '18446744073709551615'),
		state TEXT NOT NULL CHECK(state IN ('ISSUED','CONSUMED','OUTCOME_RECORDED','VOID','UNKNOWN_TOOL_OUTCOME')),
		void_reason TEXT CHECK(void_reason IN ('run_not_admitted','turn_inactive','lease_invalid','denied_above_set','expired')),
		canonical_tool_name TEXT NOT NULL,
		canonical_args_digest TEXT NOT NULL CHECK(length(canonical_args_digest)=64 AND canonical_args_digest NOT GLOB '*[^0-9a-f]*'),
		canonical_resource TEXT,
		cwd TEXT,
		effect_descriptor BLOB NOT NULL,
		issued_at INTEGER NOT NULL,
		consumed_at INTEGER,
		outcome_ref TEXT,
		UNIQUE(run_id, turn_id, tool_call_id),
		CHECK((state='VOID' AND void_reason IS NOT NULL) OR (state<>'VOID' AND void_reason IS NULL)),
		CHECK((state='OUTCOME_RECORDED' AND outcome_ref IS NOT NULL) OR (state<>'OUTCOME_RECORDED' AND outcome_ref IS NULL))
	)`,
	`CREATE TABLE IF NOT EXISTS epochs (
		run_id TEXT PRIMARY KEY REFERENCES runs(run_id),
		turn_epoch TEXT NOT NULL CHECK(length(turn_epoch)=20 AND turn_epoch NOT GLOB '*[^0-9]*' AND turn_epoch <= '18446744073709551615'),
		state_seq TEXT NOT NULL CHECK(length(state_seq)=20 AND state_seq NOT GLOB '*[^0-9]*' AND state_seq <= '18446744073709551615')
	)`,
	`CREATE TABLE IF NOT EXISTS leases (
		run_id TEXT NOT NULL REFERENCES runs(run_id),
		lease_kind TEXT NOT NULL CHECK(lease_kind IN ('worker','turn')),
		lease_id TEXT NOT NULL,
		generation_id TEXT REFERENCES workers(generation_id),
		turn_epoch TEXT NOT NULL CHECK(length(turn_epoch)=20 AND turn_epoch NOT GLOB '*[^0-9]*' AND turn_epoch <= '18446744073709551615'),
		state TEXT NOT NULL CHECK(state IN ('ACTIVE','RELEASED')),
		granted_at INTEGER NOT NULL,
		released_at INTEGER,
		PRIMARY KEY(run_id, lease_kind),
		CHECK((state='ACTIVE' AND released_at IS NULL) OR (state='RELEASED' AND released_at IS NOT NULL))
	)`,
	`CREATE TABLE IF NOT EXISTS pending_app_events (
		event_id TEXT PRIMARY KEY,
		run_id TEXT NOT NULL REFERENCES runs(run_id),
		turn_id TEXT REFERENCES turns(turn_id),
		event_bytes BLOB NOT NULL,
		reported_by TEXT NOT NULL,
		created_at INTEGER NOT NULL
	)`,
	`CREATE TABLE IF NOT EXISTS cancellations (
		cancellation_id TEXT PRIMARY KEY,
		run_id TEXT NOT NULL REFERENCES runs(run_id),
		target_kind TEXT NOT NULL CHECK(target_kind IN ('run','turn','attempt')),
		target_id TEXT NOT NULL,
		epoch TEXT NOT NULL CHECK(length(epoch)=20 AND epoch NOT GLOB '*[^0-9]*' AND epoch <= '18446744073709551615'),
		disposition TEXT NOT NULL CHECK(disposition IN ('PENDING','ACKNOWLEDGED','COMPLETED','REJECTED')),
		requested_at INTEGER NOT NULL,
		resolved_at INTEGER,
		UNIQUE(target_kind, target_id, epoch)
	)`,
	`CREATE TABLE IF NOT EXISTS wake_schedule (
		relay_id TEXT PRIMARY KEY,
		run_id TEXT NOT NULL REFERENCES runs(run_id),
		disposition TEXT NOT NULL CHECK(disposition IN ('PENDING','ADMITTED','IGNORED')),
		received_at INTEGER NOT NULL,
		admitted_turn_id TEXT REFERENCES turns(turn_id)
	)`,
	`CREATE TABLE IF NOT EXISTS broker_control (
		singleton INTEGER PRIMARY KEY CHECK(singleton=1),
		control_token TEXT NOT NULL,
		control_generation TEXT NOT NULL CHECK(length(control_generation)=20 AND control_generation NOT GLOB '*[^0-9]*' AND control_generation <= '18446744073709551615'),
		minted_at INTEGER NOT NULL
	)`,
	`CREATE TABLE IF NOT EXISTS broker_events (
		broker_instance_nonce TEXT NOT NULL,
		event_seq TEXT NOT NULL CHECK(length(event_seq)=20 AND event_seq NOT GLOB '*[^0-9]*' AND event_seq <= '18446744073709551615'),
		event_type TEXT NOT NULL CHECK(event_type IN ('boundary_cut','epoch_installed')),
		run_id TEXT NOT NULL REFERENCES runs(run_id),
		turn_epoch TEXT NOT NULL CHECK(length(turn_epoch)=20 AND turn_epoch NOT GLOB '*[^0-9]*' AND turn_epoch <= '18446744073709551615'),
		event_bytes BLOB NOT NULL,
		ack_bytes BLOB NOT NULL,
		committed_at INTEGER NOT NULL,
		PRIMARY KEY(broker_instance_nonce, event_seq)
	)`,
	`CREATE TABLE IF NOT EXISTS content_ready_receipts (
		run_id TEXT NOT NULL REFERENCES runs(run_id),
		turn_id TEXT NOT NULL REFERENCES turns(turn_id),
		attempt_id TEXT NOT NULL REFERENCES provider_attempts(attempt_id),
		round_identity TEXT NOT NULL,
		seq_hwm TEXT NOT NULL CHECK(length(seq_hwm)=20 AND seq_hwm NOT GLOB '*[^0-9]*' AND seq_hwm <= '18446744073709551615'),
		generation_id TEXT NOT NULL REFERENCES workers(generation_id),
		committed_at INTEGER NOT NULL,
		PRIMARY KEY(run_id, turn_id, attempt_id)
	)`,
	`CREATE TABLE IF NOT EXISTS turn_disclosures (
		disclosing_turn_id TEXT NOT NULL REFERENCES turns(turn_id),
		source_turn_id TEXT NOT NULL REFERENCES turns(turn_id),
		tool_call_id TEXT NOT NULL,
		ticket_id TEXT NOT NULL,
		state_at_disclosure TEXT NOT NULL CHECK(state_at_disclosure IN ('UNKNOWN_TOOL_OUTCOME','PARTIAL_TOOL_EFFECT')),
		canonical_tool_name TEXT NOT NULL,
		canonical_args_digest TEXT NOT NULL CHECK(length(canonical_args_digest)=64 AND canonical_args_digest NOT GLOB '*[^0-9a-f]*'),
		PRIMARY KEY(disclosing_turn_id, tool_call_id, ticket_id)
	)`,
}
