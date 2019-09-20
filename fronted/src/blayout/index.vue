<template>
  <el-container>
    <el-header>
      <div class="fix-header">
        <Header />
      </div>
    </el-header>

    <el-main>
      <div class="main">
        <BlogMain />
        <Aside :user="user" />
      </div>
    </el-main>
    <el-footer>
      <Footer />
    </el-footer>
  </el-container>
</template>

<script>
import { getInfo } from '@/api/user'
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
      user: {},
    }
  },
  mounted () {
    this.getUser()
  },
  methods: {
    getUser () {
      getInfo().then(response => {
        this.user = {
          nickname: response.data.nickname,
          avatar: response.data.avatar,
          introduction: response.data.introduction,
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import "~@/styles/mixin.scss";
@import "~@/styles/variables.scss";
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
  position: fixed;
  width: 100%;
  bottom: 0;
}
</style>
