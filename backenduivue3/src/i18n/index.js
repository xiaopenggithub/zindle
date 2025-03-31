import { createI18n } from 'vue-i18n'
import zhCN from '@/locales/zh-CN'
import enUS from '@/locales/en-US'
import zhTW from '@/locales/zh-TW'
import jaJP from '@/locales/ja-JP'
import koKR from '@/locales/ko-KR'
import deDE from '@/locales/de-DE'

const messages = {
  'zh-CN': zhCN,
  'en-US': enUS,
  'zh-TW': zhTW,
  'ja-JP': jaJP,
  'ko-KR': koKR,
  'de-DE': deDE
}

const i18n = createI18n({
  legacy: false, // 使用 Composition API 模式
  locale: localStorage.getItem('locale') || 'zh-CN', // 默认语言
  fallbackLocale: 'en-US', // 回退语言
  messages,
  globalInjection: true, // 全局注入 $t 函数
})

export default i18n 