import { ElMessageBox,ElMessage } from 'element-plus'
import { ethers } from 'ethers'
import { ethUtils } from '@/utils/blockchain'
export const useBatchTxHandler = ()=>{
    //这里只判断了小费设置是否合适，后期可加入其他判断
    const checkTx=(txs:any)=>{
        if(txs.length<1)return Promise.resolve(false)
        //批量操作中小费都是一样的，这里只取出一个来判断就好
        let tx = txs[0]
        return new Promise((resolve) => {
            if(tx[0].priorityFee){
                let a = ethers.utils.parseUnits(tx[0].priorityFee.toString(), "gwei")
                // let b = ethers.utils.parseUnits('4', "gwei")
                if(ethUtils.weiToGwei(a)>10 || (ethUtils.weiToGwei(a)<2)){
                    ElMessageBox.confirm(`您的小费设置为${tx.priorityFee} 您确认吗？`, '提示')
                        .then((res)=>{//‘confirm’
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
    const batchTxSender = async(raw:any,txs:[],address :[])=>{
        if(txs.length<1)return Promise.resolve({status:false,txHash:null})
        if(txs.length != address.length)return Promise.resolve({status:false,txHash:null})
        let isValidTx =  await checkTx(txs);
        if(isValidTx){
            let TxArr = []
            txs.forEach((t,i)=>{
                TxArr.push(ethUtils.createTxFromData(raw,t,address[i]))
            })
            // Promise.allSettled()
            // return new Promise(async (resolve) => {
            //     let validTx =  ethUtils.createTxForMetamask(raw,tx);
            //     try {
            //         let txHash = await (window as any).ethereum.request(validTx)
            //         resolve({status:true,txHash:txHash})
            //     } catch (error) {
            //         ElMessage({
            //             type: 'info',
            //             message: error.message,
            //         })
            //         resolve({status:false,txHash:null})
            //     }
            // })
        }else{
            return Promise.resolve({status:false,txHash:null})
        }
    }
    // const singleTxResHalder=(hash:string)=>{

    // }
    return{
        // checkAndCreateTx,
        batchTxSender,
        // singleTxResHalder,

    }
}

