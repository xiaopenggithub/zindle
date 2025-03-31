<template>
  <div class="forgot-password-container light-theme">
    <div class="forgot-password-box">
      <div class="forgot-password-header">
        <img src="@/assets/logo.png" alt="logo" class="logo" />
        <h2>{{ $t('forgotPassword.title') }}</h2>
        <p class="description">{{ $t('forgotPassword.description') }}</p>
      </div>
      <a-form
        ref="formRef"
        :model="form"
        @submit="handleSubmit"
        class="forgot-password-form"
      >
        <a-form-item
          field="email"
          :label="$t('forgotPassword.email')"
          :rules="[
            { required: true, message: $t('forgotPassword.emailRule') },
            { type: 'email', message: $t('forgotPassword.emailInvalid') }
          ]"
        >
          <a-input
            v-model="form.email"
            :placeholder="$t('forgotPassword.emailPlaceholder')"
            size="large"
            allow-clear
          >
            <template #prefix>
              <icon-email />
            </template>
          </a-input>
        </a-form-item>
        
        <a-form-item
          field="verificationCode"
          :label="$t('forgotPassword.verificationCode')"
          :rules="[{ required: true, message: $t('forgotPassword.verificationCodeRule') }]"
        >
          <a-space>
            <a-input
              v-model="form.verificationCode"
              :placeholder="$t('forgotPassword.verificationCodePlaceholder')"
              size="large"
              allow-clear
              style="width: 280px"
            >
              <template #prefix>
                <icon-safe />
              </template>
            </a-input>
            <a-button
              type="primary"
              size="large"
              :loading="sendingCode"
              :disabled="!!countdown || !form.email"
              @click="handleSendCode"
            >
              {{ countdown ? `${countdown}s` : $t('forgotPassword.getVerificationCode') }}
            </a-button>
          </a-space>
        </a-form-item>

        <a-form-item
          field="newPassword"
          :label="$t('forgotPassword.newPassword')"
          :rules="[
            { required: true, message: $t('forgotPassword.newPasswordRule') },
            { minLength: 8, message: $t('forgotPassword.passwordMinLength') }
          ]"
        >
          <a-input-password
            v-model="form.newPassword"
            :placeholder="$t('forgotPassword.newPasswordPlaceholder')"
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
          :label="$t('forgotPassword.confirmPassword')"
          :rules="[
            { required: true, message: $t('forgotPassword.confirmPasswordRule') },
            { validator: validatePassword, message: $t('forgotPassword.passwordNotMatch') }
          ]"
        >
          <a-input-password
            v-model="form.confirmPassword"
            :placeholder="$t('forgotPassword.confirmPasswordPlaceholder')"
            size="large"
            allow-clear
          >
            <template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item>
          <a-button type="primary" html-type="submit" size="large" long :loading="submitting">
            {{ $t('forgotPassword.submit') }}
          </a-button>
        </a-form-item>
        <div class="login-link">
          <router-link to="/login">{{ $t('forgotPassword.back') }}</router-link>
        </div>
      </a-form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { Message } from '@arco-design/web-vue'
import { IconEmail, IconLock, IconSafe } from '@arco-design/web-vue/es/icon'
import { useI18n } from 'vue-i18n'

const router = useRouter()
const { t } = useI18n()
const formRef = ref(null)
const submitting = ref(false)
const sendingCode = ref(false)
const countdown = ref(0)

const form = ref({
  email: '',
  verificationCode: '',
  newPassword: '',
  confirmPassword: ''
})

const validatePassword = (value) => value === form.value.newPassword

const startCountdown = () => {
  countdown.value = 60
  const timer = setInterval(() => {
    countdown.value--
    if (countdown.value <= 0) {
      clearInterval(timer)
    }
  }, 1000)
}

const handleSendCode = async () => {
  try {
    sendingCode.value = true
    // TODO: 调用发送验证码的API
    // await api.sendVerificationCode(form.value.email)
    Message.success(t('forgotPassword.codeSent'))
    startCountdown()
  } catch (error) {
    Message.error(error.message)
  } finally {
    sendingCode.value = false
  }
}

const handleSubmit = async () => {
  try {
    submitting.value = true
    // TODO: 调用重置密码的API
    // await api.resetPassword({
    //   email: form.value.email,
    //   verificationCode: form.value.verificationCode,
    //   newPassword: form.value.newPassword
    // })
    Message.success(t('forgotPassword.resetSuccess'))
    router.push('/login')
  } catch (error) {
    Message.error(error.message || t('forgotPassword.resetFailed'))
  } finally {
    submitting.value = false
  }
}
</script>

<style scoped>
.forgot-password-container {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f2f3f5;
}

.forgot-password-container.light-theme {
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

.forgot-password-box {
  width: 500px;
  padding: 40px;
  background: #ffffff;
  border-radius: 4px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
}

.forgot-password-header {
  text-align: center;
  margin-bottom: 40px;
}

.forgot-password-header h2 {
  color: var(--color-text-1);
  margin-bottom: 16px;
}

.description {
  color: var(--color-text-3);
  margin-bottom: 24px;
}

.logo {
  width: 64px;
  margin-bottom: 16px;
}

.forgot-password-form {
  width: 100%;
}

/* 表单项样式 */
.forgot-password-form :deep(.arco-form-item-label-col) {
  text-align: left;
  color: var(--color-text-2);
  min-width: 85px;
}

/* 处理浏览器自动填充的背景色 */
.forgot-password-form :deep(input:-webkit-autofill),
.forgot-password-form :deep(input:-webkit-autofill:hover),
.forgot-password-form :deep(input:-webkit-autofill:focus),
.forgot-password-form :deep(input:-webkit-autofill:active) {
  -webkit-text-fill-color: var(--color-text-1) !important;
  -webkit-box-shadow: 0 0 0 1000px transparent inset !important;
  transition: background-color 50000s ease-in-out 0s;
  background-color: transparent !important;
}

.forgot-password-form :deep(.arco-input),
.forgot-password-form :deep(.arco-input-wrapper),
.forgot-password-form :deep(.arco-input-inner-wrapper) {
  background-color: transparent !important;
  border-color: var(--color-border) !important;
}

.forgot-password-form :deep(.arco-input) {
  color: var(--color-text-1) !important;
}

.forgot-password-form :deep(.arco-input::placeholder) {
  color: var(--color-text-3) !important;
}

.forgot-password-form :deep(.arco-input-wrapper:hover),
.forgot-password-form :deep(.arco-input-wrapper:focus-within) {
  border-color: var(--color-primary) !important;
  background-color: transparent !important;
}

/* 按钮样式 */
.forgot-password-form :deep(.arco-btn-primary) {
  background-color: #165dff !important;
  border-color: #165dff !important;
  color: #ffffff !important;
}

.forgot-password-form :deep(.arco-btn-primary:hover) {
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