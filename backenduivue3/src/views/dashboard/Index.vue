<template>
  <div class="dashboard">
    <a-row :gutter="16">
      <a-col :span="6">
        <a-card class="stat-card">
          <template #title>
            <div class="card-title">
              <icon-user />
              <span>{{ t('dashboard.totalUsers') }}</span>
            </div>
          </template>
          <div class="stat-value">1,234</div>
          <div class="stat-footer">
            <span class="trend up">+12%</span>
            <span class="period">{{ t('dashboard.vsLastMonth') }}</span>
          </div>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card class="stat-card">
          <template #title>
            <div class="card-title">
              <icon-file />
              <span>{{ t('dashboard.totalLogs') }}</span>
            </div>
          </template>
          <div class="stat-value">5,678</div>
          <div class="stat-footer">
            <span class="trend up">+8%</span>
            <span class="period">{{ t('dashboard.vsLastMonth') }}</span>
          </div>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card class="stat-card">
          <template #title>
            <div class="card-title">
              <icon-settings />
              <span>{{ t('dashboard.activeRoles') }}</span>
            </div>
          </template>
          <div class="stat-value">12</div>
          <div class="stat-footer">
            <span class="trend">0%</span>
            <span class="period">{{ t('dashboard.vsLastMonth') }}</span>
          </div>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card class="stat-card">
          <template #title>
            <div class="card-title">
              <icon-menu />
              <span>{{ t('dashboard.menuItems') }}</span>
            </div>
          </template>
          <div class="stat-value">45</div>
          <div class="stat-footer">
            <span class="trend down">-2%</span>
            <span class="period">{{ t('dashboard.vsLastMonth') }}</span>
          </div>
        </a-card>
      </a-col>
    </a-row>

    <a-row :gutter="16" class="mt-16">
      <a-col :span="16">
        <a-card :title="t('dashboard.recentActivity')">
          <a-table :data="recentActivity" :columns="activityColumns" />
        </a-card>
      </a-col>
      <a-col :span="8">
        <a-card :title="t('dashboard.systemStatus')">
          <a-list>
            <a-list-item>
              <template #meta>
                <span>{{ t('dashboard.cpuUsage') }}</span>
                <a-progress :percent="65" />
              </template>
            </a-list-item>
            <a-list-item>
              <template #meta>
                <span>{{ t('dashboard.memoryUsage') }}</span>
                <a-progress :percent="45" />
              </template>
            </a-list-item>
            <a-list-item>
              <template #meta>
                <span>{{ t('dashboard.diskUsage') }}</span>
                <a-progress :percent="75" />
              </template>
            </a-list-item>
          </a-list>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import {
  IconUser,
  IconFile,
  IconSettings,
  IconMenu
} from '@arco-design/web-vue/es/icon'

const { t } = useI18n()

const activityColumns = [
  {
    title: t('dashboard.time'),
    dataIndex: 'time',
  },
  {
    title: t('dashboard.user'),
    dataIndex: 'user',
  },
  {
    title: t('dashboard.action'),
    dataIndex: 'action',
  },
  {
    title: t('dashboard.status'),
    dataIndex: 'status',
  },
]

const recentActivity = ref([
  {
    time: '2024-03-28 10:00:00',
    user: 'Admin',
    action: 'Login',
    status: t('dashboard.success'),
  },
  {
    time: '2024-03-28 09:45:00',
    user: 'User1',
    action: 'Update Profile',
    status: t('dashboard.success'),
  },
  {
    time: '2024-03-28 09:30:00',
    user: 'User2',
    action: 'Create Role',
    status: t('dashboard.success'),
  },
  {
    time: '2024-03-28 09:15:00',
    user: 'User3',
    action: 'Delete Permission',
    status: t('dashboard.failed'),
  },
])
</script>

<style scoped>
.dashboard {
  padding: 16px;
}

.mt-16 {
  margin-top: 16px;
}

.stat-card {
  height: 100%;
}

.card-title {
  display: flex;
  align-items: center;
  gap: 8px;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  margin: 16px 0;
}

.stat-footer {
  display: flex;
  align-items: center;
  gap: 8px;
  color: var(--color-text-3);
}

.trend {
  display: flex;
  align-items: center;
  gap: 4px;
}

.trend.up {
  color: var(--success-color);
}

.trend.down {
  color: var(--danger-color);
}

.period {
  font-size: 12px;
}
</style> 