<template>
  <el-sub-menu ref="subMenu" :index="routerInfo.name">

    <template #title>
      <div v-if="!isCollapse" class="gd-subMenu">
        <el-icon v-if="routerInfo.meta.icon">
          <font-awesome-icon :icon="routerInfo.meta.icon" />
        </el-icon>
        <span>{{ $t('route.'+routerInfo.name)}}</span>
      </div>
      <template v-else>
        <el-icon v-if="routerInfo.meta.icon">
          <font-awesome-icon :icon="routerInfo.meta.icon"/>
        </el-icon>
        <span>{{ $t('route.'+routerInfo.name)}}</span>
      </template>
    </template>
    <slot/>
  </el-sub-menu>
</template>

<script>
export default {
  name: 'AsyncSubmenu',
}
</script>

<script setup>
import { ref, watch } from 'vue'
// import { useI18n } from 'vue-i18n'
// const { t } = useI18n()

const props = defineProps({
  routerInfo: {
    default: function() {
      return null
    },
    type: Object
  },
  isCollapse: {
    default: function() {
      return false
    },
    type: Boolean
  },
  theme: {
    default: function() {
      return {}
    },
    type: Object
  }
})



</script>

<style lang="scss" scoped>
.el-sub-menu{
  :deep(.el-sub-menu__title){
      padding: 6px;
      color: var(--gd-aside-normal-txt);
      // color: red;
  }
}

  .is-active:not(.is-opened){
  :deep(.el-sub-menu__title) .gd-subMenu{
      flex:1;
      height: 100%;
      line-height: 44px;
      //gd-changed-tag
      background: var(--main-bg-color) !important;
      // background: v-bind(activeBackground) !important;
      border-radius: 4px;
      box-shadow: 0 0 2px 1px var(--main-bg-color)!important;
      box-shadow: 0 0 2px 1px var(--gd-aside-bg-active)!important;
      i{
        color: var(--gd-aside-active-txt);
      }
      span{
        color:var(--gd-aside-active-txt);
      }
    }
  }

.gd-subMenu {
  padding-left: 4px;
  // font-size: var(--top-menu-title-size);
}
</style>
