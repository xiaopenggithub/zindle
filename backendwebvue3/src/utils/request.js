import axios from 'axios' // 引入axios
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { emitter } from '@/utils/bus.js'
import router from '@/router/index'
import i18n from '@/i18n' // added by mohamed hassan to multilangauge

const service = axios.create({
  baseURL: import.meta.env.VITE_BASE_API,
  timeout: 99999
})
let acitveAxios = 0
let timer
const showLoading = () => {
  acitveAxios++
  if (timer) {
    clearTimeout(timer)
  }
  timer = setTimeout(() => {
    if (acitveAxios > 0) {
      emitter.emit('showLoading')
    }
  }, 400)
}

const closeLoading = () => {
  acitveAxios--
  if (acitveAxios <= 0) {
    clearTimeout(timer)
    emitter.emit('closeLoading')
  }
}
// http request 拦截器
service.interceptors.request.use(
  config => {
    if (!config.donNotShowLoading) {
      showLoading()
    }
    const userStore = useUserStore()
    config.headers = {
      'Content-Type': 'application/json',
      'authorization': userStore.token,
      'x-user-id': userStore.userInfo.ID,
      'Accept-Language': userStore.language, // added by mohame hassan to allow store selected language for multilanguage support.
      ...config.headers
    }
    return config
  },
  error => {
    if (!error.config.donNotShowLoading) {
      closeLoading()
    }
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)

// http response 拦截器
service.interceptors.response.use(
  response => {
    const userStore = useUserStore()
    if (!response.config.donNotShowLoading) {
      closeLoading()
    }
    if (response.headers['new-token']) {
      userStore.setToken(response.headers['new-token'])
    }
    if (response.data.code === 200 || response.headers.success === 'true') {
      if (response.headers.message) {
        response.data.message = decodeURI(response.headers.message)
      }
      return response.data
    } else {
      ElMessage({
        showClose: true,
        message: response.data.message || decodeURI(response.headers.message),
        type: 'error'
      })
      if (response.data.code === 401 || (response.data.data && response.data.data.reload)) {
        userStore.token = ''
        localStorage.clear()
        router.push({ name: 'Login', replace: true })
      }
      return response.data.message ? response.data : response
    }
  },
  error => {
    if (!error.config.donNotShowLoading) {
      closeLoading()
    }

    if (!error.response) {
      ElMessageBox.confirm(`
        <p>检测到请求错误</p>
        <p>${error}</p>
        `, '请求报错', {
        dangerouslyUseHTMLString: true,
        distinguishCancelAndClose: true,
        confirmButtonText: '稍后重试',
        cancelButtonText: '取消'
      })
      return
    }

    switch (error.response.status) {
      case 500:
        ElMessageBox.confirm(`
        <p>检测到接口错误${error}</p>
        <p>错误码<span style="color:red"> 500 </span>：此类错误内容常见于后台panic，请先查看后台日志，如果影响您正常使用可强制登出清理缓存</p>
        `, '接口报错', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: i18n.t('general.confirm'),
          cancelButtonText: i18n.t('general.cancel')
        })
          .then(() => {
            const userStore = useUserStore()
            userStore.token = ''
            localStorage.clear()
            router.push({ name: 'Login', replace: true })
          })
        break
      case 404:
        ElMessageBox.confirm(`
          <p>检测到接口错误${error}</p>
          <p>错误码<span style="color:red"> 404 </span>：此类错误多为接口未注册（或未重启）或者请求路径（方法）与api路径（方法）不符--如果为自动化代码请检查是否存在空格</p>
          `, '接口报错', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: i18n.t('general.confirm'),
          cancelButtonText: i18n.t('general.cancel')
        })
        break
    }

    return error
  }
)
export default service
