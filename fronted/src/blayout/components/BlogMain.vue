<template>
  <section class="blog-main">
    <transition name="fade-transform" mode="out-in">
      <keep-alive>
        <router-view :key="key" />
      </keep-alive>
    </transition>
  </section>
</template>

<script>
import { truncate } from 'fs';
export default {
  name: 'BlogMain',
  computed: {
    key () {
      return this.$route.fullPath
    }
  }
}
</script>

<style scoped>
.fixed-header + .blog-main {
  padding-top: 50px;
}
</style>

<style lang="scss">
// fix css style bug in open el-dialog
.el-popup-parent--hidden {
  .fixed-header {
    padding-right: 15px;
  }
}
</style>
