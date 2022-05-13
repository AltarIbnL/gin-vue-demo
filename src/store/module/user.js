import storageService from '../../../service/storageService';
import userService from '../../../service/userService';
import router from "@/router";

const userModule = {
  namespaced: true,
  state: {
    token: storageService.get(storageService.USER_TOKEN),
    // eslint-disable-next-line max-len
    userInfo: storageService.get(storageService.USER_TOKEN) ? JSON.parse(storageService.get(storageService.USER_INFO)) : null,
  },
  mutations: {
    SET_TOKEN(state, token) {
      // 更新本地缓存
      storageService.set(storageService.USER_TOKEN, token);
      // 更新state
      state.token = token;
    },
    SET_USERINFO(state, userInfo) {
      // 更新本地缓存
      storageService.set(storageService.USER_INFO, JSON.stringify(userInfo));
      // 更新state
      state.userInfo = userInfo;
    },
  },
  actions: {
    register(context, { name, telephone, password }) {
      return new Promise((resolve, reject) => {
        // userService.register(this.user).then((res) => { // 前面的register是用户输入完之后发给后端，then是执行接收到后端返回的数据以后,在后端的register之后
        userService.register({ name, telephone, password }).then((res) => {
          // 保存token
          console.log(res.data);
          context.commit('SET_TOKEN', res.data.token); // 将token 储存在vuex中，代码在store/index 和module user.js中
          // this.$store.commit('userModule/SET_TOKEN', res.data.token);  // 将token 储存在vuex中，代码在store/index 和module user.js中
          return userService.info();  // 读取info列表的信息，res的内容就是 info()读取的内容
        }).then((res) => { // 从后端的info页面获取数据, 链式调用
          // 保存用户信息
          // this.SET_USERINFO(response.data.data.user); // 将用户信息储存在vuex中，代码在store/index 和module user.js中
          context.commit('SET_USERINFO', res.data.data.user); // 将用户信息储存在vuex中，代码在store/index 和module user.js中
          resolve(res);
        }).catch((err) => {
          reject(err);
        });
      });
    },
    login(context, { telephone, password }) {
      return new Promise((resolve, reject) => {
        // userService.register(this.user).then((res) => { // 前面的register是用户输入完之后发给后端，then是执行接收到后端返回的数据以后,在后端的register之后
        userService.login({ telephone, password }).then((res) => {
          // 保存token
          console.log(res.data);
          context.commit('SET_TOKEN', res.data.token); // 将token 储存在vuex中，代码在store/index 和module user.js中
          return userService.info();
        }).then((res) => { // 从后端的info页面获取数据, 链式调用
          // 保存用户信息
          context.commit('SET_USERINFO', res.data.data.user); // 将用户信息储存在vuex中，代码在store/index 和module user.js中
          resolve(res);
        }).catch((err) => {
          reject(err);
        });
      });
    },
    loggout(context) {
      // 清除token
      context.commit('SET_TOKEN', '');
      storageService.set(storageService.USER_TOKEN, '');
      // 清楚用户信息
      context.commit('SET_USERINFO', '');
      storageService.set(storageService.USER_INFO, '');
      window.location.reload();
      // router.push({ name: 'home' });// 跳转主界面
    },
  },
};
export default userModule;
