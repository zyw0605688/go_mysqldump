import { defineConfig } from "vite";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";
import vue from "@vitejs/plugin-vue";
import path from "path";

export default defineConfig({
  publicDir: "./public",
  plugins: [
    vue(),
    AutoImport({
      resolvers: [ElementPlusResolver()]
    }),
    Components({
      resolvers: [ElementPlusResolver()]
    })
  ],
  base: "./",
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src")
    }
  },
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: "",
        api: "modern-compiler"
      }
    }
  },
  server: {
    open: true,
    host: "127.0.0.1",
    cors: true,
    proxy: {
      "/zmos_ze05": {
        target: "http://127.0.0.1:3028/",
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/zmos_ze05/, "")
      }
    },
    hmr: true, // 开启热更新
    strictPort: false //
  },
  build: {
    terserOptions: {
      compress: {
        drop_console: true,
        drop_debugger: true
      }
    }
  }
});
