<template>
  <div class="modal-card" style="max-width: 520px;margin:auto">
    <header class="modal-card-head">
      <p class="modal-card-title">
        {{ data._ ? '添加公司' : '修改公司信息' }}
      </p>
    </header>
    <section class="modal-card-body">
      <b-field label="公司名称">
        <b-input ref="1" v-model="data.Name" expanded required></b-input>
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
  name: 'ModalCompany',
  props: {
    which: {
      type: Object,
      default() {
        return {
          ID: '',
          Name: '',
          _: true
        }
      }
    }
  },
  data: () => ({ data: null }),
  created() {
    this.data = Object.assign({}, this.which)
  },
  methods: {
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
        xhr = this.$xhr.newCompany(this.data.Name)
      } else {
        xhr = this.$xhr.modifyCompany(this.data.ID, this.data.Name)
      }
      xhr
        .then(() => {
          this.$parent.$parent.loadAsyncData()
          this.$parent.close()
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
