<template>
  <div class="articles-list-container">
    <div v-if="total === 0" class="search-empty-state">
      <i class="el-icon-document-delete empty-icon" />
      <p v-if="listQuery.q !== undefined" v-cloak class="empty-text">
        暂无包含关键字 <span>"{{ listQuery.q }}"</span> 的文章
      </p>
      <p v-else class="empty-text">暂无文章！</p>
    </div>

    <div v-else class="articles-feed">
      <!-- 文章卡片列表 -->
      <div
        v-for="item in list"
        :key="item.id"
        class="article-card"
        @mouseenter="onMouseEnter"
        @mouseleave="onMouseLeave"
      >
        <!-- 极客 IDE 风格侧边指示线 -->
        <div class="active-gutter-line" />

        <div class="card-content">
          <!-- 极客范文件路径指示器 -->
          <div class="card-geek-header">
            <span class="file-path">
              <i class="el-icon-document" /> {{ getGeekPath(item.title) }}
            </span>
            <span class="file-size">{{ getGeekSize(item.title) }}</span>
          </div>

          <div class="card-header">
            <router-link
              class="article-title-link"
              :to="{
                name: 'CTQArticle',
                query: { id: item.id }
              }"
            >
              <span class="prompt-symbol">$</span> {{ item.title }}
            </router-link>
          </div>

          <div class="card-footer">
            <div class="meta-info">
              <span class="meta-date">
                <i class="el-icon-date icon" />
                {{ item.created_time | formatDate }}
              </span>
            </div>

            <router-link
              class="read-more-link"
              :to="{
                name: 'CTQArticle',
                query: { id: item.id }
              }"
            >
              阅读全文
              <i class="el-icon-right arrow-icon" />
            </router-link>
          </div>
        </div>
      </div>

      <div class="pagination-wrapper">
        <pagination
          :total="total"
          :page.sync="listQuery.page"
          :limit.sync="listQuery.limit"
          :pager-count="device === 'mobile' ? 5 : 7"
          layout="prev, pager, next"
          @pagination="getList"
        />
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { fetchList } from '@/api/article'
import Pagination from '@/components/Pagination'
import ResizeMixin from '@/blayout/mixin/ResizeHandler'
import gsap from 'gsap'
export default {
  name: 'ArticleList',
  components: { Pagination },
  filters: {
    formatDate(time) {
      return time ? time.split(' ')[0] : ''
    }
  },
  mixins: [ResizeMixin],
  data() {
    return {
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10,
        category: undefined,
        tag: undefined,
        q: undefined
      }
    }
  },
  computed: {
    ...mapGetters([
      'device'
    ])
  },
  created() {
    if ('category' in this.$route.query) {
      this.listQuery.category = this.$route.query.category
    }
    if ('tag' in this.$route.query) {
      this.listQuery.tag = this.$route.query.tag
    }
    if ('q' in this.$route.query) {
      this.listQuery.q = this.$route.query.q
    }
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      fetchList(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.listLoading = false
        this.animateCards()
      })
    },
    animateCards() {
      this.$nextTick(() => {
        const prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches
        if (prefersReducedMotion) return

        // Ensure elements exist before animating
        if (document.querySelectorAll('.article-card').length > 0) {
          gsap.fromTo('.article-card',
            { opacity: 0, y: 25 },
            {
              opacity: 1,
              y: 0,
              duration: 0.5,
              stagger: 0.06,
              ease: 'power2.out',
              clearProps: 'transform,opacity'
            }
          )
        }
      })
    },
    getGeekPath(title) {
      if (!title) return ''
      // 生成符合linux命名的短网址拼音/英文格式路径
      const slug = title
        .toLowerCase()
        .replace(/[^\w\u4e00-\u9fa5]/g, '-')
        .replace(/-+/g, '-')
        .replace(/^-+|-+$/g, '')
        .slice(0, 60)
      return `~/posts/${slug || 'article'}.md`
    },
    getGeekSize(title) {
      if (!title) return '1.0 KB'
      const bytes = title.length * 48 + 512
      return `${(bytes / 1024).toFixed(2)} KB`
    },
    onMouseEnter(event) {
      const prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches
      if (prefersReducedMotion) return

      const card = event.currentTarget
      const line = card.querySelector('.active-gutter-line')
      const prompt = card.querySelector('.prompt-symbol')

      gsap.to(line, {
        scaleY: 1,
        duration: 0.35,
        ease: 'power2.out',
        overwrite: 'auto'
      })

      gsap.to(prompt, {
        x: 4,
        color: '#0f172a',
        duration: 0.25,
        ease: 'back.out(1.5)',
        overwrite: 'auto'
      })
    },
    onMouseLeave(event) {
      const prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches
      if (prefersReducedMotion) return

      const card = event.currentTarget
      const line = card.querySelector('.active-gutter-line')
      const prompt = card.querySelector('.prompt-symbol')

      gsap.to(line, {
        scaleY: 0,
        duration: 0.25,
        ease: 'power2.in',
        overwrite: 'auto'
      })

      gsap.to(prompt, {
        x: 0,
        color: '#94a3b8',
        duration: 0.25,
        ease: 'power2.out',
        overwrite: 'auto'
      })
    }
  }
}
</script>

<style lang="scss">
// 全局分页样式微调，符合现代UI
.pagination-wrapper {
  margin-top: 30px;
  display: flex;
  justify-content: center;

  .el-pagination.is-background .el-pager li:not(.disabled).active {
    background-color: #1e293b !important;
    color: #fff !important;
  }

  .el-pagination.is-background .el-pager li:not(.disabled):hover {
    color: #0f172a !important;
  }

  .el-pagination.is-background .btn-next,
  .el-pagination.is-background .btn-prev,
  .el-pagination.is-background .el-pager li {
    border-radius: 8px !important;
    background-color: #ffffff !important;
    border: 1px solid #e2e8f0;
    color: #64748b !important;
    transition: all 0.3s ease;

    &:hover {
      border-color: #1e293b;
      color: #1e293b !important;
    }
  }
}

@media (max-width: 480px) {
  .pagination-wrapper {
    ::v-deep .pagination-container {
      padding: 16px 4px !important;
      background: transparent !important;
    }

    .el-pagination {
      padding: 0 !important;
      display: flex;
      justify-content: center;
      align-items: center;
    }

    .el-pagination.is-background .btn-next,
    .el-pagination.is-background .btn-prev,
    .el-pagination.is-background .el-pager li {
      min-width: 26px !important;
      height: 26px !important;
      line-height: 26px !important;
      font-size: 12px !important;
      margin: 0 2px !important;
      border-radius: 6px !important;
    }
  }
}
</style>

<style scoped lang="scss">
[v-cloak] {
  display: none;
}

.articles-list-container {
  width: 100%;
  min-width: 0;
  overflow: hidden;
}

// 空状态美化
.search-empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px 20px;
  background: #ffffff;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);

  .empty-icon {
    font-size: 48px;
    color: #94a3b8;
    margin-bottom: 16px;
  }

  .empty-text {
    font-size: 16px;
    color: #64748b;
    margin: 0;

    span {
      color: #1e293b;
      font-weight: 600;
    }
  }
}

.articles-feed {
  display: flex;
  flex-direction: column;
  gap: 20px;
  width: 100%;
  min-width: 0;
  box-sizing: border-box;
}

// 极客 IDE 风格侧边指示线
.active-gutter-line {
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  width: 4px;
  background-color: #1e293b;
  transform: scaleY(0);
  transform-origin: center;
  z-index: 2;
}

// 文章卡片美化
.article-card {
  position: relative;
  width: 100%;
  min-width: 0;
  box-sizing: border-box;
  background: #ffffff;
  border-radius: 12px;
  border: 1px solid rgba(226, 232, 240, 0.8);
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05), 0 2px 4px -1px rgba(0, 0, 0, 0.03);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  overflow: hidden;
  text-align: left;

  &:hover {
    transform: translateY(-3px);
    box-shadow: 0 12px 20px -8px rgba(30, 41, 59, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.02);
    border-color: rgba(30, 41, 59, 0.15);

    .article-title-link {
      color: #0f172a;
    }

    .read-more-link {
      color: #0f172a;

      .arrow-icon {
        transform: translateX(4px);
      }
    }
  }
}

.card-content {
  padding: 24px;
}

@media (max-width: 480px) {
  .card-content {
    padding: 16px;
  }
}

// 极客卡片头部
.card-geek-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-family: Consolas, Monaco, 'Courier New', monospace;
  font-size: 12px;
  color: #64748b;
  margin-bottom: 12px;
  padding-bottom: 8px;
  border-bottom: 1px dashed #e2e8f0;

  .file-path {
    display: flex;
    align-items: center;
    gap: 4px;
    color: #475569;
  }

  .file-size {
    opacity: 0.85;
  }
}

.card-header {
  margin-bottom: 16px;
}

.prompt-symbol {
  display: inline-block;
  color: #94a3b8;
  font-family: Consolas, Monaco, 'Courier New', monospace;
  font-weight: 500;
  opacity: 0.65;
}

.article-title-link {
  font-size: 20px;
  font-weight: 700;
  color: #0f172a;
  line-height: 1.4;
  text-decoration: none;
  transition: color 0.3s ease;
  display: block;
  word-break: break-all;
}

@media (max-width: 480px) {
  .article-title-link {
    font-size: 18px;
  }
}

.card-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-top: 1px solid #f1f5f9;
  padding-top: 16px;
  margin-top: 8px;
}

.meta-info {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #64748b;

  .icon {
    margin-right: 4px;
  }

  .meta-date, .meta-read-time {
    display: flex;
    align-items: center;
  }

  .meta-divider {
    color: #cbd5e1;
  }
}

.read-more-link {
  font-size: 13px;
  font-weight: 600;
  color: #1e293b;
  display: flex;
  align-items: center;
  gap: 4px;
  text-decoration: none;
  transition: all 0.3s ease;

  .arrow-icon {
    transition: transform 0.3s ease;
  }
}
</style>
