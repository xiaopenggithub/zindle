export default {
  path: '/403',
  name: 'Forbidden',
  component: () => import('@/views/error/403.vue'),
  meta: {
    requiresAuth: false
  }
} 