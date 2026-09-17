# Application Flows

This document describes the major implemented end-to-end flows and which layer owns each step.

## Browser login

```text
LoginPage
  -> auth store
  -> POST /api/v1/auth/login
  -> backend auth handler/service
  -> verify Argon2id password
  -> create auth session + access/refresh credentials
  -> access token returned to SPA memory
  -> refresh token set as HttpOnly cookie
  -> CSRF value/cookie established
  -> current user/permissions loaded
  -> router enters dashboard
```

The refresh credential must never be copied into browser persistent JavaScript storage.

## Browser reload / restore

```text
SPA starts with no access token
  -> session restore requests /auth/refresh
  -> browser automatically sends HttpOnly refresh cookie
  -> SPA sends matching CSRF header/value
  -> backend validates CSRF + rotates refresh credential
  -> new access token kept in memory
  -> GET /auth/me
  -> Pinia session restored
  -> protected route resolves
```

Concurrent refresh is serialized. Invalid refresh results in an unauthenticated browser state.

## Browser protected admin request

```text
Admin page/query
  -> shared API client + bearer access token
  -> backend auth middleware
  -> session/user lookup
  -> effective permissions
  -> RequirePermission
  -> admin handler
  -> admin service
  -> repository / transaction
  -> MariaDB
  -> JSON envelope
  -> TanStack Query cache
  -> UI
```

A hidden menu/button does not grant or revoke access; backend permission enforcement is authoritative.

## User management

The browser loads users and role options. Create/update submits public DTOs to the admin API. The backend validates the use case, resolves role public IDs, performs user/role changes transactionally where needed, and returns the public representation. The frontend invalidates user-related query state and reloads server truth.

## Role and permission management

The browser loads roles and permissions. A role create/update carries permission public identifiers. Backend service/repository updates role data and the role-permission join relationship in a transaction. Effective user permissions are later derived through `user_roles -> roles -> role_permissions -> permissions`.

## Menu management

Menus are navigation metadata with optional parent and permission requirement. The browser sends public parent IDs. Backend resolves them internally and prevents invalid self/cyclic hierarchy. Menu permission metadata drives navigation visibility but protected endpoints still enforce permission independently.

## Mobile login

```text
LoginPage
  -> AuthController
  -> AuthRepository.login
  -> POST /auth/login client=mobile
  -> JSON access + refresh tokens
  -> TokenStorage / platform secure storage
  -> SessionUser in Riverpod
  -> GoRouter redirects to home
```

## Mobile startup restore

```text
app starts
  -> AuthController.build
  -> AuthRepository.restore
  -> read secure access token
  -> GET /auth/me
       success -> authenticated
       401 -> refresh()
               -> read secure refresh token
               -> POST /auth/refresh JSON
               -> rotate + save both tokens
               -> retry /auth/me once
  -> Riverpod session state
  -> GoRouter redirect
```

If no access token exists but a refresh token exists, restore attempts refresh first. Invalid refresh clears credentials.

## Mobile protected request / automatic retry

```text
feature repository
  -> authenticated Dio
  -> attach bearer token
  -> protected API
       success -> return
       401 -> serialized AuthRepository.refresh
              -> save rotated tokens
              -> retry original request once
       second failure -> propagate error
```

Auth endpoints are excluded from this interceptor path and a request retry marker prevents infinite loops.

## Logout

Browser logout invalidates the backend session and clears browser refresh/CSRF cookies plus in-memory session state. Mobile logout attempts server logout with the access token and always clears local secure credentials even when the network call fails.

## Backend health/readiness

Infrastructure probes call health/readiness endpoints. Liveness should answer whether the process is running; readiness includes dependency readiness such as database connectivity. These endpoints must not depend on an authenticated user session.
