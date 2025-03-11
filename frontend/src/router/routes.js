const routes = [
  {
    path: '/',
    children: [
      {
        path: 'login',
        component: () => import('src/pages/LoginPage.vue'),
        meta: {
          requiresAuth: false,
        },
      },
      {
        path: 'register',
        component: () => import('src/pages/RegisterPage.vue'),
        meta: {
          requiresAuth: false,
        },
      },
      {
        path: '',
        component: () => import('src/pages/IndexPage.vue'),
        meta: {
          requiresAuth: false,
        },
      },
    ],
  },

  {
    path: '/',
    component: () => import('src/layouts/userLayout.vue'),
    children: [
      { path: 'create/user-tickets', component: () => import('src/pages/TicketPage.vue') },
      { path: 'user-page', component: () => import('src/pages/UserPage.vue') },
      { path: 'user/tickets', component: () => import('src/pages/UserTicketPage.vue') },
      { path: 'settings', component: () => import('src/pages/SettingsPage.vue') },

      { path: 'deneme', component: () => import('src/pages/DenemePage.vue') }, // Deneme icin chatgptye yaptirdim
    ],
  },

  {
    path: '/',
    component: () => import('src/layouts/adminLayout.vue'),
    children: [
      { path: 'admin-page', component: () => import('src/pages/AdminPage.vue') },
      { path: 'settings', component: () => import('src/pages/SettingsPage.vue') },
      { path: 'clients/list', component: () => import('src/pages/CustomersPage.vue') },
      { path: 'clients/ticket', component: () => import('src/pages/AdminTicketPage.vue') },
      { path: 'admin/settings', component: () => import('src/pages/SettingsPage.vue') },
      
    ],
  },

  // {
  //   path: '/login',
  //   component: () => import('src/pages/LoginPage.vue'),
  //   meta: {
  //     requiresAuth: false,
  //   },
  // },
  // {
  //   path: '/register',
  //   component: () => import('src/pages/RegisterPage.vue'),
  //   meta: {
  //     requiresAuth: false,
  //   },
  // },
  // {
  //   path: '',
  //   component: () => import('src/pages/IndexPage.vue'),
  //   meta: {
  //     requiresAuth: false,
  //   },
  // },

  // {
  //   path: '/register',
  //   component: () => import('src/pages/RegisterPage.vue'),
  //   meta: {
  //     requiresAuth: false,
  //   },
  // },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },

  {
    path: '/justTrying',
    component: () => import('src/layouts/denemeLayout.vue'),
  },
]

export default routes
