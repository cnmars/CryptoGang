<template>
    <div>
        <div class="gd-search-box" style="text-align: center;">
            <el-form :inline="true" :model="form" :rules="rules" ref="feedbackRef">
                <el-form-item :label="$t('ufeedback.label.type')" prop="type" style="width: 28%;">
                <el-select v-model="form.type" clearable :placeholder="$t('ufeedback.ph.type')">
                    <el-option
                    v-for="item in typeOptions"
                    :key="item.value"
                    :label="`${item.label}`"
                    :value="item.value"
                    />
                </el-select>
                </el-form-item>
                <el-form-item :label="$t('ufeedback.label.topic')" prop="topic" style="width: 28%;">
                    <el-input v-model="form.topic" :placeholder="$t('ufeedback.ph.topic')"></el-input>
                </el-form-item>
                <el-form-item :label="$t('ufeedback.label.phone')" prop="phone" style="width: 28%;">
                    <el-input v-model="form.phone" :placeholder="$t('ufeedback.ph.phone')"></el-input>
                </el-form-item>
            </el-form>
        </div>
         <!-- <div class="gd-table-box"> -->
        <div class="gd-table-box">
            <el-row style="width:100%" >
                <el-col :span="24" style="">

                    <wangEditor v-if="reloadComponentTag" />
                    <!-- <div  style="border: 1px solid rgb(144, 144, 144); width:90%;margin-left: 5%;">asasas</div>
                    <div  style="border: 1px solid rgb(144, 144, 144); width:90%;margin-left: 5%;">
                        <Toolbar
                            style="border-bottom: 1px solid #126b40"
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
                            @onChange="changeContent"
                        />
                    </div> -->
                </el-col>
            </el-row>

             <el-row style="width:100%;align-items: center;" >
                <el-col :span="24" style="align-items: center;text-align: center;margin-bottom: 18px;">
                    <!-- <el-form-item> -->
                        <el-button type="primary" class="submit-btn" style="margin-top: 10px;border-radius: 20px;height: 35px;" @click="submitForm"  :circle="true">
                            <template #icon>
                                <font-awesome-icon icon="fa-solid fa-paper-plane"/>
                            </template>
                            {{$t('ufeedback.btn.submit')}}
                        </el-button>
                    <!-- </el-form-item>  -->
                </el-col>
             </el-row>
        </div>
    </div>
  </template>
<script lang="ts">
  export default{
     name:"userFeedback"
  }
</script>
<script lang="ts" setup>
/* 参考：https://segmentfault.com/a/1190000042722618 */
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import { ref,reactive,onUnmounted ,h} from 'vue'
import { onBeforeUnmount, shallowRef} from 'vue'
import { ElMessage, ElNotification } from 'element-plus'
import  wangEditor  from "@/components/wangEditor/wangEditor.vue";
import { useUserStore } from '@/pinia/modules/userStore'
import { emitter } from '@/common/bus'
import {createFeedback} from '@/api/feedback'
import { isEmpty } from 'lodash';
import { ElLoading } from 'element-plus'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const userStore = useUserStore()
const reloadComponentTag =ref(true)
const typeOptions = ref([
  {
    value: 1,
    label: '意见 / 建议',
    type: 'success'
  },
  {
    value: 2,
    label: '系统bug',
    type: ''
  },
])
const isVlidNumber = ((rule,value,callback)=>{
    //不输入手机号也可以
    if(!value){
        return true
    }
    const reg = /^1(3\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\d|9[0-35-9])\d{8}$/
    if(reg.test(value)){
        return true
    }
    callback(new Error('请输入正确的手机号'))
    })
 const form = ref({
    type:'',
    phone: "",
    topic: "",
    content: "",
    address: ''
    })
const content= ref()
const rules = reactive({
    type:  [{ required: true, message: "请选择类型", trigger: "blur" }],
    phone: [{ required: false, message: "请输入正确的手机号码", trigger: "blur" },{validator:isVlidNumber,trigger:'blur'}],
    topic: [{ required: true, message: "请输入反馈名称", trigger: "blur" }],
})

const feedbackRef= ref()
const reloadThisComponnet = () => {
    reloadComponentTag.value = false;
    editorConfig.placeholder = t('ufeedback.ph.txt')
      setTimeout(() => {
        reloadComponentTag.value = true;
      }, 0)
    }
const submitForm= ()=> {
    feedbackRef.value.validate(async(valid:boolean) => {
        if (valid) {
        // 校验通过，可以进行提交操作
            if(isEmpty(content.value) || content.value=='<p><br></p>'){
                ElMessage.error('请录入您想要反馈的内容！')
                return
            }else{
                form.value.content = content.value
                form.value.address = userStore.userInfo.address
                const loadingInstance1 = ElLoading.service({ fullscreen: true })
                const res = await createFeedback(form.value)as any
                if (res.code === 0) {
                    ElNotification({
                        title: 'Success',
                        message: h('div', { style: 'color: green' }, '提交成功'),
                    })
                }
                form.value = {
                    type:'',
                    phone: "",
                    topic: "",
                    content: "",
                    address: ''
                }
                reloadThisComponnet()
                loadingInstance1.close()
                // window.location.reload()
            }
        } else {
        // 校验失败，提示错误信息
            ElMessage.error('请按规则录入信息！')
            return
        }
    });
    }
emitter.on('userFeddbackChange',((data)=>{
    content.value= data
    // console.log(content.value);
}))
emitter.on('changeEditorLang',((data)=>{
    reloadThisComponnet()
}))
onUnmounted(() => {
    emitter.off('userFeddbackChange')
    emitter.off('changeEditorLang')
})


// created (){


// }

 // 编辑器实例，必须用 shallowRef
 const editorRef = shallowRef()
    // 内容 HTML
    const valueHtml = ref('') //  工具栏配置 

    const toolbarConfig = {excludeKeys: [
     'fullScreen',//不显示全屏
     'group-video', // 不显示上传视频
     'insertLink',
     'todo',
     'uploadImage'
   ]}
//菜单配置
    const editorConfig = {
        placeholder: t('ufeedback.ph.txt'),
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
    const changeContent = ((data)=>{
         emitter.emit('userFeddbackChange',valueHtml.value)
    })
   
    // 组件销毁时，也及时销毁编辑器
    onBeforeUnmount(() => {
        const editor = editorRef.value
        if (editor == null) return
        editor.destroy()
        // emitter.off('changeEditorLang')
    })

    const handleCreated = (editor) => {
      editorRef.value = editor // 记录 editor 实例，重要！

        // 查看所有工具栏key，先查看可以根据实际情况进项增删
        // console.log(editor.getAllMenuKeys())

    }
    const mode ='default' // 或 'simple'

</script>

<style lang="scss" scoped>
.feedback-warp{
    width: 99.9%;
    height: auto;
    box-shadow: var(--el-box-shadow-light);
    background-color: #fff;
    align-items: center;
}

.feedback-form-wrap{
    // border:#909090 1px solid;
    // margin-top: 40px;
    width: 88%;
    padding-top: 40px;
    padding-left: 6%;
    padding-right: 40px;
    .el_input_inner{
        width: 100%;
    }
    .el-form-item__label {
      color: #000000;
      font-size: 15px;
    }
}

.submit-btn{
    border-radius: 20px;height: 35px;background-color: #126b40;
    border: 1px #126b40 solid;
    width: 135px;
} 

</style>