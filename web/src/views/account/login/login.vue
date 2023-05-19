<template>
    <div>
      <div class="big" v-loading="loading" element-loading-text="等待签名登录...">
        <div class="inner">
            <el-button type="primary" size="large" :style="{width:'160px'}" @click="ConnectAndLogin">{{ $t('com.btn.connect') }}</el-button>
        </div>
      </div>
    </div>
  </template>

<script lang="ts">
    export default{
        name:'login'
    }
</script>
<script setup lang="ts">
import { ethers } from "ethers";
import { ref } from "vue";
import { LOCAL_STORAGE } from "@/global/constants";
import { useUserStore } from '@/pinia/modules/userStore'
import { emitter } from '@/common/bus'
const userStore = useUserStore()
const loading = ref(false)

const ConnectAndLogin = async()=>{
  emitter.emit('closeLoading')
  loading.value = true;
  userStore.ConnectAndLogin().then(()=>{

  }).finally(()=>{
    loading.value = false;
  })
  // if((window as any).ethereum){
  //   emitter.emit('closeLoading')
  //   loading.value = true;
  //     (window as any).ethereum.request({
  //       method: "eth_requestAccounts",
  //     }).then(async(accounts)=>{
  //         // userStore.userInfo.address=accounts[0]
  //         userStore.userInfo.address=accounts[0];
  //         userStore.userInfo.isConnect= true
  //         window.localStorage.setItem(LOCAL_STORAGE.USER_ADDRESS,accounts[0])
  //         userStore.userProvider.httpProvider.getBalance(accounts[0]).then((res)=>{
  //           const bal=ethers.utils.formatEther(res)
  //           window.localStorage.setItem(LOCAL_STORAGE.USER_BAL,bal)
  //           userStore.userInfo.balance = bal
    
  //         })
  //         userStore.userProvider.httpProvider.lookupAddress(accounts[0]).then((res)=>{
  //           window.localStorage.setItem(LOCAL_STORAGE.USER_ENS,res)
  //           userStore.userInfo.ens =res
  //         })
  //        await userStore.ConnectAndLogin()

  //     }).finally(()=>{
  //       loading.value = false;
  //     })
  // }

}


</script>


<style lang="scss" scoped>
    .big{
        width: 100%;
        // height: calc(100vh - 260px);
        height: 100%;
        background-color: rgb(244,244,244);
        position: relative;
    }
    .inner{
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }
    .inner p {
        text-align: center;
        font-size: 24px;
    }
    .inner .leftPic{
        width: 60px;
        height: 60px;
        margin-left: 44%;
        margin-top: 20px;
    }

    .svg404{
      width: 33em;
      height: 33em;
      color: #126b40;
    }

    :deep(.el-loading-spinner .circular){
    display: inline;
    height: 100px;
    width: 100px;
    -webkit-animation: loading-rotate 2s linear infinite;
    animation: loading-rotate 2s linear infinite;
  }
  :deep(.el-loading-mask .el-loading-spinner .el-loading-text) {
  font-size: 22px;
  font-weight: bold;
}
</style>
