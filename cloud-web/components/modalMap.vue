<template>
  <div id="location-container" style="height:80vh;width:100%"></div>
</template>

<script>
import ModalLogEdgeServer from '@/components/modalLogEdgeServer'
export default {
  name: 'ModalMap',
  props: {
    location: { type: Array, default: () => [108.939621, 34.343147] }
  },
  mounted() {
    this.initMap()
  },
  methods: {
    initMap() {
      const that = this
      const map = new AMap.Map('location-container', {
        mapStyle: 'amap://styles/4261c8dec313b784c0933324313428a8666'
      })
      map.setZoomAndCenter(12, [this.location[1], this.location[0]])
      /* 绘制小地图的各个省份的分界线 */
      AMapUI.loadUI(['geo/DistrictExplorer'], function(DistrictExplorer) {
        // 创建一个实例
        const districtExplorer = new DistrictExplorer({
          map
        })
        const adcode = 100000
        districtExplorer.loadAreaNode(adcode, function(_, areaNode) {
          // 清除已有的绘制内容
          districtExplorer.clearFeaturePolygons()
          // 绘制子区域
          districtExplorer.renderSubFeatures(areaNode, function(feature, i) {
            return {
              bubble: true,
              strokeColor: '#000', // 线颜色
              strokeOpacity: 0.5, // 线透明度
              strokeWeight: 0.6, // 线宽
              fillColor: '#EEEEEE', // 填充色
              fillOpacity: 0 // 填充透明度
            }
          })
        })
      })

      this.$xhr.getLocations().then((res) => {
        const data = res.data.data
        const baseStations = data.baseStations.map((x) => {
          const [lng, lat] = x.Location.split(',')
          const m = new AMap.Marker({
            position: new AMap.LngLat(lat, lng),
            title: '边缘节点: ' + x.Name,
            x,
            icon: '/tower1.png',
            offset: new AMap.Pixel(-8, -15)
          })
          m.on(
            'click',
            (e) => {
              const obj = e.target.De
              this.$buefy.modal.open({
                component: ModalLogEdgeServer,
                hasModalCard: true,
                props: {
                  which: obj.x
                },
                fullScreen: true
              })
            },
            that
          )
          return m
        })
        const towers = data.towers.map((x) => {
          const [lng, lat] = x.Location.split(',')
          return new AMap.Marker({
            position: new AMap.LngLat(lat, lng),
            title: x.Name,
            icon: '/tower0.png',
            offset: new AMap.Pixel(-8, -15)
          })
        })
        // 将创建的点标记添加到已有的地图实例：
        map.add([...towers, ...baseStations])
        // 创建关系
        const mapIDBaseStationPosition = {}
        for (const b of data.baseStations) {
          const [lng, lat] = b.Location.split(',')
          mapIDBaseStationPosition[b.ID] = new AMap.LngLat(lat, lng)
        }
        for (const t of data.towers) {
          const [lng, lat] = t.Location.split(',')
          map.add(
            new AMap.Polyline({
              path: [
                mapIDBaseStationPosition[t.BaseStationID],
                new AMap.LngLat(lat, lng)
              ],
              borderWeight: 1, // 线条宽度，默认为 1
              strokeColor: '#4099FF' // 线条颜色
            })
          )
        }
        // 创建标记
        map.add(
          new AMap.Marker({
            position: new AMap.LngLat(this.location[1], this.location[0])
          })
        )
      })
    }
  }
}
</script>

<style lang="scss"></style>
