# Database boundary

The shared database package owns connection lifecycle and health only. Feature SQL belongs with the owning feature repository. Do not create a generic repository abstraction or hide `database/sql` behind interfaces unless a concrete application/test seam requires it. Transactions should be controlled at the use-case boundary when multiple writes must commit atomically.
