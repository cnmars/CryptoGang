<template>
  <div>
    <el-button v-show="!drawer"  class="drawer-container" color="fff" @click="showSettingDrawer" >
    <!-- <el-button v-show="!drawer"  class="drawer-container" color="var(--el-color-primary)" @click="showSettingDrawer" > -->
      <template #icon>
        <font-awesome-icon icon="fa-solid fa-screwdriver-wrench fa-2xl" color="#000"/>
        <!-- <font-awesome-icon icon="fa-solid fa-skull fa-2xl" size="" color="#000"/> -->
        <!-- <font-awesome-icon icon="fa-solid fa-g fa-2xl" size=""/> -->
      </template>
    </el-button>
    <el-drawer
      v-model="drawer"
      :direction="direction"
      :before-close="handleClose"
    >
    <template #title >
        <span style="font-size: 20px;font-weight:bald;color: #126b40;">{{$t('setting.title')}}</span>
    </template>

    <warning-bar :title="$t('setting.warn')" />
      <div class="setting_body">
        <div style="margin-bottom:20px">
          <span @click.stop="toggleDark()" style="margin-right: 10px;font-size: 18px;">{{$t('setting.mode')}}</span>
          <el-switch size="small" v-model="isDark" style="vertical-align: -10%;" @change="changeMode"/>
        </div>
        <div class="private-key-wrap">
          <el-collapse accordion="true" v-model="defaultOpen" >
                <span slot="title" style="font-size: 18px;">
                    <svg-icon icon-class="node" style="width: 1.3em; height: 1.3em;"/>
                    {{$t('setting.label.node')}}
                </span>
                <el-collapse-item :name="1" style="margin-top: -40px;">
                    <div style="display: flex;width: 100%;margin-top: 20px;">
                      <el-form-item  style="height: 30px;width:80%;">
                        <template #label>
                          <span style="width:50px;text-align: right;">HTTPS:</span>
                        </template>
                        <el-input  v-model="userStore.userInfo.httpProvider" :placeholder="$t('setting.ph.http')"  style="height: 30px;"/>
                      </el-form-item>
                      <el-form-item >
                        <el-button v-if="userStore.userInfo.useDefaultHttp" type="primary"  @click="setProvider(1)" style="height: 28px;margin-left: 5%;">
                            <template #icon>
                              <el-icon><CircleCheckFilled /></el-icon>
                              <svg-icon icon-class="confirm" style="width: 1.3em; height: 1.3em;"/>
                          </template>
                          {{$t('setting.btn.set')}}
                        </el-button>
                        <el-button v-else type="warning"  @click="setDefault('http')" style="height: 28px;margin-left: 5%;">
                            <template #icon>
                              <el-icon><RefreshLeft /></el-icon>
                              <!-- <svg-icon icon-class="confirm" style="width: 1.3em; height: 1.3em;"/> -->
                          </template>
                          {{$t('setting.btn.reset')}}
                        </el-button>
                      </el-form-item>
                    </div>

                    <div style="display: flex;width: 100%;margin-top: 10px;">
                      <el-form-item  style="height: 30px;width:80%;">
                        <template #label>
                          <span style="width:50px;text-align: right;">WSS:</span>
                        </template>
                        <el-input v-model="userStore.userInfo.wssProvider" :placeholder="$t('setting.ph.wss')" style="height: 30px;"/>
                      </el-form-item>
                      <el-form-item >
                        <el-button v-if="userStore.userInfo.useDefaultWss"  type="primary"  @click="setProvider(2)" style="height: 28px;margin-left: 5%;">
                            <template #icon>
                              <el-icon><CircleCheckFilled /></el-icon>
                              <!-- <svg-icon icon-class="confirm" style="width: 1.3em; height: 1.3em;"/> -->
                          </template>
                          {{$t('setting.btn.set')}}
                        </el-button>
                        <el-button v-else type="warning"  @click="setDefault('wss')" style="height: 28px;margin-left: 5%;">
                            <template #icon>
                              <el-icon><RefreshLeft /></el-icon>
                              <!-- <svg-icon icon-class="confirm" style="width: 1.3em; height: 1.3em;"/> -->
                          </template>
                          {{$t('setting.btn.reset')}}
                        </el-button>
                      </el-form-item>
                    </div>


                </el-collapse-item>
            </el-collapse>
            <el-collapse accordion="true" v-model="defaultOpen"  style="margin-top: 40px;">
                <span slot="title" style="font-size: 18px;">
                    <svg-icon icon-class="keys" style="width: 1.3em; height: 1.3em;"/>
                    {{$t('setting.label.pvk')}}
                </span>
                <el-collapse-item :name="1" style="margin-top: -40px;">
                    <el-row >
                            <el-col :span="19">
                                <textarea :rows=5 v-model="pvString" autosize="{minRows:2,maxRows:8}" class="pv-text-area"
                                    style ="margin-left: 2px;
                                    font-size: 12px;
                                    max-width: 100%;
                                    min-width: 100%;
                                    height: auto;
                                    min-height: 120px;
                                    line-height: 1.5715; 
                                    vertical-align: bottom; 
                                    transition: all .3s,height 0s; 
                                    transition-duration: 0.3s, 0s; 
                                    transition-timing-function: ease, ease; 
                                    transition-delay: 0s, 0s; 
                                    transition-property: all, height;
                                    outline-color:#126b40;" 
                                    :placeholder="txtPlaceholder"
                                ></textarea>
                            </el-col>
                            <el-col :span="5" class="pv-btn-list" style="display: flex;flex-direction:column;padding-left: 10px;width: 150px;">
                                <el-upload
                                    class="upload-demo"
                                    ref="upload"
                                    :on-remove="handleRemove"
                                    :on-change="handleChange" 
                                    :on-success="handleSuccess"
                                    :on-exceed="handleExceed"
                                    :auto-upload="false"
                                    :multiple="true"
                                    :limit="1"
                                    list-type="txt" style="display: flex; ">
                                    <el-button slot="trigger" size="small" type="primary" style="width: 80px;margin-left: 5px;margin-top: -2px;"> {{$t('setting.btn.choose_file')}}</el-button>
                                </el-upload>
                                <!-- <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div> -->
                                <el-button  size="small" type="info" @click="clearPvKey" style="width: 80px;margin-top: 14px;margin-left: 5px"> {{$t('setting.btn.clear_list')}}</el-button>
                                <el-button  size="small" type="success" @click="parsePvKey" style="width: 80px;margin-top: 14px;margin-left: 5px;"> {{$t('setting.btn.parse')}}</el-button>
                    
                            </el-col>
                        </el-row>

                        <el-row style="margin-top: 5px;">
                          <el-col :span="24">
                            <el-table :data="walletTable"  border 
                                  :header-cell-style="{ 'text-align': 'center','heifht':'30px' }"
                                  :cell-style="{ 'text-align': 'center' }" 
                                  style="width: 98%;height:180px;margin-left: 2px;text-align: center" 
                                  v-loading="tableLoading" :element-loading-text="$t('hashmint.ld.table')" 
                              >
                              <template #empty>
                                    <el-empty :image-size="50" :description="t('hashmint.description')"></el-empty>
                                </template>
                                <el-table-column prop="address" :label="$t('hashmint.label.address')">
                                    <template #default="scope">
                                        <div>
                                        {{`${scope.row.address.slice(0,4)}...${scope.row.address.slice(-4)}`}}
                                        </div>
                                    </template>
                                </el-table-column>
                                <el-table-column prop="balance" :label="$t('hashmint.label.balance')">
                                    <template #default="scope">
                                        <div v-loading="scope.row.isBalanceLoading">
                                            {{`${scope.row.balance.slice(0,6)}ETH`}}
                                        </div>
                                    </template>
                                </el-table-column>
                                <!-- <el-table-column prop="hash" :label="$t('hashmint.label.hash')"></el-table-column>
                                <el-table-column prop="status" :label="$t('hashmint.label.status')"></el-table-column> -->
                                <el-table-column prop="operation" :label="$t('hashmint.label.op')">
                                    <template #default="scope">
                                            <el-popconfirm 
                                                confirm-button-text="Yes"
                                                cancel-button-text="No"
                                                icon="InfoFilled"
                                                icon-color="#126b40"
                                                title="Are you sure to delete this?"
                                                @confirm="removeWallet(scope.$index, scope.row)"
                                                @cancel="">
                                            <template #reference>
                                                <el-button size="small" type="danger" plain >{{$t('hashmint.btn.delete')}}</el-button>
                                            </template>
                                        </el-popconfirm>
                                      </template>
                                </el-table-column>

                              </el-table>
                          </el-col>
                        </el-row>
                </el-collapse-item>
            </el-collapse>
        </div>
      </div>
    </el-drawer>

  </div>
</template>

<script lang="ts">
export default {
  name: 'Setting',
}
</script>

<script lang="ts" setup>
import { ref,onActivated,watch,computed } from 'vue'
import {storeToRefs} from 'pinia'
import { useUserStore } from '@/pinia/modules/userStore'
import { useWalletStore } from '@/pinia/modules/walletStore'
import WarningBar from '@/components/warningBar/warningBar.vue'
import type { UploadInstance, UploadProps, UploadRawFile } from 'element-plus'
import { genFileId ,ElMessage} from 'element-plus'
import {  ethers } from "ethers"
import { fileUtils } from '@/utils/file'
import { ethUtils } from '@/utils/blockchain'
import { emitter } from '@/common/bus'
import { useI18n } from 'vue-i18n'

import { useDark, useToggle } from '@vueuse/core'
const userStore = useUserStore()
const isDark = useDark({
  storageKey: 'theme',
  valueDark: 'dark',
  valueLight: 'light',
  initialValue:'dark',
})
const toggleDark = useToggle(isDark)
const changeMode =()=>{
  userStore.userInfo.theme = (isDark.value ? 'dark':'light')
  debugger
  userStore.setTheme(userStore.userInfo.theme)
}

const wstore =useWalletStore()
const  {walletTable} =useWalletStore()
// const  {userWallets} =storeToRefs(userStore)
// const {walletTable} = storeToRefs(wstore)
// const walletStore =useWalletStore()
const { t } = useI18n()
const drawer = ref(false)
const direction = ref('rtl')
const txtPlaceholder = ref(t('setting.ph.txt_pvk'))
let pvFile =null
const walletList = ref([])
const tableLoading = ref(false)
const upload = ref<UploadInstance>()
const defaultOpen =ref(1) //为了使el-collapse默认打开
const handleExceed: UploadProps['onExceed'] = (files) => {
  upload.value!.clearFiles()
  const file = files[0] as UploadRawFile
  file.uid = genFileId()
  upload.value!.handleStart(file)
}


//   watch(() => [...userWallets.WalletArray], (nArray, oArray)=>{
//     const n = new Map(nArray);
//     const o = new Map(oArray);
//     walletList.value=nArray
//     console.log(nArray)
//     // console.log(n.get('width'), n.get('height'), o.get('width'), o.get('height'));
// },{deep:true});
const node =ref({
  https:'',
  wss:''
})
const pvString=ref('') 
const handleClose = () => {
  drawer.value = false
}
const showSettingDrawer = () => {
  drawer.value = true
}
const setProvider = (type)=>{
  const {userProvider} = useUserStore()
  switch (type) {
    case 1:
        if(!userStore.userInfo.httpProvider){
          ElMessage({
            message:"请输入url",
            type:'error'
          })
          return
        }
        userProvider.setHttpProvider(userStore.userInfo.httpProvider)
      break;
    case 2:
          if(!userStore.userInfo.wssProvider){
            ElMessage({
              message:"请输入url",
              type:'error'
            })
            return
          }
          userProvider.setWssProvider(userStore.userInfo.wssProvider)
        break;
    default:
      break;
  }   
}
const setDefault = (type)=>{
  const {userProvider} = useUserStore()
  userProvider.setToDefault(type)
   
}
// const changeMode = (e) => {
//   if (e === null) {
//     userStore.changeSideMode('dark')
//     return
//   }
//   userStore.changeSideMode(e)
// }
const handleChange =(file,list)=>{
    const isH5 = file.name.split(".")[1] === "txt";
    const isLt2G = file.size / 1024 / 1024 < 5;
    if (!isH5) {
        ElMessage({
          message: "仅支持上传.txt文件",
          type: 'error'
        })
        upload.value!.clearFiles()
	}
    if (!isLt2G) {
        ElMessage({
          message: "上传文件大小不能超过 5KB!",
          type: 'error'
        })
        upload.value!.clearFiles()
    }
    if (list.length > 1 && file.status !== "fail") {
      list.splice(0, 1);
    } else if (file.status === "fail") {
        ElMessage({
          message: "失败，请重新上传",
          type: 'error'
        })
      list.splice(0, 1);
    }
    txtPlaceholder.value= `${t('setting.ph.selected')}${file.name}`
    pvFile = file
}
const handleSuccess =(file)=>{
    pvFile = file
}

const handleRemove: UploadProps['onRemove'] = (file, uploadFiles) => {
  console.log(file, uploadFiles)
}
const parsePvKey =()=>{
  
    walletList.value = []
    if (!pvString.value && pvFile!=null){
        fileUtils.readFile(pvFile).then((res)=>{
            pvString.value = res as any
            doParse(pvString.value)
        })
    }else{
      doParse(pvString.value)
    }
}

const doParse =(content:string)=>{
  let pv = content.trim()
  if(!pv){
    ElMessage({
      message:"请录入私钥！",
      type:'error'
    })
    return
  }
  let pv2 = pv.replace(/\r\n/g, ",").replace(/\n/g, ",")
  let pvArr = pv2.split(',')
  for (let i in pvArr){
    if(!ethUtils.isValidPVKey(pvArr[i])){
      ElMessage({
        message: '您的录入含有不正确的私钥！',
        type: 'error'
        })
      return
    }
  }
  tableLoading.value = true
  let flag =true
  Promise.allSettled(pvArr.map(ethUtils.getBalanceByPvk)).then((statuses)=>{
    statuses.forEach((i,v)=>{
      if(i.status==='fulfilled'){
        let wal = (i as any).value     
        if(!(Number(wal.balance)>0)as boolean){
          ElMessage({
                message: '已略过没有余额的钱包(无法参与mint)',
                type: 'warning'
            })
        }else{
          wal['pvk']=pvArr[v]
          walletList.value.push(wal)
          emitter.emit('addWallet',(wal))
          // addWalletByPvKey(pvArr[v])
          userStore.userWallets.addWalletByPvKey(pvArr[v],"Default")
        }
      }else{//多是网络错误
        if(flag){//因为这里一旦报错会报多个
          ElMessage({
            message: "节点连接出错！",
            type: 'info'
          });
          flag =false
        }
      
      }
    })
  }).catch((err)=>{
  }).finally(()=>{
    tableLoading.value = false
    ElMessage({
        message: "解析结束",
        type: 'success'
    });
  })

}
const clearPvKey =()=>{
    pvString.value='';
    txtPlaceholder.value=t('setting.ph.txt_pvk');
    pvFile = null
    upload.value!.clearFiles()
    //是否清空储存？
    // userStore.hashMintPvkArr = []
}

const removeWallet =(index, row)=> {
      var list = walletList.value;
      let addr =''
      for (var i = 0; i < list.length; i++) {
          if (row.address == list[i].address) {
            addr = row.address
            userStore.userWallets.removeWallet(addr)
            walletList.value.splice(i, 1);
            let wal = {
                  address:addr,
                  balance:row.balance,
              }
            //通知其他页面也移除这个钱包
            emitter.emit('removeWallet',wal)
            // hashPvmap.value.delete()
              // this.$message({ message: '删除成功', duration: 2000, type: 'success' });
           }
       }  
}


</script>

<style lang="scss" scoped>
:deep(.el-icon svg ){
    height: 1.6em;
    width: 1.6em;
}
.g-icon{
  width: 6em;
  height: 6em;
}

// 上传文件后不展示文件名和图标
:deep(.el-upload-list){
      display: none !important; 
}
.drawer-container {
  transition: all 0.2s;
  &:hover{
    right: 0
  }
  position: fixed;
  right: 0px;
  bottom: 15%;
  height: 50px;
  width: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 3001;
  color: #fff;
  background-color: #fff;
  border: 6px rgb(14, 14, 14) solid;
  border-radius: 50% !important;
  // border-radius: 4px 0 0 4px;
  cursor: pointer;
  -webkit-box-shadow: inset 0 0 6px rgba(0 ,0 ,0, 10%);
  box-shadow: inset 0 0 6px rgba(0 ,0 ,0, 10%);
}
.setting_body{
  padding: 20px;
  .setting_card{
    margin-bottom: 20px;
  }
  .setting_content{
    margin-top: 20px;
    display: flex;
    flex-direction: column;
    >.theme-box{
     display: flex;
    }
    >.color-box{
      div{
        display: flex;
        flex-direction: column;
      }
    }
    .item{
      display: flex;
      align-items: center;
      justify-content: center;
      flex-direction: column;
      margin-right: 20px;
      .item-top{
        position: relative;
      }
      .check{
        position: absolute;
        font-size: 20px;
        color: #00afff;
        right:10px;
        bottom: 10px;
      }
      p{
        text-align: center;
        font-size: 12px;
      }
    }
  }
}
// 输入框选择时阴影高度
:deep(.el-input__inner){
  height: 30px !important;
}

:deep(.el-empty__description){
  margin-top: 0px;
  margin-bottom: -10px;
}
:deep(.el-table__empty-text){
  line-height: 40px;
}
</style>
