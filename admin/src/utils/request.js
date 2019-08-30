import axios from 'axios'
import { MessageBox, Message } from 'element-ui'
import store from '@/store'
import { getToken } from '@/utils/auth'

// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000 // request timeout
})
// request interceptor
service.interceptors.request.use(
  config => {
    // do something before request is sent

    const passUrls = ['/user/login', '/user/logout', '/user/info']
    if (store.getters.token && !passUrls.includes(config.url)) {
      // let each request carry token
      // ['X-Token'] is a custom headers key
      // please modify it according to the actual situation
      if (typeof config.params === 'undefined') {
        config.params = { token: getToken() }
      } else {
        config.params.token = getToken()
      }
    }
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
   */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    const res = response.data

    // if the custom code is not 20000, it is judged as an error.
    if (res.code !== 20000) {
      Message({
        message: res.msg || 'Error',
        type: 'error',
        duration: 5 * 1000
      })

      return Promise.reject(new Error(res.msg || 'Error'))
    } else {
      return res
    }
  },
  error => {
    if (!error.response) {
      Message({
        message: error.message,
        type: 'error',
        duration: 5 * 1000
      })
    } else {
      switch (error.response.status) {
        case 401:
          if (error.response.data.code === 40002) {
            MessageBox.confirm(
              `${error.response.data.msg}，你可以停留在本页面或重新登录`,
              {
                confirmButtonText: '重新登录',
                cancelButtonText: '取消',
                type: 'warning'
              }
            )
              .then(() => {
                store.dispatch('user/resetToken').then(() => {
                  location.reload()
                })
              })
              .catch(() => {
                // pass
              })
          } else {
            Message({
              message: error.response.data.msg || 'Error',
              type: 'error',
              duration: 5 * 1000
            })
          }
          break
        default:
          Message({
            message: error.response.data.msg || 'Error',
            type: 'error',
            duration: 5 * 1000
          })
      }
    }
    return Promise.reject(error)
  }
)

export default service
