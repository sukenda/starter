# Architecture guardrails

- Prefer vertical feature modules over global technical buckets.
- Domain/application logic must not depend on Fiber, Vue, Flutter, Dio, or browser APIs.
- Introduce interfaces at real substitution/test boundaries, not preemptively.
- Avoid generic repositories, base services, and catch-all utility packages.
- Public API behavior is contract-first and versioned.
- Backend authorization is authoritative; client guards are UX only.
- Cross-cutting infrastructure is centralized and small.
- Optimize for code an engineer or AI agent can locate and understand from the repository tree.
