<template>
  <div class="modal-card" style="max-width: 520px;margin:auto">
    <header class="modal-card-head">
      <p class="modal-card-title">
        {{ data._ ? '添加塔杆' : '修改塔杆信息' }} - {{ line.Name }}
      </p>
    </header>
    <section class="modal-card-body">
      <b-field label="塔杆名称">
        <b-input
          ref="towerName"
          v-model="data.Name"
          expanded
          required
        ></b-input>
      </b-field>
      <b-field label="所属边缘节点">
        <b-select
          ref="baseStation"
          v-model="data.BaseStationID"
          expanded
          required
        >
          <option v-for="(bs, id) of mapID2BaseStation" :key="id" :value="id"
            >{{ bs.Name }}({{ bs.IP }})
          </option>
        </b-select>
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
            <button class="button is-primary" @click="locate">定位</button>
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

    <b-modal :active.sync="showLocate">
      <modal-locate
        v-if="showLocate"
        @complete="handleLocateComplete"
      ></modal-locate>
    </b-modal>
  </div>
</template>

<script>
import ModalLocate from './modalLocate'
export default {
  name: 'ModalUser',
  components: { ModalLocate },
  props: {
    which: {
      type: Object,
      default() {
        return {
          ID: '',
          Name: '',
          BaseStationID: 0,
          Location: '',
          _: true
        }
      }
    },
    mapID2BaseStation: {
      type: Object,
      default() {
        return {}
      }
    },
    line: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  data: () => ({ data: null, showLocate: false }),
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
        xhr = this.$xhr
          .newTower(
            this.data.Name,
            this.line.ID,
            parseInt(this.data.BaseStationID),
            this.data.Location.replace(/\s/, '')
          )
          .then(() => {
            this.line.TowersNum++
          })
      } else {
        xhr = this.$xhr.modifyTower(
          this.data.ID,
          this.data.Name,
          this.line.ID,
          parseInt(this.data.BaseStationID),
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
    },
    locate() {
      this.showLocate = true
    },
    handleLocateComplete(location) {
      this.data.Location = location.Q + ',' + location.R
      this.showLocate = false
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
