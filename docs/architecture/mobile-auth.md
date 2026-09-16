# Mobile Authentication Architecture

The Flutter application keeps authentication behind feature and infrastructure boundaries:

- `AuthRepository` owns backend authentication calls.
- `TokenStorage` owns credential persistence and uses platform secure storage.
- `AuthController` owns session lifecycle for the UI.
- `GoRouter` redirects between public and authenticated routes based on session state.
- Presentation widgets never read or write raw tokens.

## Token handling

Access and refresh tokens are stored with `flutter_secure_storage`, not shared preferences. On application startup the controller restores the access token and validates it through `/api/v1/auth/me`. Invalid credentials are removed.

This slice intentionally does not implement automatic refresh-token rotation. That should be introduced together with a centralized authenticated Dio interceptor so concurrent 401 responses can be serialized into one refresh operation rather than producing refresh races.

## Logout

Logout attempts server-side session revocation and always clears local credentials, including when the network request fails.

## UI

Login and Home use Material 3 primitives only. Visual styling is provisional until the shared Vue/Flutter design-system phase.
