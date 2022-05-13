const userRoutes = [
  {
    path: '/register',
    name: 'register',
    // 惰性的加载页面,不是一次性全部都加载.只有在调用的时候才加载
    component: () => import('@/views/Register/Register.vue'),
  },
  {
    path: '/login',
    name: 'login',
    // eslint-disable-next-line import/no-unresolved
    component: () => import('@/views/Login/Login.vue'),
  },
  {
    path: '/profile',
    name: 'profile',
    meta: {
      auth: true, // 需要登陆的用户才能登入
    },
    // eslint-disable-next-line import/no-unresolved
    component: () => import('@/views/Profile/Profile.vue'),
  },
];
export default userRoutes;
