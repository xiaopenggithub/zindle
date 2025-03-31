export default {
  path: 'orders',
  name: 'Orders',
  redirect: '/orders/manage',
  meta: {
    title: 'menu.orders',
    requiresAuth: true,
    icon: 'IconOrderedList',
    order: 3
  },
  children: [
    {
      path: 'manage',
      name: 'OrderManage',
      component: () => import('@/views/orders/Index.vue'),
      meta: {
        title: 'menu.orders.manage',
        requiresAuth: true,
        roles: ['admin']
      }
    },
    {
      path: 'my',
      name: 'MyOrders',
      component: () => import('@/views/memberOrders/Index.vue'),
      meta: {
        title: 'menu.orders.my',
        requiresAuth: true,
        requiresMember: true
      }
    }
  ]
} 