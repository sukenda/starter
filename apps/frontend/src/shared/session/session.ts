export interface SessionUser {
  id: string
  name: string
  email: string
}

export interface SessionSnapshot {
  user: SessionUser | null
  permissions: readonly string[]
}

export const anonymousSession: SessionSnapshot = {
  user: null,
  permissions: [],
}
