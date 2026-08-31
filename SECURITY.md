# Security Policy

## Reporting a vulnerability

Please report vulnerabilities privately via GitHub's private vulnerability reporting on [The-Frank-Organization/frank](https://github.com/The-Frank-Organization/frank/security/advisories/new), or email the maintainer at lijiaquan0@gmail.com.
Please do not open a public issue for a security problem.

## Supported versions

frank is pre-1.0.
Only the latest commit on `main` is supported; there are no maintained release branches.

## Threat-model honesty

frank's ethos is honest labeling, and that applies to frank itself.
These are the current security boundaries, stated plainly.

- frank is a local, single-host, same-trust-domain system.
- Every wire is a unix domain socket; there is no TCP listener anywhere.
- It is not hardened for external, untrusted, or multi-tenant use, and treating it as such is a standing hard blocker, not a footnote.
- The MVP has no sandbox: the worker's `bash` tool runs with the invoking user's ambient authority.
- Seat identity is confusion-resistant, not theft-proof: a process running as the same uid, outside the tool surface, can reach the store files and sockets directly (the accepted D5 residual).
- The observe layer stamps per-field evidence integrity; fields it does not observe remain `self_reported`.

If a report shows frank failing a boundary it claims to hold, that is a vulnerability and we want to hear about it.
If a report shows frank lacking a boundary it explicitly does not claim (for example, sandboxing or multi-tenant isolation), it is likely known and documented, but reports that sharpen the documented boundaries are still welcome.
