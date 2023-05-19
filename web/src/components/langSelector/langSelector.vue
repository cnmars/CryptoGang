<template>
  <el-dropdown class="lang-select" trigger="click" @command="handleSetLanguage">
    <div class="lang-select__icon">
      <svg-icon icon-class="language-solid" className='icon'/>
    </div>
    <template #dropdown>
      <el-dropdown-menu v-if="reloadComponentTag">
        <el-dropdown-item :disabled="app.language === 'zh-cn'" command="zh-cn">
          中文
        </el-dropdown-item>
        <el-dropdown-item :disabled="app.language === 'en'" command="en">
          English
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>
<script lang="ts">
  export default{
    name:"LangSelector"
  }
</script>
<script setup lang="ts">
import { computed,ref } from 'vue';
import {useAppStore} from '@/pinia/modules/app';
import { useI18n } from 'vue-i18n';
import { ElMessage } from 'element-plus';
import { emitter } from '@/common/bus'
const  app  = useAppStore();

const { locale } = useI18n();
const reloadComponentTag =ref(true)
const reloadThisComponnet = () => {
    reloadComponentTag.value = false;
      setTimeout(() => {
        reloadComponentTag.value = true;
      }, 0)
    }

const changeEditorLang = (()=>{
          emitter.emit('changeEditorLang')
    })
function handleSetLanguage(lang: string) {
  locale.value = lang;
  app.setLanguage(lang);
  if (lang == 'en') {
    ElMessage.success('Switch Language Successful!');
  } else {
    ElMessage.success('切换语言成功！');
  }
  changeEditorLang()
  reloadThisComponnet()
}
</script>

<style lang="scss" scoped>
.lang-select__icon {
  line-height: 20px;
}
.icon {
  width: 1.5em;
  height: 1.5em;
  fill: var(--el-color-primary);
}


</style>
