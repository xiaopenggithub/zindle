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
              <el-button size="small" type="primary" @click="onDelete">{{ t('general.confirm') }}</el-button>
            </div>
            <template #reference>
              <el-button icon="delete" size="small" :disabled="!apis.length" style="margin-left: 10px;" @click="deleteVisible = true">{{ t('general.delete') }}</el-button>
            </template>
          </el-popover>
        </div>
        <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
            <el-table-column type="selection" width="55" fixed="left" />
            <el-table-column label="ID" prop="id" width="70" />
            <el-table-column label="封面" width="100" align="center">
                <template #default="scope">
                <div v-if="scope.row.cover">
                    <img :src="'/uploads1/' + scope.row.cover" width="80" height="80" @click="openPriviewDialog('/uploads1/' + scope.row.cover)"/>
                </div>
                <div v-else>-无-</div>
                </template>
            </el-table-column>

            <el-table-column label="标题" prop="title" />
            <el-table-column label="简介" prop="description" />
            <el-table-column label="数量" prop="quantity" />
            <el-table-column label="排序" prop="sort" />

            <el-table-column label="创建时间" width="160">
                <template #default="scope">
                {{ formatDate(scope.row.created_at)}}
                </template>
            </el-table-column>

            <el-table-column label="更新时间" width="160">
                <template #default="scope">
                {{ formatDate(scope.row.updated_at)}}
                </template>
            </el-table-column>
          
  
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
  
      <el-dialog v-model="dialogPreviewVisible" :before-close="closePriviewDialog" title="图片预览" width="900px">
          <div>
            <embed type="image/jpg" :src="previewImg" style="width:100%"/>
          </div>

      </el-dialog>

      <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle">
        <el-form ref="apiForm" :model="form" :rules="rules" label-width="120px">          
            <el-form-item
          label="原封面:"
          v-if="form.cover && form.cover.indexOf('.') > 0"
        >
          <img :src="'uploads1/'+form.cover" width="100" height="100" />
        </el-form-item>
        
        <el-form-item :minlength="6" label="上传封面">
          <el-upload
            ref="upload_coverUrl"
            accept=".jpg,.jpeg,.png,.gif,.bmp,.pdf,.JPG,.JPEG,.PBG,.GIF,.BMP,.mp4,.MP4,.apk,.APk"
            list-type="picture-card"
            :class="{ imgHide: imgHideUpload }"
            action=""
            :limit="1"
            :auto-upload="false"
            :file-list="fileList_coverUrl"
            :on-exceed="handleExceed_coverUrl"
            :before-upload="handleBeforeUpload_coverUrl"
            :on-preview="handlePictureCardPreview_coverUrl"
            :on-success="handleSuccess_coverUrl"
            :on-remove="handleRemove_coverUrl"
            :on-change="handleChange_coverUrl"
          >
            <i class="el-icon-plus"></i>
          </el-upload>
          <el-dialog :visible.sync="dialogVisible_coverUrl" append-to-body>
            <img width="100%" :src="dialogImageUrl_coverUrl" alt="" />
          </el-dialog>
        </el-form-item>
        

        <el-form-item label="标题:" prop="title">
          <el-input
            v-model="form.title"
            clearable
            placeholder="请输入标题"
          />
        </el-form-item>

        <el-form-item label="简介:" prop="description">
          <el-input
            v-model="form.description"
            clearable
            placeholder="请输入简介"
          />
        </el-form-item>

        <el-form-item label="数量:" prop="quantity">
          <el-input
            v-model.number="form.quantity"
            clearable
            placeholder="请输入简介"
          />
        </el-form-item>

        <el-form-item label="排序:" prop="sort">
          <el-input
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
  import { formatTimeToStr } from "@/utils/date";

  import {
  bookList,
  bookDelete,
  bookDeleteBatch,
  bookOne,
  bookAdd,
  bookUpdate,
} from "@/api/book";

  import { toSQLLine } from '@/utils/stringFun'
  import { ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage
  
  const { t } = useI18n() // added by mohamed hassan to support multilanguage
  
  const methodFilter = (value) => {
    const target = methodOptions.value.filter(item => item.value === value)[0]
    return target && `${target.label}`
  }
  
  
  // 上传 begin
  const upload_coverUrl = ref(null)
  const fileList_coverUrl=ref([])

  const coverFile=ref('')
  const dialogImageUrl_coverUrl=ref('')
  
  const imgHideUpload=ref(false)
  const selectedfile_coverUrl=ref(false)
  const dialogVisible_coverUrl=ref(false)
  
  
  const handleBeforeUpload_coverUrl=(file)=>{
      console.log("before");
      if (
        !(
          file.type === "image/png" ||
          file.type === "image/gif" ||
          file.type === "image/jpg" ||
          file.type === "image/jpeg"
        )
      ) {
        ElMessage.warning('请上传格式为image/png, image/gif, image/jpg, image/jpeg的图片')
      }
      let size = file.size / 1024 / 1024 / 2;
      if (size > 10) {
        ElMessage.warning('图片大小必须小于10M')
      }
    }
  // 文件超出个数限制时的钩子
  const handleExceed_coverUrl=(files, fileList)=>{
    console.log(files, fileList);
    if (selectedfile_coverUrl.value || fileList.length >= 1) {      
      ElMessage.warning('上传的图片不能大于1张')
      return;
    }
  }
  // 点击文件列表中已上传的文件时的钩子
  const handlePictureCardPreview_coverUrl=(file)=>{
      dialogImageUrl_coverUrl.value = file.url;
      dialogVisible_coverUrl.value = true;
  }
  const handleSuccess_coverUrl=async (response, file, fileList)=>{
    if (response.code == 200) {
      console.log('handleSuccess_coverUrl---',response.data);
      // this.showUpdateCover = false;
      // 更新store中头像
      // await store.dispatch("user/changeCoverStore", response.data);
    }
  }
  const handleChange_coverUrl=(file, fileList)=>{
    coverFile.value=file.raw
    imgHideUpload.value = fileList.length >= 1;
    console.log('----handleChange_coverUrl----')
    console.log(file);
    console.log(coverFile.value);
    selectedfile_coverUrl.value = true;
  }
  // 文件列表移除文件时的钩子
  const handleRemove_coverUrl=(file, fileList)=>{
    imgHideUpload.value = fileList.length >= 1;
    fileList_coverUrl.value = [];
    selectedfile_coverUrl.value = false;
  }
  // 点击上传
  const uploadFile_coverUrl=()=>{
    if (!selectedfile_coverUrl.value) {
      ElMessage.warning('请选择一个文件')
      return;
    }
    upload_coverUrl.value.submit();
  }
  // 上传 end
    
  const apis = ref([])
  const form = ref({
    id: 0,
    title: "",
    description: "",
    quantity:0,
    cover: "",
    sort: 0,
    create_by: "",
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
    const table = await bookList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    const res = await bookDeleteBatch({ ids:ids.join(",") })
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
        id: 0,
        title: "",
        description: "",
        quantity:0,
        cover: "",
        sort: 0,
        create_by: "",
        update_by: "",
    }
  }
  
  const dialogTitle = ref('新增Api')
  const dialogFormVisible = ref(false)
  const openDialog = (key) => {
    switch (key) {
      case 'add':
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

  // preview image begin
  const dialogPreviewVisible = ref(false)
  const previewImg = ref('')
  const closePriviewDialog = () => {    
    dialogPreviewVisible.value = false
  }
  const openPriviewDialog = (url) => {    
    dialogPreviewVisible.value = true
    previewImg.value=url
  }
  // preview image end
  
  
  const editApiFunc = async(row) => {
    const res = await bookOne({ id: row.id })
    form.value = res.data.item
    openDialog('edit')
  }
  
  const enterDialog = async() => {
    apiForm.value.validate(async valid => {
      if (valid) {
        let postData = new FormData();
        postData.append('file',coverFile.value)
        postData.append('title',form.value.title)
        postData.append('description',form.value.description)
        postData.append('cover',form.value.cover)
        postData.append('sort',form.value.sort)
        postData.append('quantity',form.value.quantity)
        postData.append('create_by',form.value.create_by)
        postData.append('update_by',form.value.update_by)

        console.log('----enterDialog----')
        console.log(postData);
        console.log(coverFile.value);

        switch (type.value) {
          case 'add':
            {
                const res = await bookAdd(postData)
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
                postData.append('id',form.value.id)
                const res = await bookUpdate(postData)

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
        const res = await bookDelete(row)
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

  const  formatBoolean=(bool)=>{
    if (bool != null) {
        return bool ? "是" : "否";
      } else {
        return "";
      }
  }
  const formatDate=(time)=>{
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
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
  