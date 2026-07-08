# s6 Step-Exit Procedure Of Record

This procedure is the operator-run gate-3 script for the fixed s6 conductor. It is transport/provenance-only: success means the conductor accepted, projected, and exposed the named transport facts; it does not verify work content beyond the recorded relay protocol.

## Preconditions

- Use a fresh blessed store root.
- Use a short socket path such as `/tmp/frank-s6-step.sock`.
- Every live seat session is pre-allowed to use the three MCP tools: `mcp__frank__submit`, `mcp__frank__project`, and `mcp__frank__read`.
- The conductor branch is `s6-transport-impl` at the step-exit build under test.
- Operator captures credentials out of band; credential bytes are never pasted into relay bodies, projections, INDEX rows, or logs.

## Fixed Sequence

1. Initialize the fresh store with pinned config.

   Evidence artifact: `results/step-exit-init.log`

   ```sh
   cat > /tmp/frank-s6-engine.json <<'JSON'
   {"gc_enabled":false,"segment_rotate_bytes":4194304}
   JSON
   frank -root "$FRANK_S6_STORE" \
     -registry internal/fieldspec/registry.json \
     -engine-config /tmp/frank-s6-engine.json \
     -init
   ```

2. Bootstrap only the operator credential before serving.

   Evidence artifact: `results/step-exit-bootstrap-operator.log`

   ```sh
   frank -root "$FRANK_S6_STORE" -mint operator -role operator -operator
   ```

3. Start the conductor and keep it serving for the remaining legs.

   Evidence artifact: `results/step-exit-conductor-start.log`

   ```sh
   frank -root "$FRANK_S6_STORE" -socket "$FRANK_S6_SOCKET"
   ```

4. Apply the s5 registry/config update as the first live operator act.

   Evidence artifact: `results/step-exit-registry-apply.json`

   Operator submits an operator-authored `record_kind: config_change` for member `fieldspec`, with `new_digest` set to the composite pinned-config chain digest (`config.Digest` over all members after replacing the target member), not the target member's standalone hash. Expected reply state: `accepted`.

   Gate-day note: the operator refused to attest until this digest rule was shown and independently reproduced. That hold was the intended custody discipline.

5. Run the ROADMAP:83-85 legs against the live conductor.

   Evidence artifacts:
   - `results/step-exit-roadmap-83.json`
   - `results/step-exit-roadmap-84.json`
   - `results/step-exit-roadmap-85.json`

   Each leg records the submit reply, the visible `project` output for the addressed seat, and a `read` of the committed relay.

6. Re-drive the archived F11 dogfood traffic live.

   Evidence artifact: `results/step-exit-dogfood-redrive.json`

   Use the archived sequence as the source. Expected claim: accepted archive traffic accepts on the fixed conductor with zero parent-class bounces and zero digest re-render bounces. The race class is covered by the in-tree concurrent-accept fixture, not by timing-faithful dogfood replay.

7. Live boot a new seat without restart.

   Evidence artifact: `results/step-exit-live-boot.json`

   Sequence:
   - Operator submits `record_kind: seat_mint` with Body JSON `{"seat":"<seat>","role":"<role>","is_operator":false}`.
   - Operator captures the credential and endpoint from that one accept reply.
   - The new seat connects through `frank-mcp`.
   - The new seat renders the pre-active boot form.
   - If any hosted-seat session receives a form re-render bounce, re-read the schema before retrying. Hosted seats do not consume `tools/list_changed`; do not trust a cached tool constant after any re-render bounce.
   - The new seat submits boot fields `PHASE`, `CEREMONY_TIER`, `SUBJECT`, `charter_loaded`, and `dispatch_status`.
   - Operator calls `project` with `{"view":"roster"}` and records the row showing the seat moved from `minted` to `active`.

## Dry-Run Boundary

This task dry-runs only conductor-side setup: build, init, bootstrap operator mint, short-socket start, and clean stop on a scratch store. It does not simulate operator approval, live-seat designation, registry content application, dogfood archive ownership, or the final step-exit evidence capture.
