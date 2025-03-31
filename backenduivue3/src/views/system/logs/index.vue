<template>
  <div class="logs-container">
    <a-card class="general-card">
      <!-- 搜索区域 -->
      <div class="search-form-wrapper">
        <a-form layout="inline" class="search-form" :model="searchForm">
          <div class="search-form-left">
            <a-form-item field="keyword">
              <a-input-search
                v-model="searchForm.keyword"
                :placeholder="t('logs.placeholder.search')"
                :style="{ width: '220px' }"
                search-button
                @search="handleSearch"
              />
            </a-form-item>
            <a-form-item field="level">
              <a-select
                v-model="searchForm.level"
                :style="{ width: '130px' }"
                :placeholder="t('logs.placeholder.selectLevel')"
                allow-clear
                @change="handleLevelChange"
              >
                <a-option value="">{{ t('logs.allLevels') }}</a-option>
                <a-option value="INFO">{{ t('logs.levels.info') }}</a-option>
                <a-option value="WARNING">{{ t('logs.levels.warning') }}</a-option>
                <a-option value="ERROR">{{ t('logs.levels.error') }}</a-option>
                <a-option value="DEBUG">{{ t('logs.levels.debug') }}</a-option>
              </a-select>
            </a-form-item>
            <a-form-item field="dateRange">
              <a-range-picker
                v-model="searchForm.dateRange"
                :style="{ width: '240px' }"
                :placeholder="[t('logs.time'), t('logs.time')]"
                show-time
                format="YYYY-MM-DD HH:mm:ss"
                @change="handleDateChange"
              />
            </a-form-item>
            <a-form-item>
              <a-button @click="resetFilters">
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
              <a-button type="primary" @click="handleExport">
                <template #icon>
                  <icon-download />
                </template>
                {{ t('logs.export') }}
              </a-button>
            </a-space>
          </div>
        </a-form>
      </div>

      <a-table
        :data="logs"
        :loading="loading"
        :pagination="pagination"
        :bordered="true"
        :stripe="true"
        @page-change="onPageChange"
      >
        <template #columns>
          <a-table-column :title="t('logs.time')" data-index="time" :width="280" />
          <a-table-column :title="t('logs.level')" data-index="level" :width="100" align="center">
            <template #cell="{ record }">
              <a-tag :color="getLogLevelColor(record.level)">
                {{ t(`logs.levels.${record.level.toLowerCase()}`) }}
              </a-tag>
            </template>
          </a-table-column>
          <a-table-column :title="t('logs.user')" data-index="user" :width="120" />
          <a-table-column :title="t('logs.ip')" data-index="ip" :width="140" />
          <a-table-column :title="t('logs.action')" data-index="action" :width="120">
            <template #cell="{ record }">
              {{ t(`logs.actions.${record.action.toLowerCase()}`) }}
            </template>
          </a-table-column>
          <a-table-column :title="t('logs.details')" data-index="details">
            <template #cell="{ record }">
              <a-typography-paragraph :ellipsis="{ rows: 2, showTooltip: true }">
                {{ record.details }}
              </a-typography-paragraph>
            </template>
          </a-table-column>
        </template>
        <template #empty>
          <a-empty :description="t('logs.table.noData')" />
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script setup>
import { ref, watch, onMounted, reactive } from 'vue'
import { IconDownload, IconRefresh } from '@arco-design/web-vue/es/icon'
import { Message, Modal } from '@arco-design/web-vue'
import { useI18n } from 'vue-i18n'
import { getLogs, deleteLog, exportLogs } from '@/api/mock/logs'

const { t } = useI18n()

// 搜索表单数据
const searchForm = reactive({
  keyword: '',
  level: '',
  dateRange: []
})

const loading = ref(false)
const logs = ref([])

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

const getLogLevelColor = (level) => {
  const colors = {
    INFO: 'blue',
    WARNING: 'orange',
    ERROR: 'red',
    DEBUG: 'green'
  }
  return colors[level] || 'blue'
}

const fetchLogs = async () => {
  loading.value = true
  try {
    const params = {
      current: pagination.value.current,
      pageSize: pagination.value.pageSize,
      searchQuery: searchForm.keyword,
      logLevel: searchForm.level,
      dateRange: searchForm.dateRange ? JSON.stringify(searchForm.dateRange) : ''
    }
    
    const res = await getLogs(params)
    if (res.data.code === 200) {
      logs.value = res.data.data.list
      pagination.value.total = res.data.data.total
    } else {
      Message.error(t('logs.fetchFailed'))
    }
  } catch (error) {
    console.error('Error fetching logs:', error)
    Message.error(t('logs.fetchFailed'))
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.value.current = 1
  fetchLogs()
}

const handleLevelChange = () => {
  pagination.value.current = 1
  fetchLogs()
}

const handleDateChange = () => {
  pagination.value.current = 1
  fetchLogs()
}

const handleExport = async () => {
  try {
    const res = await exportLogs()
    if (res.data.code === 200) {
      // Here you would typically handle the export data
      // For example, download as CSV or Excel
      Message.success(t('logs.exportSuccess'))
    } else {
      Message.error(t('logs.exportFailed'))
    }
  } catch (error) {
    console.error('Error exporting logs:', error)
    Message.error(t('logs.exportFailed'))
  }
}

const handleDelete = (record) => {
  Modal.warning({
    title: t('common.confirmDelete'),
    content: t('logs.confirmDelete'),
    onOk: async () => {
      try {
        const res = await deleteLog(record.id)
        if (res.data.code === 200) {
          Message.success(t('logs.deleteSuccess'))
          fetchLogs()
        } else {
          Message.error(t('logs.deleteFailed'))
        }
      } catch (error) {
        Message.error(t('logs.deleteFailed'))
      }
    }
  })
}

const onPageChange = (page) => {
  pagination.value.current = page
  fetchLogs()
}

const resetFilters = () => {
  searchForm.keyword = ''
  searchForm.level = ''
  searchForm.dateRange = []
  handleSearch()
}

const handleRefresh = () => {
  fetchLogs()
}

// Watch for search query changes
watch(() => searchForm.keyword, () => {
  if (!searchForm.keyword) {
    handleSearch()
  }
})

// Initial data fetch
onMounted(() => {
  fetchLogs()
})
</script>

<style scoped>
.logs-container {
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
:deep(.arco-select-view),
:deep(.arco-picker) {
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