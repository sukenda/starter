import { describe, expect, it } from 'vitest'
import { ApiError } from './errors'

describe('ApiError', () => {
  it('preserves normalized server metadata', () => {
    const error = new ApiError(401, {
      code: 'unauthorized',
      message: 'Authentication is required.',
      request_id: 'request-1',
    })

    expect(error).toBeInstanceOf(Error)
    expect(error.status).toBe(401)
    expect(error.code).toBe('unauthorized')
    expect(error.requestId).toBe('request-1')
  })
})
