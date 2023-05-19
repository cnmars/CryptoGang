<template>
  <el-container class="layout-cont">
    <el-container :class="[isSider?'openside':'hideside',isMobile ? 'mobile': '']">
      <el-row :class="[isShadowBg?'shadowBg':'']" @click="changeShadow()" />
      <!-- 左边菜单 -->
      <el-aside class="main-cont main-left gd-aside">
        <!-- 应用图标和名字 -->
        <!-- <div class="tilte" :style="{background: backgroundColor,width: `${isCollapse ?'54px':'260px'}`}">
          <img alt class="logoimg" src="@/assets/img/cat.jpg">
          <div v-if="isSider" class="tit-text" :style="{color:textColor}">CryptoGang</div>
        </div> -->
        <div class="tilte" :style="{background: 'var(--el-header-bg-color)',width: `${isCollapse ?'54px':'260px'}`}" @click="goHome">
          <img class="cglogo" alt src="">
          <div v-if="isSider" class="tit-text2 cg-logo">CryptoGang</div>
        </div>

        <Aside class="aside" />
      </el-aside>
      <!-- 分块滑动功能 -->
      <el-main class="main-cont main-right">
        <transition :duration="{ enter: 800, leave: 100 }" mode="out-in" name="el-fade-in-linear">
          <div :style="{width: `calc(100% - ${isMobile?'0px':isCollapse?'54px':'260px'})`}" class="topfix">
            <el-row>
              <el-col>
                <el-header class="header-cont">
                  <el-row class="pd-0">
                    <el-col :xs="2" :lg="1" :md="1" :sm="1" :xl="1" style="z-index:100">
                      <div class="menu-total" @click="totalCollapse">
                        <svg-icon v-if="isCollapse" icon-class="arrow-double-right" className='icon' />
                        <svg-icon v-else icon-class="arrow-double-left" className='icon' />
                      </div>
                    </el-col>
                    <el-col :xs="6" :lg="5" :md="6" :sm="6" :xl="6" :pull="1">
                      <el-breadcrumb class="breadcrumb">
                        <el-breadcrumb-item v-for="item in matched.slice(1,matched.length)" :key="item.path" >
                          {{ $t('route.'+item.name) }}
                        </el-breadcrumb-item>
                      </el-breadcrumb>
                    </el-col>
                    <!-- <el-col :xs="10" :lg="11" :md="10" :sm="10" :xl="10">
                      <div style="width:100%;margin-top:10px;">
                        <SearchBar/>
                      </div>
                    </el-col> -->
                    <el-col :xs="16" :lg="18" :md="17" :sm="17" :xl="17">
                      <!-- <el-col :xs="6" :lg="7" :md="7" :sm="7" :xl="7"> -->
                        <div class="right-box">
                          <div style="min-width:600px;margin-right:10px;margin-top: -3px;">
                            <SearchBar/>
                          </div>
                        <!-- 搜索框 -->
                        <Search />
                        <el-button v-if="userStore.userInfo.isSign && userStore.userInfo.authorityId===333" type="primary" 
                         style="margin-right: 10px;" size="small" :style="{width:'80px'}" @click="upgrade">
                          <span style="font-size: 14px;">{{ $t('upgrade.btn')}}</span>
                        </el-button>
                        <el-dropdown>
                          <template v-if="userStore.userInfo.isConnect">
                            <div class="dp-flex justify-content-center align-items height-full width-full">
                              <span class="header-avatar" style="cursor: pointer">
                                <CustomPic />
                                <span style="margin-left: 5px;font-size: medium;color: var(--el-color-primary);">{{showUsername}}</span>
                                <svg-icon icon-class="arrow-down" className='icon-down' />
                              </span>
                            </div>
                          </template>
                          <template v-else>
                              <el-button type="primary"  size="small" @click="userStore.ConnectAndLogin">
                                <span style="font-size: 14px;">{{ $t('com.btn.connect') }}</span>
                                </el-button>
                              <!-- <el-button type="primary"  size="small" @click="ConnectWallet">连接钱包</el-button> -->
                          </template>
                            
                          <template #dropdown v-if="userStore.userInfo.isConnect">
                            <el-dropdown-menu class="dropdown-group">
                              <el-dropdown-item>
                                <span style="font-weight: 600;" v-if="userStore.userInfo.authorityId===333">
                                  用户角色:普通用户
                                </span>
                                 <span style="font-weight: 600;color: chocolate;" v-if="userStore.userInfo.authorityId===666">
                                  用户角色 : VIP用户
                                  <span style="font-size: 12px;font-weight: 600;color: chocolate;">{{ vipRest }}</span>
                                </span>
                                <span style="font-weight: 600;color: red;" v-if="userStore.userInfo.authorityId===888">
                                  用户角色 : 管理员
                                </span>
                              </el-dropdown-item>
                              <template v-if="userStore.userInfo.authorities">
                                <el-dropdown-item v-for="item in userStore.userInfo.authorities.filter(i=>i.authorityId!==userStore.userInfo.authorityId)" :key="item.authorityId" @click="changeUserAuth(item.authorityId)">
                                  <span>
                                    切换为:角色2
                                  </span>
                                </el-dropdown-item>
                              </template>
                              <el-dropdown-item v-if="!userStore.userInfo.isSign" @click="userStore.ConnectAndLogin">
                                  <font-awesome-icon icon="fa-solid fa-feather" /> &ensp;签名登录
                              </el-dropdown-item>
                              <el-dropdown-item v-if="userStore.userInfo.isSign" style="cursor:default" >
                                <font-awesome-icon icon="fa-solid fa-copy" @click="copy" style="cursor:pointer"/>
                                 &ensp;地址:{{showAddress}}
                              </el-dropdown-item>
                              <el-dropdown-item v-if="userStore.userInfo.isSign"  @click="toPerson">
                                <font-awesome-icon icon="fa-solid fa-piggy-bank" />
                                 &ensp;余额:{{showBal}}
                              </el-dropdown-item>
                              <el-dropdown-item v-if="userStore.userInfo.isSign" style="cursor:default">
                                <font-awesome-icon icon="fa-solid fa-globe" />
                                 &ensp;网络:{{userStore.userInfo.showNetwork}}
                              </el-dropdown-item>
                              <el-dropdown-item v-if="userStore.userInfo.isSign"  @click="userStore.LoginOut">
                                <font-awesome-icon icon="fa-solid fa-arrow-left" />
                                &ensp;登 出
                              </el-dropdown-item>
                            </el-dropdown-menu>
                          </template>
                        </el-dropdown>
                      </div>
                    </el-col>
                  </el-row>
                </el-header>
              </el-col>
            </el-row>
            <!-- 当前面包屑用路由自动生成可根据需求修改 -->
            <!--
            :to="{ path: item.path }" 暂时注释不用-->
            <HistoryComponent ref="layoutHistoryComponent" />
          </div>
        </transition>
        <router-view v-if="reloadFlag" v-slot="{ Component }" v-loading="loadingFlag" element-loading-text="正在加载中" class="admin-box" >
          <div>
          <!-- <div style="position: fixed;box-sizing: border-box;"> -->
            <transition mode="out-in" name="el-fade-in-linear">
              <keep-alive :include="routerStore.keepAliveRouters">
                <component :is="Component" />
              </keep-alive>
            </transition>
          </div>
        </router-view>
        <!-- 注意下面这里的长度280px 是两边的margin10px + 左边的菜单栏宽度得来-->
        <div :style="{width: `calc(100% - ${isMobile?'0px':isCollapse?'54px':'280px'})`}" class="bottomfix">
          <BottomBar/>
        </div>
        <!-- <BottomInfo /> -->
        <setting />
      </el-main>
    </el-container>

  </el-container>
</template>

<script>
export default {
  name: 'Layout',
}
</script>

<script setup>
import Aside from '@/views/layout/aside/index.vue'
import HistoryComponent from '@/views/layout/aside/historyComponent/history.vue'
import Search from '@/views/layout/search/search.vue'
import BottomBar from '@/views/layout/bottomInfo/bottomBar.vue'
import CustomPic from '@/components/customPic/index.vue'
import Setting from './setting/index.vue'
import { setUserAuthority,getSignMsg ,signAndLogin} from '../../api/userApi'
import { emitter } from '@/common/bus'
// import LangSelector from '@/components/langSelector/langSelector.vue'
import SearchBar from '@/components/searchBar/searchBar.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { computed, ref, onMounted, nextTick } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useRouterStore } from '@/pinia/modules/router'
import { useUserStore } from '@/pinia/modules/userStore'
import { appConfig } from '../../global/config'
import { ethers } from "ethers";
import { LOCAL_STORAGE } from "@/global/constants";
const router = useRouter()
const route = useRoute()
const routerStore = useRouterStore()
// 三种窗口适配
const isCollapse = ref(false)
const isSider = ref(true)
const isMobile = ref(false)
const userStore = useUserStore()

const initPage = () => {
  const screenWidth = document.body.clientWidth
  if (screenWidth < 1000) {
    isMobile.value = true
    isSider.value = false
    isCollapse.value = true
  } else if (screenWidth >= 1000 && screenWidth < 1200) {
    isMobile.value = false
    isSider.value = false
    isCollapse.value = true
  } else {
    isMobile.value = false
    isSider.value = true
    isCollapse.value = false
  }
}

const goHome =()=>{
  router.push({path:'/'})
}
const showBal= computed(()=>{
  if(userStore.userInfo.balance){
    return userStore.userInfo.balance.slice(0,7)+'ETH'
  }
  else{
    return 0
  }
})

initPage()
window.aa = routerStore.asyncRouters

const ConnectWallet = async()=>{
  if(window.ethereum){
      const accounts = await window.ethereum.request({
        method: "eth_requestAccounts",
      });
      // userStore.userInfo.address=accounts[0]
      userStore.userInfo.address=accounts[0];
      userStore.userInfo.isConnect= true
      window.localStorage.setItem(LOCAL_STORAGE.USER_ADDRESS,accounts[0])
      userStore.userProvider.httpProvider.getBalance(accounts[0]).then((res)=>{
        const bal=ethers.utils.formatEther(res)
        window.localStorage.setItem(LOCAL_STORAGE.USER_BAL,bal)
        userStore.userInfo.balance = bal

      })
      userStore.userProvider.httpProvider.lookupAddress(accounts[0]).then((res)=>{
        window.localStorage.setItem(LOCAL_STORAGE.USER_ENS,res)
        userStore.userInfo.ens =res
      })
  }
}

/* 登录*/
const ConnectAndLogin = async()=>{
  if(window.ethereum){
      window.ethereum.request({
        method: "eth_requestAccounts",
      }).then(async(accounts)=>{
          // userStore.userInfo.address=accounts[0]
          userStore.userInfo.address=accounts[0];
          userStore.userInfo.isConnect= true
          window.localStorage.setItem(LOCAL_STORAGE.USER_ADDRESS,accounts[0])
          userStore.userProvider.httpProvider.getBalance(accounts[0]).then((res)=>{
            const bal=ethers.utils.formatEther(res)
            window.localStorage.setItem(LOCAL_STORAGE.USER_BAL,bal)
            userStore.userInfo.balance = bal
    
          })
          userStore.userProvider.httpProvider.lookupAddress(accounts[0]).then((res)=>{
            window.localStorage.setItem(LOCAL_STORAGE.USER_ENS,res)
            userStore.userInfo.ens =res
          })
         await userStore.ConnectAndLogin()

      }).finally(()=>{
      })
  }
}
const upgrade = () => {
  router.push({ name: 'upgrade' })
}


const showAddress= computed(()=>{
    let acbf4 = userStore.userInfo.address.slice(0,4)
    let acaf4 = userStore.userInfo.address.slice(-4)
    return `${acbf4}...${acaf4}`
})
const showUsername= computed(()=>{
    if(userStore.userInfo.ens && userStore.userInfo.ens!="null")return userStore.userInfo.ens
    else return showAddress.value
    // return userStore.userInfo.ens ? userStore.userInfo.ens : showAddress
})
const copy=()=> {
     var input = document.createElement("input"); // 创建input对象
     input.value = userStore.userInfo.address; // 设置复制内容
     document.body.appendChild(input); // 添加临时实例
     input.select(); // 选择实例内容
     document.execCommand("Copy"); // 执行复制
     document.body.removeChild(input); // 删除临时实例
     ElMessage({
          message: "复制成功",
          type: 'success'
        })
}

const loadingFlag = ref(false)
onMounted(() => {
  // 挂载一些通用的事件
  emitter.emit('collapse', isCollapse.value)
  emitter.emit('mobile', isMobile.value)
  emitter.on('reload', reload)
  emitter.on('toLogin', totalCollapse)
  emitter.on('showLoading', () => {
    loadingFlag.value = true
  })
  emitter.on('closeLoading', () => {
    loadingFlag.value = false
  })
  window.onresize = () => {
    return (() => {
      initPage()
      emitter.emit('collapse', isCollapse.value)
      emitter.emit('mobile', isMobile.value)
    })()
  }
  if (userStore.loadingInstance) {
    userStore.loadingInstance.close()
  }
  //假如用户未登录打开网站时让左边菜单栏关闭
  if(!userStore.userInfo.isSign){
    totalCollapse()
  }

})



const textColor = computed(() => {
  if (userStore.sideMode === 'dark') {
    return '#c28e0c'
    // return '#fff'
  } else if (userStore.sideMode === 'light') {
    return '#c28e0c'
    // return '#191a23'
  } else {
    return '#ffffff'
  }
})

const backgroundColor = computed(() => {
  if (appConfig.sideMode === 'dark') {
    return '#191a23'
  } else if (appConfig.sideMode === 'light') {
    return '#126b40'
  } else {
    return userStore.sideMode
  }
})

const matched = computed(() => route.meta.matched)

const changeUserAuth = async(id) => {
  const res = await setUserAuthority({
    authorityId: id
  })
  if (res.code === 0) {
    window.sessionStorage.setItem('needCloseAll', 'true')
    window.location.reload()
  }
}
// const keepAlive=computed(()=>{
//   if(!routerStore.keepAliveRouters.includes('nft')){
//     routerStore.keepAliveRouters.push('nft')
//     console.log("mainpage--------:keepAlive--->:",routerStore.keepAliveRouters.join(','));
//   }
//   return routerStore.keepAliveRouters.join(',')
// })
const reloadFlag = ref(true)
let reloadTimer = null
const reload = async() => {
  if (reloadTimer) {
    window.clearTimeout(reloadTimer)
  }
  reloadTimer = window.setTimeout(async() => {
    if (route.meta.keepAlive) {
      reloadFlag.value = false
      await nextTick()
      reloadFlag.value = true
    } else {
      const title = route.meta.title
      router.push({ name: 'Reload', params: { title }})
    }
  }, 400)
}

const isShadowBg = ref(false)
const totalCollapse = () => {
  isCollapse.value = !isCollapse.value
  isSider.value = !isCollapse.value
  isShadowBg.value = !isCollapse.value
  emitter.emit('collapse', isCollapse.value)
}
const vipRest = computed(()=>{
  let ex = new Date(userStore.userInfo.expireAt);
  let now = Date.now()
  var interval = ex - now;
  if(interval<0){
    return "(已过期)"
  }
  interval /= 1000;
  let day = Math.round(interval / 60 / 60 / 24);
  return `(剩余 ${day}天)`
})
const toPerson = () => {
  router.push({ name: 'person' })
}
const changeShadow = () => {
  isShadowBg.value = !isShadowBg.value
  isSider.value = !!isCollapse.value
  totalCollapse()
}
</script>

<style>

.el-popper.is-twitter {
  /* Set padding to ensure the height is 32px */
  padding: 6px 12px;
  margin-bottom: -5px !important;
  background: skyblue;
}
.el-popper.is-twitter .el-popper__arrow::before {
  background: skyblue;
  right: 0;
  margin-bottom: -10px !important;
}
.el-popper.is-dc {
  /* Set padding to ensure the height is 32px */
  padding: 6px 12px;
  margin-bottom: -5px !important;
  background: #6370F4;
}
.el-popper.is-dc .el-popper__arrow::before {
  background: #6370F4;
  right: 0;
  margin-bottom: -10px !important;
}

</style>
<style lang="scss">
@import '@/style/mobile.scss';

.dark{
  background-color: #191a23 !important;
  color: #fff !important;
}
.light{
  background-color: #fff !important;
  color: #000 !important;
}

.icon {
  width: 1em;
  height: 1em;
  fill: var(--el-color-primary);
}
.icon-down{
  width: 1.5em;
  height: 1.5em;
  margin-left: 2px;
  fill: var(--el-color-primary);
}

</style>
