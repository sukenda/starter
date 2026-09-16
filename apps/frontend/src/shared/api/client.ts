import { ApiError, type ApiErrorPayload } from './errors'

interface Envelope<T> {
  data?: T
  error?: ApiErrorPayload
  meta?: unknown
}

const baseUrl = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080'

export async function apiRequest<T>(path: string, init?: RequestInit): Promise<T> {
  const response = await fetch(`${baseUrl}${path}`, {
    ...init,
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
      ...init?.headers,
    },
  })

  const envelope = (await response.json()) as Envelope<T>
  if (!response.ok || envelope.error) {
    throw new ApiError(response.status, envelope.error ?? {
      code: 'unexpected_response',
      message: 'The server returned an unexpected response.',
    })
  }

  return envelope.data as T
}
