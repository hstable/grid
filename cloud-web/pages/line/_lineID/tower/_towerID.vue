<template>
  <div class="columns" style="padding:2.25em 0 0.75em">
    <b-carousel
      class="flex-none"
      style="width:calc(100% - 340px - 1.5em)"
      :autoplay="false"
      with-carousel-list
      :indicator="false"
      :overlay="gallery"
      animated="fade"
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
          v-model="props.active"
          :data="items"
          :config="al"
          :refresh="gallery"
          as-indicator
          @switch="props.switch($event, false)"
        />
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
export default {
  name: 'ID',
  data() {
    return {
      events: [
        {
          date: new Date(2020, 4 - 1, 2),
          type: 'is-green'
        },
        {
          date: new Date(2020, 4 - 1, 6),
          type: 'is-red'
        },
        {
          date: new Date(2020, 4 - 1, 6),
          type: 'is-yellow'
        },
        {
          date: new Date(2020, 4 - 1, 6),
          type: 'is-green'
        },
        {
          date: new Date(2020, 4 - 1, 13),
          type: 'is-yellow'
        },
        {
          date: new Date(2020, 4 - 1, 15),
          type: 'is-green'
        }
      ],
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
      items: [
        {
          title: 'Slide 1',
          image:
            'http://sx.zhysdxl.com:8001/upload/99000843096846/202004/s/99000843096846_20200412171903.jpg'
        },
        {
          title: 'Slide 2',
          image:
            'http://sx.zhysdxl.com:8001/upload/99000843096846/202004/s/99000843096846_20200412161904.jpg'
        },
        {
          title: 'Slide 3',
          image:
            'http://sx.zhysdxl.com:8001/upload/99000843096846/202004/s/99000843096846_20200412151905.jpg'
        },
        {
          title: 'Slide 4',
          image:
            'http://sx.zhysdxl.com:8001/upload/99000843096846/202004/s/99000843096846_20200412141904.jpg'
        },
        {
          title: 'Slide 5',
          image:
            'http://sx.zhysdxl.com:8001/upload/99000843096846/202004/s/99000843096846_20200412171903.jpg'
        },
        {
          title: 'Slide 6',
          image:
            'http://sx.zhysdxl.com:8001/upload/99000843096846/202004/s/99000843096846_20200412121904.jpg'
        },
        {
          title: 'Slide 7',
          image:
            'http://sx.zhysdxl.com:8001/upload/99000843096846/202004/s/99000843096846_20200412111904.jpg'
        },
        {
          title: 'Slide 8',
          image:
            'http://sx.zhysdxl.com:8001/upload/99000843096846/202004/s/99000843096846_20200412101904.jpg'
        }
      ]
    }
  },
  created() {
    this.$xhr.getTowerMonthMarks(3, new Date())
  },
  methods: {
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
</style>
