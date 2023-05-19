<template>
    <div>
        <div class="index_container__twYoi">
            <PendingList style="width: 35%;"></PendingList>
            <MintDetail style="width: 64.2%;margin-left: 0.8%;"></MintDetail>
        </div>

    </div>
</template>

<script lang="ts">
    export default{
        name:"liveMint"
    }
</script>

<script lang="ts" setup>
import { ref,onActivated,onMounted,computed } from 'vue'
import  MintDetail from "./detail.vue";
import  PendingList from "./list.vue";

</script>

<style scoped lang="scss">

.livemint-wrap{
    width: 99.9%;
    height: auto;
    min-height: 880px;
    box-shadow: var(--el-box-shadow-light);
    background-color: rgb(81, 74, 74);
    display: flex;
    flex-direction: column;
    align-items: center;
  }
.index_container__twYoi {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    box-shadow: var(--el-box-shadow-light);
    // padding-top: 64px;
    // width: 100vw;
    height: 100%;
    background-color: var(--main-bg-color);
    overflow: hidden;
    overflow-x: hidden;
}
  .list-wrap{
    width: 30%;
  }
  .pend-wrap{
    width: 70%;
  }
</style>
