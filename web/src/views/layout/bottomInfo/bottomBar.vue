<template>
  <div class="left-info"></div>
  <div class="middle-info">
    <el-tooltip :content="`${currentName}点击查看操作手册`" raw-content placement="top" effect="doc">
      <div style="font-size: 12px;" class="info-item">
        <a class="info-item" aria-label="Cryptogang" :href="url" target="_blank" rel="noopener noreferrer">
          <!-- <el-icon style="vertical-align: -18%;color: var(--el-color-primary);"><QuestionFilled /></el-icon> -->
          <img alt="" src="@/assets/img/book.png" style="width: calc(var(--bottom-bar-h) - 3px);height: calc(var(--bottom-bar-h) - 3px);">
          <span style="color: #fff;vertical-align:+25%;">如何操作?</span>
        </a>
      </div>
      </el-tooltip>
  </div>
  <div class="right-info">
    <el-tooltip content="Twitter暂未开通" raw-content placement="top" effect="twitter">
      <a class="info-item" aria-label="Twitter" href="https://twitter.com/C0NTAC_" target="_blank" rel="noopener noreferrer">
        <svg-icon icon-class="twitter-w" class-name="twitter-icon"></svg-icon>
        <!-- <img alt="" src="@/assets/img/twitter.png" style="width: calc(var(--bottom-bar-h) - 3px);height: calc(var(--bottom-bar-h) - 3px);"> -->
        <!-- <svg-icon icon-class="twitter" class-name="twitter-icon" color="skyblue"></svg-icon> -->
      </a>
    </el-tooltip>
    <el-tooltip content="Discord暂未开通" raw-content placement="top" effect="dc">
      <a class="info-item" aria-label="Twitter" href="https://cryptogang.vip" target="_blank" rel="noopener noreferrer">
        <svg-icon icon-class="dc-w" class-name="twitter-icon"></svg-icon>
        <!-- <svg-icon icon-class="discord" class-name="twitter-icon" color="#6370F4"></svg-icon> -->
        <!-- <img alt="" src="@/assets/img/dc.png" style="width: calc(var(--bottom-bar-h) - 3px);height: calc(var(--bottom-bar-h) - 3px);"> -->
      </a>
    </el-tooltip>
  </div>
    <!-- <div style="display: flex;flex-direction: row;justify-content: space-between;">
    </div> -->
  </template>
  
  <script lang="ts">
  export default {
    name: 'BottomBar'
  }
  </script>

  <script lang="ts" setup>
  import { useRoute } from 'vue-router'
  import { computed } from 'vue'
  import { useI18n } from 'vue-i18n'
const { t } = useI18n()
  const route = useRoute()
  //对于在操作手册中没有页面的功能过滤，跳转的手册首页
  const currentName = computed(()=>{
    if(['upgrade','login','userFeedBack'].indexOf(route.name as string)< 0){
      return `${t('route.'+(route.name as string))} 如何操作？<br>`
    }else{
      return ""
    }
  })
  const url = computed(()=>{
    if(['upgrade','login','userFeedBack'].indexOf(route.name as string)< 0){
      return `https://docs.cryptogang.vip/content/${route.name as string}.html`
    }else{
      return "https://docs.cryptogang.vip"
    }
  })
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
 .el-popper.is-doc {
   /* Set padding to ensure the height is 32px */
   padding: 6px 12px;
   margin-bottom: -5px !important;
   background: rgb(111, 111, 111);
 }
 .el-popper.is-doc .el-popper__arrow::before {
   background: rgb(111, 111, 111);
   right: 0;
   margin-bottom: -10px !important;
 }
 
 </style>
 <style lang="scss">
 @import '@/style/mobile.scss';
 .left-info{
   display: flex;
   flex-direction: row;
   gap: 5px;
 }
 .middle-info{
   display: flex;
   flex-direction: row;
   gap: 2px;
  //  padding: 3px;
   padding-right: 10px;
 } 
 .right-info{
   display: flex;
   flex-direction: row;
   gap: 2px;
   padding-right: 10px;
 }
 .info-item{
   padding: 1px 5px;
 }
 .info-item:hover{
   background-color: rgb(111, 111, 111);
 }
 .twitter-icon{
   height: 1.2em;
   width: 1.2em;
 }
  </style>
  