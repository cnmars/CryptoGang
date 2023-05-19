<template>
    <el-input v-model="aa" placeholder="aaaaaaaaaa" />
</template>


<script>
    export default{
       name:"freeMint"
    }
</script>
<script setup>
import {onMounted,onUpdated,onActivated,onDeactivated,ref} from "vue";

const aa = ref('')
// onMounted(()=>{
//     alert("onMounted")

// })

// onUpdated(()=>{
//     alert("updated")

// })
// onActivated(()=>{
//     alert("activated")

// })
// onDeactivated(()=>{
//     alert("deactivated")

// })

</script>

<style lang="scss" scoped>

</style>