import axios from 'axios'
import { Message } from '@arco-design/web-vue'

const service = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '',
  timeout: 15000
})

// Request interceptor
service.interceptors.request.use(
  config => {
    // You can add auth token here if needed
    return config
  },
  error => {
    console.error('Request error:', error)
    return Promise.reject(error)
  }
)

// Response interceptor
service.interceptors.response.use(
  response => {
    const res = response.data
    
    if (res.code !== 200) {
      Message.error(res.message || 'Error')
      return Promise.reject(new Error(res.message || 'Error'))
    }
    
    return response
  },
  error => {
    console.error('Response error:', error)
    Message.error(error.message || 'Request failed')
    return Promise.reject(error)
  }
)

export default service 