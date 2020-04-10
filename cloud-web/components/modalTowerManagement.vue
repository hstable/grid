<template>
  <div class="modal-card" style="width: auto">
    <header class="modal-card-head">
      <p class="modal-card-title">塔杆管理 - {{ line.Name }}</p>
    </header>
    <section class="modal-card-body">
      <b-table
        :data="data"
        :loading="loading"
        paginated
        backend-pagination
        :total="total"
        :per-page="perPage"
        aria-next-label="Next page"
        aria-previous-label="Previous page"
        aria-page-label="Page"
        aria-current-label="Current page"
        @page-change="onPageChange"
      >
        <template slot-scope="props">
          <b-table-column field="Name" label="塔杆名称">
            {{ props.row.Name }}
          </b-table-column>

          <b-table-column field="BaseStationName" label="所属基站">
            {{ mapID2BaseStation[props.row.BaseStationID].Name }}
          </b-table-column>

          <b-table-column field="BaseStationIP" label="所属基站IP">
            {{ mapID2BaseStation[props.row.BaseStationID].IP }}
          </b-table-column>

          <b-table-column field="DevicesNum" label="设备数量" numeric>
            {{ props.row.DevicesNum }}
          </b-table-column>

          <b-table-column label="操作" width="350">
            <div class="operate-box">
              <b-button
                size="is-small"
                icon-left="map-marker"
                outlined
                type="is-second"
                @click="handleDeviceManagement(props.row)"
              >
                查看位置
              </b-button>
              <b-button
                size="is-small"
                icon-left="cctv"
                outlined
                type="is-success"
                @click="handleDeviceManagement(props.row)"
              >
                设备管理
              </b-button>
              <b-button
                size="is-small"
                icon-left="playlist-edit"
                outlined
                type="is-info"
                @click="handleModify(props.row)"
              >
                编辑
              </b-button>
              <b-button
                size="is-small"
                icon-left="delete-empty-outline"
                outlined
                type="is-delete"
                @click="handleDelete(props.row)"
              >
                删除
              </b-button>
            </div>
          </b-table-column>
        </template>
      </b-table>
    </section>
    <footer class="modal-card-foot">
      <button class="button is-primary" @click="handleNew">添加塔杆</button>
    </footer>
  </div>
</template>

<script>
import ModalTower from '@/components/modalTower'
import ModalDeviceManagement from '@/components/modalDeviceManagement'

export default {
  name: 'ModalTowerManagement',
  filters: {
    /**
     * Filter to truncate string, accepts a length parameter
     */
    truncate(value, length) {
      return value.length > length ? value.substr(0, length) + '...' : value
    }
  },
  props: {
    line: {
      type: Object,
      default: () => ({})
    }
  },
  data: () => ({
    data: [],
    total: 0,
    loading: false,
    page: 1,
    perPage: 20,
    mapID2BaseStation: {}
  }),
  mounted() {
    this.loadAsyncData()
  },
  methods: {
    /*
     * Load async data
     */
    loadAsyncData() {
      const that = this
      this.loading = true
      this.$xhr
        .getTowers(this.page, this.perPage, this.line.ID)
        .then(({ data }) => {
          this.data = []
          this.total = data.data.towers.total
          this.data = data.data.towers.data
          that.mapID2BaseStation = {}
          data.data.baseStations.forEach((c) => {
            that.mapID2BaseStation[c.ID] = c
          })
          this.loading = false
        })
        .catch((error) => {
          this.data = []
          this.total = 0
          this.loading = false
          throw error
        })
    },
    /*
     * Handle page-change event
     */
    onPageChange(page) {
      this.page = page
      this.loadAsyncData()
    },
    handleNew() {
      this.$buefy.modal.open({
        parent: this,
        component: ModalTower,
        hasModalCard: true,
        trapFocus: true,
        props: {
          line: this.line,
          mapID2BaseStation: this.mapID2BaseStation
        }
      })
    },
    handleModify(which) {
      this.$buefy.modal.open({
        parent: this,
        component: ModalTower,
        hasModalCard: true,
        trapFocus: true,
        props: {
          which,
          line: this.line,
          mapID2BaseStation: this.mapID2BaseStation
        }
      })
    },
    handleDelete(which) {
      this.$buefy.dialog.confirm({
        title: '删除确认',
        message: `删除操作是不可恢复的，请确定要删除<b>${which.Name}</b>吗？`,
        confirmText: '确定',
        cancelText: '取消',
        type: 'is-danger',
        hasIcon: true,
        onConfirm: () => {
          this.loading = true
          this.$xhr
            .delTower(which.ID)
            .then(() => {
              this.loadAsyncData()
              this.line.TowersNum--
            })
            .catch(() => {
              this.loading = false
            })
        }
      })
    },
    handleDeviceManagement(row) {
      this.$buefy.modal.open({
        parent: this,
        component: ModalDeviceManagement,
        hasModalCard: true,
        trapFocus: true,
        fullScreen: true,
        props: {
          tower: row
        }
      })
    }
  }
}
</script>

<style lang="scss">
.modal-close {
  &::before,
  &::after {
    background-color: black;
  }
}
</style>
