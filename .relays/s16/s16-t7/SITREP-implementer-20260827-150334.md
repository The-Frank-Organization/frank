## SITREP — WP1 T7 FINDING UP: the approved real-connector loopback proof has no authorized local-TLS trust mechanism

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s16-t7-tls-finding-1
PARENT_DISPATCH_ID: s16-impl
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E3
HUMAN_GATE_REQUIRED: yes — an owner/master ruling is required because every truthful repair changes the approved contract or crosses the exact write fence
GRILL_REQUIRED: no
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: T7 spec-base mismatch and fence pressure — the real frank-connector cannot authenticate a demo-owned local TLS endpoint under the approved bytes; no in-fence mechanism may make the mandatory plain-go-test loopback E3 green

Implementation has reached the T7 composed-runtime proof after the production app/worker/broker path was wired and commit-gated. The mandated real connector plus controlled loopback endpoint is not realizable under the approved base and write fence without an upstream ruling.

Mechanical evidence at the implementation tree:

- `internal/connector/endpoint/endpoint.go:14-16` accepts only `https://`; plain loopback HTTP is rejected before composition.
- `cmd/frank-connector/main.go:82-83` calls `os.Clearenv()` before the production connector service constructs its provider transport. An inherited `SSL_CERT_FILE` trust override therefore cannot reach `transport.NewClient`.
- `internal/connector/service/service.go:66` selects `transport.NewClient()` when the real connector binary is used; the production constructor has no injected root-pool or demo-CA parameter.
- `internal/connector/transport/transport.go:81-83` constructs the TLS configuration from the system roots. Its configurable TLS-root path is package-private test machinery, not reachable by `cmd/frank-connector`.
- Plan section 2.4 excludes the connector command and internal connector packages. It authorizes neither a connector trust input nor a host trust-store mutation. The latter would also make plain `go test` stateful and machine-dependent, and this seat has no authority to mutate operator trust state.

Consequently the required combination cannot all hold at once at the dispatched bytes: real `frank-connector`; full `freeze→authorize→attach→send`; demo-owned local endpoint; authenticated HTTPS; plain hermetic `go test`; and no write outside the fence. A self-signed loopback TLS server fails certificate verification. HTTP is rejected. Silently using a fake provider or test-only connector would violate section 3. Mutating the host trust store would be an unauthorized external state change and non-hermetic evidence.

Requested ruling UP: choose and authorize an owner-reviewed repair, such as an m-8-owned, non-secret test trust-root input carried into the real connector over a bounded inherited descriptor, with the write fence amended for its exact loci; or supply a different owner-approved mechanism that preserves the real-process/full-wire/loopback/plain-test requirements. Until then, T7 E3 and WP1-close are honestly HELD. No connector byte, host trust setting, dependency, frozen member, or seam instrument byte was touched.

Unaffected progress continues inside the existing grant: T8's mechanical six-loci audit now records RED at `77f8c9db76b2e2eb90252b3098ccf7b29fd8b3a9` and GREEN file:line callers at the implementation tree. The current pushed production composition commit is `283676d3`; its gate was seam `64/0/64`, plain suite green including `test/fixtures` (`231.930s`), and `go vet ./...` green. This is a finding, not a request to weaken T7 or a completion claim.

ACTIONS_GIT_REF: read-only inspection of the locked connector endpoint, process bootstrap, service, and transport loci; in-fence implementation continued on s16-integration through 283676d3; no connector, trust-store, dependency, frozen, seam, merge, deploy, or release action
FINAL_GIT_STATUS_SHORT: concurrent governance-root changes were present and untouched; the implementation worktree contained only the new in-fence composed caller proof at this report instant
