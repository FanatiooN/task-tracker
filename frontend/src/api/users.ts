import { apiFetch } from './client'
import type { User } from '@/types'

export const getUser = (id: string) => apiFetch<User>(`/users/${id}`)

export const updateUser = (id: string, name: string) =>
  apiFetch<User>(`/users/${id}`, {
    method: 'PUT',
    body: JSON.stringify({ name }),
  })

export const deleteUser = (id: string) =>
  apiFetch<void>(`/users/${id}`, { method: 'DELETE' })
