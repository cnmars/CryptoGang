<template>
  <el-input v-model="aaaa" placeholder="aaaaaaaaaa" />
</template>


<script>
  export default{
     name:"Person"
  }
</script>
<script setup>
import {onMounted,onUpdated,onActivated,onDeactivated,ref} from "vue";

const aaaa = ref('')


</script>

<style lang="scss" scoped>

</style>