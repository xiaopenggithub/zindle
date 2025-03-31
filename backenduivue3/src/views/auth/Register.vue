<template>
  <div class="register-container light-theme">
    <div class="register-box">
      <div class="register-header">
        <img src="@/assets/logo.png" alt="logo" class="logo" />
        <h2>{{ $t('register.title') }}</h2>
      </div>
      <a-form
        :model="form"
        @submit="handleSubmit"
        class="register-form"
      >
        <a-form-item field="username" :rules="[{ required: true, message: $t('register.usernameRule') }]">
          <template #label>{{ $t('register.username') }}</template>
          <a-input
            v-model="form.username"
            :placeholder="$t('register.usernamePlaceholder')"
            size="large"
            allow-clear
          >
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item 
          field="email" 
          :rules="[
            { required: true, message: $t('register.emailRule') },
            { type: 'email', message: $t('register.emailInvalid') }
          ]"
        >
          <template #label>{{ $t('register.email') }}</template>
          <a-input
            v-model="form.email"
            :placeholder="$t('register.emailPlaceholder')"
            size="large"
            allow-clear
          >
            <template #prefix>
              <icon-email />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item field="password" :rules="[{ required: true, message: $t('register.passwordRule') }]">
          <template #label>{{ $t('register.password') }}</template>
          <a-input-password
            v-model="form.password"
            :placeholder="$t('register.passwordPlaceholder')"
            size="large"
            allow-clear
          >
            <template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item 
          field="confirmPassword" 
          :rules="[
            { required: true, message: $t('register.confirmPasswordRule') },
            { validator: validatePassword, message: $t('register.passwordNotMatch') }
          ]"
        >
          <template #label>{{ $t('register.confirmPassword') }}</template>
          <a-input-password
            v-model="form.confirmPassword"
            :placeholder="$t('register.confirmPasswordPlaceholder')"
            size="large"
            allow-clear
          >
            <template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" html-type="submit" size="large" long>
            {{ $t('register.submit') }}
          </a-button>
        </a-form-item>
        <div class="login-link">
          {{ $t('register.hasAccount') }} <router-link to="/login">{{ $t('register.loginNow') }}</router-link>
        </div>
      </a-form>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import { Message } from '@arco-design/web-vue'
import { useRouter } from 'vue-router'
import { IconUser, IconLock, IconEmail } from '@arco-design/web-vue/es/icon'
import { useI18n } from 'vue-i18n'

const router = useRouter()
const { t } = useI18n()

const form = reactive({
  username: '',
  email: '',
  password: '',
  confirmPassword: ''
})

const validatePassword = (value) => value === form.password

const handleSubmit = async () => {
  try {
    // TODO: Implement actual registration logic
    Message.success(t('register.registerSuccess'))
    router.push('/login')
  } catch (error) {
    Message.error(error.message)
  }
}
</script>

<style scoped>
.register-container {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f2f3f5;
}

.register-container.light-theme {
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

.register-box {
  width: 500px;
  padding: 40px;
  background: #ffffff;
  border-radius: 4px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
}

.register-header {
  text-align: center;
  margin-bottom: 40px;
}

.register-header h2 {
  color: var(--color-text-1);
}

.logo {
  width: 64px;
  margin-bottom: 16px;
}

.register-form {
  width: 100%;
}

/* 处理浏览器自动填充的背景色 */
.register-form :deep(input:-webkit-autofill),
.register-form :deep(input:-webkit-autofill:hover),
.register-form :deep(input:-webkit-autofill:focus),
.register-form :deep(input:-webkit-autofill:active),
.register-form :deep(textarea:-webkit-autofill),
.register-form :deep(textarea:-webkit-autofill:hover),
.register-form :deep(textarea:-webkit-autofill:focus),
.register-form :deep(select:-webkit-autofill),
.register-form :deep(select:-webkit-autofill:hover),
.register-form :deep(select:-webkit-autofill:focus) {
  -webkit-text-fill-color: #1d2129 !important;
  -webkit-box-shadow: 0 0 0 1000px transparent inset !important;
  transition: background-color 50000s ease-in-out 0s;
  background-color: transparent !important;
}

/* 表单项样式 */
.register-form :deep(.arco-form-item) {
  display: flex;
  align-items: center;
  margin-bottom: 24px;
}

.register-form :deep(.arco-form-item-label-col) {
  flex: 0 0 85px;
  margin-right: 8px;
}

.register-form :deep(.arco-form-item-wrapper) {
  flex: 1;
}

/* 输入框样式 */
.register-form :deep(.arco-input-wrapper),
.register-form :deep(.arco-input) {
  background-color: transparent !important;
  border-color: #e5e6eb !important;
}

.register-form :deep(.arco-input) {
  color: #1d2129 !important;
}

.register-form :deep(.arco-input-wrapper:hover),
.register-form :deep(.arco-input-wrapper:focus-within) {
  border-color: #165dff !important;
}

/* 图标颜色 */
.register-form :deep(.arco-input-prefix),
.register-form :deep(.arco-input-password-suffix) {
  color: #4e5969 !important;
}

/* 按钮样式 */
.register-form :deep(.arco-btn-primary) {
  background-color: #165dff !important;
  border-color: #165dff !important;
  color: #ffffff !important;
}

.register-form :deep(.arco-btn-primary:hover) {
  opacity: 0.9;
}

/* 链接样式 */
.login-link {
  text-align: center;
  margin-top: 16px;
  color: var(--color-text-3);
}

.login-link :deep(a) {
  color: var(--color-primary);
  text-decoration: none;
}

.login-link :deep(a:hover) {
  text-decoration: underline;
}
</style> 