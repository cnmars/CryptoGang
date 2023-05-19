<template>
  <el-input v-model="aa" placeholder="aaaaaaaaaa" />
</template>


<script>
  export default{
     name:"Menus"
  }
</script>
<script setup>
import {onMounted,onUpdated,onActivated,onDeactivated,ref} from "vue";

const aa = ref('')


</script>

<style lang="scss" scoped>

</style>