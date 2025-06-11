import { defineConfig, loadEnv } from "vite";
import process from "node:process";
import vue from "@vitejs/plugin-vue";
import unocss from "unocss/vite";
import path from "path";
import vueJsx from '@vitejs/plugin-vue-jsx'
export default defineConfig((configEnv) => {
  const viteEnv = loadEnv(configEnv.mode, process.cwd());
  return {
    plugins: [vue(), unocss(),vueJsx()],
    build: {
      sourcemap: false,
      commonjsOptions: {
        ignoreTryCatch: false,
      },
    },
    resolve: {
      extensions: [
        ".mjs",
        ".js",
        ".ts",
        ".jsx",
        ".tsx",
        ".json",
        ".less",
        ".css",
      ],
      alias: {
        "@": path.resolve(__dirname, "src"),
      },
    },
    server: {
      host: "0.0.0.0",
      port: 80,
      client: {
        webSocketURL: 'http://nj9ahm.natappfree.cc/message/ws',
      },
      headers: {
        'Access-Control-Allow-Origin': '*',
      },
      proxy: {
       
    
        "/api": {
          target: viteEnv.VITE_SERVER,
          changeOrigin: true,
          rewrite: (path) => {
            console.log('path',222)
            return path.replace(/^\/api/, "")
          },
        },
      },
      hmr: {
        overlay: false,
      },
    },
  };
});
