import request from '@/utils/request';

// 用户注册,request 是之前创建的axios
const register = ({ name, telephone, password }) => request.post('auth/register', { name, telephone, password });

// 获取用户信息
const info = () => request.get('auth/info');

// 用户登陆,request 是之前创建的axios
const login = ({ telephone, password }) => request.post('auth/login', { telephone, password });

export default {
  register,
  login,
  info,
};
