import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
// import { fab } from '@fortawesome/free-brands-svg-icons'
import { App } from 'vue'

library.add(fas)
library.add(far)
// library.add(fab)
export default (app: App<Element>) => {
  app.component('font-awesome-icon', FontAwesomeIcon)
}