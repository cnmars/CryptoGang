<template>
  <div @click="clickFull">
    <el-tooltip v-if="isShow" :content="$t('navbar.tip.full')" effect="light" placement="bottom">
        <svg-icon  icon-class="screen-full" className='icon' />
    </el-tooltip> 
    <el-tooltip v-else :content="$t('navbar.tip.exit')" effect="light" placement="bottom">
      <svg-icon  icon-class="screen-exit" className='icon' />
    </el-tooltip>
  </div>
</template>

<script>
export default {
  name: 'Screenfull',
}
</script>

<script setup>
import screenfull from 'screenfull' // 引入screenfull
import { onMounted, onUnmounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
defineProps({
  width: {
    type: Number,
    default: 22
  },
  height: {
    type: Number,
    default: 22
  },
  fill: {
    type: String,
    default: '#48576a'
  }
})

onMounted(() => {
  if (screenfull.isEnabled) {
    screenfull.on('change', changeFullShow)
  }
})

onUnmounted(() => {
  screenfull.off('change')
})

const clickFull = () => {
  if (screenfull.isEnabled) {
    screenfull.toggle()
  }
}

const isShow = ref(true)
const changeFullShow = () => {
  isShow.value = !screenfull.isFullscreen
}

</script>

<style scoped lang="scss">
.screenfull-svg {
  width: 16px;
  height: 16px;
  cursor: pointer;
  vertical-align: middle;
  margin-right: 32px;
  fill: rgba(0, 0, 0, 0.45);
}

.icon {
  width: 1.1em;
  height: 1.1em;
  fill:var(--el-color-primary);
}

</style>
