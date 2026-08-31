# usage data

The observe layer stamps per-field `evidence_integrity` (`observed` vs `self_reported`) at send; fields it does not observe remain `self_reported`.

The frank store is the usage record. There is no analytics side channel.

Read usage through the same governance surfaces:

- `project` lists relay IDs visible to the seat mailbox.
- `read` returns a committed relay view for a relay ID.
- `records/` contains canonical sealed records.
- `projections/INDEX.md` and mailbox projections are derived read views.

There is no aggregation, dashboard, or cross-seat usage summary layer today; the surfaces above are where the durable record lives.
