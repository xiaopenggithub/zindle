<template>
  <div class="orders-container">
    <a-card class="general-card">
      <div class="search-form-wrapper">
        <a-form layout="inline" class="search-form" :model="searchForm">
          <div class="search-form-left">
            <a-form-item field="status">
              <a-radio-group v-model="searchForm.status" type="button" @change="handleStatusChange">
                <a-radio :value="'all'">{{ t('orders.status.all') }}</a-radio>
                <a-radio :value="'paid'">{{ t('orders.status.paid') }}</a-radio>
                <a-radio :value="'unpaid'">{{ t('orders.status.unpaid') }}</a-radio>
                <a-radio :value="'cancelled'">{{ t('orders.status.cancelled') }}</a-radio>
              </a-radio-group>
            </a-form-item>
            <a-form-item field="dateRange">
              <a-range-picker
                v-model="searchForm.dateRange"
                :style="{ width: '360px' }"
                :placeholder="[t('orders.startTime'), t('orders.endTime')]"
                show-time
                format="YYYY-MM-DD HH:mm:ss"
                @change="handleDateChange"
                allow-clear
              />
            </a-form-item>
            <a-form-item field="keyword">
              <a-input-search
                v-model="searchForm.keyword"
                :placeholder="t('orders.searchPlaceholder')"
                :style="{ width: '320px' }"
                @search="handleSearch"
                allow-clear
              />
            </a-form-item>
            <a-form-item>
              <a-button type="outline" status="normal" @click="resetFilters">
                {{ t('common.reset') }}
              </a-button>
            </a-form-item>
          </div>
        </a-form>
      </div>

      <a-table
        :loading="loading"
        :data="orders"
        :pagination="pagination"
        @page-change="onPageChange"
        @page-size-change="onPageSizeChange"
        :bordered="{ cell: true }"
        :stripe="true"
        size="medium"
      >
        <template #columns>
          <a-table-column :title="t('orders.orderNo')" data-index="orderNo" :width="280">
            <template #cell="{ record }">
              <a-typography-text>{{ record.orderNo }}</a-typography-text>
            </template>
          </a-table-column>
          <a-table-column :title="t('orders.memberName')" data-index="memberName" :width="150">
            <template #cell="{ record }">
              <a-link @click="handleViewMember(record)">{{ record.memberName }}</a-link>
            </template>
          </a-table-column>
          <a-table-column :title="t('orders.amount')" data-index="amount" align="right">
            <template #cell="{ record }">
              <a-typography-text type="warning">¥{{ record.amount.toFixed(2) }}</a-typography-text>
            </template>
          </a-table-column>
          <a-table-column :title="t('orders.status.label')" data-index="status" align="center" :width="100">
            <template #cell="{ record }">
              <a-tag :color="getStatusColor(record.status)">
                {{ t(`orders.status.${record.status}`) }}
              </a-tag>
            </template>
          </a-table-column>
          <a-table-column :title="t('orders.createTime')" data-index="createTime" align="center" :width="180" />
          <a-table-column :title="t('orders.payTime')" data-index="payTime" align="center" :width="180">
            <template #cell="{ record }">
              <a-typography-text type="secondary">{{ record.payTime || '-' }}</a-typography-text>
            </template>
          </a-table-column>
          <a-table-column :title="t('common.operations')" align="center" fixed="right" :width="140">
            <template #cell="{ record }">
              <a-space>
                <a-button type="text" size="small" @click="handleView(record)">
                  <template #icon><icon-eye /></template>
                  {{ t('common.view') }}
                </a-button>
                <a-button 
                  v-if="record.status === 'unpaid'"
                  type="text" 
                  status="success" 
                  size="small" 
                  @click="handlePay(record)"
                >
                  <template #icon><icon-check /></template>
                  {{ t('orders.pay') }}
                </a-button>
                <a-popconfirm
                  v-if="record.status === 'unpaid'"
                  :content="t('orders.confirmCancel')"
                  @ok="handleCancel(record)"
                  position="left"
                >
                  <a-button type="text" status="danger" size="small">
                    <template #icon><icon-close /></template>
                    {{ t('orders.cancel') }}
                  </a-button>
                </a-popconfirm>
              </a-space>
            </template>
          </a-table-column>
        </template>
      </a-table>

      <!-- 订单详情弹窗 -->
      <a-modal
        v-model:visible="detailVisible"
        :title="t('orders.orderDetail')"
        @cancel="detailVisible = false"
        :unmount-on-close="true"
        :width="800"
        :mask-closable="false"
      >
        <a-descriptions :data="orderDetailData" layout="horizontal" bordered :column="2" size="medium" />
        <template #footer>
          <a-button @click="detailVisible = false">{{ t('common.close') }}</a-button>
        </template>
      </a-modal>

      <!-- 会员信息弹窗 -->
      <a-modal
        v-model:visible="memberVisible"
        :title="t('orders.memberInfo')"
        @cancel="memberVisible = false"
        :unmount-on-close="true"
        :width="600"
        :mask-closable="false"
      >
        <a-descriptions :data="memberDetailData" layout="horizontal" bordered :column="1" size="medium" />
        <template #footer>
          <a-button @click="memberVisible = false">{{ t('common.close') }}</a-button>
        </template>
      </a-modal>
    </a-card>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { Message } from '@arco-design/web-vue'
import { IconEye, IconCheck, IconClose } from '@arco-design/web-vue/es/icon'
import { getOrders, payOrder, cancelOrder, getMemberInfo } from '@/api/mock/order'

const { t } = useI18n()
const loading = ref(false)
const orders = ref([])
const detailVisible = ref(false)
const memberVisible = ref(false)
const orderDetail = ref(null)
const memberDetail = ref(null)

// 搜索表单数据
const searchForm = reactive({
  status: 'all',
  dateRange: [],
  keyword: ''
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

// 获取状态标签颜色
const getStatusColor = (status) => {
  const colorMap = {
    paid: 'green',
    unpaid: 'orange',
    cancelled: 'red'
  }
  return colorMap[status] || 'blue'
}

// 订单详情数据转换
const orderDetailData = computed(() => {
  if (!orderDetail.value) return []
  
  return [
    {
      label: t('orders.orderNo'),
      value: orderDetail.value.orderNo
    },
    {
      label: t('orders.memberName'),
      value: orderDetail.value.memberName
    },
    {
      label: t('orders.amount'),
      value: `¥${orderDetail.value.amount.toFixed(2)}`
    },
    {
      label: t('orders.status.label'),
      value: t(`orders.status.${orderDetail.value.status}`)
    },
    {
      label: t('orders.createTime'),
      value: orderDetail.value.createTime
    },
    {
      label: t('orders.payTime'),
      value: orderDetail.value.payTime || '-'
    },
    {
      label: t('orders.cancelTime'),
      value: orderDetail.value.cancelTime || '-'
    },
    {
      label: t('orders.cancelReason'),
      value: orderDetail.value.cancelReason || '-'
    },
    {
      label: t('orders.remark'),
      value: orderDetail.value.remark || '-',
      span: 2
    }
  ]
})

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
      label: t('members.registerTime'),
      value: memberDetail.value.registerTime
    },
    {
      label: t('members.status.label'),
      value: t(`members.status.${memberDetail.value.status}`)
    }
  ]
})

// 获取订单列表
const fetchOrders = async () => {
  loading.value = true
  try {
    const res = await getOrders({
      current: pagination.current,
      pageSize: pagination.pageSize,
      status: searchForm.status,
      keyword: searchForm.keyword,
      startTime: searchForm.dateRange?.[0]?.toISOString(),
      endTime: searchForm.dateRange?.[1]?.toISOString()
    })
    
    if (res.data.code === 200) {
      orders.value = res.data.data.list
      pagination.total = res.data.data.total
    } else {
      Message.error(t('orders.fetchFailed'))
    }
  } catch (error) {
    console.error('Fetch orders error:', error)
    Message.error(t('orders.fetchFailed'))
  } finally {
    loading.value = false
  }
}

// 查看订单详情
const handleView = (record) => {
  orderDetail.value = { ...record }
  detailVisible.value = true
}

// 查看会员信息
const handleViewMember = async (record) => {
  try {
    const res = await getMemberInfo(record.memberId)
    if (res.data.code === 200) {
      memberDetail.value = res.data.data
      memberVisible.value = true
    } else {
      Message.error(t('members.fetchFailed'))
    }
  } catch (error) {
    console.error('Fetch member error:', error)
    Message.error(t('members.fetchFailed'))
  }
}

// 支付订单
const handlePay = async (record) => {
  try {
    const res = await payOrder(record.id)
    if (res.data.code === 200) {
      Message.success(t('orders.paySuccess'))
      fetchOrders()
    } else {
      Message.error(res.data.message || t('orders.payFailed'))
    }
  } catch (error) {
    console.error('Pay order error:', error)
    Message.error(t('orders.payFailed'))
  }
}

// 取消订单
const handleCancel = async (record) => {
  try {
    const res = await cancelOrder(record.id)
    if (res.data.code === 200) {
      Message.success(t('orders.cancelSuccess'))
      fetchOrders()
    } else {
      Message.error(res.data.message || t('orders.cancelFailed'))
    }
  } catch (error) {
    console.error('Cancel order error:', error)
    Message.error(t('orders.cancelFailed'))
  }
}

// 状态切换
const handleStatusChange = () => {
  pagination.current = 1
  fetchOrders()
}

// 分页改变
const onPageChange = (current) => {
  pagination.current = current
  fetchOrders()
}

// 每页条数改变
const onPageSizeChange = (pageSize) => {
  pagination.pageSize = pageSize
  fetchOrders()
}

// 日期范围改变
const handleDateChange = () => {
  pagination.current = 1
  fetchOrders()
}

// 搜索
const handleSearch = () => {
  pagination.current = 1
  fetchOrders()
}

// 重置筛选条件
const resetFilters = () => {
  searchForm.status = 'all'
  searchForm.dateRange = []
  searchForm.keyword = ''
  pagination.current = 1
  fetchOrders()
}

// 初始化
onMounted(() => {
  fetchOrders()
})
</script>

<style scoped>
.orders-container {
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

:deep(.arco-table-th) {
  background-color: var(--color-fill-2) !important;
}

:deep(.arco-table-tr:hover) {
  background-color: var(--color-fill-2) !important;
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
  
  :deep(.arco-radio-group) {
    width: 100%;
    display: flex;
    justify-content: space-between;
  }
  
  :deep(.arco-picker) {
    width: 100% !important;
  }
  
  :deep(.arco-input-search) {
    width: 100% !important;
  }
}
</style> 