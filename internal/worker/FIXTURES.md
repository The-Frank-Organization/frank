# m-9 worker fixture traceability

Status: T13 implementation evidence for `s15-m9-worker`. This is an E2 battery,
not the Step-3 exit corpus and not live E3 evidence.

Normative inputs:

- worker r7: `2026-07-19-mvp-full-worker.md`, especially §6 and §10
- lifecycle r21: `2026-07-17-mvp-lifecycle-half.md`
- D1/D2 delta r15: `2026-07-22-relock-lane2-m9-delta.md`, especially §10
- one-file writer/reader and resume notes dated 2026-07-26/27
- locked s15 plan r4 at SHA-256
  `9c5f56a974d6ae7ee7f5b7052e31ac0e711c8900b695c2fbd079f806bd8de81e`

`master/exit-fixtures/**`, the frozen exit oracle `d4580c52…`, and RLBS-1
are not read, copied, edited, or re-derived by this battery.

## Build-task battery

| Family | Green proof | Deliberately broken / anti-vacuity proof |
|---|---|---|
| T1 JCS + counters | `TestCanonicalizeRFC8785PrimitiveVector`, `TestDigestIsLowercaseSHA256AndInputOrderInvariant`, `TestCounterRoundTrip` | `TestCanonicalizeRejectsNonIJSONInput`, `TestParseCounterRejects` |
| T2 F58 catalog | `TestExpectedCatalogDigest`, `TestConsumedM2IdentityBytesArePinned` | `TestCanonicalSetEqualityRejectsAbsentMember` |
| T3 frame codec | `TestFrameRoundTripIsByteExact`, `TestFrameReplyCorrelationAndClosedBody` | `TestFrameRejectsMalformedEnvelopeClasses`, `TestFrameRejectsOversizeBeforePayloadAllocation` |
| T4 one-file journal | `TestWriterFreshCreateResumeAndCrossGenerationHandoff`, `TestRecoveryCrashTable`, `TestRoundMarkerDigestAndMembership` | `TestDescriptorBatteryRejectsSymlinkModeHardlinkAndReplacement`, `TestRecoveryFirstFaultWinsAndLaterMarkerCannotRescue`, `TestRecoveryRejectsSuspectMarkerToPriorBoundary` |
| T5 F59 executor | `TestEveryCatalogToolUsesUniformAuthorityPath`, `TestExecutedOutcomeCarriesActualInvocationIdentity` | `TestPreparedCallIsInertWithoutAuthorityPath`, `TestMutationBeforeConsumeReturnsIdentityMismatchAndLeavesTicketIssued`, `TestMutationAfterConsumeRecordsIntegrityFaultAndNeverInvokes` |
| T6 local tools | `TestRegistryDispatchesAllFiveLocalTools`, `TestWriteEditAndApplyPatch` | `TestArgumentsRejectUnknownMissingAndTrailingMembers`, `TestFileToolsStayWithinResolvedWorkspace`, `TestMalformedPatchHasNoPartialEffect` |
| T7 mapping + conductor facade | `TestLockedReferenceFingerprintAndBranchCoverage`, `TestNativeAndMCPUseSameConductFacadePayloadAndH16StillGates`, `TestShimReconnectsAndRetriesSingleCallAfterConnectionLoss` | `TestStrictParseTypedErrorIdentities`, `TestStaticRelayArgumentsFailBeforeFacadeCall`, `TestH16MCPRejectsForgedSystemHeadersBeforeConductorCall` |
| T8 turn machine | `TestOneActiveTurnAndStaleEpochAreInert`, `TestParkedUnknownComparatorTotalAndGateOrdering`, `TestCancelAckExactlyOnceAndPrecedesCancelledTerminal` | `TestMalformedAndDuplicateParkedUnknownFailClosed`, `TestEveryCompiledBoundTerminatesExhausted` |
| T9 provider cycle | `TestAttemptOrderingTotalTerminalMappingAndNoRetry`, `TestOpaqueResponseItemRoundTripsThroughJournalBytePreserved`, `TestBothCancellationCuts` | `TestEpochRejectIsAttemptInert`, `TestConnectorFaultDoesNotAutoRetry` |
| T10 context + E0 | `TestPinnedTierSurvivesEveryEvictionAndSummary`, `TestLogicalSurfaceDigestStableNoOpAndMovesOnAnySurfaceChange`, `TestProviderE0TotalTableAndNoEmissionCuts` | `TestPinnedIntegrityFailureIsFailClosed`, `TestE0RedactionRejectsSecretShapedMembers`; the E0 table test also injects contradictory `attempt_started` rows |
| T11 resume | `TestAllFiveFirstActionBranchesAreReachable`, `TestSettlementManifestClosedUnionAndIdentityGrain`, `TestPrefixOracleStopsAtFrozenFullKeyAndExcludesRecordDigest` | `TestTrustWindowViolationsNeverTrustContent`, `TestEditedSessionClassDistinctnessAndNonPromotion`; malformed union legs return no partial manifest |
| T12 governed turn | `TestOneHonestGovernedTurnE2E`, `TestWakeRelayObjectiveUsesWorkerSeatResolver` | `TestHandshakeFailuresFailClosedBeforeProviderOrToolWork`; the wake test removes the resolver and proves zero provider/tool work |

The enclosing-guard reachability rule is explicit in the named tests: H-16 is
reached through the refactored MCP path; F59 mutations occur once before each
of its two fresh derivation points; journal negatives first establish a clean
genesis/boundary; Gate-2 comparator tests begin from an admitted assembling
turn; and the governed-turn trace reaches attach, attempt-open, provider,
authorize, consume, durable tool-result marker, outcome, and terminal in order.

## Worker r7 §10 and §6

| Source obligation | Disposition and proof |
|---|---|
| §10 single F59 gate, all eight names, reads have no fast path | GREEN — `TestEveryCatalogToolUsesUniformAuthorityPath`; zero-effect controls in `TestPreparedCallIsInertWithoutAuthorityPath` and `TestAuthorizeRejectsAndEpochMismatchAreAttemptInert` |
| §10 sync authorize/consume, async record, CONSUMED gap | GREEN — `TestAuthorityCrashWindowsAreTotal`, `TestDoubleConsumeIsRejected`, `TestCancellationAndEOFBeforeConsumeFailClosed` |
| §10 five local tools, bounded bash, group teardown, descriptor hygiene | GREEN — `TestRegistryDispatchesAllFiveLocalTools`, `TestReadAndBashOutputTruncateWithExplicitMarker`, `TestBashCancellationKillsWholeProcessGroup`, `TestBashDoesNotInheritUnlistedFileDescriptor`; escaped-`setsid` containment remains the ratified named residual, never a pass claim |
| §10 native relay tool, mapping, refresh, AXI envelope | GREEN — `TestStrictValidationAndCanonicalFacadeBytes`, `TestF1RefreshBeforeRejectAllowsFreshExpansion`, `TestF2EveryRejectedSubmitRefreshesAndReRenderNormalizes`, `TestUnknownToolReturnsAXIShapedFailure` |
| §10 governed compaction | GREEN for the T10-owned mechanism — `TestPinnedTierSurvivesEveryEvictionAndSummary`, `TestPinnedIntegrityFailureIsFailClosed`, `TestDestroyClearsEveryTier`; provider item opacity is independently pinned by `TestOpaqueResponseItemRoundTripsThroughJournalBytePreserved`. Token-budget tuning and external-quality evaluation are not E2 claims. |
| §10 durable one-file replay successor to the struck no-reload row | GREEN — `TestRecoveryCrashTable`, `TestResumeTruncatesOnlyUntrustedTail`, `TestOneHonestGovernedTurnE2E` byte-equal replay |
| §10 objective acquisition | GREEN — both `AdmissionRef.Validate` kinds plus `TestWakeRelayObjectiveUsesWorkerSeatResolver`; a wake relay requires the injected worker-seat resolver, while operator input is carried verbatim. Missing/malformed references fail before provider work. |
| §10 F58 catalog | GREEN — `TestExpectedCatalogDigest`, `TestLocalSchemaDigestsMatchPinnedIdentities`, `TestConsumedM2IdentityBytesArePinned`, `TestVersionMarkersAndMappingVersionShape` |
| §10 §6 E0 total table | GREEN — `TestProviderE0TotalTableAndNoEmissionCuts` covers completed, transport failure, denial, local reject, both cancellation cuts, stream loss, both epoch-inert tokens, crashed populator, and retirement-wins. `TestE0RedactionRejectsSecretShapedMembers` is the planted leak. |
| §10 anti-fabrication | GREEN — `TestAuthorityCrashWindowsAreTotal`, `TestOutcomeRecordDomainIsClosed`, and the no-auto-retry provider battery; no test manufactures a terminal effect from missing evidence. |

## Delta §10

| Delta row | Status | Proof or non-silent disposition |
|---|---|---|
| Log/crash battery (§10 first row) | SUPERSEDED/REALIZED | The segmented 15-row topology was removed by ratified D1 r11 §0C-4/§0C-6. Its live one-file successor is `TestRecoveryCrashTable`, `TestRecoveryGenesisFaultHasEmptyTrustedPrefix`, `TestResumeTruncatesOnlyUntrustedTail`, and `TestResumeAbsentNeverRecreates`. |
| Terminal-seal, seal/boundary, late-write, identity/progression, five append-selection branches, chain identity, boundary equation, chain topology | SUPERSEDED | All depend on removed `segment_*`, rotation, chaining, boundary-equation, or terminal-seal machinery. Ratified D1 r11 replaces them with one `session.log`, one per-run fence, `create_auth_id`, physical-order recovery, and last-good G/M boundaries. No test claims the removed mechanism. |
| At-rest descriptor battery | GREEN | `TestDescriptorBatteryRejectsSymlinkModeHardlinkAndReplacement` plus `TestSyntheticWrongOwnerDescriptorRejected`; paths are not returned in typed errors. |
| S-4 settlement-manifest binding | GREEN consumer half | `TestSettlementManifestClosedUnionAndIdentityGrain`: canonical top-level v1, closed `kind × class`, tool key `{run,source turn,tool_call,args_digest}`, provider key `{run,source turn,attempt}` with no args digest, no partial decode, and duplicate exact-identity rejection. Producer totality remains m-10-owned. |
| Writer fence | GREEN | `TestSecondWriterFenceViolation`, `TestWriterFreshCreateResumeAndCrossGenerationHandoff`, and descriptor-CLOEXEC coverage. A second generation cannot write while the predecessor owns `session.lock`. |
| Round membership and full round key | GREEN | `TestRoundMarkerDigestAndMembership`, `TestRecoveryRejectsSuspectMarkerToPriorBoundary`, `TestPrefixOracleStopsAtFrozenFullKeyAndExcludesRecordDigest` |
| Duplicate stream rows | GREEN live one-file subset | `TestAdjacentByteIdenticalDuplicateCollapses`, `TestRecoveryFirstFaultWinsAndLaterMarkerCannotRescue`; removed cross-segment variants are superseded as above. |
| Provenance and class distinction | GREEN | `TestClosedRecordUnionRejectsUnknownKindAndMembers`, `TestRecoveryContentTrustClassifiesExternalAndModelBytesDifferently`, `TestEditedSessionClassDistinctnessAndNonPromotion`, `TestSettlementManifestClosedUnionAndIdentityGrain` |
| Durability ordering + `content_ready` | GREEN worker half | `TestDurabilityOrderingRequiresMarkerBeforeOutcome`, the durable invoker in `TestOneHonestGovernedTurnE2E`, and `TestContentReadyOrderingAndDispositionGate`. Receipt persistence/equivalence is m-10-owned and joint-gated below. |
| D2/resume five branches + no-work gate | GREEN | `TestAllFiveFirstActionBranchesAreReachable`, `TestTrustWindowViolationsNeverTrustContent`, `TestContentReadyOrderingAndDispositionGate` |
| Assembly-refusal five classes and zero-attempt `turn_failed` | PARTIAL, non-silent | Exact catalog/set negatives exist (`TestCanonicalSetEqualityRejectsAbsentMember`, `TestVersionMarkersAndMappingVersionShape`) and halt before provider work, but this slice does not claim a live m-10 receipt for all five synthetic array-shape failures. The existing terminal is wired by the runtime control seam; counterpart proof remains under m-10's §D-5 fixture. No COMPLETE claim is made for that counterpart leg. |
| S-1 same-key receiver totality | HELD at owner boundary | m-9 constructs the closed receipt only (`TestContentReadyOrderingAndDispositionGate`). Equivalent-duplicate-before-stale and `receipt_conflict` persistence are m-10-owned and remain under the §D join; the fake does not substitute for the canonical store. |
| `logical_surface_digest` carrier / five operand recipes | GREEN for plan T10; counterpart-held for E3 | `TestLogicalSurfaceDigestStableNoOpAndMovesOnAnySurfaceChange`, catalog/mapping exact-byte tests, and Tier-2 sentinel tests prove the worker derivation properties. m-10 attempt-row persistence and m-3 independent reconstruction are sibling-owned and are not claimed by this E2 battery. |
| E/C/B and `relay.submit` canonical-resource cell | PARTIAL/HELD | E0 and relay mapping/native parity are green locally. `content_id` unknown, full effect-descriptor persistence, `frozen_core_digest` carriage, and external E3 joining remain at their named m-8/m-10/m-3 owner gates; no local pass substitutes for them. |
| `parked_unknown_capacity_exceeded` consumer posture | OWNER-EXTERNAL | m-10 owns the run terminal and no-revival proof. m-9's turn machine has no resume/first-action token for it by construction; the five-value first-action enumeration test excludes it. |

## m-2 PV and shared-front-end rows

The Appendix-A/PV family is green at the exact locked fingerprint
`306b3149a002f0648995f659926ba0f985ee771d95bb899c7ecc81adadab28ac`
through `TestLockedReferenceFingerprintAndBranchCoverage`. Strict parser,
stable/volatile render, required/enum, CC fold, re-render, F-1/F-2 freshness,
H-16, import graph, static schema identity, and native/MCP facade parity are
named by the T7 row above.

Branch A is operative under master's R3 ruling
`RECONCILE-orchestrator-planner-20260820-234501.md` and R7 Option 1
`RECONCILE-orchestrator-planner-20260821-033914.md`; therefore there are no
Branch-B exclusions in this build. The hypothetical Branch-B omitted MCP
halves (m-2 §6 items 1/4/5/9) are not marked complete or silently reassigned:
that branch required a separate master-named owner/dispatch/gate and a
PARTIAL/HELD return, but master selected Branch A and all seven granted MCP
files carry the shared-module consumer and tests.

## Gates that this table does not retire

- m-10 ticket-schema/store and settlement-receipt producer behavior remain
  owned by m-10; the in-process fake proves only worker consumability.
- S-1/S-2/S-4/S-5 stay subject to their exact reciprocal and §D join status;
  local green tests do not assert premature normativity.
- m-3 E3 independence, evidence locators, live liveness, and the Step-3 exit
  predicate are outside this E2 battery.
- m-8 real provider translation, credentials, and wire egress are absent; the
  fake never performs network egress.
- release/export, restack, merge, push, deployment, and slice close remain at
  T14/master/operator gates.
