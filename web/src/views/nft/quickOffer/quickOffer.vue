<template>
    <el-input v-model="bbb" placeholder="aaaaaaaaaa" />
    <el-input v-model="bbb" placeholder="aaaaaaaaaa" />
</template>


<script>
    export default{
       name:"quickOffer"
    }
</script>
<script setup>
import {onMounted,onUpdated,onActivated,onDeactivated,ref} from "vue";

const bbb = ref('')


</script>

<style lang="scss" scoped>

</style>