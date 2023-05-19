import { useUserStore } from "@/pinia/modules/userStore";
import { BigNumber, ethers } from "ethers";
import { netWorkMap } from "@/global/constants";

export const ethUtils={
    switchNetwork : async()=>{
        let ethereum = (window as any).ethereum
        try {
            await ethereum.request({
              method: 'wallet_switchEthereumChain',
              params: [{ chainId: '0x1' }],
            });
          } catch (switchError) {
            // This error code indicates that the chain has not been added to MetaMask.
            console.log(switchError)
            if (switchError.code === 4902) {
              //网络未添加
            }
            // handle other "switch" errors
          }
    },
    getNetWorkNameById :(id:string)=>{
        // let name = netWorkMap[id]
        // if(!name){
        //     name = 'Unknown Network'
        // }
        return netWorkMap[id] || 'Unknown Network'
    },
    isValidTxHash :(value: any) =>{
        if (value.substring(0, 2) !== "0x") { value = "0x" + value; }
        if (typeof(value) !== "string" || !value.match(/^0x[0-9A-Fa-f]*$/)) {
            return false
        }
        if(((value.length - 2) / 2)!=32){
            return false
        }
        return true;
    },
    weiToGwei : (value:any)=>{
        if (value?._isBigNumber){
            return Number(value)/1000000000
        }
        return value/1000000000
    },
    GweiToWei : (value:any)=>{
        return value*1000000000
    },
    //单位：eth
    getCurrentTransferGas:async()=>{
        const {userProvider } = useUserStore()
        const gasPrice = await userProvider.httpProvider.getGasPrice()
        const gasLimit = 21000
        const gas = gasPrice.mul(gasLimit)
        const numGas = Number(ethers.utils.formatEther(gas.toString()))
        return numGas
    },
    //单位：eth
    getTxGasFee:async(tx)=>{
        const {userProvider } = useUserStore()
        const gasPrice = await userProvider.httpProvider.getGasPrice()
        const gasLimit = Number(await userProvider.httpProvider.estimateGas(tx))
        const gas = gasPrice.mul(gasLimit)
        const numGas = Number(ethers.utils.formatEther(gas.toString()))
        return numGas
    },
    isHexString:(value: any, length?: number): boolean => {
        if (typeof(value) !== "string" || !value.match(/^0x[0-9A-Fa-f]*$/)) {
            return false
        }
        if (length && value.length !== 2 + 2 * length) { return false; }
        return true;
    },
    //verify eth address
    isValidAddress:(value):boolean=>{
        return value.match(/0x[a-fA-F0-9]{40}/)
    },
    isValidPVKey:(value):boolean=>{
        return (typeof(value) === "string") && (value.match(/^(0x)?[0-9Fa-f]{64}$/i)) as undefined as boolean
        
    },

    createTxFromData:async (raw :any,data:any,address:string)=>{
        const userStore = useUserStore()
        let provider = userStore.userProvider.httpProvider
        let nonce = await provider.getTransactionCount(address,'pending')
        // let nonce = await provider.getTransactionCount(address,'latest')
        data.gasLimit = (typeof data.gasLimit==='string' ? data.gasLimit:data.gasLimit.toString())
        let tx ={
            from    :  address,
            to      :  raw.to,
            data    :  data.data,
            value   :  ethers.utils.parseEther(data.value),
            // gasPrice:  ethers.utils.hexlify((Number.parseFloat(data.priorityFee) + Number.parseFloat(data.baseFee))*1000000000),
            gasLimit:  ethers.utils.hexlify(Number.parseFloat(data.gasLimit)),
            chainId :raw.chainId,
            nonce   :  nonce,
            maxFeePerGas:ethers.utils.parseUnits(data.maxFee.toString(), "gwei"),
            maxPriorityFeePerGas: ethers.utils.parseUnits(data.priorityFee.toString(), "gwei"),
        }
        
        return tx
    },
    createTxForMetamask :(raw,data)=>{
        const userStore = useUserStore()
        let tx = {
            method: 'eth_sendTransaction',
            params: [{
                nonce: data.nonce, // ignored by MetaMask
                // gasPrice:  ethers.utils.hexlify((Number.parseFloat(data.priorityFee) + Number.parseFloat(data.baseFee))*1000000000), // customizable by user during MetaMask confirmation.
                gas:    ethers.utils.hexlify(Number.parseFloat(data.gasLimit)),
                // gas: data.gasLimit.toHexString(), // customizable by user during MetaMask confirmation.
                to: raw.to, // Required except during contract publications.
                from: userStore.userInfo.address, // must match user's active address.
                value: raw.value.toHexString(), // Only required to send ether to the recipient from the initiating external account.
                data:raw.data, // Optional, but used for defining smart contract creation and interaction.
                chainId: raw.chainId,
            }],
          }
        return tx
    },
    sendTxByPvKey:async(tx:any,pvk:string)=>{
        const userStore = useUserStore()
        let provider = userStore.userProvider.httpProvider
        let wal = new ethers.Wallet(pvk,provider)
        let signer = wal.connect(provider);
        let gasPrice = provider.getGasPrice()
        const transaction = await signer.sendTransaction(tx)
        return transaction
    },
    getBalanceByPvk:async(address:string)=>{
        return new Promise(async(resolve,reject)=>{
            try {
                const userStore = useUserStore()
              let addr = new ethers.Wallet(address).address
              let res = await userStore.userProvider.httpProvider.getBalance(addr) as any
              const bal=ethers.utils.formatEther(res)
              resolve({
                address:addr,
                balance:bal.slice(0,7)
              })
            } catch (error) {
              reject({
                code :-1,
                msg:error
              })
            }
        })
    },
    getBalance:async(address:string)=>{
        return new Promise(async(resolve,reject)=>{
            try {
                const userStore = useUserStore()
              let res = await userStore.userProvider.httpProvider.getBalance(address) as any
              const bal=ethers.utils.formatEther(res)
              resolve({
                address:address,
                balance:bal
              })
            } catch (error) {
                console.log("getBalance");
                console.log(error);
                
              reject({
                code :-1,
                msg:error
              })
            }
        })
    },
    getBalanceWithUint:async(address:string)=>{
        try {
            const userStore = useUserStore()
            let res = await userStore.userProvider.httpProvider.getBalance(address) as any
            const bal=ethers.utils.formatEther(res)
            return bal.slice(0,7)+'ETH'
        } catch (error) {
            console.log("getBalance");
            console.log(error);
            return 'Unknown'
        }
    },
    isContract:async (address : string)=>{
        const {userProvider} = useUserStore()
        if(!address){
            return false
        }
        try {
            const code = await userProvider.httpProvider.getCode(address);
            if (code !== '0x') return true;
          } catch (error) {
            console.log(error);
          }
          return false
    },
    getContractType:async(address:string)=>{
        if(!address.match(/0x[a-fA-F0-9]{40}/)){
            return 0
        }else if(! await ethUtils.isContract(address)){
            return 0
        }else{
            const {userProvider} = useUserStore()
            const abiERC721 = ["function supportsInterface(bytes4) public view returns(bool)"];
            const contractERC721 = new ethers.Contract(address, abiERC721, userProvider.httpProvider)
            const selectorERC721 = "0x80ac58cd"
            const selectorERC1155 = "0xd9b67a26"
            try {
                const isERC721 = await contractERC721.supportsInterface(selectorERC721)
                if(isERC721)return 721  
                const isERC1155 = await contractERC721.supportsInterface(selectorERC1155)
                if(isERC1155)return 1155
                else{
                    
                }
            } catch (error) {
                const abiERC20=["function decimals() external view returns (uint8)"];
                const contract = new ethers.Contract(address, abiERC20, userProvider.httpProvider);
                try {
                const response = await contract.decimals()
                const decimals = Number(response)
                if (isNaN(decimals)) {
                    // Most likely not an ERC20 contract
                    return 0
                }
                if (decimals >= 0) {
                    // Most likely an ERC20 contract
                    return 20
                }
                } catch (error) {
                // Unable to determine if it's ERC20 due to error
                // This might mean that the contract is not an ERC20, or that something else happened
                console.log(error);
                    return 0
                }
            }
        }
    },
    isErc20ByAbi :(abi)=> {//global下的ERC20abi文件
        const signatures = new Set();
        for (const i of abi['abi']) {
          signatures.add(`${ i.name }(${ i.inputs.map(inp => inp.type).join(",") })`);
        }
        return signatures.has("transfer(address,uint256)")
          && signatures.has("transferFrom(address,address,uint256)")
          && signatures.has("approve(address,uint256)")
          && signatures.has("decimals()")
          && signatures.has("totalSupply()")
          && signatures.has("balanceOf(address)");
      },
    
      isERC20ByAddress :async (address:string) => {
        if(!address.match(/0x[a-fA-F0-9]{40}/)){
            return false
        }
        else if(! await ethUtils.isContract(address)){return false}
        else{
            const {userProvider} = useUserStore()
            const abiERC20=["function decimals() external view returns (uint8)"];
            const contract = new ethers.Contract(address, abiERC20, userProvider.httpProvider);
            try {
              const response = await contract.decimals()
              const decimals = Number(response)
              if (isNaN(decimals)) {
                // Most likely not an ERC20 contract
                return false
              }
              if (decimals >= 0) {
                // Most likely an ERC20 contract
                return true
              }
            } catch (error) {
              // Unable to determine if it's ERC20 due to error
              // This might mean that the contract is not an ERC20, or that something else happened
              console.log(error);
              return false
            }
        }
    },
    isERC721ByAddress :async(address:string)=>{
        if(!address.match(/0x[a-fA-F0-9]{40}/)){
            return false
        }
        else if(! await ethUtils.isContract(address)){return false}
        else{
            const {userProvider} = useUserStore()
            const abiERC721 = [
                // "function name() view returns (string)",
                // "function symbol() view returns (string)",
                "function supportsInterface(bytes4) public view returns(bool)",
            ];
            const contractERC721 = new ethers.Contract(address, abiERC721, userProvider.httpProvider)
            const selectorERC721 = "0x80ac58cd"
            /**
             * const _INTERFACE_ID_IERC165 = '0x01ffc9a7';
                const _INTERFACE_ID_IERC721 = '0x80ac58cd';
                const _INTERFACE_ID_IERC721METADATA = '0x5b5e139f';
                const _INTERFACE_ID_IERC721ENUMERABLE = '0x780e9d63';
                1155:0xd9b67a26
             */
            const isERC721 = await contractERC721.supportsInterface(selectorERC721)
            return isERC721
    
        }
        
    }    

}