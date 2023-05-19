import { createApp, ref } from 'vue'
import 'element-plus/dist/index.css'
import './style/element_visiable.scss'
import './style/gd_global.scss'
import './style/dark.scss'

import App from '@/App.vue'
import router from '@/router/router'
import { store } from '@/pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import { initDom } from '@/utils/positionToCode.js'
import svgIcon from "@/components/svgIcon/index.vue";
import { useUserStore } from '@/pinia/modules/userStore'
import { ElNotification } from 'element-plus'
import { h} from 'vue'
import installFontawesome from '@/plugins/fontAwesome'
import '@/permission'
import 'virtual:svg-icons-register';
import i18n from '@/lang/index';
import { ethUtils } from "@/utils/blockchain";
initDom()
/**
 * @description 导入加载进度条，防止首屏加载时间过长，用户等待
 *
 * */
const app = createApp(App)
//引入图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
installFontawesome(app)
app.use(router)
    .use(ElementPlus, { locale: zhCn })
    .use(store)
    .use(i18n)
    .mount('#app')
app.component('svg-icon', svgIcon);

(window as any).ethereum.on('networkChanged', function (networkId) {
  // Time to reload your interface with the new networkId
  const userStore = useUserStore()
  let networkName = ethUtils.getNetWorkNameById(networkId)
  userStore.userInfo.showNetwork = networkName
  if(networkId!=='1'){
    ElNotification({
      title: 'Network',
      message: h('div', { style: 'color: red' }, '当小狐狸网络不是以太主网时，只能用私钥进行交易，而不能采用小狐狸交易。(当前只支持以太主网)'),
    })
  }
});
(window as any).ethereum.on('accountsChanged', function (accounts) {
  // Time to reload your interface with the new networkId
  const userStore = useUserStore()
  if(userStore.userInfo.address !==accounts[0]){
    ElNotification({
      title: '切换账号',
      message: h('div', { style: 'color: red' }, '切换账号后原账号任保留登录状态，但是此后调用小狐狸进行的交易会在新账号上进行'),
    })
  }
})
