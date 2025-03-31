export default {
  path: 'members',
  name: 'Members',
  component: () => import('@/views/members/Index.vue'),
  meta: {
    title: 'menu.members',
    requiresAuth: true,
    icon: 'icon-user-group',
    order: 2
  }
} 