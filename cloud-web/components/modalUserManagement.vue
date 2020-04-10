<template>
  <div class="modal-card" style="width: auto">
    <header class="modal-card-head">
      <p class="modal-card-title">用户管理</p>
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
          <b-table-column field="Sub" label="用户名">
            {{ props.row.Sub }}
          </b-table-column>

          <b-table-column field="Name" label="姓名">
            {{ props.row.Name }}
          </b-table-column>

          <b-table-column field="CompanyID" label="公司" numeric>
            {{ mapCompanyID2Name[props.row.CompanyID] }}
          </b-table-column>

          <b-table-column field="Department" label="部门" numeric>
            {{ props.row.Department }}
          </b-table-column>

          <b-table-column field="PhoneNumber" label="手机号" centered>
            {{ props.row.PhoneNumber }}
          </b-table-column>

          <b-table-column field="WechatID" label="微信ID">
            {{ props.row.WechatID }}
          </b-table-column>

          <b-table-column field="EnterpriseID" label="企业号">
            {{ props.row.EnterpriseID }}
          </b-table-column>
          <b-table-column label="操作" width="150">
            <div class="operate-box">
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
      <button class="button is-primary" @click="handleNew">添加用户</button>
    </footer>
  </div>
</template>

<script>
import ModalUser from '@/components/modalUser'
export default {
  name: 'ModalUserManagement',
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
    perPage: 20,
    mapCompanyID2Name: {}
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
        .getUsers(this.page, this.perPage)
        .then(({ data }) => {
          this.data = []
          this.total = data.data.users.total
          this.data = data.data.users.data
          that.mapCompanyID2Name = {}
          data.data.companies.forEach((c) => {
            that.mapCompanyID2Name[c.ID] = c.Name
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
        component: ModalUser,
        hasModalCard: true,
        trapFocus: true,
        props: {
          mapCompanyID2Name: this.mapCompanyID2Name
        }
      })
    },
    handleModify(which) {
      this.$buefy.modal.open({
        parent: this,
        component: ModalUser,
        hasModalCard: true,
        trapFocus: true,
        props: {
          which,
          mapCompanyID2Name: this.mapCompanyID2Name
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
            .delUser(which.ID)
            .then(() => {
              this.loadAsyncData()
            })
            .catch(() => {
              this.loading = false
            })
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
