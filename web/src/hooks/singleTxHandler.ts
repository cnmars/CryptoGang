import {ref} from 'vue'
import { ElMessageBox,ElMessage } from 'element-plus'
import { ethers } from 'ethers'
import { ethUtils } from '@/utils/blockchain'
export const useSingleTxHandler = ()=>{
    const checkTx=(tx:any)=>{
        return new Promise((resolve) => {
            //基本上to地址都是来自rawtx
            // if(!tx.to){
            //     ElMessageBox.alert("交易没有指定To地址！")
            //     resolve(false)
            //     ElMessage({
            //         type: 'info',
            //         message: '已取消',
            //     })
            // }
            if(tx.priorityFee){
                let a = ethers.utils.parseUnits(tx.priorityFee.toString(), "gwei")
                let b = ethers.utils.parseUnits('4', "gwei")
                // ethUtils.weiToGwei(a)
                if(ethUtils.weiToGwei(a)>4){
                    ElMessageBox.confirm(`您的小费设置为${tx.priorityFee} 您确认吗？`, '提示')
                        .then((res)=>{//‘confirm’
                            // ethUtils.createTxFromData(raw,tx,address)
                            resolve(true)
                        })
                        .catch((e) => {//'cancel' or 'close'
                            resolve(false)
                            ElMessage({
                                type: 'info',
                                message: '已取消',
                            })
                        })
                }else{
                    resolve(true)
                }
            }
        })
    }
    const singleTxMmSender = async(raw:any,tx:any,address ?:string)=>{
        let isValidTx =  await checkTx(tx);
        if(isValidTx){
            return new Promise(async (resolve) => {

                let validTx =  ethUtils.createTxForMetamask(raw,tx);
                try {
                    let txHash = await (window as any).ethereum.request(validTx)
                    resolve({status:true,txHash:txHash})
                } catch (error) {
                    ElMessage({
                        type: 'info',
                        message: error.message,
                    })
                    resolve({status:false,txHash:null})
                }
            })
        }else{
            return Promise.resolve({status:false,txHash:null})
        }
    }
    // const singleTxResHalder=(hash:string)=>{

    // }
    return{
        // checkAndCreateTx,
        singleTxMmSender,
        // singleTxResHalder,

    }
}

