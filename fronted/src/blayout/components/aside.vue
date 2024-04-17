<template>
  <div>
    <div v-if="anchors.length !== 0" class="anchors">
      <div class="user-education user-bio-section">
        <p
          v-for="anchor in anchors"
          :key="anchor.id"
          :style="{ padding: `5px 5px 5px ${anchor.indent * 20 + 5}px` }"
          :class="anchor.text===heightTitle?'title-active':'title-inactive'"
          @click="scrollToAnchor(anchor.text)"
        >{{ anchor.text }}</p>
      </div>
    </div>
    <el-card v-else ref="ele">
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
            <img :src="xiaochengxu">
          </div>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script>
import poorops from '@/assets/img/poorops.jpg'
import xiaochengxu from '@/assets/img/xiaochengxu.jpg'
import { mapGetters } from 'vuex'
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
      xiaochengxu: xiaochengxu,
      heightTitle: ''
    }
  },
  computed: {
    ...mapGetters([
      'anchors'
    ])
  },
  methods: {
    scrollToAnchor(anchorText) {
      const headings = document.querySelectorAll('h1, h2, h3, h4, h5, h6')
      for (var i = 0; i < headings.length; i++) {
        var heading = headings[i]
        if (heading.textContent.trim() === anchorText) {
          console.log('Found the heading:', heading)
          this.heightTitle = anchorText
          heading.scrollIntoView({ behavior: 'smooth', block: 'center' })
          break
        }
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.title-active {
  border-left: 2px solid #777;
  box-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3); /* 水平偏移量，垂直偏移量，模糊半径，阴影颜色 */
  // padding: 10px; /* 添加内边距以使阴影效果更明显 */
  background-color: #000; /* 设置背景色以突出显示阴影效果 */
  color: #fff;
  cursor: pointer;
  margin: 5px;
}
.title-inactive {
  cursor: pointer;
  margin: 3px;
}
.title-inactive:hover {
  background-color: #e4e4e4;
}
.anchors {
  text-align: left;
  width: 24%;
  position: fixed;
  left: 75%;
  overflow-y: auto;
  max-height: 80vh;
  font-size: small;
  border: 1px #ccc;
  box-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
}
.el-card {
  float: right;
  width: 24%;
  position: fixed;
  left: 75%;
}

img {
  width: 100px;
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

.user-bio {
  // margin-top: 20px;
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
