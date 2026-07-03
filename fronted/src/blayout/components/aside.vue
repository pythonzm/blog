<template>
  <div class="aside-wrapper">
    <!-- 文章目录 -->
    <div v-if="anchors.length !== 0" class="anchors-card">
      <div class="aside-section-header">
        <i class="el-icon-notebook-2" style="margin-right: 6px; color: #1e293b;" />
        <span>目录导航</span>
      </div>
      <div class="anchors-list">
        <p
          v-for="anchor in anchors"
          :key="anchor.id"
          :style="{ paddingLeft: `${anchor.indent * 12 + 12}px` }"
          :class="anchor.text === heightTitle ? 'anchor-active' : 'anchor-inactive'"
          @click="scrollToAnchor(anchor.text)"
        >
          <span class="anchor-bullet" />
          {{ anchor.text }}
        </p>
      </div>
    </div>

    <!-- 侧边栏内容 -->
    <div v-else class="bio-cards">
      <div class="aside-card soup-card">
        <div class="aside-section-header">
          <i class="el-icon-chat-line-round" style="margin-right: 6px; color: #1e293b;" />
          <span>每日一汤</span>
        </div>
        <div class="soup-body">
          <span class="quote-mark">“</span>
          <p class="soup-text">{{ soup.content }}</p>
          <span class="quote-mark-end">”</span>
        </div>
      </div>

      <div class="aside-card qr-card">
        <div class="aside-section-header">
          <i class="el-icon-mobile-phone" style="margin-right: 6px; color: #1e293b;" />
          <span>扫码关注</span>
        </div>
        <div class="qr-container">
          <div class="qr-item">
            <img :src="poorops" alt="公众号">
            <p class="qr-label">微信公众号</p>
          </div>
          <div class="qr-item">
            <img :src="xiaochengxu" alt="小程序">
            <p class="qr-label">微信小程序</p>
          </div>
        </div>
      </div>
    </div>
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
.aside-wrapper {
  width: 100%;
}

.aside-card {
  background: #ffffff;
  border-radius: 12px;
  border: 1px solid rgba(226, 232, 240, 0.8);
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05), 0 2px 4px -1px rgba(0, 0, 0, 0.03);
  padding: 20px;
  margin-bottom: 20px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  text-align: left;

  &:hover {
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.05), 0 4px 6px -2px rgba(0, 0, 0, 0.02);
    border-color: rgba(30, 41, 59, 0.15);
  }
}

.aside-section-header {
  display: flex;
  align-items: center;
  font-size: 15px;
  font-weight: 700;
  color: #1e293b;
  border-bottom: 1px solid #f1f5f9;
  padding-bottom: 12px;
  margin-bottom: 14px;
}

// 每日一汤卡片
.soup-card {
  background: linear-gradient(to bottom right, #ffffff, #fdfeff);

  .soup-body {
    position: relative;
    padding: 8px 12px;
  }

  .quote-mark {
    font-size: 32px;
    font-family: Georgia, serif;
    color: #cbd5e1;
    opacity: 0.6;
    line-height: 1;
    position: absolute;
    top: -10px;
    left: -4px;
  }

  .quote-mark-end {
    font-size: 32px;
    font-family: Georgia, serif;
    color: #cbd5e1;
    opacity: 0.6;
    line-height: 1;
    position: absolute;
    bottom: -22px;
    right: -4px;
  }

  .soup-text {
    font-size: 14px;
    color: #475569;
    line-height: 1.6;
    margin: 4px 0;
    font-style: italic;
    white-space: pre-line;
    position: relative;
    z-index: 1;
  }
}

// 关注二维码卡片
.qr-card {
  .qr-container {
    display: flex;
    justify-content: space-around;
    align-items: center;
    padding: 8px 0;
  }

  .qr-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;

    img {
      width: 105px;
      height: 105px;
      border-radius: 8px;
      border: 1px solid #e2e8f0;
      padding: 4px;
      background: #ffffff;
      transition: all 0.3s ease;

      &:hover {
        transform: scale(1.05);
        border-color: #1e293b;
        box-shadow: 0 4px 12px rgba(30, 41, 59, 0.1);
      }
    }

    .qr-label {
      font-size: 12px;
      color: #64748b;
      margin-top: 8px;
      margin-bottom: 0;
      font-weight: 500;
    }
  }
}

// 文章目录卡片
.anchors-card {
  background: #ffffff;
  border-radius: 12px;
  border: 1px solid rgba(226, 232, 240, 0.8);
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
  padding: 20px;
  max-height: 80vh;
  display: flex;
  flex-direction: column;
  text-align: left;

  .anchors-list {
    overflow-y: auto;
    flex: 1;
    padding-right: 4px;

    /* 自定义滚动条 */
    &::-webkit-scrollbar {
      width: 4px;
    }
    &::-webkit-scrollbar-thumb {
      background: #cbd5e1;
      border-radius: 4px;
    }
  }

  p {
    font-size: 13px;
    line-height: 1.5;
    margin: 8px 0;
    cursor: pointer;
    border-radius: 6px;
    padding: 6px 12px;
    transition: all 0.2s ease;
    display: flex;
    align-items: center;
    position: relative;
  }

  .anchor-bullet {
    width: 5px;
    height: 5px;
    border-radius: 50%;
    background: #94a3b8;
    margin-right: 8px;
    display: inline-block;
    transition: all 0.2s ease;
  }

  .anchor-inactive {
    color: #64748b;

    &:hover {
      background-color: #f1f5f9;
      color: #1e293b;

      .anchor-bullet {
        background: #1e293b;
        transform: scale(1.2);
      }
    }
  }

  .anchor-active {
    background-color: rgba(30, 41, 59, 0.06);
    color: #0f172a;
    font-weight: 600;

    .anchor-bullet {
      background: #0f172a;
      transform: scale(1.4);
    }
  }
}
</style>
