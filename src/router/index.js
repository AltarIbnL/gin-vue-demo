import Vue from 'vue';
import VueRouter from 'vue-router';
// import Register from '@/views/Register/Register.vue';
// import Login from '@/views/Login/Login.vue';
import store from '@/store';
import HomeView from '../views/HomeView.vue';
import userRoutes from './module/user';

Vue.use(VueRouter);
// 给每个页面做指向
const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue'),
  },
  ...userRoutes,
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});
// to 是去的路由，from为来的路由,这个是在进入用户详情页的时候做判断,如果没有登陆则进入登陆界面
router.beforeEach((to, from, next) => {
  if (to.meta.auth) { // 判断用户是否需要登陆
    // 判断用户是否登录
    if (store.state.userModule.token) {
      // 判断token的有效性，比如有没有过期，需要后台发放token的时候，带上token的有效期
      // 如果token 无效，需要请求token
      next();
    } else {
      // 跳转登陆
      router.push({ name: 'login' });
    }
  } else {
    next();
  }
});

export default router;
