<template>
  <div>
    <div class="gd-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="路径">
          <el-input v-model="searchInfo.path" placeholder="路径" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="searchInfo.description" placeholder="描述" />
        </el-form-item>
        <el-form-item label="API组">
          <el-input v-model="searchInfo.apiGroup" placeholder="api组" />
        </el-form-item>
        <el-form-item label="请求">
          <el-select v-model="searchInfo.method" clearable placeholder="请选择">
            <el-option
              v-for="item in methodOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gd-table-box">
      <div class="gd-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog('addApi')">新增</el-button>
        <el-popover v-model="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" :disabled="!apis.length" style="margin-left: 10px;" @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column align="left" label="id" min-width="60" prop="ID" sortable="custom" />
        <el-table-column align="left" label="API路径" min-width="150" prop="path" sortable="custom" />
        <el-table-column align="left" label="API分组" min-width="150" prop="apiGroup" sortable="custom" />
        <el-table-column align="left" label="API简介" min-width="150" prop="description" sortable="custom" />
        <el-table-column align="left" label="请求" min-width="150" prop="method" sortable="custom">
          <template #default="scope">
            <div>
              {{ scope.row.method }} / {{ methodFilter(scope.row.method) }}
            </div>
          </template>
        </el-table-column>

        <el-table-column align="left" fixed="right" label="操作" width="200">
          <template #default="scope">
            <el-button
              icon="edit"
              size="small"
              type="primary"
              link
              @click="editApiFunc(scope.row)"
            >编辑</el-button>
            <el-button
              icon="delete"
              size="small"
              type="primary"
              link
              @click="deleteApiFunc(scope.row)"
            >删除</el-button>
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

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle">
      <warning-bar title="新增API，需要在角色管理内配置权限才可使用" />
      <el-form ref="apiForm" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="路径" prop="path">
          <el-input v-model="form.path" autocomplete="off" />
        </el-form-item>
        <el-form-item label="请求" prop="method">
          <el-select v-model="form.method" placeholder="请选择" style="width:100%">
            <el-option
              v-for="item in methodOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="api分组" prop="apiGroup">
          <el-input v-model="form.apiGroup" autocomplete="off" />
        </el-form-item>
        <el-form-item label="api简介" prop="description">
          <el-input v-model="form.description" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Api',
}
</script>

<script setup>

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
