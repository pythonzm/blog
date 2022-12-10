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
      views: 0
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
          const toc = this.genToc(response.data.article.html)
          console.log(toc)
          store.dispatch('app/setToc', toc)
        })
        .catch(err => {
          console.log(err)
        })
    },
    genToc(articleHtml) {
      const toc = articleHtml.match(/<[hH][1-6]>.*?<\/[hH][1-6]>/g)
      console.log(toc)
      const levelStack = []
      let result = ''
      const addStartUL = () => { result += '<ul class="catalog-list">' }
      const addEndUL = () => { result += '</ul>\n' }
      const addLI = (index, itemText) => { result += '<li><a name="link" class="toc-link' + '-#' + index + '" href="#' + index + '">' + itemText + '</a></li>\n' }
      toc.forEach((item, index) => {
        const _toc = `<div name='toc-title' id='${index}'>${item} </div>`
        articleHtml = articleHtml.replace(item, _toc)
      })

      toc.forEach(function(item, index) {
        const itemText = item.replace(/<[^>]+>/g, '') // 匹配h标签的文字
        const itemLabel = item.match(/<\w+?>/)[0] // 匹配h?标签<h?>
        let levelIndex = levelStack.indexOf(itemLabel) // 判断数组里有无<h?>
        // 没有找到相应<h?>标签，则将新增ul、li
        if (levelIndex === -1) {
          levelStack.unshift(itemLabel)
          addStartUL()
          addLI(index, itemText)
        }
        // 找到了相应<h?>标签，并且在栈顶的位置则直接将li放在此ul下
        else if (levelIndex === 0) {
          addLI(index, itemText)
        }
        // 找到了相应<h?>标签，但是不在栈顶位置，需要将之前的所有<h?>出栈并且打上闭合标签，最后新增li
        else {
          while (levelIndex--) {
            levelStack.shift()
            addEndUL()
          }
          addLI(index, itemText)
        }
      })
      // 如果栈中还有<h?>，全部出栈打上闭合标签
      while (levelStack.length) {
        levelStack.shift()
        addEndUL()
      }
      return result
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
