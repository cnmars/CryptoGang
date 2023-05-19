import { useUserStore } from '@/pinia/modules/userStore'
import { useRouterStore } from '@/pinia/modules/router'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import getPageTitle from '@/utils/page'
import router from '@/router/router'
import { LOCAL_STORAGE } from "@/global/constants";
import { unfinished } from '@/global/config'
import { ElMessage } from 'element-plus'
let asyncRouterFlag = 0

const whiteList = ['login', 'Init']


async function handleKeepAlive(to) {
  if (to.matched.some(item => item.meta.keepAlive)) {
    if (to.matched && to.matched.length > 2) {
      for (let i = 1; i < to.matched.length; i++) {
        const element = to.matched[i - 1]
        if (element.name === 'layout') {
          to.matched.splice(i, 1)
          await handleKeepAlive(to)
        }
        // 如果没有按需加载完成则等待加载
        if (typeof element.components.default === 'function') {
          await element.components.default()
          await handleKeepAlive(to)
        }
      }
    }
  }
}

const getRouter = async(userStore,address) => {
    const routerStore = useRouterStore()
    await routerStore.SetAsyncRouter()
    await userStore.GetUserInfo(address)
    const asyncRouters = routerStore.asyncRouters
    asyncRouters.forEach(asyncRouter => {
      router.addRoute(asyncRouter)
    })
  }
const backFillMsg =async(token:string)=>{
  const userStore = useUserStore()
  const address = localStorage.getItem(LOCAL_STORAGE.USER_ADDRESS)
  if(address==null || address =='' || !address){
    userStore.LoginOut()
  }else{
    await userStore.GetUserInfo(address)
      userStore.userInfo.isConnect = true
      if(token){
        userStore.userInfo.isSign=true
      }else{
        userStore.userInfo.isSign=false
      }
  }
}

router.beforeEach(async(to:any, from:any) => {
    NProgress.start()
    if(unfinished.indexOf(to.name) > -1){
      ElMessage.error("该功能未开放")
      return { path: from.path }
    }
    const userStore = useUserStore()
    to.meta.matched = [...to.matched]
    handleKeepAlive(to)
    const token = userStore.token
    // await backFillMsg(token)
    // 在白名单中的判断情况
    document.title = getPageTitle(to.meta.title, to)
    if (whiteList.indexOf(to.name) > -1) {
      return true
    } else {
       await backFillMsg(token)
      // 不在白名单中并且已经登录的时候
      if (token) {
        // 添加flag防止多次获取动态路由和栈溢出
        if ((!asyncRouterFlag && whiteList.indexOf(from.name) < 0)) {
          asyncRouterFlag++
          await getRouter(userStore,"")
          if (userStore.token) {
            return { ...to, replace: true }
          } else {
            return {
              name: 'login',
              query: { redirect: to.href }
            }
          }
        } else {
          if (to.matched.length) {
            return true
          } else {
            return { path: '/layout/404' }
          }
        }
      }
      // 不在白名单中并且未登录的时候
      if (!token) {
        await getRouter(userStore,"")
        return {
          name: 'login',
          query: {
            redirect: document.location.hash
          }
        }
      }
    }
  })
router.afterEach(() => {
  // 路由加载完成后关闭进度条
  NProgress.done()
})

router.onError(() => {
  // 路由发生错误后销毁进度条
  NProgress.remove()
})
