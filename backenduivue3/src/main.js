import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ArcoVue from '@arco-design/web-vue'
import App from '@/App.vue'
import router from '@/router'
import i18n from '@/i18n'
import '@arco-design/web-vue/dist/arco.css'
import '@/styles/index.css'

const app = createApp(App)
const pinia = createPinia()

app.use(ArcoVue)
app.use(router)
app.use(i18n)
app.use(pinia)

// 初始化主题
const theme = localStorage.getItem('theme') || 'light'
document.body.setAttribute('arco-theme', theme)

app.mount('#app')
