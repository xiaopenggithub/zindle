<template>
  <div class="category-container">
    <a-card class="general-card">
      <!-- 搜索区域 -->
      <div class="search-form-wrapper">
        <a-form :model="searchForm" layout="inline" class="search-form">
          <div class="search-form-left">
            <a-form-item field="name" :label="$t('product.category.name')" hide-label>
              <a-input v-model="searchForm.name" :placeholder="$t('product.category.validation.nameRequired')" allow-clear>
                <template #prefix>
                  <icon-search />
                </template>
              </a-input>
            </a-form-item>
            <!-- 新增：创建时间范围选择 -->
            <a-form-item field="createdAt" :label="$t('product.category.createdAt')" hide-label>
              <a-range-picker
                v-model="searchForm.createdAt"
                :placeholder="[$t('common.startDate'), $t('common.endDate')]"
                style="width: 240px"
              />
            </a-form-item>
            <a-form-item>
              <a-space>
                <a-button type="primary" @click="handleSearch">
                  <template #icon><icon-search /></template>
                  {{ $t('common.search') }}
                </a-button>
                <a-button @click="handleReset">
                  <template #icon><icon-refresh /></template>
                  {{ $t('common.reset') }}
                </a-button>
              </a-space>
            </a-form-item>
          </div>
          <div class="search-form-right">
            <a-space>
              <a-button type="primary" @click="handleAdd">
                <template #icon><icon-plus /></template>
                {{ $t('product.category.add') }}
              </a-button>
            </a-space>
          </div>
        </a-form>
      </div>
      
      <a-table
        :data="categoryList"
        :loading="loading"
        :pagination="pagination"
        @page-change="onPageChange"
        :bordered="true"
        :stripe="true"
        :hoverable="true"
      >
        <template #columns>
          <a-table-column title="ID" data-index="id" :width="80" align="center" />
          <a-table-column :title="$t('product.category.name')" data-index="name" />
          <a-table-column :title="$t('product.category.description')" data-index="description" :width="300" />
          <a-table-column :title="$t('product.category.createdAt')" data-index="created_at" :width="180" align="center" />
          <a-table-column :title="$t('common.operations')" align="center" :width="160" fixed="right">
            <template #cell="{ record }">
              <a-space>
                <a-button type="text" size="small" @click="handleEdit(record)">
                  <template #icon><icon-edit /></template>
                  {{ $t('common.edit') }}
                </a-button>
                <a-divider direction="vertical" />
                <a-popconfirm
                  :content="$t('product.category.deleteConfirm')"
                  @ok="handleDelete(record)"
                  position="left"
                >
                  <a-button type="text" status="danger" size="small">
                    <template #icon><icon-delete /></template>
                    {{ $t('common.delete') }}
                  </a-button>
                </a-popconfirm>
              </a-space>
            </template>
          </a-table-column>
        </template>
      </a-table>
    </a-card>

    <!-- 添加/编辑分类的抽屉 -->
    <a-drawer
      :visible="drawerVisible"
      :width="500"
      @cancel="closeDrawer"
      @ok="handleSubmit"
      unmountOnClose
    >
      <template #title>
        {{ drawerTitle }}
      </template>
      <a-form ref="formRef" :model="formData" :rules="rules" class="drawer-form">
        <a-form-item field="name" :label="$t('product.category.name')" validate-trigger="blur">
          <a-input v-model="formData.name" :placeholder="$t('product.category.validation.nameRequired')" allow-clear />
        </a-form-item>
        <a-form-item field="description" :label="$t('product.category.description')">
          <a-textarea
            v-model="formData.description"
            :placeholder="$t('product.category.validation.descriptionMaxLength')"
            :auto-size="{ minRows: 3, maxRows: 5 }"
            allow-clear
          />
        </a-form-item>
      </a-form>
    </a-drawer>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { Message } from '@arco-design/web-vue'
import { IconPlus, IconEdit, IconDelete, IconSearch, IconRefresh } from '@arco-design/web-vue/es/icon'
import { useI18n } from 'vue-i18n'
import { getCategories, addCategory, updateCategory, deleteCategory } from '@/api/mock/category'

const { t } = useI18n()
const loading = ref(false)
const drawerVisible = ref(false)
const drawerTitle = ref('')
const formRef = ref(null)
const categoryList = ref([])
const currentId = ref(null)

const searchForm = reactive({
  name: '',
  createdAt: []      // 新增：用于存放日期范围
})

const pagination = reactive({
  total: 0,
  current: 1,
  pageSize: 10,
  showTotal: true,
  showJumper: true,
  showPageSize: true,
  pageSizeOptions: [10, 20, 50, 100],
  baseSize: 10
})

const formData = reactive({
  name: '',
  description: '',
})

const rules = {
  name: [
    { required: true, message: t('product.category.validation.nameRequired') },
    { maxLength: 50, message: t('product.category.validation.nameMaxLength') }
  ],
  description: [
    { maxLength: 200, message: t('product.category.validation.descriptionMaxLength') }
  ]
}

// 获取分类列表
const fetchCategoryList = async () => {
  loading.value = true
  try {
    const params = {
      current: pagination.current,
      pageSize: pagination.pageSize,
      name: searchForm.name
    }
    // 新增：如果有选择日期范围，则加入参数
    if (searchForm.createdAt && searchForm.createdAt.length === 2) {
      params.startDate = searchForm.createdAt[0]
      params.endDate   = searchForm.createdAt[1]
    }

    const res = await getCategories(params)
    
    if (res.data.code === 200) {
      categoryList.value = res.data.data.list
      pagination.total = res.data.data.total
    } else {
      Message.error(t('common.fetchFailed'))
    }
  } catch (error) {
    console.error('获取分类列表失败:', error)
    Message.error(t('common.fetchFailed'))
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  pagination.current = 1
  fetchCategoryList()
}

// 重置搜索
const handleReset = () => {
  searchForm.name = ''
  searchForm.createdAt = []   // 新增：清空日期
  handleSearch()
}

// 添加分类
const handleAdd = () => {
  drawerTitle.value = t('product.category.add')
  currentId.value = null
  formData.name = ''
  formData.description = ''
  drawerVisible.value = true
}

// 编辑分类
const handleEdit = (record) => {
  drawerTitle.value = t('product.category.edit')
  currentId.value = record.id
  formData.name = record.name
  formData.description = record.description
  drawerVisible.value = true
}

// 删除分类
const handleDelete = async (record) => {
  try {
    const res = await deleteCategory(record.id)
    if (res.data.code === 200) {
      Message.success(t('product.category.deleteSuccess'))
      // 如果当前页只有一条数据，且不是第一页，则跳转到上一页
      if (categoryList.value.length === 1 && pagination.current > 1) {
        pagination.current--
      }
      fetchCategoryList()
    } else {
      Message.error(t('product.category.deleteFailed'))
    }
  } catch (error) {
    console.error('删除分类失败:', error)
    Message.error(t('product.category.deleteFailed'))
  }
}

// 提交表单
const handleSubmit = async () => {
  try {
    await formRef.value.validate()
    const apiMethod = currentId.value ? updateCategory : addCategory
    const res = await (currentId.value ? 
      apiMethod(currentId.value, formData) : 
      apiMethod(formData)
    )
    
    if (res.data.code === 200) {
      Message.success(currentId.value ? t('common.updateSuccess') : t('common.addSuccess'))
      closeDrawer()
      fetchCategoryList()
    } else {
      Message.error(currentId.value ? t('common.updateFailed') : t('common.addFailed'))
    }
  } catch (error) {
    console.error('保存分类失败:', error)
    Message.error(currentId.value ? t('common.updateFailed') : t('common.addFailed'))
  }
}

// 关闭抽屉
const closeDrawer = () => {
  drawerVisible.value = false
  formRef.value?.resetFields()
}

// 页码改变
const onPageChange = (current) => {
  pagination.current = current
  fetchCategoryList()
}

// 初始化
fetchCategoryList()
</script>

<style scoped>
.category-container {
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

.drawer-form {
  padding: 16px;
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