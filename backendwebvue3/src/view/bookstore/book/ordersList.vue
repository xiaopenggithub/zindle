<template>
    <div>
        <!--
        TODO:分别查询用户名/图书标题
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
        -->
      <div class="gva-table-box">
        <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">          
            <el-table-column label="ID" prop="id" width="70" />

            <el-table-column label="用户">
                <template #default="scope">
                <span>[{{scope.row.member_id}}]{{scope.row.member_name}}</span>
                </template>
            </el-table-column>
            
            <el-table-column label="书名">
                <template #default="scope">
                <span>[{{scope.row.book_id}}]{{scope.row.book_name}}</span>
                </template>
            </el-table-column>

            <el-table-column label="还书日期" width="160">
                <template #default="scope">
                {{ formatDate(scope.row.return_date)}}
                </template>
            </el-table-column>

            <el-table-column label="创建时间" width="160">
                <template #default="scope">
                {{ formatDate(scope.row.created_at) }}
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
        <warning-bar :title="t('view.api.newApiNote')" />
        <el-form ref="apiForm" :model="form" :rules="rules" label-width="120px">
          <el-form-item :label="t('view.api.path')" prop="path">
            <el-input v-model="form.path" autocomplete="off" />
          </el-form-item>
          <el-form-item :label="t('general.request')" prop="method">
            <el-select v-model="form.method" :placeholder="t('general.pleaseSelect')" style="width:100%">
              <el-option
                v-for="item in methodOptions"
                :key="item.value"
                :label="`${item.label}(${item.value})`"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item :label="t('view.api.apiGrouping')" prop="api_group">
            <el-input v-model="form.api_group" autocomplete="off" />
          </el-form-item>
          <el-form-item :label="t('view.api.apiDescrpition')" prop="description">
            <el-input v-model="form.description" autocomplete="off" />
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
  bookOrderList,
  bookOrderDelete,
  bookOrderDeleteBatch,
  bookOrderOne,
  bookOrderAdd,
  bookOrderUpdate,
} from "@/api/bookOrder";
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
    path: '',
    api_group: '',
    method: '',
    description: ''
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
    const table = await bookOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  
  const editApiFunc = async(row) => {
    const res = await systemApiOne({ id: row.id })
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
              const res = await systemApiUpdate(form.value)
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
        return "未验证";
      } else if (v == 1) {
        return "已验证";
      } else {
        return "验证错误";
      }
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
  