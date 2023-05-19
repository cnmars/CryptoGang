import {createRouter,createWebHashHistory,createWebHistory, RouteRecordRaw} from 'vue-router'
// import Home from '@/views/home.vue'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import getPageTitle from '../utils/page'
import { useUserStore } from '../pinia/modules/userStore'
import { useRouterStore } from '../pinia/modules/router'
import { defineAsyncComponent } from 'vue';
import { LOCAL_STORAGE } from "../global/constants";
//扩展原生对象，目前用不上
export type ApprouterREcordRaw = RouteRecordRaw & {
  hiden?: boolean
}
// const _import = (path: string) => defineAsyncComponent(() => import(`../views/${path}.vue`));
export const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    redirect:'/layout/nft/hashMint',
    // component: () => import('@/views/home.vue')
    // component: Home
  },
] as ApprouterREcordRaw[]

const router = createRouter({
  history: createWebHistory(),
 /*  history: createWebHashHistory(),//浏览器路径带# */
  routes
})
export default router