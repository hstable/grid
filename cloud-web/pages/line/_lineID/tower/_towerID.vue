<template>
  <div class="columns" style="padding:2.25em 0 0.75em">
    <b-carousel
      :key="items.length"
      class="flex-none"
      style="width:calc(100% - 340px - 1.5em)"
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
                borderColor: props2.list.borderColor
              }"
            >
              <img :src="props2.list.image" :title="props2.list.title" />
            </figure>
          </template>
        </b-carousel-list>
      </template>
      <template slot="overlay">
        <div class="has-text-centered has-text-white">
          Hello i'am overlay!
        </div>
      </template>
    </b-carousel>
    <div class="flex-none" style="width:350px;padding:0 0.75em">
      <b-datepicker v-model="date" inline :events="events" indicators="dots" />
      <b-collapse style="margin-top: 0.75em" class="panel" animation="slide">
        <div slot="trigger" class="panel-heading" role="button">
          <strong>设备操作</strong>
        </div>
        <p class="panel-tabs">
          <a
            v-for="(c, i) of cameras"
            :key="c.ID"
            :class="{ 'is-active': indexCamera === i }"
            @click="indexCamera = i"
            >{{ c.Name }}</a
          >
        </p>
        <div class="panel-block" style="padding: 0">
          <div style="width:100%" class="card-footer">
            <b-button
              type="is-primary"
              style="height:100%;border-radius:0;border-color: whitesmoke"
              class="card-footer-item"
              outlined
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
      events: [],
      date: null,
      indexCamera: 0,
      cameras: [
        { ID: 0, Name: '摄像头1' },
        { ID: 1, Name: '摄像头2' },
        { ID: 2, Name: '摄像头3' }
      ],
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
      loading: false
    }
  },
  computed: {
    itemsToShow() {
      return this.$refs.indicator.settings.itemsToShow
    }
  },
  created() {
    this.initDatePicker()
  },
  mounted() {
    this.loadCarouselData()
  },
  methods: {
    handleIndexChange(val) {
      if (this.items.length - (val + this.itemsToShow) < 2) {
        const after =
          this.items.length > 0 ? this.items[this.items.length - 1].ID : null
        this.loadCarouselData(after)
      }
    },
    loadCarouselData(after) {
      if (this.end || this.loading) {
        return
      }
      this.loading = true
      this.$nextTick(() => {
        this.$xhr
          .getTreeTowerImages(
            after,
            this.itemsToShow + 2,
            this.$route.params.towerID
          )
          .then((res) => {
            const interpreterColor = {
              danger: 'rgba(255, 0, 0, 0.73)',
              warning: 'rgb(255, 200, 0)',
              info: 'rgb(80, 109, 164)'
            }
            const data = res.data.data
            if (data.images) {
              this.items.push(
                ...data.images.map((x) => ({
                  ID: x.ID,
                  title: x.Name,
                  image: `${apiRoot}/image/${x.Filename}`,
                  borderColor: interpreterColor[x.Mark]
                }))
              )
            } else {
              if (!this.items.length) {
                this.items = [
                  {
                    title: '暂无图片',
                    image: 'http://sx.zhysdxl.com:8001/images/noImage.jpg'
                  }
                ]
              }
              this.end = true
            }
          })
          .finally(() => {
            this.loading = false
          })
      })
    },
    initDatePicker() {
      this.$xhr
        .getTowerMonthMarks(this.$route.params.towerID, new Date())
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
        })
    },
    switchGallery(value) {
      this.gallery = value
      if (value) {
        return document.documentElement.classList.add('is-clipped')
      } else {
        return document.documentElement.classList.remove('is-clipped')
      }
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
</style>
