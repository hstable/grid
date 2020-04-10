<template>
  <div class="modal-card" style="max-width: 520px;margin:auto">
    <header class="modal-card-head">
      <p class="modal-card-title">
        {{ data._ ? '添加用户' : '修改用户信息' }}
      </p>
    </header>
    <section class="modal-card-body">
      <b-field label="用户名">
        <b-input ref="1" v-model="data.Sub" expanded required></b-input>
      </b-field>
      <b-field label="密码">
        <b-input
          ref="password"
          v-model="data.Password"
          type="password"
          expanded
          required
          @focus="handleFocusPassword"
          @blur="handleBlurPassword"
        ></b-input>
      </b-field>
      <b-field label="姓名">
        <b-input ref="2" v-model="data.Name" expanded required></b-input>
      </b-field>
      <b-field label="公司">
        <b-select ref="3" v-model="data.CompanyID" expanded required>
          <option v-for="(name, id) of mapCompanyID2Name" :key="id" :value="id"
            >{{ name }}
          </option>
        </b-select>
      </b-field>
      <b-field label="部门">
        <b-input ref="4" v-model="data.Department" expanded></b-input>
      </b-field>
      <b-field label="手机号">
        <b-input
          ref="5"
          v-model="data.PhoneNumber"
          pattern="\d*"
          minlength="7"
          expanded
        ></b-input>
      </b-field>
      <b-field label="微信ID">
        <b-input ref="6" v-model="data.WechatID" expanded></b-input>
      </b-field>
      <b-field label="企业号">
        <b-input ref="7" v-model="data.EnterpriseID" expanded></b-input>
      </b-field>
      <b-field label="管理员">
        <b-select ref="8" v-model="data.Admin" expanded required>
          <option :value="true">是</option>
          <option :value="false">否</option>
        </b-select>
      </b-field>
    </section>
    <footer class="modal-card-foot flex-end">
      <button class="button" type="button" @click="$parent.close()">
        取消
      </button>
      <button class="button is-primary" @click="handleClickSubmit">
        确定
      </button>
    </footer>
  </div>
</template>

<script>
export default {
  name: 'ModalUser',
  props: {
    which: {
      type: Object,
      default() {
        return {
          ID: '',
          Sub: '',
          Password: '',
          CompanyID: 0,
          Department: '',
          PhoneNumber: '',
          WechatID: '',
          EnterpriseID: '',
          Name: '',
          Admin: true,
          _: true
        }
      }
    },
    mapCompanyID2Name: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  data: () => ({ data: null }),
  created() {
    this.data = Object.assign({}, this.which)
  },
  methods: {
    handleFocusPassword() {
      if (this.data.Password === '$__hidden_PASSWORD') {
        this.data.Password = ''
      }
    },
    handleBlurPassword() {
      if (this.data.Password === '' && !this._) {
        this.data.Password = '$__hidden_PASSWORD'
        this.$nextTick(() => {
          this.$refs.password.checkHtml5Validity()
        })
      }
    },
    handleClickSubmit() {
      let valid = true
      for (const k in this.$refs) {
        // eslint-disable-next-line no-prototype-builtins
        if (!this.$refs.hasOwnProperty(k)) {
          continue
        }
        const x = this.$refs[k]
        if (!x) {
          continue
        }
        if (
          // eslint-disable-next-line no-prototype-builtins
          x.hasOwnProperty('checkHtml5Validity') &&
          typeof x.checkHtml5Validity === 'function' &&
          !x.checkHtml5Validity()
        ) {
          valid = false
        }
      }
      if (!valid) {
        return
      }
      const loading = this.$buefy.loading.open()
      let xhr
      if (this.which._) {
        xhr = this.$xhr.newUser(
          this.data.Sub,
          this.data.Password,
          this.data.Name,
          parseInt(this.data.CompanyID),
          this.data.Department,
          this.data.PhoneNumber,
          this.data.WechatID,
          this.data.EnterpriseID,
          this.data.Admin
        )
      } else {
        xhr = this.$xhr.modifyUser(
          this.data.ID,
          this.data.Sub,
          this.data.Password,
          this.data.Name,
          parseInt(this.data.CompanyID),
          this.data.Department,
          this.data.PhoneNumber,
          this.data.WechatID,
          this.data.EnterpriseID,
          this.data.Admin
        )
      }
      xhr
        .then(() => {
          this.$parent.close()
          this.$parent.$parent.loadAsyncData()
        })
        .finally(() => {
          loading.close()
        })
    }
  }
}
</script>

<style lang="scss">
.is-twitter .is-active a {
  color: #4099ff !important;
}

.readonly {
  pointer-events: none;
}

.same-width-5 li {
  width: 5em;
}
</style>
