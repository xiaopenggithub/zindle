<template>
  <div class="profile-page">
    <div class="page-content">
      <div class="profile-card">
        <div class="profile-header">
          <a-avatar :size="64" :image-url="userAvatar" />
          <div class="profile-info">
            <h3>{{ username }}</h3>
            <p>{{ email }}</p>
          </div>
        </div>
        <a-divider />
        <a-form :model="form" @submit="handleSubmit" class="profile-form">
          <a-form-item 
            field="username" 
            :label="$t('profile.username')"
            :rules="[{ required: true, message: $t('profile.usernamePlaceholder') }]"
          >
            <a-input 
              v-model="form.username" 
              :placeholder="$t('profile.usernamePlaceholder')" 
            />
          </a-form-item>
          <a-form-item 
            field="email" 
            :label="$t('profile.email')"
            :rules="[
              { required: true, message: $t('profile.emailPlaceholder') },
              { type: 'email', message: $t('register.emailInvalid') }
            ]"
          >
            <a-input 
              v-model="form.email" 
              :placeholder="$t('profile.emailPlaceholder')" 
            />
          </a-form-item>
          <a-form-item 
            field="currentPassword" 
            :label="$t('profile.currentPassword')"
            :rules="[{ required: true, message: $t('profile.currentPasswordRequired') }]"
          >
            <a-input-password 
              v-model="form.currentPassword" 
              :placeholder="$t('profile.currentPasswordPlaceholder')" 
            />
          </a-form-item>
          <a-form-item 
            field="newPassword" 
            :label="$t('profile.newPassword')"
            :rules="[{ required: true, message: $t('profile.newPasswordRequired') }]"
          >
            <a-input-password 
              v-model="form.newPassword" 
              :placeholder="$t('profile.newPasswordPlaceholder')"
            />
          </a-form-item>
          <a-form-item 
            field="confirmPassword" 
            :label="$t('profile.confirmPassword')"
            :rules="[
              { required: true, message: $t('profile.confirmPasswordRequired') },
              { validator: validatePassword, message: $t('profile.passwordNotMatch') }
            ]"
          >
            <a-input-password 
              v-model="form.confirmPassword"
              :placeholder="$t('profile.confirmPasswordPlaceholder')" 
            />
          </a-form-item>
          <div class="form-actions">
            <a-button type="primary" html-type="submit">
              {{ $t('profile.save') }}
            </a-button>
          </div>
        </a-form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useStore } from '@/store'
import { Message } from '@arco-design/web-vue'
import { useI18n } from 'vue-i18n'

const store = useStore()
const { t } = useI18n()

const userAvatar = computed(() => {
  return store.user?.avatar || 'https://avatars.githubusercontent.com/u/1?v=4'
})

const username = computed(() => store.user?.username || '')
const email = computed(() => store.user?.email || '')

const form = ref({
  username: username.value,
  email: email.value,
  currentPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const validatePassword = (value) => value === form.value.newPassword

const handleSubmit = async () => {
  try {
    // TODO: Implement profile update logic
    Message.success(t('profile.updateSuccess'))
  } catch (error) {
    Message.error(t('profile.updateFailed'))
  }
}
</script>

<style scoped>
.profile-page {
  height: 100%;
  background: var(--color-bg-1);
}

.page-content {
  padding: 24px;
  max-width: 800px;
  margin: 0 auto;
}

.profile-card {
  background: var(--color-bg-2);
  border: 1px solid var(--color-border);
  border-radius: 4px;
  padding: 24px;
}

.profile-header {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 24px;
}

.profile-info {
  h3 {
    margin: 0;
    color: var(--color-text-1);
  }
  p {
    margin: 4px 0 0;
    color: var(--color-text-3);
  }
}

.form-actions {
  margin-top: 24px;
  text-align: right;
}
/* 处理浏览器自动填充的背景色 */
.profile-form :deep(input:-webkit-autofill),
.profile-form :deep(input:-webkit-autofill:hover),
.profile-form :deep(input:-webkit-autofill:focus),
.profile-form :deep(input:-webkit-autofill:active),
.profile-form :deep(textarea:-webkit-autofill),
.profile-form :deep(textarea:-webkit-autofill:hover),
.profile-form :deep(textarea:-webkit-autofill:focus),
.profile-form :deep(select:-webkit-autofill),
.profile-form :deep(select:-webkit-autofill:hover),
.profile-form :deep(select:-webkit-autofill:focus) {
  -webkit-text-fill-color: var(--color-text-1) !important;
  -webkit-box-shadow: 0 0 0 1000px transparent inset !important;
  transition: background-color 50000s ease-in-out 0s;
  background-color: transparent !important;
}

.profile-form :deep(.arco-input),
.profile-form :deep(.arco-input-password) {
  color: var(--color-text-1) !important;
}

.profile-form :deep(.arco-input-wrapper),
.profile-form :deep(.arco-input-inner-wrapper) {
  background-color: transparent !important;
}

.profile-form :deep(.arco-input::placeholder),
.profile-form :deep(.arco-input-password input::placeholder) {
  color: var(--color-text-3);
}
</style> 