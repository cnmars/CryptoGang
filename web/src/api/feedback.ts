import service from '@/utils/request.js'

export const createFeedback = (data:any) => {
    return service({
      url: '/system/feedback/createFb',
      method: 'post',
      data: data
    })
  }
  export const deleteFbById = (data:any) => {
    return service({
      url: '/system/feedback/deleteFb',
      method: 'post',
      data: data
    })
  }
  export const setRead = (data:any) => {
    return service({
      url: '/system/feedback/setRead',
      method: 'post',
      data: data
    })
  } 
  export const getFbList = (data:any) => {
    return service({
      url: '/system/feedback/getFbList',
      method: 'post',
      data: data
    })
  }