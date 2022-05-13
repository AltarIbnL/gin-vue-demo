// eslint-disable-next-line import/no-extraneous-dependencies
import axios from 'axios';
import storageService from '../../service/storageService';
import {config} from "@vue/cli-plugin-eslint/eslintOptions";
// process 是vue底下的一个进程，env是保存的所有环境变量文件，VUE_APP 是VUE创建变量的前缀
// 定义个一个axios
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL,  // 默认地址是VUE_APP_BASE_URL
  timeout: 1000 * 5,
  // headers: { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` },
});
// 添加请求拦截器
// 添加请求拦截器
// eslint-disable-next-line no-shadow
service.interceptors.request.use((config) => {
  // 在发送请求之前做些什么
  Object.assign(config.headers, { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` });
  return config;
}, (error) => {
  // 对请求错误做些什么
  return Promise.reject(error);
});
export default service;
