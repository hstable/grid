import Vue from 'vue'
import Tree from './tree'
import User from './user'
import Company from './company'
import Power from './power'
import Line from './line'
import Tower from './tower'
import BaseStation from './baseStation'
import Device from './device'
import Visualization from './visualization'
import Mark from './mark'

export default function({ $axios }) {
  Vue.prototype.$xhr = {
    ...User({ $axios }),
    ...Company({ $axios }),
    ...Power({ $axios }),
    ...Tree({ $axios }),
    ...Line({ $axios }),
    ...Tower({ $axios }),
    ...BaseStation({ $axios }),
    ...Device({ $axios }),
    ...Visualization({ $axios }),
    ...Mark({ $axios })
  }
}
