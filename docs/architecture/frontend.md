# Frontend Architecture

## Stack

The browser application lives in `apps/frontend` and uses Vue 3 + Vite + TypeScript. Pinia owns application/session state, Vue Router owns navigation, and TanStack Query owns server-state fetching/mutations for admin data.

## Source layout

```text
src/main.ts       application bootstrap
src/App.vue       router outlet/root
src/router/       route definitions and auth/permission guards
src/layouts/      authenticated application shell
src/features/     feature-owned pages/components/state/API usage
src/shared/api/   HTTP transport, errors and shared API primitives
src/shared/       reusable non-feature concerns
src/styles/       global/design-token adapters
```

The application is feature-first. Domain-specific code belongs under its feature; cross-feature transport or UI primitives belong under `shared`.

## Session model

Browser access tokens are memory-only. They are not persisted to localStorage/sessionStorage. The backend stores the refresh token in an HttpOnly cookie. A readable CSRF cookie/value is paired with refresh requests. On reload, the SPA can restore the session by asking the backend to refresh from the cookie and then loading `/auth/me`.

Concurrent browser refresh is serialized so several failed requests do not race the single-use refresh rotation. Authentication endpoints and transport retry logic must avoid recursive refresh loops.

## Routing and authorization UX

Route guards restore/check authentication before entering protected screens. Routes/navigation can declare permission requirements. The admin layout renders only allowed navigation. This is UX only: the backend must independently enforce the same permission.

## Server state

Admin users/roles/permissions/menus are fetched through typed API modules and TanStack Query. Mutations invalidate the relevant query keys so list state is refreshed from the server rather than manually patched in multiple places.

## Admin UI

The current baseline provides login, dashboard, users, roles, and menus screens. Dialog components support create/update workflows. Users support multiple roles; roles support permission assignment; menus support parent relationships and permission metadata.

## Design system

Reusable `Ds*` components implement common inputs, buttons, cards, badges, state displays, page headers, dialogs and selects. Visual values originate from the shared design-token vocabulary. Feature pages should compose these components rather than introduce unrelated one-off visual systems.

## Error boundary

`shared/api` normalizes HTTP/API/network failures into frontend errors. Callers should handle expected mutation/query errors at the feature boundary. Transport details should not be duplicated in individual pages.

## Extension pattern

A new frontend feature should get a directory under `features`, keep API/query logic close to that feature, reuse the shared API client, declare routing/permission metadata centrally, compose design-system primitives, and add focused tests for state/transport/critical UI behavior. Any backend contract change must be reflected in the API contract and frontend types together.
