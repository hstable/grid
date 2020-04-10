<template>
  <div id="modal-login" class="modal-card" style="max-width: 450px;margin:auto">
    <header class="modal-card-head">
      <p class="modal-card-title">
        登录
      </p>
    </header>
    <section class="modal-card-body">
      <p style="text-align: center">
        <img src="~assets/logo.png" alt="V2RayA" />
      </p>
      <b-field label="用户名" type="is-success">
        <b-input v-model="username" @keyup.enter.native="handleEnter"></b-input>
      </b-field>
      <b-field label="密码" type="is-success">
        <b-input
          v-model="password"
          type="password"
          maxlength="32"
          @keyup.enter.native="handleEnter"
        ></b-input>
      </b-field>
    </section>
    <footer class="modal-card-foot flex-end">
      <b-button class="is-primary" @click="handleClickSubmit">
        登录
      </b-button>
    </footer>
  </div>
</template>

<script>
export default {
  name: 'ModalLogin',
  data: () => ({
    username: '',
    password: '',
    loading: null
  }),
  methods: {
    handleClickSubmit() {
      // login
      this.loading = this.$buefy.loading.open({
        container: document.querySelector('#modal-login')
      })
      this.$xhr
        .login(this.username, this.password)
        .then((res) => {
          window.location.reload()
        })
        .catch(() => {
          this.loading.close()
        })
    },
    handleEnter() {
      this.handleClickSubmit()
    }
  }
}
</script>

<style lang="scss">
.modal-login .modal-background {
  background-color: rgba(10, 10, 10, 0.95) !important;
}
</style>
<style lang="scss" scoped>
.after-line-dot5 {
  font-size: 14px;
  p {
    font-size: 14px;
  }
}
.flex-end {
  justify-content: flex-end;
}
</style>
