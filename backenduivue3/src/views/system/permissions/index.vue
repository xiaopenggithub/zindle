<template>
  <div class="permissions-container">
    <a-card class="general-card">
      <!-- 搜索区域 -->
      <div class="search-form-wrapper">
        <a-form layout="inline" class="search-form" :model="searchForm">
          <div class="search-form-left">
            <a-form-item field="searchQuery">
              <a-input-search
                v-model="searchForm.searchQuery"
                :placeholder="t('permissions.placeholder.search')"
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
                {{ t('permissions.addPermission') }}
              </a-button>
            </a-space>
          </div>
        </a-form>
      </div>

      <a-table
        :data="permissions"
        :loading="loading"
        :pagination="pagination"
        :empty="t('permissions.table.noData')"
        :bordered="true"
        :stripe="true"
        @page-change="onPageChange"
      >
        <template #columns>
          <a-table-column :title="t('permissions.name')" data-index="name" />
          <a-table-column :title="t('permissions.code')" data-index="code" />
          <a-table-column :title="t('permissions.type')" data-index="type">
            <template #cell="{ record }">
              <a-tag :color="getTypeColor(record.type)">
                {{ t(`permissions.types.${record.type.toLowerCase()}`) }}
              </a-tag>
            </template>
          </a-table-column>
          <a-table-column :title="t('permissions.description')" data-index="description" />
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

      <!-- 新增/编辑权限模态框 -->
      <a-modal
        v-model:visible="modalVisible"
        :title="t(modalType === 'add' ? 'permissions.addPermission' : 'permissions.editPermission')"
        @ok="handleModalOk"
        @cancel="modalVisible = false"
      >
        <a-form
          ref="formRef"
          :model="formData"
          :rules="rules"
          @submit="handleModalOk"
        >
          <a-form-item field="name" :label="t('permissions.name')" :validate-trigger="['change', 'blur']">
            <a-input
              v-model="formData.name"
              :placeholder="t('permissions.placeholder.name')"
            />
          </a-form-item>
          <a-form-item field="code" :label="t('permissions.code')" :validate-trigger="['change', 'blur']">
            <a-input
              v-model="formData.code"
              :placeholder="t('permissions.placeholder.code')"
            />
          </a-form-item>
          <a-form-item field="type" :label="t('permissions.type')" :validate-trigger="['change', 'blur']">
            <a-select
              v-model="formData.type"
              :placeholder="t('permissions.placeholder.type')"
            >
              <a-option value="Menu">{{ t('permissions.types.menu') }}</a-option>
              <a-option value="Button">{{ t('permissions.types.button') }}</a-option>
              <a-option value="API">{{ t('permissions.types.api') }}</a-option>
            </a-select>
          </a-form-item>
          <a-form-item field="description" :label="t('permissions.description')" :validate-trigger="['change', 'blur']">
            <a-textarea
              v-model="formData.description"
              :placeholder="t('permissions.placeholder.description')"
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
import { getPermissions, deletePermission, addPermission, updatePermission } from '@/api/mock/permission'

const { t } = useI18n()

const searchForm = reactive({
  searchQuery: ''
})
const loading = ref(false)
const permissions = ref([])
const modalVisible = ref(false)
const modalType = ref('add')
const formRef = ref(null)
const formData = ref({
  name: '',
  code: '',
  type: '',
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

const getTypeColor = (type) => {
  const colors = {
    Menu: 'blue',
    Action: 'green',
    Data: 'purple'
  }
  return colors[type] || 'blue'
}

const fetchPermissions = async () => {
  loading.value = true
  try {
    const res = await getPermissions({
      current: pagination.value.current,
      pageSize: pagination.value.pageSize,
      keyword: searchForm.searchQuery
    })
    
    if (res.data.code === 200) {
      permissions.value = res.data.data.list
      pagination.value.total = res.data.data.total
    } else {
      Message.error(t('permissions.fetchFailed'))
    }
  } catch (error) {
    console.error('Error fetching permissions:', error)
    Message.error(t('permissions.fetchFailed'))
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.value.current = 1
  fetchPermissions()
}

const onPageChange = (page) => {
  pagination.value.current = page
  fetchPermissions()
}

// Watch for search query changes
watch(() => searchForm.searchQuery, (newVal) => {
  if (!newVal) {
    handleSearch()
  }
})

// Initial data fetch
onMounted(() => {
  fetchPermissions()
})

const resetForm = () => {
  formData.value = {
    name: '',
    code: '',
    type: '',
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
    content: t('permissions.confirmDelete'),
    okText: t('common.confirm'),
    cancelText: t('common.cancel'),
    hideCancel: false,
    onOk: async () => {
      try {
        const res = await deletePermission(record.id)
        if (res.data.code === 200) {
          Message.success(t('permissions.deleteSuccess'))
          fetchPermissions()
        } else {
          Message.error(res.data.message || t('permissions.deleteFailed'))
        }
      } catch (error) {
        console.error('Delete permission error:', error)
        Message.error(t('permissions.deleteFailed'))
      }
    }
  })
}

const handleModalOk = async () => {
  try {
    await formRef.value.validate()
    const api = modalType.value === 'add' ? addPermission : updatePermission
    const params = modalType.value === 'add' ? [formData.value] : [currentRecord.value.id, formData.value]
    
    const res = await api(...params)
    if (res.data.code === 200) {
      Message.success(modalType.value === 'add' ? t('permissions.addSuccess') : t('permissions.updateSuccess'))
      modalVisible.value = false
      fetchPermissions()
    } else {
      Message.error(res.data.message || (modalType.value === 'add' ? t('permissions.addFailed') : t('permissions.updateFailed')))
    }
  } catch (error) {
    console.error('Form validation failed:', error)
  }
}

const rules = {
  name: [{ required: true, message: t('permissions.validation.nameRequired') }],
  code: [{ required: true, message: t('permissions.validation.codeRequired') }],
  type: [{ required: true, message: t('permissions.validation.typeRequired') }],
  description: [{ required: true, message: t('permissions.validation.descriptionRequired') }]
}
</script>

<style scoped>
.permissions-container {
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