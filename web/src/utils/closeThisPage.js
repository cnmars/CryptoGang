import { emitter } from '@/common/bus'

export const closeThisPage = () => {
  emitter.emit('closeThisPage')
}
