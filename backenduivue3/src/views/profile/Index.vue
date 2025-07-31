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
        <a-form 
          ref="formRef"
          :model="form" 
          :rules="rules"
          @submit="handleSubmit" 
          class="profile-form"
        >
          <a-form-item 
            field="nickname" 
            :label="$t('profile.nickname')"
            validate-trigger="blur"
          >
            <a-input 
              v-model="form.nickname" 
              :placeholder="$t('profile.nicknamePlaceholder')" 
            />
          </a-form-item>
          <div class="form-actions">
            <a-button type="outline" @click="showPasswordModal">
              {{ $t('profile.changePassword') }}
            </a-button>
            <a-button type="primary" html-type="submit" :loading="loading">
              {{ $t('profile.save') }}
            </a-button>
          </div>
        </a-form>

        <!-- 修改密码弹窗 -->
        <a-modal
          v-model:visible="passwordModalVisible"
          :title="$t('profile.changePassword')"
          @before-ok="handlePasswordModalOk"
          @cancel="handlePasswordModalCancel"
          :mask-closable="false"
          :ok-loading="passwordLoading"
        >
          <a-form 
            ref="passwordFormRef"
            :model="passwordForm" 
            :rules="passwordRules"
          >
            <a-form-item
              field="currentPassword"
              :label="$t('profile.currentPassword')"
              validate-trigger="blur"
            >
              <a-input-password
                v-model="passwordForm.currentPassword"
                :placeholder="$t('profile.currentPasswordPlaceholder')"
              />
            </a-form-item>
            <a-form-item
              field="newPassword"
              :label="$t('profile.newPassword')"
              validate-trigger="blur"
            >
              <a-input-password
                v-model="passwordForm.newPassword"
                :placeholder="$t('profile.newPasswordPlaceholder')"
              />
            </a-form-item>
            <a-form-item
              field="confirmPassword"
              :label="$t('profile.confirmPassword')"
              validate-trigger="blur"
            >
              <a-input-password
                v-model="passwordForm.confirmPassword"
                :placeholder="$t('profile.confirmPasswordPlaceholder')"
              />
            </a-form-item>
          </a-form>
        </a-modal>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useStore } from '@/store'
import { Message } from '@arco-design/web-vue'
import { useI18n } from 'vue-i18n'
import { useRouter } from 'vue-router'
import { changePassword, updateInfo } from '@/api/user'

const store = useStore()
const router = useRouter()
const { t } = useI18n()

const formRef = ref(null)
const passwordFormRef = ref(null)
const loading = ref(false)
const passwordLoading = ref(false)

const userAvatar = computed(() => {
  return store.user?.avatar || 'https://avatars.githubusercontent.com/u/1?v=4'
})

const nickname = computed(() => store.user?.nickname || '')
// const email = computed(() => store.user?.email || '')

const form = ref({
  nickname: nickname.value,  
})

const rules = {
  nickname: [
    { required: true, message: t('profile.nicknameRequired') },
    { minLength: 3, message: t('profile.nicknameeMinLength') }
  ],
  // email: [
  //   { required: true, message: t('profile.emailRequired') },
  //   { type: 'email', message: t('profile.emailInvalid') }
  // ]
}

const passwordModalVisible = ref(false)
const passwordForm = ref({
  currentPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const passwordRules = {
  currentPassword: [
    { required: true, message: t('profile.currentPasswordRequired') },
    { minLength: 8, message: t('profile.passwordMinLength') }
  ],
  newPassword: [
    { required: true, message: t('profile.newPasswordRequired') },
    { minLength: 8, message: t('profile.passwordMinLength') }
  ],
  confirmPassword: [
    { required: true, message: t('profile.confirmPasswordRequired') },
    { minLength: 8, message: t('profile.passwordMinLength') },
    {
      validator: (value, cb) => {
        if (value !== passwordForm.value.newPassword) {
          cb(t('profile.passwordsDoNotMatch'))
        } else {
          cb()
        }
      }
    }
  ]
}

const showPasswordModal = () => {
  passwordModalVisible.value = true
}

const handlePasswordModalCancel = () => {
  passwordForm.value = {
    currentPassword: '',
    newPassword: '',
    confirmPassword: ''
  }
  passwordFormRef.value?.resetFields()
  passwordModalVisible.value = false
}

const handlePasswordModalOk = async (done) => {
  try {
    // 先进行表单验证     
    const validRes = await passwordFormRef.value.validate()
    // validRes 为 undefined 时表示验证通过
    if (validRes) {
      // 验证失败
      done(false)
      return
    }
    
    // 验证通过后再调用接口
    passwordLoading.value = true
    const res = await changePassword({
      currentPassword: passwordForm.value.currentPassword,
      newPassword: passwordForm.value.newPassword,
      confirmPassword: passwordForm.value.confirmPassword
    })

    if (res.data.code === 200) {
      Message.success(t('profile.passwordChangeSuccess'))
      handlePasswordModalCancel()
      await store.logout()
      router.push('/login')
      done()
    } else {
      throw new Error(res.data.message || t('profile.passwordChangeFailed'))
    }
  } catch (error) {
    // 验证失败或 API 调用失败
    done(false)
    // API 调用失败时显示错误消息
    if (error.response) {
      Message.error(error.message || t('profile.passwordChangeFailed'))
    }
  } finally {
    passwordLoading.value = false
  }
}

const handleSubmit = async ({ values, errors }) => {
  if (errors) {
    return
  }
  
  try {
    loading.value = true
    const res = await updateInfo({
      nickname: values.nickname,      
    })
    if (res.data.code === 200) {  
      Message.success(t('profile.updateSuccess'))
      store.user.nickname = values.nickname
      localStorage.setItem('userInfo', JSON.stringify(store.user))
    } else {
      throw new Error(res.data.message || t('profile.updateFailed'))
    }
  } catch (error) {
    Message.error(t('profile.updateFailed'))
  } finally {
    loading.value = false
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
  display: flex;
  justify-content: flex-end;
  gap: 16px;
  margin-top: 24px;
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