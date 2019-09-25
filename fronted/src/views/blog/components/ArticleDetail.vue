<template>
  <div id="article">
    <div class="article-warp">
      <div class="article-message">
        <p class="article-title">
          {{ article.title }}
        </p>
        <div class="article-info">
          <svg-icon icon-class="calendar" />
          发表于 {{ article.created_time }} •
          <svg-icon icon-class="category" />
          <span class="classify"
            ><router-link
              :to="{
                name: 'CTArticle',
                query: { category: category.category_name }
              }"
              >{{ category.category_name }}</router-link
            ></span
          >
          •
          <svg-icon icon-class="eye-open" />{{ views }}次围观
        </div>
      </div>
      <div class="article-view">
        <div v-html="article.html" v-mhighlight></div>
      </div>

      <div class="tags">
        <div
          v-for="(tag, index) in tags"
          :key="index"
          class="tag"
          @click="
            $router.push({ name: 'CTArticle', query: { tag: tag.tag_name } })
          "
        >
          <svg-icon icon-class="tag" />
          {{ tag.tag_name }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import { fetchArticle } from '@/api/article'
import hljs from "highlight.js";
import "highlight.js/styles/monokai-sublime.css";
export default {
  name: 'ArticleDetail',
  data () {
    return {
      article: {},
      category: {},
      tags: {},
      views: 0,
    }
  },
  created () {
    const id = this.$route.query && this.$route.query.id
    this.fetchData(id)
  },
  methods: {
    fetchData (id) {
      fetchArticle(id)
        .then(response => {
          this.article = response.data.article
          this.category = response.data.category
          this.tags = response.data.tags
          this.views = response.data.views
        })
        .catch(err => {
          console.log(err)
        })
    }
  },
  directives: {
    'mhighlight': function (el) {
      let blocks = el.querySelectorAll("pre code");
      blocks.forEach(block => {
        hljs.highlightBlock(block);
      });
    }
  }

}
</script>

<style scoped>
.article {
  padding: 30px 10px;
}
.article-wrap {
  position: relative;
  animation: show-data-v-ce9b41d0 0.8s;
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
</style>