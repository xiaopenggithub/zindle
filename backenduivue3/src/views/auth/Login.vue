<template>
  <div class="login-container light-theme">
    <div class="login-box">
      <div class="login-header">
        <img src="@/assets/logo.png" alt="logo" class="logo" />
        <h2>{{ $t('login.title') }}</h2>
      </div>
      <a-form
        ref="formRef"
        :model="form"
        :rules="rules"
        @submit="handleSubmit"
        class="login-form"
      >
        <a-form-item field="username" validate-trigger="blur">
          <template #label>{{ $t('login.username') }}</template>
          <a-input
            v-model="form.username"
            :placeholder="$t('login.usernamePlaceholder')"
            size="large"
            allow-clear
          >
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item field="password" validate-trigger="blur">
          <template #label>{{ $t('login.password') }}</template>
          <a-input-password
            v-model="form.password"
            :placeholder="$t('login.passwordPlaceholder')"
            size="large"
            allow-clear
          >
            <template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item field="captcha" validate-trigger="blur">
          <template #label>{{ $t('login.captcha') }}</template>
          <div class="captcha-container">
            <a-input
              v-model="form.captcha"
              :placeholder="$t('login.captchaPlaceholder')"
              size="large"
            >
              <template #prefix>
                <icon-safe />
              </template>
            </a-input>
            <div class="captcha-image" @click="refreshCaptcha">
              <img :src="captchaImage" alt="验证码" />
            </div>
          </div>
        </a-form-item>
        <div class="login-options">
          <div class="login-form-remember">
            <a-checkbox>{{ $t('login.rememberMe') }}</a-checkbox>
          </div>
          <router-link to="/forgotPassword">{{ $t('login.forgotPassword') }}</router-link>
        </div>
        <a-form-item>
          <a-button type="primary" html-type="submit" size="large" long :loading="loading">
            {{ $t('common.login') }}
          </a-button>
        </a-form-item>
        <div class="register-link">
          {{ $t('login.noAccount') }} <router-link to="/register">{{ $t('login.registerNow') }}</router-link>
        </div>
      </a-form>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Message } from '@arco-design/web-vue'
import { IconUser, IconLock, IconSafe } from '@arco-design/web-vue/es/icon'
import { useStore } from '@/store'
import { useI18n } from 'vue-i18n'
import { getCaptcha } from '@/api/user'
import { getToken } from '@/utils/auth'

const router = useRouter()
const store = useStore()
const { t } = useI18n()
const formRef = ref(null)
const loading = ref(false)

// 检查登录状态并自动跳转
const checkLoginStatus = async () => {
  const token = getToken()
  if (token) {
    try {
      // 获取上次访问的页面
      const lastPath = localStorage.getItem('lastPath')
      // 如果有上次访问的页面且不是登录相关页面，则跳转到该页面
      if (lastPath && !['/login', '/register', '/'].includes(lastPath)) {
        router.replace(lastPath)
      } else {
        router.replace('/dashboard')
      }
    } catch (error) {
      console.error('Failed to get user info:', error)
      // 登录页面刷新验证码
      refreshCaptcha()
    }
  } else {
    // 没有登录状态，刷新验证码
    refreshCaptcha()
  }  
}

const form = reactive({
  username: '',
  password: '',
  captcha: '',
  captcha_id: ''
})

const rules = {
  username: [
    { required: true, message: t('login.usernameRule') }
  ],
  password: [
    { required: true, message: t('login.passwordRule') }
  ],
  captcha: [
    { required: true, message: t('login.captchaRule') }
  ]
}

const captchaImage = ref('')

const refreshCaptcha = async () => {
  try {
    const res = await getCaptcha()
    if (res.data.code === 200) {
      captchaImage.value = res.data.data.captcha_code
      form.captcha_id = res.data.data.captcha_id
    } else {
      Message.error(res.data.message || t('login.captchaError'))
    }
  } catch (error) {
    Message.error(t('login.captchaError'))
    console.error('获取验证码失败:', error)
  }
}

// 在组件挂载时检查登录状态
onMounted(() => {
  checkLoginStatus()  
})

const handleSubmit = async ({ values, errors }) => {
  if (errors) {
    return
  }
  
  try {
    loading.value = true
    await store.login(values)
    Message.success(t('login.loginSuccess'))
    // 获取重定向地址
    const redirect = router.currentRoute.value.query.redirect || '/'
    // 等待一小段时间确保 token 和角色信息已保存    
    setTimeout(() => {
      router.push(redirect)
    }, 100)
  } catch (error) {
    // Message.error(error.message)
    console.log('登录失败:', error.message)
    refreshCaptcha() // 登录失败时刷新验证码
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f2f3f5;
}

.login-container.light-theme {
  --color-bg-1: #ffffff;
  --color-bg-2: #f2f3f5;
  --color-text-1: #1d2129;
  --color-text-2: #4e5969;
  --color-text-3: #86909c;
  --color-border: #e5e6eb;
  --color-primary: #165dff;
  --color-primary-light: #e8f3ff;
  color: var(--color-text-1);
}

.login-box {
  width: 400px;
  padding: 40px;
  background: #ffffff;
  border-radius: 4px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
}

.login-header {
  text-align: center;
  margin-bottom: 40px;
}

.login-header h2 {
  color: var(--color-text-1);
}

.logo {
  width: 64px;
  margin-bottom: 16px;
}

.login-form {
  width: 100%;
}

/* 处理浏览器自动填充的背景色 */
.login-form :deep(input:-webkit-autofill),
.login-form :deep(input:-webkit-autofill:hover),
.login-form :deep(input:-webkit-autofill:focus),
.login-form :deep(input:-webkit-autofill:active),
.login-form :deep(textarea:-webkit-autofill),
.login-form :deep(textarea:-webkit-autofill:hover),
.login-form :deep(textarea:-webkit-autofill:focus),
.login-form :deep(select:-webkit-autofill),
.login-form :deep(select:-webkit-autofill:hover),
.login-form :deep(select:-webkit-autofill:focus) {
  -webkit-text-fill-color: #1d2129 !important;
  -webkit-box-shadow: 0 0 0 1000px transparent inset !important;
  transition: background-color 50000s ease-in-out 0s;
  background-color: transparent !important;
}

/* 表单项样式 */
.login-form :deep(.arco-form-item-label-col) {
  text-align: left;
  color: var(--color-text-2);
}

/* 输入框样式 */
.login-form :deep(.arco-input-wrapper),
.login-form :deep(.arco-input),
.captcha-container :deep(.arco-input-wrapper) {
  background-color: #ffffff !important;
  border-color: #e5e6eb !important;
}

.login-form :deep(.arco-input) {
  color: #1d2129 !important;
}

.login-form :deep(.arco-input-wrapper:hover),
.login-form :deep(.arco-input-wrapper:focus-within) {
  border-color: #165dff !important;
}

/* 图标颜色 */
.login-form :deep(.arco-input-prefix),
.login-form :deep(.arco-input-password-suffix) {
  color: #4e5969 !important;
}

/* 复选框样式 */
.login-form :deep(.arco-checkbox-label) {
  color: #4e5969 !important;
}

.login-form :deep(.arco-checkbox-mask) {
  background-color: #ffffff !important;
}

/* 按钮样式 */
.login-form :deep(.arco-btn-primary) {
  background-color: #165dff !important;
  border-color: #165dff !important;
  color: #ffffff !important;
}

.login-form :deep(.arco-btn-primary:hover) {
  opacity: 0.9;
}

/* 布局样式 */
.login-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  padding-left: 0;
}

.login-form-remember {
  margin-left: 75px;
}

/* 链接样式 */
.login-options :deep(a),
.register-link :deep(a) {
  color: var(--color-primary);
  text-decoration: none;
}

.login-options :deep(a:hover),
.register-link :deep(a:hover) {
  text-decoration: underline;
}

.register-link {
  text-align: center;
  margin-top: 16px;
  color: var(--color-text-3);
}

/* 验证码样式 */
.captcha-container {
  display: flex;
  gap: 12px;
  align-items: center;
}

.captcha-container :deep(.arco-input-wrapper) {
  flex: 1;
}

.captcha-image {
  width: 100px;
  height: 35px;
  cursor: pointer;
  border: 1px solid #e5e6eb !important;
  border-radius: 4px;
  overflow: hidden;
  background-color: #ffffff !important;
}

.captcha-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
</style> 