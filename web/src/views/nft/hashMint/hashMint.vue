<template>
        <div style="height: auto;">
         <message-bar :title="$t('hashmint.message')" v-show="props.showWarnBar" />
         <div class="hash-mint-warp">
             <div class="hash-mint-form-wrap">
                  <el-form ref="formRef" :inline="true" :model="form" :rules="rules" label-width="130px" v-loading="loading" element-loading-text="Loading...">
                     <el-row style="width:100%">
                         <el-col :span="16" >
                             <el-form-item  style="width: 100%;" :label="$t('hashmint.label.hash')" prop="txHash">
                                 <el-tooltip :content="$t('hashmint.tip.hash')" placement="top" effect="light" style="float: right;">
                                     <el-icon><QuestionFilled /></el-icon>
                                 </el-tooltip>
                                 <span style="width: 95%;">
                                     <el-input class="gd-input" v-model="form.txHash" clearable  autocomplete="off" :placeholder="$t('hashmint.ph.input_hash')" @input="changeTxHash" />
                                 </span>
                             </el-form-item>
                         </el-col>
                         <el-col :span="4" style="display: flex;flex-direction: row;justify-content: flex-end;">
                             <el-button size="medium" type="danger" @click="clearForm()" style="width: 150px;">{{$t('hashmint.btn.reset')}}</el-button>
                         </el-col>
                     </el-row>
                    <el-row style="padding-top: 20px;">
                        <el-col :span="20">
                            <el-form-item :label="$t('hashmint.label.contract_addr')" prop="contractAddress" style="width: 100%;">
                                <el-input class="gd-input" v-model="form.contractAddress" autocomplete="off" :placeholder="$t('hashmint.ph.input_address')" @click="copyAddress"  readonly  disabled="true"/>
                            </el-form-item>
                        </el-col>
                    </el-row>
                     <el-row style="padding-top: 20px;">
                         <el-col :xl="props.fromOut ? 10 :4" :lg="10" :md="10" :sm="10" :xs="20" >
                             <el-form-item label="Value:" prop="value"  style="width: 100%;">
                                 <el-input class="gd-input" v-model="form.value" autocomplete="off" style="min-width: 60px;" placeholder="(ETH)" readonly  disabled="true"/>
                             </el-form-item>
                         </el-col> 
                         <el-col :xl="props.fromOut ? 10 :4" :lg="10" :md="10" :sm="10" :xs="20">
                             <el-form-item prop="maxFee" style="width: 100%;">
                                <template #label>
                                    <span style="width:80px;text-align: right;">Max fee:</span>
                                </template>
                                 <el-input class="gd-input" v-model="form.maxFee" autocomplete="off" style="min-width: 60px;" :placeholder="$t('hashmint.ph.input_maxfee')" oninput="value=value.replace(/[^0-9.]/g,'')"  />
                             </el-form-item>
                         </el-col>
                         <el-col :xl="props.fromOut ? 10 :4" :lg="10" :md="10" :sm="10" :xs="20">
                             <el-form-item label="Priority fee:" prop="priorityFee" style="width: 100%;">
                                 <el-input class="gd-input" v-model="form.priorityFee" autocomplete="off" style="min-width: 60px;" :placeholder="$t('hashmint.ph.input_prifee')" oninput="value=value.replace(/[^0-9.]/g,'')" />
                             </el-form-item>
                         </el-col>
                         <el-col :xl="props.fromOut ? 10 :4" :lg="10" :md="10" :sm="10" :xs="20">
                             <el-form-item prop="gasLimit" style="width: 100%;">
                                <template #label>
                                    <span style="width:80px;text-align: right;">Gas limit:</span>
                                </template>
                                 <el-input class="gd-input" v-model="form.gasLimit" autocomplete="off" style="min-width: 60px;" :placeholder="$t('hashmint.ph.input_gaslimit')" oninput="value=value.replace(/[^0-9.]/g,'')" />
                             </el-form-item>
                         </el-col>
                         <el-col :xl="props.fromOut ? 10 :4" :lg="10" :md="10" :sm="10" :xs="20">
                             <el-form-item :label="$t('hashmint.label.max_cost')" style="width: 100%;">
                                 <el-input class="gd-input" v-model="maxCost" autocomplete="off" style="min-width: 60px;"  :placeholder="$t('hashmint.ph.input_cost')" readonly  disabled="true"/>
                             </el-form-item>
                         </el-col>
                     </el-row>
                     <el-row style="padding-top: 20px;">
                         <el-col :span="20">
                             <el-form-item label="data:" prop="data"  style="width: 100%;">
                                 <textarea rows="4" v-model="form.data" autosize="{minRows:2,maxRows:8}" class="pv-text-area"
                                     style ="max-width: 100%;
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
                                     :placeholder="$t('hashmint.ph.txtarea_data')"
                                 ></textarea>
                             </el-form-item>
                         </el-col>
                     </el-row>
             </el-form>
             </div>
             
             <div class="start-wrap">
                
                 <el-row>
                     <el-col :span="21">
                         <el-form :model="mmForm" ref="mmFormRef" v-if="useMm" inline="true">
                             <el-row type="flex" justify="right">
                                 <!-- <el-col :span="8"  ></el-col> -->
                                 <el-col :span="9">
                                 </el-col>
                                 <el-col :span="15"  >
                                     <el-form-item label="Hash:" >
                                          <span style="color: #126b40;">{{mmForm.useMmHash}}
                                             <a aria-label="etherscan"
                                              :href="`${URL_PREFIX.ES_TX}${mmForm.useMmHash}`" role="link" rel="noopener noreferrer" target="_blank">
                                              <svg-icon icon-class="etherscan-logo-circle" className="etherscan"></svg-icon>
                                              </a>
                                         
                                         </span>
                                     </el-form-item>
                                     <el-form-item label="Result:">
                                         <span v-if="mmForm.status" style="color: #126b40;" v-loading="!mmForm.isEnd">{{mmForm.useMmRes}}
                                         </span>
                                         <span v-if="!mmForm.status" style="color: red;" v-loading="!mmForm.isEnd">{{mmForm.useMmRes}}</span>
                                     </el-form-item>
                                 </el-col>
                             </el-row>
                         </el-form>
                     </el-col>
                         <el-col :span="3">
                             <el-button type="primary" class="start-btn" style="margin-top: -20px;border-radius: 20px;height: 35px;"  @click="metamaskMmint">
                                     <template #icon>
                                         <svg-icon icon-class="metamask-icon" class ="metamask-icon" style="width: 4em; height: 4em;"/>
                                     </template>
                                     {{$t('hashmint.btn.mm_mint')}}
                                 </el-button>
                         </el-col>
                 </el-row>
                 <el-row>
                     <el-col :span="24">
                         <hr style="width: 75%;"/>
                     </el-col>
                 </el-row>
                 <el-row>
                   <el-col :span="21"></el-col>
                   <el-col :span="3">
                     <el-button type="primary" class="start-btn" style="margin-left: 0px;border-radius: 20px;height: 35px;" @click="batchMint"  :circle="true">
                             <template #icon>
                                 <font-awesome-icon icon="fa-solid fa-bars-progress"/>
                             </template>
                             {{$t('hashmint.btn.batch_mint')}}
                         </el-button>
                   </el-col>
                 </el-row>
             </div>
             <div class="result-show-wrap">
             <el-table :data="hmWalletList"  :border="true" 
                 :header-cell-style="{ 'text-align': 'center' }"
                 :cell-style="{ 'text-align': 'center' }" 
                 style="width: 74%;margin-left: 140px;text-align: center" 
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
                         <div>
                             {{`${scope.row.balance.slice(0,6)}ETH`}}
                         </div>
                     </template>
                 </el-table-column>
                 <el-table-column :label="$t('hashmint.label.hash')" >
                     <template #default="scope" >
                         <div v-if="scope.row.hash && scope.row.hash!=='showLoading' && scope.row.hash.slice(0,2)=='0x'" style="color: #126b40;" v-loading ="scope.row.hash==='showLoading'" :style="{'minWidth':'36px','minHeight':'36px'}">
                             {{scope.row.hash=='showLoading' ? '':scope.row.hash}}
                             <a v-if="scope.row.hash && scope.row.hash!=='showLoading' && scope.row.hash.slice(0,2)=='0x'" aria-label="etherscan" :href="`${URL_PREFIX.ES_TX}${scope.row.hash}`" role="link" rel="noopener noreferrer" target="_blank">
                                 <svg-icon icon-class="etherscan-logo-circle" className="etherscan"></svg-icon>
                             </a>
                         </div>  
                         <div v-else style="color: red;" v-loading ="scope.row.hash==='showLoading'" :style="{'minWidth':'36px','minHeight':'36px'}">
                             {{scope.row.hash=='showLoading' ? '':scope.row.hash}}
                         </div>
                     </template>
                 </el-table-column>
                 <el-table-column :label="$t('hashmint.label.status')">
                     <template  #default="scope">
                         <div v-if="scope.row.status=='Success'" v-loading="scope.row.status=='showLoading'" element-loading-text="waiting" style="color: #126b40;minWidth:20px;minHeight:20px;">
                             {{scope.row.status=='showLoading' ? '':scope.row.status}}
                         </div>
                          <div v-if="scope.row.status!='Success' || !scope.row.status" v-loading="scope.row.status=='showLoading'" element-loading-text="waiting" style="color: red;minWidth:40px;minHeight:40px;">
                             {{scope.row.status=='showLoading' ? '':scope.row.status}}
                          </div>
                     </template>
                 
                 </el-table-column>
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
                                 <el-button size="mini" type="danger" plain >{{$t('hashmint.btn.delete')}}</el-button>
                             </template>
                         </el-popconfirm>
                       </template>
                 </el-table-column>
     
             </el-table>
             </div>
         </div>
        </div>
</template>


<script lang="ts">
    export default{
       name:"hashMint"
    }
</script>
<script setup lang="ts">
import { ref,reactive,computed,onMounted,h} from 'vue'
import { BigNumber, errors, ethers } from "ethers" 
import { ElMessage ,ElNotification} from 'element-plus'
import { useUserStore } from '@/pinia/modules/userStore'
import MessageBar from '@/components/messageBar/messageBar.vue'
import { ethUtils } from '@/utils/blockchain'
import { URL_PREFIX } from "@/global/constants";
import { useSingleTxHandler } from '@/hooks/singleTxHandler'
import { emitter } from '@/common/bus'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const useMm = ref(false)
const startBtMint = ref(true)

const props=defineProps({
    txHash:{
      type: String,
      default: '',
    },
    showWarnBar: {
      type: Boolean,
      default: true
    }, 
    fromOut: {
      type: Boolean,
      default: false
    }
  })




const mmForm = ref({
    useMmRes:'',
    useMmHash:'',
    status:true,
    isEnd:false,
})
const mmFormRef = ref()

const loading = ref(false)
const tableLoading = ref(false)
var rawTx ={}
const formRef = ref()
const hmWalletList = ref([])
const userStore = useUserStore()
    const form = reactive({
        alreadyParseHash:false,
        gasLimit:'',
        txHash:'',
        value:'',
        gasPrice:'',
        baseFee:'',
        priorityFee:'',
        maxFee:'',
        contractAddress:'',
        data:''
    })

    const clearForm = ()=>{
        formRef.value.clearValidate()
        formRef.value.resetFields()
        maxCost ='' as any
        let list = hmWalletList.value
        for(let i in list){
            hmWalletList.value[i].status=''
        }
    }

    const copyAddress=()=> {
     var input = document.createElement("input"); // 创建input对象
     input.value = form.contractAddress; // 设置复制内容
     document.body.appendChild(input); // 添加临时实例
     input.select(); // 选择实例内容
     document.execCommand("Copy"); // 执行复制
     document.body.removeChild(input); // 删除临时实例
     ElMessage({
          message: t('hashmint.msg.copy'),
          type: 'success'
        })
}
    const changeTxHash = (async (value)=>{
        form.alreadyParseHash = false
        startBtMint.value=false
        useMm.value=false
        mmFormRef.value = {
            useMmRes:'',
            useMmHash:'',
            status:true,
            isEnd:false,
        }
        if(ethUtils.isValidTxHash(value)){
            loading.value=true
            try {
                let tx = await userStore.userProvider.httpProvider.getTransaction(value)
                rawTx = tx
                loading.value=false
               form.gasLimit = Number(tx.gasLimit)as any
               form.gasPrice = ethUtils.weiToGwei(tx.gasPrice) as any
               form.priorityFee = ethUtils.weiToGwei(Number(tx.maxPriorityFeePerGas)) as any
               form.baseFee = ethUtils.weiToGwei(tx.gasPrice) as any - ethUtils.weiToGwei(Number(tx.maxPriorityFeePerGas)) as any
               form.maxFee =ethUtils.weiToGwei(tx.maxFeePerGas) as any
               form.value = ethers.utils.formatEther(tx.value)
               form.data = tx.data
               form.contractAddress = tx.to
               form.alreadyParseHash = true
           } catch (error) {
                ElMessage({
                    type:'error',
                    message:'parse hash error!'
                })
           }
        }
        
    })
    var maxCost = computed(()=>{
        //公式： gasPrice = baseFee + PriorityFee , 这里的最大花费不以MaxFee去计算，而以gasPrice去计算
        //gasPrice * gasLimit = maxCost, 
        let p = form.priorityFee//gwei
        if(p==='' ||(typeof p ==='string'&& p.lastIndexOf('.')+1===form.priorityFee.length))return
        let b = form.baseFee //gwei
        let gasPrice = ethUtils.GweiToWei(Number.parseFloat(p)+Number.parseFloat(b))// wei
        let weiCost =  BigNumber.from(gasPrice).mul(BigNumber.from(form.gasLimit))
        let ethCost =  ethers.utils.formatEther(weiCost)
        return (Number.parseFloat(ethCost)+Number.parseFloat(form.value)).toString().slice(0,7)+'ETH'
    })
    const isVlidHash = ((rule,value,callback)=>{
        if(ethUtils.isValidTxHash(value)){
            return callback()
        }
        callback(new Error(t('hashmint.msg.hash'),))
    })
  
    const rules = reactive({
        txHash: [{required: true,message: t('hashmint.msg.hash'),trigger: 'blur'},{validator:isVlidHash,trigger:'change'}],
    })

emitter.on('addWallet',((wal:any)=>{
    let v = hmWalletList.value.findIndex((item)=>item.address==wal.address)
   if(v<0){
       hmWalletList.value.push(wal)
   }
}))
//右侧设置栏移除一个钱包时，hashmint也要把那个钱包移除，
emitter.on('removeWallet',((wal:any)=>{
    var list = hmWalletList.value;
      for (var i = 0; i < list.length; i++) {
          if (wal.address == list[i].address) {
            hmWalletList.value.splice(i, 1);
           }
       }  
}))
//监听点击mint事件

//移除hashmint操作的钱包时，不会在全局移除那个钱包
const removeWallet =(index, row)=> {
      var list = hmWalletList.value;
    //   let addr =''
      for (var i = 0; i < list.length; i++) {
          if (row.address == list[i].address) {
            // addr = row.address
            //移除hashmint操作的钱包时，不会在全局移除那个钱包
            // userStore.userWallets.removeFollowMintPvKey(addr)
            hmWalletList.value.splice(i, 1);
            // hashPvmap.value.delete()
              // this.$message({ message: '删除成功', duration: 2000, type: 'success' });
           }
       }  
}

const metamaskMmint =async()=>{
    
    if(!form.alreadyParseHash){
        ElMessage({
            type:"error",
            message:"未解析到交易Hash!"
        })
        return
    }
    const {singleTxMmSender} =useSingleTxHandler()
    singleTxMmSender(rawTx,form).then((res:any)=>{
        //{"status": true,"txHash": "0xe8dadcbd163a535a75170c18cc98c277ef55888922fe450f35c04132f83aafd1"}
        if(res.status){
            useMm.value=true
            mmForm.value.useMmHash = res.txHash
            userStore.userProvider.httpProvider.waitForTransaction(res.txHash,null,20000).then((data)=>{
                mmForm.value.isEnd = true
                if(data.status===1){
                    mmForm.value.useMmRes = 'Success'
                }else{//===0:failed
                    mmForm.value.useMmRes = 'Failed'
                }
            }).catch((err)=>{
                ElMessage({
                    type:'error',
                    message:'Get Transaction Status Failed!'
                })
            })
            // userStore.userProvider.httpProvider.getTransactionReceipt(res.txHash).then((res)=>{})
        }
   })
    // let txPar = ethUtils.createTxForMetamask(rawTx,form);
    // let a = await (window as any).ethereum.request(txPar)
}

const batchMint =()=>{
    // let walMap = userStore.userWallets.getHashWallets()
    if(!form.alreadyParseHash){
        ElMessage({
            type:"error",
            message:"未解析到交易Hash!"
        })
        return
    }else{
        startBtMint.value=true
        clearResult()
        let list = hmWalletList.value;
    // let cost =Number.parseFloat(maxCost.value.slice(0,7))
    // let validTx = []
    // for (var i = 0; i < list.length; i++) {
    //     if(cost > list[i].balance){
    //         hmWalletList.value[i]['status']=t('hashmint.msg.not_send')
    //     }else{
    //         validTx.push(list[i])
    //         let signer = userStore.userWallets.getHashMintSigner(list[i].address)
    //         let tx = ethUtils.createTxFromData(rawTx,form,list[i].address);
    //         if(signer._isSigner){
    //             signer.sendTransaction(tx).then((tx)=>{
    //                 let aa = tx
    //             }).catch((err:errors)=>{
    //                 let e =err;
    //             })
    //         }
    //     }
    // }

    //表格展示加载中...
     
    new Promise((resolve) => {
        ElMessage({
                type:'success',
                message:'Hash Mint Started!'
            })
        for (let i = 0; i < list.length; i++) {
            hmWalletList.value[i]['hash']='showLoading'
            hmWalletList.value[i]['status']='showLoading'
        }
        resolve(true)
    }).then(()=>{
        Promise.all(list.map(doMint)).then((res)=>{
            //记录发送交易成功的表格数据下标，
            let indexArr =[]
            let hashArr =[]
            res.forEach((item:any,ind)=>{
                if(item.status){
                    indexArr.push(ind)
                    hashArr.push((item as any).tx.hash)
                    hmWalletList.value[ind]['hash']=(item as any).tx.hash
                }else{
                    hmWalletList.value[ind]['status']=(item as any).msg
                }
            })
            Promise.all(hashArr.map(getBatchMintRes)).then((res)=>{
                res.forEach((item:any,ind)=>{
                    if(item.status){
                        //根据indexArr找到在表格中所属行
                        hmWalletList.value[indexArr[ind]]['status']= 'Success'
                    }else{// "rejected"
                        hmWalletList.value[indexArr[ind]]['status']= 'Failed'
                    }
                })
            })
        }).catch((err)=>{

            ElMessage({
                type:'error',
                message:err
            })
        }).finally(()=>{    
            ElNotification({
                title: 'HashMint',
                message: h('i', { style: 'color: teal' }, 'Hash Mint Finished'),
            })
        }).finally(()=>{

        })
    })
   
    }
    
}
const  clearResult=()=>{
    hmWalletList.value.forEach((item)=>{
        item.status = undefined
        item.hash = undefined
    })
}
const doMint=async(item:any)=>{
    let cost =Number.parseFloat(maxCost.value.slice(0,7))
    return new Promise(async (resolve) => {
        if(cost > item.balance){
            let list = hmWalletList.value;
            for (let i = 0; i < list.length; i++) {
                    if(item.address===hmWalletList.value[i]['address']){
                        hmWalletList.value[i]['hash']='Send Transaction Failed'
                        hmWalletList.value[i]['status']=t('hashmint.msg.not_send')
                    }
                }
            resolve({status:false,msg:t('hashmint.msg.not_send')})
        }else{
            try {
                let signer = userStore.userWallets.getSigner(item.address)
                let txdata = await ethUtils.createTxFromData(rawTx,form,item.address);
                if(signer._isSigner){
                     let tx = await signer.sendTransaction(txdata)
                     resolve({status:true,tx:tx})
                }else{
                    resolve({status:false,msg:'not signer'})
                }
            } catch (error) {
                
                let list = hmWalletList.value;
                for (let i = 0; i < list.length; i++) {
                    if(item.address===hmWalletList.value[i]['address'])
                    hmWalletList.value[i]['hash']='Send Transaction Failed'
                }
                resolve({status:false,msg:'error'})
                //{'code': -32000, 'message': 'already known'} :已经存在于内存池
                console.log(error);
            }
        }
    }).catch((err)=>{
        let list = hmWalletList.value;
        for (let i = 0; i < list.length; i++) {
            if(item.address===hmWalletList.value[i]['address']){
                hmWalletList.value[i]['hash']='Send Transaction Failed'
            }
        }
        console.log(err);
    
    })
}

const getBatchMintRes=async(hash)=>{
    return new Promise(async(resolve) => {
        try {
        //    let rec = await userStore.userProvider.httpProvider.waitForTransaction(hash,null,20000)
           let rec = await userStore.userProvider.httpProvider.waitForTransaction(hash)
           if(rec.status===1){
               resolve({status:true})
           }else{
            resolve({status:false})
           }
        } catch (error) {
            resolve({status:false})
        }
    }).catch((err)=>{
        
    })
}
/*
rec:
{
    "to": "0x1c4E5c4c6F8AaDF5F3B90EcE84F750f7293c188a",
    "from": "0xF406d85fc1e6D32E2d123D4ce50C940330F668c6",
    "contractAddress": null,
    "transactionIndex": 131,
    "gasUsed": {
        "type": "BigNumber",
        "hex": "0x5208"
    },
    "logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    "blockHash": "0xefc239ba5295094343e65254d3a3db926aaf508e972a5f486e5546101cdf90c4",
    "transactionHash": "0x30df89ab3d816350060b3583b482bab08921847335cf4e8ed87eb752450c40dd",
    "logs": [],
    "blockNumber": 8127282,
    "confirmations": 1,
    "cumulativeGasUsed": {
        "type": "BigNumber",
        "hex": "0x01c2ba1c"
    },
    "effectiveGasPrice": {
        "type": "BigNumber",
        "hex": "0x5cb09b93"
    },
    "status": 1,
    "type": 2,
    "byzantium": true
}

*/
//打开页面时，加载全局私钥钱包
onMounted(()=>{
    const {userWallets} =useUserStore()
    const {userProvider}=useUserStore()
    const {userInfo}=useUserStore()
    if(userInfo.isSign){
        const wals = userWallets.getAllWallets()
        wals.forEach(async w => {
            let addr = w.address;
                let wal ={
                    address:addr,
                    balance:w.balance,
                    hash:'',
                    status:'',
                }
                hmWalletList.value.push(wal)
        });
        
    }
    emitter.on('callHashMint',((hash:any)=>{
        form.txHash= hash;
        changeTxHash(hash);
    }))
})



</script>

<style lang="scss" scoped>
.hash-mint-warp{
    // width: 99.9%;
    height: 100%;
    box-shadow: var(--el-box-shadow-light);
    background-color: var(--main-bg-color);
    padding: 24px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    justify-items: center;
    align-items: center;
}
.hash-mint-form-wrap{
    border:#126b40 1px solid;
    // box-shadow: var(--el-box-shadow-light);
    background-color: var(--gd-gray700);
    width: 80%;
    min-width: 800px;
    padding: 40px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    .el_input_inner{
        width: 100%;
    }
    .el-form-item__label {
      color: #000000;
      font-size: 15px;
    }
}
.private-key-wrap{
    padding-top: 40px;
    padding-left: 40px;
    padding-right: 40px;

}
.pv-btn-list{
    padding-left: 90px;
}
.start-btn-list{
    padding-left: 40px;
    display: flex;
    flex-direction: column;
}  
.start-btn{
    border-radius: 20px;height: 35px;background-color: #126b40;
    border: 1px #126b40 solid;
    width: 135px;
} 
.start-btn:hover{
   border: #c7cd0b 2px solid;
   background-color: #c7cd0b;
}    
.start-wrap{
   border:#126b40 1px solid;
   margin-top: 20px;
   background-color: var(--gd-gray700);
    // box-shadow: var(--el-box-shadow-light);
    width: 80%;
    min-width: 800px;
    padding: 40px;
    padding-bottom: 20px;
    display: flex;
    flex-direction: column;
    justify-content: center;
}
.result-show-wrap{
    border:#126b40 1px solid;
   margin-top: 20px;
   background-color: var(--gd-gray700);
    // box-shadow: var(--el-box-shadow-light);
    width: 80%;
    min-width: 800px;
    padding: 40px;
    display: flex;
    flex-direction: column;
    justify-content: center;
}
.result-table{
    width: 100%;
    text-align: left;
    border-radius: 2px 2px 0 0;
    border-collapse: initial;
    border-spacing: 0;
}

:deep(.el-upload){
    margin-left: 150px !important;
}
:deep(.metamask-icon){

  height: 20px;
  width: 20px;
}

:deep(.el-upload-list__item .el-upload-list__item-info){
    display: inline-flex;
    justify-content: center;
    flex-direction: column;
     width: calc(100%); 
    margin-left: 4px;
}
:deep(.el-loading-spinner .circular){
    display: inline;
    height: 2em;
    width: 2em;
    -webkit-animation: loading-rotate 2s linear infinite;
    animation: loading-rotate 2s linear infinite;
}
.etherscan{
    width: 1.2em;
    height: 1.2em;
    vertical-align: -10%;
}
</style>