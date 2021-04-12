<template>
  <div class="modal-card">
    <header class="modal-card-head">
      <p class="modal-card-title">告警清单</p>
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
          <b-table-column field="CreatedAt" label="时间">
            {{ props.row.CreatedAt | timeFormat }}
          </b-table-column>

          <b-table-column field="LineName" label="线路名称">
            {{ props.row.LineName }}
          </b-table-column>

          <b-table-column field="TowerName" label="塔杆名称">
            {{ props.row.TowerName }}
          </b-table-column>

          <b-table-column field="PowerName" label="电力级别">
            {{ props.row.PowerName }}
          </b-table-column>

          <b-table-column field="BaseStationName" label="所属边缘节点">
            {{ props.row.BaseStationName }}
          </b-table-column>

          <b-table-column field="Mark" label="告警级别">
            <b-tag :type="mark2color(props.row.Mark)">{{
              props.row.Mark | mark2name
            }}</b-tag>
          </b-table-column>

          <b-table-column field="Annotate" label="告警事由">
            {{ props.row.Annotate }}
          </b-table-column>

          <b-table-column label="操作" width="100">
            <div class="operate-box">
              <b-button
                size="is-small"
                icon-left="camera-image"
                outlined
                type="is-success"
                @click="handleOpenPreview(props.row)"
              >
                图片
              </b-button>
            </div>
          </b-table-column>
        </template>
      </b-table>
    </section>
  </div>
</template>

<script>
import dayjs from 'dayjs'
export default {
  name: 'ModalLogMarksRisky',
  filters: {
    mark2name(mark) {
      return {
        danger: '非常危险',
        warning: '危险',
        info: '警惕',
        safe: '安全'
      }[mark]
    },
    timeFormat(time) {
      return dayjs(time).format('YYYY-MM-DD HH:mm:ss')
    }
  },
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
  data: () => ({
    loading: false,
    total: 0,
    page: 1,
    perPage: 20,
    data: []
  }),
  created() {
    this.loadAsyncData()
  },
  methods: {
    mark2color(mark) {
      return {
        danger: 'is-red',
        warning: 'is-yellow',
        info: 'is-twitter',
        safe: 'is-green'
      }[mark]
    },
    handleOpenPreview(row) {
      this.$buefy.modal.open(
        `<p class="image is-4by3">
              <img src="${apiRoot}/edgeServer/${row.BaseStationID}/image/${row.Filename}">
          </p>`
      )
    },
    onPageChange(page) {
      this.page = page
      this.loadAsyncData()
    },
    loadAsyncData() {
      this.loading = true
      this.$xhr
        .getMarksRisky(this.page, this.perPage)
        .then(({ data }) => {
          this.data = []
          this.total = data.data.marks.total
          this.data = data.data.marks.data.map((x) => {
            x.CreatedTime = dayjs(x.CreatedAt).format('YYYY-MM-DD HH:mm:ss')
            return x
          })
        })
        .catch((error) => {
          this.data = []
          this.total = 0
          throw error
        })
        .finally(() => {
          this.loading = false
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
