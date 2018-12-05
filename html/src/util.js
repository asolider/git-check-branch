import axios from 'axios'
import qs from 'qs'

axios.interceptors.request.use(
  // Do something before request is sent
  config => {
    if (config.method == 'post') {
        config.headers['Content-Type'] = 'application/x-www-form-urlencoded';
        config.data = qs.stringify(config.data);
    }
    return config;
  },
  error => {
    // Do something with request error
    return Promise.reject(error);
  });

// Add a response interceptor
axios.interceptors.response.use(
  response => {
    // Do something with response data
    if (response.status == 200) {
        return response.data;
    } else {
        return Promise.reject("api响应异常");
    }
  },
  error => {
    // Do something with response error
    return Promise.reject(error);
  });

export default {
    http: axios,
}
