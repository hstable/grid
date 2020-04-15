<template>
  <div>
    <v-chart :force-fit="true" :height="height" :data="data" :scale="scale">
      <v-tooltip :show-title="false" data-key="item*percent" />
      <v-axis />
      <v-legend data-key="item" />
      <v-pie
        position="percent"
        color="item"
        :v-style="pieStyle"
        :label="labelConfig"
      />
      <v-coord type="theta" :radius="0.75" :inner-radius="0.6" />
    </v-chart>
  </div>
</template>

<script>
const DataSet = require('@antv/data-set')

const interpreter = {
  danger: '非常危险',
  warning: '危险',
  info: '警惕',
  safe: '安全'
}

const scale = [
  {
    dataKey: 'percent',
    min: 0,
    formatter: '.0%'
  }
]

export default {
  name: 'MarkPie',
  data() {
    return {
      scale,
      height: 300,
      sourceData: [],
      pieStyle: {
        stroke: '#fff',
        lineWidth: 1
      },
      labelConfig: [
        'percent',
        {
          formatter: (val, item) => {
            return item.point.item + ': ' + val
          }
        }
      ]
    }
  },
  computed: {
    data() {
      const dv = new DataSet.View().source(
        this.sourceData.map((x) => {
          return { item: interpreter[x.Mark], count: x.Count }
        })
      )
      dv.transform({
        type: 'percent',
        field: 'count',
        dimension: 'item',
        as: 'percent'
      })
      return dv.rows
    }
  },
  created() {
    this.$xhr.getMarks().then((res) => {
      this.sourceData = res.data.data.result
    })
  }
}
</script>
