# System Architecture

## Purpose

Starter is a reusable full-stack monorepo. It deliberately keeps backend, browser frontend, and mobile application independently buildable while sharing contracts and design vocabulary at repository level.

```text
Browser / Vue SPA ───────┐
                         │ HTTPS / JSON
Flutter mobile ──────────┼────────> Go API ─────> MariaDB / MySQL
                         │             │
Health probes ───────────┘             └─ migrations + persistence

packages/api-contract  ---> HTTP contract reference
packages/design-tokens ---> visual source values mirrored by web/mobile adapters
```

## Repository boundaries

```text
apps/backend    Go HTTP API, authentication, authorization, admin domain, database
apps/frontend   Vue SPA for login/dashboard/admin management
apps/mobile     Flutter native client for login/authenticated home foundation
packages/       cross-application contracts/tokens, not runtime application code
docs/           architecture, flows, decisions, development and operations
.github/        independent CI pipelines for backend/frontend/mobile
```

Each application owns its framework-specific dependencies and build tooling. The root `Makefile` orchestrates common developer tasks instead of introducing a monorepo framework.

## Runtime trust boundaries

The Go API is the security boundary. A browser/mobile permission check is only a UX optimization; protected backend routes must enforce authentication and required permission server-side.

Browser and native clients intentionally use different refresh-token transports. Browser access tokens live in memory while refresh tokens use HttpOnly cookies plus a CSRF token. Flutter stores native access/refresh tokens in secure storage and sends the refresh token in the JSON refresh request.

## Request lifecycle

A normal protected request enters the API, receives request/security middleware, authenticates the bearer access token, resolves the session/user and permissions, checks route permission, invokes a domain handler/service/repository path, accesses MariaDB when needed, and returns the standard JSON envelope.

## Application architecture direction

Backend dependencies point inward from HTTP handlers to services to repository/store interfaces. Frontend is feature-oriented with shared transport/design primitives. Mobile is feature-oriented with Riverpod application state, repositories for remote/session behavior, secure storage for credentials, and GoRouter for navigation.

## Shared contracts

`packages/api-contract` is the intended canonical HTTP contract surface. Runtime implementation and contract must be changed together. `packages/design-tokens` is the canonical visual vocabulary; Vue CSS variables/components and Flutter token/theme adapters consume equivalent values.

## Current product flows

Implemented baseline flows are browser login/session restoration/logout, mobile login/session restoration/refresh/logout, dashboard entry, and browser admin management of users, roles, permissions, and menus. Detailed sequences live under `docs/flows/`.

## Current production-readiness boundary

The repository contains foundational CI, configuration validation, migrations, request IDs, structured access logging, security headers, health/readiness endpoints, authentication/RBAC, browser CSRF session transport, native secure storage and refresh retry foundations. Deployment images, full observability, release automation, backups/restore, rate limiting, complete integration coverage, admin pagination/audit, and other hardening items remain explicit follow-up work until implemented.
