<template>
  <div class="modal-card" style="width: auto">
    <header class="modal-card-head">
      <p class="modal-card-title">基站管理</p>
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
          <b-table-column field="Name" label="基站名称">
            {{ props.row.Name }}
          </b-table-column>

          <b-table-column field="IPAddress" label="IP地址">
            {{ props.row.IP }}
          </b-table-column>

          <b-table-column field="Online" label="在线状态">
            {{ props.row.Online ? '正常' : '无法正常通信' }}
          </b-table-column>

          <b-table-column field="TowersNum" label="塔杆数量" numeric>
            {{ props.row.TowersNum }}
          </b-table-column>

          <b-table-column label="操作" width="250">
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
      <button class="button is-primary" @click="handleNew">添加基站</button>
    </footer>
  </div>
</template>

<script>
import ModalBaseStation from '@/components/modalBaseStation'

export default {
  name: 'ModalBaseStationManagement',
  filters: {
    /**
     * Filter to truncate string, accepts a length parameter
     */
    truncate(value, length) {
      return value.length > length ? value.substr(0, length) + '...' : value
    }
  },
  data: () => ({
    data: [],
    total: 0,
    loading: false,
    page: 1,
    perPage: 20
  }),
  mounted() {
    this.loadAsyncData()
  },
  methods: {
    /*
     * Load async data
     */
    loadAsyncData() {
      this.loading = true
      this.$xhr
        .getBaseStations(this.page, this.perPage)
        .then(({ data }) => {
          this.data = []
          this.total = data.data.baseStations.total
          this.data = data.data.baseStations.data
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
        component: ModalBaseStation,
        hasModalCard: true,
        trapFocus: true
      })
    },
    handleModify(which) {
      this.$buefy.modal.open({
        parent: this,
        component: ModalBaseStation,
        hasModalCard: true,
        trapFocus: true,
        props: {
          which
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
            .delBaseStation(which.ID)
            .then(() => {
              this.loadAsyncData()
            })
            .catch(() => {
              this.loading = false
            })
        }
      })
    },
    handleDeviceManagement() {
      // TODO
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
