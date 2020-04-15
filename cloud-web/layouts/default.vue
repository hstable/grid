<template>
  <div>
    <b-navbar ref="navs" fixed-top shadow type="is-light">
      <template slot="brand">
        <b-navbar-item tag="nuxt-link" to="/" style="padding: 0">
          <img src="~assets/logo.png" alt="陕西电网" class="logo no-select" />
        </b-navbar-item>
      </template>
      <template slot="end">
        <b-navbar-dropdown hoverable arrowless>
          <template slot="label">
            <b-icon icon="cog-outline" size="is-small"></b-icon>
            <span>系统配置</span>
          </template>
          <b-navbar-item @click="handleClickUserManagement">
            <b-icon icon="account" size="is-small"></b-icon>
            <span>用户管理</span>
          </b-navbar-item>
          <b-navbar-item @click="handleClickCompanyManagement">
            <b-icon icon="domain" size="is-small"></b-icon>
            <span>公司管理</span>
          </b-navbar-item>
          <b-navbar-item @click="handleClickBaseStationManagement">
            <b-icon icon="transmission-tower" size="is-small"></b-icon>
            <span>边缘节点管理</span>
          </b-navbar-item>
          <b-navbar-item @click="handleClickPowerManagement">
            <b-icon icon="flash" size="is-small"></b-icon>
            <span>电力级别管理</span>
          </b-navbar-item>
          <b-navbar-item @click="handleClickLineManagement">
            <b-icon icon="transit-connection-variant" size="is-small"></b-icon>
            <span>线路管理</span>
          </b-navbar-item>
        </b-navbar-dropdown>
        <b-dropdown
          position="is-bottom-left"
          aria-role="menu"
          style="margin-right:10px"
          class="menudropdown"
          hoverable
        >
          <a slot="trigger" class="navbar-item" role="button">
            <span class="no-select">{{ name }}</span>
            <b-icon icon="menu-down"></b-icon>
          </a>

          <b-dropdown-item custom aria-role="menuitem">
            当前登录用户为 {{ sub }}
          </b-dropdown-item>
          <hr class="dropdown-divider" />
          <b-dropdown-item
            value="logout"
            aria-role="menuitem"
            class="no-select"
            @click="handleClickLogout"
          >
            <b-icon icon="logout" size="is-small"></b-icon>
            注销
          </b-dropdown-item>
        </b-dropdown>
      </template>
    </b-navbar>

    <section class="main-content columns" style="margin-bottom: 0">
      <b-menu ref="menu" class="column section">
        <b-menu-list label="MENU">
          <c-menu-item
            icon="home"
            label="系统总览"
            :to="{ name: 'index' }"
            tag="nuxt-link"
            exact-active-class="is-exact"
          ></c-menu-item>
          <c-menu-item icon="eye" tag="menuitem" :expanded.sync="tree.expanded">
            <template slot="label" slot-scope="props">
              通道可视化
              <b-icon
                class="is-pulled-right"
                :icon="props.expanded ? 'menu-down' : 'menu-up'"
              ></b-icon>
            </template>
            <c-menu-item
              v-for="company of tree.root"
              :key="company.CompanyID"
              :active.sync="company.active"
              :expanded.sync="company.expanded"
              icon="domain"
              tag="menuitem"
              :disabled="company._Disabled"
            >
              <template slot="label" slot-scope="props">
                {{ tree.companyMap[company.CompanyID] }}
                <b-icon
                  class="is-pulled-right"
                  :icon="props.expanded ? 'menu-down' : 'menu-up'"
                ></b-icon>
              </template>
              <c-menu-item
                v-for="power of company.data"
                :key="power.PowerID"
                :active.sync="power.active"
                :expanded.sync="power.expanded"
                icon="flash"
                tag="menuitem"
              >
                <template slot="label" slot-scope="props">
                  {{ tree.powerMap[power.PowerID] }}
                  <b-icon
                    class="is-pulled-right"
                    :icon="props.expanded ? 'menu-down' : 'menu-up'"
                  ></b-icon>
                </template>
                <c-menu-item
                  v-for="line of power.data"
                  :key="line.ID"
                  :active.sync="line.active"
                  :expanded.sync="line.expanded"
                  icon="transit-connection-variant"
                  tag="menuitem"
                >
                  <template slot="label" slot-scope="props">
                    {{ line.Name }}
                    <b-icon
                      class="is-pulled-right"
                      :icon="props.expanded ? 'menu-down' : 'menu-up'"
                    ></b-icon>
                  </template>
                  <c-menu-item
                    v-for="tower of line.data"
                    :key="tower.ID"
                    icon="transmission-tower"
                    :label="tower.Name"
                    :disabled="tower._Disabled"
                    tag="nuxt-link"
                    exact-active-class="is-exact"
                    :to="`/line/${line.ID}/tower/${tower.ID}`"
                  >
                  </c-menu-item>
                </c-menu-item>
              </c-menu-item>
            </c-menu-item>
          </c-menu-item>
        </b-menu-list>
      </b-menu>

      <nuxt class="column is-10" />
    </section>
    <div id="login"></div>
  </div>
</template>

<script>
import ModalUserManagement from '@/components/modalUserManagement'
import ModalCompanyManagement from '@/components/modalCompanyManagement'
import ModalPowerManagement from '@/components/modalPowerManagement'
import ModalLineManagement from '@/components/modalLineManagement'
import ModalBaseStationManagement from '@/components/modalBaseStationManagement'
import CMenuItem from '@/components/buefy/MenuItem'

export default {
  components: { CMenuItem },
  data() {
    return {
      sub: localStorage.sub,
      name: localStorage.name,
      tree: {}
    }
  },
  created() {
    this.setupTree()
  },
  methods: {
    inactiveAll() {
      const menu = this.$refs.menu
      function f(component) {
        for (const x of component.$children) {
          if (typeof x.newActive === 'boolean') {
            if (x.newActive || x.newExpanded) {
              x.newActive = false
              x.newExpanded = false
            }
            f(x)
          }
        }
      }
      f(menu)
    },
    handleLineActiveChanged(line) {
      if (line.active) {
        if (!line.data.length) {
          line.data = [
            {
              Name: 'loading...',
              _Disabled: true
            }
          ]
        }
        this.$xhr
          .getTreeTowers(line.ID)
          .then((res) => {
            line.data = res.data.data.towers
          })
          .catch(() => {
            line.data = [
              {
                Name: '加载失败',
                _Disabled: true
              }
            ]
          })
      }
    },
    setupTree() {
      const that = this
      this.tree = {
        root: [{ CompanyID: -1, _Disabled: true }],
        companyMap: { '-1': 'loading...' }
      }
      this.$xhr
        .getTreeLines()
        .then((res) => {
          const data = res.data.data.lines
          const tree = []
          const companyMap = {}
          const powerMap = {}
          data.forEach((x) => {
            // 建立ID2Company映射
            if (!(x.CompanyID in companyMap)) {
              companyMap[x.CompanyID] = x.CompanyName
            }
            // 建立ID2Power映射
            if (!(x.PowerID in powerMap)) {
              powerMap[x.PowerID] = x.PowerName
            }
          })
          // 建立第一层tree并按CompanyID递增排序
          for (const cid in companyMap) {
            tree.push({
              CompanyID: parseInt(cid),
              data: [],
              ID2Index: {},
              expanded: false,
              active: false
            })
          }
          tree.sort((a, b) => {
            if (a.CompanyID < b.CompanyID) {
              return -1
            } else if (a.CompanyID === b.CompanyID) {
              return 0
            }
            return 1
          })
          const ID2Index = {}
          for (const index in tree) {
            ID2Index[tree[index].CompanyID] = index
          }

          // 建立第n层tree并按PowerID递增排序
          function setupTreeLayer(
            data,
            layer,
            sortKeyword,
            locateVal2Index,
            locateKeyword,
            reservedProperties
          ) {
            if (!reservedProperties) {
              reservedProperties = []
            }
            data.forEach((x) => {
              if (
                !(
                  x[sortKeyword] in
                  layer[locateVal2Index[x[locateKeyword]]].ID2Index
                )
              ) {
                layer[locateVal2Index[x[locateKeyword]]].ID2Index[
                  x[sortKeyword]
                ] = -1
                const obj = {
                  [sortKeyword]: x[sortKeyword],
                  data: [],
                  ID2Index: {},
                  expanded: false,
                  active: false
                }
                for (const p of reservedProperties) {
                  obj[p] = x[p]
                }
                layer[locateVal2Index[x[locateKeyword]]].data.push(obj)
              }
            })
            layer.forEach((x) => {
              x.data.sort((a, b) => {
                if (a[sortKeyword] < b[sortKeyword]) {
                  return -1
                } else if (a[sortKeyword] === b[sortKeyword]) {
                  return 0
                }
                return 1
              })
              for (const index in x.data) {
                x.ID2Index[x.data[index][sortKeyword]] = parseInt(index)
              }
            })
          }

          setupTreeLayer(data, tree, 'PowerID', ID2Index, 'CompanyID')
          tree.forEach((c) => {
            setupTreeLayer(
              data.filter((d) => d.CompanyID === c.CompanyID),
              c.data,
              'ID',
              c.ID2Index,
              'PowerID',
              ['Name', 'CompanyID', 'CompanyName', 'PowerID', 'PowerName']
            )
            c.data.forEach((power) => {
              power.data.forEach((line) => {
                let active = false
                Object.defineProperty(line, 'active', {
                  enumerable: true,
                  configurable: true,
                  set(newValue) {
                    active = newValue
                    that.handleLineActiveChanged(line)
                    return active
                  },
                  get() {
                    return active
                  }
                })
              })
            })
          })
          this.tree = {
            root: tree,
            companyMap,
            powerMap,
            expanded: false
          }
          this.expandMenu()
        })
        .catch(() => {
          this.tree = {
            root: [{ CompanyID: -1, _Disabled: true }],
            companyMap: { '-1': '加载失败' }
          }
        })
    },
    expandMenu() {
      // eslint-disable-next-line no-unused-vars
      let { lineID } = this.$route.params
      if (!lineID) {
        return
      }
      lineID = parseInt(lineID)
      const tree = this.tree
      if (
        tree.root.some((com) => {
          if (
            com.data.some((power) => {
              if (
                power.data.some((line) => {
                  if (line.ID === lineID) {
                    line.expanded = true
                    line.active = true
                    return true
                  }
                  return false
                })
              ) {
                power.expanded = true
                power.active = true
                return true
              }
              return false
            })
          ) {
            com.expanded = true
            com.active = true
            return true
          }
          return false
        })
      ) {
        tree.expanded = true
      }
    },
    handleClickLogout() {
      this.$buefy.loading.open()
      localStorage.removeItem('token')
      localStorage.removeItem('name')
      localStorage.removeItem('sub')
      window.location.reload()
    },
    handleClickUserManagement() {
      this.$buefy.modal.open({
        parent: this,
        component: ModalUserManagement,
        hasModalCard: true,
        trapFocus: true,
        fullScreen: true
      })
    },
    handleClickCompanyManagement() {
      this.$buefy.modal.open({
        parent: this,
        component: ModalCompanyManagement,
        hasModalCard: true,
        trapFocus: true,
        fullScreen: true
      })
    },
    handleClickPowerManagement() {
      this.$buefy.modal.open({
        parent: this,
        component: ModalPowerManagement,
        hasModalCard: true,
        trapFocus: true,
        fullScreen: true
      })
    },
    handleClickLineManagement() {
      this.$buefy.modal.open({
        parent: this,
        component: ModalLineManagement,
        hasModalCard: true,
        trapFocus: true,
        fullScreen: true
      })
    },
    handleClickBaseStationManagement() {
      this.$buefy.modal.open({
        parent: this,
        component: ModalBaseStationManagement,
        hasModalCard: true,
        trapFocus: true,
        fullScreen: true
      })
    }
  }
}
</script>
<style lang="scss" scoped>
#app {
  margin: 0;
}

.menucontainer {
  padding: 20px;
}

.logo {
  max-height: 3.25em;
  margin-left: 1em;
  margin-right: 1em;
}

.navbar-item .iconfont {
  margin-right: 0.15em;
}

#statusTag:hover {
  cursor: pointer;
}
</style>

<style lang="scss">
.navbar-dropdown {
  border-top: none;
}
.carousel.is-overlay {
  width: 100% !important;
  padding: 0;

  .modal-close {
    &:before,
    &:after {
      background-color: white;
      filter: drop-shadow(0px 0px 3px #000);
    }
  }
}

menuitem {
  display: block;
  cursor: pointer;
  user-select: none;

  border-radius: 2px;
  color: #4a4a4a;
  padding: 0.5em 0.75em;

  &:hover {
    background-color: rgba(37, 104, 187, 0.15);
  }
}

.menu-list li ul {
  margin-right: 0 !important;
}

.menu-list a:hover,
.menu-list .is-exact {
  background-color: rgba(37, 104, 187, 0.15);
}

html {
  //  &::-webkit-scrollbar {
  //    // 去掉讨厌的滚动条
  //    display: none;
  //  }

  #app {
    height: calc(100vh - 3.25rem);
    /*overflow-y: auto;*/
    //overflow-scrolling: touch;
    //-webkit-overflow-scrolling: touch;
  }
}

@media screen and (max-width: 1023px) {
  .dropdown.is-mobile-modal .dropdown-menu {
    // 修复modal模糊问题
    left: 0 !important;
    right: 0 !important;
    margin: auto;
    transform: unset !important;
    z-index: 1000;
  }
}

.dropdown-item:focus {
  // 不要丑丑的outline
  outline: none !important;
}

.no-select {
  user-select: none;
  -webkit-user-drag: none;
}

.menudropdown .dropdown-item {
  font-size: 0.8em;
}

.leave-right {
  margin-right: 0.6rem;
}

.lazy {
  p {
    margin: 0.5em 0;
  }
}

a.navbar-item:focus,
a.navbar-item:focus-within,
a.navbar-item:hover,
a.navbar-item.is-active,
.navbar-link:focus,
.navbar-link:focus-within,
.navbar-link:hover,
.navbar-link.is-active,
.is-link,
a {
  $success: #506da4;
  color: $success;
}

.icon-loading_ico-copy {
  font-size: 2.5rem;
  color: rgba(0, 0, 0, 0.45);
  animation: loading-rotate 2s infinite linear;
}

.modal-custom-ports {
  z-index: 999;
}

.after-line-dot5 {
  p {
    margin-bottom: 0.5em;
  }
}

.about-small {
  font-size: 0.85em;
  text-indent: 1em;
  color: rgba(0, 0, 0, 0.6);
}

a.dropdown-item.is-active,
.dropdown .dropdown-menu .has-link a.is-active,
button.dropdown-item.is-active {
  background-color: unset;
}
</style>
