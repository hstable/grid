<template>
  <div class="modal-card" style="max-width: 520px;margin:auto">
    <header class="modal-card-head">
      <p class="modal-card-title">
        {{ data._ ? '添加线路' : '修改线路信息' }}
      </p>
    </header>
    <section class="modal-card-body">
      <b-field label="线路名称">
        <b-input ref="2" v-model="data.Name" expanded required></b-input>
      </b-field>
      <b-field label="所属公司">
        <b-select ref="3" v-model="data.CompanyID" expanded required>
          <option v-for="(name, id) of mapCompanyID2Name" :key="id" :value="id"
            >{{ name }}
          </option>
        </b-select>
      </b-field>
      <b-field label="电力级别">
        <b-select ref="4" v-model="data.PowerID" expanded required>
          <option v-for="(name, id) of mapPowerID2Name" :key="id" :value="id"
            >{{ name }}
          </option>
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
          Name: '',
          CompanyID: 0,
          PowerID: 0,
          _: true
        }
      }
    },
    mapCompanyID2Name: {
      type: Object,
      default() {
        return {}
      }
    },
    mapPowerID2Name: {
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
        xhr = this.$xhr.newLine(
          this.data.Name,
          parseInt(this.data.CompanyID),
          parseInt(this.data.PowerID)
        )
      } else {
        xhr = this.$xhr.modifyLine(
          this.data.ID,
          this.data.Name,
          parseInt(this.data.CompanyID),
          parseInt(this.data.PowerID)
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
