<template>
  <div v-if="ifDetail">
    <div v-html="articleToc" />
  </div>
  <div v-else>
    <el-card ref="ele">
      <div class="user-bio">
        <div class="user-education user-bio-section">
          <div class="user-bio-section-header">
            <svg-icon icon-class="education" />
            <span>每日一汤</span>
          </div>
          <div class="user-bio-section-body">
            <div class="text-muted">{{ soup.content }}</div>
          </div>
        </div>
      </div>

      <div class="user-bio">
        <div class="user-education user-bio-section">
          <div class="user-bio-section-header">
            <svg-icon icon-class="wx" />
            <span>扫不出吃亏，扫不出上当</span>
          </div>
          <div class="user-bio-section-body">
            <img :src="poorops">
          </div>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script>
import poorops from '@/assets/img/poorops.jpg'
import { mapState } from 'vuex'
// import {onMounted, nextTick, onBeforeUnmount} from 'vue'
// const rootEl = document.getElementById('article-html')
export default {
  name: 'Aside',
  props: {
    soup: {
      type: Object,
      default: () => {
        return {
          content: ''
        }
      }
    }
  },
  data() {
    return {
      poorops: poorops,
      ifDetail: true
      // articleHtml: ''
    }
  },
  computed: mapState({
    articleToc: (state) => state.app && state.app.articleToc
  }),
  created() {
    if (!this.$route.query.hasOwnProperty('id')) {
      this.ifDetail = false
    }
  },
  methods: {

  }
}
</script>

<style  lang="scss">
.catalog-list {
  font-weight: 600;
  padding-left: 10px;
  position: relative;
  font-size: 15px;
  &:first-child::before {
    content: "";
    position: absolute;
    top: 10px;
    left: 12px;
    bottom: 0;
    width: 2px;
    background-color: #ebedef;
    opacity: .8;
  }
  & > li > a {
    position: relative;
    padding-left: 16px;
    line-height: 20px;
    // @include catalogRound(0, 6px);
  }
  ul, li {
    padding: 0;
    margin: 0;
    list-style: none;
  }
  ul > li > a {
    font-size: 14px;
    color: #333333;
    padding-left: 36px;
    font-weight: 500;
    position: relative;
    // @include catalogRound(20px, 5px);
  }
  ul > ul > li > a {
    line-height: 20px;
    font-size: 14px;
    color: #333333;
    padding-left: 50px;
    font-weight: normal;
    // @include catalogRound;
  }
  a {
    color: #000;
    display: block;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    padding: 4px 0 4px 12px;
    &:hover {
      background-color: #ebedef;
    }
  }
}
</style>

<style lang="scss" scoped>
.el-card {
  float: right;
  width: 24%;
  position: fixed;
  left: 75%;
}
.box-center {
  margin: 0 auto;
  display: table;
}

.text-muted {
  color: #777;
  white-space: pre-line;
  line-height: 1.5em;
}

.user-profile {
  .user-name {
    padding: 15px;
    font-size: 20px;
    font-weight: 700;
  }

  .box-center {
    padding-top: 10px;
  }
}

.user-bio {
  margin-top: 20px;
  color: #606266;

  span {
    padding-left: 4px;
  }

  .user-bio-section {
    font-size: 14px;
    padding: 15px 0;

    .user-bio-section-header {
      border-bottom: 1px solid #dfe6ec;
      padding-bottom: 10px;
      margin-bottom: 10px;
      font-weight: bold;
    }
  }
}
</style>
