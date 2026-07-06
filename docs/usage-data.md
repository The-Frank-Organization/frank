# usage data

transport/provenance only; done-state and `record_integrity` remain `self_reported` until Step-2 observe.

The frank store is the usage record. There is no analytics side channel in this slice.

Read usage through the same governance surfaces:

- `project` lists relay IDs visible to the seat mailbox.
- `read` returns a committed relay view for a relay ID.
- `records/` contains canonical sealed records.
- `projections/INDEX.md` and mailbox projections are derived read views.

Aggregation, dashboards, and cross-seat usage summaries are s5 work. This slice only documents where the durable record already lives.
