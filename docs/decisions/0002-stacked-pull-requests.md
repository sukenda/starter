# ADR 0002: Use stacked pull requests for starter milestones

- Status: Accepted
- Date: 2026-09-16

## Context

The starter spans backend, frontend, and mobile. Large pull requests make architecture review difficult and make AI-assisted changes harder to reason about.

## Decision

Use stacked pull requests. Every new milestone branches from the latest active milestone rather than repeatedly branching from `main`.

Each pull request has one primary concern and targets its direct predecessor. Once a predecessor is merged, the next pull request is retargeted to the new base as needed.

The intended sequence is monorepo baseline -> engineering foundation -> authentication/RBAC -> frontend admin shell -> mobile authentication/home -> design system.

## Consequences

- Reviews remain focused and incremental.
- Dependencies between milestones are explicit.
- CI failures can be attributed to a smaller change set.
- A later PR must not be merged before its prerequisite stack has landed.
