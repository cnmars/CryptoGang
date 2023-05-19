<template>
    <div v-if="isShow" class="message-bar" :class="href&&'can-click'" @click="open">
      <el-icon>
        <warning-filled />
      </el-icon>
      <span>
        {{ title }}
      </span>
      <el-icon  style="color:red;padding-top: 5px; right:-4px;width: 12px;height: 12px;cursor:pointer;" @click="isShow=false">
        <circle-close-filled />
      </el-icon>
    </div>
  </template>
  <script setup>
  import { WarningFilled } from '@element-plus/icons-vue'
  import { ref } from 'vue';
  const prop = defineProps({
    title: {
      type: String,
      default: ''
    },
    href: {
      type: String,
      default: ''
    }
  })
  const isShow=ref(true)
  const open = () => {
    if (prop.href) {
      window.open(prop.href)
    }
  }
  const close=()=>{
    // console.log("ppppppppppppppppppp");
    // alert("aaaaa")
  }
  </script>
  <style lang="scss" scoped>
  
  .can-click{
    cursor: pointer;
  }
  </style>
  