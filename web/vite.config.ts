import { defineConfig,loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
// 如果编辑器提示 path 模块找不到，则可以安装一下 @types/node -> npm i @types/node -D
import { resolve as rov} from 'path'
import path from 'path'
import viteCompression from 'vite-plugin-compression'
import {createSvgIconsPlugin } from 'vite-plugin-svg-icons';
export default defineConfig(({command,mode})=>{
  let base: string
  let wssUrl:string
  if (command === 'build') {
    base = 'https://cryptogang.vip/'
    wssUrl = 'wss://cryptogang.vip/system/ws'
    console.log("build");
  } else {
    base = './'
    wssUrl = 'ws://localhost:4728/system/ws'
  }
  const envConfig = loadEnv(mode, './',["VITE","VENUS","GD"]);

  let config ={
    
  }
  return{
    base,
    envPrefix: ["VITE", "VENUS","GD"],
    plugins: [vue(),createSvgIconsPlugin({
      // 指定需要缓存的图标文件夹
    iconDirs: [path.resolve(process.cwd(), 'src/assets/svg')],
      // 指定symbolId格式
      symbolId: 'icon-[dir]-[name]',
    }) as any,viteCompression({
      verbose: true,
      disable: false,
      threshold: 10240,
      algorithm: 'gzip',
      ext: '.gz',
    })],
    resolve: {
      alias: {
        '@': rov(__dirname, 'src') // 设置 `@` 指向 `src` 目录
      }
    },
    /*  base: './', */  
    define: {
      'process.env': {...envConfig,WSS_URL:wssUrl}
    },
     server: {
      port: 4444, // 设置服务启动端口号
      open: false, // 设置服务启动时是否自动打开浏览器
      cors: true ,// 允许跨域
  
      // 设置代理
      proxy: {
        '/api': {
          target: 'http://127.0.0.1:4728/',
          changeOrigin: true,
          secure: false,
          rewrite: (path) => path.replace('/api/', '/'),
        },
      },
    },
    build: {
      sourcemap: false,
      minify: 'terser',
      outDir: 'dist',
      chunkSizeWarningLimit: 1600,
      terserOptions: {
        compress: {
          drop_console: true,
          drop_debugger: true
        }
      },
      rollupOptions: {
        output: {
          manualChunks(id) {
            if (id.includes('node_modules')) {
              return id
                .toString()
                .split('node_modules/')[1]
                .split('/')[0]
                .toString();
            }
          },
          chunkFileNames: (chunkInfo) => {
            const facadeModuleId = chunkInfo.facadeModuleId
              ? chunkInfo.facadeModuleId.split('/')
              : [];
            const fileName =
              facadeModuleId[facadeModuleId.length - 2] || '[name]';
            return `js/${fileName}/[name].[hash].js`;
          }
        }
      }
    },
  }
})