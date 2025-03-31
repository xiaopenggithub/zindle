<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from '@/store'

const store = useStore()
const router = useRouter()

onMounted(() => {
  // 检查登录状态
  if (!store.isAuthenticated && router.currentRoute.value.path !== '/login') {
    router.push('/login')
  }
})
</script>

<template>
  <router-view v-slot="{ Component }">
    <transition name="fade" mode="out-in">
      <component :is="Component" />
    </transition>
  </router-view>
</template>

<style>
html,
body {
  margin: 0;
  padding: 0;
  height: 100%;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto,
    'Helvetica Neue', Arial, 'Noto Sans', sans-serif, 'Apple Color Emoji',
    'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji';
}

#app {
  height: 100vh;
}

.page-container {
  padding: 16px;
}

.page-header {
  margin-bottom: 16px;
}

.search-bar {
  margin-bottom: 16px;
}

.table-operations {
  margin-bottom: 16px;
}

.form-actions {
  margin-top: 24px;
  text-align: right;
}

.form-actions .arco-btn + .arco-btn {
  margin-left: 8px;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
