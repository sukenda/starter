# Pull request stack

Keep milestone branches connected in dependency order.

Current stack:

```text
main
  <- feat/production-monorepo-starter
       <- feat/engineering-foundation
            <- feat/auth-rbac-foundation
                 <- feat/frontend-admin-shell
                      <- feat/mobile-auth-home
```

Do not merge a child before its parent. When a parent lands, rebase or retarget the direct child onto the parent's new destination before merging it. New work branches from the latest active milestone so AI agents always receive the complete current architecture.
