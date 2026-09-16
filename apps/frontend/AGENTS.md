# Frontend AI Rules

Applies to `apps/frontend`.

## Stack

- Vue 3 Composition API
- TypeScript strict mode
- Vite
- Vue Router
- TanStack Query for server state
- Pinia for true global client state
- Zod for boundary/form schemas
- shadcn-vue/Tailwind may be added as the design-system layer

## Structure

```text
src/
  app/          app bootstrap/providers
  components/   reusable UI/layout components
  features/     business features
  composables/  genuinely cross-feature composables
  lib/          API client and infrastructure
  router/       route definitions/guards
  stores/       global client state only
  types/        shared TypeScript types
  utils/        small pure utilities
```

## Rules

- Use `<script setup lang="ts">`.
- Keep TypeScript strict; avoid `any` unless interacting with an untyped external boundary and document why.
- Do not call `fetch` directly from Vue components.
- Server/API state belongs in TanStack Query, not Pinia.
- Do not duplicate server state in Pinia.
- Keep feature-specific code under `src/features/<feature>`.
- Reusable visual primitives belong in `src/components/ui`.
- Prefer composition over giant configurable components.
- Pages coordinate features; they should not contain large business workflows.
- Validate external/untrusted payloads at meaningful boundaries when runtime validation is required.
- Use route-level lazy loading for application pages.

## API

Public request/response behavior must match `/packages/api-contract/openapi.yaml`. Centralize base URL, headers, auth handling, error normalization, and request cancellation in `src/lib/api`.

## Testing

- Unit test pure/composable behavior with Vitest.
- Component test meaningful interactions rather than implementation details.
- Add E2E tooling only when an actual end-to-end workflow is introduced.
