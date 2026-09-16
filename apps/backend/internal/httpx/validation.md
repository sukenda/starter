# Request validation

Transport validation reports field violations using stable `validation_failed` errors. Validation of syntax/shape belongs at the HTTP boundary; business invariants remain in application/domain logic. Feature handlers should not leak validator-library-specific error structures into the public contract.
