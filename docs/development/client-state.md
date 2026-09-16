# Client state ownership

Frontend server state belongs in TanStack Query; Pinia is for durable client/session state that is not simply a cache of server resources. Flutter uses Riverpod for dependency composition and application state. Keep ephemeral form/widget state local unless multiple screens genuinely need it.

Do not duplicate server collections into global stores without a concrete reason.
