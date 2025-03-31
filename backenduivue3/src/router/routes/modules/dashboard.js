export default {
  path: 'dashboard',
  name: 'Dashboard',
  component: () => import('@/views/dashboard/Index.vue'),
  meta: {
    title: 'menu.dashboard',
    requiresAuth: true,
    icon: 'icon-dashboard',
    order: 1
  }
} 