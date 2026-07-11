import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { fileURLToPath, URL } from 'node:url'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  server: {
    proxy: {
      '/login': 'http://localhost:8080',
      '/logout': 'http://localhost:8080',
      '/register': 'http://localhost:8080',
      '/refresh': 'http://localhost:8080',
      '/users': 'http://localhost:8080',
      '/tasks': 'http://localhost:8080',
    },
  },
})
