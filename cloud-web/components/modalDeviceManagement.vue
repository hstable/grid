<template>
  <div class="modal-card" style="width: auto">
    <header class="modal-card-head">
      <p class="modal-card-title">设备管理 - {{ tower.Name }}</p>
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
          <b-table-column field="Name" label="设备名称">
            {{ props.row.Name }}
          </b-table-column>

          <b-table-column field="SerialNumber" label="设备号">
            {{ props.row.SerialNumber }}
          </b-table-column>

          <b-table-column field="Manufacture" label="厂商">
            {{ props.row.Manufacture }}
          </b-table-column>

          <b-table-column field="ModelName" label="型号">
            {{ props.row.ModelName }}
          </b-table-column>

          <b-table-column
            field="Disabled"
            label="状态"
            @click="handleStatusControl(props.row)"
          >
            {{ props.row.Disabled ? '禁用' : '正常' }}
          </b-table-column>

          <b-table-column label="操作" width="250">
            <div class="operate-box">
              <b-button
                size="is-small"
                :icon-left="props.row.Disabled ? 'play' : 'stop'"
                outlined
                :type="props.row.Disabled ? 'is-success' : 'is-danger'"
                @click="handleStatusControl(props.row)"
              >
                {{ props.row.Disabled ? '启用' : '禁用' }}
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
      <button class="button is-primary" @click="handleNew">添加设备</button>
    </footer>
  </div>
</template>

<script>
import ModalDevice from '@/components/modalDevice'

export default {
  name: 'ModalDeviceManagement',
  filters: {
    /**
     * Filter to truncate string, accepts a length parameter
     */
    truncate(value, length) {
      return value.length > length ? value.substr(0, length) + '...' : value
    }
  },
  props: {
    tower: {
      type: Object,
      default: () => ({})
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
        .getDevices(this.page, this.perPage, this.tower.ID)
        .then(({ data }) => {
          this.data = []
          this.total = data.data.devices.total
          this.data = data.data.devices.data
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
        component: ModalDevice,
        hasModalCard: true,
        trapFocus: true,
        props: {
          tower: this.tower
        }
      })
    },
    handleModify(which) {
      this.$buefy.modal.open({
        parent: this,
        component: ModalDevice,
        hasModalCard: true,
        trapFocus: true,
        props: {
          which,
          tower: this.tower
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
            .delDevice(which.ID)
            .then(() => {
              this.loadAsyncData()
              this.tower.DevicesNum--
            })
            .catch(() => {
              this.loading = false
            })
        }
      })
    },
    handleStatusControl(row) {
      this.loading = true
      this.$xhr
        .changeDeviceStatus(row.ID, row.Disabled)
        .then(() => {
          this.loadAsyncData()
        })
        .catch(() => {
          this.loading = false
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
