<template>
  <div>
    <v-chart :force-fit="true" :height="height" :data="data">
      <v-tooltip />
      <v-axis />
      <v-legend />
      <v-bar position="公司名称*设备数量" color="name" :adjust="adjust" />
    </v-chart>
  </div>
</template>

<script>
const DataSet = require('@antv/data-set')

export default {
  name: 'DeviceGroupedBar',
  data() {
    return {
      height: 300,
      sourceData: [],
      adjust: [
        {
          type: 'dodge',
          marginRatio: 1 / 32
        }
      ]
    }
  },
  computed: {
    data() {
      const sourceData1 = { name: '图像类' }
      this.sourceData.forEach((x) => {
        sourceData1[x.CompanyName] = x.Count
      })
      const fields = []
      for (const k in sourceData1) {
        if (k === 'name') {
          continue
        }
        fields.push(k)
      }
      const dv = new DataSet.View().source([sourceData1])
      dv.transform({
        type: 'fold',
        fields,
        key: '公司名称',
        value: '设备数量'
      })
      return dv.rows
    }
  },
  created() {
    this.$xhr.getCompanyDevices().then((res) => {
      this.sourceData = res.data.data.result
    })
  }
}
</script>
