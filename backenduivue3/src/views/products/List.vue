<template>
  <div class="list-container">
    <a-card class="general-card">

      <!-- 搜索表单 -->
      <div class="search-form-wrapper">
        <a-form :model="searchForm" layout="inline" class="search-form">
          <div class="search-form-left">
            <a-form-item field="name" :label="$t('product.list.name')">
              <a-input v-model="searchForm.name" :placeholder="$t('product.list.validation.nameRequired')" />
            </a-form-item>
            <a-form-item field="category_id" :label="$t('product.list.category')">
              <a-select v-model="searchForm.category_id" :placeholder="$t('product.list.validation.categoryRequired')">
                <a-option
                  v-for="item in categoryOptions"
                  :key="item.id"
                  :value="item.id"
                >
                  {{ item.name }}
                </a-option>
              </a-select>
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
              <a-button type="outline" @click="handleExport">
                <template #icon><icon-download /></template>
                {{ $t('common.export') }}
              </a-button>
              <a-button type="primary" @click="handleAdd">
                <template #icon><icon-plus /></template>
                {{ $t('product.list.add') }}
              </a-button>
            </a-space>
          </div>
        </a-form>
      </div>
      
      <a-table
        :data="productList"
        :loading="loading"
        :pagination="pagination"
        @page-change="onPageChange"
      >
        <template #columns>
          <a-table-column title="ID" data-index="id" />
          <a-table-column :title="$t('product.list.name')" data-index="name" />
          <a-table-column :title="$t('product.list.category')" data-index="category_name" />
          <a-table-column :title="$t('product.list.price')" data-index="price">
            <template #cell="{ record }">
              ¥{{ record.price }}
            </template>
          </a-table-column>
          <a-table-column :title="$t('product.list.stock')" data-index="stock" />
          <a-table-column :title="$t('common.status')">
            <template #cell="{ record }">
              <a-tag :color="record.status === 1 ? 'green' : 'red'">
                {{ record.status === 1 ? $t('product.list.onSale') : $t('product.list.offSale') }}
              </a-tag>
            </template>
          </a-table-column>
          <a-table-column :title="$t('common.operations')">
            <template #cell="{ record }">
              <a-space>
                <a-button type="text" size="small" @click="handleView(record)">
                  <template #icon><icon-eye /></template>
                  {{ $t('product.list.view') }}
                </a-button>
                <a-button type="text" size="small" @click="handleEdit(record)">
                  <template #icon><icon-edit /></template>
                  {{ $t('common.edit') }}
                </a-button>
                <a-popconfirm
                  :content="$t('product.list.deleteConfirm')"
                  @ok="handleDelete(record)"
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

    <!-- 添加/编辑产品的抽屉 -->
    <a-drawer
      :visible="drawerVisible"
      :width="600"
      @cancel="closeDrawer"
      @ok="handleSubmit"
      unmountOnClose
    >
      <template #title>
        {{ drawerTitle }}
      </template>
      <a-form ref="formRef" :model="formData" :rules="rules">
        <a-form-item field="name" :label="$t('product.list.name')" validate-trigger="blur">
          <a-input v-model="formData.name" :placeholder="$t('product.list.validation.nameRequired')" />
        </a-form-item>
        <a-form-item field="category_id" :label="$t('product.list.category')" validate-trigger="change">
          <a-select v-model="formData.category_id" :placeholder="$t('product.list.validation.categoryRequired')">
            <a-option
              v-for="item in categoryOptions"
              :key="item.id"
              :value="item.id"
            >
              {{ item.name }}
            </a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="price" :label="$t('product.list.price')" validate-trigger="blur">
          <a-input-number
            v-model="formData.price"
            :placeholder="$t('product.list.validation.priceRequired')"
            :precision="2"
            :step="0.1"
            :min="0"
          />
        </a-form-item>
        <a-form-item field="stock" :label="$t('product.list.stock')" validate-trigger="blur">
          <a-input-number
            v-model="formData.stock"
            :placeholder="$t('product.list.validation.stockRequired')"
            :min="0"
          />
        </a-form-item>
        <a-form-item field="description" :label="$t('product.list.description')">
          <a-textarea
            v-model="formData.description"
            :placeholder="$t('product.list.validation.descriptionMaxLength')"
            :auto-size="{ minRows: 3, maxRows: 5 }"
          />
        </a-form-item>
        <a-form-item field="status" :label="$t('common.status')">
          <a-radio-group v-model="formData.status">
            <a-radio :value="1">{{ $t('product.list.onSale') }}</a-radio>
            <a-radio :value="0">{{ $t('product.list.offSale') }}</a-radio>
          </a-radio-group>
        </a-form-item>
      </a-form>
    </a-drawer>

    <!-- 产品详情抽屉 -->
    <a-drawer
      :visible="detailVisible"
      :width="600"
      @cancel="closeDetail"
      unmountOnClose
    >
      <template #title>
        {{ $t('product.list.view') }}
      </template>
      <a-descriptions :data="detailData" layout="inline-horizontal" />
    </a-drawer>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { Message } from '@arco-design/web-vue'
import {
  IconPlus,
  IconEdit,
  IconDelete,
  IconSearch,
  IconRefresh,
  IconDownload,
  IconEye
} from '@arco-design/web-vue/es/icon'
import { useI18n } from 'vue-i18n'
import { getProducts, addProduct, updateProduct, deleteProduct, exportProducts } from '@/api/mock/product'

const { t } = useI18n()

const loading = ref(false)
const drawerVisible = ref(false)
const detailVisible = ref(false)
const drawerTitle = ref('')
const formRef = ref(null)
const productList = ref([])
const categoryOptions = ref([])
const currentId = ref(null)
const detailData = ref([])

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

const searchForm = reactive({
  name: '',
  category_id: undefined,
})

const formData = reactive({
  name: '',
  category_id: undefined,
  price: 0,
  stock: 0,
  description: '',
  status: 1,
})

const rules = {
  name: [
    { required: true, message: t('product.list.validation.nameRequired') },
    { maxLength: 100, message: t('product.list.validation.nameMaxLength') }
  ],
  category_id: [
    { required: true, message: t('product.list.validation.categoryRequired') }
  ],
  price: [
    { required: true, message: t('product.list.validation.priceRequired') }
  ],
  stock: [
    { required: true, message: t('product.list.validation.stockRequired') }
  ],
  description: [
    { maxLength: 500, message: t('product.list.validation.descriptionMaxLength') }
  ]
}

// 获取产品列表
const fetchProductList = async () => {
  loading.value = true
  try {
    const res = await getProducts({
      current: pagination.current,
      pageSize: pagination.pageSize,
      name: searchForm.name,
      category_id: searchForm.category_id
    })
    
    if (res.data.code === 200) {
      productList.value = res.data.data.list
      pagination.total = res.data.data.total
    } else {
      Message.error(t('common.fetchFailed'))
    }
  } catch (error) {
    console.error('获取产品列表失败:', error)
    Message.error(t('common.fetchFailed'))
  } finally {
    loading.value = false
  }
}

// 获取分类选项
const fetchCategoryOptions = async () => {
  try {
    // 模拟API调用延迟
    await new Promise(resolve => setTimeout(resolve, 500))
    categoryOptions.value = mockCategories
  } catch (error) {
    Message.error('获取分类列表失败')
  }
}

// 搜索
const handleSearch = () => {
  pagination.current = 1
  fetchProductList()
}

// 重置搜索
const handleReset = () => {
  searchForm.name = ''
  searchForm.category_id = undefined
  handleSearch()
}

// 添加产品
const handleAdd = () => {
  drawerTitle.value = t('product.list.add')
  currentId.value = null
  Object.assign(formData, {
    name: '',
    category_id: undefined,
    price: 0,
    stock: 0,
    description: '',
    status: 1,
  })
  drawerVisible.value = true
}

// 编辑产品
const handleEdit = (record) => {
  drawerTitle.value = t('product.list.edit')
  currentId.value = record.id
  Object.assign(formData, record)
  drawerVisible.value = true
}

// 查看产品详情
const handleView = (record) => {
  detailData.value = [
    { label: t('product.list.name'), value: record.name },
    { label: t('product.list.category'), value: record.category_name },
    { label: t('product.list.price'), value: `¥${record.price}` },
    { label: t('product.list.stock'), value: record.stock },
    { label: t('common.status'), value: record.status === 1 ? t('product.list.onSale') : t('product.list.offSale') },
    { label: t('product.list.description'), value: record.description },
    { label: t('product.list.createdAt'), value: record.created_at },
    { label: t('product.list.updatedAt'), value: record.updated_at },
  ]
  detailVisible.value = true
}

// 删除产品
const handleDelete = async (record) => {
  try {
    const res = await deleteProduct(record.id)
    if (res.data.code === 200) {
      Message.success(t('product.list.deleteSuccess'))
      // 如果当前页只有一条数据，且不是第一页，则跳转到上一页
      if (productList.value.length === 1 && pagination.current > 1) {
        pagination.current--
      }
      fetchProductList()
    } else {
      Message.error(t('product.list.deleteFailed'))
    }
  } catch (error) {
    console.error('删除产品失败:', error)
    Message.error(t('product.list.deleteFailed'))
  }
}

// 导出产品
const handleExport = async () => {
  try {
    const res = await exportProducts()
    if (res.data.code === 200) {
      Message.success(t('common.exportSuccess'))
      // 这里可以处理导出的数据，比如下载Excel文件
    } else {
      Message.error(t('common.exportFailed'))
    }
  } catch (error) {
    console.error('导出产品失败:', error)
    Message.error(t('common.exportFailed'))
  }
}

// 提交表单
const handleSubmit = async () => {
  try {
    await formRef.value.validate()
    const apiMethod = currentId.value ? updateProduct : addProduct
    const res = await (currentId.value ? 
      apiMethod(currentId.value, formData) : 
      apiMethod(formData)
    )
    
    if (res.data.code === 200) {
      Message.success(currentId.value ? t('common.updateSuccess') : t('common.addSuccess'))
      closeDrawer()
      fetchProductList()
    } else {
      Message.error(currentId.value ? t('common.updateFailed') : t('common.addFailed'))
    }
  } catch (error) {
    console.error('保存产品失败:', error)
    Message.error(currentId.value ? t('common.updateFailed') : t('common.addFailed'))
  }
}

// 关闭抽屉
const closeDrawer = () => {
  drawerVisible.value = false
  formRef.value?.resetFields()
}

// 关闭详情抽屉
const closeDetail = () => {
  detailVisible.value = false
}

// 页码改变
const onPageChange = (current) => {
  pagination.current = current
  fetchProductList()
}

// 初始化
fetchProductList()
fetchCategoryOptions()
</script>

<style scoped>
.list-container {
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
  padding:14px 14px 5px;
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
  
  :deep(.arco-select) {
    width: 100%;
  }
}
</style> 