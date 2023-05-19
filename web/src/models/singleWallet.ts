import { ethers } from "ethers";
export class Wallet{//这个类不能写方法，不然数据没法响应式
    address         : string
    pvk             : string
    ens             : string
    balance         : string
    isBalanceLoading: boolean
    isEnsLoading    : boolean
    tag             : string//标记分组
    signer          : ethers.Wallet
    opValue         : string//操作金额（代币分发 归集时的value）
    opHash          : string//交易hash
    opResult        : number//操作结果
    isOpFinished    : boolean//操作是否在pending中
    constructor(pvk:string,address:string,tag ?:string){
        this.isBalanceLoading =true
        this.isEnsLoading =true
        this.isOpFinished = false
        this.address = address
        this.pvk = pvk
        this.balance ='0'
        if(tag){
            this.tag =tag
        }
    }
}