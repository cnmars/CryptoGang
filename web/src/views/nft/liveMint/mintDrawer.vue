<template>
    <div>
      <el-drawer
        v-model="drawer"
        :direction="direction"
        :before-close="handleClose"
        append-to-body="true"
        modal-class=""
      >
      <template #title >
          <span style="font-size: 20px;font-weight:bald;color: #126b40;">{{$t('setting.title')}}</span>
      </template>
  
      <warning-bar :title="$t('setting.warn')" />
        <div class="setting_body">
          
        </div>
      </el-drawer>
  
    </div>
  </template>
  
  <script lang="ts">
  export default {
    name: 'mintDrawer',
  }
  </script>
  
  <script lang="ts" setup>
  import { ref,onActivated,watch,computed } from 'vue'
  import {storeToRefs} from 'pinia'
  import { useUserStore } from '@/pinia/modules/userStore'
  import { useWalletStore } from '@/pinia/modules/walletStore'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import type { UploadInstance, UploadProps, UploadRawFile } from 'element-plus'
  import { genFileId ,ElMessage} from 'element-plus'
  import {  ethers } from "ethers"
  import { fileUtils } from '@/utils/file'
  import { ethUtils } from '@/utils/blockchain'
  import { emitter } from '@/common/bus'
  import { useI18n } from 'vue-i18n'
  
  const userStore = useUserStore()
  const wstore =useWalletStore()
  const  {walletTable} =useWalletStore()
  const  {userWallets} =storeToRefs(userStore)
  // const {walletTable} = storeToRefs(wstore)
  const walletStore =useWalletStore()
  const { t } = useI18n()
  const drawer = ref(true)
  const direction = ref('btt')

  const props=defineProps({
    show: {
      type: Boolean,
      default: true
    }
  })
  
  const handleClose = () => {
    drawer.value = false
  }
  const showSettingDrawer = () => {
    drawer.value = true
  }

 
 
  
  </script>
  
  <style lang="scss" scoped>
  // 上传文件后不展示文件名和图标
  :deep(.el-upload-list){
        display: none !important; 
  }
  .drawer-container {
    transition: all 0.2s;
    &:hover{
      right: 0
    }
    position: fixed;
    right: -5px;
    bottom: 15%;
    height: 40px;
    width: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    // z-index: 999;
    color: #fff;
    border-radius: 4px 0 0 4px;
    cursor: pointer;
    -webkit-box-shadow: inset 0 0 6px rgba(0 ,0 ,0, 10%);
    box-shadow: inset 0 0 6px rgba(0 ,0 ,0, 10%);
  }
  .setting_body{
    padding: 20px;
    .setting_card{
      margin-bottom: 20px;
    }
    .setting_content{
      margin-top: 20px;
      display: flex;
      flex-direction: column;
      >.theme-box{
       display: flex;
      }
      >.color-box{
        div{
          display: flex;
          flex-direction: column;
        }
      }
      .item{
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
        margin-right: 20px;
        .item-top{
          position: relative;
        }
        .check{
          position: absolute;
          font-size: 20px;
          color: #00afff;
          right:10px;
          bottom: 10px;
        }
        p{
          text-align: center;
          font-size: 12px;
        }
      }
    }
  }
  // 输入框选择时阴影高度
  :deep(.el-input__inner){
    height: 30px !important;
  }
  
  :deep(.el-empty__description){
    margin-top: 0px;
    margin-bottom: -10px;
  }
  :deep(.el-table__empty-text){
    line-height: 40px;
  }
  </style>
  