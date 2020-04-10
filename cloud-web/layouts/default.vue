<template>
  <div>
    <b-navbar ref="navs" fixed-top shadow type="is-light">
      <template slot="brand">
        <b-navbar-item href="/" style="padding: 0">
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
            <span>基站管理</span>
          </b-navbar-item>
          <b-navbar-item @click="handleClickPowerManagement">
            <b-icon icon="flash" size="is-small"></b-icon>
            <span>电力管理</span>
          </b-navbar-item>
          <b-navbar-item @click="handleClickLineManagement">
            <b-icon icon="help-network" size="is-small"></b-icon>
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

    <section class="main-content columns">
      <b-menu class="column is-2 section">
        <b-menu-list label="MENU">
          <b-menu-item
            icon="home"
            label="系统总览"
            :to="{ name: 'index' }"
            tag="nuxt-link"
            exact-active-class="is-exact"
          ></b-menu-item>
          <b-menu-item icon="eye">
            <template slot="label" slot-scope="props">
              通道可视化
              <b-icon
                class="is-pulled-right"
                :icon="props.expanded ? 'menu-down' : 'menu-up'"
              ></b-icon>
            </template>
            <b-menu-item
              v-for="company of list"
              :key="company.id"
              icon="domain"
              :label="company.name"
            ></b-menu-item>
          </b-menu-item>
        </b-menu-list>
      </b-menu>

      <div class="container column is-10">
        <nuxt />
      </div>
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

export default {
  data() {
    return {
      sub: localStorage.sub,
      name: localStorage.name,
      list: []
    }
  },
  created() {
    this.$xhr.getLines().then((res) => {
      // const data = res.data.data
    })
  },
  methods: {
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
.menu-list .is-exact {
  background-color: #2263ab;
  color: #fff;
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
