<template>
  <div class="login-container light-theme">
    <div class="login-box">
      <div class="login-header">
        <img src="@/assets/logo.png" alt="logo" class="logo" />
        <h2>{{ $t('login.title') }}</h2>
      </div>
      <a-form
        :model="form"
        @submit="handleSubmit"
        class="login-form"
      >
        <a-form-item field="username" :rules="[{ required: true, message: $t('login.usernameRule') }]">
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
        <a-form-item field="password" :rules="[{ required: true, message: $t('login.passwordRule') }]">
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
        <a-form-item field="captcha" :rules="[{ required: true, message: $t('login.captchaRule') }]">
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
          <a-button type="primary" html-type="submit" size="large" long>
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

const router = useRouter()
const store = useStore()
const { t } = useI18n()
const form = reactive({
  username: '',
  password: '',
  captcha: ''
})

import captchaImg from '@/assets/logo.png';
const captchaImage = ref(captchaImg);

const refreshCaptcha = () => {
  // 这里应该调用后端接口获取验证码图片
  // 示例：captchaImage.value = await getCaptchaImage()
}

// 在组件挂载时获取验证码
onMounted(() => {
  refreshCaptcha()
})

const handleSubmit = async () => {    
  try {
    await store.login(form.username, form.password, form.captcha)
    Message.success(t('login.loginSuccess'))
    // 获取重定向地址
    const redirect = router.currentRoute.value.query.redirect || '/'
    // 等待一小段时间确保 token 和角色信息已保存
    setTimeout(() => {
      router.push(redirect)
    }, 100)
  } catch (error) {
    Message.error(error.message)
    refreshCaptcha() // 登录失败时刷新验证码
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