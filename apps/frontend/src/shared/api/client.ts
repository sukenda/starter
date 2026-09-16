import { ApiError, type ApiErrorPayload } from './errors'

interface Envelope<T> { data?: T; error?: ApiErrorPayload; meta?: unknown }
const baseUrl = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080'

export async function apiRequest<T>(path: string, init?: RequestInit): Promise<T> {
  const headers = new Headers(init?.headers)
  headers.set('Accept', 'application/json')
  if (init?.body && !(init.body instanceof FormData) && !headers.has('Content-Type')) headers.set('Content-Type', 'application/json')
  let response: Response
  try { response = await fetch(`${baseUrl}${path}`, { ...init, headers }) }
  catch { throw new ApiError(0, { code: 'network_error', message: 'Unable to reach the server. Check your connection and try again.' }) }
  if (response.status === 204) return undefined as T
  const contentType = response.headers.get('content-type') ?? ''
  let envelope: Envelope<T> = {}
  if (contentType.includes('application/json')) {
    try { envelope = await response.json() as Envelope<T> }
    catch { throw new ApiError(response.status, { code: 'invalid_response', message: 'The server returned an invalid JSON response.' }) }
  }
  if (!response.ok || envelope.error) throw new ApiError(response.status, envelope.error ?? { code: 'unexpected_response', message: 'The server returned an unexpected response.' })
  return envelope.data as T
}
