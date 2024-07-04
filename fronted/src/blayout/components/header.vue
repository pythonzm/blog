<template>
  <div class="header-container">
    <div v-if="!mobile" class="logo">
      <a href="/">{{ logo }}</a>
    </div>
    <el-menu
      :default-active="activeIndex"
      mode="horizontal"
      active-text-color="rgb(255, 255, 255)"
      router
      style="display: inline-flex;"
    >
      <el-menu-item
        v-for="(item, key) in navOptions"
        :key="key"
        :index="item.index"
        style="background-color: unset;"
      >{{ item.label }}</el-menu-item>
      <div style="padding: 0 10px;">
        <svg-icon icon-class="search" class="search" @click="dialogVisible = true" />
      </div>

    </el-menu>

    <el-dialog
      :visible.sync="dialogVisible"
      append-to-body
      :show-close="showClose"
      :fullscreen="mobile"
    >
      <AlgoliaSearch class="el-dialog-div" />
    </el-dialog>
  </div>
</template>

<script>
// import Search from '@/components/HeaderSearch'
import AlgoliaSearch from '@/components/AlgoliaSearch'
export default {
  name: 'Header',
  components: {
    AlgoliaSearch
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
      showClose: false
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
.search {
  color: gray;
  cursor: pointer;
}

.search:hover {
  color: #fff;
}

.logo {
  padding: 0 20px;
  font-size: 1.4em;
  color: #fff;
  text-decoration: none;
  float: left;
}
.header-search {
  float: right;
  padding: 0 20px;
}
.header-container {
  text-align: center;
}
.el-menu.el-menu--horizontal {
  border-bottom: unset;
  float: right;
}
.el-menu {
  background-color: unset;
}
.el-menu--horizontal > .el-menu-item:hover {
  color: #fff;
}
.el-dialog-div {
  max-height: 70vh;
  overflow: auto;
}
</style>
