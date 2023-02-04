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
        <el-button size="small" type="primary" icon="plus" @click="openDialog('add')">{{ t('general.add') }}</el-button>
        <el-popover v-model="deleteVisible" placement="top" width="160">
          <p>{{ t('general.deleteConfirm') }}</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">{{ t('general.cancel') }}</el-button>
            <el-button size="small" type="primary" @click="deleteBatchFunc">{{ t('general.confirm') }}</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" :disabled="!selectedData.length" style="margin-left: 10px;" @click="deleteVisible = true">{{ t('general.delete') }}</el-button>
          </template>
        </el-popover>
      </div>
      <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column align="left" label="ID" min-width="60" prop="id"/>
        <el-table-column align="left" label="活动名称" min-width="150" prop="title"/>
        <el-table-column align="left" label="活动地址" min-width="150" prop="address"/>
        <el-table-column label="开始时间" width="160">
          <template #default="scope">
            {{ formatDate(scope.row.time_start) }}
          </template>
        </el-table-column>
        <el-table-column label="结束时间" width="160">
          <template #default="scope">
            {{ formatDate(scope.row.time_end) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="简介" min-width="150" prop="description"/>
        <el-table-column align="left" label="数量" min-width="150" prop="quantity"/>
        <el-table-column align="left" label="封面" min-width="150" prop="cover"/>  
        <el-table-column align="left" label="排序" min-width="150" prop="sort"/>
        <el-table-column align="left" fixed="right" :label="t('general.operations')" width="200">
          <template #default="scope">
            <el-button
              icon="edit"
              size="small"
              type="primary"
              link
              @click="editFunc(scope.row)"
            >{{ t('general.edit') }}</el-button>
            <el-button
              icon="delete"
              size="small"
              type="primary"
              link
              @click="deleteFunc(scope.row)"
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
      <el-form ref="editForm" :model="form" :rules="rules" label-width="120px">                    
          <el-form-item label="活动名称:" prop="title">
          <el-input
              v-model="form.title"
              clearable
              placeholder="请输入活动名称"
          />
          </el-form-item>

          <el-form-item label="活动地址:" prop="address">
          <el-input
              v-model="form.address"
              clearable
              placeholder="请输入活动地址"
          />
          </el-form-item>

          <el-form-item label="简介:" prop="description">
          <el-input
              v-model="form.description"
              clearable
              placeholder="请输入简介"
          />
          </el-form-item>

          <el-form-item label="封面:" prop="cover">
          <el-input
              v-model="form.cover"
              clearable
              placeholder="请输入封面"
          />
          </el-form-item>


          <el-form-item label="开始时间:" prop="time_start">          
            <el-date-picker
              v-model="form.time_start"
              type="datetime"
              placeholder="选择开始时间"
              align="right"
              format="YYYY-MM-DD HH:mm:ss"            
            >
            </el-date-picker>
          </el-form-item>

          <el-form-item label="结束时间:" prop="time_end">
            <el-date-picker
                v-model="form.time_end"
                type="datetime"
                placeholder="请输入结束时间"
                align="right"
                format="YYYY-MM-DD HH:mm:ss"            
              >
            </el-date-picker>
          </el-form-item>
          
          <el-form-item label="数量:" prop="quantity">
          <el-input
              type="number"
              v-model.number="form.quantity"
              clearable
              placeholder="请输入数量"
          />
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
  import { ref } from 'vue'  
  import {
    activityList,
    activityDelete,
    activityDeleteBatch,
    activityOne,
    activityAdd,
    activityUpdate,
  } from "@/api/activity";
  import moment from "moment"; //导入模块
  

  import { ElMessage, ElMessageBox } from 'element-plus'
  
  import { useI18n } from 'vue-i18n'
  const { t } = useI18n()


  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({})

  const actionType = ref('')
  
  const selectedData = ref([])
  const form = ref({
    id: 0,
    create_by: "",
    update_by: "",
    sort: 0,
    cover: "",
    quantity: 0,
    description: "",        
    title: "",
    address: "",
    time_start: 0,
    time_end: 0,  
  })


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
    const table = await activityList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 200) {
      tableData.value = table.data.list
      total.value = table.data.total
      // page.value = table.data.page
      // pageSize.value = table.data.pageSize
    }
  }
  // 批量操作
  const handleSelectionChange = (val) => {
    selectedData.value = val
  }
  getTableData()

  const deleteVisible = ref(false)
  const deleteBatchFunc = async() => {
    const ids = selectedData.value.map(item => item.id)
    const res = await activityDeleteBatch({ ids:ids.join(",") })
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
  const editForm = ref(null)
  const initForm = () => {
    editForm.value.resetFields()
    form.value = {        
      id: 0,
      create_by: "",
      update_by: "",
      sort: 0,
      cover: "",
      quantity: 0,
      description: "",        
      title: "",
      address: "",
      time_start: 0,
      time_end: 0,        
    }
  }

  const dialogTitle = ref('新增Api')
  const dialogFormVisible = ref(false)
  const openDialog = (key) => {
    switch (key) {
      case 'add':
        dialogTitle.value = '新增'
        break
      case 'edit':
        dialogTitle.value = '编辑'
        break
      default:
        break
    }
    actionType.value = key
    dialogFormVisible.value = true
  }
  const closeDialog = () => {
    initForm()
    dialogFormVisible.value = false
  }
  const editFunc = async(row) => {    
    const res = await activityOne({ id: row.id })
    console.log(res)    
    form.value = res.data
    openDialog('edit')
  }
  
  const formatDate=(time)=>{
    if (time != null && time != "" && time>100000000) {
      return moment(time).format("yyyy-MM-DD HH:mm:ss")
    }
    return time;
    
  }

  const enterDialog = async() => {
    editForm.value.validate(async valid => {
      if (valid) {
        if(form.value.time_start){
          form.value.time_start=parseInt(moment(form.value.time_start).format('x'))
        }
        if(form.value.time_end){
          form.value.time_end=parseInt(moment(form.value.time_end).format('x'))
        }
        switch (actionType.value) {
          case 'add':
            {
              const res = await activityAdd(form.value)
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
              const res = await activityUpdate(form.value)
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
  
  const deleteFunc = async(row) => {
    ElMessageBox.confirm(t('view.api.deleteApiConfirm'), t('general.hint'), {
      confirmButtonText: t('general.confirm'),
      cancelButtonText: t('general.cancel'),
      type: 'warning'
    })
      .then(async() => {
        const res = await activityDelete(row)
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
</script>