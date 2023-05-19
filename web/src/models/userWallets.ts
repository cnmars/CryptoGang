import { ethers } from "ethers";
import { ref,markRaw,toRefs ,toRaw} from "vue";
import { useUserStore } from '../pinia/modules/userStore'
import { ethUtils } from '@/utils/blockchain'
import { Wallet } from './singleWallet'
import { useWalletStore } from '@/pinia/modules/walletStore'
// import { httpProvider } from "../common/httpProvider";


export class UserWallets{
    loginAddress :string
    ens          :string
    balance      :string
    isBalanceLoading:boolean
    isEnsLoading:boolean
    walletMap    :Map<string,ethers.Wallet>
    constructor(address:string){
        const {userProvider} = useUserStore()
        this.isBalanceLoading = true
        this.isEnsLoading = true
        this.loginAddress = address
        let that = this
        userProvider.httpProvider.lookupAddress(address).then((ens)=>{
            that.ens= ens
        }).finally(()=>{
            this.isEnsLoading = false
        })
        ethUtils.getBalanceWithUint(address).then((bal)=>{
            that.balance = bal
        }).finally(()=>{
            this.isBalanceLoading = false
        })
        this.walletMap = new Map<string,ethers.Wallet>()
    }
   //传进来的是私钥，但是map的key是钱包地址
    public addWalletByPvKey=(pvk:string,tag ?:string)=>{
        const {walletTable} =useWalletStore()
        const userStore = useUserStore()
        let wal = new ethers.Wallet(pvk,userStore.userProvider.httpProvider)
        if(!this.getSigner(wal.address)){
            let w = new Wallet(pvk,wal.address,tag)
            this.walletMap.set(wal.address,wal)
            walletTable.push(w)
            this.updateWalletBalance(wal.address)
        }
    }
    //不传地址就是所有
    private updateWalletBalance =async(address ?:string) =>{
        const {walletTable} =useWalletStore()
        if(address){
            let i = walletTable.findIndex((item)=>item.address===address)
            if(!(i<0)){
             ethUtils.getBalanceWithUint(address).then((bal)=>{
                    walletTable[i].balance = bal
                    walletTable[i].isBalanceLoading = false
                }).catch((err)=>{
                    walletTable[i].balance = 'Unknown'
                    walletTable[i].isBalanceLoading = false
                })
            }
        }else{
            walletTable.forEach((row)=>{
                ethUtils.getBalanceWithUint(address).then((bal)=>{
                    row.balance = bal
                    row.isBalanceLoading = false
                }).catch((err)=>{
                    row.balance = 'Unknown'
                    row.isBalanceLoading = false
                })
            })
        }
    } 
    //不传地址就是所有
    private updateWalletEns =async(address ?:string) =>{
        const {walletTable} =useWalletStore()
        const {userProvider} = useUserStore()
        if(address){
            let i = walletTable.findIndex((item)=>item.address===address)
            if(!(i<0)){
                userProvider.httpProvider.lookupAddress(address).then((ens)=>{
                    walletTable[i].ens = ens
                    walletTable[i].isEnsLoading = false
                }).catch((err)=>{
                    walletTable[i].ens = 'Search Failed'
                    walletTable[i].isEnsLoading = false
                })
            }
        }else{
            walletTable.forEach((row)=>{
                userProvider.httpProvider.lookupAddress(address).then((ens)=>{
                    row.ens = ens
                    row.isEnsLoading = false
                }).catch((err)=>{
                    row.ens = 'Search Failed'
                    row.isEnsLoading = false
                })
            })
        }
            
    }
    //不传地址就是所有
     public waitTxFinish =async(address ?:string)=>{
         const {walletTable} =useWalletStore()
         const {userProvider} = useUserStore()
         if(address){
            let i = walletTable.findIndex((item)=>item.address===address)
            if(!(i<0)){
                let row = walletTable[i]
                if(row.opHash){
                    userProvider.httpProvider.waitForTransaction(row.opHash).then((res)=>{
                        row.opResult = res.status
                    }).catch ((error)=>{
                        console.log(address,error)
                    }).finally(()=>{
                        row.isOpFinished = true
                    })
                }
           }
        }else{
            walletTable.forEach((row)=>{
                if(row.opHash){
                    userProvider.httpProvider.waitForTransaction(row.opHash).then((res)=>{
                        row.opResult = res.status
                    }).catch ((error)=>{
                        console.log(address,error)
                    }).finally(()=>{
                        row.isOpFinished = true
                    })
                }
            })
        }
    }
    public getWalletByTag(tag :string){
        let walArr =[]
        const {walletTable} =useWalletStore()
        walletTable.forEach((row)=>{
            if(row.tag===tag){
                walArr.push(row)
            }
        })
        return walArr
    }
    public getWallet(address:string){
        const {walletTable} =useWalletStore()
        return walletTable
    }
    public removeWallet(address:string){
        this.walletMap.set(address,null)
        this.walletMap.delete(address)
        const {walletTable} =useWalletStore()
        walletTable.forEach((row,i)=>{
            if(row.address === address){
                walletTable.splice(i,1)
            }
        })
    }
    public getSigner(address:string){
        return this.walletMap.get(address)
    }
    public getAllWallets(){
        const {walletTable} =useWalletStore()
        return walletTable
    }
    //更新所有余额
    public updateAllBalance =async()=>{
        try {
            const {walletTable} =useWalletStore()
            walletTable.forEach((row)=>{
                row.updateBalance()
           })
        } catch (error) {
            console.log(error);
        }finally{
        }
    }
    //获取交易结果
    public waitAllTxFinish =async(provider:ethers.providers.BaseProvider)=>{
        try {
            const {walletTable} =useWalletStore()
            walletTable.forEach((row)=>{
                row.updateBalance()
           })
        } catch (error) {
            console.log(error);
        }finally{
        }
    }
}
