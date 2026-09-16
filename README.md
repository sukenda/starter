# Fullstack Starter Monorepo

Production-oriented, AI-friendly starter for one product across Go, Vue, and Flutter applications.

## Included baseline

- Go + Fiber API with MariaDB/MySQL, migrations, health/readiness, structured logging and graceful shutdown
- database-backed authentication with Argon2id passwords, opaque access/refresh tokens, rotation and revocation
- multi-role RBAC with permissions and permission-aware navigation
- Vue authenticated admin application with dashboard, users, roles and menus management
- Flutter authentication and authenticated home foundation
- cross-platform semantic design tokens and reusable UI primitives
- OpenAPI contract, Make automation and separate GitHub Actions workflows for backend/frontend/mobile
- root and application `AGENTS.md` guidance for AI-assisted development

## Repository layout

```text
apps/
  backend/       Go API
  frontend/      Vue web application
  mobile/        Flutter application
packages/
  api-contract/  Canonical OpenAPI contract
  design-tokens/ Cross-platform semantic design tokens
docs/
  architecture/  Architecture and security documentation
  decisions/     Architecture Decision Records
  development/   Development and delivery guides
scripts/          Repository automation
.github/          CI workflows
```

## Local quick start

Requirements: Go 1.27+, Node.js 22+, pnpm 10+, Flutter 3.47+, Docker with Compose.

```bash
make setup
make infra-up
```

Copy `apps/backend/.env.example` if local overrides are needed. Never commit production credentials. Then run applications in separate terminals:

```bash
make dev-backend
make dev-frontend
make dev-mobile
```

MariaDB development credentials are defined in `docker-compose.yml` and are for local use only.

## Quality gate

Before opening or merging a pull request:

```bash
make fmt
make check
```

GitHub Actions independently validates backend, frontend and mobile changes. A PR should not be treated as production-ready until the relevant checks pass.

## Production checklist

Before deployment, configure production secrets outside the repository, use TLS at the ingress/reverse proxy, restrict MariaDB from public access, set environment-specific API origins, run migrations as a controlled release step, configure centralized logs/metrics/alerts, establish database backups with restore testing, and define rollback procedures. Do not reuse development credentials or defaults.

Authentication sessions are server-side and revocable. Refresh tokens are single-use: rotation is an atomic compare-and-swap so concurrent reuse of an already-rotated refresh token is rejected. Keep access-token lifetimes short and review session TTLs for each product's risk profile.

## AI workflow

Read these before implementing a feature: `/AGENTS.md`, `/ARCHITECTURE.md`, the closest `apps/<app>/AGENTS.md`, `packages/api-contract/openapi.yaml` for API work, and `docs/development/feature-workflow.md`. Keep changes vertical, preserve security boundaries, update contracts/docs with behavior changes, and finish each scope through reviewable CI-backed pull requests.
