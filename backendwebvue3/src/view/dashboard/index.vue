<template>
  <div class="page">
    <div class="gva-card-box">
      <div class="gva-card gva-top-card">
        <div class="gva-top-card-left">
          <div class="gva-top-card-left-title">
            <!--{{ t('view.dashboard.title') }}-->
            <h4>
              早安，管理员:{{ userStore.userInfo.nickName }}，
              请开始您一天的工作吧！
            </h4>
          </div>
          <div class="gva-top-card-left-dot">{{ weatherInfo }}</div>
          <div class="gva-top-card-left-rows">
            <el-row>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dashboard-icon">
                    <sort />
                  </el-icon>
                  馆藏图书({{summary.totalBook}})
                </div>
              </el-col>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dashboard-icon">
                    <avatar />
                  </el-icon>
                  累计借阅({{summary.totalBorrow}})
                </div>
              </el-col>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dashboard-icon">
                    <comment />
                  </el-icon>
                  待归还({{summary.totalShouldReturn}})
                </div>
              </el-col>
            </el-row>
            <!-- 新增统计行 -->
            <el-row style="margin-top: 16px;">
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dashboard-icon">
                    <calendar />
                  </el-icon>
                  本月新增({{summary.monthlyNew}})
                </div>
              </el-col>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dashboard-icon">
                    <user />
                  </el-icon>
                  读者总数({{summary.totalReader}})
                </div>
              </el-col>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dashboard-icon">
                    <warning />
                  </el-icon>
                  逾期未还({{summary.overdueCount}})
                </div>
              </el-col>
            </el-row>
          </div>
          <!--
          <div>
            <div class="gva-top-card-left-item">
              {{ t('view.dashboard.instructionalUse') }}
              <a
                style="color:#409EFF"
                target="view_window"
                href="https://www.bilibili.com/video/BV1Rg411u7xH/"
              >https://www.bilibili.com/video/BV1Rg411u7xH</a>
            </div>
            <div class="gva-top-card-left-item">
              {{ t('view.dashboard.pluginRepo') }}
              <a
                style="color:#409EFF"
                target="view_window"
                href="https://plugin.gin-vue-admin.com/#/layout/home"
              >https://plugin.gin-vue-admin.com</a>
            </div>
          </div>
          -->
        </div>
        <img src="@/assets/dashboard.png" class="gva-top-card-right" alt>
      </div>
    </div>
    <div class="gva-card-box">
      <el-card class="gva-card quick-entrance">
        <template #header>
          <div class="card-header">
            <span>{{ t('view.dashboard.quickEntry') }}</span>
          </div>
        </template>
        <el-row :gutter="20">
          <el-col
            v-for="(card, key) in toolCards"
            :key="key"
            :span="4"
            :xs="8"
            class="quick-entrance-items"
            @click="toTarget(card.name)"
          >
            <div class="quick-entrance-item">
              <div class="quick-entrance-item-icon" :style="{ backgroundColor: card.bg }">
                <el-icon>
                  <component :is="card.icon" :style="{ color: card.color }" />
                </el-icon>
              </div>
              <p>{{ card.label }}</p>
            </div>
          </el-col>
        </el-row>
      </el-card>
    <!-- <div class="quick-entrance-title"></div> -->
    </div>
    <!-- 启用图表和数据统计区域 -->
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>{{ t('view.dashboard.statistics') || '数据统计' }}</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <el-col :xs="24" :sm="18">
              <echarts-line />
            </el-col>
            <el-col :xs="24" :sm="6">
              <dashboard-table />
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
    <!-- 添加图书分类统计卡片 -->
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>图书分类统计</span>
        </div>
        <div class="category-statistics">
          <el-row :gutter="20">
            <el-col
              v-for="(category, index) in categoryStats"
              :key="index"
              :span="6"
              :xs="12"
            >
              <div class="category-card">
                <div class="category-name">{{ category.name }}</div>
                <div class="category-count">{{ category.count }} 本</div>
                <div class="category-progress">
                  <el-progress :percentage="category.percentage" :stroke-width="6" :show-text="false"/>
                </div>
              </div>
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import EchartsLine from '@/view/dashboard/dashboardCharts/echartsLine.vue'
import DashboardTable from '@/view/dashboard/dashboardTable/dashboardTable.vue'
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useWeatherInfo } from '@/view/dashboard/weather.js'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage
import { Calendar, User, Warning } from '@element-plus/icons-vue'

import { useUserStore } from '@/pinia/modules/user'

import { counts } from "@/api/book";

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const weatherInfo = useWeatherInfo()

const toolCards = ref([  
  {
    label: '图书管理',
    icon: 'cpu',
    name: 'books',
    color: '#ffd666',
    bg: 'rgba(255, 214, 102,.3)'
  },
  {
    label: '借阅记录',
    icon: 'document-checked',
    name: 'bookOrder',
    color: '#ff85c0',
    bg: 'rgba(255, 133, 192,.3)'
  },  
  {
    label: t('view.dashboard.userManage'),
    icon: 'user',
    name: 'readers',
    color: '#ff9c6e',
    bg: 'rgba(255, 156, 110,.3)'
  },  
  {
    label: t('view.dashboard.menuManage'),
    icon: 'menu',
    name: 'systemMenuList',
    color: '#b37feb',
    bg: 'rgba(179, 127, 235,.3)'
  },
  {
    label: 'API管理',
    icon: 'monitor',
    name: 'systemApiList',
    color: '#5cdbd3',
    bg: 'rgba(92, 219, 211,.3)'
  },
  {
    label: t('view.dashboard.roleManage'),
    icon: 'setting',
    name: 'systemRoleList',
    color: '#69c0ff',
    bg: 'rgba(105, 192, 255,.3)'
  }
])

const router = useRouter()

const toTarget = (name) => {
  router.push({ name })
}

const userStore = useUserStore()

// 图书汇总
const summary = ref({
  totalBook:0,
  totalBorrow:0,
  totalShouldReturn:0,
  monthlyNew: 0,      // 新增：本月新增图书数量
  totalReader: 0,     // 新增：读者总数
  overdueCount: 0     // 新增：逾期未还数量
})

// 新增：图书分类统计数据
const categoryStats = ref([
  { name: '文学艺术', count: 356, percentage: 35 },
  { name: '社会科学', count: 289, percentage: 29 },
  { name: '自然科学', count: 187, percentage: 19 },
  { name: '工程技术', count: 168, percentage: 17 }
])

onMounted(()=>{  
  loadCounts()
})
const loadCounts = async () => {
  const res = await counts({ id: 0 });  
  if (res.code == 200) {
    summary.value = {
      ...summary.value, 
      ...res.data.counts 
    }
    // 如果API返回的数据不完整，使用模拟数据补充
    if (!summary.value.monthlyNew) {
      // 模拟本月新增图书数据
      summary.value.monthlyNew = Math.floor(Math.random() * 50) + 10;
    }
    if (!summary.value.totalReader) {
      // 模拟读者总数数据
      summary.value.totalReader = Math.floor(Math.random() * 500) + 200;
    }
    if (!summary.value.overdueCount) {
      // 模拟逾期未还数据
      summary.value.overdueCount = Math.floor(Math.random() * 20);
    }
  } else {
    // 如果API调用失败，使用模拟数据
    summary.value = {
      totalBook: 1250,
      totalBorrow: 3280,
      totalShouldReturn: 45,
      monthlyNew: 32,
      totalReader: 486,
      overdueCount: 12
    }
  }
}
</script>

<style lang="scss" scoped>
@mixin flex-center {
    display: flex;
    align-items: center;
}
.page {
    background: #f0f2f5;
    padding: 0;
    .gva-card-box{
      padding: 12px 16px;
      &+.gva-card-box{
        padding-top: 0px;
      }
    }
    .gva-card {
      box-sizing: border-box;
        background-color: #fff;
        border-radius: 2px;
        height: auto;
        padding: 26px 30px;
        overflow: hidden;
        box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
    }
    .gva-top-card {
        height: auto;
        min-height: 260px;
        @include flex-center;
        justify-content: space-between;
        color: #777;
        &-left {
          height: auto;
          min-height: 260px;
          @include flex-center;
          justify-content: space-between;
          color: #777;
          &-left {
            height: 100%;
            display: flex;
            flex-direction: column;
              &-title {
                  font-size: 22px;
                  color: #343844;
              }
              &-dot {
                  font-size: 16px;
                  color: #6B7687;
                  margin-top: 24px;
              }
              &-rows {
                  // margin-top: 15px;
                  margin-top: 18px;
                  color: #6B7687;
                  width: 600px;
                  align-items: center;
              }
              &-item{
                +.gva-top-card-left-item{
                  margin-top: 24px;
                }
                margin-top: 14px;
              }
          }
          &-right {
              height: 300px;
              width: 300px;
              margin-top: 28px;
              object-fit: contain;
          }
      }
      // 新增：分类统计卡片样式
      .category-statistics {
        padding: 10px 0;
        .category-card {
          background: #f8f9fa;
          padding: 20px;
          border-radius: 8px;
          text-align: center;
          .category-name {
            font-size: 16px;
            font-weight: 500;
            margin-bottom: 8px;
          }
          .category-count {
            font-size: 24px;
            font-weight: bold;
            color: #409EFF;
            margin-bottom: 12px;
          }
          .category-progress {
            width: 100%;
          }
        }
      }
      ::v-deep(.el-card__header){
          padding:0;
          border-bottom: none;
        }
        .card-header{
          padding-bottom: 20px;
          border-bottom: 1px solid #e8e8e8;
        }
    .quick-entrance-title {
        height: 30px;
        font-size: 22px;
        color: #333;
        width: 100%;
        border-bottom: 1px solid #eee;
    }
    .quick-entrance-items {
        @include flex-center;
        justify-content: center;
        text-align: center;
        color: #333;
        .quick-entrance-item {
          padding: 16px 28px;
          margin-top: -16px;
          margin-bottom: -16px;
          border-radius: 4px;
          transition: all 0.2s;
          &:hover{
            box-shadow: 0px 0px 7px 0px rgba(217, 217, 217, 0.55);
          }
            cursor: pointer;
            height: auto;
            text-align: center;
            // align-items: center;
            &-icon {
                width: 50px;
                height: 50px !important;
                border-radius: 8px;
                @include flex-center;
                justify-content: center;
                margin: 0 auto;
                i {
                    font-size: 24px;
                }
            }
            p {
                margin-top: 10px;
            }
        }
    }
    .echart-box{
      padding: 14px;
    }
}
.dashboard-icon {
    font-size: 20px;
    color: rgb(85, 160, 248);
    width: 30px;
    height: 30px;
    margin-right: 10px;
    @include flex-center;
}
.flex-center {
    @include flex-center;
}

//小屏幕不显示右侧，将登录框居中
@media (max-width: 750px) {
    .gva-card {
        padding: 20px 10px !important;
        .gva-top-card {
            height: auto;
            &-left {
                &-title {
                    font-size: 20px !important;
                }
                &-rows {
                    margin-top: 15px;
                    align-items: center;
                }
            }
            &-right {
                display: none;
            }
        }
        .gva-middle-card {
            &-item {
                line-height: 20px;
            }
        }
        .dashboard-icon {
            font-size: 18px;
        }
    }
}
</style>