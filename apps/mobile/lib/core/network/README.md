# Mobile network boundary

Feature repositories obtain Dio through Riverpod rather than constructing clients in widgets. This boundary owns base URL, transport timeouts, shared headers, and normalization of server error envelopes. Authentication interceptors are added in the authentication PR after token lifecycle is defined.
