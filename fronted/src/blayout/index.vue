<template>
  <el-container>
    <el-header>
      <div class="fix-header">
        <Header :mobile="mobile" />
      </div>
    </el-header>

    <el-main>
      <a
        href="https://github.com/pythonzm"
        target="_blank"
        style="position: fixed;"
      >
        <img
          width="149"
          height="149"
          :src="forkme"
          class="attachment-full size-full"
          alt="Fork me on GitHub"
          data-recalc-dims="1"
        />
      </a>
      <div class="main">
        <BlogMain :style="{ width: widthVar }" />
        <div v-if="device !== 'mobile'">
          <Aside :soup="soup" />
        </div>
      </div>
    </el-main>
    <el-footer>
      <Footer />
    </el-footer>
  </el-container>
</template>

<script>
import { fetchRandSoup } from "@/api/soup";
import Header from "./components/header";
import BlogMain from "./components/BlogMain";
import Aside from "./components/aside";
import Footer from "./components/footer";
import ResizeMixin from "./mixin/ResizeHandler";
import forkme from "@/assets/img/forkme.png";
export default {
  name: "Blayout",
  components: {
    Header,
    BlogMain,
    Aside,
    Footer
  },
  mixins: [ResizeMixin],
  computed: {
    device() {
      return this.$store.state.app.device;
    },
    classObj() {
      return {
        mobile: this.device === "mobile"
      };
    }
  },
  data() {
    return {
      soup: {},
      widthVar: "75%",
      mobile: false,
      forkme: forkme
    };
  },
  mounted() {
    this.getSoup();
    this.widthVar = this.device === "mobile" ? "100%" : "75%";
    this.mobile = this.device === "mobile" ? true : false;
  },
  methods: {
    getSoup() {
      fetchRandSoup().then(response => {
        this.soup = response.data;
      });
    }
  }
};
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
