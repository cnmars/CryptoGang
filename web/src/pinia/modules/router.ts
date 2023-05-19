import { asyncRouterHandle } from '@/utils/asyncRouter'
import { emitter } from '@/common/bus'
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useUserStore } from "@/pinia/modules/userStore";
import { asyncMenu } from '@/api/menu'
const routerListArr = []
const notLayoutRouterArr = []
const keepAliveRoutersArr = []
const nameMap = {}
const formatRouter = (routes, routeMap) => {
  routes && routes.forEach(item => {
    //没有子菜单或者子菜单是隐藏状态
    if ((!item.children || item.children.every(ch => ch.hidden)) && item.name !== '404' && !item.hidden) {
      if(item?.meta?.title){
        routerListArr.push({ label: item.meta.title, value: item.name })

      }
    }
    item.meta.btns = item.btns
    item.meta.hidden = item.hidden
    if (item.meta.defaultMenu === true) {
      notLayoutRouterArr.push({
        ...item,
        path: `/${item.path}`,
      })
    } else {
      routeMap[item.name] = item
      if (item.children && item.children.length > 0) {
        formatRouter(item.children, routeMap)
      }
    }
  })
}

const KeepAliveFilter = (routes) => {
  routes && routes.forEach(item => {
    // 子菜单中有 keep-alive 的，父菜单也必须 keep-alive，否则无效。这里将子菜单中有 keep-alive 的父菜单也加入。
    if ((item.children && item.children.some(ch => ch.meta.keepAlive) || item.meta.keepAlive)) {
      
      item.component && item.component().then(val => {
        keepAliveRoutersArr.push(val.default.name)
        nameMap[item.name] = val.default.name
      })
    }
    if (item.children && item.children.length > 0) {
      KeepAliveFilter(item.children)
    }
  })
}

export const useRouterStore = defineStore('router', () => {
  const keepAliveRouters = ref([])
  const setKeepAliveRouters = (history) => {
    const keepArrTemp = []
    history.forEach(item => {
      if (nameMap[item.name]) {
        keepArrTemp.push(nameMap[item.name])
      }
    })
    
    keepAliveRouters.value = Array.from(new Set(keepArrTemp))
  }
  emitter.on('setKeepAlive', setKeepAliveRouters)

  const asyncRouters = ref([])
  const routeMap = ({})
  // 从后台获取动态路由
  const SetAsyncRouter = async() => {
    const baseRouter = [{
      path: '/layout',
      name: 'layout',
      component: 'views/layout/mainPage.vue',
      meta: {
        title: '底层layout'
      },
      children: [
        {
          path: "404",
          name: "404",
          hidden: true,
          meta: {
            title: "迷路了*。*",
            closeTab: true,
            hidden: true
          },
          component: "views/error/index.vue"
        },
        {
          path: "login",
          name: "login",
          hidden: true,
          meta: {
            title: "签名登录",
            closeTab: true,
            hidden: true
          },
          component: "views/account/login/login.vue"
        },
        {
          path: "reload",
          name: "Reload",
          hidden: true,
          meta: {
            title: "",
            closeTab: true,
            hidden: true
          },
          component: "views/error/reload.vue"
        }
      ]
    }]
    const userStore = useUserStore()
    var asyncRouterRes=null
    var asyncRouter =[]
    if(userStore.token){
       asyncRouterRes = await asyncMenu()
       asyncRouter = asyncRouterRes.data.menus
    }
    if (!asyncRouter){
      asyncRouter=[]
    }
    //后台来的路由添加到现有layout下面
    // let preChild = router.options.routes[0].children
    // baseRouter[0].children = preChild.concat(asyncRouter) as any
    baseRouter[0].children = baseRouter[0].children.concat(asyncRouter)
    formatRouter(baseRouter, routeMap)
    baseRouter.push({
      path: '/',
      redirect: '/layout/login'
    }as any) 
    baseRouter.push({
      path: '/:catchAll(.*)',
      redirect: '/layout/404'

    }as any)
    // baseRouter[0].children = baseRouter[0].children.concat(asyncRouter)
    if (notLayoutRouterArr.length !== 0) {
      baseRouter[0].children.push(...notLayoutRouterArr)
    }
    asyncRouterHandle(baseRouter)
    // KeepAliveFilter(asyncRouter)
    KeepAliveFilter(asyncRouter.concat(baseRouter[0].children))
    asyncRouters.value = baseRouter as any
    //routerList用于搜索栏跳转功能
    return true
  }

  return {
    asyncRouters,
    keepAliveRouters,
    SetAsyncRouter,
    routeMap
  }
})

