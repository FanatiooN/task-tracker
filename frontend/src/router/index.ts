import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import AuthView from '@/views/AuthView.vue'
import TasksView from '@/views/TasksView.vue'
import ProfileView from '@/views/ProfileView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/auth',
      component: AuthView,
      meta: { public: true },
    },
    {
      path: '/tasks',
      component: TasksView,
    },
    {
      path: '/profile',
      component: ProfileView,
    },
    {
      path: '/',
      redirect: '/tasks',
    },
  ],
})

router.beforeEach((to) => {
  const auth = useAuthStore()

  if (to.query.access_token && to.query.refresh_token) {
    auth.setTokens(to.query.access_token as string, to.query.refresh_token as string)
    return { path: '/tasks', query: {} }
  }

  if (!auth.isLoggedIn && !to.meta.public) return '/auth'
  if (auth.isLoggedIn && to.path === '/auth') return '/tasks'
})

export default router
