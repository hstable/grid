import Vue from 'vue'
// eslint-disable-next-line no-unused-vars
import { ToastProgrammatic, Modal } from 'buefy'
import ModalLogin from '@/components/modalLogin'
export default function({ $axios, redirect }) {
  $axios.onRequest((config) => {
    // eslint-disable-next-line no-prototype-builtins
    if (localStorage.hasOwnProperty('token')) {
      config.headers.Authorization = `${localStorage.token}`
    }
  })
  $axios.onResponseError((err) => {
    // console.log('!!', err.name, err.message)
    // console.log(Object.assign({}, err))
    if (err.response && err.response.status === 401) {
      // 401未授权
      new Vue({
        components: { Modal, ModalLogin },
        data: () => ({
          show: true
        }),
        render() {
          return (
            <b-modal
              active={this.show}
              trap-focus={true}
              has-modal-card={true}
              aria-role="dialog"
              aria-modal={true}
              full-screen={false}
              style="z-index:1000"
              class="modal-login"
              id="login"
            >
              <ModalLogin
                onClose={() => {
                  this.show = false
                }}
              />
            </b-modal>
          )
        }
      }).$mount('#login')
    } else {
      // ToastProgrammatic.open({
      //   message: err.message,
      //   type: 'is-warning',
      //   position: 'is-top',
      //   duration: 5000,
      //   queue: false
      // })
    }
  })
}
