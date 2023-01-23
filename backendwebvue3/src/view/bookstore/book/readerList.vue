<template>
    <div>
      <div class="gva-search-box">
        <el-form ref="searchForm" :inline="true" :model="searchInfo">
          <el-form-item label="搜索关键词">
            <el-input v-model="searchInfo.keyword" placeholder="输入搜索关键词" clearable/>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary" icon="search" @click="onSubmit">{{ t('general.search') }}</el-button>
            <el-button size="small" icon="refresh" @click="onReset">{{ t('general.reset') }}</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="gva-table-box">
        <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">          
            <el-table-column label="ID" prop="id" width="70" />

            <el-table-column label="用户账号" prop="username" />
            <el-table-column label="手机" prop="phone" />
            <el-table-column label="用户邮箱" prop="email" />
            <el-table-column label="帐号状态" width="80" align="center">
                <template #default="scope">
                {{ formatStatus(scope.row.status)}}
                </template>
            </el-table-column>
            <el-table-column label="最后登录IP" prop="login_ip" />

            <el-table-column label="最后登录" width="160">
                <template #default="scope">
                {{ formatDate(scope.row.login_date)}}
                </template>
            </el-table-column>

            <el-table-column label="备注" prop="remark" />

            <el-table-column label="创建时间" width="160">
                <template #default="scope">
                {{ formatDate(scope.row.created_at)}}
                </template>
            </el-table-column>

            <el-table-column label="操作" fixed="right" width="180" align="center">
                <template #default="scope">
                    <el-button
                        icon="edit"
                        size="small"
                        type="primary"
                        link
                        @click="editFunc(scope.row)">{{ t('general.edit') }}</el-button>
                </template>
            </el-table-column>
        </el-table>
        <div class="gva-pagination">
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
        <el-form ref="apiForm" :model="form" :rules="rules" label-width="120px">
            <el-form-item label="用户账号:" prop="username">
            <el-input
                v-model="form.username"
                :disabled="true"
                clearable
                placeholder="请输入用户账号"
            />
            </el-form-item>

            <el-form-item label="密码:" prop="password">
            <el-input
                v-model="form.password"
                clearable
                placeholder="请输入密码"
                :disabled="true"
            />
            </el-form-item>

            <el-form-item label="手机:" prop="phone">
            <el-input
                v-model="form.phone"
                clearable
                placeholder="请输入手机"
                :disabled="true"
            />
            </el-form-item>

            <el-form-item label="用户邮箱:" prop="email">
            <el-input
                v-model="form.email"
                clearable
                placeholder="请输入用户邮箱"
                :disabled="true"
            />
            </el-form-item>

            <el-form-item label="帐号状态:" prop="status">
            <el-radio-group v-model="form.status">
                <el-radio :label="0" name="status">正常</el-radio>
                <el-radio :label="1" name="status">停用</el-radio>
            </el-radio-group>
            </el-form-item>

            <el-form-item label="备注:" prop="remark">
            <el-input
                v-model="form.remark"
                clearable
                placeholder="请输入备注"
            />
            </el-form-item>
        </el-form>
        <template #footer>
          <div class="dialog-footer">
            <el-button size="small" @click="closeDialog">{{ t('general.close') }}</el-button>
            <el-button size="small" type="primary" @click="enterDialog">{{ t('general.confirm') }}</el-button>
          </div>
        </template>
      </el-dialog>
    </div>
  </template>
  
  <script setup>
  import {
    systemApiOne,
    systemApiList,
    systemApiAdd,
    systemApiUpdate,
    systemApiDelete,
    systemApiDeleteBatch,  
  } from "@/api/systemApi";

import {
  readerList,
  readerDelete,
  readerDeleteBatch,
  readerOne,
  readerAdd,
  readerUpdate,
} from "@/api/reader";
import { formatTimeToStr } from "@/utils/date";

  import { toSQLLine } from '@/utils/stringFun'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage
  
  const { t } = useI18n() // added by mohamed hassan to support multilanguage
  
  const methodFilter = (value) => {
    const target = methodOptions.value.filter(item => item.value === value)[0]
    return target && `${target.label}`
  }
  
  const apis = ref([])
  const form = ref({
    id: 0,
    username: "",
    password: "",
    nickname: "",
    phone: "",
    email: "",
    status: 0,
    login_ip: "",
    login_date: "",
    remark: "",
  })
  const methodOptions = ref([
    {
      value: 'POST',
      label: t('view.api.create'),
      type: 'success'
    },
    {
      value: 'GET',
      label: t('view.api.view'),
      type: ''
    },
    {
      value: 'PUT',
      label: t('view.api.update'),
      type: 'warning'
    },
    {
      value: 'DELETE',
      label: t('general.delete'),
      type: 'danger'
    }
  ])
  
  const type = ref('')
  const rules = ref({
    path: [{ required: true, message: t('view.api.enterApiPath'), trigger: 'blur' }],
    api_group: [
      { required: true, message: t('view.api.enterGroupName'), trigger: 'blur' }
    ],
    method: [
      { required: true, message: t('view.api.selectRequestMethod'), trigger: 'blur' }
    ],
    description: [
      { required: true, message: t('view.api.enterApiDescription'), trigger: 'blur' }
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
      if (prop === 'ID') {
        prop = 'id'
      }
      searchInfo.value.orderKey = toSQLLine(prop)
      searchInfo.value.desc = order === 'descending'
    }
    getTableData()
  }
  
  // 查询
  const getTableData = async() => {    
    const table = await readerList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 200) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }
  
  getTableData()
  
  // 批量操作
  const handleSelectionChange = (val) => {
    apis.value = val
  }
  
  
  
  // 弹窗相关
  const apiForm = ref(null)
  const initForm = () => {
    apiForm.value.resetFields()
    form.value = {
      path: '',
      apiGroup: '',
      method: '',
      description: ''
    }
  }
  
  const dialogTitle = ref('新增Api')
  const dialogFormVisible = ref(false)
  const openDialog = (key) => {
    switch (key) {
      case 'addApi':
        dialogTitle.value = t('view.api.newApi')
        break
      case 'edit':
        dialogTitle.value = t('view.api.editApi')
        break
      default:
        break
    }
    type.value = key
    dialogFormVisible.value = true
  }
  const closeDialog = () => {
    initForm()
    dialogFormVisible.value = false
  }
  
  const editFunc = async(row) => {
    const res = await readerOne({ id: row.id })
    form.value = res.data.item
    openDialog('edit')
  }
  
  const enterDialog = async() => {
    apiForm.value.validate(async valid => {
      if (valid) {
        switch (type.value) {
          case 'addApi':
            {
              const res = await systemApiAdd(form.value)
              if (res.code === 200) {
                ElMessage({
                  type: 'success',
                  message: t('general.addSuccess'),
                  showClose: true
                })
              }
              getTableData()
              closeDialog()
            }  
            break
          case 'edit':
            {
              form.value.login_date=formatDate(form.value.login_date)
              const res = await readerUpdate(form.value)
              if (res.code === 200) {
                ElMessage({
                  type: 'success',
                  message: t('general.editSuccess'),
                  showClose: true
                })
              }
              getTableData()
              closeDialog()
            }
            break
          default:
            // eslint-disable-next-line no-lone-blocks
            {
              ElMessage({
                type: 'error',
                message: t('view.api.unknownOperation'),
                showClose: true
              })
            }
            break
        }
      }
    })
  }
  
  const deleteApiFunc = async(row) => {
    ElMessageBox.confirm(t('view.api.deleteApiConfirm'), t('general.hint'), {
      confirmButtonText: t('general.confirm'),
      cancelButtonText: t('general.cancel'),
      type: 'warning'
    })
      .then(async() => {
        const res = await systemApiDelete(row)
        if (res.code === 200) {
          ElMessage({
            type: 'success',
            message: t('general.deleteSuccess')
          })
          if (tableData.value.length === 1 && page.value > 1) {
            page.value--
          }
          getTableData()
        }
      })
  }

    const formatDate=(time)=>{
        if (time != null && time != "") {
            var date = new Date(time);
            return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
        } else {
            return "";
        }
    }

    const formatType=(v)=>{
      if (v == 0) {
        return "手机";
      }
      return "邮箱";
    }
    const formatStatus=(v)=>{
        if (v == 0) {
            return "正常";
        }
        return "停用";
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
  