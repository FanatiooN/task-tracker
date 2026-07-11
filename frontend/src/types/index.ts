export type TaskStatus =
  | 'TASK_STATUS_IN_PROGRESS'
  | 'TASK_STATUS_DONE'
  | 'TASK_STATUS_CANCELLED'
  | 'TASK_STATUS_UNSPECIFIED'

export interface Task {
  id: string
  title: string
  description?: string
  status: TaskStatus
  user_id: string
  created_at: string
  updated_at: string
}

export interface User {
  id: string
  name: string
  email?: string
  created_at?: string
  updated_at?: string
}

export interface Tokens {
  access_token: string
  refresh_token: string
}

export interface ListTasksResponse {
  tasks: Task[]
  nextPageToken?: string
}
