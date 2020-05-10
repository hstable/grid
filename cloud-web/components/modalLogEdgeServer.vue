<template>
  <div class="modal-card">
    <header class="modal-card-head">
      <p class="modal-card-title">边缘服务器日志 - {{ which.Name }}</p>
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

          <b-table-column field="TowerName" label="塔杆名称">
            {{ props.row.TowerName }}
          </b-table-column>

          <b-table-column field="Mark" label="结果">
            <b-tag :type="mark2color(props.row.Mark)">{{
              props.row.Mark | mark2name
            }}</b-tag>
          </b-table-column>

          <b-table-column field="Annotate" label="备注">
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
  name: 'ModalLogEdgeServer',
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
    data: [
      {
        CreatedAt: dayjs('2020-04-18 08:00:00').toDate(),
        TowerID: 1,
        TowerName: '鱼镐Ⅰ、Ⅱ线/005号大号侧',
        Mark: 'safe',
        Annotate: '',
        Filename: '1.jpg'
      },
      {
        CreatedAt: dayjs('2020-04-18 07:59:00').toDate(),
        TowerID: 1,
        TowerName: '信上I II线/224号（小号侧）',
        Mark: 'safe',
        Annotate: '',
        Filename: '2.jpg'
      },
      {
        CreatedAt: dayjs('2020-04-18 07:58:00').toDate(),
        TowerID: 1,
        TowerName: '鱼镐Ⅰ、Ⅱ线/005号大号侧',
        Mark: 'danger',
        Annotate: '塔吊',
        Filename: '1.jpg'
      },
      {
        CreatedAt: dayjs('2020-04-18 07:57:00').toDate(),
        TowerID: 1,
        TowerName: '鱼镐Ⅰ、Ⅱ线/005号大号侧',
        Mark: 'warning',
        Annotate: '起重机',
        Filename: '1.jpg'
      },
      {
        CreatedAt: dayjs('2020-04-18 07:56:00').toDate(),
        TowerID: 1,
        TowerName: '信上I II线/214号',
        Mark: 'info',
        Annotate: '塔吊',
        Filename: '3.jpg'
      },
      {
        CreatedAt: dayjs('2020-04-18 07:55:00').toDate(),
        TowerID: 1,
        TowerName: '鱼镐Ⅰ、Ⅱ线/005号大号侧',
        Mark: 'safe',
        Annotate: '',
        Filename: '1.jpg'
      },
      {
        CreatedAt: dayjs('2020-04-18 07:54:00').toDate(),
        TowerID: 1,
        TowerName: '鱼镐Ⅰ、Ⅱ线/005号大号侧',
        Mark: 'safe',
        Annotate: '',
        Filename: '1.jpg'
      },
      {
        CreatedAt: dayjs('2020-04-18 07:53:00').toDate(),
        TowerID: 1,
        TowerName: '鱼镐Ⅰ、Ⅱ线/005号大号侧',
        Mark: 'safe',
        Annotate: '',
        Filename: '1.jpg'
      },
      {
        CreatedAt: dayjs('2020-04-18 07:52:00').toDate(),
        TowerID: 1,
        TowerName: '鱼镐Ⅰ、Ⅱ线/005号大号侧',
        Mark: 'safe',
        Annotate: '',
        Filename: '1.jpg'
      },
      {
        CreatedAt: dayjs('2020-04-18 07:51:00').toDate(),
        TowerID: 1,
        TowerName: '鱼镐Ⅰ、Ⅱ线/005号大号侧',
        Mark: 'safe',
        Annotate: '',
        Filename: '1.jpg'
      },
      {
        CreatedAt: dayjs('2020-04-18 07:50:00').toDate(),
        TowerID: 1,
        TowerName: '鱼镐Ⅰ、Ⅱ线/005号大号侧',
        Mark: 'safe',
        Annotate: '',
        Filename: '1.jpg'
      }
    ]
  }),
  created() {
    // TODO
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
              <img src="${apiRoot}/edgeServer/${row.TowerID}/image/${row.Filename}">
          </p>`
      )
    },
    onPageChange(page) {
      this.page = page
      this.loadAsyncData()
    },
    loadAsyncData() {
      // const that = this
      // this.loading = true
      // this.$xhr
      //   .getLines(this.page, this.perPage)
      //   .then(({ data }) => {
      //     this.data = []
      //     this.total = data.data.lines.total
      //     this.data = data.data.lines.data
      //     that.mapCompanyID2Name = {}
      //     data.data.companies.forEach((c) => {
      //       that.mapCompanyID2Name[c.ID] = c.Name
      //     })
      //     that.mapPowerID2Name = {}
      //     data.data.powers.forEach((c) => {
      //       that.mapPowerID2Name[c.ID] = c.Name
      //     })
      //     this.loading = false
      //   })
      //   .catch((error) => {
      //     this.data = []
      //     this.total = 0
      //     this.loading = false
      //     throw error
      //   })
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
