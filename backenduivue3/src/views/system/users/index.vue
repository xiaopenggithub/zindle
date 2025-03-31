<template>
  <div class="users-container">
    <a-card class="general-card">
      <!-- 搜索区域 -->
      <div class="search-form-wrapper">
        <a-form layout="inline" class="search-form" :model="searchForm">
          <div class="search-form-left">
            <a-form-item field="searchQuery">
              <a-input-search
                v-model="searchForm.searchQuery"
                :placeholder="t('users.placeholder.search')"
                search-button
                @search="handleSearch"
              />
            </a-form-item>
            <a-form-item field="roleFilter">
              <a-select
                v-model="searchForm.roleFilter"
                :placeholder="t('users.filterByRole')"
                :style="{ width: '180px' }"
                allow-clear
                @change="handleRoleChange"
              >
                <a-option value="admin">{{ t('users.roles.admin') }}</a-option>
                <a-option value="user">{{ t('users.roles.user') }}</a-option>
                <a-option value="guest">{{ t('users.roles.guest') }}</a-option>
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
          <a-table-column :title="t('users.email')" data-index="email" />
          <a-table-column :title="t('users.role')" data-index="role">
            <template #cell="{ record }">
              <a-tag>{{ t(`users.roles.${record.role}`) }}</a-tag>
            </template>
          </a-table-column>
          <a-table-column :title="t('common.status')" data-index="status">
            <template #cell="{ record }">
              <a-tag :color="record.status === 'active' ? 'green' : 'red'">
                {{ t(`common.${record.status}`) }}
              </a-tag>
            </template>
          </a-table-column>
          <a-table-column :title="t('users.createdAt')" data-index="createdAt" />
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
        @ok="handleModalOk"
        @cancel="handleModalCancel"
      >
        <a-form ref="formRef" :model="formData" :rules="rules">
          <a-form-item field="username" :label="t('users.username')" validate-trigger="blur">
            <a-input v-model="formData.username" :placeholder="t('users.placeholder.username')" />
          </a-form-item>
          <a-form-item field="email" :label="t('users.email')" validate-trigger="blur">
            <a-input v-model="formData.email" :placeholder="t('users.placeholder.email')" />
          </a-form-item>
          <a-form-item field="role" :label="t('users.role')" validate-trigger="blur">
            <a-select v-model="formData.role" :placeholder="t('users.placeholder.role')">
              <a-option value="admin">{{ t('users.roles.admin') }}</a-option>
              <a-option value="user">{{ t('users.roles.user') }}</a-option>
              <a-option value="guest">{{ t('users.roles.guest') }}</a-option>
            </a-select>
          </a-form-item>
          <a-form-item field="status" :label="t('common.status')" validate-trigger="blur">
            <a-select v-model="formData.status">
              <a-option value="active">{{ t('common.active') }}</a-option>
              <a-option value="inactive">{{ t('common.inactive') }}</a-option>
            </a-select>
          </a-form-item>
          <a-form-item v-if="modalType === 'add'" field="password" :label="t('users.password')" validate-trigger="blur">
            <a-input-password v-model="formData.password" :placeholder="t('users.placeholder.password')" />
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
import { getUsers, deleteUser, addUser, updateUser } from '@/api/mock/user'

const { t } = useI18n()

const searchForm = reactive({
  searchQuery: '',
  roleFilter: ''
})
const loading = ref(false)
const users = ref([])
const modalVisible = ref(false)
const modalType = ref('add')
const formRef = ref(null)
const currentRecord = ref(null)

const formData = ref({
  username: '',
  email: '',
  role: '',
  status: 'active',
  password: ''
})

const rules = {
  username: [{ required: true, message: t('users.validation.usernameRequired') }],
  email: [
    { required: true, message: t('users.validation.emailRequired') },
    { type: 'email', message: t('users.validation.emailInvalid') }
  ],
  role: [{ required: true, message: t('users.validation.roleRequired') }],
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

const fetchUsers = async (page = 1) => {
  try {
    loading.value = true
    const res = await getUsers({
      current: page,
      pageSize: pagination.value.pageSize,
      keyword: searchForm.searchQuery,
      role: searchForm.roleFilter
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
    Message.error(t('users.fetchFailed'))
  } finally {
    loading.value = false
  }
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
    email: '',
    role: '',
    status: 'active',
    password: ''
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
        const res = await deleteUser(record.id)
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

const handleModalOk = async () => {
  try {
    await formRef.value.validate()
    const api = modalType.value === 'add' ? addUser : updateUser
    const params = modalType.value === 'add' ? [formData.value] : [currentRecord.value.id, formData.value]
    
    const res = await api(...params)
    if (res.data.code === 200) {
      Message.success(modalType.value === 'add' ? t('users.addSuccess') : t('users.updateSuccess'))
      modalVisible.value = false
      fetchUsers(pagination.value.current)
    } else {
      Message.error(res.data.message || (modalType.value === 'add' ? t('users.addFailed') : t('users.updateFailed')))
    }
  } catch (error) {
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
watch(() => [searchForm.searchQuery, searchForm.roleFilter], () => {
  fetchUsers(1)
})

// Initial data fetch
onMounted(() => {
  fetchUsers()
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