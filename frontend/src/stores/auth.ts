import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

function parseUserId(token: string): string | null {
  try {
    const payload = JSON.parse(atob(token.split('.')[1]))
    return payload.UserID ?? null
  } catch {
    return null
  }
}

export const useAuthStore = defineStore('auth', () => {
  const accessToken = ref<string | null>(localStorage.getItem('access_token'))
  const refreshToken = ref<string | null>(localStorage.getItem('refresh_token'))
  const userId = ref<string | null>(localStorage.getItem('user_id'))

  const isLoggedIn = computed(() => !!accessToken.value)

  function setTokens(access: string, refresh: string) {
    accessToken.value = access
    refreshToken.value = refresh
    userId.value = parseUserId(access)
    localStorage.setItem('access_token', access)
    localStorage.setItem('refresh_token', refresh)
    if (userId.value) localStorage.setItem('user_id', userId.value)
  }

  function clearTokens() {
    accessToken.value = null
    refreshToken.value = null
    userId.value = null
    localStorage.clear()
  }

  return { accessToken, refreshToken, userId, isLoggedIn, setTokens, clearTokens }
})
