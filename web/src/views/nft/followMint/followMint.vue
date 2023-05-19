<template>
    <el-input class="gd-input" v-model="aaaa"  autocomplete="off" placeholder="粘贴交易哈希,不建议手动输入" @input="changeTxHash" />
                           
    
</template>


<script>
    export default{
       name:"followMint"
    }
</script>
<script setup>
import { ref,reactive,toRef, onActivated,onMounted } from 'vue'

var aaaa =ref()

</script>

<style lang="scss" scoped>
.hash-mint-warp{
    width: 99.9%;
    height: auto;
    box-shadow: var(--el-box-shadow-light);
    background-color: #fff;
}
.hash-mint-form-wrap{
    // border:#909090 1px solid;
    margin-top: 60px;
    padding-left: 40px;
    padding-right: 40px;
    .el_input_inner{
        width: 100%;
    }
    .el-form-item__label {
      color: #000000;
      font-size: 15px;
    }
}
.private-key-wrap{
    padding-top: 80px;
    padding-left: 40px;
    padding-right: 40px;

}
.pv-btn-list{
    padding-left: 90px;
}
.start-wrap{
    margin-top: 40px;
    
}
.result-show-wrap{
    margin-top: 30px;
    padding-bottom: 160px;
}
.result-table{
    width: 100%;
    text-align: left;
    border-radius: 2px 2px 0 0;
    border-collapse: initial;
    border-spacing: 0;
}


</style>