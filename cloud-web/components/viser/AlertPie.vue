<template>
  <div>
    <v-chart
      v-if="ok"
      :force-fit="true"
      :height="height"
      :data="data"
      :scale="scale"
    >
      <v-tooltip :show-title="false" :item-tpl="itemTpl" />
      <v-coord type="theta" :radius="0.5" />
      <v-pie
        position="percent"
        :color="'type'"
        :label="label"
        :select="false"
        :v-style="style"
        :tooltip="tooltip"
      />
      <v-view :data="viewData" :scale="scale">
        <v-coord type="theta" :radius="0.75" :inner-radius="0.5 / 0.75" />
        <v-pie
          position="percent"
          :color="color"
          label="name"
          :select="false"
          :v-style="style"
          :tooltip="tooltip"
        />
      </v-view>
    </v-chart>
  </div>
</template>

<script>
const DataSet = require('@antv/data-set')

const scale = {
  dataType: 'percent',
  formatter: '.2%'
}

const itemTpl =
  '<li><span style="background-color:{color};" class="g2-tooltip-marker"></span>{name}: {value}</li>'

const style = {
  lineWidth: 1,
  stroke: '#fff'
}

const label = ['type', { offset: -10 }]

const tooltip = [
  'name*percent',
  (item, percent) => {
    percent = (percent * 100).toFixed(2) + '%'
    return {
      name: item,
      value: percent
    }
  }
]

const color = [
  'name',
  ['#BAE7FF', '#7FC9FE', '#71E3E3', '#ABF5F5', '#8EE0A1', '#BAF5C4']
]

export default {
  name: 'AlertPie',
  data() {
    return {
      scale,
      height: 300,
      itemTpl,
      tooltip,
      color,
      label,
      style,
      ok: false
    }
  },
  computed: {
    data() {
      const dv = new DataSet.View().source(this.sourceData)
      dv.transform({
        type: 'percent',
        field: 'value',
        dimension: 'type',
        as: 'percent'
      })
      return dv.rows
    },
    viewData() {
      const viewDv = new DataSet.View().source(this.sourceData)
      viewDv.transform({
        type: 'percent',
        field: 'value',
        dimension: 'name',
        as: 'percent'
      })
      return viewDv.rows
    }
  },
  created() {
    this.$xhr.getTypeCounts().then((res) => {
      this.sourceData = res.data.data.result
      this.ok = true
    })
  }
}
</script>
