export default {
  path: 'system',
  name: 'System',
  redirect: '/system/users',
  meta: {
    title: 'menu.system',
    requiresAuth: true,
    icon: 'IconSettings',
    order: 99
  },
  children: [
    {
      path: 'users',
      name: 'Users',
      component: () => import('@/views/system/users/Index.vue'),
      meta: {
        title: 'menu.system.users',
        requiresAuth: true,
        roles: ['admin']
      }
    },
    {
      path: 'roles',
      name: 'Roles',
      component: () => import('@/views/system/roles/Index.vue'),
      meta: {
        title: 'menu.system.roles',
        requiresAuth: true,
        roles: ['admin']
      }
    },
    {
      path: 'permissions',
      name: 'Permissions',
      component: () => import('@/views/system/permissions/Index.vue'),
      meta: {
        title: 'menu.system.permissions',
        requiresAuth: true,
        roles: ['admin']
      }
    },
    {
      path: 'menus',
      name: 'Menus',
      component: () => import('@/views/system/menus/Index.vue'),
      meta: {
        title: 'menu.system.menus',
        requiresAuth: true,
        roles: ['admin']
      }
    },
    {
      path: 'logs',
      name: 'Logs',
      component: () => import('@/views/system/logs/Index.vue'),
      meta: {
        title: 'menu.system.logs',
        requiresAuth: true,
        roles: ['admin']
      }
    }
  ]
} 