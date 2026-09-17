# Starter Documentation

This directory is the map of the architecture and operating knowledge for the monorepo. Documentation describes the code that exists today; future work must be called out explicitly rather than documented as if already implemented.

## Start here

1. [`../README.md`](../README.md) — repository purpose and quick start.
2. [`../ARCHITECTURE.md`](../ARCHITECTURE.md) — system-level architecture and boundaries.
3. [`architecture/system-overview.md`](architecture/system-overview.md) — runtime topology and end-to-end request paths.
4. [`architecture/backend.md`](architecture/backend.md) — Go API architecture.
5. [`architecture/frontend.md`](architecture/frontend.md) — Vue SPA architecture.
6. [`architecture/mobile.md`](architecture/mobile.md) — Flutter architecture.
7. [`flows/README.md`](flows/README.md) — user and security flows.
8. [`data-model.md`](data-model.md) — persistence model and relationships.
9. [`api.md`](api.md) — HTTP API conventions and endpoint ownership.
10. [`security.md`](security.md) — authentication, authorization, browser/mobile session boundaries.
11. [`development/README.md`](development/README.md) — local development and quality gates.
12. [`operations.md`](operations.md) — runtime configuration, health, migrations, and production gaps.

## Existing focused documents

- `architecture/auth-rbac.md` — RBAC implementation details.
- `architecture/admin-management.md` — admin management boundaries.
- `architecture/frontend-admin.md` — admin SPA details.
- `architecture/mobile-auth.md` — mobile authentication details.
- `architecture/web-session-security.md` — browser token/cookie/CSRF model.
- `design-system.md` — cross-platform design system.
- `decisions/` — architecture decision records.

## Documentation rule

Any PR that changes a public endpoint, persistence relationship, application boundary, authentication/session behavior, major user flow, configuration contract, or deployment assumption must update the corresponding document in the same PR.
