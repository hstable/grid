<template>
  <div class="columns" style="padding:2.25em 0 0.75em">
    <div style="position: relative;width:calc(100% - 340px - 1.5em)">
      <b-carousel
        v-if="items.length"
        v-model="carouselValue"
        class="flex-none"
        :autoplay="false"
        with-carousel-list
        :indicator="false"
        :overlay="gallery"
        animated="fade"
        :repeat="false"
        @change="handleIndexChange"
      >
        <b-carousel-item v-for="(item, i) in items" :key="i">
          <figure class="image" @click="switchGallery(true)">
            <img :src="item.image" />
          </figure>
        </b-carousel-item>
        <span
          v-if="gallery"
          class="modal-close is-large"
          @click="switchGallery(false)"
        />
        <template slot="list" slot-scope="props">
          <b-carousel-list
            ref="indicator"
            v-model="props.active"
            :data="items"
            :config="al"
            :refresh="gallery"
            as-indicator
            @switch="props.switch($event, false)"
          >
            <template slot="item" slot-scope="props2">
              <figure
                class="image"
                :style="{
                  borderStyle: props2.list.borderColor ? 'solid' : 'none',
                  borderWidth: '3px',
                  borderColor: props2.list.borderColor,
                  height: '100%'
                }"
              >
                <img
                  :src="props2.list.image"
                  :title="props2.list.title"
                  style="height: 100%"
                />
              </figure>
            </template>
          </b-carousel-list>
        </template>
      </b-carousel>
      <b-loading
        :is-full-page="false"
        :active="loading && !loadmore"
        :can-cancel="false"
      ></b-loading>
    </div>
    <div class="flex-none" style="width:350px;padding:0 0.75em">
      <div>
        <b-datepicker
          v-model="date"
          class="towerID-date-picker"
          inline
          :events="events"
          indicators="dots"
          :selectable-dates="selectableDates"
          :unselectable-days-of-week="unselectableDaysOfWeek"
          @change-month="handleChangeMonth"
          @change-year="handleChangeYear"
        >
          <b-loading
            :is-full-page="false"
            :active.sync="loadingDatePicker"
            :can-cancel="false"
          ></b-loading
        ></b-datepicker>
      </div>
      <b-collapse
        style="margin-top: 0.75em;position: relative"
        class="panel"
        animation="slide"
      >
        <b-loading
          :is-full-page="false"
          :active.sync="loadingDeviceOperations"
          :can-cancel="false"
        ></b-loading>
        <div slot="trigger" class="panel-heading" role="button">
          <strong>设备操作</strong>
        </div>
        <p class="panel-tabs">
          <a
            v-for="(c, i) of devices"
            :key="c.ID"
            :class="{ 'is-active': indexDevice === i }"
            @click="indexDevice = i"
            >{{ c.Name }}</a
          >
        </p>
        <div v-if="devices.length > 0">
          <div class="panel-block" style="padding: 0">
            <div style="width:100%" class="card-footer">
              <b-button
                type="is-primary"
                style="height:100%;border-radius:0;border-color: whitesmoke"
                class="card-footer-item"
                outlined
                @click="newPhoto"
                >图像抓拍
              </b-button>
              <b-button
                type="is-light"
                style="height:100%;border-radius:0;border-color: whitesmoke;color:gray"
                class="card-footer-item"
                outlined
                >参数设置
              </b-button>
              <!--            <a class="card-footer-item">参数设置</a>-->
            </div>
          </div>
          <div class="panel-block" style="padding: 0">
            <div style="width:100%" class="card-footer">
              <b-button
                type="is-second"
                style="height:100%;border-radius:0;border-color: whitesmoke"
                class="card-footer-item"
                outlined
                >视频拍摄
              </b-button>

              <b-button
                type="is-danger"
                style="height:100%;border-radius:0;border-color: whitesmoke"
                class="card-footer-item"
                outlined
                >禁用
              </b-button>
            </div>
          </div>
        </div>
        <p v-else class="panel-block">
          <span v-if="loadingDeviceOperations">加载中...</span>
          <span v-else>未在该塔杆下检测到设备</span>
        </p>
      </b-collapse>
    </div>
  </div>
</template>

<script>
import dayjs from 'dayjs'

export default {
  name: 'ID',
  data() {
    return {
      carouselValue: 0,
      events: [],
      date: null,
      indexDevice: 0,
      devices: [],
      gallery: false,
      al: {
        hasGrayscale: true,
        itemsToShow: 2,
        breakpoints: {
          768: {
            hasGrayscale: false,
            itemsToShow: 4
          },
          960: {
            hasGrayscale: true,
            itemsToShow: 6
          }
        }
      },
      items: [],
      end: false,
      loading: false,
      loadmore: false,
      loadingDatePicker: false,
      loadingDeviceOperations: false,
      selectableDates: [],
      year: dayjs().year()
    }
  },
  computed: {
    itemsToShow() {
      return this.$refs.indicator
        ? this.$refs.indicator.settings.itemsToShow
        : 6
    },
    unselectableDaysOfWeek() {
      if (this.selectableDates.length) {
        return []
      } else {
        return [0, 1, 2, 3, 4, 5, 6]
      }
    }
  },
  watch: {
    date(val) {
      this.beginTime = val
      this.endTime = dayjs(this.beginTime)
        .add(1, 'day')
        .toDate()
      this.end = false
      this.loading = false
      this.loadCarouselData()
    }
  },
  mounted() {
    this.loadCarouselData()
    this.loadDatePicker()
    this.initDevicesList()
  },
  methods: {
    handleChangeYear(YY) {
      this.year = YY
    },
    handleChangeMonth(MM) {
      this.$nextTick(() => {
        this.loadDatePicker(dayjs(`${this.year}-${MM + 1}-01`).toDate())
      })
    },
    initDevicesList() {
      this.loadingDeviceOperations = true
      return this.$xhr
        .getDevices(1, 20, this.$route.params.towerID)
        .then((res) => {
          this.devices = res.data.data.devices.data
        })
        .finally(() => {
          this.loadingDeviceOperations = false
        })
    },
    handleIndexChange(val) {
      if (this.items.length - (val + this.itemsToShow) < 2) {
        const after =
          this.items.length > 0 ? this.items[this.items.length - 1].ID : null
        this.loadCarouselData(after, true)
      }
    },
    loadCarouselData(after, loadmore) {
      if (this.end || this.loading) {
        return
      }
      this.loading = true
      loadmore && (this.loadmore = true)
      this.$nextTick(() => {
        return this.$xhr
          .getTreeTowerImages(
            after,
            this.itemsToShow + 2,
            this.$route.params.towerID,
            this.beginTime,
            this.endTime
          )
          .then((res) => {
            const interpreterColor = {
              danger: 'rgba(255, 0, 0, 0.73)',
              warning: 'rgb(255, 200, 0)',
              info: 'rgb(80, 109, 164)',
              safe: 'transparent'
            }
            const data = res.data.data
            data.images == null && (data.images = [])
            if (loadmore) {
              this.items == null && (this.items = [])
            } else {
              this.end = false
              this.items = []
            }
            if (data.images.length) {
              this.items.push(
                ...data.images
                  .filter(
                    (x) =>
                      !this.items.length ||
                      x.ID < this.items[this.items.length - 1].ID
                  )
                  .map((x) => ({
                    ID: x.ID,
                    title: x.Name,
                    image: `${apiRoot}/image/${x.Filename}`,
                    borderColor: interpreterColor[x.Mark]
                  }))
              )
              this.$refs.indicator.total += data.images.length
            } else {
              this.end = true
            }
            if (!this.items.length) {
              this.items = [
                {
                  title: '暂无图片',
                  image: '/noImage.jpg'
                }
              ]
              this.end = true
            }
          })
          .finally(() => {
            this.loading = false
            this.loadmore = false
          })
      })
    },
    loadDatePicker(date) {
      this.loadingDatePicker = true
      !date && (date = new Date())
      return this.$xhr
        .getTowerMonthMarks(this.$route.params.towerID, date)
        .then((res) => {
          const interpreter = {
            danger: 'is-red',
            warning: 'is-yellow',
            info: 'is-green'
          }
          const data = res.data.data
          this.events = data.marks.map((x) => ({
            date: dayjs(x.CreatedAt).toDate(),
            type: interpreter[x.Mark]
          }))
          this.selectableDates = this.events.map((x) => x.date)
        })
        .finally(() => {
          this.loadingDatePicker = false
        })
    },
    switchGallery(value) {
      this.gallery = value
      if (value) {
        return document.documentElement.classList.add('is-clipped')
      } else {
        return document.documentElement.classList.remove('is-clipped')
      }
    },
    newPhoto() {
      const loading = this.$buefy.loading.open()
      const deviceID = this.devices[this.indexDevice].ID
      this.$xhr
        .newPhoto(deviceID)
        .then((res) => {
          this.beginTime = null
          this.endTime = null
          this.loadCarouselData()
        })
        .finally(() => {
          loading.close()
        })
    }
  }
}
</script>

<style lang="scss">
.flex-none {
  flex: none;
}

.carousel-item {
  &.fade-leave-to,
  &.fade-leave-from,
  &.fade-leave-active,
  &.fade-enter-to,
  &.fade-enter-from,
  &.fade-enter-active {
    opacity: unset !important;
  }
}

@media screen and (min-width: 768px) {
  .carousel-items img {
    max-height: calc((100vh - 4.75em - 5px) / 5 * 4);
  }
  .carousel-list img {
    max-height: calc((100vh - 4.75em - 5px) / 5);
  }
}

@media screen and (min-width: 960px) {
  .carousel-items img {
    max-height: calc((100vh - 4.75em - 5px) / 7 * 6);
  }
  .carousel-list img {
    max-height: calc((100vh - 4.75em - 5px) / 7);
  }
}

.datepicker
  .datepicker-table
  .datepicker-body
  .datepicker-cell.is-selectable:focus:not(.is-selected) {
  background-color: transparent;
}

.datepicker
  .datepicker-table
  .datepicker-body.has-events
  .datepicker-cell.has-event.dots {
  .events {
    padding: 0;
    margin: 0 0.1em;
    overflow: hidden;
    width: calc(100% - 0.2em);
  }

  .event {
    flex-shrink: 0;
  }
}
.towerID-date-picker .datepicker-footer {
  padding: 0;
  border: none;
  margin: 0;
}
</style>
