import { fmtTitle } from '@/utils/fmtRouterTitle'
import {appConfig} from '../global/config.ts'
export default function getPageTitle(pageTitle, route) {
  if (pageTitle) {
    const title = fmtTitle(pageTitle, route)
    return `${title} - ${appConfig.appName}`
  }
  return `${appConfig.appName}`
}
