<template>
  <a-layout class="layout">
    <a-layout-header class="header">
      <div class="header-left">
        <div class="logo-container">
          <div class="logo">
            <img src="@/assets/logo.png" alt="logo" />
          </div>
        </div>
        <div class="nav-container">
          <a-breadcrumb class="breadcrumb">
            <template v-if="currentMenu">
              <a-breadcrumb-item>{{ $t(currentMenu.title) }}</a-breadcrumb-item>
              <a-breadcrumb-item v-if="currentSubmenu">{{ $t(currentSubmenu.title) }}</a-breadcrumb-item>
            </template>
          </a-breadcrumb>
        </div>
      </div>
      <div class="header-right">
        <a-dropdown trigger="hover">
          <a-button type="text" class="language-selector">
            <template #icon>
              <icon-language />
            </template>
            {{ languageLabels[currentLanguage] }}
          </a-button>
          <template #content>
            <div class="language-dropdown">
              <a-doption
                v-for="(label, value) in languageLabels"
                :key="value"
                :value="value"
                :class="{ active: value === currentLanguage }"
                @click="handleLanguageChange(value)"
              >
                {{ label }}
              </a-doption>
            </div>
          </template>
        </a-dropdown>
        <a-switch
          v-model="isDarkTheme"
          :default-checked="false"
          @change="handleThemeChange"
          checked-color="#17171a" unchecked-color="#373739"
        >
          <template #checked>
            <icon-moon-fill />
          </template>
          <template #unchecked>
            <icon-sun-fill style="color: orangered;"/>
          </template>
        </a-switch>
        <a-dropdown>
          <a-avatar :size="32" style="cursor: pointer">
            <img alt="avatar" :src="userAvatar" />
          </a-avatar>
          <template #content>
            <a-doption @click="handleProfile">
              {{ $t('common.profile') }}
            </a-doption>
            <a-doption @click="handleLogout">
              {{ $t('common.logout') }}
            </a-doption>
          </template>
        </a-dropdown>
      </div>
    </a-layout-header>
    <a-layout>
      <a-layout-sider
        :collapsed="collapsed"
        collapsible
        :width="MENU_WIDTH"
        breakpoint="xl"
        @collapse="handleCollapse"
        :show-trigger="false"
      >
        <a-menu
          :default-selected-keys="['1']"
          :style="{ width: '100%' }"
          :collapsed="collapsed"
          accordion
          :default-open-keys="openKeys"
          v-model:open-keys="openKeys"
          @menu-item-click="handleMenuItemClick"
        >
          <template v-for="menu in menus" :key="menu.key">
            <template v-if="menu.children && menu.children.length">
              <a-sub-menu :key="menu.key">
                <template #icon>
                  <component :is="menu.icon" v-if="menu.icon" />
                </template>
                <template #title>{{ menu.title }}</template>
                <a-menu-item
                  v-for="child in menu.children"
                  :key="child.key"
                >
                  <template #icon>
                    <component :is="child.icon" v-if="child.icon" />
                  </template>
                  {{ child.title }}
                </a-menu-item>
              </a-sub-menu>
            </template>
            <template v-else>
              <a-menu-item :key="menu.key">
                <template #icon>
                  <component :is="menu.icon" v-if="menu.icon" />
                </template>
                {{ menu.title }}
              </a-menu-item>
            </template>
          </template>
        </a-menu>
      </a-layout-sider>
      <a-layout-content>
        <div class="content">
          <router-view v-slot="{ Component }">
            <transition name="fade" mode="out-in">
              <component :is="Component" />
            </transition>
          </router-view>
        </div>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useStore } from '@/store'
import { useI18n } from 'vue-i18n'
import {
  IconDashboard,
  IconUser,
  IconUserGroup,
  IconApps,
  IconSettings,
  IconFile,
  IconMoonFill,
  IconSunFill,
  IconMenuFold,
  IconMenuUnfold,
  IconLanguage,
  IconOrderedList
} from '@arco-design/web-vue/es/icon'

const { locale } = useI18n()
const store = useStore()
const router = useRouter()
const route = useRoute()
const MENU_WIDTH = 220
const collapsed = ref(false)
const isDarkTheme = ref(store.theme === 'dark')
const currentMenu = ref(null)
const currentSubmenu = ref(null)
const currentLanguage = ref(localStorage.getItem('locale') || 'zh-CN')
const openKeys = ref([])
const selectedKey = ref('1')

const userAvatar = computed(() => {
  return store.user?.avatar || 'https://avatars.githubusercontent.com/u/1?v=4'
})

// 转换菜单数据为组件需要的格式
const menus = computed(() => {
  if (!store.userMenus || !Array.isArray(store.userMenus)) {
    return []
  }
  return store.userMenus.map(menu => ({
    key: menu.id.toString(),
    title: menu.name,
    icon: resolveIcon(menu.icon),
    path: menu.path,
    children: menu.children?.map(child => ({
      key: child.id.toString(),
      title: child.name,
      path: child.path,
      icon: resolveIcon(child.icon)
    }))
  }))
})

// 解析图标组件
const resolveIcon = (iconName) => {
  const iconMap = {
    'icon-dashboard': IconDashboard,
    'icon-user-group': IconUserGroup,
    'icon-order': IconOrderedList,
    'icon-product': IconApps,
    'icon-setting': IconSettings
  }
  return iconMap[iconName] || IconApps
}

// 添加一个特殊路由映射，用于面包屑
const specialRoutes = {
  '/profile': { title: 'common.profile' }
}

const languageLabels = {
  'zh-CN': '简体中文',
  'en-US': 'English',
  'zh-TW': '繁體中文',
  'ja-JP': '日本語',
  'ko-KR': '한국어',
  'de-DE': 'Deutsch',
  'fr-FR': 'Français',
  'id-ID': 'Bahasa Indonesia',
  'th-TH': 'ไทย',
  'es-ES': 'Español'
}

const handleCollapse = (val) => {
  collapsed.value = val
}

const handleMenuItemClick = (key) => {
  const menu = findMenuByKey(key)
  if (menu) {
    router.push(menu.path)
  }
}

const handleThemeChange = (checked) => {
  const theme = checked ? 'dark' : 'light'
  store.setTheme(theme)
  isDarkTheme.value = checked
}

const handleProfile = () => {
  router.push('/profile')
}

const handleLogout = () => {
  // 保存当前路径，但不保存根路径
  if (route.path !== '/') {
    localStorage.setItem('lastPath', route.path)
  }
  store.logout()
  router.push('/login')
}

const handleLanguageChange = (value) => {
  currentLanguage.value = value
  locale.value = value
  localStorage.setItem('locale', value)
}

// 根据key查找菜单项
const findMenuByKey = (key) => {
  for (const menu of menus.value) {
    if (menu.key === key) {
      return menu
    }
    if (menu.children) {
      const child = menu.children.find(child => child.key === key)
      if (child) {
        return child
      }
    }
  }
  return null
}

// 根据路径查找并设置当前菜单
const setCurrentMenu = (path) => {
  for (const menu of menus.value) {
    if (menu.path === path) {
      selectedKey.value = menu.key
      return
    }
    if (menu.children) {
      const child = menu.children.find(child => child.path === path)
      if (child) {
        selectedKey.value = child.key
        openKeys.value = [menu.key]
        return
      }
    }
  }
}

// 更新面包屑
const updateBreadcrumb = (path) => {
  const pathParts = path.split('/').filter(Boolean)
  currentMenu.value = null
  currentSubmenu.value = null
  
  // 检查是否是特殊路由
  if (specialRoutes[path]) {
    currentMenu.value = specialRoutes[path]
    return
  }
  
  // 查找一级菜单
  const mainMenu = menus.value.find(menu => {
    if (!menu.children) {
      return menu.path === path
    }
    // 如果有子菜单，检查是否有匹配的子菜单
    if (menu.children) {
      const subMenu = menu.children.find(child => child.path === path)
      if (subMenu) {
        currentSubmenu.value = subMenu
        return true
      }
    }
    return false
  })
  
  if (mainMenu) {
    currentMenu.value = mainMenu
  }
}

// 监听路由变化
watch(() => route.path, (newPath) => {
  updateBreadcrumb(newPath)
  setCurrentMenu(newPath)
}, { immediate: true })

onMounted(() => {
  // 初始化语言
  const savedLocale = localStorage.getItem('locale')
  if (savedLocale) {
    locale.value = savedLocale
    currentLanguage.value = savedLocale
  }
  
  // 初始化主题
  isDarkTheme.value = store.theme === 'dark'
  store.setTheme(store.theme)
})
</script>

<style scoped>
.layout {
  height: 100vh;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0;
  background: var(--color-bg-2);
  border-bottom: 1px solid var(--color-border);
}

.header-left {
  display: flex;
  align-items: center;
  flex: 1;
}

.logo-container {
  width: v-bind(MENU_WIDTH + 'px');
  height: 64px;
  display: flex;
  align-items: center;
  padding: 0 16px;
  background: var(--color-bg-2);
  position: relative;
  flex-shrink: 0;
  box-sizing: border-box;
  z-index: 1;
  box-shadow: 1px 0 2px 0 rgba(0, 0, 0, 0.08);
}

/* 添加阴影效果 */
/* .logo-container::after {
  content: '';
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  width: 1px;
  background: linear-gradient(180deg, 
    rgba(0, 0, 0, 0.05) 0%,
    rgba(0, 0, 0, 0.02) 100%);
} */

.logo {
  height: 32px;
  display: flex;
  align-items: center;
  img {
    height: 100%;
    width: auto;
  }
}

.nav-container {
  display: flex;
  align-items: center;
  height: 64px;
  flex: 1;
  padding: 0 16px;
  margin-left: 0; /* 移除之前的负边距 */
}

.breadcrumb {
  line-height: 64px;
}

.header-right {
  display: flex;
  align-items: center;
  padding-right: 24px; /* 与内容区右边距对齐 */
  gap: 16px;
}

.content {
  padding: 0;
  background: var(--color-bg-1);
  min-height: 100%;
}

:deep(.arco-layout-content) {
  padding: 0;
}

:deep(.page-header) {
  margin: 0;
  padding: 16px 24px;
  background: var(--color-bg-2);
  border-bottom: 1px solid var(--color-border);
  color: var(--color-text-1);
}

:deep(.page-content) {
  padding: 16px 24px;
}

:deep(.arco-typography) {
  color: var(--color-text-1);
}

.language-selector {
  display: flex;
  align-items: center;
  gap: 4px;
  padding: 5px 8px;
  border-radius: 4px;
  transition: background-color 0.2s;
}

.language-selector:hover {
  background: var(--color-fill-2);
}

:deep(.language-dropdown) {
  max-height: 400px;
  overflow-y: auto;
}

:deep(.arco-dropdown-option) {
  padding: 8px 12px;
}

:deep(.arco-dropdown-option.active) {
  color: rgb(var(--primary-6));
  background-color: var(--color-fill-2);
}

:deep(.arco-dropdown-option:hover) {
  background-color: var(--color-fill-2);
}
</style> 