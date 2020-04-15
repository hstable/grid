<template>
  <div>
    <v-chart ref="chart" :force-fit="true" :height="height" :data="data">
      <v-legend :custom="true" :clickable="false" :items="items" />
      <v-axis />
      <v-tooltip />
      <v-bar
        position="PowerName*Count"
        shape="waterfall"
        :color="color"
        :tooltip="tooltip"
      />
    </v-chart>
  </div>
</template>

<script>
import { Global, registerShape } from 'viser-vue'
import deepcopy from 'deepcopy'

function getRectPath(points) {
  const path = []
  for (let i = 0; i < points.length; i++) {
    const point = points[i]
    if (point) {
      const action = i === 0 ? 'M' : 'L'
      path.push([action, point.x, point.y])
    }
  }
  const first = points[0]
  path.push(['L', first.x, first.y])
  path.push(['z'])
  return path
}

function getFillAttrs(cfg) {
  const defaultAttrs = Global.shape.interval
  const attrs = Object.assign(
    {},
    defaultAttrs,
    {
      fill: cfg.color,
      stroke: cfg.color,
      fillOpacity: cfg.opacity
    },
    cfg.style
  )
  return attrs
}

registerShape('interval', 'waterfall', {
  draw(cfg, container) {
    const attrs = getFillAttrs(cfg)
    let rectPath = getRectPath(cfg.points)
    rectPath = this.parsePath(rectPath)
    const interval = container.addShape('path', {
      attrs: Object.assign(attrs, {
        path: rectPath
      })
    })

    if (cfg.nextPoints) {
      let linkPath = [
        ['M', cfg.points[2].x, cfg.points[2].y],
        ['L', cfg.nextPoints[0].x, cfg.nextPoints[0].y]
      ]

      if (cfg.nextPoints[0].y === 0) {
        linkPath[1] = ['L', cfg.nextPoints[1].x, cfg.nextPoints[1].y]
      }
      linkPath = this.parsePath(linkPath)
      container.addShape('path', {
        attrs: {
          path: linkPath,
          stroke: '#8c8c8c',
          lineDash: [4, 2]
        }
      })
    }

    return interval
  }
})

const items = [
  { value: '塔杆数量', fill: '#1890FF', marker: 'square' },
  { value: '塔杆总数', fill: '#8c8c8c', marker: 'square' }
]

const color = [
  'PowerName',
  (PowerName) => {
    if (PowerName === '塔杆总数') {
      return '#8c8c8c'
    }
    return '#1890FF'
  }
]

const tooltip = [
  'PowerName*Count',
  (PowerName, Count) => {
    if (Array.isArray(Count)) {
      return {
        name: '数量',
        value: Count[1] - Count[0]
      }
    }

    return {
      name: '数量',
      value: Count
    }
  }
]

export default {
  name: 'PowerWaterfall',
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
      height: 300,
      items,
      color,
      tooltip
    }
  },
  computed: {
    data() {
      const data = deepcopy(this.input)
      for (let i = 0; i < data.length; i++) {
        const item = data[i]

        if (i > 0 && i < data.length - 1) {
          if (Array.isArray(data[i - 1].Count)) {
            item.Count = [
              data[i - 1].Count[1],
              item.Count + data[i - 1].Count[1]
            ]
          } else {
            item.Count = [data[i - 1].Count, item.Count + data[i - 1].Count]
          }
        }
      }
      return data
    }
  }
}
</script>
