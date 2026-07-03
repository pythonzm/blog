<template>
  <div class="article-wrap">
    <div class="article-message">
      <h1 class="article-title">
        {{ article.title }}
      </h1>
      <div class="article-info">
        <span class="meta-item">
          <i class="el-icon-date" />
          发表于 {{ article.created_time | formatDateTime }}
        </span>
        <span v-if="article.updated_time" class="meta-divider">•</span>
        <span v-if="article.updated_time" class="meta-item">
          <i class="el-icon-edit" />
          更新于 {{ article.updated_time | formatDateTime }}
        </span>
        <span class="meta-divider">•</span>
        <span class="meta-item classify-badge">
          <i class="el-icon-folder-opened" />
          <router-link
            :to="{
              name: 'CTQArticle',
              query: { category: category.category_name }
            }"
          >{{ category.category_name }}</router-link>
        </span>
        <span class="meta-divider">•</span>
        <span class="meta-item">
          <i class="el-icon-view" />
          {{ views }} 次阅读
        </span>
        <span class="meta-divider">•</span>
        <span class="meta-item">
          <i class="el-icon-time" />
          预计阅读 {{ getReadTime(article.content) }} 分钟
        </span>
      </div>
    </div>

    <div class="article-view">
      <div v-mhighlight v-viewer class="md-body markdown-content" v-html="article.html" />
    </div>

    <div v-if="tags && tags.length > 0" class="tags-container">
      <div class="tags-label">
        <i class="el-icon-price-tag" />
        标签：
      </div>
      <div class="tags-list">
        <div
          v-for="(tag, index) in tags"
          :key="index"
          class="tag-capsule"
          @click="$router.push({ name: 'CTQArticle', query: { tag: tag.tag_name } })"
        >
          {{ tag.tag_name }}
        </div>
      </div>
    </div>

    <div class="comments-section">
      <Comments />
    </div>

    <script type="application/ld+json" v-html="jsonld">
      {}
    </script>
  </div>
</template>

<script>
import { fetchArticle } from '@/api/article'
import Comments from './comment'
import '@/assets/md.css'
import CodeCopy from '@/components/CodeCopy'
import Vue from 'vue'
import store from '@/store'

export default {
  name: 'ArticleDetail',
  components: {
    Comments
  },
  filters: {
    formatDateTime(time) {
      if (!time) return ''
      return time.includes(' ') ? time.split(' ')[0] : time
    }
  },
  data() {
    return {
      article: {},
      category: {},
      tags: [],
      views: 0,
      anchors: [],
      heightTitle: '',
      jsonld: {}
    }
  },
  watch: {
    '$route'(to, from) {
      if (!to.fullPath.includes('id=')) {
        store.dispatch('anchors/updateAnchors', [])
      } else {
        this.generateTOC()
      }
    }
  },
  created() {
    this.$nextTick(() => {
      const id = this.$route.query && this.$route.query.id
      this.fetchData(id)
    })
  },
  updated() {
    setTimeout(() => {
      document.querySelectorAll('[class*="lang-"]').forEach(el => {
        if (el.classList.contains('code-copy-added')) return
        const ComponentClass = Vue.extend(CodeCopy)
        const instance = new ComponentClass()
        instance.code = el.innerText
        instance.parent = el
        instance.$mount()
        el.classList.add('code-copy-added')
        el.appendChild(instance.$el)
      })
    }, 100)
  },
  destroyed() {
    store.dispatch('anchors/updateAnchors', [])
  },
  methods: {
    fetchData(id) {
      fetchArticle(id)
        .then(response => {
          this.article = response.data.article
          this.category = response.data.category
          this.tags = response.data.tags || []
          this.views = response.data.views
          this.generateTOC()
          document.title = `${this.article.title} - POOROPS`
          const created_time = this.convertToISO8601(this.article.created_time)
          const updated_time = this.article.updated_time ? this.convertToISO8601(this.article.updated_time) : created_time
          this.jsonld = {
            '@context': 'https://schema.org',
            '@type': 'Article',
            'headline': this.article.title,
            'datePublished': created_time,
            'dateModified': updated_time,
            'author': {
              '@type': 'Person',
              'name': 'POOROPS',
              'url': window.location.href
            }
          }
        })
        .catch(err => {
          console.log(err)
        })
    },
    generateTOC() {
      const parser = new DOMParser()
      const doc = parser.parseFromString(this.article.html, 'text/html')
      const headings = doc.querySelectorAll('h1, h2, h3, h4, h5, h6')
      const titles = Array.from(headings).filter((title) => !!title.innerText.trim())
      if (!titles.length) {
        store.dispatch('anchors/updateAnchors', [])
        return
      }
      const hTags = Array.from(new Set(titles.map((title) => title.tagName))).sort()

      this.anchors = Array.from(headings).map((heading) => {
        if (!heading.id) {
          heading.id = heading.textContent.trim().replace(/\s+/g, '-')
        }
        return {
          id: heading.id,
          text: heading.textContent,
          indent: hTags.indexOf(heading.tagName)
        }
      })
      store.dispatch('anchors/updateAnchors', this.anchors)
    },
    convertToISO8601(dateString, timezoneOffset = '+08:00') {
      const date = new Date(dateString)
      const isoDateString = date.toISOString()
      return isoDateString.slice(0, -1) + timezoneOffset
    },
    getReadTime(content) {
      if (!content) return 0
      const count = content.length
      return Math.ceil(count / 400) || 1
    }
  }
}
</script>

<style scoped lang="scss">
.code-copy-added {
  background-color: #1e293b;
  color: #f8fafc;
  padding: 30px 20px 20px;
  margin: 16px 0;
  text-align: left;
  border-radius: 8px;
  position: relative;
  border: 1px solid #334155;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
}
.code-copy-added:hover .copy-btn {
  opacity: 1;
}

.article-wrap {
  background-color: #ffffff;
  border-radius: 12px;
  border: 1px solid rgba(226, 232, 240, 0.8);
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
  padding: 40px;
  width: 100%;
  box-sizing: border-box;
  text-align: left;
}

.article-message {
  border-bottom: 1px solid #f1f5f9;
  padding-bottom: 24px;
  margin-bottom: 30px;
}

.article-title {
  font-size: 30px;
  font-weight: 800;
  color: #0f172a;
  line-height: 1.35;
  margin: 0 0 16px 0;
}

.article-info {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 12px;
  font-size: 13.5px;
  color: #64748b;

  .meta-item {
    display: inline-flex;
    align-items: center;
    gap: 4px;
  }

  .meta-divider {
    color: #e2e8f0;
  }

  .classify-badge {
    a {
      color: #1e293b;
      font-weight: 600;
      transition: color 0.2s ease;

      &:hover {
        color: #0f172a;
      }
    }
  }
}

.article-view {
  width: 100%;
  margin-top: 10px;
}

// 高级Markdown样式覆写
.markdown-content {
  font-size: 16px;
  color: #334155;
  line-height: 1.8;
  word-break: break-word;

  ::v-deep {
    p {
      margin-top: 0;
      margin-bottom: 20px;
    }

    h1, h2, h3, h4, h5, h6 {
      color: #0f172a;
      font-weight: 700;
      margin-top: 32px;
      margin-bottom: 16px;
      line-height: 1.3;
    }

    h1 {
      font-size: 24px;
      border-bottom: 1px solid #e2e8f0;
      padding-bottom: 8px;
    }

    h2 {
      font-size: 20px;
      border-bottom: 1px solid #f1f5f9;
      padding-bottom: 6px;
    }

    h3 {
      font-size: 18px;
      border-bottom: none;
    }

    a {
      color: #1e293b !important;
      text-decoration: none;
      border-bottom: 1px dashed rgba(30, 41, 59, 0.4);
      transition: all 0.2s ease;

      &:hover {
        color: #0f172a !important;
        border-bottom-style: solid;
      }
    }

    blockquote {
      background: linear-gradient(to right, #f8fafc, rgba(248, 250, 252, 0.2));
      border-left: 4px solid #1e293b;
      color: #475569;
      padding: 12px 20px;
      margin: 20px 0;
      border-radius: 0 8px 8px 0;
      font-style: italic;
    }

    code {
      background-color: #f1f5f9;
      color: #e11d48;
      padding: 3px 6px;
      border-radius: 4px;
      font-size: 0.9em;
      font-family: Menlo, Monaco, Consolas, "Courier New", monospace;
    }

    pre {
      background-color: #1e293b;
      padding: 16px 20px;
      border-radius: 8px;
      overflow-x: auto;
      margin: 20px 0;

      code {
        background-color: transparent;
        color: #f8fafc;
        padding: 0;
        border-radius: 0;
        font-size: 14px;
      }
    }

    img {
      max-width: 100%;
      border-radius: 8px;
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
      margin: 16px auto;
      display: block;
    }

    ul, ol {
      padding-left: 20px;
      margin-bottom: 20px;

      li {
        margin-bottom: 6px;
      }
    }

    table {
      width: 100%;
      border-collapse: collapse;
      margin: 24px 0;
      font-size: 14.5px;

      th {
        background-color: #f8fafc;
        border: 1px solid #e2e8f0;
        color: #0f172a;
        font-weight: 600;
        padding: 10px 12px;
      }

      td {
        border: 1px solid #e2e8f0;
        padding: 10px 12px;
        color: #475569;
      }

      tr:nth-child(even) {
        background-color: #fafbfc;
      }
    }
  }
}

.tags-container {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  padding: 24px 0;
  margin-top: 30px;
  border-top: 1px solid #f1f5f9;
  border-bottom: 1px solid #f1f5f9;

  .tags-label {
    font-size: 14px;
    font-weight: 600;
    color: #475569;
    display: flex;
    align-items: center;
    gap: 4px;
  }

  .tags-list {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    margin-left: 8px;
  }
}

.tag-capsule {
  background-color: #f1f5f9;
  color: #475569;
  font-size: 12.5px;
  font-weight: 500;
  padding: 4px 12px;
  border-radius: 20px;
  cursor: pointer;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);

  &:hover {
    background-color: rgba(30, 41, 59, 0.08);
    color: #1e293b;
    transform: translateY(-1px);
  }
}

.comments-section {
  margin-top: 30px;
}

@media (max-width: 768px) {
  .article-wrap {
    padding: 20px;
  }

  .article-title {
    font-size: 24px;
  }
}
</style>
