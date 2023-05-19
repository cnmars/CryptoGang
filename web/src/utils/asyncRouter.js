const viewModules = import.meta.glob('@/views/**/*.vue')
const pluginModules = import.meta.glob('../plugin/**/*.vue')
//把后端返回来的views/xxx转换成 ../views/xxx
export const asyncRouterHandle = (asyncRouter) => {
  asyncRouter.forEach(item => {
    if (item.component && typeof item.component==='string' ) {
      if (item.component.split('/')[0] === 'views') {
        item.component = dynamicImport(viewModules, item.component)
      } else if (item.component.split('/')[0] === 'plugin') {
        item.component = dynamicImport(pluginModules, item.component)
      }
    } else {
      delete item['component']
    }
    if (item.children) {
      asyncRouterHandle(item.children)
    }
  })
}

function dynamicImport(dynamicViewsModules,component) {
  const keys = Object.keys(viewModules)
  const matchKeys = keys.filter((key) => {
    const k = key.replace('/src/', '')
    // const k = key.replace('../', '')
    return k === component
  })
  const matchKey = matchKeys[0]

  return viewModules[matchKey]
}
