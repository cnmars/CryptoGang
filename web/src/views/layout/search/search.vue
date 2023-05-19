<template>
  <div class="search-component">
    <div v-if="btnShow" class="user-box" >
      <div  :class="[reload ? 'reloading' : '']" @click="handleReload" >
        <el-tooltip :content="$t('navbar.tip.ref')" effect="light" placement="bottom">
          <svg-icon icon-class="refresh" className='icon' />
        </el-tooltip>
      </div>
    </div>
    <div v-if="btnShow" class="user-box" >
      <Screenfull class="search-icon" :style="{cursor:'pointer'}" />
    </div>
    <div v-if="btnShow" class="user-box" >
      <el-tooltip :content="$t('navbar.tip.sev')" effect="light" placement="bottom">
        <svg-icon icon-class="service" className='icon' @click="centerDialogVisible = true"/>
        <!-- <svg-icon icon-class="service" className='icon' @click="toService"/> -->
        </el-tooltip>
    </div>
    <div v-if="btnShow" class="user-box" >
      <LangSelector class="search-icon" style="cursor:pointer"/>
    </div>

    <el-dialog v-model="centerDialogVisible" title="wechat" width="260px" center>

      <img src="@/assets/img/contac.jpg" alt="" style="width:200px;height:200px;">
  </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'BtnBox',
}
</script>

<script setup>
import Screenfull from '@/views/layout/screenfull/index.vue'
import LangSelector from '@/components/langSelector/langSelector.vue'
import { ElMessageBox } from 'element-plus'
import { emitter } from '@/common/bus'
import { ref, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { useRouterStore } from '@/pinia/modules/router'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const centerDialogVisible = ref(false)
const router = useRouter()

const routerStore = useRouterStore()

const value = ref('')
const changeRouter = () => {
  router.push({ name: value.value })
  value.value = ''
}

const show = ref(false)
const btnShow = ref(true)
const hiddenSearch = () => {
  show.value = false
  setTimeout(() => {
    btnShow.value = true
  }, 500)
}

const searchInput = ref(null)
const showSearch = async() => {
  btnShow.value = false
  show.value = true
  await nextTick()
  searchInput.value.focus()
}

const reload = ref(false)
const handleReload = () => {
  reload.value = true
  emitter.emit('reload')
  setTimeout(() => {
    reload.value = false
  }, 500)
}
const toService = () => {

  ElMessageBox.confirm(`
        <div style="flex-direction:row;display:flex;justify-content:flex-between;gap:5px;">
        <div style="display:flex;justify-content:flex-end;">
          <img src="${image}" alt="" style="width:60px;height:60px;">
        </div>
        </div>
          `, 'Wechat', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: '我知道了',
          cancelButtonText: '取消'
        })
  // window.open('https://twitter.com/C0NTAC_')
}

</script>
<style scoped lang="scss">

:deep(.el-dialog .el-dialog__header) {
    // padding: 2px 20px 12px 20px;
     border-bottom: none;
}

:deep(.el-dialog .el-dialog__body) {
    padding: 4px 1px;
    display: flex;
    justify-content: center;
}
.reload{
  font-size: 18px;
}

.reloading{
  animation:turn 0.5s linear infinite;
}
@keyframes turn {
  0%{-webkit-transform:rotate(0deg);}
  25%{-webkit-transform:rotate(90deg);}
  50%{-webkit-transform:rotate(180deg);}
  75%{-webkit-transform:rotate(270deg);}
  100%{-webkit-transform:rotate(360deg);}
}

.right-menu-item {
      display: inline-block;
      padding: 0 8px;
      height: 100%;
      font-size: 14px;
      color: #5a5e66;
      vertical-align: text-bottom;

      &.hover-effect {
        cursor: pointer;
        transition: background 0.3s;

        &:hover {
          background: rgba(0, 0, 0, 0.025);
        }
      }
    }
.icon {
  width: 1.5em;
  height: 1.5em;
  fill:var(--el-color-primary);
}

</style>
