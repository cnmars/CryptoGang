import axios from 'axios' // 引入axios
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserStore } from '@/pinia/modules/userStore'
import { emitter } from '@/common/bus'
import router from '@/router/router'
import image from "@/assets/img/contac.jpg";
const service = axios.create({
  baseURL: import.meta.env.VITE_BASE_API,
  timeout: 99999
})
let acitveAxios = 0
let timer
const showLoading = () => {
  acitveAxios++
  if (timer) {
    clearTimeout(timer)
  }
  timer = setTimeout(() => {
    if (acitveAxios > 0) {
      emitter.emit('showLoading')
    }
  }, 200)
}

const closeLoading = () => {
  acitveAxios--
  if (acitveAxios <= 0) {
    clearTimeout(timer)
    emitter.emit('closeLoading')
  }
}
// http request 拦截器
service.interceptors.request.use(
  config => {
    if (!config.donNotShowLoading) {
      showLoading()
    }
    const userStore = useUserStore()
    config.headers = {
      'Content-Type': 'application/json',
      'x-token': userStore.token,
      'x-user-id': userStore.userInfo.ID,
      ...config.headers
    }
    return config
  },
  error => {
    closeLoading()
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)

// http response 拦截器
service.interceptors.response.use(
  response => {
    const userStore = useUserStore()
    closeLoading()
    if (response.headers['new-token']) {
      userStore.setToken(response.headers['new-token'])
    }
    if (response.data.code === 0 || response.headers.success === 'true') {
      if (response.headers.msg) {
        response.data.msg = decodeURI(response.headers.msg)
      }if(response.data.msg){
        // ElMessage({
        //   message: response.data.msg,
        //   type: 'success'
        // })
      }
      return response.data
    } else {
      ElMessage({
        message: response.data.msg || decodeURI(response.headers.msg),
        type: 'error'
      })
      if (response.data.data && response.data.data.reload) {
        userStore.token = ''
        localStorage.clear()
        //登录过期，这里应跳转到首页
        router.push({ name: 'login', replace: true })
      }
      return response.data.msg ? response.data : response
    }
  },
  error => {
    closeLoading()
    if (!error.response) {
      ElMessageBox.confirm(`
        
        <div style="flex-direction:row;display:flex;justify-content:space-between;gap:5px;">
          <div>
          <p>检测到请求错误</p>
          <p>${error}</p>
            <p>请通知我处理</p>
          </div>
          <div style="display:flex;justify-content:flex-end;">
            <img src="${image}" alt="" style="width:60px;height:60px;">
          </div>
        </div>
        `, '请求报错', {
        dangerouslyUseHTMLString: true,
        distinguishCancelAndClose: true,
        confirmButtonText: '重试',
        cancelButtonText: '取消'
      })
      return
    }
    
    switch (error.response.status) {
      case 500:
        /* <p>${error}</p>
        <p>错误码<span style="color:red"> 500 </span>：请通知我处理</p> */
        ElMessageBox.confirm(`
        <div style="flex-direction:row;display:flex;justify-content:space-between;gap:5px;">
          <div>
            <p>服务器未开启，即将推出新版本，敬请期待</p>
              <p>如有紧急事情，请联系我</p>
          </div>
          <div style="display:flex;justify-content:flex-end;">
            <img src="${image}" alt="" style="width:60px;height:60px;">
          </div>
        </div>
        `, '系统错误', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: '清理缓存',
          cancelButtonText: '取消'
        })
          .then(() => {
            const userStore = useUserStore()
            userStore.token = ''
            localStorage.clear()
            router.push({ name: 'login', replace: true })
          })
        break
      case 404:
        ElMessageBox.confirm(`
        <div style="flex-direction:row;display:flex;justify-content:space-between;gap:5px;">
        <div>
          <p>${error}</p>
          <p>错误码<span style="color:red"> 404 </span>：请通知我处理</p>
        </div>
        <div style="display:flex;justify-content:flex-end;">
          <img src="${image}" alt="" style="width:60px;height:60px;">
        </div>
        </div>
          `, '接口报错', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: '我知道了',
          cancelButtonText: '取消'
        })
        break
      case 502:
        /* <p>${error}</p>
        <p>错误码<span style="color:red"> 502 </span>：请通知我处理</p> */
        ElMessageBox.confirm(`
          <div style="flex-direction:row;display:flex;justify-content:space-between;gap:5px;">
            <div>
              <p>服务器未开启，即将推出新版本，敬请期待</p>
              <p>如有紧急事情，请联系我</p>
            </div>
            <div style="display:flex;justify-content:flex-end;">
              <img src="${image}" alt="" style="width:60px;height:60px;">
            </div>
          </div>
          `, '服务器错误', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: '我知道了',
          cancelButtonText: '取消'
        })
        break
    }

    return error
  }
)
export default service
