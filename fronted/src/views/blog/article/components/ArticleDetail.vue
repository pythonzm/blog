<template>
  <div id="article">
    <div class="article-wrap">
      <div class="article-message">
        <p class="article-title">
          {{ article.title }}
        </p>
        <div class="article-info">
          <svg-icon icon-class="calendar" />
          发表于 {{ article.created_time }} •
          <span
            v-if="article.updated_time !== ''"
          ><svg-icon icon-class="calendar" /> 更新于
            {{ article.updated_time }} •</span>
          <svg-icon icon-class="category" />
          <span
            class="classify"
          ><router-link
            :to="{
              name: 'CTQArticle',
              query: { category: category.category_name }
            }"
          >{{ category.category_name }}</router-link></span>
          •
          <svg-icon icon-class="eye-open" />{{ views }}次围观
        </div>
      </div>
      <div class="article-view">
        <div v-mhighlight v-viewer class="md-body" v-html="article.html" />
      </div>

      <div class="tags">
        <div
          v-for="(tag, index) in tags"
          :key="index"
          class="tag"
          @click="
            $router.push({ name: 'CTQArticle', query: { tag: tag.tag_name } })
          "
        >
          <svg-icon icon-class="tag" />
          {{ tag.tag_name }}
        </div>
      </div>
      <Comments />
    </div>
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
  data() {
    return {
      article: {},
      category: {},
      tags: {},
      views: 0,
      anchors: [],
      heightTitle: ''
    }
  },
  watch: {
    // 监听$route对象的变化
    '$route'(to, from) {
      // 当路由变化时，更新页面标题
      console.log(to.fullPath)
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
        /* 手动挂载 */
        instance.$mount()
        el.classList.add('code-copy-added')
        el.appendChild(instance.$el)
      })
    }, 100)
  },
  methods: {
    fetchData(id) {
      fetchArticle(id)
        .then(response => {
          this.article = response.data.article
          this.category = response.data.category
          this.tags = response.data.tags
          this.views = response.data.views
          this.generateTOC()
        })
        .catch(err => {
          console.log(err)
        })
    },
    generateTOC() {
      // 使用浏览器DOM API解析HTML内容
      const parser = new DOMParser()
      const doc = parser.parseFromString(this.article.html, 'text/html')
      // 提取所有标题元素
      const headings = doc.querySelectorAll('h1, h2, h3, h4, h5, h6')
      const titles = Array.from(headings).filter((title) => !!title.innerText.trim())
      if (!titles.length) {
        return
      }
      const hTags = Array.from(new Set(titles.map((title) => title.tagName))).sort()

      // 生成目录结构
      this.anchors = Array.from(headings).map((heading) => {
        // 确保每个标题都有唯一的ID
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
      console.log(this.anchors)
    }
  }
}
</script>

<style scoped>
.code-copy-added {
  background-color: #282c34;
  color: white;
  padding: 25px 20px;
  margin: 10px 0;
  text-align: left;
  border-radius: 3px;
  position: relative;
}
.code-copy-added:hover .copy-btn {
  opacity: 1;
}

.article {
  padding: 30px 10px;
}
.article-wrap {
  position: relative;
  padding: 30px;
  width: 100%;
  background-color: #fff;
  box-shadow: 0 0 5px 0 rgba(38, 42, 48, 0.1);
}
.article-message {
  display: flex;
  -ms-flex-direction: column;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.article-title {
  font-size: 26px;
  font-weight: 700;
  margin: 0;
}
.article-info {
  font-size: 14px;
  margin: 20px 0;
  color: #666;
  display: flex;
  flex-direction: row;
  justify-content: center;
  flex-wrap: wrap;
}
svg {
  margin: 0 5px;
}
.article-view {
  position: relative;
  width: 100%;
  margin-top: 10px;
  text-align: left;
}
.tags {
  width: 100%;
  padding: 10px 0;
  display: flex;
  flex-direction: row;
  align-items: center;
  flex-wrap: wrap;
  border-bottom: 1px solid #eee;
}
.tag {
  color: #fff;
  padding: 5px;
  background-color: #262a30;
  font-size: 12px;
  margin-right: 5px;
  border-radius: 5px;
  position: relative;
  line-height: 1;
  transition: all 0.3s;
  cursor: pointer;
  -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
}

.classify {
  color: #444;
  border-bottom: 1px solid #262a30;
  cursor: pointer;
  -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
  margin-right: 5px;
}
.cl-wrapper {
  position: relative;
}

.cl-wrapper ul,
.cl-wrapper li {
  margin: 0;
  -moz-padding-start: 12px;
  -webkit-padding-start: 12px;
  list-style: none;
}

.cl-wrapper li > .cl-link.cl-link-active {
  color: rgba(66, 185, 131, 0.9);
  transition: 0.5s;
}

.cl-wrapper li > .cl-transform.cl-link-active {
  transform: translate(3px);
}

.cl-wrapper .cl-link {
  cursor: pointer;
  color: rgba(52, 73, 94, 0.5);
  font-size: 13px;
  transition: all 0.3s cubic-bezier(0.23, 1, 0.32, 1);
}

.cl-wrapper .cl-marker {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
}

.cl-wrapper .cl-marker path {
  transition: all 0.3s ease;
}
</style>
