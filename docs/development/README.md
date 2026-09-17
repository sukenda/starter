# Development Guide

## Prerequisites

Use the toolchain versions expected by each application (`go.mod`, frontend package metadata, Flutter/Dart metadata/CI). MariaDB/MySQL is required for backend flows that touch persistence.

## Repository workflow

The root `Makefile` is the orchestration entrypoint. Application-native tools remain authoritative inside each app. Keep changes scoped, run the affected application's checks, update contracts/docs when boundaries change, and submit through a PR.

## Backend

Work from `apps/backend`. Copy `.env.example` to the local environment shape required by the app, provide database/auth configuration, apply migrations, then run/build/test with Go/Make targets. Configuration validation intentionally fails fast for invalid production-critical values.

Before PR: format Go code, run `go vet ./...`, run tests including race checks used by CI, and verify migrations/contract changes.

## Frontend

Work from `apps/frontend`. Install dependencies with the package manager declared by the project, configure the API base URL through the documented Vite environment contract, then use package scripts for development, lint, tests and build.

Before PR: lint, unit tests and production TypeScript/Vite build must pass. Do not bypass browser session transport by persisting access/refresh tokens in localStorage.

## Mobile

Work from `apps/mobile`. Fetch Flutter dependencies, configure API base URL as expected by app config, and run on an emulator/device.

Before PR: run `dart format` using the CI-compatible SDK, `flutter analyze`, and `flutter test`. Protected feature repositories should use the authenticated Dio provider rather than custom refresh logic.

## Database migrations

Never edit a migration that has already been applied in a shared environment. Add a new ordered migration. Keep data invariants in schema constraints/indexes where appropriate and keep application behavior compatible with deployment ordering.

## API changes

Update backend implementation, `packages/api-contract`, client types/usages, tests and `docs/api.md`/flow docs in the same PR. Do not let OpenAPI/contract drift become a separate undocumented task.

## Architecture changes

Update `ARCHITECTURE.md` and the relevant `docs/architecture/*.md`. If the change is a durable design choice with alternatives/tradeoffs, add an ADR under `docs/decisions/`.

## AI/Codex changes

Read root `AGENTS.md` and the app-local `AGENTS.md` before editing. Preserve dependency direction and security boundaries. Prefer small reviewable PRs. Documentation must describe actual merged behavior, not intended future behavior.
