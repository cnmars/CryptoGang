import { ethers } from "ethers";
import { useUserStore } from '../pinia/modules/userStore'
import{setProviderDefault,setProviderUrl} from '@/api/userApi'
import { ElMessage } from "element-plus";
import { ElLoading } from 'element-plus'

export class UserProvider{

    httpProvider  : ethers.providers.BaseProvider
    // httpProvider  : ethers.providers.JsonRpcProvider
    wssProvider   : ethers.providers.BaseProvider
    // wssProvider   : ethers.providers.WebSocketProvider

    constructor(httpUrl ?:string,wssUrl ?:string){
        if(httpUrl){
            this.httpProvider =  new ethers.providers.JsonRpcProvider(httpUrl)
        }else{
            //开发阶段使用georli测试网,正式环境为eth主网
            this.httpProvider =  new ethers.providers.JsonRpcProvider(process.env.GD_RPC_URL)
        }
        if(wssUrl){
            this.wssProvider  =  new ethers.providers.WebSocketProvider(wssUrl)
        }else{
            this.wssProvider  = new ethers.providers.WebSocketProvider(process.env.GD_WSS_URL)
        }
    }
    public setHttpProvider =async(url:string)=>{
        if(!url)return
        const loadingInstance = ElLoading.service({text:"正在连接节点...",fullscreen:false})
        this.httpProvider =  new ethers.providers.JsonRpcProvider(url)
        this.httpProvider.getBalance('0x474A7A1F35D018bF81A1F38c748D54D7F0Dd78bd').then(async (data)=>{
            const  userStore = useUserStore()
                userStore.userInfo.httpProvider = url
                userStore.userInfo.useDefaultHttp = false
                let res = await setProviderUrl({urlType:1,url:url}) as any
                if (res.code ===0){
                    loadingInstance.close()
                    return true
                }
        }).catch((err)=>{
            loadingInstance.close()
            this.httpProvider =  new ethers.providers.JsonRpcProvider(process.env.GD_RPC_URL)
            if(err.code==='NETWORK_ERROR'){
                ElMessage({
                    message:"请输入正确的地址！",
                    type:'error'
                })
                return
            }
        })
    } 
    public setWssProvider =async(url:string)=>{
        if(!url)return
        try {
            this.wssProvider =  new ethers.providers.WebSocketProvider(url)
        } catch (error) {
            ElMessage({
                message:"请输入正确的地址！",
                type:'error'
            })
            this.wssProvider  = new ethers.providers.WebSocketProvider(process.env.GD_WSS_URL)
            return false
        }
        const  userStore = useUserStore()
        userStore.userInfo.wssProvider = url
        userStore.userInfo.useDefaultWss = false
        let res = await setProviderUrl({urlType:2,url:url}) as any
        if (res.code ===0){
                return true
        }else{
            return false
        }
        
    }

    public setToDefault = (type:string)=>{
        const  userStore = useUserStore()
        switch (type) {
            case 'http':
                this.httpProvider =  new ethers.providers.JsonRpcProvider(process.env.GD_RPC_URL)
                userStore.userInfo.useDefaultHttp = true
                break;
            case 'wss':
                userStore.userInfo.useDefaultWss = true
                this.wssProvider  = new ethers.providers.WebSocketProvider(process.env.GD_WSS_URL)
                break;
        
            default:
                break;
        }
        setProviderDefault({urlType:type})
    }
}
