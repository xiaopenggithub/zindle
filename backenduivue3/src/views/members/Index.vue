<template>
  <div class="members-container" v-if="!isUnmounted">
    <a-card class="general-card">
      <div class="search-form-wrapper">
        <a-form layout="inline" class="search-form" :model="searchForm">
          <div class="search-form-left">
            <a-form-item>
              <a-input-search
                v-model="searchForm.keyword"
                :placeholder="t('members.placeholder.search')"
                :style="{ width: '300px' }"
                @search="handleSearch"
              />
            </a-form-item>
            <a-form-item>
              <a-space>
                <a-button type="primary" @click="handleAdd">
                  <template #icon>
                    <icon-plus />
                  </template>
                  {{ t('members.addMember') }}
                </a-button>
                <a-button type="primary" status="success" @click="handleExport">
                  <template #icon>
                    <icon-download />
                  </template>
                  {{ t('members.export') }}
                </a-button>
              </a-space>
            </a-form-item>
          </div>
        </a-form>
      </div>

      <a-table
        :loading="loading"
        :data="members"
        :pagination="pagination"
        @page-change="onPageChange"
        @page-size-change="onPageSizeChange"
        :bordered="{ cell: true }"
        :stripe="true"
        size="medium"
      >
        <template #columns>
          <a-table-column title="ID" data-index="id" :width="80" />
          <a-table-column :title="t('members.username')" data-index="username" />
          <a-table-column :title="t('members.nickname')" data-index="nickname" />
          <a-table-column :title="t('members.mobile')" data-index="mobile" />
          <a-table-column :title="t('members.email')" data-index="email" />
          <a-table-column :title="t('members.status.label')" data-index="status">
            <template #cell="{ record }">
              <a-tag :color="record.status === 1 ? 'green' : 'red'">
                {{ record.status === 1 ? t('members.status.active') : t('members.status.inactive') }}
              </a-tag>
            </template>
          </a-table-column>
          <a-table-column :title="t('members.registerTime')" data-index="registerTime" />
          <a-table-column :title="t('common.operations')" align="center" fixed="right" :width="200">
            <template #cell="{ record }">
              <a-space>
                <a-button type="text" size="small" @click="handleView(record)">
                  {{ t('common.view') }}
                </a-button>
                <a-button type="text" size="small" @click="handleEdit(record)">
                  {{ t('common.edit') }}
                </a-button>
                <a-popconfirm
                  :content="t('members.confirmDelete')"
                  @ok="handleDelete(record)"
                >
                  <a-button type="text" status="danger" size="small">
                    {{ t('common.delete') }}
                  </a-button>
                </a-popconfirm>
              </a-space>
            </template>
          </a-table-column>
        </template>
      </a-table>

      <!-- 添加/编辑会员弹窗 -->
      <a-modal
        v-if="modalVisible"
        v-model:visible="modalVisible"
        :title="modalTitle"
        @ok="handleModalOk"
        @cancel="handleModalCancel"
        :unmount-on-close="true"
        :width="600"
      >
        <a-form ref="formRef" :model="form" :rules="rules" layout="horizontal">
          <a-form-item field="username" :label="t('members.username')" :rules="[{ required: true }]">
            <a-input v-model="form.username" :placeholder="t('members.placeholder.username')" />
          </a-form-item>
          <a-form-item field="nickname" :label="t('members.nickname')" :rules="[{ required: true }]">
            <a-input v-model="form.nickname" :placeholder="t('members.placeholder.nickname')" />
          </a-form-item>
          <a-form-item field="mobile" :label="t('members.mobile')" :rules="[{ required: true }]">
            <a-input v-model="form.mobile" :placeholder="t('members.placeholder.mobile')" />
          </a-form-item>
          <a-form-item field="email" :label="t('members.email')" :rules="[{ required: true, type: 'email' }]">
            <a-input v-model="form.email" :placeholder="t('members.placeholder.email')" />
          </a-form-item>
          <a-form-item field="status" :label="t('members.status.label')" :rules="[{ required: true }]">
            <a-radio-group v-model="form.status">
              <a-radio :value="1">{{ t('members.status.active') }}</a-radio>
              <a-radio :value="0">{{ t('members.status.inactive') }}</a-radio>
            </a-radio-group>
          </a-form-item>
        </a-form>
      </a-modal>

      <!-- 查看会员详情弹窗 -->
      <a-modal
        :width="800"
        v-if="detailVisible"
        v-model:visible="detailVisible"
        :title="t('members.memberDetail')"
        @cancel="detailVisible = false"
        :unmount-on-close="true"
      >
        <a-descriptions :data="memberDetailData" layout="vertical" bordered />
      </a-modal>
    </a-card>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onBeforeUnmount } from 'vue'
import { useI18n } from 'vue-i18n'
import { Message } from '@arco-design/web-vue'
import { IconPlus, IconDownload } from '@arco-design/web-vue/es/icon'
import { getMembers, addMember, updateMember, deleteMember, exportMembers } from '@/api/mock/members'

const { t } = useI18n()
const loading = ref(false)
const members = ref([])
const modalVisible = ref(false)
const detailVisible = ref(false)
const isEdit = ref(false)
const memberDetail = ref({})
const isUnmounted = ref(false)
const searchQuery = ref('')
const currentId = ref(null)

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

const form = reactive({
  username: '',
  nickname: '',
  mobile: '',
  email: '',
  status: 1
})

const rules = {
  username: [
    { required: true, message: t('members.rules.username') }
  ],
  nickname: [
    { required: true, message: t('members.rules.nickname') }
  ],
  mobile: [
    { required: true, message: t('members.rules.mobile') }
  ],
  email: [
    { required: true, type: 'email', message: t('members.rules.email') }
  ],
  status: [
    { required: true, message: t('members.rules.status') }
  ]
}

const formRef = ref(null)

const modalTitle = computed(() => {
  return isEdit.value ? t('members.editMember') : t('members.addMember')
})

// 搜索表单数据
const searchForm = reactive({
  keyword: ''
})

// 获取会员列表
const fetchMembers = async () => {
  if (isUnmounted.value) return
  loading.value = true
  try {
    const res = await getMembers({
      current: pagination.value.current,
      pageSize: pagination.value.pageSize,
      searchQuery: searchQuery.value
    })
    
    if (!isUnmounted.value && res.data.code === 200) {
      members.value = res.data.data.list
      pagination.value.total = res.data.data.total
    }
  } catch (error) {
    if (!isUnmounted.value) {
      console.error('Fetch members error:', error)
      Message.error(t('members.fetchFailed'))
    }
  } finally {
    if (!isUnmounted.value) {
      loading.value = false
    }
  }
}

// 修改搜索处理函数
const handleSearch = (value) => {
  searchQuery.value = searchForm.keyword
  pagination.value.current = 1
  fetchMembers()
}

// 添加会员
const handleAdd = () => {
  isEdit.value = false
  form.username = ''
  form.nickname = ''
  form.mobile = ''
  form.email = ''
  form.status = 1
  modalVisible.value = true
}

// 编辑会员
const handleEdit = (record) => {
  isEdit.value = true
  currentId.value = record.id
  form.username = record.username
  form.nickname = record.nickname
  form.mobile = record.mobile
  form.email = record.email
  form.status = record.status
  modalVisible.value = true
}

// 会员详情数据转换
const memberDetailData = computed(() => {
  if (!memberDetail.value) return []
  
  return [
    {
      label: t('members.username'),
      value: memberDetail.value.username
    },
    {
      label: t('members.nickname'),
      value: memberDetail.value.nickname
    },
    {
      label: t('members.mobile'),
      value: memberDetail.value.mobile
    },
    {
      label: t('members.email'),
      value: memberDetail.value.email
    },
    {
      label: t('members.status'),
      value: memberDetail.value.status === 1 ? t('members.status.active') : t('members.status.inactive')
    },
    {
      label: t('members.registerTime'),
      value: memberDetail.value.registerTime
    }
  ]
})

// 查看会员详情
const handleView = (record) => {
  memberDetail.value = { ...record }
  detailVisible.value = true
}

// 删除会员
const handleDelete = async (record) => {
  if (isUnmounted.value || !record.id) return
  try {
    const res = await deleteMember(record.id)
    if (!isUnmounted.value && res.data.code === 200) {
      Message.success(t('members.deleteSuccess'))
      // 如果当前页只有一条数据，且不是第一页，则跳转到上一页
      if (members.value.length === 1 && pagination.value.current > 1) {
        pagination.value.current--
      }
      fetchMembers()
    }
  } catch (error) {
    if (!isUnmounted.value) {
      console.error('Delete member error:', error)
      Message.error(t('members.deleteFailed'))
    }
  }
}

// 导出会员
const handleExport = async () => {
  try {
    const res = await exportMembers()
    if (res.data.code === 200) {
      Message.success(t('members.exportSuccess'))
      // 这里可以处理导出的数据，比如下载Excel文件
    }
  } catch (error) {
    Message.error(t('members.exportFailed'))
  }
}

// 弹窗确认
const handleModalOk = async () => {
  if (!formRef.value || isUnmounted.value) return
  
  try {
    await formRef.value.validate()
    const apiMethod = isEdit.value ? updateMember : addMember
    const res = await (isEdit.value ? apiMethod(currentId.value, form) : apiMethod(form))
    
    if (!isUnmounted.value && res.data.code === 200) {
      if (isEdit.value) {
        Message.success(t('members.updateSuccess'))
      } else {
        Message.success(t('members.addSuccess'))
      }
      modalVisible.value = false
      fetchMembers()
    }
  } catch (error) {
    if (!isUnmounted.value) {
      console.error('Save member error:', error)
      Message.error(isEdit.value ? t('members.updateFailed') : t('members.addFailed'))
    }
  }
}

// 弹窗取消
const handleModalCancel = () => {
  modalVisible.value = false
}

// 分页改变
const onPageChange = (current) => {
  pagination.value.current = current
  fetchMembers()
}

// 每页条数改变
const onPageSizeChange = (pageSize) => {
  pagination.value.pageSize = pageSize
  fetchMembers()
}

// 初始化
onMounted(() => {
  isUnmounted.value = false
  fetchMembers()
})

onBeforeUnmount(() => {
  // 先设置卸载标志
  isUnmounted.value = true
  
  // 清理所有状态
  loading.value = false
  members.value = []
  modalVisible.value = false
  detailVisible.value = false
  isEdit.value = false
  memberDetail.value = {}
  
  // 重置分页
  pagination.value.total = 0
  pagination.value.current = 1
  pagination.value.pageSize = 10
  
  // 重置表单
  Object.keys(form).forEach(key => {
    form[key] = key === 'status' ? 1 : ''
  })
})
</script>

<style scoped>
.members-container {
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

:deep(.arco-table) {
  border: none;
}

:deep(.arco-table-th) {
  background-color: var(--color-fill-2) !important;
  border-top: none;
}

:deep(.arco-table-tr:hover) {
  background-color: var(--color-fill-2) !important;
}

:deep(.arco-table-td) {
  border-right: 1px solid var(--color-border) !important;
}

:deep(.arco-table-cell) {
  padding: 9px 16px;
}

:deep(.arco-typography) {
  margin-bottom: 0;
}

:deep(.arco-btn-text) {
  padding: 0 4px;
}

:deep(.arco-space-horizontal) {
  gap: 4px !important;
}

:deep(.arco-input-wrapper),
:deep(.arco-select-view),
:deep(.arco-picker) {
  background-color: var(--color-bg-2);
}

:deep(.arco-link) {
  font-size: 14px;
}

@media screen and (max-width: 768px) {
  .search-form {
    flex-direction: column;
  }
  
  .search-form-left {
    width: 100%;
  }
  
  :deep(.arco-form-item) {
    margin-right: 0;
    width: 100%;
  }
  
  :deep(.arco-input-search) {
    width: 100% !important;
  }
  
  :deep(.arco-space) {
    width: 100%;
    justify-content: flex-end;
  }
}
</style> 