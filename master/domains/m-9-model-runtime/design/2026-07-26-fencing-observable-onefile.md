# m-9 writer-fence observable — one-file-per-run successor (Route-4 / close4)

**Status:** PROPOSED — m-9-authored exact successor contract for the writer-fence observable and its negative fixture under the operator-ratified one-file-per-run D1 re-scope (Decisions 1–3).
Fresh design artifact (own `DESIGN_DOC_ID`, own SHA) — supersedes the fence text this doc replaces; binds no lock; r17 `01b885fe…` UNMOVED.
Awaiting m-9.implementer exact-byte review → m-10 Route-4 reciprocal → m-3 evidence confirm → §7 fixture row (record count / sample weight left to m-3+m-10+l4).

## 1. The fence, unchanged in mechanism, re-based to one file
Under one-file-per-run there is exactly one journal per run and exactly one writer permitted to it.
The fence is the **per-run advisory `flock`** on a **dedicated lock file** — NOT on the journal — retained verbatim in mechanism from r17 §1.6:139-143.
The advisory-lock residual (a non-cooperating process can still write the bytes) is honestly named and is the accepted MVP posture (confusion-not-malice); a mandatory-lock or fd-passing hardening is Step-4.

## 2. The exact acquisition sequence (M9-CLOSE4-R1-F2)
The observable is the **order**, stated exactly:
1. open the dedicated lock file `<per-run runtime dir>/session.lock` with `O_CLOEXEC`;
2. attempt `flock(fd, LOCK_EX | LOCK_NB)`;
3. **only after success** open / read-for-trust / attach the single per-run journal, then append.

On `flock` failure (would-block): perform **no read-for-trust, no journal attach, and no write**; report through the bounded fault path.
The lock is on `session.lock`, acquired **before** any journal open/read/attach — the journal open is NOT the lock acquisition. Failure is fail-closed.

## 3. The negative fixture is a TWO-ACTOR sequence (M9-CLOSE4-R1-F3)
The WRONG_LEASE / double-writer condition is two distinct actors with two distinct consequences on two distinct surfaces. One actor cannot both receive no `assign`/`turn_open` and then acquire the worker lock. The m-9 sub-observation is actor B only; actor A's rejection is m-10-owned.

| actor | who | what happens | surface / evidence locator | zero-work claim |
|---|---|---|---|---|
| **A — disposed predecessor generation** | the generation whose lease was released, attempting to bind/admit again | m-10's **admission gate rejects it pre-assign** — no `assign`, no `turn_open` is ever emitted to it (`...route4-m10-ans-1/...033500.md:24-35`) | m-10 admission-reject record (m-10-owned; NOT an m-9 observable) | n/a — it never reaches the worker |
| **B — legitimate current replacement** | the correctly-assigned replacement generation | it **is** assigned + `turn_open`'d, then attempts `session.lock` per §2; `flock` fails because the **still-live predecessor retains the open file description**; it performs **no recovery-read, no journal attach, no write, no provider call, no tool call, no conductor verb**, and emits the specified bounded fault | the m-9 fault record from the fail-closed path in §2 (the m-9 sub-observation) | zero recovery/journal/provider/tool/conductor work, actor-scoped to B |

The m-9 join is the **actor-B sub-observation** (assigned replacement blocked at `session.lock`, zero downstream work, bounded fault). If the final N.b fixture is instead defined to observe **only** actor A's pre-worker admission rejection, then there is **no m-9 observable in it** and m-9's join to that fixture closes (nothing for m-9 to bind).

## 4. What m-9 does NOT own here
m-9 does not own m-10's admission predicate, the disposal/lease-release decision, or the final fixture cardinality / record count / sample weight — those are m-10+m-3+l4. m-9 owns only the acquisition order (§2) and the actor-B zero-work fault emission (§3).

## 5. Boundaries
No lock moved; r17 `01b885fe…`, rev16 `3e3c5192…`, §D amendment, interface lock `cbd1893c…` UNMOVED. The §7 fixture row + sample weight are bound by m-3+m-10+l4, not this doc. H-12 hard-blocks external use.
