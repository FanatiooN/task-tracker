<template>
  <div class="profile-page">
    <div class="page-header">
      <h2 class="page-title">Profile</h2>
    </div>

    <div v-if="loadingUser" class="state-msg">Загрузка...</div>

    <div v-else-if="user" class="profile-card">
      <form @submit.prevent="handleUpdate">
        <div class="field-group">
          <div class="field">
            <label class="label">Имя</label>
            <input v-model="editName" type="text" class="input" placeholder="Ваше имя" required />
          </div>

          <div v-if="user.email" class="field">
            <label class="label">Email</label>
            <input :value="user.email" type="email" class="input input-readonly" readonly />
          </div>
        </div>

        <p v-if="updateError" class="error-msg">{{ updateError }}</p>
        <p v-if="updateSuccess" class="success-msg">Изменения сохранены</p>

        <div class="form-actions">
          <button
            type="submit"
            class="btn-primary"
            :disabled="saving || editName === user.name"
          >
            {{ saving ? 'Сохраняем...' : 'Сохранить' }}
          </button>
        </div>
      </form>

      <div class="danger-zone">
        <div v-if="!confirmDelete">
          <button class="btn-danger" @click="confirmDelete = true">Удалить аккаунт</button>
        </div>
        <div v-else class="confirm-row">
          <span class="confirm-text">Вы уверены? Это необратимо.</span>
          <button class="btn-danger-solid" @click="handleDelete" :disabled="deleting">
            {{ deleting ? 'Удаляем...' : 'Да, удалить' }}
          </button>
          <button class="btn-ghost" @click="confirmDelete = false">Отмена</button>
        </div>
      </div>
    </div>

    <div v-else class="state-msg">Не удалось загрузить профиль</div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import * as usersApi from '@/api/users'
import { AuthError } from '@/api/client'
import type { User } from '@/types'

const auth = useAuthStore()
const router = useRouter()

const user = ref<User | null>(null)
const loadingUser = ref(true)
const editName = ref('')
const saving = ref(false)
const updateError = ref('')
const updateSuccess = ref(false)
const confirmDelete = ref(false)
const deleting = ref(false)

onMounted(async () => {
  if (!auth.userId) return
  try {
    user.value = await usersApi.getUser(auth.userId)
    editName.value = user.value.name
  } catch (e: unknown) {
    if (e instanceof AuthError) { router.push('/auth'); return }
    console.error('Ошибка загрузки профиля:', e)
  } finally {
    loadingUser.value = false
  }
})

async function handleUpdate() {
  if (!auth.userId || !user.value) return
  updateError.value = ''
  updateSuccess.value = false
  saving.value = true
  try {
    user.value = await usersApi.updateUser(auth.userId, editName.value)
    updateSuccess.value = true
    setTimeout(() => (updateSuccess.value = false), 3000)
  } catch (e: unknown) {
    updateError.value = e instanceof Error ? e.message : 'Ошибка сохранения'
  } finally {
    saving.value = false
  }
}

async function handleDelete() {
  if (!auth.userId) return
  deleting.value = true
  try {
    await usersApi.deleteUser(auth.userId)
    auth.clearTokens()
    router.push('/auth')
  } catch (e: unknown) {
    console.error('Ошибка удаления аккаунта:', e)
  } finally {
    deleting.value = false
  }
}
</script>

<style scoped>
.profile-page {
  max-width: 480px;
  margin: 0 auto;
  padding: 40px 24px;
}

.page-header {
  margin-bottom: 24px;
}

.page-title {
  font-size: 18px;
  font-weight: 600;
  color: var(--text-primary);
}

.state-msg {
  padding: 48px 0;
  text-align: center;
  color: var(--text-muted);
  font-size: 13px;
}

.profile-card {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius);
  overflow: hidden;
}

.field-group {
  padding: 28px 28px 0;
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.label {
  font-size: 12px;
  font-weight: 500;
  color: var(--text-secondary);
  letter-spacing: 0.02em;
}

.input {
  width: 100%;
  padding: 9px 12px;
  border: 1px solid var(--border);
  border-radius: var(--radius);
  font-size: 14px;
  color: var(--text-primary);
  background: var(--surface);
  outline: none;
  transition: border-color var(--transition);
}

.input:focus {
  border-color: var(--border-hover);
}

.input-readonly {
  background: var(--bg);
  color: var(--text-secondary);
  cursor: default;
}

.error-msg {
  font-size: 12px;
  color: var(--danger);
  padding: 0 28px;
  margin-top: 8px;
}

.success-msg {
  font-size: 12px;
  color: #16a34a;
  padding: 0 28px;
  margin-top: 8px;
}

.form-actions {
  padding: 20px 28px 28px;
}

.btn-primary {
  padding: 9px 20px;
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

.danger-zone {
  border-top: 1px solid var(--border);
  padding: 20px 28px;
}

.btn-danger {
  font-size: 13px;
  color: var(--danger);
  transition: opacity var(--transition);
}

.btn-danger:hover {
  opacity: 0.7;
}

.btn-danger-solid {
  padding: 7px 16px;
  background: var(--danger);
  color: #fff;
  font-size: 13px;
  border-radius: var(--radius);
  transition: background var(--transition);
}

.btn-danger-solid:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.btn-danger-solid:not(:disabled):hover {
  background: var(--danger-hover);
}

.btn-ghost {
  padding: 7px 12px;
  font-size: 13px;
  color: var(--text-secondary);
  transition: color var(--transition);
}

.btn-ghost:hover {
  color: var(--text-primary);
}

.confirm-row {
  display: flex;
  align-items: center;
  gap: 12px;
  flex-wrap: wrap;
}

.confirm-text {
  font-size: 13px;
  color: var(--text-secondary);
}
</style>
