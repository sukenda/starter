# Frontend Admin Architecture

The Vue administration application uses feature boundaries for auth, dashboard and administration. Shared API/session infrastructure stays outside product features.

## Session

Access and refresh tokens are intentionally held in Pinia memory in this slice, not localStorage. This avoids creating a persistent XSS token target. A browser refresh requires re-authentication until a cookie-based web refresh transport is introduced. Mobile uses its own secure-storage boundary.

## Authorization

Route guards and menu visibility use permissions returned by `/api/v1/auth/me`. They are presentation controls only. Backend permission middleware remains authoritative.

## Data fetching

Server state is loaded through TanStack Query and the shared API client. Product pages do not call `fetch` directly.

## Presentation

The current CSS is deliberately provisional. It establishes layout and usability without defining the final visual language. Shared design tokens and reusable Vue/Flutter components will be introduced in the dedicated design-system phase.
