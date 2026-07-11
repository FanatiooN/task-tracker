<template>
  <nav v-if="auth.isLoggedIn" class="nav">
    <span class="nav-brand">Task Tracker</span>
    <div class="nav-links">
      <RouterLink to="/tasks" class="nav-link">Tasks</RouterLink>
      <RouterLink to="/profile" class="nav-link">Profile</RouterLink>
      <button class="nav-logout" @click="handleLogout">Logout</button>
    </div>
  </nav>
  <RouterView />
</template>

<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import { useRouter } from 'vue-router'
import * as authApi from '@/api/auth'

const auth = useAuthStore()
const router = useRouter()

async function handleLogout() {
  if (auth.refreshToken) {
    await authApi.logout(auth.refreshToken).catch(() => {})
  }
  auth.clearTokens()
  router.push('/auth')
}
</script>

<style scoped>
.nav {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 32px;
  height: 52px;
  background: var(--surface);
  border-bottom: 1px solid var(--border);
  position: sticky;
  top: 0;
  z-index: 10;
}

.nav-brand {
  font-size: 13px;
  font-weight: 600;
  letter-spacing: 0.04em;
  text-transform: uppercase;
  color: var(--text-primary);
}

.nav-links {
  display: flex;
  align-items: center;
  gap: 24px;
}

.nav-link {
  font-size: 13px;
  color: var(--text-secondary);
  transition: color var(--transition);
}

.nav-link:hover,
.nav-link.router-link-active {
  color: var(--text-primary);
}

.nav-logout {
  font-size: 13px;
  color: var(--text-secondary);
  transition: color var(--transition);
  cursor: pointer;
}

.nav-logout:hover {
  color: var(--danger);
}
</style>
