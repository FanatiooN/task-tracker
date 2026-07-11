<template>
  <div class="task-item" :class="{ done: task.status === 'TASK_STATUS_DONE', selected }">
    <label class="select-wrap">
      <input
        type="checkbox"
        class="select-input"
        :checked="selected"
        @change="$emit('toggle-select')"
      />
      <span class="select-box"></span>
    </label>

    <div class="task-body">
      <span class="task-title">{{ task.title }}</span>
      <span v-if="task.description" class="task-desc">{{ task.description }}</span>
    </div>

    <button
      class="done-btn"
      :class="{ checked: task.status === 'TASK_STATUS_DONE' }"
      @click="$emit('toggle-done')"
      :title="task.status === 'TASK_STATUS_DONE' ? 'Вернуть в работу' : 'Отметить выполненной'"
    >
      <svg width="18" height="18" viewBox="0 0 18 18" fill="none">
        <circle cx="9" cy="9" r="8" stroke="currentColor" stroke-width="1.4" />
        <path
          v-if="task.status === 'TASK_STATUS_DONE'"
          d="M5.5 9l2.5 2.5 4.5-4.5"
          stroke="currentColor"
          stroke-width="1.4"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
      </svg>
    </button>
  </div>
</template>

<script setup lang="ts">
import type { Task } from '@/types'

defineProps<{ task: Task; selected: boolean }>()
defineEmits<{
  'toggle-select': []
  'toggle-done': []
}>()
</script>

<style scoped>
.task-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  border-bottom: 1px solid var(--border);
  transition: background var(--transition);
}

.task-item:hover {
  background: #faf9f6;
}

.task-item.selected {
  background: #f0efeb;
}

.task-item.done .task-title {
  color: var(--done-text);
  text-decoration: line-through;
  text-decoration-color: var(--border-hover);
}

.task-item.done .task-desc {
  color: var(--text-muted);
}

.select-wrap {
  display: flex;
  align-items: center;
  cursor: pointer;
  flex-shrink: 0;
}

.select-input {
  position: absolute;
  opacity: 0;
  width: 0;
  height: 0;
}

.select-box {
  width: 15px;
  height: 15px;
  border: 1px solid var(--border-hover);
  border-radius: 2px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: border-color var(--transition), background var(--transition);
}

.select-input:checked + .select-box {
  background: var(--text-primary);
  border-color: var(--text-primary);
}

.select-input:checked + .select-box::after {
  content: '';
  display: block;
  width: 4px;
  height: 7px;
  border: 1.5px solid #fff;
  border-top: none;
  border-left: none;
  transform: rotate(45deg) translateY(-1px);
}

.task-body {
  flex: 1;
  min-width: 0;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.task-title {
  font-size: 14px;
  color: var(--text-primary);
  transition: color var(--transition);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.task-desc {
  font-size: 12px;
  color: var(--text-secondary);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  transition: color var(--transition);
}

.done-btn {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border-radius: 50%;
  color: var(--border-hover);
  transition: color var(--transition), background var(--transition);
}

.done-btn:hover {
  color: var(--text-secondary);
  background: var(--border);
}

.done-btn.checked {
  color: var(--accent);
}

.done-btn.checked:hover {
  color: var(--text-secondary);
}
</style>
