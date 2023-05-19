import { getUserInfoFromServer } from '../../api/userApi';
import { saveTheme } from '../../api/userApi';
import router from '../../router/router';
import { ElMessageBox, ElMessage,ElNotification } from 'element-plus';
import { defineStore } from 'pinia';
import { ref, computed, watch,markRaw,h} from 'vue';
import { useRouterStore } from './router';
import { ethers } from "ethers";
import { getSignMsg ,signAndLogin} from '../../api/userApi';
import { jsonInBlacklist} from '../../api/jwt';
import { LOCAL_STORAGE } from "../../global/constants";
import { UserWallets } from '@/models/userWallets';
import { UserProvider } from '@/models/userProvider';
import { ethUtils } from "@/utils/blockchain";
import type { Action } from 'element-plus';
import { emitter } from '@/common/bus';
import { useAppStore } from "@/pinia/modules/app";
//标记是否已经设置用户提供的provider
let initProviderFlag = false
export const useUserStore = defineStore('user', () => {
  const loadingInstance = ref()

  const userInfo = ref({
    address:'',
    balance:'',
    ens:'',
    httpProvider:'',
    wssProvider:'',
    showNetwork:'',
    useDefaultHttp:true,
    useDefaultWss:false,
    avatar:'',
    expireAt:'',
    theme:'dark',
    isSign:false,
    isConnect:false,
    showBal:"0",
    authorityId:333,
    authority: {},
    sideMode: 'dark',
    activeColor: '#4D70FF',
    baseColor: '#fff',
    
  })
  const userWallets = ref()

  var userProvider= new UserProvider()
  userProvider = markRaw(userProvider)

  const setUserProvider=()=>{
    if(userInfo.value.httpProvider && !userInfo.value.useDefaultHttp){
      userProvider.setHttpProvider(userInfo.value.httpProvider)
    }
    if(userInfo.value.wssProvider && !userInfo.value.useDefaultWss){
      userProvider.setWssProvider(userInfo.value.wssProvider)
    }
  }
  const token = ref(window.localStorage.getItem('token') || '')
  const setUserInfo = (val:any) => {
    userInfo.value = val
    userInfo.value.isConnect = true
    userInfo.value.isSign = true
    if(!initProviderFlag){
      setUserProvider()
      initProviderFlag = true
    }
  }

  const NeedInit = () => {
    token.value = ''
    window.localStorage.removeItem('token')
    localStorage.clear()
    router.push({ name: 'Init', replace: true })
  }

  const ResetUserInfo = (value = {}) => {
    userInfo.value = {
      ...userInfo.value,
      ...value
    }
  }
  /* 获取用户信息*/
  const GetUserInfo = async(address:string) => {
    if(!address)return
    
    if(!userWallets || !userWallets.value ){
      userWallets.value = new UserWallets(address)
    }
    let oldAuthid = userInfo.value.authorityId
    getUserInfoFromServer({"address":address}).then(async(res:any)=>{
      if (res?.code === 0) {
        //这个if是后期加的，为了不重新写user数据倒是顶部栏因为isConnect和isSign变化而展示出按钮
        if(res.data.userInfo.address == userInfo.value.address){
          return userInfo
        }
        setUserInfo(res.data.userInfo)
        if(window.localStorage.getItem(LOCAL_STORAGE.USER_ENS)){ 
          userInfo.value.ens = window.localStorage.getItem(LOCAL_STORAGE.USER_ENS)
        }else{
  
        }
        if(window.localStorage.getItem(LOCAL_STORAGE.USER_BAL)){ 
          userInfo.value.balance = window.localStorage.getItem(LOCAL_STORAGE.USER_BAL)
        }else{
          userProvider.httpProvider.getBalance(address).then((res)=>{
            const bal=ethers.utils.formatEther(res)
            window.localStorage.setItem(LOCAL_STORAGE.USER_BAL,bal)
            userInfo.value.balance = bal
    
          })
        }
        let netid = (window as any).ethereum.networkVersion
        let networkName = ethUtils.getNetWorkNameById(netid)
        userInfo.value.showNetwork = networkName
        window.localStorage.setItem(LOCAL_STORAGE.USER_NET,networkName)
        if(window.localStorage.getItem(LOCAL_STORAGE.USER_AVATAR)){ 
          userInfo.value.avatar = window.localStorage.getItem(LOCAL_STORAGE.USER_AVATAR)
        }else{
          userProvider.httpProvider.getResolver(userInfo.value.ens).then((resolver)=>{
            resolver?.getText("avatar").then((ava)=>{
              if(ava){
                userInfo.value.avatar = ava
                window.localStorage.setItem(LOCAL_STORAGE.USER_AVATAR,ava)
              }
            })
          })
        }
        if(oldAuthid!==userInfo.value.authorityId){//说明会员过期了
          emitter.emit('reload')//刷新网页会重新加载菜单
        }
      }
    })
  }

/* 登录*/
const ConnectAndLogin = async()=>{
  if(!(window as any).ethereum){
    ElMessageBox.confirm(`
        <div style="flex-direction:row;display:flex;justify-content:flex-between;gap:5px;vertical-align:middle;">
          <div>
            <p>请安装Metamask(小狐狸)插件使用</p>
          </div>
        </div>
        `, '系统错误', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: '清理缓存',
          cancelButtonText: '取消'
        })
          .then(() => {
            /* if((window as any).ActiveXObject) {
              console.log("IE")
            }else  */if((document as any).getBoxObjectFor) {
              //console.log("Firefox")
              window.open('https://addons.mozilla.org/firefox/addon/ether-metamask/?utm_source=addons.mozilla.org&utm_medium=referral&utm_content=search')
            }else if(window.MessageEvent && !(document as any).getBoxObjectFor) {
              //console.log("Chrome")
              window.open('https://chrome.google.com/webstore/detail/metamask/nkbihfbeogaeaoehlefnkodbefgpgknn')
            }else if((window as any).openDatebase) {
              //console.log("safair")
            }else{
              ElMessage.info("除Chrome和FireFox外，其他浏览器请自行搜索安装")
            }
            return
          }).catch((action: Action) => {
            return
            // ElMessage({
            //   type: 'info',
            //   message:
            //   /* ESC 和 点击关闭分别对应'cancel' 和 'close'。 */
            //     action === 'cancel' ? "": "",
            // })
          })
  }else{
    let win = (window as any)
    const provider = await new ethers.providers.Web3Provider(win.ethereum)
    const accounts = await provider.send("eth_requestAccounts",[])
    if(accounts[0]){
      userWallets.value = new UserWallets(accounts[0])
      const msg = await getSignMsg({"address":accounts[0]}) as any
      const signer = provider.getSigner(); //获取签名者信息
      let userPk = await signer.getAddress(); // 获取签名者公钥
      if(msg.code !==0)return
      let signature = await signer.signMessage(msg.data.msg+msg.data.nonce) // metamask 对“登录网站进行签名”
      let res = await signAndLogin({address:accounts[0],signature:signature}) as any
      if (res.code === 0) {

        setUserInfo(res.data.user)
        setToken(res.data.token)
        window.localStorage.setItem(LOCAL_STORAGE.THEME,res.data.user.siteTheme)
        let netid = (window as any).ethereum.networkVersion
        let networkName = ethUtils.getNetWorkNameById(netid)
        userInfo.value.showNetwork = networkName
        if(netid!=='1'){
          ElNotification({
            title: 'Network',
            message: h('div', { style: 'color: red' }, '当小狐狸网络不是以太主网时，只能用私钥进行交易，而不能采用小狐狸交易。'),
          })
        }
        window.localStorage.setItem(LOCAL_STORAGE.USER_ADDRESS,accounts[0])
        userInfo.value.isConnect = true
        userInfo.value.isSign = true
        provider.getBalance(accounts[0]).then((res)=>{
          const bal=ethers.utils.formatEther(res)
          window.localStorage.setItem(LOCAL_STORAGE.USER_BAL,bal)
          userInfo.value.balance = bal
  
        })
        provider.lookupAddress(accounts[0]).then((res)=>{
          window.localStorage.setItem(LOCAL_STORAGE.USER_ENS,res)
          userInfo.value.ens =res
        })
        const routerStore = useRouterStore()
        await routerStore.SetAsyncRouter()
        const asyncRouters = routerStore.asyncRouters
        router.options.routes= asyncRouters 
        asyncRouters.forEach(asyncRouter => {
          router.addRoute(asyncRouter)
        })
        const {InitWebsocket} = useAppStore()
        InitWebsocket('wss://cryptogang.vip/system/ws')
        await router.push({ name: 'hashMint',replace:true})
        window.location.reload()
      }
    }

  }
}
/* 登出*/
const LoginOut = async() => {
  jsonInBlacklist().then(async(res:any)=>{
    if (res.code === 0) {
      token.value = ''
      sessionStorage.clear()
      localStorage.clear()
      setUserInfo({})
      emitter.emit("closeAllPage");
      await router.push({ name: 'login', replace: true })
      window.location.reload()
    }
  })
}

  const setToken = (val) => {
    token.value = val
  }
  const setTheme =async(val:string)=>{
    const msg = await saveTheme({theme:val}) as any
    if(msg.code !==0){
      ElMessage.error("主题保存失败")
    }
  }

  /* 清理数据 */
  const ClearStorage = async() => {
    token.value = ''
    sessionStorage.clear()
    localStorage.clear()
  }
 
  const mode = computed(() => userInfo.value.sideMode)
  const sideMode = computed(() => {
    if (userInfo.value.sideMode === 'dark') {
      return '#191a23'
    } else if (userInfo.value.sideMode === 'light') {
      return '#fff'
    } else {
      return userInfo.value.sideMode
    }
  })
  const baseColor = computed(() => {
    if (userInfo.value.sideMode === 'dark') {
      return '#fff'
    } else if (userInfo.value.sideMode === 'light') {
      return '#191a23'
    } else {
      return userInfo.value.baseColor
    }
  })
  watch(() => token.value, () => {
    window.localStorage.setItem('token', token.value)
  })

  return {
    userInfo,
    userWallets,
    token,
    userProvider,
    setTheme,
    // NeedInit,
    // ResetUserInfo,
    ConnectAndLogin,
    LoginOut,
    GetUserInfo,
    setToken,
    setUserInfo,
    mode,
    sideMode,
    baseColor,
    // activeColor,
    loadingInstance,
    ClearStorage
  }
})