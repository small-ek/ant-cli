import {defineConfig} from 'vite'
import {resolve} from 'path'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite';
import {ArcoResolver} from 'unplugin-vue-components/resolvers';
import {vitePluginForArco} from '@arco-plugins/vite-vue'

import viteCompression from 'vite-plugin-compression'

// https://vitejs.dev/config/
export default defineConfig({
    resolve: {
        alias: {
            '@': resolve(__dirname, 'src'),
        }
    },
    plugins: [vue(), AutoImport({
        resolvers: [ArcoResolver()],
    }),
        Components({
            resolvers: [ArcoResolver({
                sideEffect: true
            })]
        }),
        vitePluginForArco({
            style: 'css'
        }),
        viteCompression({
            verbose: true, // 默认即可
            disable: false, // 开启压缩(不禁用)，默认即可
            deleteOriginFile: true, // 删除源文件
            threshold: 5120, // 压缩前最小文件大小
            algorithm: 'gzip', // 压缩算法
            ext: '.gz' // 文件类型
        })],
    server: {
        // 服务器主机名，如果允许外部访问，可设置为 "0.0.0.0"
        host: "0.0.0.0",
    },
    base: "web",
    build: {
        target: 'es2015',
        cssTarget: 'chrome80',
        cssCodeSplit: true, // 启用 CSS 代码拆分
        minify: "terser", // 必须开启：使用terserOptions才有效果
        rollupOptions: {
            output: {// 分包
                chunkFileNames: 'static/js/[name]-[hash].js',
                entryFileNames: 'static/js/[name]-[hash].js',
                assetFileNames: 'static/[ext]/[name]-[hash].[ext]',
                manualChunks(id) {
                    if (id.includes("node_modules")) {
                        return id.toString().split("node_modules/")[1].split("/")[0].toString()
                    }
                },
            }
        },
        terserOptions: {
            compress: {
                keep_infinity: true,  // 防止 Infinity 被压缩成 1/0，这可能会导致 Chrome 上的性能问题
                drop_console: true,   // 生产环境去除 console
                drop_debugger: true   // 生产环境去除 debugger
            },
        },
        brotliSize: false, // 进行压缩计算
        chunkSizeWarningLimit: 1500, // chunk 大小警告的限制（以 kbs 为单位）
    }
})
