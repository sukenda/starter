# Fullstack Starter Monorepo

Production-oriented starter for applications that share one repository across:

- Backend: Go
- Frontend: Vue 3 + TypeScript
- Mobile: Flutter

The repository is designed to be AI-friendly and intentionally keeps each ecosystem native while sharing architecture, contracts, documentation, and automation at the repository root.

## Repository layout

```text
apps/
  backend/      Go API
  frontend/     Vue 3 web app
  mobile/       Flutter app
packages/
  api-contract/ OpenAPI contract shared across applications
  design-tokens/Shared design tokens and conventions
docs/
  architecture/ Architecture documentation
  decisions/    Architecture Decision Records (ADR)
  development/  Development guides
scripts/         Repository automation
.github/         CI workflows and pull request templates
```

## Design principles

1. **One product, one context** — backend, frontend, mobile, contracts, and architecture live together.
2. **Native tooling** — Go uses Go tooling, Vue uses pnpm/Vite, Flutter uses Flutter tooling.
3. **Feature-oriented code** — organize application code around features/domains rather than technical folders alone.
4. **API contract first** — `packages/api-contract/openapi.yaml` is the shared HTTP contract.
5. **Thin delivery layer** — handlers/controllers and widgets/components should delegate business logic.
6. **AI guardrails** — read `AGENTS.md` and the closest application-level `AGENTS.md` before changing code.
7. **Progressive architecture** — avoid abstractions that do not yet solve a real problem.

## Quick start

Requirements:

- Go 1.25+
- Node.js 22+
- pnpm 10+
- Flutter stable / Dart 3+
- Docker + Docker Compose

```bash
make setup
make dev
```

Or run applications independently:

```bash
make dev-backend
make dev-frontend
make dev-mobile
```

## Quality commands

```bash
make lint
make test
make check
```

Each command delegates to the native ecosystem tooling.

## Example feature

The starter includes a minimal `health` path as the first vertical slice. New features should follow the same repository rules and extend the OpenAPI contract when the public API changes.

## AI usage

Before implementing a feature, an AI coding agent should read:

1. `/AGENTS.md`
2. `/ARCHITECTURE.md`
3. The relevant `apps/<app>/AGENTS.md`
4. `packages/api-contract/openapi.yaml` when API behavior is involved

See `docs/development/feature-workflow.md` for the expected implementation flow.
