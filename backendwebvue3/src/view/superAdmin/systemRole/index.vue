<template>
  <div class="authority">
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="addAuthority(0)">{{ t('authority.addRole') }}</el-button>
      </div>
      <el-table
        :data="tableData"
        row-key="id"
        style="width: 100%"
      >
        <el-table-column :label="t('authority.roleID')" width="70" prop="id" />
        <el-table-column align="left" :label="t('authority.roleName')" prop="name" />
        <el-table-column label="父级ID" prop="parent_id" width="70" />
        <el-table-column label="排序" prop="sort" width="70" />
        <el-table-column align="left" :lable="t('general.operations')" width="260">
          <template #default="scope">
            <el-button
              icon="setting"
              size="small"
              type="primary"
              link
              @click="opdendrawer(scope.row)"
            >{{ t('authority.setPermissions') }}</el-button>
            <el-button
              icon="edit"
              size="small"
              type="primary"
              link
              @click="editAuthority(scope.row)"
            >{{ t('general.edit') }}</el-button>
            <el-button
              icon="delete"
              size="small"
              type="primary"
              link
              @click="deleteAuth(scope.row)"
            >{{ t('general.delete') }}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- 新增角色弹窗 -->
    <el-dialog v-model="dialogFormVisible" :title="dialogTitle">
      <el-form ref="authorityForm" :model="form" :rules="rules" label-width="100px">
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

        <el-form-item label="角色名:" prop="name">
          <el-input
            v-model="form.name"
            clearable
            placeholder="请输入角色名"
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

    <el-drawer v-if="drawer" v-model="drawer" custom-class="auth-drawer" :with-header="false" size="40%" :title="t('authority.roleConfig')">
      <el-tabs :before-leave="autoEnter" type="border-card">
        <el-tab-pane :label="t('authority.roleAPI')">
          <Apis ref="apis" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>
        <el-tab-pane :label="t('authority.roleMenu')">
          <Menus ref="menus" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>                
      </el-tabs> 
    </el-drawer>
  </div>
</template>

<script setup>
import {
  systemRoleList,
  systemRoleParent,
  systemRoleDelete,
  // systemRoleDeleteBatch,
  // systemRoleOne,
  systemRoleAdd,
  systemRoleUpdate,
} from "@/api/systemRole";

import Menus from '@/view/superAdmin/systemRole/components/menus.vue'
import Apis from '@/view/superAdmin/systemRole/components/apis.vue'

import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const mustUint = (rule, value, callback) => {
  if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
    return callback(new Error(t('authority.positiveIntNote')))
  }
  return callback()
}

const AuthorityOption = ref([
  {
    authorityId: 0,
    authorityName: t('authority.rootRole')
  }
])
const drawer = ref(false)
const dialogType = ref('add')
const activeRow = ref({})

const dialogTitle = ref(t('authority.addRole'))
const dialogFormVisible = ref(false)
const apiDialogFlag = ref(false)

const form = ref({
  authorityId: 0,
  authorityName: '',
  parentId: 0
})
const rules = ref({
  authorityId: [
    { required: true, message: t('authority.roleIdNote'), trigger: 'blur' },
    { validator: mustUint, trigger: 'blur', message: '必须为正整数' }
  ],
  authorityName: [
    { required: true, message: t('authority.roleNameNote'), trigger: 'blur' }
  ],
  parentId: [
    { required: true, message: t('authority.roleSelectMethod'), trigger: 'blur' },
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const tableData = ref([])
const searchInfo = ref({})

// 查询
const getTableData = async() => {
  const table = await systemRoleList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 200) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()


const parents =ref([])
const getParent=async(page, pageSize)=>{
  const res = await systemRoleParent({ page, pageSize });  
  if (res.code == 200) {
    parents.value = res.data.list;
  }
}
getParent(1, 200)


const changeRow = (key, value) => {
  activeRow.value[key] = value
}
const menus = ref(null)
const apis = ref(null)
const datas = ref(null)
const autoEnter = (activeName, oldActiveName) => {
  const paneArr = [menus, apis, datas]
  if (oldActiveName) {
    if (paneArr[oldActiveName].value.needConfirm) {
      paneArr[oldActiveName].value.enterAndNext()
      paneArr[oldActiveName].value.needConfirm = false
    }
  }
}
const opdendrawer = (row) => {
  drawer.value = true
  activeRow.value = row
}
// 删除角色
const deleteAuth = (row) => {
  ElMessageBox.confirm(t('authority.roleDeleteConfirm'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    type: 'warning'
  })
    .then(async() => {
      const res = await systemRoleDelete({ id: row.id })
      if (res.code === 200) {
        ElMessage({
          type: 'success',
          message: res.message
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: t('general.undeleted')
      })
    })
}
// 初始化表单
const authorityForm = ref(null)
const initForm = () => {
  if (authorityForm.value) {
    authorityForm.value.resetFields()
  }
  form.value = {
    id: 0,
    name: "",
    parent_id: 0,
    sort: 0,
  }
}
// 关闭窗口
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
  apiDialogFlag.value = false
}
// 确定弹窗

const enterDialog = () => {
  form.value.authorityId = Number(form.value.authorityId)
  if (form.value.authorityId === 0) {
    ElMessage({
      type: 'error',
      message: t('authority.roleId0Error')
    })
    return false
  }
  authorityForm.value.validate(async valid => {
    if (valid) {
      switch (dialogType.value) {
        case 'add':
          {
            const res = await systemRoleAdd(form.value)
            if (res.code === 200) {
              ElMessage({
                type: 'success',
                message: t('general.addSuccess')
              })
              getTableData()
              closeDialog()
            }
          }
          break
        case 'edit':
          {
            const res = await systemRoleUpdate(form.value)
            if (res.code === 200) {
              ElMessage({
                type: 'success',
                message: t('general.updateSuccess')
              })
              getTableData()
              closeDialog()
            }
          }
          break
      }

      initForm()
      dialogFormVisible.value = false
    }
  })
}
const setOptions = () => {
  AuthorityOption.value = [
    {
      authorityId: 0,
      authorityName: t('authority.rootRole')
    }
  ]
  setAuthorityOptions(tableData.value, AuthorityOption.value, false)
}
const setAuthorityOptions = (AuthorityData, optionsData, disabled) => {
  form.value.authorityId = String(form.value.authorityId)
  AuthorityData &&
        AuthorityData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              disabled: disabled || item.authorityId === form.value.authorityId,
              children: []
            }
            setAuthorityOptions(
              item.children,
              option.children,
              disabled || item.authorityId === form.value.authorityId
            )
            optionsData.push(option)
          } else {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              disabled: disabled || item.authorityId === form.value.authorityId
            }
            optionsData.push(option)
          }
        })
}
// 增加角色
const addAuthority = (parentId) => {
  initForm()
  dialogTitle.value = t('authority.addRole')
  dialogType.value = 'add'
  form.value.parentId = parentId
  setOptions()
  dialogFormVisible.value = true
}
// 编辑角色
const editAuthority = (row) => {
  setOptions()
  dialogTitle.value = t('authority.editRole')
  dialogType.value = 'edit'
  for (const key in form.value) {
    form.value[key] = row[key]
  }
  setOptions()
  dialogFormVisible.value = true
}

</script>

<style lang="scss">
.authority {
  .el-input-number {
    margin-left: 15px;
    span {
      display: none;
    }
  }
}
.tree-content{
  overflow: auto;
  height: calc(100vh - 100px);
  margin-top: 10px;
}

.auth-drawer{
  .el-drawer__body{
    overflow: hidden;
  }
}
</style>
