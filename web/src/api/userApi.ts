import service from '@/utils/request.js'
// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data:any) => {
   return{
    data:{
      user:{
        uuid: '',
      nickName: 'CONTAC',
      avatar: '',
      authority: {},
      sideMode: 'dark',
      activeColor: '#4D70FF',
      baseColor: '#fff'
      }
  },
  code:0
}
  // return service({
  //   url: '/base/login',
  //   method: 'post',
  //   data: data
  // })
}

// @Summary 获取验证码
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/captcha [post]
export const captcha = (data:any) => {
  return service({
    url: '/base/captcha',
    method: 'post',
    data: data
  })
}
export const getSignMsg = (data:any) => {
  return service({
    url: '/system/base/getNonce',
    method: 'post',
    data: data
  })
}
export const signAndLogin = (data:any) => {
  return service({
    url: '/system/base/login',
    method: 'post',
    data: data
  })
}
export const setProviderUrl = (data:any) => {
  return service({
    url: '/system/user/setProviderUrl',
    method: 'post',
    data: data
  })
}
export const saveTheme = (data:any) => {
  return service({
    url: '/system/user/setTheme',
    method: 'post',
    data: data
  })
}
export const setProviderDefault = (data:any) => {
  return service({
    url: '/system/user/setProviderDefault',
    method: 'post',
    data: data
  })
}

// @Summary 用户注册
export const register = (data:any) => {
  return service({
    url: '/user/admin_register',
    method: 'post',
    data: data
  })
}

export const changePassword = (data:any) => {
  return service({
    url: '/user/changePassword',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary 分页获取用户列表
export const getUserList = (data:any) => {
  return service({
    url: '/user/getUserList',
    method: 'post',
    data: data
  })
}
export const upgrade = (data:any) => {
  return service({
    url: '/system/user/upgrade',
    method: 'post',
    data: data
  })
}

// @Summary 设置用户权限
export const setUserAuthority = (data:any) => {
  return service({
    url: '/user/setUserAuthority',
    method: 'post',
    data: data
  })
}

export const deleteUser = (data:any) => {
  return service({
    url: '/user/deleteUser',
    method: 'delete',
    data: data
  })
}

export const setUserInfo = (data:any) => {
  return service({
    url: '/user/setUserInfo',
    method: 'put',
    data: data
  })
}

// @Tags SysUser
// @Summary 设置用户信息
export const setSelfInfo = (data:any) => {
  return service({
    url: '/user/setSelfInfo',
    method: 'put',
    data: data
  })
}

// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.setUserAuthorities true "设置用户权限"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthorities [post]
export const setUserAuthorities = (data:any) => {
  return service({
    url: '/user/setUserAuthorities',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserInfo [get]
export const getUserInfoFromServer = (data:any) => {
  return service({
    url: '/system/user/getUserInfo',
    method: 'get',
    data:data
  })
}

export const resetPassword = (data:any) => {
  return service({
    url: '/user/resetPassword',
    method: 'post',
    data: data
  })
}
