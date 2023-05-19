<template>
    <div style="border: 1px solid rgb(144, 144, 144); width:90%;margin-left: 5%;">
      <Toolbar
        style="border-bottom: 1px solid #ccc"
        :editor="editorRef"
        :defaultConfig="toolbarConfig"
        :mode="mode"
      />
      <Editor
      style="height: 680px; overflow-y: hidden;max-height: 680;"
        v-model="valueHtml"
        :defaultConfig="editorConfig"
        :mode="mode"
        @onCreated="handleCreated"
        @onChange="handleChange"
      />
    </div>
  </template>
  
  <script>
  export default{
      name:'wangEditor'
  }
  </script>
  <script setup>
  import '@wangeditor/editor/dist/css/style.css' // 引入 css
  import { onBeforeUnmount, ref, shallowRef, onMounted, nextTick, watch } from 'vue'
  import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
  import { ElMessage} from 'element-plus'
  import { emitter } from '@/common/bus'
  import { useI18n } from 'vue-i18n'
  
  const { t } = useI18n()
  /**
    initValue: 父组件传递过来的富文本框初始值，这里会实时监听，更新初始值，放置在弹窗中使用，没有钩子函数的更新。
    getEditorContent() 方法，父组件通过这个方法获取富文本编辑器的内容，包括数组格式的和html格式的内容
  */
  // 初始值
  const props = defineProps({
    initValue: String
  })
  const emits = defineEmits(['getEditorContent'])
  // const emits = defineEmits([''])
  let mode = ref('default')
  // 编辑器实例，必须用 shallowRef
  const editorRef = shallowRef()
  // 内容 HTML
  const valueHtml = ref('')
  // 模拟 ajax 异步获取内容
  onMounted(() => {
    nextTick(() => { // 界面节点更新完以后再修改值。
      valueHtml.value = props.initValue
    })
    emitter.on('changeEditorLang',(()=>{
              editorConfig.placeholder = t('ufeedback.ph.txt')
          }))
  })
  // 工具栏配置
  const toolbarConfig = {
      excludeKeys: [
       'fullScreen',//不显示全屏
       'group-video', // 不显示上传视频
       'insertLink',
       'todo',
       'uploadImage'
     ]
  }
  const editorConfig = {
      placeholder: t('ufeedback.ph.txt'),
    // readOnly: true,
    MENU_CONF:{
          // 配置默认字号
          // 配置上传图片
          uploadImage:{//现在没有使用这个，只能根据链接上传网络图片
                  // 上传图片请求接口路径
                  server: "/api/v1/upload/file", 
                  // 后端接收的文件名称
                  fieldName: "multipartFile",
                  maxFileSize: 4 * 1024 * 1024, // 上传图片4M
                  // 上传的图片类型
                  allowedFileTypes: ["image/*"],
                  // 小于该值就插入 base64 格式（而不上传），默认为 0
                  base64LimitSize: 4 * 1024, // 10MB
                  // 自定义上传图片返回格式【后端返回的格式】
                  customInsert(res, insertFn) {
                  if(res.code != 200 ){
                      ElMessage.error("上传文件失败，"+res.message)
                      return
                  }
                  // 从 res 中找到 url alt href ，然后插入图片 ,根据后端实际返回的字段来
                  insertFn(res.data.url, res.data.alt,res.data.url)
                  },
                  
                  // 将 meta 拼接到 url 参数中，默认 false
                  metaWithUrl: true,
                  // 单个文件上传成功之后
                  onSuccess(file, res) {
                  if(res.code == 200){
                      ElMessage.success(`${file.name} 上传成功`)
                      return
                  }else {
                      ElMessage.warning(`${file.name} 上传出了点异常`)
                      return
                  }
                  // console.log(`${file.name} 上传成功`, res)
                  //ElMessage.success(`${file.name} 上传成功`, res)
                  },
                  // 单个文件上传失败
                  onFailed(file, res) {
                      console.log(res)
                      ElMessage.error(`${file.name} 上传失败`)
                  },
                  // 上传错误，或者触发 timeout 超时
                  onError(file, err, res) {
                      console.log(err, res)
                      ElMessage.error(`${file.name} 上传出错`)
                  },
              },
          },
  }
  // 组件销毁时，也及时销毁编辑器
  onBeforeUnmount(() => {
      const editor = editorRef.value
      if (editor == null) return
      editor.destroy()
  })
  const handleCreated = (editor) => {
    editorRef.value = editor // 创建富文本编辑器
  }
  const handleChange = (info) => {
    // info.children 数组包含了数据类型和内容，valueHtml.value内容的html格式
  //    emitter.emit('userFeddbackChange',valueHtml.value)
  emitter.emit('userFeddbackChange',valueHtml.value)
  }
  watch(()=>props.initValue, (value) => {  // 父组件获取初始值
    valueHtml.value = value
  })
  </script>