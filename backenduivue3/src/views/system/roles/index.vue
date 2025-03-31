<template>
  <div class="roles-container">
    <a-card class="general-card">
      <!-- 搜索区域 -->
      <div class="search-form-wrapper">
        <a-form layout="inline" class="search-form" :model="searchForm">
          <div class="search-form-left">
            <a-form-item field="searchQuery">
              <a-input-search
                v-model="searchForm.searchQuery"
                :placeholder="t('roles.placeholder.search')"
                search-button
                @search="handleSearch"
              />
            </a-form-item>
          </div>
          <div class="search-form-right">
            <a-space>
              <a-button type="primary" @click="handleAdd">
                <template #icon>
                  <icon-plus />
                </template>
                {{ t('roles.addRole') }}
              </a-button>
            </a-space>
          </div>
        </a-form>
      </div>

      <a-table
        :data="roles"
        :loading="loading"
        :pagination="pagination"
        :empty="t('roles.table.noData')"
        :bordered="true"
        :stripe="true"
        @page-change="onPageChange"
      >
        <template #columns>
          <a-table-column :title="t('roles.name')" data-index="name" />
          <a-table-column :title="t('roles.code')" data-index="code" />
          <a-table-column :title="t('roles.description')" data-index="description" />
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

      <!-- 新增/编辑角色模态框 -->
      <a-modal
        v-model:visible="modalVisible"
        :title="t(modalType === 'add' ? 'roles.addRole' : 'roles.editRole')"
        @ok="handleModalOk"
        @cancel="modalVisible = false"
      >
        <a-form
          ref="formRef"
          :model="formData"
          :rules="rules"
          @submit="handleModalOk"
        >
          <a-form-item field="name" :label="t('roles.name')" :validate-trigger="['change', 'blur']">
            <a-input
              v-model="formData.name"
              :placeholder="t('roles.placeholder.name')"
            />
          </a-form-item>
          <a-form-item field="code" :label="t('roles.code')" :validate-trigger="['change', 'blur']">
            <a-input
              v-model="formData.code"
              :placeholder="t('roles.placeholder.code')"
            />
          </a-form-item>
          <a-form-item field="description" :label="t('roles.description')" :validate-trigger="['change', 'blur']">
            <a-textarea
              v-model="formData.description"
              :placeholder="t('roles.placeholder.description')"
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
import { getRoles, deleteRole, addRole, updateRole } from '@/api/mock/role'

const { t } = useI18n()

const searchForm = reactive({
  searchQuery: ''
})
const loading = ref(false)
const roles = ref([])
const modalVisible = ref(false)
const modalType = ref('add')
const formRef = ref(null)
const formData = ref({
  name: '',
  code: '',
  description: ''
})

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

const fetchRoles = async () => {
  loading.value = true
  try {
    const res = await getRoles({
      current: pagination.value.current,
      pageSize: pagination.value.pageSize,
      keyword: searchForm.searchQuery
    })
    
    if (res.data.code === 200) {
      roles.value = res.data.data.list
      pagination.value.total = res.data.data.total
    } else {
      Message.error(t('roles.fetchFailed'))
    }
  } catch (error) {
    console.error('Error fetching roles:', error)
    Message.error(t('roles.fetchFailed'))
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.value.current = 1
  fetchRoles()
}

const onPageChange = (page) => {
  pagination.value.current = page
  fetchRoles()
}

// Watch for search query changes
watch(() => searchForm.searchQuery, (newVal) => {
  if (!newVal) {
    handleSearch()
  }
})

// Initial data fetch
onMounted(() => {
  fetchRoles()
})

const resetForm = () => {
  formData.value = {
    name: '',
    code: '',
    description: ''
  }
  if (formRef.value) {
    formRef.value.resetFields()
  }
}

const handleAdd = () => {
  modalType.value = 'add'
  resetForm()
  modalVisible.value = true
}

const handleEdit = (record) => {
  modalType.value = 'edit'
  formData.value = { ...record }
  modalVisible.value = true
}

const handleDelete = (record) => {
  Modal.warning({
    title: t('common.confirm'),
    content: t('roles.confirmDelete'),
    okText: t('common.confirm'),
    cancelText: t('common.cancel'),
    hideCancel: false,
    onOk: async () => {
      try {
        const res = await deleteRole(record.id)
        if (res.data.code === 200) {
          Message.success(t('roles.deleteSuccess'))
          fetchRoles()
        } else {
          Message.error(res.data.message || t('roles.deleteFailed'))
        }
      } catch (error) {
        console.error('Delete role error:', error)
        Message.error(t('roles.deleteFailed'))
      }
    }
  })
}

const handleModalOk = async () => {
  try {
    await formRef.value.validate()
    const api = modalType.value === 'add' ? addRole : updateRole
    const params = modalType.value === 'add' ? [formData.value] : [currentRecord.value.id, formData.value]
    
    const res = await api(...params)
    if (res.data.code === 200) {
      Message.success(modalType.value === 'add' ? t('roles.addSuccess') : t('roles.updateSuccess'))
      modalVisible.value = false
      fetchRoles()
    } else {
      Message.error(res.data.message || (modalType.value === 'add' ? t('roles.addFailed') : t('roles.updateFailed')))
    }
  } catch (error) {
    console.error('Form validation failed:', error)
  }
}

const rules = {
  name: [{ required: true, message: t('roles.validation.nameRequired') }],
  code: [{ required: true, message: t('roles.validation.codeRequired') }],
  description: [{ required: true, message: t('roles.validation.descriptionRequired') }]
}
</script>

<style scoped>
.roles-container {
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