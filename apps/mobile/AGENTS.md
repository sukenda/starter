# Mobile AI Rules

Applies to `apps/mobile`.

## Stack

- Flutter / Dart
- Riverpod
- GoRouter
- Dio
- Freezed + json_serializable when immutable models/code generation provide value
- flutter_secure_storage when authentication secrets are introduced

## Structure

```text
lib/
  app/          bootstrap, router, top-level providers
  core/         network, config, errors, storage, theme
  shared/       reusable widgets/models that are truly cross-feature
  features/     business features
```

For substantial features, prefer:

```text
features/<feature>/
  data/
  domain/
  presentation/
```

Do not force all three layers for trivial features.

## Rules

- Widgets render state and emit user intent; keep networking/business logic out of widgets.
- Access remote APIs through repositories/data sources using the shared Dio clients.
- Use `dioProvider` for public/auth transport and `authenticatedDioProvider` for protected application APIs.
- Do not manually attach bearer tokens or implement per-feature refresh/retry logic; the authenticated client owns that boundary.
- Use Riverpod for dependency injection and application state.
- Use GoRouter for navigation and route guards/redirects.
- Model loading, data, empty, and error states explicitly.
- Prefer immutable state.
- Do not use `BuildContext` inside repositories/services.
- Cancel/ignore obsolete asynchronous work where appropriate.
- Keep platform-specific code behind narrow abstractions.
- Never persist access credentials in plain preferences.

## API

HTTP behavior must match `/packages/api-contract/openapi.yaml`. Normalize transport failures into application-level failures before presentation code handles them.

Protected requests must use `authenticatedDioProvider`. It attaches the current access token, performs at most one refresh-and-retry after a 401, and relies on `AuthRepository` to serialize refresh rotation. Authentication endpoints stay on the base `dioProvider` so refresh requests cannot recursively trigger refresh handling.

## Testing

- Unit test domain/controller/repository behavior.
- Widget test important UI interactions.
- Use `integration_test` for critical cross-screen workflows when those workflows exist.
