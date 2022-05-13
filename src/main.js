import Vue from 'vue';
// eslint-disable-next-line import/no-extraneous-dependencies
import Vuelidate from 'vuelidate';
// eslint-disable-next-line import/no-extraneous-dependencies
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import App from './App.vue';
import router from './router';
import store from './store';
import './assets/scss/index.scss';
import axios from 'axios';
import VueAxios from 'vue-axios';
// app.js

Vue.config.productionTip = false;

// Install BootstrapVue
Vue.use(BootstrapVue);
Vue.use(Vuelidate);
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin);

// axios
Vue.use(VueAxios, axios);
new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
