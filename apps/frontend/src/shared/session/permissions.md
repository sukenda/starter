# Permission semantics

Permission codes are opaque backend-issued capabilities such as `users.read`. Frontend helpers perform exact-code checks only. Do not infer authorization from role names or menu presence, and never treat client-side permission checks as the system security boundary.
