export interface ApiErrorPayload {
  code: string
  message: string
  request_id?: string
  details?: unknown
}

export class ApiError extends Error {
  readonly code: string
  readonly status: number
  readonly requestId: string | undefined
  readonly details: unknown

  constructor(status: number, payload: ApiErrorPayload) {
    super(payload.message)
    this.name = 'ApiError'
    this.code = payload.code
    this.status = status
    this.requestId = payload.request_id
    this.details = payload.details
  }
}
