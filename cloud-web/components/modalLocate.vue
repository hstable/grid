<template>
  <div style="height:80vh;width:100%">
    <div id="locate-container" style="height:100%;width:100%"></div>
    <div class="locate-info">
      <h4 id="locate-status">{{ status }}</h4>
      <div v-show="showResult">
        <hr />
        <!--eslint-disable-next-line-->
        <p id="locate-result" v-html="result"></p>
        <hr />
      </div>
      <div v-if="showButtons">
        <button
          class="button is-primary is-small"
          @click="$emit('complete', position)"
        >
          定位准确，使用该定位
        </button>
        <button class="button is-small" @click="$parent.close()">取消</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ModalLocate',
  data: () => ({
    position: null,
    showButtons: false,
    showResult: false,
    result: '',
    status: '定位中...'
  }),
  mounted() {
    this.initMap()
  },
  methods: {
    initMap() {
      const that = this
      const map = new AMap.Map('locate-container', {
        resizeEnable: true
      })
      AMap.plugin('AMap.Geolocation', function() {
        const geolocation = new AMap.Geolocation({
          enableHighAccuracy: true, // 是否使用高精度定位，默认:true
          timeout: 10000, // 超过10秒后停止定位，默认：5s
          buttonPosition: 'RB', // 定位按钮的停靠位置
          buttonOffset: new AMap.Pixel(10, 20), // 定位按钮与设置的停靠位置的偏移量，默认：Pixel(10, 20)
          zoomToAccuracy: true // 定位成功后是否自动调整地图视野到定位点
        })
        map.addControl(geolocation)
        geolocation.getCurrentPosition(function(status, result) {
          that.showResult = true
          if (status === 'complete') {
            onComplete(result)
          } else {
            onError(result)
          }
        })
      })
      // 解析定位结果
      function onComplete(data) {
        that.status = '定位成功'
        const str = []
        str.push('定位结果：' + data.position)
        that.position = data.position
        str.push('定位类别：' + data.location_type)
        if (data.accuracy) {
          str.push('精度：' + data.accuracy + ' 米')
        } // 如为IP精确定位结果则没有精度信息
        that.result = str.join('<br>')
        that.showButtons = true
      }
      // 解析定位错误信息
      function onError(data) {
        that.status = '定位失败'
        that.result = '失败原因排查信息:' + data.message
      }
    }
  }
}
</script>

<style lang="scss">
.locate-info {
  & {
    box-sizing: border-box;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto,
      'Helvetica Neue', Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji',
      'Segoe UI Symbol', 'Noto Color Emoji';
    line-height: 1.5;
    font-weight: 300;
    color: #111213;
  }
  font-size: 12px;
  max-width: 98vw;
  padding: 0.75rem 1.25rem;
  margin-bottom: 1rem;
  border-radius: 0.25rem;
  position: fixed;
  top: 1vw;
  right: 1vw;
  background-color: white;
  width: 32em;
  border-width: 0;
  box-shadow: 0 2px 6px 0 rgba(114, 124, 245, 0.5);
  hr {
    margin: 0.5rem 0;
    box-sizing: content-box;
    height: 0;
    overflow: visible;
    border: 0;
    border-top: 1px solid rgba(0, 0, 0, 0.1);
  }
  p {
    margin-top: 0;
    margin-bottom: 0;
  }
}
</style>
