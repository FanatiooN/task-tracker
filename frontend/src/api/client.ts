export class AuthError extends Error {
  constructor() {
    super('Unauthorized')
    this.name = 'AuthError'
  }
}

export async function apiFetch<T>(path: string, options: RequestInit = {}): Promise<T> {
  const token = localStorage.getItem('access_token')

  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
    ...(options.headers as Record<string, string>),
  }

  if (token) {
    headers['Authorization'] = `Bearer ${token}`
  }

  const res = await fetch(path, { ...options, headers })

  if (res.status === 204) return undefined as T

  const data = await res.json().catch(() => null)

  if (res.status === 401) {
    localStorage.clear()
    throw new AuthError()
  }

  if (!res.ok) {
    const message = data?.error ?? data?.message ?? `HTTP ${res.status}`
    throw new Error(message)
  }

  return data as T
}
