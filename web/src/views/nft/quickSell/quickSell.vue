<template>
    <div class="scroll">
    <el-row class="background" type="flex" justify="space-between">
      <el-col span="2" style="height: 40px">
        <el-image style="width: 20px;height: 20px;margin-top: 10px;margin-left: 12px" radius fit="cover"></el-image>
      </el-col>
      <el-col span="21">
        <ul id="ulll" :class="{ anim: animate == true }">
          <li v-for="(item, index) in items" :key="index">
            <el-row type="flex" justify="space-between">
              <el-col span="5" class="single_line">
                {{"["+ item.userName + "]" }}
              </el-col>
              <el-col span="13" class="single_line" align="center">
                {{$t('chenggonggoumai_nft')}}
              </el-col>
              <el-col span="6" style="color: #35C79B;padding-right: 8px" class="single_line" align="end">
                {{ item.time + $t('fenzhongqian') }}
              </el-col>
            </el-row>
          </li>
        </ul>
      </el-col>
    </el-row>
 
</div>
</template>

<script lang="ts">
    export default{
        name:"quickSell"
    }
</script>

<script lang="ts" setup>
import { ref,onActivated,onMounted,computed } from 'vue'
import { useAppStore } from "@/pinia/modules/app";
const hot = "🔥 14 Gwei"
const v = "Ξ 0.026"
const animate= ref(false)
const items= ref([
        {userName:'小明',time:1},
        {userName:'小明1',time:2},
        {userName:'小明2',time:3},
        {userName:'小明2',time:3},
        {userName:'小明2',time:3},
        {userName:'小明2',time:3},
        {userName:'小明2',time:3},
        {userName:'小明2',time:3},
      ])
 const newItem= ref([
        {userName:'小明',time:1},
        {userName:'小明1',time:2},
        {userName:'小明2',time:3},
      ])
     
const scroll=()=> {
    if(newItem.value.length>0){
        animate.value = true; // 为true时将动画样式赋值给布局，见布局处引用
        setTimeout(() => {
          //  这里直接使用了es6箭头写法，省去了处理this指向偏移问题
          // items.value.push(items[0]); // 将数组的第一个元素添加到数组的
          items.value.unshift(newItem.value[0])
          // items.value.push(newItem.value[0]); // 将数组的第一个元素添加到数组的
          newItem.value.shift()
          // items.value.shift()
    
    
    
          items.value.pop(); //删除数组的第一个元素
          animate.value = false; // margin-top为0的时候取消过渡动画，实现无缝滚动
        }, 500);//每条滚动时长0.5s

    }
    }
const getScrollMessage=()=>{
      //这里用网络请求替换真实数据，把下面的代码放到响应成功的线程里
    //   setInterval(scroll, 2000);
    setInterval(()=>{
        newItem.value.push({userName:'1111',time:1})
    },1000)
    setInterval(scroll, 2000)
}

onMounted(() => {

  
  // const ws = new WebSocket('ws://localhost:3002/ws')
// const wsInit =async()=>{
//     const ws = window.WebSocket ? new WebSocket('ws://localhost:4728/system/ws') : null
//     ws.onopen = function () {
//         console.log('success!')
//         ws.send('{"action":"subscribe","topic":"gas"}')
//         // ws.send('startting...')
//       }
//       ws.onclose = function () {
//         console.log('close')
//       }
//       ws.onerror = function () {
//         console.log('error')
//       }
//       ws.onmessage = function (e) {
//         console.log(e.data)
//       }
//     //      
// }
// wsInit()



})
</script>

<style scoped lang="scss">
.scroll {
  /*几个高度要一致*/
  height: 40px;
  font-family: SourceHanSansSC-Regular;
  font-size: 10px;
  font-weight: normal;
  font-stretch: normal;
  line-height: 0vw;
  letter-spacing: 0vw;
  color: #666666;
  overflow: hidden;
}
.anim {
  transition: all 0.5s;
  /*几个高度要一致*/
  margin-top: 40px;
//   margin-bottom: -40px;
}
#ulll li {
  list-style: none;
  /*几个高度要一致*/
  line-height: 40px;
  /*几个高度要一致*/
  height: 40px;
}
.background{
  background: #FFFFFF;
  border-radius: 8px;
  height: 40px;
}
.single_line {
    white-space:nowrap;	/*溢出不换行*/
    text-overflow: ellipsis;
    overflow:hidden;/*溢出部分隐藏*/
    line-clamp:1;
    -webkit-line-clamp:1;				/*显示的行数*/
 
}
</style>
