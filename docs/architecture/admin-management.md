# Admin Management

The admin API is an authenticated, permission-protected boundary for maintaining users, roles and navigation menus.

## Authorization

Every route first authenticates the database-backed session and then requires an explicit permission. The backend is authoritative; future frontend/mobile permission checks only control presentation.

Permissions use resource/action codes: `users.read`, `users.write`, `roles.read`, `roles.write`, `menus.read`, `menus.write`.

## User management

Users can hold multiple roles through `user_roles`. Passwords are only accepted when creating a user in this first slice and are stored using the existing Argon2id password service. User status can disable authentication without deleting audit-relevant records.

## Role management

Roles can hold multiple permissions. Role and permission assignments are updated transactionally so partial assignments cannot be committed.

## Menu management

Menus are navigation metadata, not an authorization mechanism. They support parent/child relationships, route, icon, ordering, active state and an optional permission code used by clients to decide visibility. API authorization remains independent of menus.

## Next slice

The frontend admin shell consumes these APIs for login, dashboard, user management, role management and menu management. A later design-system PR will replace provisional presentation with shared semantic tokens and polished components.
