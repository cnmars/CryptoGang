import { localStorage } from '@/utils/storage';
import { defineStore } from 'pinia';
import { ref } from 'vue';
import { getLanguage } from '@/lang/index';
import {GdWebsocket} from '@/models/websocket'
export const useAppStore = defineStore('app', () => {

  const language = ref(getLanguage())
  const setLanguage=(lang: string) =>{
    language.value = lang;
    localStorage.set('language', language.value);
  }
  
  const ws = ref(new GdWebsocket())
  const PendingDataArr = ref([])
  const GasData = ref({
    baseFee:'0.0',
    fast:{
      maxFeePerGas:'0.0',
      maxPriorityFeePerGas:'0.0'
    },
    medium:{
      maxFeePerGas:'0.0',
      maxPriorityFeePerGas:'0.0'
    },
    slow:{
      maxFeePerGas:'0.0',
      maxPriorityFeePerGas:'0.0'
    },
  })
  const CoinPriceData = ref({ethereum:{cny:0.000,usd:0.000}})
  const InitWebsocket=(url:string)=>{
    ws.value.init(url)
  }
  const FillPendingData=(data)=>{
    PendingDataArr.value.push(data)
  }
  const FillGasData=(data)=>{
    GasData.value = data
  }
  const FillCoinPriceData=(data)=>{
    CoinPriceData.value = data
  }
  return {
    language,
    setLanguage,
    InitWebsocket,
    GasData,
    CoinPriceData,
    FillGasData,
    FillCoinPriceData,
    ws,
  }
})
