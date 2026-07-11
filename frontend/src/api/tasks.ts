import { apiFetch } from './client'
import type { Task, ListTasksResponse } from '@/types'

export const listTasks = (pageToken = '', status = '') => {
  const params = new URLSearchParams({ pageSize: '20' })
  if (status) params.set('taskStatus', status)
  if (pageToken) params.set('pageToken', pageToken)
  return apiFetch<ListTasksResponse>(`/tasks?${params}`)
}

export const createTask = (title: string, description?: string) =>
  apiFetch<Task>('/tasks', {
    method: 'POST',
    body: JSON.stringify({ title, ...(description ? { description } : {}) }),
  })

export const updateTask = (
  id: string,
  patch: { title?: string; description?: string; taskStatus?: string },
) =>
  apiFetch<Task>(`/tasks/${id}`, {
    method: 'PUT',
    body: JSON.stringify(patch),
  })

export const deleteTasks = (ids: string[]) =>
  apiFetch<void>('/tasks', {
    method: 'DELETE',
    body: JSON.stringify({ ids }),
  })
