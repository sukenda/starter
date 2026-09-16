# Fullstack Starter Monorepo

Production-oriented, AI-friendly starter for one product across three native applications.

## Stack baseline

- **Backend:** Go 1.27.1, Fiber v3, `database/sql`, MariaDB/MySQL driver
- **Database:** MariaDB 12.3 LTS (MySQL protocol compatible)
- **Frontend:** Vue 3.5, TypeScript, Vite 8, Vue Router, Pinia, TanStack Query, Zod
- **Mobile:** Flutter 3.47 stable, Riverpod, GoRouter, Dio, Freezed/json_serializable
- **Contract:** OpenAPI in `packages/api-contract/openapi.yaml`
- **Automation:** Make + GitHub Actions; each ecosystem keeps its native toolchain

We intentionally use stable releases for the starter rather than RC/beta releases.

## Repository layout

```text
apps/
  backend/      Go API
  frontend/     Vue web application
  mobile/       Flutter application
packages/
  api-contract/ Canonical OpenAPI contract
  design-tokens/Reserved for the shared design system phase
docs/
  architecture/ Architecture documentation
  decisions/    Architecture Decision Records
  development/  Development guides
scripts/         Repository automation
.github/         CI workflows
```

## Design principles

1. **One product, one context** — backend, frontend, mobile, contracts, and architecture live together.
2. **Native tooling** — do not force Go or Flutter through a JavaScript monorepo orchestrator.
3. **Feature-oriented code** — organize product code around domains/features.
4. **API contract first** — public API changes must update OpenAPI.
5. **Thin delivery layer** — handlers and UI delegate business logic.
6. **AI guardrails** — agents read root and application-level `AGENTS.md` files first.
7. **Progressive architecture** — abstractions must solve a demonstrated problem.
8. **Design-system ready** — frontend and mobile structure leaves visual primitives centralized for the next phase.

## Quick start

Requirements: Go 1.27+, Node.js 22+, pnpm 10+, Flutter 3.47+, Docker with Compose.

```bash
make setup
make infra-up
```

Then run each application in its own terminal:

```bash
make dev-backend
make dev-frontend
make dev-mobile
```

MariaDB is exposed on `localhost:3306` with development-only credentials defined in `docker-compose.yml`. Copy `apps/backend/.env.example` when you need local overrides; never commit secrets.

## Quality gate

```bash
make fmt
make check
```

`make check` runs native type/lint/analyze and test commands for all three applications.

## AI usage

Before implementing a feature, read in order:

1. `/AGENTS.md`
2. `/ARCHITECTURE.md`
3. the closest `apps/<app>/AGENTS.md`
4. `packages/api-contract/openapi.yaml` for API work
5. `docs/development/feature-workflow.md`

The starter currently establishes infrastructure and a health vertical slice. Authentication/users and the shared frontend/mobile design system should be added as explicit subsequent vertical slices rather than hidden framework magic.
