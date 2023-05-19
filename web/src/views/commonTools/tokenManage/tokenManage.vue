<template>
    <div>
        <div class="token-wrap">
            <div class="dis-form-wrap">
                <div>
                    <span style="font-size: 18px;color: #7e8983;">{{$t('tm.title.dis')}}</span>
                </div>
                <div style="display: flex;flex-direction: row;align-items: center;margin-top: 10px;margin-bottom: 20px;width: 100%;" 
                v-loading="disLoading" element-loading-text="Sending..." 
                >
                    <div style="width:45%; display: flex;flex-direction: column;align-items: flex-start;width: 100%;margin-top: -52px;">
                        <el-form ref="disRef" :model="disForm" :rules="rules" size="medium" label-width="120px">
                            <el-row  style="">
                                <el-col :span="24">
                                    <el-form-item label-width="120px" :label="$t('tm.label.mm_transfer')" prop="useMm">
                                    <el-switch v-model="disForm.useMm"></el-switch>
                                    </el-form-item>
                                </el-col>
                            </el-row>
                        <el-row v-if="!disForm.useMm">
                            <el-col :span="24">
                                <el-form-item :label="$t('tm.label.main_wal')" label-width="120px" prop="mainWal">
                                    <el-select v-model="disForm.mainWal" :placeholder="$t('tm.ph.main_wal')" :style="{width: '100%'}" @change="changeMainWal" :disabled="disableSelector">
                                        <el-option v-for="(item, index) in mainWalOptions" :key="index" :label="item.label" :value="item.value">
                                        <span style="float: left">{{ item.label }}</span>
                                        <span style="float: right;color: var(--el-text-color-secondary);font-size: 13px;">
                                            {{ $t('com.balance')+':'+item.balance.slice(0,5)+'ETH' }}
                                        </span>
                                    </el-option>
                                    </el-select>
                                </el-form-item>
                            </el-col>
                        </el-row>
                        
                        <el-row v-if="!disForm.useMm" style="margin-top: 10px;">
                            <el-col :span="24">
                                <el-form-item :label="$t('tm.label.main_wal_pvk')" label-width="120px" prop="pvk">
                                    <el-input v-model="disForm.pvk"  :placeholder="$t('tm.ph.pvk')" clearable :style="{width: '100%'}" @change="changePvk">
                                    </el-input>
                                </el-form-item>
                            </el-col>
                        </el-row>
                        <el-row style="margin-top: 10px;">
                            <el-col :span="24">
                                <el-form-item :label="$t('tm.label.token_addr')" label-width="120px" prop="token">
                                    <el-input v-model="disForm.token" placeholder="Only eth so far" clearable :style="{width: '100%'}" disabled="true" >
                                    <!-- <el-input v-model="disForm.token" :placeholder="$t('tm.ph.token')" clearable :style="{width: '100%'}" > -->
                                    </el-input>
                                </el-form-item>
                            </el-col>
                        </el-row>
                        <el-row style="margin-top: 10px;">
                            <el-col :span="8">
                                <el-form-item :label="$t('tm.label.ava_bal')" label-width="120px" prop="priFee">
                                    <el-input v-model="disForm.avaliable"  v-loading="balLoading" :style="{width: '100%'}" disabled="true">
                                    </el-input>
                                </el-form-item>
                            </el-col>
                            <el-col :span="8">
                                <el-form-item :label="$t('tm.label.pri_fee')" label-width="100px" prop="priFee">
                                    <el-input v-model="disForm.priFee" placeholder="Priority Fee" clearable :style="{width: '100%'}" oninput="value=value.replace(/[^0-9.]/g,'')">
                                    </el-input>
                                </el-form-item>
                            </el-col>
                            <el-col :span="8">
                                <el-form-item label="Max Fee" label-width="80px" prop="maxFee">
                                    <el-input v-model="disForm.maxFee" placeholder="Max Fee" clearable :style="{width: '100%'}" oninput="value=value.replace(/[^0-9.]/g,'')">
                                    </el-input>
                                </el-form-item>
                            </el-col>
                        </el-row>
                    </el-form>
                    </div>
                    <div style="width:10%; display: flex;flex-direction: column;align-items: flex-end;vertical-align: middle;margin-top: 20px;">
                        <svg-icon icon-class="arrow-double-right" className='icon' />
                    </div>
                    <div style="width:45%; display: flex;flex-direction: column;align-items: flex-end;width: 100%;margin-left: 20px;margin-top: 10px;">
                        <el-row style="height: 80%;">
                            <div style="width: 100%;">
                                <el-button icon="Plus" type="primary"  @click="openDialogUp" :style="{width:'80px',height:'25px','margin-bottom':'5px'}"><span style="font-size: 10px;">Add Wallet</span></el-button>
                                <el-table :data="rightTbData"  :border="true" 
                                    :header-cell-style="{ 'text-align': 'center','heifht':'30px' }"
                                    :cell-style="{'padding': '0','text-align': 'center' }"
                                    v-loading="tableAddrLoading" :element-loading-text="$t('hashmint.ld.table')" 
                                    height="220"
                                >
                                    <template #empty>
                                        <!-- <el-empty :image-size="50" :description="t('hashmint.description')"></el-empty> -->
                                        <div style="    padding-left: 30%;padding-top: 25%;padding-bottom: 25%;">
                                            <el-upload
                                                class="upload-demo"
                                                ref="upload"
                                                :on-change="hanldFileChange"
                                                :on-success="handleAddrSuccess"
                                                :on-exceed="handleExceed"
                                                :auto-upload="false"
                                                :multiple="true"
                                                :limit="1"
                                                list-type="txt" style="display: flex; ">
                                                <el-button slot="trigger" size="small" type="info" style="width: 120px;margin-left: 5px;margin-top: -2px;"> {{$t('tm.btn.choose_file')}}</el-button>
                                                <el-tooltip placement="top" effect="light" style="float: right;">
                                                    <template #content>
                                                        <img src="@/assets/img/format.png" alt="">
                                                    </template>
                                                    <el-icon><QuestionFilled /></el-icon>
                                                </el-tooltip>
                                            </el-upload>
                                        </div>

                                    </template>
                                    <el-table-column prop="address" :label="$t('hashmint.label.address')" style="width: 60%;">
                                        <template #default="scope">
                                            <div>
                                            {{`${scope.row.address.slice(0,4)}...${scope.row.address.slice(-4)}`}}
                                            </div>
                                        </template>
                                    </el-table-column>
                                    <el-table-column prop="balance" :label="$t('hashmint.label.balance')" style="width: 10%;">
                                        <template #default="scope">
                                            <div v-loading="scope.row.loading" >
                                                <!-- {{scope.row.balance ? scope.row.balance.slice(0,7)+'ETH' :null}} -->
                                                {{scope.row.balance}}
                                            </div>
                                        </template>
                                    </el-table-column>
                                    <el-table-column prop="value" :label="$t('tm.label.transfer')">
                                        <template #default="scope">
                                            <el-input v-model="scope.row.value" clearable style="width: 80px;height: 35px;padding: 5px 5px;" :disabled="disableUpInput"> </el-input>
                                        </template>
                                    </el-table-column>
                                    <!-- <el-table-column prop="result" :label="$t('hashmint.label.res')"></el-table-column> -->

                                </el-table>
                            </div>
                        </el-row>
                        <el-row style="" >
                            <el-col :span="16" style="width:70%;display: flex;justify-items: start;">
                                <el-form-item :label="$t('tm.label.each_wal')" label-width="160px" prop="eachWal">
                                    <el-input v-model="disForm.eachWal" placeholder="" @change="eachChange" clearable :style="{width: '60%'}" oninput="value=value.replace(/[^0-9.]/g,'')">
                                    </el-input>
                                </el-form-item>
                            </el-col>
                            <el-col :span="8" style="width:30%;display: flex;justify-items: end;">
                                <el-form-item size="small">
                                    <el-button type="primary" @click="submitForm" :disabled="!subAva">{{$t('com.btn.confirm')}}</el-button>
                                    <el-button @click="resetForm" :disabled="!(dataList.length>0)">{{$t('com.btn.reset')}}</el-button>
                                </el-form-item>
                            </el-col>
                        </el-row>
                    </div>
                </div>
            </div>
        
            <div class="col-form-wrap">
                <div>
                    <span style="font-size: 18px;color: #7e8983;">{{$t('tm.title.coll')}}</span>
                </div>
                <div style="display: flex;flex-direction: row;align-items: center;margin-top: 10px;margin-bottom: 20px;width: 100%;" 
                v-loading="colLoading" element-loading-text="Sending...">
                    <div style="width:45%; display: flex;flex-direction: column;align-items: flex-start;width: 100%;margin-top: -52px;">
                        <el-form ref="pvkRef" :model="colForm" :rules="rulesPvk" size="medium" label-width="120px">
                        <el-row>
                            <el-col :span="24">
                                <el-form-item :label="$t('tm.label.main_wal_c')" label-width="120px" prop="mainWal">
                                    <el-select v-model="colForm.mainWal" :placeholder="$t('tm.ph.main_wal_c')" :style="{width: '100%'}" 
                                    @change="changePvkMainWal" 
                                    :disabled="disableAddrSelector">
                                        <el-option v-for="(item, index) in walColOption" :key="index" :label="item.label" :value="item.value">
                                        <span style="float: left">{{ item.label }}</span>
                                        <span style="float: right;color: var(--el-text-color-secondary);font-size: 13px;">
                                            {{ $t('com.balance')+':'+item.balance.slice(0,5)+'ETH' }}
                                        </span>
                                    </el-option>
                                    </el-select>
                                </el-form-item>
                            </el-col>
                        </el-row>
                        
                        <el-row  style="margin-top: 10px;">
                            <el-col :span="24">
                                <el-form-item :label="$t('tm.label.main_wal_pvk')" label-width="120px" prop="pvk">
                                    <el-input v-model="colForm.pvk"  :placeholder="$t('tm.ph.pvk_c')" clearable :style="{width: '100%'}" @change="changeMainAddress">
                                    </el-input>
                                </el-form-item>
                            </el-col>
                        </el-row>
                        <el-row style="margin-top: 10px;">
                            <el-col :span="24">
                                <el-form-item :label="$t('tm.label.token_addr')" label-width="120px" prop="token">
                                    <el-input v-model="colForm.token" placeholder="Only eth so far" clearable :style="{width: '100%'}" disabled="true" >
                                    </el-input>
                                </el-form-item>
                            </el-col>
                        </el-row>
                        <el-row style="margin-top: 10px;">
                            <el-col :span="12">
                                <el-form-item :label="$t('tm.label.pri_fee')" label-width="120px" prop="priFee">
                                    <el-input v-model="colForm.priFee" placeholder="Priority Fee" clearable :style="{width: '100%'}" oninput="value=value.replace(/[^0-9.]/g,'')">
                                    </el-input>
                                </el-form-item>
                            </el-col>
                            <el-col :span="12">
                                <el-form-item label="Max Fee" label-width="80px" prop="maxFee">
                                    <el-input v-model="colForm.maxFee" placeholder="Max Fee" clearable :style="{width: '100%'}" oninput="value=value.replace(/[^0-9.]/g,'')">
                                    </el-input>
                                </el-form-item>
                            </el-col>
                        </el-row>
                    </el-form>
                    </div>
                    <div style="width:10%; display: flex;flex-direction: column;align-items: flex-end;vertical-align: middle;margin-top: 20px;">
                        <svg-icon icon-class="arrow-double-left" className='icon' />
                    </div>
                    <div style="width:45%; display: flex;flex-direction: column;align-items: flex-end;width: 100%;margin-left: 20px;margin-top: 10px;">
                        <el-row style="height: 80%;">
                            <div style="width: 100%;">
                                <el-button icon="Plus" type="primary"  @click="openDialogDown" :style="{width:'80px',height:'25px','margin-bottom':'5px'}"><span style="font-size: 10px;">Add Wallet</span></el-button>
                                <el-table :data="pvkTable"  :border="true" 
                                    :header-cell-style="{ 'text-align': 'center','heifht':'30px' }"
                                    :cell-style="{'padding': '0','text-align': 'center' }"
                                    v-loading="tablePvkLoading" :element-loading-text="$t('hashmint.ld.table')" 
                                    height="220"
                                >
                                    <template #empty>
                                        <!-- <el-empty :image-size="50" :description="t('hashmint.description')"></el-empty> -->
                                        <div style="    padding-left: 30%;padding-top: 25%;padding-bottom: 25%;">
                                            <el-upload
                                                class="upload-demo"
                                                ref="uploadPvk"
                                                :on-change="hanldPvkFileChange"
                                                :on-success="handlePvkSuccess"
                                                :on-exceed="handlePvkExceed"
                                                :auto-upload="false"
                                                :multiple="true"
                                                :limit="1"
                                                list-type="txt" style="display: flex; ">
                                                <el-button slot="trigger" size="small" type="info" style="width: 120px;margin-left: 5px;margin-top: -2px;"> {{$t('tm.btn.choose_file_c')}}</el-button>
                                                <el-tooltip placement="top" :content="$t('tm.tip.file')" effect="light" style="float: right;">
                                                    <el-icon><QuestionFilled /></el-icon>
                                                </el-tooltip>
                                            </el-upload>
                                        </div>

                                    </template>
                                    <el-table-column prop="address" :label="$t('hashmint.label.address')" style="width: 60%;">
                                        <template #default="scope">
                                            <div>
                                            {{`${scope.row.address.slice(0,4)}...${scope.row.address.slice(-4)}`}}
                                            </div>
                                        </template>
                                    </el-table-column>
                                    <el-table-column prop="balance" :label="$t('hashmint.label.balance')" style="width: 10%;">
                                        <template #default="scope">
                                            <div v-loading="scope.row.loading" >
                                                <!-- {{scope.row.balance ? scope.row.balance.slice(0,7)+'ETH' :null}} -->
                                                {{scope.row.balance}}
                                            </div>
                                        </template>
                                    </el-table-column>
                                    <el-table-column prop="value" :label="$t('tm.label.transferout')">
                                        <template #default="scope">
                                            <el-input v-model="scope.row.value" clearable style="width: 80px;height: 35px;padding: 5px 5px;" :disabled="valueDownDis"> </el-input>
                                        </template>
                                    </el-table-column>
                                    <!-- <el-table-column prop="result" :label="$t('hashmint.label.res')"></el-table-column> -->

                                </el-table>
                            </div>
                        </el-row>
                        <el-row style="width: 100%;" >
                            <el-col :span="16" style="width:70%;display: flex;justify-items: start;">
                                <el-form-item label-width="120px" :label="$t('tm.label.etra_all')" prop="extrAll">
                                    <el-switch v-model="colForm.extrAll" @change="changeExtr"></el-switch>
                                </el-form-item>
                            </el-col>
                            <el-col :span="8" style="width: 30%;display: flex;justify-items: end;">
                                <el-form-item size="small">
                                    <el-button type="primary" @click="submitColForm" :disabled="!subPvkAva">{{$t('com.btn.confirm')}}</el-button>
                                    <el-button @click="resetPvkForm" >{{$t('com.btn.reset')}}</el-button>
                                </el-form-item>
                            </el-col>
                        </el-row>
                    </div>
                </div>

            </div>
        </div>
    <el-dialog v-model="addWalletVisible" title="添加钱包" draggable width="40%">
      <warning-bar title="请按格式录入" />
      <el-row >
            <el-col :span="24">
                <textarea :rows=5 v-model="addWalletString" autosize="{minRows:2,maxRows:8}" class="pv-text-area"
                    style ="margin-left: 2px;
                    font-size: 18px;
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
                    :placeholder="addWalletPlaceholder"
                ></textarea>
            </el-col>
        </el-row>

      <template #footer>
        <div class="dialog-footer" style="text-align: center;">
            <el-button @click="cancelAddWallet" >Cancel</el-button>
            <el-button type="primary" @click="confirmAddWallet">
            Confirm
            </el-button>
        </div>
      </template>
    </el-dialog>   
    </div>
    
  </template>
<script lang="ts">
    export default{
        name:"tokenManage"
    }
</script>
  <script lang="ts" setup>
import { ref ,onMounted,computed,h} from "vue";
import { useI18n } from 'vue-i18n'
import { genFileId ,ElMessage,ElNotification} from 'element-plus'
import {  ethers } from "ethers" 
import { URL_PREFIX } from "@/global/constants";
import { fileUtils } from '@/utils/file'
import { ethUtils } from '@/utils/blockchain'
import { useUserStore } from '@/pinia/modules/userStore'
import type { UploadInstance, UploadProps, UploadRawFile } from 'element-plus'
import { emitter } from '@/common/bus'
import { useDisperse} from '@/hooks/useDisperse'
const addWalletPlaceholder = ref()
// const walletList = ref([])
const upload = ref<UploadInstance>()
const uploadPvk = ref<UploadInstance>()
let addressFile =null
let pvkFile =null
let balLoading =ref(false)
let addressArr =[]
let valueArr =[]
let colValueArr =[]
let colAddrArr =[]
const isExtrAllArr =[]
const disRef = ref()
const pvkRef = ref()
const dataList = ref([])
const addWalletVisible = ref(false)
const disLoading = ref(false)
const colLoading = ref(false)
const openType = ref(0)
const addWalletString=ref('') 
const disableSelector =ref(false)
const disableAddrSelector =ref(false)
const rightTbData = ref([])
const pvkTable = ref([])
const tableAddrLoading = ref(false)
const tablePvkLoading = ref(false)
const { t } = useI18n()
const disForm=ref( {
    useMm: false,
    mainWal: '',
    pvk: '',
    token: '',
    avaliable: '',
    priFee: '',
    maxFee: '',
    eachWal: '',
})
const colForm=ref( {
    mainWal: '',
    pvk: '',
    token: '',
    priFee: '',
    maxFee: '',
    extrAll: true,
})
const disableUpInput =ref(false)
const eachChange=(val)=>{
    
    if(val){
        disableUpInput.value = true
    }else{
        disableUpInput.value =false
    }
}
const valueDownDis =ref(false)
const changeExtr=(val)=>{
    
    if(val){
        valueDownDis.value = true
    }else{
        valueDownDis.value =false
    }
}

// const disableUpInput = computed(()=>(disForm.value.eachWal=='' || !(disForm.value.eachWal)))
const mainWalV = ((rule,value,callback)=>{
    if(disForm.value.useMm)return true
    if(disForm.value.pvk && ethUtils.isValidPVKey(disForm.value.pvk))return true
    if(value)return true
    callback(new Error(t('com.msg.wrong_addr'),))

    return false
})

const pvkV = ((rule,value,callback)=>{

    if(disForm.value.useMm)return true
    if(disForm.value.mainWal)return true
    if(ethUtils.isValidPVKey(value))return true
    callback(new Error(t('com.msg.wrong_pvk'),))

    return false
})

const tokenV = ((rule,value,callback)=>{
    if(!value)return true
    if(value && ethUtils.isValidAddress(value)){
        return true
    }else{
        callback(new Error(t('com.msg.wrong_addr'),))
        return false
    }
})
const rules= {
    mainWal : [{required: false,message: t('hashmint.msg.hash'),trigger: 'blur'},{validator:mainWalV,trigger:'change'}],
    pvk     : [{required: false,message: t('hashmint.msg.hash'),trigger: 'blur'},{validator:pvkV,trigger:'change'}],
    token   : [{required: false,message: t('hashmint.msg.hash'),trigger: 'blur'},{validator:tokenV,trigger:'change'}],
}

const tokenVP = ((rule,value,callback)=>{
    if(!value)return true
    if(value && ethUtils.isValidAddress(value)){
        return true
    }else{
        callback(new Error(t('com.msg.wrong_addr'),))
        return false
    }
})
const mainWalVP = ((rule,value,callback)=>{
    if(colForm.value.pvk && ethUtils.isValidAddress(colForm.value.pvk))return true
    if(value)return true
    callback(new Error(t('com.msg.wrong_addr'),))

    return false
})
const addrV = ((rule,value,callback)=>{
    if(colForm.value.mainWal)return true
    if(ethUtils.isValidAddress(value))return true
    callback(new Error(t('com.msg.wrong_addr'),))

    return false
})

const rulesPvk= {
    mainWal : [{required: false,message: t('hashmint.msg.hash'),trigger: 'blur'},{validator:mainWalVP,trigger:'change'}],
    pvk     : [{required: false,message: t('hashmint.msg.hash'),trigger: 'blur'},{validator:addrV,trigger:'change'}],
    token   : [{required: false,message: t('hashmint.msg.hash'),trigger: 'blur'},{validator:tokenVP,trigger:'change'}],
}
const mainWalOptions =ref([])
const walColOption =ref([])
//提交分发
const submitForm=()=> {
    new Promise(async (resolve, reject) => {
        if(!(await disRef.value.validate())){
        ElMessage({
            type:'error',
            message:'请正确录入信息！'
        })
        resolve(false)
    }
    if(!checksub())resolve(false)
    else{
        disLoading.value= true
        let signer = null
        if(disForm.value.useMm){
            const provider = new ethers.providers.Web3Provider((window as any).ethereum);
            signer = provider.getSigner();
        }else{
            let mainWal =disForm.value.mainWal
            let pvk = disForm.value.pvk
            const {userProvider} = useUserStore() 
            if(mainWal){
                const {userWallets} = useUserStore()
                signer = userWallets.getSigner(mainWal)
            }else{
                signer =  new ethers.Wallet(pvk,userProvider.httpProvider)
            }
        }
        resolve(signer)
    }
    }).then((signer)=>{
        if(signer){
            const {handleEthDisperse} = useDisperse()
            handleEthDisperse(addressArr as any,valueArr as any,signer).then((res)=>{
                disLoading.value= false
                    if(res.status){
                        
                        ElNotification({
                            title: 'Success',
                            message: h('i', { style: 'color: teal' }, res.message),
                            onClick(){
                                window.open(`${URL_PREFIX.ES_TX}${res.tx.hash}`)
                            }
                        })
                    }else{
                        ElNotification({
                            title: 'Error',
                            type:'error',
                            message: h('i', { style: 'color: teal' }, res.message),
                        })
                        rightTbData.value.forEach((item)=>{
                            item.value=0
                        })
                        disForm.value.eachWal=''
                    }
                })
        }
    })
}
//提交归集
const submitColForm=()=> {
    new Promise(async (resolve, reject) => {
        if(!(await pvkRef.value.validate())){
            ElMessage({
                type:'error',
                message:'请正确录入信息！'
            })
            resolve(false)
        }
        if(!checkPvkSub()){
            resolve(false)
        }
        else{
            colLoading.value= true
        let address = colForm.value.mainWal
            resolve(address)
        }
    }).then(async (address)=>{
        if(address){
            const {handleEthCollect} = useDisperse()
            handleEthCollect((colAddrArr as any),(colValueArr as any),(isExtrAllArr as any),(address as string)).then((res)=>{
                if((res as any).status){
                    ElNotification({
                        title: 'Success',
                        message: h('i', { style: 'color: teal' },"全部完成"),
                    })
                }else{
                    if((res as any).failed.length>0){
                        ElNotification({
                        type:'warning',
                        title: '以下地址失败',
                        message: (res as any).failed,
                    }) 
                    }
                }
            }).catch((err)=>{
                ElMessage({
                    type:'error',
                    message:err
                })
            }).finally(()=>{
                colLoading.value= false
            })
             
        }
    })
}
const subAva=computed(()=>{
    if(rightTbData.value.length<1)return false
    let avaliable = disForm.value.avaliable
    let total= 0
    if(disForm.value.eachWal){
        total = rightTbData.value.length * (typeof disForm.value.eachWal!= 'number' ?parseFloat(disForm.value.eachWal) :disForm.value.eachWal)
    }else{
        rightTbData.value.forEach((item)=>{
            total += item.value
        })
    }
    if(total<=0)return false
    if((typeof avaliable =='string' ? parseFloat(avaliable) :avaliable) <=total)return false
    return true
})
const subPvkAva=computed(()=>{
    if(pvkTable.value.length<1)return false
    let total= 0
    if(!colForm.value.extrAll){
        pvkTable.value.forEach((item)=>{
            total += item.value
        })
        if(total<=0)return false
    }
    return true
})
const checksub =()=>{
    if(!(disRef.value.validate()))return false
    if(rightTbData.value.length<1)return false
    
    let avaliable = disForm.value.avaliable
    let total= 0
    addressArr=[]
    valueArr=[]
    if(disForm.value.eachWal){
        total = rightTbData.value.length * (typeof disForm.value.eachWal== 'string' ? parseFloat(disForm.value.eachWal) :disForm.value.eachWal)
        rightTbData.value.forEach((item)=>{
            valueArr.push(disForm.value.eachWal)
            addressArr.push(item.address)
        })
    }else{
        rightTbData.value.forEach((item)=>{
            if(item.value){
                total += (typeof item.value =='string' ? parseFloat(item.value) : item.value)
                valueArr.push(item.value)
                addressArr.push(item.address)
            }
        })
    }
    if(total<=0)return false
    if((typeof avaliable =='string' ? parseFloat(avaliable) :avaliable) <=total){
        ElMessage({
            type:'error',
            message:'Balance Not Enough!'
        })
        return false
    }
    return true
}
const checkPvkSub =()=>{
    if(pvkTable.value.length<1)return false
    let total= 0
    addressArr=[]
    valueArr=[]
    if(colForm.value.extrAll){
        pvkTable.value.forEach((item)=>{
            let bal =0;
            try {
                bal =parseFloat(item.balance.slice(0,5))
            } catch (error) {
                bal=0
            }
            if(bal>0.005){
                total += bal
                colValueArr.push(bal)
                colAddrArr.push(item.address)
                isExtrAllArr.push(true)//标记全部提取，假如是提取则金额预留gas费
            }else{
                ElMessage.warning("余额小于0.005ETH的账户将会忽略")
            }
        })
    }else{
        pvkTable.value.forEach((item)=>{
            if(item.value){
                let v =0
                try {
                    v =(typeof item.value =='string' ? parseFloat(item.value) : item.value)
                } catch (error) {
                    v = 0
                }
                if(v >=0.005){
                    total += v
                    colValueArr.push(item.value)
                    colAddrArr.push(item.address)
                    if((parseFloat(item.balance.slice(0,5)) - v) <= 0.0001){
                        isExtrAllArr.push(true)//标记全部提取，假如是提取则金额预留gas费
                    }else{
                        isExtrAllArr.push(false)
                    }
                }else{
                    ElMessage.error("归集金额必须大于或等于0.005ETH")
                }
            }
        })
    }
    if(total<=0)return false
    return true
}

const parseAddress =async()=>{
    if (addressFile!=null){
        try {
            let s = await fileUtils.readFile(addressFile) as string
            let ss = s.trim()
            let adds = ss.replace(/\r\n/g, "*").replace(/\n/g, "*")
            let addArr =[]
            let temArr = adds.split('*')
            tableAddrLoading.value=true
            temArr.forEach((i)=>{
                let ar = i.split(',')
                if(!ar || !ethUtils.isValidAddress(ar[0])){
                    upload.value.clearFiles()
                    throw new Error("wrong file!");
                }
                addArr.push(ar[0])
                let v = dataList.value.findIndex((item)=>item.address==ar[0])
                if(v<0){
                    dataList.value.push({
                        address:ar[0],
                        value:ar[1] ? ar[1] : 0,
                        balance:0,
                        loading:true,
                    })
                }
            })
            tableAddrLoading.value=false
            Promise.allSettled(addArr.map(ethUtils.getBalance)).then((res)=>{
                res.forEach((i,v)=>{
                    if(i.status==='fulfilled'){
                        let wal = (i as any).value
                        let vv = dataList.value.findIndex((item)=>item.address==wal.address)
                        if(!(vv<0)){
                            dataList.value[vv].balance = wal.balance.slice(0,7)+'ETH'   
                            dataList.value[vv].loading = false 
                        }
                    }else{
                        dataList.value[v].balance = 'unknown'   
                    }
                })
                // tableAddrLoading.value=false
            })
            
        } catch (error) {
            upload.value.clearFiles()
            tableAddrLoading.value=false
            ElMessage({
                message:error,
                type:'error'
            })
        }finally{
            rightTbData.value = dataList.value
        }
        
    }
}
const parsePvk =async()=>{
    
    if (pvkFile!=null){
        try {
            let s = await fileUtils.readFile(pvkFile) as string
            let ss = s.trim()
            let pv2 = ss.replace(/\r\n/g, ",").replace(/\n/g, ",")
            let pvArr = pv2.split(',')
            let pvkArr =[]
            tablePvkLoading.value=true
            const {userWallets} = useUserStore()
            pvArr.forEach((i)=>{
                if(!ethUtils.isValidPVKey(i)){
                    uploadPvk.value.clearFiles()
                    throw new Error("wrong file!");
                }
                pvkArr.push(i)
                userWallets.addWalletByPvKey(i,"TokenManage")
                let addr = new ethers.Wallet(i).address
                let v = dataList.value.findIndex((item)=>item.address==addr)
                if(v<0){
                    dataList.value.push({
                        address:addr,
                        value:0,
                        balance:0,
                        loading:true,
                    })
                }
            })
            tablePvkLoading.value=false
            Promise.allSettled(pvkArr.map(ethUtils.getBalanceByPvk)).then((res)=>{
                res.forEach((i,v)=>{
                    if(i.status==='fulfilled'){
                        let wal = (i as any).value
                        let vv = dataList.value.findIndex((item)=>item.address==wal.address)
                        addPvkOption(wal)
                        if(!(vv<0)){
                            dataList.value[vv].balance = wal.balance.slice(0,7)+'ETH'   
                            dataList.value[vv].loading = false 
                        }
                    }else{
                        dataList.value[v].balance = 'unknown'   
                    }
                })
                // tableAddrLoading.value=false
            })
            
        } catch (error) {
            uploadPvk.value.clearFiles()
            tablePvkLoading.value=false
            ElMessage({
                message:error,
                type:'error'
            })
        }finally{
            pvkTable.value = dataList.value
        }
        
    }
}
const addPvkOption =(wal)=>{
    if(walColOption.value.findIndex((item)=>item.address==wal.address)<0){
        walColOption.value.push({
            label:wal.address,
            value:wal.address,
            balance:wal.balance.slice(0,7)+'ETH',
        })
    }
}
const changeMainWal =(value)=>{
    if(!value)return
    let index = dataList.value.findIndex((item)=>item.address==value)
    if(!(index<0)){
        disForm.value.avaliable= dataList.value[index].balance
    }
    // disForm.value.avaliable = 
    disForm.value.pvk=""
    rightTbData.value =[...dataList.value]
    rightTbData.value.forEach((item,v)=>{
        if(item.address == value){
            rightTbData.value.splice(v, 1)
        }
    })
}
const changePvkMainWal =(value)=>{
    if(!value)return
    let index = dataList.value.findIndex((item)=>item.address==value)
    if(!(index<0)){
        disForm.value.avaliable= dataList.value[index].balance
    }
    // disForm.value.avaliable = 
    colForm.value.pvk=""
    pvkTable.value =[...dataList.value]
    pvkTable.value.forEach((item,v)=>{
        if(item.address == value){
            pvkTable.value.splice(v, 1)
        }
    })
}

const openDialogUp =()=>{
    addWalletPlaceholder.value =`Example:
0x6d0068652E7bFBE823f506c8bA945A4f4Ae77974,1.2
0x76906e0aB3d95EDa86a19CB92B69542c4d61d761,0.5`
    if(openType.value!==1){
        addWalletString.value =''
    }
    openType.value = 1
    addWalletVisible.value = true
}
const openDialogDown =()=>{
    addWalletPlaceholder.value =`每行一个`
    if(openType.value!==2){
        addWalletString.value =''
    }
    openType.value = 2
    addWalletVisible.value = true
}
const cancelAddWallet=()=>{
    addWalletVisible.value = false
}
const confirmAddWallet=()=>{
    if(openType.value===1){
        let tempArray = addWalletString.value.split(/[\s,;:\t\r\n]+/);
        let oddArray = tempArray.filter((v,i) => i%2);//金额
        let evenArray = tempArray.filter((v,i) => !(i%2));//地址
    
        let addArray = [];
        let amtArray = [];
        let total = 0.0;
        const UNIT = 1000000000000000000;
        let dataArr=[]
        for (let i = 0; i < evenArray.length; ++i) {
            if(!ethUtils.isValidAddress(evenArray[i])){
                ElMessage({
                    type:'error',
                    message:'包含错误地址!'
                })
                return
            }
            if (!isNaN((oddArray[i] as any)) && (parseFloat(oddArray[i]) > 0.0 )) {
                addArray.push(evenArray[i].trim());
                amtArray.push(('0x'+((oddArray[i] as any)*UNIT).toString(16)));
                total += parseFloat(oddArray[i]);
                dataArr.push({
                    address:evenArray[i],
                    value:oddArray[i]
                })
            }else{
                ElMessage({
                    type:'error',
                    message:'包含错误金额!'
                })
                return
            }
        }
        addWalToTable(dataArr)      
        
    }else if(openType.value===2){
        let pv2 = addWalletString.value.replace(/\r\n/g, ",").replace(/\n/g, ",")
        let pvArr = pv2.split(',')
        let dataArr =[]
        const {userWallets} = useUserStore()
        for (let i = 0; i < pvArr.length; ++i) {
            if(!ethUtils.isValidPVKey(pvArr[i])){
                ElMessage({
                    type:'error',
                    message:'包含错误私钥!'
                })
                return
            }
            let wal ={
                address:new ethers.Wallet(pvArr[i]).address,
                balance:0,
            }
            userWallets.addWalletByPvKey(pvArr[i],"TokenManage")
            dataArr.push(wal)
        }
        addWalToPvkTable(dataArr)
    }
    addWalletVisible.value = false
}
const addWalToTable=(wals:any)=>{
    let tempArr =[]
    let ri =[]
    let di =[]
    wals.forEach((wal)=>{
        let v = dataList.value.findIndex((item)=>item.address==wal.address)
        if(v<0){
            let w ={
                balance:0,
                loading:true,
                address:wal.address,
                value:wal.value
            }
            ri.push(rightTbData.value.length)
            di.push(dataList.value.length)
            rightTbData.value.push(w)
            dataList.value.push(w)
            tempArr.push(w.address)
        }
    })
    Promise.allSettled(tempArr.map(ethUtils.getBalance)).then((statuses)=>{
            statuses.forEach((i,v)=>{
                if(i.status==='fulfilled'){
                    rightTbData.value[ri[v]]['balance']=(i as any).value.balance.slice(0,7)+'ETH'
                    rightTbData.value[ri[v]]['loading']=false
                    dataList.value[di[v]]['balance']=(i as any).value.balance.slice(0,7)+'ETH'
                    dataList.value[di[v]]['loading']=false
                }else{//多是网络错误
                    rightTbData.value[ri[v]]['balance']='unknow'
                    rightTbData.value[ri[v]]['loading']=false
                    dataList.value[di[v]]['balance']='unknow'
                    dataList.value[di[v]]['loading']=false
                }
            })
        })
}
const addWalToPvkTable=(wals:any)=>{
    let tempArr =[]
    let ri =[]
    let di =[]
    wals.forEach((wal)=>{
        let v = dataList.value.findIndex((item)=>item.address==wal.address)
        if(v<0){
            let w ={
                balance:0,
                loading:true,
                address:wal.address,
                value:wal.value
            }
            ri.push(pvkTable.value.length)
            di.push(dataList.value.length)
            pvkTable.value.push(w)
            dataList.value.push(w)
            tempArr.push(w.address)
        }
    })
    Promise.allSettled(tempArr.map(ethUtils.getBalance)).then((statuses)=>{
            statuses.forEach((i,v)=>{
                if(i.status==='fulfilled'){
                    pvkTable.value[ri[v]]['balance']=(i as any).value.balance.slice(0,7)+'ETH'
                    pvkTable.value[ri[v]]['loading']=false
                    dataList.value[di[v]]['balance']=(i as any).value.balance.slice(0,7)+'ETH'
                    dataList.value[di[v]]['loading']=false
                    addPvkOption((i as any).value)
                }else{//多是网络错误
                    pvkTable.value[ri[v]]['balance']='unknow'
                    pvkTable.value[ri[v]]['loading']=false
                    dataList.value[di[v]]['balance']='unknow'
                    dataList.value[di[v]]['loading']=false
                }
            })
        })
}

const resetForm=() =>{
    // dataList.value=[]
    rightTbData.value=[]
    disRef.value.resetFields()
}
const resetPvkForm=() =>{
    // dataList.value=[]
    pvkTable.value=[]
    pvkRef.value.resetFields()
    
}

const changePvk=(val)=>{
    if(ethUtils.isValidPVKey(val)){
        disForm.value.mainWal=""
        disableSelector.value = true
        balLoading.value=true
        ethUtils.getBalanceByPvk(val).then((res)=>{
            balLoading.value=false
            disForm.value.avaliable= (res as any).balance
        })
    }else{
        disableSelector.value = false
        balLoading.value=false
    }
    
}
const changeMainAddress=(val)=>{
    if(ethUtils.isValidAddress(val)){
        colForm.value.mainWal=""
        disableAddrSelector.value = true
    }else{
        disableAddrSelector.value = false
    }

}
const handleExceed: UploadProps['onExceed'] = (files) => {
  upload.value!.clearFiles()
  const file = files[0] as UploadRawFile
  file.uid = genFileId()
  upload.value!.handleStart(file)
}
const hanldFileChange =(file,list)=>{
    const isH5 = file.name.split(".")[1] === "txt";
    const isLt2G = file.size / 1024 / 1024 < 5;
    if (!isH5) {
        ElMessage({
          message: "仅支持上传.txt文件",
          type: 'error'
        })
        upload.value!.clearFiles()
        return
	}
    if (!isLt2G) {
        ElMessage({
          message: "上传文件大小不能超过 5KB!",
          type: 'error'
        })
        upload.value!.clearFiles()
        return
    }
    if (list.length > 1 && file.status !== "fail") {
      list.splice(0, 1);
    } else if (file.status === "fail") {
        ElMessage({
          message: "失败，请重新上传",
          type: 'error'
        })
      list.splice(0, 1);
      return
    }
    addressFile = file
    parseAddress()
}
const handleAddrSuccess =(file)=>{
    addressFile = file
    parseAddress()
}
const handlePvkExceed: UploadProps['onExceed'] = (files) => {
    uploadPvk.value!.clearFiles()
  const file = files[0] as UploadRawFile
  file.uid = genFileId()
  uploadPvk.value!.handleStart(file)
}
const hanldPvkFileChange =(file,list)=>{
    const isH5 = file.name.split(".")[1] === "txt";
    const isLt2G = file.size / 1024 / 1024 < 5;
    if (!isH5) {
        ElMessage({
          message: "仅支持上传.txt文件",
          type: 'error'
        })
        uploadPvk.value!.clearFiles()
        return
	}
    if (!isLt2G) {
        ElMessage({
          message: "上传文件大小不能超过 5KB!",
          type: 'error'
        })
        uploadPvk.value!.clearFiles()
        return
    }
    if (list.length > 1 && file.status !== "fail") {
      list.splice(0, 1);
    } else if (file.status === "fail") {
        ElMessage({
          message: "失败，请重新上传",
          type: 'error'
        })
      list.splice(0, 1);
      return
    }
    pvkFile = file
    parsePvk()
}
const handlePvkSuccess =(file)=>{
    pvkFile = file
    parsePvk()
}

emitter.on('addWallet',((wal:any)=>{
   wal['value'] = 0;
   let v = dataList.value.findIndex((item)=>item.address==wal.address)
   if(v<0){
        mainWalOptions.value.push({
            label:wal.address,
            value:wal.address,
            balance:wal.balance.slice(0,7)+'ETH',
        })
       addPvkOption(wal)
       dataList.value.push(wal)
       rightTbData.value=[...dataList.value]
       pvkTable.value=[...dataList.value]
   }
}))
//打开页面时，加载全局私钥钱包
onMounted(()=>{
    const {userWallets} =useUserStore()
    const {userProvider}=useUserStore()
    const wals = userWallets.getAllWallets()
    wals.forEach(async w => {
        let wal ={
            address:w.address,
            balance:w.balance,
            value:0,
        }
        mainWalOptions.value.push({
            label:wal.address,
            value:wal.address,
            balance:wal.balance,
        }) 
        addPvkOption(wal)
        dataList.value.push(wal)
        rightTbData.value.push(wal)
        pvkTable.value.push(wal)

    });

})


  </script>
<style lang="scss" scoped>
  .token-wrap{
    height: auto;
    min-height: 880px;
    box-shadow: var(--el-box-shadow-light);
    background-color: var(--bg-color);
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .dis-form-wrap{
    // border:#909090 1px solid;
    padding-top: 15px;
    padding-left: 20px;
    padding-right: 40px;
    margin-top: 40px;
    width: 70%;
    height: 370px;
    border:  2px solid #7e8983;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    align-items: center;
    
    // box-shadow: var(--el-box-shadow-dark);
    .el_input_inner{
        width: 100%;
    }
    .el-form-item__label {
      color: #000000;
      font-size: 15px;
    }
}
.col-form-wrap{
    // border:#909090 1px solid;
    padding-top: 15px;
    padding-left: 40px;
    padding-right: 40px;
    margin-top: 40px;
    margin-bottom: 40px;
    width: 70%;
    height: 370px;
    border:  2px solid #7e8983;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    align-items: center;
    // box-shadow: var(--el-box-shadow-dark);
    .el_input_inner{
        width: 100%;
    }
    .el-form-item__label {
      color: #000000;
      font-size: 15px;
    }
}

:deep(.el-table__header-wrapper) {
    flex-shrink: 0;
    padding: 0px;
    width: 100% !important;
    // width: calc(100% - 10px) !important;
    height: 30px !important;
}
:deep(.el-table__header) {
    width: calc(100% - 40px) !important;
}
:deep(.el-table th .cell ){
    /* color: rgba(0, 0, 0, 0.85); */
    font-size: 14px;
    padding: 0px;
    /* line-height: 40px; */
    /* min-height: 40px; */
}

.icon{
    margin-top: -40px;
    width: 4em;
    height: 4em;
}
  </style>
  