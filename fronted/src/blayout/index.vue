<template>
  <el-container style="height:100%">
    <el-header>
      <div class="fix-header">
        <Header />
      </div>
    </el-header>

    <el-main>
      <div class="main">
        <BlogMain />
        <Aside :soup="soup" />
      </div>
    </el-main>
    <el-footer style="display:flex">
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
.fix-header {
  width: 80%;
  margin-left: auto;
  margin-right: auto;
}

.main {
  width: 80%;
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
  padding-left: unset;
  padding-right: unset;
}

.el-container {
  margin: auto;
}
.blog-main {
  width: 70%;
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
