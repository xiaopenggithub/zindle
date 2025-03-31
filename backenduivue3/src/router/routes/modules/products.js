export default {
  path: 'products',
  name: 'Products',
  redirect: '/products/list',
  meta: {
    title: 'menu.products._',
    requiresAuth: true,
    icon: 'IconApps',
    order: 4
  },
  children: [
    {
      path: 'category',
      name: 'ProductCategory',
      component: () => import('@/views/products/Category.vue'),
      meta: {
        title: 'menu.products.category',
        requiresAuth: true
      }
    },
    {
      path: 'list',
      name: 'ProductList',
      component: () => import('@/views/products/List.vue'),
      meta: {
        title: 'menu.products.list',
        requiresAuth: true
      }
    }
  ]
} 