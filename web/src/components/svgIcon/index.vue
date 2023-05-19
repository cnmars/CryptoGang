<template>
    <svg aria-hidden="true" :class="svgClass">
      <use :xlink:href="symbolId" :fill="color" />
    </svg>
  </template>
  
  <script setup lang="ts">
  import { computed } from 'vue';
  
  const props=defineProps({
    prefix: {
      type: String,
      default: 'icon',
    },
    iconClass: {
      type: String,
      required: true,
    },
    className:{
      type: String,
      default: 'svg-icon',
    },
    color: {
      type: String,
      default: ''
    }
  })
  
  const symbolId = computed(() => `#${props.prefix}-${props.iconClass}`);
  const svgClass =computed(()=>{
    return props.className
  }) 
  </script>
  
  <style scoped>
  .svg-icon {
    width: 2em;
    height: 2em;
    vertical-align: -0.15em;
    overflow: hidden;
    fill: currentColor;
    /* fill: v-bind(color); */
  }
  </style>