<template>
  <div :style="{ background: userStore.sideMode }">
    <el-scrollbar style="height: calc(100vh - 110px);overflow: hidden;">
      <transition
        :duration="{ enter: 800, leave: 100 }"
        mode="out-in"
        name="el-fade-in-linear"
      >
        <el-menu
          :collapse="isCollapse"
          :collapse-transition="false"
          :default-active="active"
          :background-color="theme.background"
          :active-text-color="theme.active"
          class="el-menu-vertical"
          unique-opened
          @select="selectMenuItem"
        >
          <template v-for="item in routerStore.asyncRouters[0].children">
            <aside-component
              v-if="!item.hidden"
              :key="item.name"
              :is-collapse="isCollapse"
              :router-info="item"
              :theme="theme"
            />
          </template>
        </el-menu>
      </transition>
    </el-scrollbar>
    <GasAndPriceVue/>

  </div>
</template>

<script>
export default {
  name: 'Aside',
}
</script>

<script setup>
import AsideComponent from '@/views/layout/aside/asideComponent/index.vue'
import GasAndPriceVue from './gasComponent/gasAndPrice.vue'
import { emitter } from '@/common/bus'
import { ref, watch, onUnmounted, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/userStore'
import { appConfig } from '../../../global/config'
import { useRouterStore } from '@/pinia/modules/router'

const route = useRoute()
const router = useRouter()
// const openeds =["nft","erc20","chainLog","liquidManage"]
const userStore = useUserStore()
const routerStore = useRouterStore()
// console.log('sideMode=============================')
// console.log(userStore.sideMode)
// console.log('routerStore=============================')
// console.log(routerStore.asyncRouters)
const theme = ref({})
// const handleOpen=(index)=> {
//   console.log(index);
// }
const getTheme = () => {
  switch (appConfig.sideMode) {
    case 'dark':
      theme.value = {
        background: '#fff',
        activeBackground: '#4D70FF',
        activeText: '#fff',
        normalText: '#333',
        hoverBackground: 'rgba(64, 158, 255, 0.08)',
        hoverText: '#333',
      }
      break
    // case 'light':
    //   theme.value = {
    //     background: '#191a23',
    //     activeBackground: '#4D70FF',
    //     activeText: '#fff',
    //     normalText: '#fff',
    //     hoverBackground: 'rgba(64, 158, 255, 0.08)',
    //     hoverText: '#fff',
    //   }
    case 'light':
      theme.value = {
        background: '#fff',
        activeBackground: '#fff',
        activeText: '#c28e0c',
        normalText: '#126b40',
        hoverBackground: '#ebe8b9',
        hoverText: '#126b40',
      }
      break
  }
}

getTheme()
const active = ref('')
watch(() => route, () => {
  active.value = route.meta.activeName || route.name
}, { deep: true })

// watch(() => userStore.sideMode, () => {
//   getTheme()
// })

const isCollapse = ref(false)
const initPage = () => {
  active.value = route.meta.activeName || route.name
  const screenWidth = document.body.clientWidth
  if (screenWidth < 1000) {
    isCollapse.value = !isCollapse.value
  }

  emitter.on('collapse', (item) => {
    isCollapse.value = item
  })
}

initPage()

onUnmounted(() => {
  emitter.off('collapse')
})

//菜单跳转路由
const selectMenuItem = (index, _, ele, aaa) => {
  const query = {}
  const params = {}
 routerStore.routeMap[index]?.parameters &&
    routerStore.routeMap[index]?.parameters.forEach((item) => {
      if (item.type === 'query') {
        query[item.key] = item.value
      } else {
        params[item.key] = item.value
      }
    })
 if (index === route.name) return
 if (index.indexOf('http://') > -1 || index.indexOf('https://') > -1) {
   window.open(index)
 } else {
   router.push({ name: index, query, params })
 }
}




</script>

<style lang="scss">

.el-sub-menu__title:hover,
.el-menu-item:hover {
  background: transparent;
}

.el-scrollbar {
  .el-scrollbar__view {
    height: 100%;
  }
}






.menu-info {
  .menu-contorl {
    line-height: 52px;
    font-size: 20px;
    display: table-cell;
    vertical-align: middle;
  }
}
</style>
