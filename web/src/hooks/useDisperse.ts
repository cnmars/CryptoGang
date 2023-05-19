import {ref} from 'vue'
import { ElMessageBox,ElMessage } from 'element-plus'
import { ethers } from 'ethers'
import { ethUtils } from '@/utils/blockchain'
import { useUserStore } from '@/pinia/modules/userStore'
export const useDisperse = ()=>{
    const handleTokenDisperse = async (tokenAddress:string,addressArray:[],amountArray:[],provider) => {
        // const provider = new ethers.providers.Web3Provider(window.ethereum);
        try {
            const signer = provider.getSigner();
            const disperseAbi = [
                "function disperseEther(address[] recipients, uint256[] values) external payable",
                "function disperseToken(IERC20 token, address[] recipients, uint256[] values) external",
                ]
            const disperseAddress = "0xD152f549545093347A162Dce210e7293f1452150"   
            const disperseContract = new ethers.Contract(disperseAddress,disperseAbi,provider);
            const disperseSigned = disperseContract.connect(signer);
        
            if ((addressArray.length === amountArray.length) && addressArray.length > 0) {
                let tx = await disperseSigned.disperseToken(tokenAddress,addressArray,amountArray);
                return tx
            } else {
                ElMessage({
                    type:'error',
                    message:'Please enter at least one valid transaction'
                })
                return null
            }
        } catch (error) {
            console.log(error);
            
            return null
        }
    }
 const handleEthDisperse = async (addressArray:[],amountArray:[],signer,isMm ?:boolean) => {
        try {
            const disperseAbi = [
                "function disperseEther(address[] recipients, uint256[] values) external payable",
                "function disperseToken(IERC20 token, address[] recipients, uint256[] values) external",
                ]
            const disperseAddress = "0xD152f549545093347A162Dce210e7293f1452150"   
            const disperseContract = new ethers.Contract(disperseAddress,disperseAbi,signer);
            const disperseSigned = disperseContract.connect(signer);
            const UNIT = 1000000000000000000;
            let valArr =[]
            let sum =0
            for (let i = 0; i < amountArray.length; ++i) {
                if (!isNaN((amountArray[i] as any)) && (parseFloat(amountArray[i]) > 0.0 )) {
                    valArr.push(ethers.utils.parseEther(amountArray[i]));
                    // valArr.push(('0x'+((amountArray[i] as any)*UNIT).toString(16)));
                    sum += parseFloat(amountArray[i])
                }
            }
            if ((addressArray.length === valArr.length) && addressArray.length > 0) {
                let tx = await disperseSigned.disperseEther(addressArray,valArr,{value:('0x'+((sum+0.001)*UNIT).toString(16))});
                return {status:true,tx:tx,message:"已完成！"}
            } else {
                return {status:false,tx:null,message:'Please enter at least one valid transaction'}
            }
        } catch (error) {
            console.log(error);
            return {status:false,tx:null,message:error}
        }
    }
    const handleEthCollect = async (addressArray:[],amountArray:[],isExtrAll:[],receiver:string) => {
       return new Promise(async (resolve, reject) => {
        try {
            let failedArr =[]
            if ((addressArray.length === amountArray.length) && (addressArray.length === isExtrAll.length) && addressArray.length > 0) {
                const {userProvider } = useUserStore()
                const gasPrice = await userProvider.httpProvider.getGasPrice()
                const gasLimit = 21000
                const gas = gasPrice.mul(gasLimit)
                const numGas = Number(ethers.utils.formatEther(gas.toString()))
                let data = []
                amountArray.forEach((item,v)=>{
                    let amt = item as number
                    if(isExtrAll[v]){//假如是提取则金额预留gas费
                        amt = item -numGas
                    }
                    data.push({
                        from:addressArray[v],
                        value:amt,
                        to:receiver
                    })
                })
                Promise.allSettled(data.map(doTransfer)).then((res)=>{
                   res.forEach((item,v)=>{
                        if(item.status!=='fulfilled'){
                            failedArr.push(item.reason.address)
                        }else{
                            console.log((item as any).value.tx.hash);
                        }
                   })
                   if(failedArr.length<1){
                    resolve({status:true,message:'',failed:failedArr})
                   }else{
                       resolve({status:false,message:'',failed:failedArr})
                   }
                })
            } else {
                resolve({status:false,message:'Please enter at least one valid transaction'})
            }
            
        } catch (error) {
            reject(error)
        }
       })
    }
    const doTransfer=async(data:any)=>{
        return new Promise(async(resolve, reject) => {
            try {
                const {userWallets } = useUserStore()
                const {userProvider } = useUserStore()
                let signer = userWallets.getSigner(data.from)
                if(signer._isSigner){
                    let a = ethers.utils.parseUnits(data.value.toString(),'ether')
                    const nonce = await userProvider.httpProvider.getTransactionCount(data.from,'latest')
                    let tx = {
                            to:data.to,
                            nonce:nonce,
                            // gasLimit:ethers.utils.hexlify(100000),
                            value:ethers.utils.parseUnits(data.value.toString(),'ether')
                        }
                    let r = await signer.sendTransaction(tx)
                    resolve({status:true,tx:tx,address:data.from}) 
                }else{
                    reject({status:false,msg:'not signer',address:data.from})
                }
            } catch (error) {
                reject({status:false,msg:error,address:data.from})
            }
        })
    }
    return{
        handleTokenDisperse,
        handleEthDisperse,
        handleEthCollect
    }
}
