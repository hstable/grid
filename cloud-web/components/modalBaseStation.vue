<template>
  <div class="modal-card" style="max-width: 520px;margin:auto">
    <header class="modal-card-head">
      <p class="modal-card-title">
        {{ data._ ? '添加边缘节点' : '修改边缘节点信息' }}
      </p>
    </header>
    <section class="modal-card-body">
      <b-field label="边缘节点名称">
        <b-input ref="name" v-model="data.Name" expanded required></b-input>
      </b-field>
      <b-field label="IPv4地址">
        <b-input
          ref="ipadress"
          v-model="data.IP"
          pattern="\d+\.\d+\.\d+\.\d+"
          expanded
          required
        ></b-input>
      </b-field>
      <b-field label="地理位置">
        <b-field>
          <b-input
            ref="location"
            v-model="data.Location"
            pattern="\d+\.\d+,\s*\d+\.\d+"
            expanded
            required
          ></b-input>
          <p class="control">
            <button class="button is-primary">定位</button>
          </p>
        </b-field>
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
          Online: true,
          IP: '',
          Location: '',
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
        xhr = this.$xhr.newBaseStation(
          this.data.Name,
          this.data.IP,
          this.data.Location.replace(/\s/, '')
        )
      } else {
        xhr = this.$xhr.modifyBaseStation(
          this.data.ID,
          this.data.Name,
          this.data.IP,
          this.data.Location.replace(/\s/, '')
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
.field > .field.has-addons {
  position: relative;
  .help {
    position: absolute;
    bottom: -1.5em;
  }
}
</style>
