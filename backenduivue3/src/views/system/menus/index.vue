<template>
  <div class="menus-container">
    <a-card class="general-card">
      <!-- 搜索区域 -->
      <div class="search-form-wrapper">
        <a-form layout="inline" class="search-form" :model="searchForm">
          <div class="search-form-left">
            <a-form-item field="searchQuery">
              <a-input-search
                v-model="searchForm.searchQuery"
                :placeholder="t('menus.placeholder.search')"
                :style="{ width: '220px' }"
                search-button
                @search="handleSearch"
              />
            </a-form-item>
            <a-form-item>
              <a-button @click="resetSearch">
                {{ t('common.reset') }}
              </a-button>
            </a-form-item>
          </div>
          <div class="search-form-right">
            <a-space>
              <a-button type="primary" @click="handleRefresh">
                <template #icon>
                  <icon-refresh />
                </template>
                {{ t('common.refresh') }}
              </a-button>
              <a-button type="primary" @click="handleAdd">
                <template #icon>
                  <icon-plus />
                </template>
                {{ t('menus.addMenu') }}
              </a-button>
            </a-space>
          </div>
        </a-form>
      </div>

      <a-table
        :data="menus"
        :loading="loading"
        :pagination="false"
        :tree-props="{ children: 'children' }"
        :bordered="true"
        :stripe="true"
      >
        <template #columns>
          <a-table-column :title="t('menus.name')" data-index="name" />
          <a-table-column :title="t('menus.path')" data-index="path" />
          <a-table-column :title="t('menus.component')" data-index="component" />
          <a-table-column :title="t('menus.icon')" data-index="icon" :width="80">
            <template #cell="{ record }">
              <component :is="record.icon" v-if="record.icon" />
            </template>
          </a-table-column>
          <a-table-column :title="t('menus.sort')" data-index="sort" :width="80" align="center" />
          <a-table-column :title="t('common.status')" data-index="status" :width="100" align="center">
            <template #cell="{ record }">
              <a-tag :color="record.status === 'active' ? 'green' : 'red'">
                {{ record.status === 'active' ? t('common.active') : t('common.inactive') }}
              </a-tag>
            </template>
          </a-table-column>
          <a-table-column :title="t('common.operations')" align="center" :width="120">
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
        <template #empty>
          <a-empty :description="t('menus.table.noData')" />
        </template>
      </a-table>
    </a-card>

    <!-- Add/Edit Menu Modal -->
    <a-modal
      v-model:visible="modalVisible"
      :title="modalType === 'add' ? t('menus.addMenu') : t('menus.editMenu')"
      @ok="handleModalOk"
      @cancel="handleModalCancel"
    >
      <a-form ref="formRef" :model="formData" :rules="rules">
        <a-form-item field="name" :label="t('menus.name')" validate-trigger="blur">
          <a-input v-model="formData.name" :placeholder="t('menus.placeholder.name')" />
        </a-form-item>
        <a-form-item field="path" :label="t('menus.path')" validate-trigger="blur">
          <a-input v-model="formData.path" :placeholder="t('menus.placeholder.path')" />
        </a-form-item>
        <a-form-item field="component" :label="t('menus.component')" validate-trigger="blur">
          <a-input v-model="formData.component" :placeholder="t('menus.placeholder.component')" />
        </a-form-item>
        <a-form-item field="icon" :label="t('menus.icon')" validate-trigger="blur">
          <a-input v-model="formData.icon" :placeholder="t('menus.placeholder.icon')" />
        </a-form-item>
        <a-form-item field="sort" :label="t('menus.sort')" validate-trigger="blur">
          <a-input-number
            v-model="formData.sort"
            :min="0"
            :max="999"
            :placeholder="t('menus.placeholder.sort')"
          />
        </a-form-item>
        <a-form-item field="status" :label="t('common.status')" validate-trigger="blur">
          <a-select v-model="formData.status">
            <a-option value="active">{{ t('common.active') }}</a-option>
            <a-option value="inactive">{{ t('common.inactive') }}</a-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, watch, onMounted, reactive } from 'vue'
import { IconPlus, IconEdit, IconDelete, IconRefresh } from '@arco-design/web-vue/es/icon'
import { Message, Modal } from '@arco-design/web-vue'
import { useI18n } from 'vue-i18n'
import { getMenus, deleteMenu, addMenu, updateMenu } from '@/api/mock/menu'

const { t } = useI18n()

const searchForm = reactive({
  searchQuery: ''
})
const loading = ref(false)
const menus = ref([])
const modalVisible = ref(false)
const modalType = ref('add')
const formRef = ref(null)
const currentRecord = ref(null)

const formData = ref({
  name: '',
  path: '',
  component: '',
  icon: '',
  sort: 0,
  status: 'active'
})

const rules = {
  name: [{ required: true, message: t('menus.validation.nameRequired') }],
  path: [{ required: true, message: t('menus.validation.pathRequired') }],
  component: [{ required: true, message: t('menus.validation.componentRequired') }],
  sort: [{ required: true, message: t('menus.validation.sortRequired') }],
  status: [{ required: true, message: t('menus.validation.statusRequired') }]
}

const fetchMenus = async () => {
  loading.value = true
  try {
    const res = await getMenus({
      keyword: searchForm.searchQuery
    })
    
    if (res.data.code === 200) {
      menus.value = res.data.data
    } else {
      Message.error(t('menus.fetchFailed'))
    }
  } catch (error) {
    console.error('Error fetching menus:', error)
    Message.error(t('menus.fetchFailed'))
  } finally {
    loading.value = false
  }
}

const resetSearch = () => {
  searchForm.searchQuery = ''
  fetchMenus()
}

const handleRefresh = () => {
  fetchMenus()
}

const handleSearch = () => {
  fetchMenus()
}

const handleAdd = () => {
  modalType.value = 'add'
  formData.value = {
    name: '',
    path: '',
    component: '',
    icon: '',
    sort: 0,
    status: 'active'
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
    content: t('menus.confirmDelete'),
    okText: t('common.confirm'),
    cancelText: t('common.cancel'),
    hideCancel: false,
    onOk: async () => {
      try {
        const res = await deleteMenu(record.id)
        if (res.data.code === 200) {
          Message.success(t('menus.deleteSuccess'))
          fetchMenus()
        } else {
          Message.error(res.data.message || t('menus.deleteFailed'))
        }
      } catch (error) {
        console.error('Delete menu error:', error)
        Message.error(t('menus.deleteFailed'))
      }
    }
  })
}

const handleModalOk = async () => {
  try {
    await formRef.value.validate()
    const api = modalType.value === 'add' ? addMenu : updateMenu
    const params = modalType.value === 'add' ? [formData.value] : [currentRecord.value.id, formData.value]
    
    const res = await api(...params)
    if (res.data.code === 200) {
      Message.success(modalType.value === 'add' ? t('menus.addSuccess') : t('menus.updateSuccess'))
      modalVisible.value = false
      fetchMenus()
    } else {
      Message.error(res.data.message || (modalType.value === 'add' ? t('menus.addFailed') : t('menus.updateFailed')))
    }
  } catch (error) {
    console.error('Form validation failed:', error)
  }
}

const handleModalCancel = () => {
  modalVisible.value = false
  formRef.value?.resetFields()
}

// Watch for search query changes
watch(() => searchForm.searchQuery, (newVal) => {
  if (!newVal) {
    fetchMenus()
  }
})

// Initial data fetch
onMounted(() => {
  fetchMenus()
})
</script>

<style scoped>
.menus-container {
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

:deep(.arco-input-wrapper) {
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