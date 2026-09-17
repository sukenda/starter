# Mobile Architecture

## Stack

The native client lives in `apps/mobile` and uses Flutter, Riverpod, GoRouter, Dio, and platform secure storage.

## Source layout

```text
lib/main.dart          process/application bootstrap
lib/app/               MaterialApp/router/application composition
lib/core/              config, networking and cross-feature infrastructure
lib/design_system/     Flutter tokens/theme/components
lib/features/auth/     auth application/data/domain/presentation
lib/features/home/     authenticated home presentation
```

Features should own their application/data/domain/presentation behavior. `core` is reserved for infrastructure used by multiple features.

## Authentication state

`AuthController` is Riverpod application state. It restores a session during startup, performs login/logout, and exposes the current `SessionUser`. `AuthRepository` owns HTTP/session orchestration. `TokenStorage` owns secure persistence of native access and refresh tokens.

## Native token lifecycle

Login identifies the client as mobile and receives access + refresh tokens in JSON. Both are stored using platform secure storage. On startup, the repository reads the access token and calls `/auth/me`; if access is absent or receives `401`, it rotates the refresh token and retries `/auth/me`. Invalid refresh credentials clear local credentials. Transient network errors are not intentionally converted into an invalid-session result.

Refresh calls inside the repository are serialized with an in-flight Future so concurrent callers share one refresh rotation. This is required because refresh tokens are single-use/rotating server-side.

## Authenticated networking

The authenticated Dio provider attaches the stored bearer access token to non-auth requests. A `401` can trigger the serialized refresh path and retry the original request once with the rotated access token. A retry marker prevents loops. Login and refresh endpoints are excluded from automatic auth retry.

Feature repositories that call protected APIs should use the authenticated client rather than implementing token injection/refresh themselves.

## Routing

GoRouter redirects based on authentication state. Unauthenticated users are kept on login; authenticated users enter the home/application route. Startup restoration is part of auth state, so routing must not invent a second session source of truth.

## Design system

Flutter design tokens/theme mirror the shared design vocabulary. Screens should use the centralized theme/tokens and reusable components instead of hardcoded visual values where a token exists.

## Testing

The current suite covers config/network/secure-store smoke primitives. Focused refresh/retry regression coverage is the next testing scope; do not describe it as complete until merged.

## Extension pattern

For a new mobile feature, create a feature directory, put remote persistence behind a repository/provider, use `authenticatedDioProvider` for protected HTTP calls, keep Riverpod state at the application boundary, use GoRouter for navigation, reuse design-system values, and add tests for state transitions and network edge cases.
