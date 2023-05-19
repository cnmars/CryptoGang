# 技术架构
    vite+vue3+ts+pinia+elementui-plus+axios

## 配置过程
    npm i vue-router@4
    npm install pinia pinia-plugin-persist
    npm install element-plus --save
    npm i axios
    //安装图标
    npm install @element-plus/icons-vue
    图标在使用时需要引入:import{Search}from '@element-plus/icons-vue'

    //使用fontAwesome图标库
        1. 安装图标
            npm i --save @fortawesome/fontawesome-svg-core
            # Free icons styles
            npm i --save @fortawesome/free-solid-svg-icons
            npm i --save @fortawesome/free-regular-svg-icons
            npm i --save @fortawesome/free-brands-svg-icons
        2. 安装图标库
            # for Vue 2.x
            npm i --save @fortawesome/vue-fontawesome@latest-2
            # for Vue 3.x
            npm i --save @fortawesome/vue-fontawesome@latest-3
        3. 创建plugins/fontAwesome.ts文件，文件内容如下：
            import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
            import { library } from '@fortawesome/fontawesome-svg-core'
            import { fas } from '@fortawesome/free-solid-svg-icons'
            import { far } from '@fortawesome/free-regular-svg-icons'
            import { fab } from '@fortawesome/free-brands-svg-icons'
            import { App } from 'vue'

            library.add(fas)
            library.add(far)
            library.add(fab)
            export default (app: App<Element>) => {
            app.component('font-awesome-icon', FontAwesomeIcon)
            }
        4. main.ts配置
            import installFontawesome from './plugins/fontAwesome'
            installFontawesome(app)



主色调：
    #126b40
    

组件传值使用pinia，弃用了vue3自身的那一套，所以vue文件代码中普遍采用<script setup>

系统默认文件的修改都打上了 "gd-change-tag"