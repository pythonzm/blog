<template>
  <div class="header-container">
    <div v-if="!mobile" class="logo">
      <a href="/">{{ logo }}</a>
    </div>
    <search id="header-search" />
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      active-text-color="rgb(255, 255, 255)"
      router
    >
      <el-menu-item
        v-for="(item, key) in navOptions"
        :key="key"
        :index="item.index"
        style="background-color: unset"
      >{{ item.label }}</el-menu-item>

    </el-menu>

  </div>
</template>

<script>
import Search from '@/components/HeaderSearch'
export default {
  name: 'Header',
  components: {
    Search
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
      activeIndex: '/'
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
</style>

<style scoped>
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
</style>
