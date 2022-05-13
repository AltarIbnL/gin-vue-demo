<template>
  <div>
    <b-navbar
      toggleable="lg"
      type="dark"
      variant="info"
    >
    <b-container>
<!--      <li class="nav-item active" @click="$router.replace({name:'home'})">-->
<!--        <a class="navbar-brand" href="#">U-Learn</a>-->
<!--      </li>-->
        <b-navbar-brand @click="$router.push({name: 'home'})">U-Learn</b-navbar-brand>
        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
      <b-collapse
        id="nav-collapse"
        is-nav
      >
      <b-navbar-nav class="ml-auto">
        <b-nav-item-dropdown
          right
          v-if="userinfo"
        >
          <!-- Using 'button-content' slot -->
          <template v-slot:button-content>
            <em>{{userinfo.name}}</em>
          </template>
          <b-dropdown-item @click="$router.replace({name:'profile'})">个人主页</b-dropdown-item>
          <b-dropdown-item @click="loggout">退出</b-dropdown-item>
        </b-nav-item-dropdown>
<!--        没有 userinfo才显示注册和登录-->
        <div v-if="!userinfo">
<!--          在登陆页面不显示登录按钮-->
          <b-nav-item v-if="$route.name !='login'" @click="$router.replace({name:'login'})">登录</b-nav-item>
          <b-nav-item v-if="$route.name !='register'" @click="$router.replace({name:'register'})">注册</b-nav-item>
        </div>
      </b-navbar-nav>
      </b-collapse>
      </b-container>
    </b-navbar>
  </div>
</template>
<script>
import { mapState,mapActions} from 'vuex';
export default {
  computed: mapState({
    // 从缓存中获取改为从vuex 中获取，因为vuex是响应式的，
    userinfo: (state) => state.userModule.userInfo,
    // return this.$store.state.userModule.userInfo;
  }),
  methods: mapActions('userModule', ['loggout']), //从userModule里面找到 loggout函数

};
</script>
<style lang="scss" scoped></style>
