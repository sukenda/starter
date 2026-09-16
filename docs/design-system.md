# Design System

The starter uses one semantic design language across Vue and Flutter while allowing each platform to remain native.

## Source of truth

`packages/design-tokens/tokens.json` is the canonical vocabulary for color, spacing, radius, typography, elevation and motion. Platform adapters expose equivalent values as CSS variables and Dart constants/themes.

Do not use raw brand colors, arbitrary spacing or one-off radii inside product features when a semantic token exists.

## Principles

1. Semantic over literal: use `surface`, `text-secondary`, `danger`, etc. rather than implementation color names.
2. Accessible by default: visible focus, sufficient contrast, touch targets, labels and native semantics are mandatory.
3. Platform-native behavior: Vue and Flutter should feel related, not pixel-identical.
4. Small primitives first: button, card, page header, fields, feedback and navigation primitives precede complex product widgets.
5. Composition over variants: avoid components with large boolean prop surfaces.
6. Product features consume the design system; they do not redefine it.

## Web boundary

Reusable primitives live under `apps/frontend/src/shared/ui`. Global semantic variables live under `src/styles`. Feature-specific components remain inside their feature.

## Mobile boundary

Theme and reusable visual primitives live under `apps/mobile/lib/design_system`. Feature presentation should prefer `Theme.of(context)` and design-system constants instead of literal styling.

## Component roadmap

Foundation: tokens, theme, button, card and page header. Next: input/field, badge, alert, loading/empty/error states, dialog, table primitives, app navigation and form patterns. Complex components are added only when at least one real feature needs them.

## AI rules

When AI changes UI it must reuse existing tokens/primitives first, keep authorization/business logic outside visual components, preserve keyboard/screen-reader semantics on web, preserve Material semantics on mobile, and update this document when introducing a new cross-platform visual convention.
