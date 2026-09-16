# Architecture

## Context

This repository contains three deployable applications that form one product:

```text
                 packages/api-contract
                         |
          +--------------+--------------+
          |              |              |
          v              v              v
   apps/backend    apps/frontend    apps/mobile
       Go API         Vue SPA         Flutter
```

Each application owns its runtime and dependencies. The monorepo exists to share product context, contracts, engineering rules, documentation, and automation—not to couple implementation details.

## Core principles

### Feature/domain orientation

Application code should be grouped around business capabilities. Shared technical infrastructure stays in explicit core/platform packages.

### Contract-first HTTP

Public API changes start in `packages/api-contract/openapi.yaml`. Clients and server must agree with that contract. Generated clients may be introduced later, but generated code must remain reproducible and must not become the architecture itself.

### Dependency direction

Business/domain logic must not depend on HTTP frameworks, database drivers, Vue components, or Flutter widgets.

### Progressive complexity

Start with the simplest architecture that preserves boundaries. Add queues, caching, CQRS, event sourcing, microservices, or elaborate abstraction layers only when requirements justify them.

## Backend

Initial flow:

```text
HTTP request
  -> handler
  -> service/domain behavior
  -> repository when persistence is required
  -> PostgreSQL
```

Handlers translate transport concerns. Services coordinate business behavior. Repositories encapsulate persistence queries. Small features may omit a layer when it adds no value.

## Frontend

```text
route/page
  -> feature component
  -> feature composable/query
  -> shared API client
  -> backend
```

TanStack Query owns server state. Pinia is reserved for genuine cross-route client state. UI primitives are shared; feature-specific components remain inside the feature.

## Mobile

```text
route/screen
  -> presentation
  -> Riverpod provider/controller
  -> repository/data source
  -> shared Dio client
  -> backend
```

Widgets render state and emit intent; network/business behavior stays outside widgets.

## Shared packages

Shared packages are deliberately small. Do not create a cross-language "shared business logic" package. Share protocols and design tokens, not runtime implementation.

## Deployment

Each application is independently buildable and deployable. Root automation orchestrates local development and CI while retaining native commands for every ecosystem.

## Observability baseline

Backend services should support structured logging, request IDs, health/readiness checks, graceful shutdown, and future metrics/tracing integration. Clients should centralize error reporting boundaries so observability can be added consistently.

## Security baseline

- Secrets come from environment/runtime secret management.
- Validate untrusted input.
- Apply least privilege.
- Keep authentication and authorization distinct.
- Never log credentials, tokens, passwords, or sensitive payloads.
- Pin/lock dependencies and review automated dependency updates.

Architecture decisions that materially change these principles should be recorded in `docs/decisions/` as ADRs.
