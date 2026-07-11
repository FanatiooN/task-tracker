<template>
  <div class="tasks-page">
    <div class="page-header">
      <h2 class="page-title">Tasks</h2>
      <div class="header-actions">
        <button
          v-if="selectedIds.size > 0"
          class="btn-danger"
          @click="handleDelete"
          :disabled="store.loading"
        >
          Удалить ({{ selectedIds.size }})
        </button>
        <button class="btn-secondary" @click="toggleCreateForm">
          {{ showCreate ? 'Отмена' : 'Новая задача' }}
        </button>
      </div>
    </div>

    <div v-if="showCreate" class="create-form">
      <form @submit.prevent="handleCreate">
        <input
          v-model="newTitle"
          class="create-input"
          placeholder="Название задачи"
          required
          ref="titleInput"
          autofocus
        />
        <input
          v-model="newDesc"
          class="create-input"
          placeholder="Описание (необязательно)"
        />
        <p v-if="createError" class="error-msg">{{ createError }}</p>
        <div class="create-actions">
          <button type="submit" class="btn-primary" :disabled="store.loading || !newTitle.trim()">
            {{ store.loading ? 'Создаём...' : 'Создать' }}
          </button>
          <button type="button" class="btn-ghost" @click="toggleCreateForm">Отмена</button>
        </div>
      </form>
    </div>

    <div class="task-list">
      <div v-if="store.loading && store.tasks.length === 0" class="state-msg">
        Загрузка...
      </div>

      <div v-else-if="store.tasks.length === 0" class="state-msg">
        Нет активных задач
      </div>

      <TaskItem
        v-for="task in store.tasks"
        :key="task.id"
        :task="task"
        :selected="selectedIds.has(task.id)"
        @toggle-select="toggleSelect(task.id)"
        @toggle-done="store.toggleDone(task.id)"
      />
    </div>

    <div v-if="store.nextPageToken" class="load-more">
      <button class="btn-ghost" @click="store.loadMore()" :disabled="store.loading">
        {{ store.loading ? 'Загрузка...' : 'Загрузить ещё' }}
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, nextTick } from 'vue'
import { useTasksStore } from '@/stores/tasks'
import TaskItem from '@/components/TaskItem.vue'

const store = useTasksStore()
const selectedIds = reactive(new Set<string>())
const showCreate = ref(false)
const newTitle = ref('')
const newDesc = ref('')
const createError = ref('')
const titleInput = ref<HTMLInputElement>()

onMounted(() => store.fetchTasks(true))

function toggleSelect(id: string) {
  if (selectedIds.has(id)) {
    selectedIds.delete(id)
  } else {
    selectedIds.add(id)
  }
}

async function toggleCreateForm() {
  showCreate.value = !showCreate.value
  if (showCreate.value) {
    newTitle.value = ''
    newDesc.value = ''
    createError.value = ''
    await nextTick()
    titleInput.value?.focus()
  }
}

async function handleCreate() {
  if (!newTitle.value.trim()) return
  createError.value = ''
  try {
    await store.addTask(newTitle.value.trim(), newDesc.value.trim() || undefined)
    showCreate.value = false
    newTitle.value = ''
    newDesc.value = ''
  } catch (e: unknown) {
    createError.value = e instanceof Error ? e.message : 'Ошибка создания задачи'
  }
}

async function handleDelete() {
  if (selectedIds.size === 0) return
  try {
    await store.removeTasks([...selectedIds])
    selectedIds.clear()
  } catch (e: unknown) {
    console.error('Ошибка удаления:', e)
  }
}
</script>

<style scoped>
.tasks-page {
  max-width: 720px;
  margin: 0 auto;
  padding: 40px 24px;
}

.page-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 24px;
}

.page-title {
  font-size: 18px;
  font-weight: 600;
  color: var(--text-primary);
}

.header-actions {
  display: flex;
  gap: 8px;
  align-items: center;
}

.btn-primary {
  padding: 7px 16px;
  background: var(--text-primary);
  color: #fff;
  font-size: 13px;
  font-weight: 500;
  border-radius: var(--radius);
  transition: opacity var(--transition);
}

.btn-primary:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.btn-primary:not(:disabled):hover {
  opacity: 0.8;
}

.btn-secondary {
  padding: 7px 16px;
  background: var(--surface);
  color: var(--text-primary);
  font-size: 13px;
  border: 1px solid var(--border);
  border-radius: var(--radius);
  transition: border-color var(--transition), background var(--transition);
}

.btn-secondary:hover {
  border-color: var(--border-hover);
  background: var(--bg);
}

.btn-danger {
  padding: 7px 16px;
  background: var(--surface);
  color: var(--danger);
  font-size: 13px;
  border: 1px solid #fca5a5;
  border-radius: var(--radius);
  transition: background var(--transition), border-color var(--transition);
}

.btn-danger:hover:not(:disabled) {
  background: #fff5f5;
  border-color: var(--danger);
}

.btn-danger:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.btn-ghost {
  padding: 7px 16px;
  font-size: 13px;
  color: var(--text-secondary);
  transition: color var(--transition);
}

.btn-ghost:hover {
  color: var(--text-primary);
}

.btn-ghost:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.create-form {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius);
  padding: 20px;
  margin-bottom: 16px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.create-form form {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.create-input {
  width: 100%;
  padding: 9px 12px;
  border: 1px solid var(--border);
  border-radius: var(--radius);
  font-size: 14px;
  color: var(--text-primary);
  background: var(--bg);
  outline: none;
  transition: border-color var(--transition);
}

.create-input:focus {
  border-color: var(--border-hover);
  background: var(--surface);
}

.create-input::placeholder {
  color: var(--text-muted);
}

.create-actions {
  display: flex;
  gap: 8px;
  align-items: center;
  margin-top: 4px;
}

.error-msg {
  font-size: 12px;
  color: var(--danger);
}

.task-list {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius);
  overflow: hidden;
}

.state-msg {
  padding: 48px 24px;
  text-align: center;
  color: var(--text-muted);
  font-size: 13px;
}

.load-more {
  display: flex;
  justify-content: center;
  padding: 16px;
}
</style>
