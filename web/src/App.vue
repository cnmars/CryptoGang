<template>
  <el-config-provider :locale="locale">
    <div id="app">
        <router-view />
    </div>
  </el-config-provider>
</template>


<script lang="ts">
export default {
  name: 'App'
}
</script>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { ElConfigProvider } from 'element-plus';

import {useAppStore} from '@/pinia/modules/app';

// 导入 Element Plus 语言包
import zhCn from 'element-plus/es/locale/lang/zh-cn';
import en from 'element-plus/es/locale/lang/en';

const app = useAppStore();

const language = computed(() => app.language);

const locale = ref();

watch(
  language,
  value => {
    locale.value = value == 'en' ? en : zhCn;
  },
  {
    // 初始化立即执行
    immediate: true
  }
);
</script>

<style lang="scss">
// 引入初始化样式
@import '@/style/main.scss';
@import '@/style/base.scss';
@import '@/style/mobile.scss';
#app {
  background: #eee;
  height: 100vh;
  overflow: hidden;
  font-weight: 400 !important;
}
.el-button{
  font-weight: 400 !important;
}
</style>
