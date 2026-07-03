<template>
  <el-container class="blog-container">
    <el-header height="64px">
      <div class="fix-header">
        <Header :mobile="device === 'mobile'" />
      </div>
    </el-header>

    <el-main class="blog-main-wrapper">
      <div class="main-layout">
        <BlogMain class="main-content" />
        <div v-if="device !== 'mobile'" class="main-aside">
          <Aside :soup="soup" />
        </div>
      </div>
    </el-main>

    <el-footer height="60px">
      <Footer />
    </el-footer>
  </el-container>
</template>

<script>
import { mapGetters } from 'vuex'
import { fetchRandSoup } from '@/api/soup'
import Header from './components/header'
import BlogMain from './components/BlogMain'
import Aside from './components/aside'
import Footer from './components/footer'
import ResizeMixin from './mixin/ResizeHandler'

export default {
  name: 'Blayout',
  components: {
    Header,
    BlogMain,
    Aside,
    Footer
  },
  mixins: [ResizeMixin],
  data() {
    return {
      soup: {},
      widthVar: '75%'
    }
  },
  computed: {
    ...mapGetters([
      'device'
    ])
  },
  mounted() {
    this.getSoup()
    this.widthVar = this.device === 'mobile' ? '100%' : '75%'
  },
  methods: {
    getSoup() {
      fetchRandSoup().then(response => {
        this.soup = response.data
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.blog-container {
  min-height: 100vh;
  background-color: #f8fafc;
  display: flex;
  flex-direction: column;
}

.fix-header {
  width: 100%;
}

.el-header {
  padding: 0;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 1000;
  box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
}

.blog-main-wrapper {
  padding-top: 88px; // 64px header + 24px padding
  padding-bottom: 40px;
  padding-left: 20px;
  padding-right: 20px;
  flex: 1 0 auto;
  display: flex;
  justify-content: center;
  overflow: visible !important;
}

@media (max-width: 768px) {
  .blog-main-wrapper {
    padding-left: 0 !important;
    padding-right: 0 !important;
  }
}

.main-layout {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 16px;
  display: flex;
  gap: 28px;
  align-items: flex-start;
  box-sizing: border-box;
}

.main-content {
  flex: 1;
  min-width: 0; // 防止flex item溢出
}

.main-aside {
  width: 300px;
  flex-shrink: 0;
  position: sticky;
  top: 88px;
}

.el-footer {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #0f172a;
  color: #94a3b8;
  border-top: 1px solid #1e293b;
  flex-shrink: 0;
}

@media (max-width: 992px) {
  .main-layout {
    flex-direction: column;
    gap: 20px;
  }

  .main-aside {
    width: 100%;
    position: static;
  }
}
</style>
