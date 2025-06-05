<template>
  <div class="users-container">
    <a-card class="general-card">
      <!-- 搜索区域 -->
      <div class="search-form-wrapper">
        <a-form layout="inline" class="search-form" :model="searchForm">
          <div class="search-form-left">
            <a-form-item field="username">
              <a-input-search
                v-model="searchForm.username"
                :placeholder="t('users.placeholder.search')"
                @search="handleSearch"
                allow-clear
              />
            </a-form-item>
            <a-form-item field="role_id">
              <a-select
                v-model="searchForm.role_id"
                :placeholder="t('users.filterByRole')"
                :style="{ width: '180px' }"
                allow-clear
                @change="handleRoleChange"
              >
                <a-option v-for="role in rolesList" :key="role.id" :value="role.id">
                  {{ role.name }}
                </a-option>
              </a-select>
            </a-form-item>
            <!-- 状态 -->
            <a-form-item field="status">
              <a-select
                v-model="searchForm.status"
                placeholder="按状态筛选"
                :style="{ width: '180px' }"
                allow-clear
                @change="handleStatusChange"
              >
                <a-option :value="0">全部</a-option>
                <a-option :value="1">正常</a-option>
                <a-option :value="2">禁用</a-option>
              </a-select>
            </a-form-item>
          </div>
          <div class="search-form-right">
            <a-space>
              <a-button type="primary" @click="handleAdd">
                <template #icon>
                  <icon-plus />
                </template>
                {{ t('users.addUser') }}
              </a-button>
            </a-space>
          </div>
        </a-form>
      </div>

      <a-table
        :data="users"
        :loading="loading"
        :pagination="pagination"
        :empty="t('users.table.noData')"
        :bordered="true"
        :stripe="true"
        @page-change="onPageChange"
      >
        <template #columns>
          <a-table-column :title="t('users.username')" data-index="username" />
          <a-table-column :title="t('users.nickname')" data-index="nickname" />
          <a-table-column :title="t('users.role')" data-index="role_id">
            <template #cell="{ record }">
              <a-tag>{{ getRoleName(record.role_id) }}</a-tag>
            </template>
          </a-table-column>
          <a-table-column :title="t('common.status')" data-index="status">
            <template #cell="{ record }">
              <a-tag :color="record.status === 1 ? 'green' : 'red'">
                {{ getStatusText(record.status) }}
              </a-tag>
            </template>
          </a-table-column>
          <a-table-column title="允许IP" data-index="allowed_ips" />
          <a-table-column title="添加日期" data-index="created_at" />
          <a-table-column :title="t('common.operations')" align="right" :width="120">
            <template #cell="{ record }">
              <a-space>
                <a-button
                  type="text"
                  size="small"
                  @click="handleEdit(record)"
                >
                  <template #icon><icon-edit /></template>
                  {{ t('common.edit') }}
                </a-button>
                <a-button
                  type="text"
                  status="danger"
                  size="small"
                  @click="handleDelete(record)"
                >
                  <template #icon><icon-delete /></template>
                  {{ t('common.delete') }}
                </a-button>
              </a-space>
            </template>
          </a-table-column>
        </template>
      </a-table>

      <!-- Add/Edit User Modal -->
      <a-modal
        v-model:visible="modalVisible"
        :title="t(modalType === 'add' ? 'users.addUser' : 'users.editUser')"
        :ok-text="t('common.confirm')"
        :cancel-text="t('common.cancel')"
        @before-ok="handleModalOk"
        @cancel="handleModalCancel"
      >
        <a-form ref="formRef" :model="formData" :rules="rules">
          <a-form-item field="username" :label="t('users.username')" validate-trigger="blur">
            <a-input v-model="formData.username" :placeholder="t('users.placeholder.username')" />
          </a-form-item>
          <a-form-item field="nickname" label="昵称" validate-trigger="blur">
            <a-input v-model="formData.nickname" :placeholder="t('users.placeholder.username')" />
          </a-form-item>
          <!--
          <a-form-item field="email" :label="t('users.email')" validate-trigger="blur">
            <a-input v-model="formData.email" :placeholder="t('users.placeholder.email')" />
          </a-form-item>
          -->
          <a-form-item field="role_id" :label="t('users.role')" validate-trigger="blur">
            <a-select v-model="formData.role_id" :placeholder="t('users.placeholder.role')">
              <a-option v-for="role in rolesList" :key="role.id" :value="role.id">
                {{ role.name }}
              </a-option>
            </a-select>
          </a-form-item>
          <a-form-item field="status" :label="t('common.status')" validate-trigger="blur">
            <a-select v-model="formData.status">
              <a-option :value="1">{{ t('common.active') }}</a-option>
              <a-option :value="2">{{ t('common.inactive') }}</a-option>
            </a-select>
          </a-form-item>
          <a-form-item v-if="modalType === 'add'" field="password" :label="t('users.password')" validate-trigger="blur">
            <a-input-password v-model="formData.password" :placeholder="t('users.placeholder.password')" />
          </a-form-item>
          <a-form-item v-if="modalType === 'edit'" :label="t('users.password')">
            <a-input-password v-model="formData.password" :placeholder="t('users.placeholder.password')" />
          </a-form-item>

          <a-form-item label="IP限定" :validate-trigger="['change', 'blur']">
            <a-textarea
              v-model="formData.allowed_ips"
              placeholder="输入限定IP，多个IP用逗号分隔，留空表示不限制"
              :rows="2"
            />
          </a-form-item>
        </a-form>
      </a-modal>
    </a-card>
  </div>
</template>

<script setup>
import { ref, watch, onMounted, reactive } from 'vue'
import { IconPlus, IconEdit, IconDelete } from '@arco-design/web-vue/es/icon'
import { Message, Modal } from '@arco-design/web-vue'
import { useI18n } from 'vue-i18n'
// import { updateUser } from '@/api/mock/user'
import { list,deleteUser,addUser,updateUser } from '@/api/user'
import { getRoles } from '@/api/role'

const { t } = useI18n()

const searchForm = reactive({
  username: '',
  role_id: null,
  status:0,
})
const loading = ref(false)
const users = ref([])
const modalVisible = ref(false)
const modalType = ref('add')
const formRef = ref(null)
const currentRecord = ref(null)

const formData = ref({
  username: '',
  nickname:'',
  role_id:1,
  status: 1,
  password: '',
  allowed_ips: '',
})

const rules = {
  username: [{ required: true, message: t('users.validation.usernameRequired') }],
  nickname: [{ required: true, message: t('users.validation.usernameRequired') }],
  // email: [
  //   { required: true, message: t('users.validation.emailRequired') },
  //   { type: 'email', message: t('users.validation.emailInvalid') }
  // ],
  role_id: [{ required: true, message: t('users.validation.roleRequired') }],
  status: [{ required: true, message: t('users.validation.statusRequired') }],
  password: [{ required: true, message: t('users.validation.passwordRequired') }]
}

const pagination = ref({
  total: 0,
  current: 1,
  pageSize: 10,
  showTotal: true,
  showJumper: true,
  showPageSize: true,
  pageSizeOptions: [10, 20, 50, 100],
  baseSize: 10
})

// 添加角色列表数据
const rolesList = ref([])

const fetchRoles = async () => {
  try {
    const res = await getRoles()
    if (res.data.code === 200) {
      rolesList.value = res.data.data.list
    }
  } catch (error) {
    console.error('Failed to fetch roles:', error)
  }
}

const getRoleName = (roleId) => {
  const role = rolesList.value.find(r => r.id === roleId)
  return role ? role.name : roleId
}

const fetchUsers = async (page = 1) => {
  try {
    loading.value = true
    const res = await list({
      page: page,
      page_size: pagination.value.pageSize,
      username: searchForm.username,
      role_id: searchForm.role_id,
      status: searchForm.status
    })

    if (res.data.code === 200) {
      users.value = res.data.data.list
      pagination.value.total = res.data.data.total
      pagination.value.current = page
    } else {
      Message.error(t('users.fetchFailed'))
    }
  } catch (error) {
    console.error('Failed to fetch users:', error)
    // Message.error(t('users.fetchFailed'))
  } finally {
    loading.value = false
  }
}

const getStatusText = (type) => {
  const typeMap = {
    1: t('common.active'),
    2: t('common.inactive'),
  }
  return typeMap[type] || type
}

const handleSearch = () => {
  fetchUsers(1)
}

const handleRoleChange = () => {
  fetchUsers(1)
}

const handleAdd = () => {
  modalType.value = 'add'
  formData.value = {
    username: '',
    nickname:'',
    role_id: 1,
    status: 1,
    password: '',
    allowed_ips: '',
  }
  modalVisible.value = true
}

const handleEdit = (record) => {
  modalType.value = 'edit'
  currentRecord.value = record
  formData.value = { ...record }
  modalVisible.value = true
}

const handleDelete = (record) => {
  Modal.warning({
    title: t('common.confirm'),
    content: t('users.deleteConfirm'),
    okText: t('common.confirm'),
    cancelText: t('common.cancel'),
    hideCancel: false,
    onOk: async () => {
      try {
        const res = await deleteUser({ids:[record.id]})
        if (res.data.code === 200) {
          Message.success(t('users.deleteSuccess'))
          fetchUsers(pagination.value.current)
        } else {
          Message.error(res.data.message || t('users.deleteFailed'))
        }
      } catch (error) {
        console.error('Delete user error:', error)
        Message.error(t('users.deleteFailed'))
      }
    }
  })
}

const handleModalOk = async (done) => {
  try {
    const validRes = await formRef.value.validate()
    // validRes 为 undefined 时表示验证通过
    if (validRes) {
      // 验证失败
      done(false)
      return
    }
    const api = modalType.value === 'add' ? addUser : updateUser
    const params = modalType.value === 'add' ? [formData.value] : [formData.value, currentRecord.value.id]
    const res = await api(...params)
    if (res.data.code === 200) {
      Message.success(modalType.value === 'add' ? t('users.addSuccess') : t('users.updateSuccess'))
      modalVisible.value = false
      fetchUsers(pagination.value.current)
      done()
    } else {
      Message.error(res.data.message || (modalType.value === 'add' ? t('users.addFailed') : t('users.updateFailed')))
    }
  } catch (error) {
    done(false)
    console.error('Form validation failed:', error)
  }
}

const handleModalCancel = () => {
  modalVisible.value = false
  formRef.value?.resetFields()
}

const onPageChange = (page) => {
  fetchUsers(page)
}

// Watch for filter changes
watch(() => [searchForm.username, searchForm.role_id,searchForm.status], () => {
  fetchUsers(1)
})

// Initial data fetch
onMounted(() => {
  fetchUsers()
  fetchRoles()
})
</script>

<style scoped>
.users-container {
  background-color: var(--color-fill-2);
  min-height: 100%;
  width: calc(100% - 2px);
  margin-left: 2px;
}

.general-card {
  border-radius: 0;
  background-color: var(--color-bg-2);
  border:none !important;
}

.search-form-wrapper {
  margin-bottom: 16px;
  padding: 14px 14px 5px;
  background-color: var(--color-fill-2);
  border-radius: 4px;
}

.search-form {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  flex-wrap: wrap;
  gap: 16px;
}

.search-form-left {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
  flex: 1;
}

.search-form-right {
  flex-shrink: 0;
}

:deep(.arco-table-th) {
  background-color: var(--color-fill-2) !important;
}

:deep(.arco-table-tr:hover) {
  background-color: var(--color-fill-2) !important;
}

:deep(.arco-input-wrapper),
:deep(.arco-select-view) {
  background-color: var(--color-bg-2);
}

@media screen and (max-width: 768px) {
  .search-form {
    flex-direction: column;
  }
  
  .search-form-left {
    width: 100%;
  }
  
  .search-form-right {
    width: 100%;
    display: flex;
    justify-content: flex-end;
  }
  
  :deep(.arco-form-item) {
    margin-right: 0;
    width: 100%;
  }
}
</style> 