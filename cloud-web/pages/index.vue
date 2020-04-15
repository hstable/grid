<template>
  <section style="padding:1.5em 0 0;position: relative;left:-1.25em">
    <div class="tile is-ancestor" style="margin-bottom: 0">
      <div class="tile is-vertical">
        <AlertPie></AlertPie>
        <PowerBar :input="PowerCounts"></PowerBar>
      </div>
      <div class="tile is-7 is-vertical is-parent">
        <div id="map-container" style="height:100%;width:100%"></div>
      </div>
      <div class="tile is-vertical">
        <MarkPie></MarkPie>
        <div class="tile is-child" style="position: relative">
          <b-table :data="data" :columns="columns"></b-table>
          <p
            v-show="!data.length"
            style="position: absolute;height:100%;width:100%;text-align: center"
          >
            暂无数据
          </p>
        </div>
      </div>
    </div>
    <div class="columns">
      <div class="column" style="padding: 0">
        <PowerWaterfall :input="PowerCounts"></PowerWaterfall>
      </div>
      <div class="column is-7" style="padding: 0">
        <BlueBlockGraph></BlueBlockGraph>
      </div>
      <div
        class="column"
        style="padding-left: 0;padding-top: 0;padding-bottom: 0"
      >
        <DeviceGroupedBar></DeviceGroupedBar>
      </div>
    </div>
  </section>
</template>

<script>
import PowerWaterfall from '../components/viser/PowerWaterfall'
import AlertPie from '../components/viser/AlertPie'
import DeviceGroupedBar from '../components/viser/DeviceGroupedBar'
import PowerBar from '../components/viser/PowerBar'
import MarkPie from '../components/viser/MarkPie'
import BlueBlockGraph from '@/components/viser/BlueBlockGraph'
export default {
  name: 'HomePage',
  components: {
    MarkPie,
    PowerBar,
    DeviceGroupedBar,
    AlertPie,
    PowerWaterfall,
    BlueBlockGraph
  },
  data: () => ({
    PowerCounts: [],
    data: [],
    columns: [
      {
        field: 'Name',
        label: '塔杆名称',
        centered: true
      },
      {
        field: 'CreatedTime',
        label: '巡视时间',
        centered: true
      },
      {
        field: 'Annotate',
        label: '告警事由',
        centered: true
      }
    ]
  }),
  created() {
    this.$parent.$parent.inactiveAll()
  },
  mounted() {
    this.initMap()
    this.loadPowerCounts()
  },
  methods: {
    loadPowerCounts() {
      this.$xhr.getPowerCounts().then((res) => {
        const data = res.data.data
        const sum = {
          PowerName: '塔杆总数',
          Count: 0
        }
        data.result.forEach((x) => {
          sum.Count += x.Count
        })
        data.result.push(sum)
        this.PowerCounts = data.result
      })
    },
    initMap() {
      const map = new AMap.Map('map-container', {
        mapStyle: 'amap://styles/4261c8dec313b784c0933324313428a8666'
      })
      map.setZoomAndCenter(8, [108.939621, 34.343147])
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
          return new AMap.Marker({
            position: new AMap.LngLat(lng, lat),
            title: '边缘节点: ' + x.Name,
            icon: 'http://sx.zhysdxl.com:8001/images/v3.0/tower1.png'
          })
        })
        const towers = data.towers.map((x) => {
          const [lng, lat] = x.Location.split(',')
          return new AMap.Marker({
            position: new AMap.LngLat(lng, lat),
            title: x.Name,
            icon: 'http://sx.zhysdxl.com:8001/images/v3.0/tower0.png'
          })
        })
        // 将创建的点标记添加到已有的地图实例：
        map.add([...towers, ...baseStations])
      })
    }
  }
}
</script>

<style lang="scss">
.amap-logo {
  z-index: inherit !important;
}
</style>
