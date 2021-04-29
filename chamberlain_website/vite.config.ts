import { defineConfig } from 'vite'
// @ts-ignore
import vue from '@vitejs/plugin-vue'

export default defineConfig({
    plugins: [vue()],
    server: {
        open: true,
        proxy: {
            '/api':{
                target : "http://localhost:8080",
                changeOrigin: true,
                rewrite: (path) => path.replace("\/api", "")
            }
        }
    }
})