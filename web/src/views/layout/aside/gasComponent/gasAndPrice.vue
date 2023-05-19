<template>
   <div class="gas-wrap">
    <el-tooltip :content="'Base Fee:'+GasData.baseFee" raw-content placement="top" effect="base">
      <div class="gas-icon-wrap">
        <div style="width: 27px;">
          <svg-icon icon-class="gas"></svg-icon>
        </div>
        <div class="gas-base" :style="{backgroundColor: `${GasData.baseFee as any > 15 ?'#7d2a2a':''}`}"><span style="vertical-align: -5%;" >{{GasData.baseFee}}</span></div>
      </div>
    </el-tooltip>
      <div style="display: flex;">
        <el-tooltip 
        :content="'Fast<br>Max Fee: '+GasData.fast.maxFeePerGas +'<br> Priority Fee:' +GasData.fast.maxPriorityFeePerGas"
           raw-content placement="top" effect="fast">
            <div class="gas-fast-wrap">
              <div class="gas-fast-max">{{ GasData.fast.maxFeePerGas.toString().slice(0,4) }}
                <div class="gas-unit">Gwei</div>
              </div>
              <div class="gas-fast-pri">{{ GasData.fast.maxPriorityFeePerGas.toString().slice(0,3) }}</div>
            </div>
        </el-tooltip>

        <el-tooltip :content="'Medium<br>Max Fee: '+ GasData.medium.maxFeePerGas+ '<br> Priority Fee: '+GasData.medium.maxPriorityFeePerGas" raw-content placement="top" effect="mid">
         <div class="gas-mid-wrap">
          <div class="gas-mid-max">{{ GasData.medium.maxFeePerGas.toString().slice(0,4) }}
            <div class="gas-unit">Gwei</div>
          </div>
          <div class="gas-mid-pri">{{ GasData.medium.maxPriorityFeePerGas.toString().slice(0,3) }}</div>
        </div>
        </el-tooltip>

        <el-tooltip :content="'Slow<br>Max Fee: '+GasData.slow.maxFeePerGas+'<br> Priority Fee: '+GasData.slow.maxPriorityFeePerGas" raw-content placement="top" effect="low">
         <div class="gas-low-wrap">
          <div class="gas-low-max">{{ GasData.slow.maxFeePerGas.toString().slice(0,4) }}
            <div class="gas-unit">Gwei</div>
          </div>
          <div class="gas-low-pri">{{ GasData.slow.maxPriorityFeePerGas.toString().slice(0,3) }}</div>
        </div>
        </el-tooltip>
      </div>
        <div class="coin-price-wrap">
            <div class="price-title">
                <svg-icon icon-class="eth-logo" class-name="eth-coin-icon"></svg-icon>
                <span style="vertical-align: +35%;font-size: 11px;">ETH Price</span>
            </div>
            <div class="price-title">
                <span style="vertical-align: +35%;font-size: 16px;padding-left:7px;padding-top: 10px;">{{ `$${CoinPriceData.ethereum.usd.toString().slice(0,6)}` }}</span>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    export default{
        name:'GasAndPrice'
    }
</script>
<script lang="ts" setup>
import { useAppStore } from "@/pinia/modules/app";
import { storeToRefs } from "pinia";
const {ws} = useAppStore()
const appStore = useAppStore()
const {CoinPriceData} = storeToRefs(appStore)
const {GasData} = storeToRefs(appStore)
const {InitWebsocket} = useAppStore()
  // const {ws} = useAppStore()import.meta.env.VITE_BASE_API
  
  InitWebsocket(process.env.WSS_URL)
  // InitWebsocket('wss://cryptogang.vip/system/ws')
  // InitWebsocket('ws://cryptogang.vip:4728/system/ws')
  // const appStore = useAppStore()
  // setTimeout(() => {
  //   ws.Subscribe("coin")
    
  // }, 2000);
  const dealBaseFee = (GasData:any)=>{
    if((GasData as any).blockPrices){
    let a =  (GasData as any).blockPrices[0].baseFeePerGas.toString().slice(0,2)
    if(a.indexOf('.')>-1){
      a =  a.slice(0,1)
    }
    return a
  }else{
    return '13'
  }
  }
// const baseFee = ref('13')
// watch(() => GasData, () => {
//   if((GasData as any).blockPrices){
//     let a =  (GasData as any).blockPrices[0].baseFeePerGas.toString().slice(0,2)
//     if(a.indexOf('.')>-1){
//       a =  a.slice(0,1)
//     }
//     baseFee.value = a
//   }else{
//     baseFee.value= '13'
//   }
//   },{deep:true})
</script>

<style>
.el-popper.is-base {
  /* Set padding to ensure the height is 32px */
  padding: 6px 12px;
  margin-bottom: -5px !important;
  background: #adc79b;
}

.el-popper.is-base .el-popper__arrow::before {
  background: #adc79b;
  right: 0;
  margin-bottom: -10px !important;
}
.el-popper.is-fast {
  /* Set padding to ensure the height is 32px */
  padding: 6px 12px;
  margin-bottom: -5px !important;
  background: linear-gradient(90deg, rgb(145, 202, 139), rgb(216, 229, 129));
}

.el-popper.is-fast .el-popper__arrow::before {
  background: linear-gradient(45deg, #b2e68d, #e0e689);
  right: 0;
  margin-bottom: -10px !important;
}

.el-popper.is-mid {
  /* Set padding to ensure the height is 32px */
  padding: 6px 12px;
  margin-bottom: -5px !important;
  background: linear-gradient(90deg,#e0e689, #7b7b73);
}

.el-popper.is-mid .el-popper__arrow::before {
  background: linear-gradient(45deg, #e0e689, #7b7b73);
  right: 0;
  margin-bottom: -10px !important;
}

.el-popper.is-low {
  /* Set padding to ensure the height is 32px */
  padding: 6px 12px;
  margin-bottom: -5px !important;
  background: linear-gradient(90deg, #7b7b73, #ebebe8);
}

.el-popper.is-low .el-popper__arrow::before {
  background: linear-gradient(45deg, #7b7b73, #ebebe8);
  right: 0;
  margin-bottom: -10px !important;
}


</style>
<style lang="scss" scoped>
:deep(.el-tooltip__popper){
  
  border-radius: 4px !important;
  border-color:#7d2a2a !important;
  box-shadow: 0px 0px 7px 0px rgba(42, 42, 42, 0.2) !important;
}

.gas-wrap{
  height: 50px;
  width: 260px;
  background-color: var(--main-bg-color);
  // padding: 5px;
  display: flex;
  flex-direction: row;
  cursor: default;
  // grid-template-columns: 0.5fr 1fr 1fr 1fr;

}
.gas-icon-wrap{
  display: flex;
  width: 54px;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  border-right: 2px var(--gd-gas-border-right-color) solid;
  background-color: var(--gd-gas-bg-color);

}
.coin-price-wrap{
  display: flex;
  flex-direction: column;
  padding: 5px;
  padding-top: 10px;
  justify-content: center;
  align-items: center;
  background-color: var(--gd-coin-bg-color);background-size: 60px 60px;
}
.price-title{
  display: inline;
}
.eth-coin-icon{
  width: 1.3em;
  height: 1.3em;
}
.gas-base{
  font-size: medium;
  border: 2px green solid;
  background-color: green;
  padding: 1px;
  width: 27px;
  // background-color: #7d2a2a;
  text-align: center;
  vertical-align: middle;
  border-radius: 50%;
}
.gas-unit{
  font-size: 9px;
  color: var(--el-text-color-regular);
}
.gas-fast-wrap{
  text-align: center;
  align-items: center;
  width: 30px;
  padding: 5px;
  display: flex;
  flex-direction: column;
  flex-shrink: 0;
  border-right: 2px var(--gd-gas-border-right-color) solid;
  background-color: var(--gd-fast-bg-color);
}
.gas-fast-max{
  font-size: 14px;
  vertical-align: middle;
  text-align: center;
  border-bottom: 1px yellow solid;
}
.gas-fast-pri{
  font-size: 14px;
}

.gas-mid-wrap{
  text-align: center;
  align-items: center;
  padding: 5px;
  width: 30px;
  display: flex;
  flex-direction: column;
  flex-shrink: 0;
  border-right: 2px var(--gd-gas-border-right-color) solid;
  background-color: var(--gd-mid-bg-color);
}
.gas-mid-max{
  font-size: 14px;
  vertical-align: middle;
  text-align: center;
  border-bottom: 1px yellow solid;
}
.gas-mid-pri{
  font-size: 14px;
}

.gas-low-wrap{
  text-align: center;
  align-items: center;
  padding: 5px;
  width: 30px;
  display: flex;
  flex-direction: column;
  flex-shrink: 0;
  
  background-color: var(--gd-slow-bg-color);
}
.gas-low-max{
  font-size: 14px;
  vertical-align: middle;
  text-align: center;
  border-bottom: 1px yellow solid;
}
.gas-low-pri{
  font-size: 14px;
}

</style>