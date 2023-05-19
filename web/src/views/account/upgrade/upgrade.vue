<template>
    <div>
        <div class="content-wrapper" v-loading="upgradeLoading" :element-loading-text="loadText">
            <h1 class="text-center">{{$t('upgrade.title')}}</h1>
            <div class="container-fluid">
            <div class="row">
                <div class="card shadow-lg" >
                    <div class="card-body">
                    <div class="text-center p-3">
                        <h5 class="card-title" style="font-size:2rem !important">$100 {{$t('upgrade.month')}}</h5>
                    </div>
                    <p class="card-text" style="text-align:center;">{{$t('upgrade.tip')}}</p>
                    </div>
                    <ul class="list-group list-group-flush">
                        <template v-for="(item,key) in monthItems" :key="key">
                            <li class="list-group-item" style="color: #000 !important;" v-if="item.own && !item.exclusive">
                                <svg-icon icon-class="select" class-name="select-icon"></svg-icon> 
                               {{ $t('route.'+item.name) }}
                            </li>
                            <li class="list-group-item" style="background-color: #ffaaaa !important;color: #000 !important;" v-if="!item.own">
                                <svg-icon icon-class="no" class-name="select-icon"></svg-icon>
                                {{ $t('route.'+item.name) }}
                            </li>
                            <li class="list-group-item" style="background-color: rgb(203, 184, 17) !important;color: #000 !important;" v-if="item.own && item.exclusive">
                                <svg-icon icon-class="no" class-name="select-icon"></svg-icon>
                                {{ $t('route.'+item.name) }}
                            </li>
                        </template>
                    </ul>
                    <div class="card-body text-center">
                        <span class="btn-outline-primary"  style="border-radius:30px" @click="buy(1)">{{$t('upgrade.buy')}}</span>
                    </div>
                </div>

                <div class="card shadow-lg">
                    <div class="card-body">
                    <div class="text-center p-3">
                        <h5 class="card-title" style="font-size:2rem !important">$1000 {{$t('upgrade.year')}}</h5>
                    </div>
                    <p class="card-text" style="text-align:center;">{{$t('upgrade.tip')}}</p>
                    </div>
                    <ul class="list-group list-group-flush" >
                        <template v-for="(item,key) in yearItems" :key="key">
                            <li class="list-group-item" style="color: #000 !important;" v-if="item.own && !item.exclusive">
                                <svg-icon icon-class="select" class-name="select-icon"></svg-icon> 
                                {{ $t('route.'+item.name) }}
                            </li>
                            <li class="list-group-item" style="background-color: #ffaaaa !important;color: #000 !important;" v-if="!item.own">
                                <svg-icon icon-class="no" class-name="select-icon"></svg-icon>
                                {{ $t('route.'+item.name) }}
                            </li>
                            <li class="list-group-item" style="background-color: rgb(203, 184, 17) !important;color: #000 !important;" v-if="item.own && item.exclusive">
                                <svg-icon icon-class="select" class-name="select-icon"></svg-icon>
                                {{ $t('route.'+item.name) }}
                            </li>
                        </template>
                    </ul>
                    <div class="card-body text-center">
                    <span class="btn-outline-primary"  style="border-radius:30px" @click="buy(2)">{{$t('upgrade.buy')}}</span>
                    </div>
                </div>
                </div>
            </div>
          </div>
    </div>
</template>

<script lang="ts">
export default{
    name:"upgrade"
}
</script>
<script lang="ts" setup>
import { ref } from "vue";
import { useUserStore } from '@/pinia/modules/userStore'
import { useRouterStore } from '@/pinia/modules/router'
import { ElMessage ,ElNotification} from 'element-plus'
import { h} from 'vue'
import { getUpgradePrice } from "@/api/system";
import { upgrade } from "@/api/userApi";
import { ethers } from "ethers";
import { useRouter } from 'vue-router'
import { emitter } from '@/common/bus'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const router = useRouter()
const upgradeLoading = ref(false)
const loadText = ref(t('upgrade.wait'))
//exclusive :独有
const monthItems =[
    { name:"quickOffer", own:true,exclusive:false },
    { name:"quickSell", own:true,exclusive:false },
    { name:"freeMint", own:true,exclusive:false },
    { name:"followMint", own:true,exclusive:false },
    { name:"sweepFloor", own:true,exclusive:false },
    { name:"hotToken", own:true,exclusive:false },
    { name:"bigFlow", own:true,exclusive:false },
    { name:"tokenAnalysis", own:true,exclusive:false },
    { name:"walletAnalysis",own:true,exclusive:false},
    { name:"frontRun",own:true,exclusive:false},
    { name:"newFeatures",own:false,exclusive:false},
]
const yearItems =[
    { name:"quickOffer", own:true,exclusive:false },
    { name:"quickSell", own:true,exclusive:false },
    { name:"freeMint", own:true,exclusive:false },
    { name:"followMint", own:true,exclusive:false },
    { name:"sweepFloor", own:true,exclusive:false },
    { name:"hotToken", own:true,exclusive:false },
    { name:"bigFlow", own:true,exclusive:false },
    { name:"tokenAnalysis", own:true,exclusive:false },
    { name:"walletAnalysis",own:true,exclusive:false},
    { name:"frontRun",own:true,exclusive:false},
    { name:"newFeatures",own:true,exclusive:true},
]

const buy = async(type:number)=>{
    //不管小狐狸当前账号是什么，生效账号为用户登录账号
    //先判断网络
    let netid = (window as any).ethereum.networkVersion
    if(netid !==1){
        try{
            let id =1;
            await (window as any).ethereum.request({
                method: 'wallet_switchEthereumChain',
                params: [{ chainId:('0x'+id.toString(16)) }]
            })
            upgradeLoading.value=true
            const userStore = useUserStore()
            const provider = new ethers.providers.Web3Provider((window as any).ethereum);
            Promise.allSettled([
                function(){
                return new Promise(async(resolve) => {
                    try {
                        let bal = await provider.getBalance(userStore.userInfo.address)
                        // let bal = await userStore.userProvider.httpProvider.getBalance(userStore.userInfo.address)
                        const balance=ethers.utils.formatEther(bal)
                        resolve(balance)
                    } catch (error) {
                        resolve("")
                    }
                    })
                }(), 
                function(){
                return new Promise(async(resolve) => {
                    try {
                        //后端获取当前ETH价格，（当前eth价格、套餐价格应该都写成后端接口）
                        //套餐价格后端已经处理成Eth
                        let res = await getUpgradePrice()
                        resolve(res)
                    } catch (error) {
                        resolve("")
                    }
                    })
                }(),
            ]).then(async(s:any)=>{
                let balance = s[0].value
                let res = s[1].value
                if((res as any).code ===0){
                let data =  (res as any).data.upgradePrice
                let month = data.priceMonth
                let year = data.priceYear
                let need
                let value
                switch (type) {
                    case 1:
                        need = month
                        value = ethers.utils.parseUnits(month.toString(),'ether')
                        break;
                    case 2:
                        need = year
                        value = ethers.utils.parseUnits(year.toString(),'ether')
                        break;
                    
                    default:
                        break;
                }
                if(data.unit =='eth'){
                    //检查余额
                    if(month> balance){
                        ElNotification({
                            title: 'Failed',
                            message: h('div', { style: 'color: red' }, '余额不足'),
                            })
                    }else{
                        //发送转账交易
                        // const provider = new ethers.providers.Web3Provider((window as any).ethereum);
                        let signer = provider.getSigner();
                        const nonce = await provider.getTransactionCount((window as any).ethereum.selectedAddress,'latest')
                        let tx = {
                                from: (window as any).ethereum.selectedAddress,
                                to:data.receiptor,
                                nonce:nonce,
                                // gasLimit:ethers.utils.hexlify(100000),
                                value:value
                            }
                        let restx = await signer.sendTransaction(tx)
                        document.querySelector(".el-loading-text").innerHTML = t('upgrade.cnf')
                        // loadText.value=`${t('upgrade.cnf')}`
                        let hash = restx.hash
                        let to = restx.to
                        let rec = await provider.waitForTransaction(hash)
                        if(rec.status===1 && rec.to.toLowerCase()==to.toLowerCase() && hash.toLowerCase()==rec.transactionHash.toLowerCase()){
                            let r = await upgrade({type:type,hash:hash})
                            if((r as any).code ===0){
                                ElNotification({
                                    title: 'Success',
                                    message: h('div', { style: 'color: green' }, '购买成功！'),
                                })
                            }else{
                                ElNotification({
                                    title: 'Failed',
                                    message: h('div', { style: 'color: red' }, '购买失败！如出现已付款成功请联系客服解决'),
                                })
                            }
                            const routerStore = useRouterStore()
                            await routerStore.SetAsyncRouter()
                            await userStore.GetUserInfo(userStore.userInfo.address)
                            const asyncRouters = routerStore.asyncRouters
                            router.options.routes= asyncRouters 
                            asyncRouters.forEach(asyncRouter => {
                                router.addRoute(asyncRouter)
                            })
                            emitter.emit('closeThisPage')
                            emitter.emit('reload')
                        }else{
                            ElNotification({
                                title: 'Upgrade',
                                message: h('div', { style: 'color: red' }, '购买失败！'),
                            })
                        }
                        //等待交易并判断交易结果（这里是否要防止恶意攻击，比如替换交易？）
                    }
                }else{
                    ElMessage.error("内部错误")
                }
                }
            }).finally(()=>{
                upgradeLoading.value=false
            })             
        }catch(err){
            upgradeLoading.value=false
            ElMessage.error("购买失败！")
            if (err.code === 4902) {
                ElMessage.error("小狐狸未添加以太坊主网");
            }
        }finally{
            // upgradeLoading.value=false
        }
    }
}

</script>

<style lang="scss" scoped>
:deep(.el-loading-spinner .circular) {
    display: inline;
    height: 80px;
    width: 80px;
    -webkit-animation: loading-rotate 2s linear infinite;
    animation: loading-rotate 2s linear infinite;
}
:deep(.el-loading-mask .el-loading-spinner .el-loading-text) {
  font-size: 25px;
  font-weight: bold;
}
.card {
      border:none;
      padding: 10px 50px;
    }

    .card::after {
      position: absolute;
      z-index: -1;
      opacity: 0;
      -webkit-transition: all 0.6s cubic-bezier(0.165, 0.84, 0.44, 1);
      transition: all 0.6s cubic-bezier(0.165, 0.84, 0.44, 1);
    }
    .card:hover {
      transform: scale(1.02, 1.02);
      -webkit-transform: scale(1.02, 1.02);
      backface-visibility: hidden;
      will-change: transform;
      box-shadow: 0 1rem 3rem rgba(0,0,0,.75) !important;
    }

    .card:hover::after {
      opacity: 1;
    }

    .card:hover .btn-outline-primary{
      color:var(--el-color-primary);
      font-size: 18px;
      background:#fff;
    //background:#3fad3d;
    
      padding: 0.6rem 3rem;
    // font-size: 0.875rem;
    line-height: 1.2;
    border-radius: 0.1875rem;
    filter: drop-shadow(2px 2px 5px #fff);
    // cursor: pointer !important;
    }
    .btn-outline-primary{
       color:var(--el-color-primary);
      font-size: 18px;
      background:rgb(212, 213, 210);
    //background:#3fad3d;
    
      padding: 0.6rem 3rem;
    // font-size: 0.875rem;
    line-height: 1.2;
    border-radius: 0.1875rem;
    
    cursor: pointer !important;
}
    .content-wrapper {
    height: auto;
    min-height: 98%;
    box-shadow: var(--el-box-shadow-light);
    background-color: var(--bg-color);
    display: flex;
    flex-direction: column;
    align-items: center;
}
.select-icon{
    height: 1.2rem;
    width: 1.2rem;
    vertical-align: -15%;
}

.row {
    display: flex;
    flex-direction: row;
    gap: 80px;
    flex-wrap: wrap;
    justify-content: center;
}

.grid-margin {
    margin-bottom: 1.5rem;
}

.card {
    border: none;
    padding: 10px 50px;
}
.card {
    border-radius: 0.25rem;
}
.card {
    position: relative;
    display: flex;
    flex-direction: column;
    min-width: 0;
    word-wrap: break-word;
    // background-color: #191c24;
    background-color: var(--gd-card-bg);
    background-clip: border-box;
    border: 1px solid var(--gd-card-bg);
    // border: 1px solid rgba(0, 0, 0, 0.125);
    border-radius: 0.25rem;
}

.card .card-body {
    padding: 1.75rem 1.5625rem;
}

.card-body {
    flex: 1 1 auto;
    min-height: 1px;
    padding: 1.25rem;
}
.text-center {
    color: var(--el-color-primary);
    text-align: center !important;
}

h1, .h1 {
    font-size: 2.19rem;
}


.container-fluid {
    width: 100%;
    padding-right: 0.75rem;
    padding-left: 0.75rem;
    margin-right: auto;
    margin-left: auto;
}

.shadow-lg {
    box-shadow: 0 1rem 3rem rgb(0 0 0 / 18%) !important;
    min-width: 290px;
}
.card .card-body {
    padding: 1.75rem 1.5625rem;
}
.card-body {
    flex: 1 1 auto;
    min-height: 1px;
    padding: 1.25rem;
}

.p-3 {
    padding: 1rem !important;
}
.card-text:last-child {
    margin-bottom: 0;
}
.card .card-title {
    color: #ffffff;
    margin-bottom: 1.125rem;
    margin-top: 0.8rem;
    text-transform: capitalize;
}
.card-text:last-child {
    margin-bottom: 0;
}
p {
    font-size: 0.875rem;
}
.card > .list-group {
    border-top: inherit;
    border-bottom: inherit;
}

.list-group-flush {
    border-radius: 0;
}
.list-group {
    display: flex;
    flex-direction: column;
    padding-left: 0;
    margin-bottom: 0;
    border-radius: 0.25rem;
}
ol, ul, dl {
    margin-top: 0;
    margin-bottom: 1rem;
}
.list-group-item {
    position: relative;
    display: block;
    padding: 0.75rem 1.25rem;
    background-color: #fff;
    background-color: var(--gd-card-item-bg) !important;
    border: 1px solid rgba(0, 0, 0, 0.125);
}
.list-group-flush > .list-group-item {
    border-width: 0 0 1px;
}
</style>