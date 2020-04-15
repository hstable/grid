<template>
  <div>
    <v-chart :force-fit="true" :height="height" :data="data" :scale="scale">
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

const sourceData = [
  { value: 149, type: '灾害', name: '烟' },
  { value: 5, type: '灾害', name: '火灾' },
  { value: 2517, type: '施工', name: '铲车' },
  { value: 2260, type: '施工', name: '工程车辆' },
  { value: 5765, type: '施工', name: '起重机' },
  { value: 3479, type: '施工', name: '吊车' }
]

const dv = new DataSet.View().source(sourceData)
dv.transform({
  type: 'percent',
  field: 'value',
  dimension: 'type',
  as: 'percent'
})
const data = dv.rows

const viewDv = new DataSet.View().source(sourceData)
viewDv.transform({
  type: 'percent',
  field: 'value',
  dimension: 'name',
  as: 'percent'
})
const viewData = viewDv.rows

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
      data,
      scale,
      viewData,
      height: 300,
      itemTpl,
      tooltip,
      color,
      label,
      style
    }
  }
}
</script>
