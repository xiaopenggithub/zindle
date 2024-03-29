import { login, logout, getUserInfo, setSelfInfo } from '@/api/user'
import { jsonInBlacklist } from '@/api/jwt'
import router from '@/router/index'
import { ElLoading, ElMessage } from 'element-plus'
import { defineStore } from 'pinia'
import { ref, computed, watch } from 'vue'
import { useRouterStore } from './router'

export const useUserStore = defineStore('user', () => {
  const loadingInstance = ref(null)

  const userInfo = ref({
    uuid: '',
    nickName: '',
    headerImg: '',
    authority: {},
    sideMode: 'dark',
    activeColor: '#4D70FF',
    baseColor: '#fff'
  })
  const token = ref(window.localStorage.getItem('token') || '')
  const language = ref(window.localStorage.getItem('langauge') || 'en') // added by mohamed hassan to allow store selected language for multilanguage support.

  const setUserInfo = (val) => {
    userInfo.value = val
  }

  const setToken = (val) => {
    token.value = val
  }

  // added by mohame hassan to allow store selected language for multilanguage support.
  const setLanguage = (val) => {
    console.log('setLanguage called with value: ' + val)
    language.value = val
  }

  const getLanguage = () => {
    return language.value
  }

  const NeedInit = () => {
    token.value = ''
    window.localStorage.removeItem('token')
    localStorage.clear()
    router.push({ name: 'Init', replace: true })
  }

  const ResetUserInfo = (value = {}) => {
    userInfo.value = {
      ...userInfo.value,
      ...value
    }
  }
  /* 获取用户信息*/
  const GetUserInfo = async() => {
    const res = await getUserInfo()
    if (res.code === 200) {
      setUserInfo(res.data)
    }
    return res
  }
  /* 登录*/
  const LoginIn = async(loginInfo) => {
    loadingInstance.value = ElLoading.service({
      fullscreen: true,
      text: '登录中，请稍候...',
    })
    try {
      const res = await login(loginInfo)
      if (res.code === 200) {        
        console.log('----userinfo----')
        console.log(res.data.user)
        console.log('----userinfo----')
        setUserInfo(res.data.user)
        setToken(res.data.token)
        const routerStore = useRouterStore()
        await routerStore.SetAsyncRouter()
        const asyncRouters = routerStore.asyncRouters
        asyncRouters.forEach(asyncRouter => {
          router.addRoute(asyncRouter)
        })        
        console.log('userInfo.value.authority.defaultRouter:'+userInfo.value.authority.defaultRouter)
        await router.push({ name: userInfo.value.authority.defaultRouter })
        loadingInstance.value.close()
        return true
      }
    } catch (e) {
      loadingInstance.value.close()
    }
    loadingInstance.value.close()
  }
  /* 登出*/
  const LoginOut = async() => {
    const res = await logout()        
    if (res.code === 200) {
      token.value = ''
      sessionStorage.clear()
      localStorage.clear()
      router.push({ name: 'Login', replace: true })
      window.location.reload()
    }
  }
  /* 清理数据 */
  const ClearStorage = async() => {
    token.value = ''
    sessionStorage.clear()
    localStorage.clear()
  }
  /* 设置侧边栏模式*/
  const changeSideMode = async(data) => {
    const res = await setSelfInfo({ sideMode: data })
    if (res.code === 0) {
      userInfo.value.sideMode = data
      ElMessage({
        type: 'success',
        message: '设置成功'
      })
    }
  }

  const mode = computed(() => userInfo.value.sideMode)
  const sideMode = computed(() => {
    if (userInfo.value.sideMode === 'dark') {
      return '#191a23'
    } else if (userInfo.value.sideMode === 'light') {
      return '#fff'
    } else {
      return userInfo.value.sideMode
    }
  })
  const baseColor = computed(() => {
    if (userInfo.value.sideMode === 'dark') {
      return '#fff'
    } else if (userInfo.value.sideMode === 'light') {
      return '#191a23'
    } else {
      return userInfo.value.baseColor
    }
  })
  const activeColor = computed(() => {
    if (userInfo.value.sideMode === 'dark' || userInfo.value.sideMode === 'light') {
      return '#4D70FF'
    }
    return userInfo.activeColor
  })

  watch(() => token.value, () => {
    window.localStorage.setItem('token', token.value)
  })

  return {
    userInfo,
    token,
    language, // added by mohame hassan to allow store selected language for multilanguage support.
    NeedInit,
    ResetUserInfo,
    GetUserInfo,
    LoginIn,
    LoginOut,
    setLanguage, // added by mohame hassan to allow store selected language for multilanguage support.
    getLanguage, // added by mohame hassan to allow store selected language for multilanguage support.
    changeSideMode,
    mode,
    sideMode,
    setToken,
    baseColor,
    activeColor,
    loadingInstance,
    ClearStorage
  }
})
