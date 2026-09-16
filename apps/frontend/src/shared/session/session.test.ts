import { describe, expect, it } from 'vitest'
import { anonymousSession } from './session'

describe('anonymousSession', () => {
  it('starts without identity or permissions', () => {
    expect(anonymousSession.user).toBeNull()
    expect(anonymousSession.permissions).toEqual([])
  })
})
