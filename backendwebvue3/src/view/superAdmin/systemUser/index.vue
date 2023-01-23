<template>
  <div>
    <div class="gva-search-box">
        <el-form ref="searchForm" :inline="true" :model="searchInfo">
          <el-form-item label="用户名">
            <el-input v-model="searchInfo.keyword" placeholder="输入用户名" clearable/>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary" icon="search" @click="onSubmit">{{ t('general.search') }}</el-button>
            <el-button size="small" icon="refresh" @click="onReset">{{ t('general.reset') }}</el-button>
          </el-form-item>
        </el-form>
      </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog('add')">{{ t('user.addUser') }}</el-button>

        <el-popover v-model="deleteVisible" placement="top" width="160">
            <p>{{ t('general.deleteConfirm') }}</p>
            <div style="text-align: right; margin-top: 8px;">
              <el-button size="small" type="primary" link @click="deleteVisible = false">{{ t('general.cancel') }}</el-button>
              <el-button size="small" type="primary" @click="onDelete">{{ t('general.confirm') }}</el-button>
            </div>
            <template #reference>
              <el-button icon="delete" size="small" :disabled="!user_ids.length" style="margin-left: 10px;" @click="deleteVisible = true">{{ t('general.delete') }}</el-button>
            </template>
          </el-popover>
      </div>
      <el-table :data="tableData" @selection-change="handleSelectionChange">
        <el-table-column
            type="selection"
            width="55"
          />
        <el-table-column label="头像" width="100" align="center">
          <template #default="scope">
            <div v-if="scope.row.avatar">
              <img :src="'/uploads/'+scope.row.avatar" width="80" height="80"/>
            </div>
            <div v-else>-无-</div>
          </template>
        </el-table-column>

        <el-table-column label="用户账号" prop="user_name" width="120" fixed="left"/>

        <el-table-column label="部门" prop="department_name" width="120" />
        <el-table-column label="角色" prop="role_name" width="120" />

        <el-table-column label="用户昵称" prop="nick_name" width="120" />

        <el-table-column label="用户类型" width="80" align="center">
          <template #default="scope">
            {{ formatuserType(scope.row.user_type) }}
          </template>
        </el-table-column>

        <el-table-column label="用户邮箱" prop="email" width="120" />

        <el-table-column label="手机号码" prop="phonenumber" width="120" />

        <el-table-column label="性别" width="60" align="center">
          <template #default="scope">
            {{ formatSex(scope.row.sex) }}
          </template>
        </el-table-column>


        <el-table-column label="帐号状态" width="80" align="center">
          <template #default="scope">
            {{ formatStatus(scope.row.status)}}
          </template>
        </el-table-column>

        <el-table-column label="删除标志" width="80" align="center">
          <template #default="scope">
            {{ formatdelFlag(scope.row.del_flag) }}
          </template>
        </el-table-column>

        <el-table-column label="最后登录IP" prop="login_ip" width="120" />

        <el-table-column label="备注" prop="remark" width="120" />

        <el-table-column label="最后登录时间" width="160">
          <template #default="scope">
            {{ formatDate(scope.row.login_date) }}
          </template>
        </el-table-column>

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
    <el-dialog
      v-model="dialogFormVisible"
      :title="t('user.addUser')"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <div style="height:60vh;overflow:auto;padding:0 12px;">
        <el-form ref="userForm" :rules="rules" :model="userInfo" label-width="80px">
          <el-form-item label="部门:" prop="dept_id">
            <el-select v-model="form.dept_id" placeholder="请选择部门">
              <el-option label="-不指定-" :value="0"></el-option>
              <el-option
                :label="item.name"
                :value="item.id"
                v-for="(item, index) in departments"
                :key="index"
              ></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="角色:" prop="role_id">
            <el-select v-model="form.role_id" placeholder="请选择角色">
              <el-option label="-不指定-" :value="0"></el-option>
              <el-option
                :label="item.name"
                :value="item.id"
                v-for="(item, index) in roles"
                :key="index"
              ></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="用户账号:" prop="user_name">
            <el-input
              v-model="form.user_name"
              clearable
              placeholder="请输入用户账号"
            />
          </el-form-item>

          <el-form-item label="密码:" prop="password">
            <el-input
              type="password"
              v-model="form.password"
              clearable
              placeholder="请输入密码"
            />
            <div style="color: red" v-if="type != 'create'">密码留空不修改</div>
          </el-form-item>

          <el-form-item label="手机号码:" prop="phonenumber">
            <el-input
              v-model="form.phonenumber"
              clearable
              placeholder="请输入手机号码"
            />
          </el-form-item>

          <el-form-item label="用户昵称:" prop="nick_name">
            <el-input
              v-model="form.nick_name"
              clearable
              placeholder="请输入用户昵称"
            />
          </el-form-item>

          <el-form-item label="用户邮箱:" prop="email">
            <el-input
              v-model="form.email"
              clearable
              placeholder="请输入用户邮箱"
            />
          </el-form-item>

          <el-form-item label="删除:" prop="del_flag">
            <el-radio-group v-model="form.del_flag">
              <el-radio :label="0" name="del_flag">保留</el-radio>
              <el-radio :label="1" name="del_flag">删除</el-radio>
            </el-radio-group>
          </el-form-item>

          <el-form-item label="帐号状态:" prop="status">
            <el-radio-group v-model="form.status">
              <el-radio :label="0" name="status">正常</el-radio>
              <el-radio :label="1" name="status">停用</el-radio>
            </el-radio-group>
          </el-form-item>

          <el-form-item label="性别:" prop="sex">
            <el-radio-group v-model="form.sex">
              <el-radio :label="0" name="sex">男</el-radio>
              <el-radio :label="1" name="sex">女</el-radio>
              <el-radio :label="2" name="sex">未知</el-radio>
            </el-radio-group>
          </el-form-item>

          <el-form-item label="用户类型:" prop="status">
            <el-radio-group v-model="form.user_type">
              <el-radio :label="0" name="user_type">系统用户</el-radio>
              <el-radio :label="1" name="user_type">其它</el-radio>
            </el-radio-group>
          </el-form-item>

          <el-form-item label="备注:" prop="remark">
            <el-input
              type="textarea"
              v-model="form.remark"
              clearable
              placeholder="请输入备注"
            />
          </el-form-item>
        </el-form>
      </div>
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
  getUserList,
  setUserAuthorities,
  register,
  deleteUser
} from '@/api/user'
import {
  getSystemUserList,
  findSystemUser,
  createSystemUser,
  deleteSystemUser,
  deleteSystemUserByIds,
  updateSystemUser,  
} from "@/api/systemUser";
import { systemDepartmentList } from "@/api/systemDepartment";
import { systemRoleList } from "@/api/systemRole";

// import { getAuthorityList } from '@/api/authority'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { setUserInfo, resetPassword } from '@/api/user.js'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage
import { formatTimeToStr } from "@/utils/date";

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const path = ref(import.meta.env.VITE_BASE_API + '/')
// 初始化相关
const setAuthorityOptions = (AuthorityData, optionsData) => {
  AuthorityData &&
        AuthorityData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              children: []
            }
            setAuthorityOptions(item.children, option.children)
            optionsData.push(option)
          } else {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName
            }
            optionsData.push(option)
          }
        })
}


const form = ref({
    avatar: "",
    create_by: "",
    del_flag: 0,
    dept_id: 0,
    role_id: 0,
    email: "",
    id: 0,
    login_ip: "",
    nick_name: "",
    password: "",
    phonenumber: "",
    remark: "",
    sex: 0,
    status: 0,
    update_by: "",
    user_name: "",
    user_type: 0,
  })
  const userForm = ref(null)
  const initForm = () => {
    userForm.value.resetFields()
    form.value = {
      avatar: "",
      create_by: "",
      del_flag: 0,
      dept_id: 0,
      role_id: 0,
      email: "",
      id: 0,
      login_ip: "",
      nick_name: "",
      password: "",
      phonenumber: "",
      remark: "",
      sex: 0,
      status: 0,
      update_by: "",
      user_name: "",
      user_type: 0,
    }
  }

  // 批量操作
  const user_ids = ref([])
  const handleSelectionChange = (val) => {
    user_ids.value = val
  }
  const deleteVisible = ref(false)
  const onDelete = async() => {
    const ids = user_ids.value.map(item => item.id)
    const res = await deleteSystemUserByIds({ ids:ids.join(",") })
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





  const departments = ref([])
  const roles = ref([])
  //加载部门
  const getDepartments=async (page, pageSize)=>{
    const res = await systemDepartmentList({ page, pageSize });
    console.log(res);
    if (res.code == 200) {
      departments.value = res.data.list;
    }
  }
  //加载角色
  const getRoles=async(page, pageSize)=>{
    const res = await systemRoleList({ page, pageSize });
    console.log(res);
    if (res.code == 200) {
      roles.value = res.data.list;
    }
  }

  // 加载部门
  getDepartments(1, 200);
  // 加载角色
  getRoles(1, 200);


  const type = ref('')
  const dialogTitle = ref('新增')
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

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getSystemUserList({ page: page.value, pageSize: pageSize.value,...searchInfo.value })
  if (table.code === 200) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const deleteFunc = async(row) => {
  ElMessageBox.confirm(t('view.api.deleteApiConfirm'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteSystemUser(row)
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

const editFunc = async(row) => {
  const res = await findSystemUser({ id: row.id })
  form.value = res.data.item
  openDialog('edit')
}

const enterDialog = async() => {
  userForm.value.validate(async valid => {
      if (valid) {
        switch (type.value) {
          case 'add':
            {
              const res = await createSystemUser(form.value)
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
              const res = await updateSystemUser(form.value)
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

watch(() => tableData.value, () => {
  setAuthorityIds()
})

const initPage = async() => {
  getTableData()
  // const res = await getAuthorityList({ page: 1, pageSize: 999 })
  // setOptions(res.data.list)
}

initPage()

const resetPasswordFunc = (row) => {
  ElMessageBox.confirm(
    t('user.resetPasswordConfrim'),
    t('general.warning'),
    {
      confirmButtonText: t('general.confirm'),
      cancelButtonText: t('general.cancel'),
      type: 'warning',
    }
  ).then(async() => {
    const res = await resetPassword({
      ID: row.ID,
    })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: res.msg,
      })
    } else {
      ElMessage({
        type: 'error',
        message: res.msg,
      })
    }
  })
}
const setAuthorityIds = () => {
  tableData.value && tableData.value.forEach((user) => {
    const authorityIds = user.authorities && user.authorities.map(i => {
      return i.authorityId
    })
    user.authorityIds = authorityIds
  })
}

const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}

// const authOptions = ref([])
// const setOptions = (authData) => {
//   authOptions.value = []
//   setAuthorityOptions(authData, authOptions.value)
// }

const backNickName = ref('')
const openEidt = (row) => {
  if (tableData.value.some(item => item.editFlag)) {
    ElMessage(t('user.anotherUserEdit'))
    return
  }
  backNickName.value = row.nickName
  row.editFlag = true
}

const enterEdit = async(row) => {
  const res = await setUserInfo({ nickName: row.nickName, ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: t('general.setupSuccess')
    })
  }
  backNickName.value = ref('')
  row.editFlag = false
}

const closeEdit = (row) => {
  row.nickName = backNickName.value
  backNickName.value = ''
  row.editFlag = false
}

const deleteUserFunc = async(row) => {
  const res = await deleteUser({ id: row.id })
  if (res.code === 0) {
    ElMessage.success(t('general.deleteSuccess'))
    await getTableData()
    row.visible = false
  }
}

// 弹窗相关
const userInfo = ref({
  username: '',
  password: '',
  nickName: '',
  headerImg: '',
  authorityId: '',
  authorityIds: [],
  enable: 1,
})

const validatePassword = (rule, value, callback) => {  
  if (type.value == "add") {
    if (form.value.password.length < 6) {
      callback(new Error("密码长度不少于6位"));
    } else {
      callback();
    }
  } else {
    if (form.value.password.length > 0 && form.value.password.length < 6) {
      callback(new Error("密码长度不少于6位"));
    } else {
      callback();
    }
  }
}
const rules = ref({
  username: [
    { required: true, message: t('user.userNameNote'), trigger: 'blur' },
    { min: 5, message: t('user.userNameLenNote'), trigger: 'blur' }
  ],
  password: [
    { required: true, trigger: "blur", validator: validatePassword },
  ],
  nickName: [
    { required: true, message: t('user.nickNameNote'), trigger: 'blur' }
  ],
  authorityId: [
    { required: true, message: t('user.userRoleNote'), trigger: 'blur' }
  ]
})
const enterAddUserDialog = async() => {
  userInfo.value.authorityId = userInfo.value.authorityIds[0]
  userForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...userInfo.value
      }
      if (dialogFlag.value === 'add') {
        const res = await register(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: t('user.userAddedNote') })
          await getTableData()
          closeAddUserDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await setUserInfo(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: t('user.userEditedNote') })
          await getTableData()
          closeAddUserDialog()
        }
      }
    }
  })
}

const addUserDialog = ref(false)
const closeAddUserDialog = () => {
  userForm.value.resetFields()
  userInfo.value.headerImg = ''
  userInfo.value.authorityIds = []
  addUserDialog.value = false
}

const dialogFlag = ref('add')

const addUser = () => {
  addUserDialog.value = true
}

const tempAuth = {}
const changeAuthority = async(row, flag, removeAuth) => {
  if (flag) {
    if (!removeAuth) {
      tempAuth[row.ID] = [...row.authorityIds]
    }
    return
  }
  await nextTick()
  const res = await setUserAuthorities({
    ID: row.ID,
    authorityIds: row.authorityIds
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: t('user.roleSetNote') })
  } else {
    if (!removeAuth) {
      row.authorityIds = [...tempAuth[row.ID]]
      delete tempAuth[row.ID]
    } else {
      row.authorityIds = [removeAuth, ...row.authorityIds]
    }
  }
}

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  userInfo.value = JSON.parse(JSON.stringify(row))
  addUserDialog.value = true
}

const switchEnable = async(row) => {
  userInfo.value = JSON.parse(JSON.stringify(row))
  await nextTick()
  const req = {
    ...userInfo.value
  }
  const res = await setUserInfo(req)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: `${req.enable === 2 ? '禁用' : '启用'}成功` })
    await getTableData()
    userInfo.value.headerImg = ''
    userInfo.value.authorityIds = []
  }
}


const formatStatus=(v)=>{
  if (v == 0) {
    return "正常";
  }
  return "停用";
}


const formatSex=(v)=>{
  if (v == 0) {
    return "男";
  } else if (v == 1) {
    return "女";
  } else {
    return "未知";
  }
}


const formatdelFlag=(v)=>{
  if (v == 0) {
    return "保留";
  }
  return "删除";
}

const formatuserType=(v)=>{
  if (v == 0) {
    return "系统用户";
  }
  return "其它";
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

<style lang="scss">
.user-dialog {
  .header-img-box {
  width: 200px;
  height: 200px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
.nickName{
  display: flex;
  justify-content: flex-start;
  align-items: center;
}
.pointer{
  cursor: pointer;
  font-size: 16px;
  margin-left: 2px;
}
</style>
