import { describe, expect, it } from 'vitest'
import { ApiError } from './errors'

describe('ApiError', () => {
  it('preserves stable API error metadata', () => {
    const error = new ApiError(422, {
      code: 'validation_failed',
      message: 'Request is invalid.',
      request_id: 'request-1',
    })

    expect(error.status).toBe(422)
    expect(error.code).toBe('validation_failed')
    expect(error.requestId).toBe('request-1')
  })
})
