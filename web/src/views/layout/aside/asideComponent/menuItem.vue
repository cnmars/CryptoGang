<template>
  <el-menu-item :index="routerInfo.name" :disabled="disabled">
    <!-- 折叠 -->
    <template v-if="isCollapse">
      <!-- <el-tooltip class="box-item " effect="light" :content="content" placement="right"> -->
      <el-tooltip class="box-item " effect="light" :content="disabled ? $t('route.unavailable') :$t('route.'+routerInfo.name)"  placement="right">
        <!-- Live Mint菜单项特殊 -->
        <div v-if="routerInfo.name=='liveMint'" class="flash-block live-mints live-mints-flicker"/>
        <el-icon v-else-if="routerInfo.meta.icon">
          <font-awesome-icon :icon="routerInfo.meta.icon" />
          <!-- <component :is="routerInfo.meta.icon"/> -->
        </el-icon>
      </el-tooltip>
    </template>
    <!-- 非折叠 -->
    <template v-else>
        <el-tooltip v-if="disabled" class="box-item " effect="light" :content="$t('route.unavailable')" placement="right">
          <div class="gd-menu-item">
            <div v-if="routerInfo.name=='liveMint'" class="flash-block live-mints live-mints-flicker"/>
            <el-icon v-else-if="routerInfo.meta.icon">
              <font-awesome-icon :icon="routerInfo.meta.icon" :size="routerInfo.meta.level>1 ? '2xs' : ''"/>
              <!-- <component :is="routerInfo.meta.icon" /> -->
            </el-icon>
            <!-- <span class="gd-menu-item-title" :style="routerInfo.meta.level>1 ? {'font-weight': 100} : {}">{{disabled ? $t('route.'+routerInfo.name)+$t('route.unavailable') :$t('route.'+routerInfo.name)}}</span> -->
            <span class="gd-menu-item-title" :style="routerInfo.meta.level>1 ? {'font-weight': 100} : {}">{{$t('route.'+routerInfo.name)}}
              <!-- <span :style="routerInfo.meta.level>1 ? {'font-size': '12px'} : {}">{{$t('route.unavailable')}}</span> -->
            </span>
          </div>
        </el-tooltip>
      <div class="gd-menu-item" v-else>
        <div v-if="routerInfo.name=='liveMint'" class="flash-block live-mints live-mints-flicker"/>
        <el-icon v-if="routerInfo.meta.icon">
          <font-awesome-icon :icon="routerInfo.meta.icon" :size="routerInfo.meta.level>1 ? '2xs' : ''"/>
          <!-- <component :is="routerInfo.meta.icon" /> -->
        </el-icon>
        <!-- <span class="gd-menu-item-title" :style="routerInfo.meta.level>1 ? {'font-weight': 100} : {}">{{disabled ? $t('route.'+routerInfo.name)+$t('route.unavailable') :$t('route.'+routerInfo.name)}}</span> -->
        <span class="gd-menu-item-title" :style="routerInfo.meta.level>1 ? {'font-weight': 100} : {}">{{$t('route.'+routerInfo.name)}}
          <!-- <span :style="routerInfo.meta.level>1 ? {'font-size': '12px'} : {}">{{$t('route.unavailable')}}</span> -->
        </span>
      </div>
    </template>
  </el-menu-item>
</template>

<script>
export default {
  name: 'MenuItem',
}
</script>

<script setup>
import { ref, watch,computed } from 'vue'
import { unfinished } from '@/global/config'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
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

const disabled = computed(()=>{
  if(unfinished.indexOf(props.routerInfo.name) > -1){
    return true
  }else{
    return false
  }
})
const content = computed(()=>{
  console.log(t('route.'+routerInfo.name)+t('route.unavailable'));
  if(unfinished.indexOf(props.routerInfo.name) > -1){
    return t('route.'+routerInfo.name)+t('route.unavailable')
  }else{
    return t('route.'+routerInfo.name)
  }
})
</script>

<style lang="scss" scoped>
 .gd-menu-item{
   width: 100%;
    flex:1;
    height: 44px;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    // padding-left: 20px;
.gd-menu-item-title {
    flex:1;
    // font-size: var(--top-menu-title-size);
  }
}

.el-menu--collapse{
  .el-menu-item.is-active{
    color: var(--main-bg-color);
    //gd-changed-tag
    // color: v-bind(activeBackground);
  }
}

.el-menu-item{
  color: var(--gd-aside-normal-txt);
}

.el-menu-item.is-active{
  .gd-menu-item{
    //gd-changed-tag
    background: var(--main-bg-color);
    // background: v-bind(activeBackground) !important;
    border-radius: 4px;
    box-shadow: 0 0 2px 1px var(--main-bg-color) !important;
    // box-shadow: 0 0 2px 1px v-bind(activeBackground) !important;
    i{
      color: var(--gd-aside-active-txt);
    }
    span{
      color: var(--gd-aside-active-txt);
    }
  }
}

.el-menu-item:hover{
  .gd-menu-item{
    background: var(--gd-aside-hover-bg);
    border-radius: 4px;
    box-shadow: 0 0 2px 1px var(--gd-aside-hover-bg);
    color: var(--gd-aside-hover-txt);
  }
}

.flash-block{
  margin-left: 7px;
  margin-right: 7px;
}
.live-mints {
    width: 14px;
    height: 14px;
    // background-color: #126b40;
    background-color: var(--live-mint-dot-color);
    position: relative;
    border-radius: 50%;
  }
 
  /* 设置动画前颜色 */
  .live-mints-flicker:after {
    // background-color: #126b40;
    background-color: var(--live-mint-dot-color);
  }
 
  /* 设置动画后颜色 */
  .live-mints-flicker:before {
    background-color: rgba(255, 255, 255, 0.2);
  }
 
  /* 设置动画 */
  .live-mints-flicker:before,
  .live-mints-flicker:after {
    content: '';
    width: 16px;
    height: 16px;
    position: absolute;
    left: 50%;
    top: 50%;
    margin-left: -8px;
    margin-top: -8px;
    border-radius: 50%;
    animation: warn 1.5s ease-out 0s infinite;
  }
 
  @keyframes warn {
    0% {
      transform: scale(0.5);
      opacity: 1;
    }
 
    30% {
      opacity: 1;
    }
 
    100% {
      transform: scale(1.4);
      opacity: 0;
    }
  }

</style>
