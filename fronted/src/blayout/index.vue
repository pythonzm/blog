<template>
  <el-container>
    <el-header>
      <div class="fix-header">
        <Header />
      </div>
    </el-header>

    <el-main>
      <a
        href="https://github.com/pythonzm"
        target="_blank"
        style="position: fixed;"
        ><img
          width="149"
          height="149"
          src="https://github.blog/wp-content/uploads/2008/12/forkme_right_darkblue_121621.png?resize=149%2C149"
          class="attachment-full size-full"
          alt="Fork me on GitHub"
          data-recalc-dims="1"
      /></a>
      <div class="main">
        <BlogMain />
        <Aside :soup="soup" />
      </div>
    </el-main>
    <el-footer>
      <Footer />
    </el-footer>
  </el-container>
</template>

<script>
import { fetchRandSoup } from '@/api/soup'

import Header from './components/header'
import BlogMain from './components/BlogMain'
import Aside from './components/aside'
import Footer from './components/footer'
export default {
  name: 'Blayout',
  components: {
    Header,
    BlogMain,
    Aside,
    Footer
  },
  data () {
    return {
      soup: {},
    }
  },
  mounted () {
    this.getSoup()
  },
  methods: {
    getSoup () {
      fetchRandSoup().then(response => {
        this.soup = response.data
      })
    }
  }
}
</script>

<style lang="scss" scoped>
a {
  position: fixed;
  right: 0;
  z-index: 1;
  margin-top: -20px;
}
.fix-header {
  margin-left: auto;
  margin-right: auto;
}

.main {
  margin-left: auto;
  margin-right: auto;
}
.el-header {
  text-align: center;
  line-height: 60px;
  padding: unset;
  background-color: #000000;
  position: fixed;
  top: 0;
  right: 0;
  z-index: 9;
  width: 100%;
}

.el-main {
  // background-color: lightgreen;
  text-align: center;
  padding-top: 80px;
}

.el-container {
  margin: auto;
  min-height: 100vh;
}
.blog-main {
  width: 75%;
  float: left;
}
.el-aside {
  width: 30%;
}
.el-footer {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #000000;
}
</style>
