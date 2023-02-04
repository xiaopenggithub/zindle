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
        <div class="gva-btn-list">
          <el-button size="small" type="primary" icon="plus" @click="openDialog('addApi')">{{ t('general.add') }}</el-button>
          <el-popover v-model="deleteVisible" placement="top" width="160">
            <p>{{ t('general.deleteConfirm') }}</p>
            <div style="text-align: right; margin-top: 8px;">
              <el-button size="small" type="primary" link @click="deleteVisible = false">{{ t('general.cancel') }}</el-button>
              <el-button size="small" type="primary" @click="onDelete">{{ t('general.confirm') }}</el-button>
            </div>
            <template #reference>
              <el-button icon="delete" size="small" :disabled="!apis.length" style="margin-left: 10px;" @click="deleteVisible = true">{{ t('general.delete') }}</el-button>
            </template>
          </el-popover>
        </div>
        <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
          <el-table-column
            type="selection"
            width="55"
          />
          <el-table-column align="left" label="ID" min-width="60" prop="id"/>
          <el-table-column align="left" label="部门名称" min-width="150" prop="name"/>
          <el-table-column align="left" label="负责人" min-width="150" prop="leader"/>
          <el-table-column align="left" label="邮箱" min-width="150" prop="email"/>
          <el-table-column align="left" label="联系电话" min-width="150" prop="phone"/>
          <el-table-column align="left" label="父级ID" min-width="150" prop="parent_id"/>
          <el-table-column align="left" label="祖级列表" min-width="150" prop="ancestors"/>
          <el-table-column align="left" label="部门状态" min-width="150" prop="description"/>
          <el-table-column label="部门状态" width="80">
                <template #default="scope">
                {{ formatBoolean(scope.row.status)}}
                </template>
            </el-table-column>
          <el-table-column align="left" label="排序" min-width="150" prop="sort"/>
          
  
          <el-table-column align="left" fixed="right" :label="t('general.operations')" width="200">
            <template #default="scope">
              <el-button
                icon="edit"
                size="small"
                type="primary"
                link
                @click="editApiFunc(scope.row)"
              >{{ t('general.edit') }}</el-button>
              <el-button
                icon="delete"
                size="small"
                type="primary"
                link
                @click="deleteApiFunc(scope.row)"
              >{{ t('general.delete') }}</el-button>
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
            <el-form-item label="父级:" prop="parent_id">
            <el-select v-model="form.parent_id" placeholder="请选择父级">
                <el-option label="顶级" :value="0"></el-option>
                <el-option
                :label="item.name"
                :value="item.id"
                v-for="(item, index) in parents"
                :key="index"
                ></el-option>
            </el-select>
            </el-form-item>
            <el-form-item label="部门名称:" prop="name">
            <el-input
                v-model="form.name"
                clearable
                placeholder="请输入部门名称"
            />
            </el-form-item>

            <el-form-item label="负责人:" prop="leader">
            <el-input
                v-model="form.leader"
                clearable
                placeholder="请输入负责人"
            />
            </el-form-item>

            <el-form-item label="联系电话:" prop="phone">
            <el-input
                v-model="form.phone"
                clearable
                placeholder="请输入联系电话"
            />
            </el-form-item>

            <el-form-item label="邮箱:" prop="email">
            <el-input
                v-model="form.email"
                clearable
                placeholder="请输入邮箱"
            />
            </el-form-item>

            <el-form-item label="部门状态:" prop="status">
            <el-radio-group v-model="form.status">
                <el-radio :label="0" name="status">正常</el-radio>
                <el-radio :label="1" name="status">停用</el-radio>
            </el-radio-group>
            </el-form-item>

            <el-form-item label="排序:" prop="sort">
            <el-input
                type="number"
                v-model.number="form.sort"
                clearable
                placeholder="请输入排序"
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
    systemDepartmentOne,
    systemDepartmentList,
    systemDepartmentAdd,
    systemDepartmentUpdate,
    systemDepartmentDelete,
    systemDepartmentDeleteBatch,
    systemDepartmentParent,
  } from "@/api/systemDepartment";

  import { toSQLLine } from '@/utils/stringFun'
  import { ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage
  
  const { t } = useI18n() // added by mohamed hassan to support multilanguage
  
  const methodFilter = (value) => {
    const target = methodOptions.value.filter(item => item.value === value)[0]
    return target && `${target.label}`
  }
  
  const parents = ref([])
  
  const apis = ref([])
  const form = ref({
    ancestors: "",
    create_by: "",
    email: "",
    id: 0,
    leader: "",
    name: "",
    parent_id: 0,
    phone: "",
    sort: 0,
    status: 0,
    update_by: "",
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
    const table = await systemDepartmentList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  
  const deleteVisible = ref(false)
  const onDelete = async() => {
    const ids = apis.value.map(item => item.id)
    const res = await systemDepartmentDeleteBatch({ ids:ids.join(",") })
    if (res.code === 200) {
      ElMessage({
        type: 'success',
        message: res.message
      })
      if (tableData.value.length === ids.length && page.value > 1) {
        page.value--
      }
      deleteVisible.value = false
      getTableData()
    }
  }
  
  // 弹窗相关
  const apiForm = ref(null)
  const initForm = () => {
    apiForm.value.resetFields()
    form.value = {
        ancestors: "",
        create_by: "",
        email: "",
        id: 0,
        leader: "",
        name: "",
        parent_id: 0,
        phone: "",
        sort: 0,
        status: 0,
        update_by: "",
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
    const res = await systemDepartmentOne({ id: row.id })
    form.value = res.data.item
    openDialog('edit')
  }
  
  const enterDialog = async() => {
    apiForm.value.validate(async valid => {
      if (valid) {
        switch (type.value) {
          case 'addApi':
            {
              const res = await systemDepartmentAdd(form.value)
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
              const res = await systemDepartmentUpdate(form.value)
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
        const res = await systemDepartmentDelete(row)
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

  const  formatBoolean=(v)=>{
    return v == 0 ? "正常" : "停用";
  }

  const loadParents=async(page, pageSize)=>{
    const res = await systemDepartmentParent({ page, pageSize })    
    if (res.code == 200) {
        parents.value = res.data.list;
    }
  }
  loadParents(1,200)
  
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
  