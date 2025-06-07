import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'

import dotenv from 'dotenv'
dotenv.config()

const allowedHosts =
    process.env.VITE_ALLOWED_HOSTS?.split(',').map(host => host.trim()) ?? []

export default defineConfig({
    plugins: [react()],
    resolve: {
        alias: {
            '@': path.resolve(__dirname, 'src'),
        },
    },
    server: {
        host: true,
        port: 5173,
        allowedHosts,
        proxy: {
            '/ws': {
                target: 'ws://pricefeed:8080',
                ws: true,
                changeOrigin: true,
            },
        },
    },
})
