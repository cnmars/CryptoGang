import service from '@/utils/request'

export const getUpgradePrice = () => {
  return service({
    url: '/system/system/getUpgradePrice',
    method: 'post',
    donNotShowLoading: true,
  }as any)
}
export const getEthPrice = () => {
  return service({
    url: '/system/system/getEthPrice',
    method: 'post'
  })
}