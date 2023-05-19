import { ethers } from "ethers";
import { useUserStore } from '../pinia/modules/userStore'
import{setProviderDefault,setProviderUrl} from '@/api/userApi'
import { ElMessage } from "element-plus";
import { UserProvider } from "./userProvider";
import { ElLoading } from 'element-plus'

export class Executor{

    httpProvider  : ethers.providers.JsonRpcProvider
    wssProvider   : ethers.providers.WebSocketProvider
    operationType : Array<string>
    results       : Array<Object>

    constructor(){
            // const {userProvider} = useUserStore()
            // this.httpProvider =  userProvider.httpProvider
            // this.wssProvider  =  userProvider.wssProvider
    }
    public preCheck=()=>{

    }
    public sendTx=()=>{

    }
    public operationRecord=()=>{
        
    }
}
