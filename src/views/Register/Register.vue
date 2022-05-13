<template>
  <div class="login">
    <b-row class="mt-5">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
    <b-card title="注册">
    <b-form >
      <b-form-group label="姓名">
        <!-- 定义一个v的对象，根据我们定义的validation生成了属性，每个属性都有自己的model  -->
       <b-form-input
          v-model="$v.user.name.$model"
          type="text"
          required
          placeholder="输入用户的名称(选填)"
        ></b-form-input>
      </b-form-group>
      <b-form-group label="手机号">
        <b-form-input
          v-model="$v.user.telephone.$model"
          type="number"
          required
          placeholder="输入手机号"
          :state="validateState('telephone')"
        ></b-form-input>
        <b-form-invalid-feedback :state="validateState('telephone')">
          手机号必须为11位
        </b-form-invalid-feedback>
      </b-form-group>
      <b-form-group label="密码">
        <b-form-input
          v-model="$v.user.password.$model"
          type="password"
          required
          placeholder="输入您的密码"
          :state="validateState('password')"
        ></b-form-input>
        <b-form-invalid-feedback :state="validateState('password')">
          密码必须大于等于6位
        </b-form-invalid-feedback>
      </b-form-group>
      <b-form-group>
       <b-button @click="register" variant="primary" block>注册</b-button>  <!-- block让按钮占满容器-->
      </b-form-group>
    </b-form>
    </b-card>
      </b-col>
    </b-row>
  </div>
</template>
<script>
// eslint-disable-next-line import/no-extraneous-dependencies
import { required, minLength,maxLength } from 'vuelidate/lib/validators';
// import { required, minLength } from 'vuelidate/lib/validators';
import {mapMutations,mapActions} from 'vuex';
import userModule from "@/store/module/user";
// import customValidator from '@/helper/validator';
// import storageService from '../../../service/storageService';
import userService from '../../../service/userService';
export default {
  data() {
    return {
      user: {
        name: '',
        telephone: '',
        password: '',
      },
      validation: null,
    };
  },
  // vuelidate来验证当前的密码为多少位
  validations: {
    user: {
      name: {
      },
      telephone: {
        required,
        // telephone: customValidator.telephoneValidator,
        minLength: minLength(11),
        maxLength: maxLength(11),
      },
      password: {
        required,
        minLength: minLength(6),
      },
    },
  },
  methods: {
    // ...mapMutations('userModule',['SET_TOKEN','SET_USERINFO']),
    ...mapActions('userModule', {userRegister: 'register' }),
    validateState(name) {
      // 这里是es6 解构赋值
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      // 验证数据,在没有用户与表单交互的情况下依旧可以触发表单验证
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求,监听9090端口
      // const api = 'http://localhost:9090/api/auth/register';
      // 返回用户填入的 this.user 包含  name telephone password
      // 这一步将register 的操作放入./store/module/user.js 里面，包括用户token的保存和用户信息的保存
      // this.$store.dispatch('userModule/register', this.user).then(() => {
      // 这种方式与上面dispatch实现的效果是一样的
      this.userRegister(this.user).then(() => {
        // 跳转主页
        this.$router.replace({ name: 'home' });
      }).catch((err) => {
        // 跳转失败,并出现弹窗显示数据验证失败
        console.log('err:', err.response.data.msg);
        this.$bvToast.toast(err.response.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });
      console.log('register');
    },
  },
};
</script>
<style lang="scss" scoped></style>
