import { apiFetch } from './client'
import type { Tokens } from '@/types'

export const login = (email: string, password: string) =>
  apiFetch<Tokens>('/login', {
    method: 'POST',
    body: JSON.stringify({ email, password }),
  })

export const register = (name: string, email: string, password: string) =>
  apiFetch<Tokens>('/register', {
    method: 'POST',
    body: JSON.stringify({ name, email, password }),
  })

export const logout = (token: string) =>
  apiFetch<void>('/logout', {
    method: 'POST',
    body: JSON.stringify({ token }),
  })

export const refresh = (token: string) =>
  apiFetch<Tokens>('/refresh', {
    method: 'POST',
    body: JSON.stringify({ token }),
  })

export const loginWithTelegram = (idToken: string) =>
  apiFetch<Tokens>('/login/telegram', {
    method: 'POST',
    body: JSON.stringify({ id_token: idToken }),
  })
