# Architecture

## Context

This repository contains three independently buildable applications that form one product:

```text
                 packages/api-contract
                         |
          +--------------+--------------+
          |              |              |
          v              v              v
   apps/backend    apps/frontend    apps/mobile
       Go API         Vue SPA         Flutter
```

Each application owns its runtime and dependencies. The monorepo shares product context, contracts, engineering rules, documentation, automation, and design tokens—not runtime implementation details.

> Detailed documentation map: [`docs/README.md`](docs/README.md). Runtime topology, per-app architecture, data model, API, security and end-to-end flows are maintained there.

## Core principles

### Feature/domain orientation
Application code is grouped around business capabilities. Shared technical infrastructure stays in explicit core/platform packages.

### Contract-first HTTP
Public API changes start in `packages/api-contract/openapi.yaml`. HTTP DTOs are explicit transport models: database structs, nullable SQL types, and internal numeric IDs must not leak into the public contract.

### Dependency direction
Business/application behavior must not depend on HTTP frameworks, database drivers, Vue components, or Flutter widgets. Services consume narrow interfaces owned by the consuming feature; concrete repositories implement those interfaces.

### Progressive complexity
Use the simplest architecture that preserves useful boundaries. Add queues, caching, CQRS, event sourcing, microservices, or additional abstraction layers only when requirements justify them.

## Backend

```text
HTTP request
  -> handler / transport DTO
  -> application service
  -> repository interface
  -> MariaDB repository
```

Handlers own transport parsing, validation, status codes, and response DTOs. Services coordinate business behavior. Repositories encapsulate SQL and transaction details. Internal numeric database IDs remain persistence details; public APIs use stable public IDs or domain codes.

Repository interfaces are introduced at behavior boundaries where they improve isolation and testing, not mechanically for every type. Transactions belong in the persistence/application operation that must remain atomic. Authorization is enforced server-side; client permission checks are only a UX optimization.

Detailed backend architecture: [`docs/architecture/backend.md`](docs/architecture/backend.md).

## Frontend

```text
route/page
  -> feature component
  -> feature query/mutation
  -> feature API adapter
  -> shared HTTP client
  -> backend
```

TanStack Query owns server state. Pinia is reserved for genuine cross-route client state. Feature API adapters isolate UI models from transport DTOs. UI primitives are shared; feature-specific components remain inside the feature.

Browser access credentials are kept in memory. Refresh uses an HttpOnly cookie with CSRF protection and session restoration after reload. Mobile uses a separate native refresh transport.

Detailed frontend architecture: [`docs/architecture/frontend.md`](docs/architecture/frontend.md).

## Mobile

```text
route/screen
  -> presentation
  -> Riverpod controller
  -> repository/data source
  -> shared Dio client
  -> backend
```

Widgets render state and emit intent; network and business behavior stay outside widgets. Credentials are stored only through the secure-storage abstraction. Authenticated requests use the centralized Dio auth path and serialized refresh rotation.

Detailed mobile architecture: [`docs/architecture/mobile.md`](docs/architecture/mobile.md).

## Data integrity

MariaDB constraints enforce relationships that must never become orphaned. Application validation complements database constraints for domain rules such as preventing cyclic menu hierarchies. List endpoints must avoid per-row N+1 queries and should add pagination before unbounded collections become large.

Detailed persistence relationships: [`docs/data-model.md`](docs/data-model.md).

## Shared packages

Shared packages are deliberately small. Do not create a cross-language shared business-logic package. Share protocols, schemas, configuration conventions, and design tokens instead.

## End-to-end flows

Implemented login, restore, refresh, logout, protected admin, RBAC and mobile retry sequences are documented in [`docs/flows/README.md`](docs/flows/README.md). Security details are in [`docs/security.md`](docs/security.md).

## Deployment and reproducibility

Each application is independently buildable and deployable. Root automation orchestrates local development and CI while retaining native commands for every ecosystem. Application dependency lockfiles should be committed when the ecosystem produces them; CI should use deterministic dependency installation and run relevant tests before merge. Remaining operational hardening is tracked in [`docs/operations.md`](docs/operations.md).

## Observability baseline

Backend services currently provide structured request/access logging, request IDs and health/readiness foundations. Full metrics/tracing/deployment monitoring remains production-hardening work and is not represented as complete.

## Security baseline

- Secrets come from environment/runtime secret management.
- Validate untrusted input and enforce authorization on the backend.
- Apply least privilege and keep authentication distinct from authorization.
- Never log credentials, tokens, passwords, or sensitive payloads.
- Rotate refresh credentials atomically and reject replay/concurrent reuse.
- Browser refresh secrets are HttpOnly and CSRF protected; native credentials use secure storage.
- Rate limiting and additional abuse controls remain required before broad internet exposure.
- Pin/lock dependencies and review automated dependency updates.

Architecture decisions that materially change these principles should be recorded in `docs/decisions/` as ADRs.
