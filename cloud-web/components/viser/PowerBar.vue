<template>
  <div>
    <v-chart :force-fit="true" :height="height" :data="data" :scale="scale">
      <v-tooltip />
      <v-axis />
      <v-bar position="PowerName*Count" :tooltip="tooltip" />
    </v-chart>
  </div>
</template>

<script>
import deepcopy from 'deepcopy'
const tooltip = [
  'PowerName*Count',
  (PowerName, Count) => ({
    name: '数量',
    value: Count
  })
]

const scale = [
  {
    dataKey: 'Count',
    tickInterval: 20
  }
]

export default {
  name: 'PowerBar',
  props: {
    input: {
      type: Array,
      default() {
        return []
      }
    }
  },
  data() {
    return {
      scale,
      tooltip,
      height: 300
    }
  },
  computed: {
    data() {
      return deepcopy(this.input)
    }
  }
}
</script>
