import axios from "axios";

// create an axios instance
const service = axios.create({
  // baseURL: 'http://localhost:9000', // api 的 base_url
  timeout: 10000 // request timeout
});

// request interceptor
service.interceptors.request.use(
  config => {
    return config;
  },
  error => {
    // Do something with request error
    console.log(error); // for debug
    Promise.reject(error);
  }
);
// response interceptor
service.interceptors.response.use(
  response => {
    const res = response.data;
    return response.data;
  },
  error => {
    return Promise.reject(error);
  }
);

export default service;
