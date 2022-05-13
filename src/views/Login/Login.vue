<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="登陆">
          <b-form >
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
              <b-button @click="login" variant="primary" block>登陆</b-button>  <!-- block让按钮占满容器-->
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
import {mapMutations,mapActions} from 'vuex';
// const telephoneValidator = (value) => /^\d(11)$/.test(value);
import customValidator from '@/helper/validator';
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
    ...mapActions('userModule', { userlogin: 'login' }),
    validateState(name) {
      // 这里是es6 解构赋值
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    login() {
      // 验证数据,在没有用户与表单交互的情况下依旧可以触发表单验证
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      //请求
      this.userlogin(this.user).then(() => {
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
      console.log('login');
    },
  },
};
</script>
<style lang="scss" scoped></style>
