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
        <el-button size="mini" type="primary" icon="plus" @click="openDialog('add')">{{ t('menu.addRootMenu') }}</el-button>
      </div>

      <!-- 由于此处菜单跟左侧列表一一对应所以不需要分页 pageSize默认999 -->
      <el-table :data="tableData" 
        row-key="id"
        default-expand-all
        :tree-props="{children:'children'}"
      >
        <el-table-column align="left" label="ID" min-width="100" prop="id" />
        <el-table-column align="left" :label="t('menu.displayName')" min-width="120" prop="authorityName">
          <template #default="scope">
            <span style="color:burlywood;font-weight:bold;" v-if="scope.row.parent_id==0">{{ scope.row.title }}</span>
            <span v-else>&nbsp;&nbsp;&nbsp;&nbsp;{{ scope.row.title }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('menu.icon')" min-width="140" prop="authorityName">
          <template #default="scope">
            <div class="icon-column">
              <el-icon>
                <component :is="scope.row.icon" />
              </el-icon>
              <span>{{ scope.row.icon }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('menu.routeName')" show-overflow-tooltip min-width="160" prop="name" />
        <el-table-column align="left" :label="t('menu.routePath')" show-overflow-tooltip min-width="160" prop="path" />
        <el-table-column align="left" :label="t('menu.visibility')" min-width="100" prop="hidden">
          <template #default="scope">
            <span>{{ scope.row.hidden? t('menu.hide') : t('menu.show') }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('menu.parent')" min-width="90" prop="parent_id" />
        <el-table-column align="left" :label="t('menu.sort')" min-width="70" prop="sort" />
        <el-table-column align="left" :label="t('menu.filePath')" min-width="360" prop="component" />
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
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="handleClose" :title="dialogTitle" width="800px">      
      <el-form
        v-if="dialogFormVisible"
        ref="menuForm"
        :model="form"
        :rules="rules"
        label-position="right"
        label-width="100px"
      >
      <el-form-item label="父级:" prop="parent_id">
          <el-select v-model="form.parent_id" placeholder="请选择父级">
            <el-option label="顶级" :value="0"></el-option>
            <el-option
              :label="item.title"
              :value="item.id"
              v-for="(item, index) in parents"
              :key="index"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="菜单名称:" prop="title">
          <el-input
            v-model="form.title"
            clearable
            placeholder="请输入附加属性"
          />
        </el-form-item>

        <el-form-item label="vue路径:" prop="component">
          <el-input
            v-model="form.component"
            clearable
            placeholder="请输入对应前端vue文件路径"
          />
        </el-form-item>

        <el-form-item label="图标:" prop="icon">
          <icon :meta="form">
            <template slot="prepend">el-icon-</template>
          </icon>
        </el-form-item>

        <el-form-item label="路由name:" prop="name">
          <el-input
            v-model="form.name"
            clearable
            placeholder="请输入路由name"
          />
        </el-form-item>

        <el-form-item label="路由path:" prop="path">
          <el-input
            v-model="form.path"
            clearable
            placeholder="请输入路由path"
          />
        </el-form-item>

        <el-form-item label="列表隐藏:" prop="hidden">
          <el-radio-group v-model="form.hidden">
            <el-radio :label="0" name="hidden">否</el-radio>
            <el-radio :label="1" name="hidden">是</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="排序:" prop="sort">
          <el-input
            type="number"
            v-model.number="form.sort"
            clearable
            placeholder="请输入排序标记"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">{{ t('general.cancel') }}</el-button>
          <el-button size="small" type="primary" @click="enterDialog">{{ t('general.confirm') }}</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
// import {
//   updateBaseMenu,
//   getMenuList,
//   addBaseMenu,
//   deleteBaseMenu,
//   getBaseMenuById
// } from '@/api/menu'
import {
  systemMenuList,
  systemMenuParent,
  systemMenuDelete,
  systemMenuDeleteBatch,
  systemMenuOne,
  systemMenuAdd,
  systemMenuUpdate,
} from "@/api/systemMenu";
import icon from '@/view/superAdmin/systemMenu/icon.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
// import { canRemoveAuthorityBtnApi } from '@/api/authorityBtn'
import { reactive, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const rules = reactive({
  path: [{ required: true, message: t('menu.enterMenuNameNote'), trigger: 'blur' }],
  component: [
    { required: true, message: t('menu.enterFilePathNote'), trigger: 'blur' }
  ],
  'meta.title': [
    { required: true, message: t('menu.enterMenuDisplayNameNote'), trigger: 'blur' }
  ]
})


const onReset = () => {
  searchInfo.value = {}
}
// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 999
  getTableData()
}

const parents = ref([])
//获取父级
const getParent=async (page, pageSize)=>{
  const res = await systemMenuParent({ page, pageSize });
  console.log(res);
  if (res.code == 200) {
    parents.value = res.data.list;
  }
}
//加载父级
getParent(1, 200)

const deleteFunc = async(row) => {
  ElMessageBox.confirm(t('view.api.deleteApiConfirm'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    type: 'warning'
  })
    .then(async() => {
      const res = await systemMenuDelete(row)
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

const type = ref('')
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

const editFunc = async(row) => {
  const res = await systemMenuOne({ id: row.id })
  form.value = res.data.item
  openDialog('edit')
}

const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const tableData = ref([])
const searchInfo = ref({})
// 查询
const getTableData = async() => {
  const table = await systemMenuList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 200) {
    let menu_data=[]
    for(let m of table.data.list){
      if(m.parent_id==0){
        m.children=[]
        menu_data.push(m)
      }
    }
    for(let p of menu_data){
      for(let m of table.data.list){
        if(m.parent_id==p.id){
          p.children.push(m)
        }
      }
    }
    tableData.value = menu_data
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// 新增参数
const addParameter = (form) => {
  if (!form.parameters) {
    form.parameters = []
  }
  form.parameters.push({
    type: 'query',
    key: '',
    value: ''
  })
}

const fmtComponent = () => {
  form.value.component = form.value.component.replace(/\\/g, '/')
}

// 删除参数
const deleteParameter = (parameters, index) => {
  parameters.splice(index, 1)
}

// 新增可控按钮
const addBtn = (form) => {
  if (!form.menuBtn) {
    form.menuBtn = []
  }
  form.menuBtn.push({
    name: '',
    desc: '',
  })
}
// 删除可控按钮
const deleteBtn = async(btns, index) => {
  const btn = btns[index]
  if (btn.ID === 0) {
    btns.splice(index, 1)
    return
  }  
}

const form = ref({
  ID: 0,
  path: '',
  name: '',
  hidden: false,
  parentId: '',
  component: '',
  meta: {
    activeName: '',
    title: '',
    icon: '',
    defaultMenu: false,
    closeTab: false,
    keepAlive: false
  },
  parameters: [],
  menuBtn: []
})
const changeName = () => {
  form.value.path = form.value.name
}

const handleClose = (done) => {
  initForm()
  done()
}
// 删除菜单
// const deleteMenu = (ID) => {
//   ElMessageBox.confirm(t('menu.deleteAllRolesConfirm'), t('general.hint'), {
//     confirmButtonText: t('general.confirm'),
//     cancelButtonText: t('general.cancel'),
//     type: 'warning'
//   })
//     .then(async() => {
//       const res = await deleteBaseMenu({ ID })
//       if (res.code === 0) {
//         ElMessage({
//           type: 'success',
//           message: t('general.deleteSuccess')
//         })
//         if (tableData.value.length === 1 && page.value > 1) {
//           page.value--
//         }
//         getTableData()
//       }
//     })
//     .catch(() => {
//       ElMessage({
//         type: 'info',
//         message: t('general.undeleted')
//       })
//     })
// }
// 初始化弹窗内表格方法
const menuForm = ref(null)
const checkFlag = ref(false)
const initForm = () => {
  checkFlag.value = false
  menuForm.value.resetFields()
  form.value = {
    ID: 0,
    path: '',
    name: '',
    hidden: false,
    parentId: '',
    component: '',
    meta: {
      title: '',
      icon: '',
      defaultMenu: false,
      closeTab: false,
      keepAlive: false
    }
  }
}
// 关闭弹窗

const dialogFormVisible = ref(false)
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}
// 添加menu
const enterDialog = async() => {
  menuForm.value.validate(async valid => {
    if (valid) {
      let res
      if (type.value=='edit') {
        res = await systemMenuUpdate(form.value)
      } else {
        res = await systemMenuAdd(form.value)
      }
      if (res.code === 200) {
        ElMessage({
          type: 'success',
          message: isEdit.value ? t('general.editSuccess') : t('general.addSuccess')
        })
        getTableData()
      }
      initForm()
      dialogFormVisible.value = false
    }
  })
}

const menuOption = ref([
  {
    ID: '0',
    title: t('menu.rootMenu')
  }
])
const setOptions = () => {
  menuOption.value = [
    {
      ID: '0',
      title: t('menu.rootDirctory')
    }
  ]
  setMenuOptions(tableData.value, menuOption.value, false)
}
const setMenuOptions = (menuData, optionsData, disabled) => {
  menuData &&
        menuData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
              disabled: disabled || item.ID === form.value.ID,
              children: []
            }
            setMenuOptions(
              item.children,
              option.children,
              disabled || item.ID === form.value.ID
            )
            optionsData.push(option)
          } else {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
              disabled: disabled || item.ID === form.value.ID
            }
            optionsData.push(option)
          }
        })
}

// 添加菜单方法，id为 0则为添加根菜单
const isEdit = ref(false)
const dialogTitle = ref(t('menu.addMenu'))
const addMenu = (id) => {
  dialogTitle.value = t('menu.addMenu')
  form.value.parentId = String(id)
  isEdit.value = false
  setOptions()
  dialogFormVisible.value = true
}
// 修改菜单方法
// const editMenu = async(id) => {
//   dialogTitle.value = t('menu.editMenu')
//   const res = await getBaseMenuById({ id })
//   form.value = res.data.menu
//   isEdit.value = true
//   setOptions()
//   dialogFormVisible.value = true
// }

</script>

<script>
export default {
  name: 'Menus',
}
</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}
.icon-column{
  display: flex;
  align-items: center;
  .el-icon{
    margin-right: 8px;
  }
}
</style>
