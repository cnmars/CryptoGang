<template>
  <div>
    <router-view v-slot="{ Component,route }" >
      <transition mode="out-in" name="el-fade-in-linear">
        <keep-alive :include="routerStore.keepAliveRouters">
          <component :is="Component" :key="route.fullPath"/>
        </keep-alive>
      </transition>
    </router-view>
  </div>
</template>

<script>
export default {
  name: 'nft'
}
</script>
<script setup>
import { useRouterStore } from '@/pinia/modules/router'
const routerStore = useRouterStore()

// const keepAlive=computed(()=>{
//   if(!routerStore.keepAliveRouters.includes('nft')){
//     routerStore.keepAliveRouters.push('nft')
//   }
//   return routerStore.keepAliveRouters.join(',')
// })
</script>
