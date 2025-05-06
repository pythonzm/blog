<template>
  <div class="header-container">
    <div v-if="!mobile" class="logo">
      <a href="/">{{ logo }}</a>
    </div>

    <div class="nav-wrapper">
      <el-menu
        :default-active="activeIndex"
        mode="horizontal"
        active-text-color="#fff"
        router
        class="nav-menu"
      >
        <el-menu-item
          v-for="(item, key) in navOptions"
          :key="key"
          :index="item.index"
          style="background-color: unset;"
        >{{ item.label }}</el-menu-item>

        <el-menu-item class="search-item">
          <svg-icon
            v-if="algoliaSearch"
            class-name="search-icon"
            icon-class="search"
            @click="dialogVisible = true"
          />

          <HeaderSearch v-else id="header-search" />
        </el-menu-item>
      </el-menu>
    </div>

    <el-dialog
      :visible.sync="dialogVisible"
      append-to-body
      :show-close="showClose"
      :fullscreen="mobile"
      @opened="focusAlgolia"
      @closed="blurAlgolia"
    >
      <AlgoliaSearch class="el-dialog-div" />
    </el-dialog>
  </div>
</template>

<script>
import HeaderSearch from '@/components/HeaderSearch'
import AlgoliaSearch from '@/components/AlgoliaSearch'
import defaultSettings from '@/settings'

const { algoliaSearch } = defaultSettings
export default {
  name: 'Header',
  components: {
    AlgoliaSearch,
    HeaderSearch
  },
  props: {
    mobile: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      logo: 'PoorOPS',
      navOptions: [
        { label: '首页', index: '/' },
        { label: '分类', index: '/category' },
        { label: '标签', index: '/tag' },
        { label: '藏宝阁', index: '/collection' },
        { label: '关于', index: '/about' }
      ],
      activeIndex: '/',
      dialogVisible: false,
      showClose: false,
      algoliaSearch
    }
  },
  created() {
    // init the default  selected path
    const path = this.$route.path
    const names = ['category', 'tag', 'collection', 'about']
    for (const name of names) {
      if (path.indexOf(name) !== -1) {
        this.activeIndex = `/${name}`
        return
      } else {
        this.activeIndex = '/'
      }
    }
  },
  methods: {
    focusAlgolia() {
      document.getElementsByClassName('ais-SearchBox-input')[0].focus()
    },
    blurAlgolia() {
      document.getElementsByClassName('ais-SearchBox-input')[0].blur()
    }
  }
}
</script>

<style>
.header-search-select .el-input__inner {
  font-size: 13px;
  height: 25px;
  line-height: 25px;
}
.el-dialog__header {
  display: none;
}
</style>

<style scoped>
.search-icon {
  font-size: 18px;
  color: #909399;
}

.search-icon:hover {
  color: #fff;
}

.logo {
  padding: 0 20px;
  font-size: 1.4em;
  color: #fff;
  text-decoration: none;
  flex-shrink: 0;
}
.header-search {
  flex-shrink: 0;
}
.header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  overflow-x: auto;
  box-sizing: border-box;
}
.el-menu.el-menu--horizontal {
  display: flex !important;
  flex-wrap: nowrap !important;
  flex: 1; 
  justify-content: space-between;
  align-items: center;
  background-color: unset;
  border-bottom: unset;
}
.el-menu.el-menu--horizontal>.el-menu-item:not(.is-disabled):hover {
  background-color: unset;
}
.el-menu {
  background-color: unset;
  overflow-x: auto;
  -webkit-overflow-scrolling: touch;
}
.el-menu-item {
  white-space: nowrap;
  flex-shrink: 0;
  padding: 0 15px;
}

.el-dialog-div {
  max-height: 70vh;
  overflow: auto;
}
.nav-wrapper {
  display: flex;
  align-items: center;
  flex-shrink: 0;
}

.nav-menu .el-menu-item {
  white-space: nowrap;
  flex-shrink: 0;
}

.search-item {
  cursor: pointer;
}

@media (max-width: 480px) {
  .header-container .el-menu-item {
    font-size: 14px;
  }
  .nav-wrapper {
    width: 100%;
  }
}
</style>
