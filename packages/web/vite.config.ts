import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { config } from 'dotenv'
import tailwindcss from '@tailwindcss/vite'
import { fileURLToPath } from 'url'

config({
  path: '../../.env',
})

const port = Number(process.env.WEB_PORT || 7003)

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue(), tailwindcss()],
  server: {
    port,
    host: '0.0.0.0',
  },
  preview: {
    port,
    host: '0.0.0.0',
    allowedHosts: true,
  },
  resolve: {
    alias: {
      '#': fileURLToPath(new URL('../ui/src', import.meta.url)),
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
})
