<template>
  <div class="auth-page">
    <div class="auth-card">
      <h1 class="auth-logo">Task Tracker</h1>

      <div class="tabs">
        <button
          class="tab"
          :class="{ active: activeTab === 'login' }"
          @click="activeTab = 'login'"
        >
          Вход
        </button>
        <button
          class="tab"
          :class="{ active: activeTab === 'register' }"
          @click="activeTab = 'register'"
        >
          Регистрация
        </button>
      </div>

      <form v-if="activeTab === 'login'" @submit.prevent="handleLogin" class="form">
        <div class="field">
          <label class="label">Email</label>
          <input
            v-model="loginForm.email"
            type="email"
            class="input"
            placeholder="you@example.com"
            required
            autocomplete="email"
          />
        </div>
        <div class="field">
          <label class="label">Пароль</label>
          <input
            v-model="loginForm.password"
            type="password"
            class="input"
            placeholder="••••••••"
            required
            autocomplete="current-password"
          />
        </div>
        <p v-if="error" class="error-msg">{{ error }}</p>
        <button type="submit" class="btn-primary" :disabled="loading">
          {{ loading ? 'Входим...' : 'Войти' }}
        </button>
      </form>

      <form v-else @submit.prevent="handleRegister" class="form">
        <div class="field">
          <label class="label">Имя</label>
          <input
            v-model="registerForm.name"
            type="text"
            class="input"
            placeholder="Иван Иванов"
            required
            autocomplete="name"
          />
        </div>
        <div class="field">
          <label class="label">Email</label>
          <input
            v-model="registerForm.email"
            type="email"
            class="input"
            placeholder="you@example.com"
            required
            autocomplete="email"
          />
        </div>
        <div class="field">
          <label class="label">Пароль</label>
          <input
            v-model="registerForm.password"
            type="password"
            class="input"
            placeholder="••••••••"
            required
            autocomplete="new-password"
          />
        </div>
        <p v-if="error" class="error-msg">{{ error }}</p>
        <button type="submit" class="btn-primary" :disabled="loading">
          {{ loading ? 'Создаём...' : 'Зарегистрироваться' }}
        </button>
      </form>

      <div class="divider"><span>или</span></div>

      <div class="oauth-buttons">
        <form method="POST" action="/login/google">
          <button type="submit" class="btn-oauth">Войти через Google</button>
        </form>
        <div ref="telegramContainer"></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import * as authApi from '@/api/auth'

const router = useRouter()
const auth = useAuthStore()

const activeTab = ref<'login' | 'register'>('login')
const loading = ref(false)
const error = ref('')
const telegramContainer = ref<HTMLDivElement>()

const loginForm = reactive({ email: '', password: '' })
const registerForm = reactive({ name: '', email: '', password: '' })

function onSuccess(access: string, refresh: string) {
  auth.setTokens(access, refresh)
  router.push('/tasks')
}

async function handleLogin() {
  error.value = ''
  loading.value = true
  try {
    const tokens = await authApi.login(loginForm.email, loginForm.password)
    onSuccess(tokens.access_token, tokens.refresh_token)
  } catch (e: unknown) {
    error.value = e instanceof Error ? e.message : 'Ошибка входа'
  } finally {
    loading.value = false
  }
}

async function handleRegister() {
  error.value = ''
  loading.value = true
  try {
    const tokens = await authApi.register(
      registerForm.name,
      registerForm.email,
      registerForm.password,
    )
    onSuccess(tokens.access_token, tokens.refresh_token)
  } catch (e: unknown) {
    error.value = e instanceof Error ? e.message : 'Ошибка регистрации'
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  // Init Telegram widget
  if (!telegramContainer.value) return

  ;(window as Window & { onTelegramAuth?: (data: { id_token?: string }) => void }).onTelegramAuth =
    async (data) => {
      if (!data.id_token) return
      try {
        const tokens = await authApi.loginWithTelegram(data.id_token)
        onSuccess(tokens.access_token, tokens.refresh_token)
      } catch (e: unknown) {
        error.value = e instanceof Error ? e.message : 'Ошибка Telegram входа'
      }
    }

  const btn = document.createElement('button')
  btn.className = 'tg-auth-button'
  btn.textContent = 'Войти через Telegram'
  telegramContainer.value.appendChild(btn)

  const script = document.createElement('script')
  script.src = 'https://telegram.org/js/telegram-login.js'
  script.setAttribute('data-client-id', '8995212014')
  script.setAttribute('data-onauth', 'onTelegramAuth(data)')
  script.setAttribute('data-request-access', 'write')
  script.async = true
  telegramContainer.value.appendChild(script)
})
</script>

<style scoped>
.auth-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
  background: var(--bg);
}

.auth-card {
  width: 100%;
  max-width: 380px;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius);
  padding: 40px 36px;
}

.auth-logo {
  font-size: 15px;
  font-weight: 600;
  letter-spacing: 0.05em;
  text-transform: uppercase;
  color: var(--text-primary);
  margin-bottom: 32px;
}

.tabs {
  display: flex;
  border-bottom: 1px solid var(--border);
  margin-bottom: 28px;
}

.tab {
  flex: 1;
  padding: 8px 0;
  font-size: 13px;
  font-weight: 500;
  color: var(--text-muted);
  border-bottom: 2px solid transparent;
  margin-bottom: -1px;
  transition: color var(--transition), border-color var(--transition);
}

.tab.active {
  color: var(--text-primary);
  border-bottom-color: var(--text-primary);
}

.form {
  display: flex;
  flex-direction: column;
  gap: 16px;
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

.input::placeholder {
  color: var(--text-muted);
}

.error-msg {
  font-size: 12px;
  color: var(--danger);
}

.btn-primary {
  width: 100%;
  padding: 10px;
  background: var(--text-primary);
  color: #fff;
  font-size: 13px;
  font-weight: 500;
  border-radius: var(--radius);
  transition: opacity var(--transition);
  margin-top: 4px;
}

.btn-primary:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.btn-primary:not(:disabled):hover {
  opacity: 0.85;
}

.divider {
  display: flex;
  align-items: center;
  gap: 12px;
  margin: 24px 0;
  color: var(--text-muted);
  font-size: 12px;
}

.divider::before,
.divider::after {
  content: '';
  flex: 1;
  height: 1px;
  background: var(--border);
}

.oauth-buttons {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.btn-oauth {
  width: 100%;
  padding: 9px;
  background: var(--surface);
  color: var(--text-primary);
  font-size: 13px;
  border: 1px solid var(--border);
  border-radius: var(--radius);
  transition: border-color var(--transition), background var(--transition);
}

.btn-oauth:hover {
  border-color: var(--border-hover);
  background: var(--bg);
}
</style>
