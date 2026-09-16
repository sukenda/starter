import { describe, expect, it } from 'vitest'
import { hasPermission } from './permissions'

describe('hasPermission', () => {
  it('checks exact permission codes', () => {
    expect(hasPermission(['users.read'], 'users.read')).toBe(true)
    expect(hasPermission(['users.read'], 'users.write')).toBe(false)
  })
})
