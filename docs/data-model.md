# Data Model

## Core identity and RBAC

```text
users ──< user_roles >── roles ──< role_permissions >── permissions
  │
  └──< auth_sessions

menus ── parent_id ──> menus
  └── permission_id/permission relation ──> permissions
```

The exact physical schema is defined by ordered SQL migrations in `apps/backend/migrations`; this document describes semantic ownership and relationships.

## Users

A user is an authenticated identity with public ID, email/name/status and password credential data. A user can have zero or more roles through `user_roles`. Public API DTOs must not expose internal database identifiers or password hashes.

## Roles

Roles group permissions and are assignable to many users. Role code/name identify the role at the application boundary. A role can contain many permissions through `role_permissions`.

## Permissions

Permissions are atomic authorization capabilities. Backend protected routes check these capabilities. They may also be referenced by menu metadata for navigation visibility.

## user_roles

Join relation implementing many-to-many users-to-roles. This is why application code must not assume a single `role_id` on a user.

## role_permissions

Join relation implementing many-to-many roles-to-permissions. Role update flows replace/update this relation transactionally so partially applied permission sets are avoided.

## auth_sessions

Server-side session records connect issued opaque credentials to a user and expiration/revocation/rotation state. Raw access/refresh secrets are not stored; hashes are persisted. Refresh rotation uses conditional state transition semantics so a credential cannot be successfully consumed more than once.

## menus

Menus describe hierarchical application navigation. A menu can have a parent menu and an optional permission requirement. Parent references must not form self-links/cycles. Menus do not replace route authorization.

## IDs

Public HTTP resources use stable public string IDs. Internal SQL primary/foreign key representation is a repository concern and must not leak into API DTOs.

## Migration policy

Schema is changed only through a new ordered migration. Migrations should establish foreign keys, uniqueness and indexes required to preserve invariants. Application changes that depend on a new schema must ship the migration and corresponding contract/docs together.
