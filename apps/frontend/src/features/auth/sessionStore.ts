import { defineStore } from 'pinia'
import { apiRequest } from '@/shared/api/client'

export interface CurrentUser {
  id: string
  email: string
  name: string
  permissions: string[]
}
interface Tokens {
  access_token: string
  token_type: string
  access_expires_at: string
  refresh_expires_at: string
}
interface LoginResponse {
  user: Omit<CurrentUser, 'permissions'>
  tokens: Tokens
}

function csrfToken() {
  const item = document.cookie.split('; ').find((value) => value.startsWith('starter_csrf='))
  return item ? decodeURIComponent(item.slice('starter_csrf='.length)) : ''
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as CurrentUser | null,
    accessToken: null as string | null,
    initialized: false,
    refreshing: null as Promise<boolean> | null,
  }),
  getters: {
    isAuthenticated: (state) => Boolean(state.user && state.accessToken),
    hasPermission: (state) => (code: string) => Boolean(state.user?.permissions.includes(code)),
  },
  actions: {
    async login(email: string, password: string) {
      const response = await apiRequest<LoginResponse>('/api/v1/auth/login', {
        method: 'POST',
        credentials: 'include',
        body: JSON.stringify({ email, password, client: 'web' }),
      })
      this.accessToken = response.tokens.access_token
      await this.loadMe()
    },
    async refresh() {
      if (this.refreshing) return this.refreshing
      this.refreshing = (async () => {
        try {
          const response = await apiRequest<Tokens>('/api/v1/auth/refresh', {
            method: 'POST',
            credentials: 'include',
            headers: { 'X-CSRF-Token': csrfToken() },
            body: '{}',
          })
          this.accessToken = response.access_token
          return true
        } catch {
          this.clear()
          return false
        } finally {
          this.refreshing = null
        }
      })()
      return this.refreshing
    },
    async loadMe() {
      if (!this.accessToken && !(await this.refresh())) {
        this.initialized = true
        return
      }
      try {
        this.user = await apiRequest<CurrentUser>('/api/v1/auth/me', {
          headers: { Authorization: `Bearer ${this.accessToken}` },
        })
      } catch {
        if (await this.refresh()) {
          try {
            this.user = await apiRequest<CurrentUser>('/api/v1/auth/me', {
              headers: { Authorization: `Bearer ${this.accessToken}` },
            })
          } catch {
            this.clear()
          }
        } else {
          this.clear()
        }
      } finally {
        this.initialized = true
      }
    },
    async logout() {
      if (this.accessToken) {
        try {
          await apiRequest('/api/v1/auth/logout', {
            method: 'POST',
            credentials: 'include',
            headers: { Authorization: `Bearer ${this.accessToken}` },
          })
        } catch {
          // Local logout still completes when the server is unavailable.
        }
      }
      this.clear()
    },
    clear() {
      this.user = null
      this.accessToken = null
      this.initialized = true
    },
  },
})
