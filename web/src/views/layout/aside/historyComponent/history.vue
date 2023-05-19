<template>
  <div class="router-history">
    <el-tabs
      v-model="activeValue"
      :closable="!(historys.length === 1 && $route.name === defaultRouter)"
      type="card"
      @contextmenu.prevent="openContextMenu($event)"
      @tab-change="changeTab"
      @tab-remove="removeTab"
    >
      <el-tab-pane v-for="item in historys" :key="name(item)" :label="$t('route.'+item.name)" :name="name(item)" :tab="item" class="gd-tab" >
        <template #label>
          <span :tab="item" :style="{ color: activeValue === name(item) ? 'var(--el-color-primary)' : 'var(--el-text-color-primary)', }">
            <i class="dot" :style="{ backgroundColor: activeValue === name(item) ? 'var(--el-color-primary)' : '#ddd', }" />
            {{ $t('route.'+item.name) }}
          </span>
        </template>
      </el-tab-pane>
    </el-tabs>

    <!--自定义右键菜单html代码-->
    <ul
      v-show="contextMenuVisible"
      :style="{ left: left + 'px', top: top + 'px' }"
      class="contextmenu"
    >
      <li @click="closeAll">关闭所有</li>
      <li @click="closeLeft">关闭左侧</li>
      <li @click="closeRight">关闭右侧</li>
      <li @click="closeOther">关闭其他</li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'HistoryComponent',
}
</script>

<script setup>
import { emitter } from '@/common/bus'
import { computed, onUnmounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/userStore'
import { fmtTitle } from '@/utils/fmtRouterTitle'

const route = useRoute()
const router = useRouter()

const getFmtString = (item) => {
  return item.name + JSON.stringify(item.query) + JSON.stringify(item.params)
}

const historys = ref([])
const activeValue = ref('')
const contextMenuVisible = ref(false)

const userStore = useUserStore()
userStore.activeColor='#126b40'
const name = (item) => {
  return (
    item.name + JSON.stringify(item.query) + JSON.stringify(item.params)
  )
}

const left = ref(0)
const top = ref(0)
const isCollapse = ref(false)
const isMobile = ref(false)
const rightActive = ref('')
console.log() 
const defaultRouter = import.meta.env.GD_DEFAULT_ROUTER
// const defaultRouter = computed(() => userStore.userInfo.authority.defaultRouter)
const openContextMenu = (e) => {
  if (historys.value.length === 1 && route.name === defaultRouter) {
    return false
  }
  let id = ''
  if (e.srcElement.nodeName === 'SPAN') {
    id = e.srcElement.offsetParent.id
  } else {
    id = e.srcElement.id
  }
  if (id) {
    contextMenuVisible.value = true
    let width
    if (isCollapse.value) {
      width = 54
    } else {
      width = 220
    }
    if (isMobile.value) {
      width = 0
    }
    left.value = e.clientX - width
    top.value = e.clientY + 10
    rightActive.value = id.substring(4)
  }
}
//关闭所有
const closeAll = () => {
  historys.value = [
    {
      name: defaultRouter,
      meta: {
        title: '首页',
      },
      query: {},
      params: {},
    },
  ]
  router.push({ name: defaultRouter })
  contextMenuVisible.value = false
  sessionStorage.setItem('historys', JSON.stringify(historys.value))
}
//关闭左边
const closeLeft = () => {
  let right
  const rightIndex = historys.value.findIndex((item) => {
    if (getFmtString(item) === rightActive.value) {
      right = item
    }
    return getFmtString(item) === rightActive.value
  })
  const activeIndex = historys.value.findIndex(
    (item) => getFmtString(item) === activeValue.value
  )
  historys.value.splice(0, rightIndex)
  if (rightIndex > activeIndex) {
    router.push(right)
  }
  sessionStorage.setItem('historys', JSON.stringify(historys.value))
}
//关闭右边
const closeRight = () => {
  let right
  const leftIndex = historys.value.findIndex((item) => {
    if (getFmtString(item) === rightActive.value) {
      right = item
    }
    return getFmtString(item) === rightActive.value
  })
  const activeIndex = historys.value.findIndex(
    (item) => getFmtString(item) === activeValue.value
  )
  historys.value.splice(leftIndex + 1, historys.value.length)
  if (leftIndex < activeIndex) {
    router.push(right)
  }
  sessionStorage.setItem('historys', JSON.stringify(historys.value))
}
//关闭其他
const closeOther = () => {
  let right
  historys.value = historys.value.filter((item) => {
    if (getFmtString(item) === rightActive.value) {
      right = item
    }
    return getFmtString(item) === rightActive.value
  })
  router.push(right)
  sessionStorage.setItem('historys', JSON.stringify(historys.value))
}
//判断两个路径是不是相同
const isSame = (route1, route2) => {
  if (route1.name !== route2.name) {
    return false
  }
  if (Object.keys(route1.query).length !== Object.keys(route2.query).length || Object.keys(route1.params).length !== Object.keys(route2.params).length) {
    return false
  }
  for (const key in route1.query) {
    if (route1.query[key] !== route2.query[key]) {
      return false
    }
  }
  for (const key in route1.params) {
    if (route1.params[key] !== route2.params[key]) {
      return false
    }
  }
  return true
}
//设置标签
const setTab = (route) => {
  if (!historys.value.some((item) => isSame(item, route))) {
    const obj = {}
    obj.name = route.name
    obj.meta = { ...route.meta }
    delete obj.meta.matched
    obj.query = route.query
    obj.params = route.params
    historys.value.push(obj)
  }
  window.sessionStorage.setItem('activeValue', getFmtString(route))
}

const historyMap = ref({})

//切换标签
const changeTab = (name) => {
  const tab = historyMap.value[name]
  //解决在菜单已经打开，在顶部有标签的情况下，这时候去点击菜单会出现两次路由，菜单一次，标签栏一次
  if(route.name!==tab.name){
      router.push({
      name: tab.name,
      query: tab.query,
      params: tab.params,
    })
  }
  // router.push({
  //   name: tab.name,
  //   query: tab.query,
  //   params: tab.params,
  // })
}
//关闭标签
const removeTab = (tab) => {
  const index = historys.value.findIndex(
    (item) => getFmtString(item) === tab
  )
  if (getFmtString(route) === tab) {
    if (historys.value.length === 1) {
      router.push({ name: defaultRouter })
    } else {
      if (index < historys.value.length - 1) {
        router.push({
          name: historys.value[index + 1].name,
          query: historys.value[index + 1].query,
          params: historys.value[index + 1].params,
        })
      } else {
        router.push({
          name: historys.value[index - 1].name,
          query: historys.value[index - 1].query,
          params: historys.value[index - 1].params,
        })
      }
    }
  }
  historys.value.splice(index, 1)
}

watch(() => contextMenuVisible.value, () => {
  if (contextMenuVisible.value) {
    document.body.addEventListener('click', () => {
      contextMenuVisible.value = false
    })
  } else {
    document.body.removeEventListener('click', () => {
      contextMenuVisible.value = false
    })
  }
})

watch(() => route, (to, now) => {
  if (to.name === 'Login' || to.name === 'Reload') {
    return
  }
  historys.value = historys.value.filter((item) => !item.meta.closeTab)
  setTab(to)
  sessionStorage.setItem('historys', JSON.stringify(historys.value))
  activeValue.value = window.sessionStorage.getItem('activeValue')
}, { deep: true })

watch(() => historys.value, () => {
  sessionStorage.setItem('historys', JSON.stringify(historys.value))
  historyMap.value = {}
  historys.value.forEach((item) => {
    historyMap.value[getFmtString(item)] = item
  })
  emitter.emit('setKeepAlive', historys.value)
}, {
  deep: true
})

//初始化页面
const initPage = () => {
  // 全局监听 关闭当前页面函数
  emitter.on('closeThisPage', () => {
    // console.log("name(route)==>",name(route));
    removeTab(name(route))
  })
  // 全局监听 关闭所有页面函数
  emitter.on('closeAllPage', () => {
    closeAll()
  })
  emitter.on('mobile', (data) => {
    isMobile.value = data
  })
  emitter.on('collapse', (data) => {
    isCollapse.value = data
  })
  const initHistorys = [
    {
      name: defaultRouter,
      meta: {
        title: '首页',
      },
      query: {},
      params: {},
    },
  ]
  historys.value = JSON.parse(sessionStorage.getItem('historys')) || initHistorys
  if (!window.sessionStorage.getItem('activeValue')) {
    activeValue.value = getFmtString(route)
  } else {
    activeValue.value = window.sessionStorage.getItem('activeValue')
  }
  setTab(route)
  //在切换用户的时候需要关闭所有标签
  if (window.sessionStorage.getItem('needCloseAll') === 'true') {
    closeAll()
    window.sessionStorage.removeItem('needCloseAll')
  }
}
initPage()

onUnmounted(() => {
  emitter.off('collapse')
  emitter.off('mobile')
})
</script>

<style lang="scss">
.contextmenu {
  width: 100px;
  margin: 0;
  border: 1px solid #ccc;
  background: var(--main-bg-color);
  z-index: 3000;
  position: absolute;
  list-style-type: none;
  padding: 5px 0;
  border-radius: 4px;
  font-size: 14px;
  color: var( --el-text-color-primary);
  // color: #333;
  box-shadow: 2px 2px 3px 0 rgba(0, 0, 0, 0.2);
}
.el-tabs__item .el-icon-close {
  color: initial !important;
}
.el-tabs__item .dot {
  content: "";
  width: 9px;
  height: 9px;
  margin-right: 8px;
  display: inline-block;
  border-radius: 50%;
  transition: background-color 0.2s;
}

.contextmenu li {
  margin: 0;
  padding: 7px 16px;
}
.contextmenu li:hover {
  background: #f2f2f2;
  cursor: pointer;
}
</style>
