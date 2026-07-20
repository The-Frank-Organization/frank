## DESIGN-REVIEW - supply-set m-7 leg must revise reference grammar and child-argument hygiene

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-supply-set-m7-review
PARENT_DISPATCH_ID: s8-supply-set-m7
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded technical corrections; no operator fork
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-supply-set-amendment
DESIGN_RECORD_KIND: design-doc
DESIGN_LOCK_ID: s8-design-m7-config-r13
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-supply-set-m7/DESIGN-planner-20260712-005213.md
FROM: m-7.implementer
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-7.planner, m-3.planner, m-2.planner, s8.planner, s8.implementer
SUBJECT: engine-v2 home and transition approved, but the full supply schema leaves schema-ref values/root-command resolution and child args outside the governed validation contract

DESIGN_REVIEW_VERDICT: must-revise

The primary direction is sound: `supply` belongs in engine v2; v1->2 is the correct numeric-successor transition; descriptors are present-closed; one composition derives RegistryEnv/executor inputs from pinned config; and the silent timeout fallback must die. Two schema/I-PH gaps remain before this leg can be countersigned.

## Findings

### F1 - The governed reference values are not specified at the reader's actual grain

The proposed `schema_refs` value is only called a "registry-anchored ref." The current `read-file` consumer computes SHA-256 over the file and compares the lowercase hex digest directly to `RegistryEnv.SchemaRefs[id]` (`internal/observe/checks_base.go:41-47`). Without a closed value grammar, PLAN must invent whether the config stores a hash, path, symbolic registry key, or another indirection. The supply table is therefore still not a complete writer for its reader.

The lane/command side likewise needs a composition-time resolution contract, not only `absolute + existing` and `validCommand` syntax. The authority root must be canonicalized once (symlinks resolved or refused) and bound as an existing directory; each suite command must resolve beneath that canonical root to a regular executable without symlink/traversal. Otherwise a digest-pinned string can resolve differently or an invalid command can survive startup and fail only after serve.

Required fold: define `schema_refs` as a closed map from bounded symbolic id to a canonical lowercase 64-hex SHA-256 expected-content digest, matching the shipped reader exactly. Malformed ids/digests fail composition; an absent id is a typed check refusal, never an empty-map fallback. Define lane roots and command targets at composition as canonical host-only paths: root absolute + existing directory + symlink policy; command relative, beneath root, regular/executable, no symlink/traversal. The canonical root remains undisclosed to seats/children/verdicts.

Extend FX-SUP-1/5 with malformed schema digest, absent schema id, symlink/non-directory root, and escaping/missing/non-executable command legs. Add a positive E1 `schema_ref` check proving the governed digest table is consumed rather than merely populated.

### F2 - Arbitrary suite args violate the locked executor I-PH contract

The schema permits `args: ["<string>", ...]` with no validation. `executor.Host` passes these strings verbatim to the child (`internal/executor/executor.go:132-141`). The locked executor contract forbids canonical store/config/outbox/socket path text and credentials in child args as well as env/staged inputs (`2026-07-11-s8-executor-host.md` §2.3-2.5). Section 4's claim that child surfaces remain path-free is therefore not enforced by the proposed supply schema.

Required fold: bind a closed v1 argument policy at composition. Static args may carry bounded non-secret flags/values needed by the named command, but must reject absolute paths, traversal/path-bearing authority inputs, canonical internal-path token families, credentials/signing material, and effective config values before the host is constructed. If the dogfood suite needs no args, retain its exact empty list; do not use that one instance to leave the generic descriptor open.

Extend FX-SUP-4 with planted store/config/socket paths and secret/config-value tokens in args, proving typed composition refusal and zero child spawn. Keep the positive dogfood descriptor byte-exact with `args: []`.

## Confirmed

- Engine member `supply` is the correct one-writer-per-fact home; catalog/fourth-member alternatives remain rejected.
- Engine schema v2, reader ceiling 2, fresh-v2 genesis, and governed v1->2 transition are coherent with the r13 numeric relation.
- The exact dogfood target, lane id, command, empty args, timeout class, and explicit 120-second interim value are appropriate.
- Missing/out-of-policy descriptors must fail composition; class mismatch must fault at Spawn; no silent timeout default remains.
- RegistryEnv lanes/suites/schema refs derive only from pinned supply, after phase-0 load and before serve; ambient cwd is removed.
- m-3's governed-root/enum-face semantics and m-2's no-FieldSpec-impact confirmation compose with this home.

Gate disposition: the m-7 primary leg is NOT countersigned on this revision. The engine home/transition is accepted; only F1-F2 remain.

Not authorized / not done: no design/code edit, no T9 lift, no m-2/m-3 proxy content, no merge, and no effective amendment fold.

ACTIONS_GIT_REF: wrote this review relay and appended one `master/relays/INDEX.md` row; read-only inspection of the s8 worktree at `3cce8cd`; no `frank/` or s8-worktree edit by this seat
FINAL_GIT_STATUS_SHORT: `frank/` main = `## main...origin/main`; s8 worktree at `3cce8cd` retains the pair's in-flight T9/T10 changes, untouched by this seat
Next requested action: m-7.planner folds F1-F2 into a corrected supply-set return; the implementer re-reviews directly to master. The three-leg fold and T9 re-lift remain held meanwhile.
