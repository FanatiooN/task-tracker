import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { Task } from '@/types'
import * as tasksApi from '@/api/tasks'
import { AuthError } from '@/api/client'
import router from '@/router'

export const useTasksStore = defineStore('tasks', () => {
  const tasks = ref<Task[]>([])
  const nextPageToken = ref('')
  const loading = ref(false)

  async function fetchTasks(reset = false) {
    if (reset) {
      tasks.value = []
      nextPageToken.value = ''
    }
    loading.value = true
    try {
      const res = await tasksApi.listTasks(
        nextPageToken.value,
        'TASK_STATUS_IN_PROGRESS',
      )
      const incoming = res.tasks ?? []
      tasks.value = reset ? incoming : [...tasks.value, ...incoming]
      nextPageToken.value = res.nextPageToken ?? ''
    } catch (e) {
      if (e instanceof AuthError) {
        router.push('/auth')
      }
    } finally {
      loading.value = false
    }
  }

  async function loadMore() {
    if (!nextPageToken.value || loading.value) return
    await fetchTasks(false)
  }

  async function addTask(title: string, description?: string) {
    const task = await tasksApi.createTask(title, description)
    tasks.value = [task, ...tasks.value]
  }

  async function toggleDone(id: string) {
    const idx = tasks.value.findIndex((t) => t.id === id)
    if (idx === -1) return

    const task = tasks.value[idx]
    const newStatus: Task['status'] =
      task.status === 'TASK_STATUS_DONE' ? 'TASK_STATUS_IN_PROGRESS' : 'TASK_STATUS_DONE'

    tasks.value.splice(idx, 1, { ...task, status: newStatus })

    try {
      await tasksApi.updateTask(id, { taskStatus: newStatus })
    } catch {
      tasks.value.splice(idx, 1, task)
    }
  }

  async function removeTasks(ids: string[]) {
    await tasksApi.deleteTasks(ids)
    tasks.value = tasks.value.filter((t) => !ids.includes(t.id))
  }

  return { tasks, nextPageToken, loading, fetchTasks, loadMore, addTask, toggleDone, removeTasks }
})
