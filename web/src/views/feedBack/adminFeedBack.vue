<template>
  <div>
    <div class="gd-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="反馈类别">
          <el-select v-model="searchInfo.type" clearable placeholder="请选择">
            <el-option
              v-for="item in typeOptions"
              :key="item.value"
              :label="`${item.label}`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择">
            <el-option
              v-for="item in statusOptions"
              :key="item.value"
              :label="`${item.label}`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="反馈名称">
          <el-input v-model="searchInfo.topic" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gd-table-box">
      <div class="gd-btn-list">

        <el-popover v-model="deleteVisible" placement="top" width="160">
          <p>确定全部设置为已读吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="readVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onRead">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="SuccessFilled" type="primary" size="small" :disabled="!fbs.length" @click="readVisible = true">已读</el-button>
          </template>
        </el-popover>

        <el-popover v-model="deleteVisible" placement="top" width="160">
          <p>确定全部删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" type="danger" size="small" :disabled="!fbs.length" style="margin-left: 10px;" @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table :data="tableData" border @sort-change="sortChange" @selection-change="handleSelectionChange">
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column align="center" type="index" label="序号"  min-width="60" />
        <!-- <el-table-column align="center" label="id" min-width="60" prop="id" sortable="custom" /> -->
        <el-table-column align="center" label="创建日期" min-width="150" prop="createAt" sortable="custom" >
          <template #default="scope">
            <div>
              {{`${scope.row.createAt.slice(0,16).replace('T',' ')}`}}
            </div>
          </template>
        </el-table-column>
        <el-table-column align="center" label="反馈主题" min-width="150" prop="topic" sortable="custom" >
          <template #default="scope">
            <div>
              {{ scope.row.topic.length>=10 ? `${scope.row.topic.slice(0,10)}...` :  scope.row.topic }}
            </div>
          </template>
        </el-table-column>
        <el-table-column align="center" label="反馈类型" min-width="150" prop="type" sortable="custom" >
          <template #default="scope">
            <div>
              {{ scope.row.type==1 ? '意见 / 建议' : '系统bug' }}
            </div>
          </template>
        </el-table-column>
        <el-table-column align="center" label="提交者" min-width="150" prop="address" sortable="custom">
          <template #default="scope">
            <div>
              {{`${scope.row.address.slice(0,4)}...${scope.row.address.slice(-4)}`}}
              <el-button icon="CopyDocument" size="small" type="primary" link @click="copyAddress(scope.row)" >复制</el-button>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="center" label="反馈状态" min-width="120" prop="status" sortable="custom">
          <template #default="scope">
            <div>
              {{ scope.row.status==2 ? '未读' : '已读' }}
            </div>
          </template>
        </el-table-column>

        <el-table-column align="center" fixed="right" label="操作" width="230">
          <template #default="scope">
            <el-button icon="view" size="small" type="primary" link @click="editApiFunc(scope.row)" >查看</el-button>
            <el-popconfirm 
                confirm-button-text="Yes"
                cancel-button-text="No"
                icon="InfoFilled"
                icon-color="#126b40"
                title="Are you sure to delete this?"
                @confirm="setReadFb(scope.row)"
                @cancel="">
            <template #reference>
            <el-button v-if="(scope.row.status ==2)" icon="SuccessFilled" size="small" type="primary" link >已读</el-button>
            </template>
           </el-popconfirm>
          <el-popconfirm 
                  confirm-button-text="Yes"
                  cancel-button-text="No"
                  icon="InfoFilled"
                  icon-color="#126b40"
                  title="Are you sure to delete this?"
                  @confirm="deleteFb(scope.row)"
                  @cancel="">
              <template #reference>
                <el-button icon="DeleteFilled" size="small" type="primary" link>删除</el-button>
            </template>
          </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <div class="gd-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>

    </div>
<!-- 弹出框 -->
    <el-dialog v-model="dialogFormVisible" :title="dialogTitle">
      <warning-bar title="处理反馈后请标记为已读" />
      <div style=" align-items: center; margin-top: 20px;margin-left: 20px;">
          <el-row>
            <el-col :span="12">
                <span style="font-weight:bold;"> 反馈主题:</span>
                <span style="margin-left:20px ;">{{form.topic}}</span>
            </el-col>
            <el-col :span="12">
                <span style="font-weight:bold;"> 反馈者电话:</span>
                <span style="margin-left:20px ;">{{form.phone}}</span>
            </el-col>
          </el-row>
          <el-row style="margin-top: 20px;">
            <el-col :span="24" >
              <span style="font-weight:bold;"> 反馈内容:</span>
              <div style="margin-left:60px ;margin-top: 40px;">
                <div v-html="form.content"></div>

              </div>
            </el-col>
          </el-row>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" type="primary" @click="enterDialog">返回</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'adminFeedBack',
}
</script>

<script setup>
import {createFeedback,deleteFbById,getFbList,setRead} from '@/api/feedback'
import { toSQLLine } from '@/utils/stringFun'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const methodFilter = (value) => {
  const target = typeOptions.value.filter(item => item.value === value)[0]
  return target && `${target.label}`
}

const fbs = ref([])
const form = ref({})
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
  {
    value: 0,
    label: '全部',
    type: ''
  },
])
const statusOptions = ref([
  {
    value: 2,
    label: '未读',
    type: 'success'
  },
  {
    value: 1,
    label: '已读',
    type: ''
  }, 
  {
    value: 0,
    label: '全部',
    type: ''
  },
])

const type = ref('')
const rules = ref({
  path: [{ required: true, message: '请输入api路径', trigger: 'blur' }],
  apiGroup: [
    { required: true, message: '请输入组名称', trigger: 'blur' }
  ],
  method: [
    { required: true, message: '请选择请求方式', trigger: 'blur' }
  ],
  description: [
    { required: true, message: '请输入api介绍', trigger: 'blur' }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
}
// 搜索

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    if (prop === 'id') {
      prop = 'id'
    }
    searchInfo.value.orderKey = toSQLLine(prop)
    searchInfo.value.desc = order === 'descending'
  }
  getTableData()
}

// 查询
const getTableData = async() => {
  
  const table = await getFbList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// 批量操作
const handleSelectionChange = (val) => {
  fbs.value = val
}

const deleteVisible = ref(false)
const onDelete = async() => {
  const ids = fbs.value.map(item => item.id)
  const res = await deleteApisByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: res.msg
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

const deleteFb=async(row)=>{
  const id = row.id
  const res = await deleteFbById({id:id})
  if (res.code === 0) {
    getTableData()
    ElMessage({
      type: 'success',
      message: res.msg
    })
  }else{
    ElMessage({
      type: 'error',
      message: res.msg || '错误'
    })
  }
}
const setReadFb=async(row)=>{
  const id = row.id
  const res = await setRead({id:id})
  if (res.code === 0) {
    getTableData()
    ElMessage({
      type: 'success',
      message: res.msg
    })
  }else{
    ElMessage({
      type: 'error',
      message: res.msg || '错误'
    })
  }
}
const readVisible = ref(false)
const onRead = async() => {
  const ids = fbs.value.map(item => item.id)
  const res = await deleteApisByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: res.msg
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    readVisible.value = false
    getTableData()
  }
}

// 弹窗相关


const dialogTitle = ref('新增Api')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
  switch (key) {
    case 'addApi':
      dialogTitle.value = '新增Api'
      break
    case 'view':
      dialogTitle.value = '查看详情'
      break
    default:
      break
  }
  type.value = key
  dialogFormVisible.value = true
}

const editApiFunc = async(row) => {
  form.value = row
  openDialog('view')
}
const copyAddress=(row)=> {
     var input = document.createElement("input"); // 创建input对象
     input.value = row.address; // 设置复制内容
     document.body.appendChild(input); // 添加临时实例
     input.select(); // 选择实例内容
     document.execCommand("Copy"); // 执行复制
     document.body.removeChild(input); // 删除临时实例
     ElMessage({
          message: "复制成功",
          type: 'success'
        })
}

const enterDialog = async() => {
  form.value={}
  dialogFormVisible.value = false
}

const deleteApiFunc = async(row) => {
  ElMessageBox.confirm('此操作将永久删除所有角色下该api, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteApi(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
}

</script>

<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.warning {
  color: #dc143c;
}
</style>
