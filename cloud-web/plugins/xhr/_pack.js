import { ToastProgrammatic } from 'buefy'
import { handleResponse } from '@/assets/js/utils'

export default function({ $axios }) {
  function pack(params) {
    return new Promise((resolve, reject) => {
      $axios(params)
        .then((res) => {
          handleResponse(
            res,
            this,
            () => {
              resolve(res)
            },
            () => {
              ToastProgrammatic.open({
                message: res.data.message,
                type: 'is-warning',
                position: 'is-top',
                queue: false,
                duration: 5000
              })
              reject(res.data.message)
            }
          )
        })
        .catch((err) => {
          function notsuretip() {
            ToastProgrammatic.open({
              message: '请求服务器时发生网络错误',
              type: 'is-warning',
              position: 'is-top',
              queue: false,
              duration: 5000
            })
          }
          if (!err.response) {
            notsuretip()
          } else {
            switch (err.response.status) {
              case 401:
                break
              case 400:
                ToastProgrammatic.open({
                  message: err.response.data.message,
                  type: 'is-warning',
                  position: 'is-top',
                  queue: false,
                  duration: 5000
                })
                break
              default:
                notsuretip()
            }
          }

          reject(err)
        })
    })
  }

  return pack
}
