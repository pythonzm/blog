<template>
  <div id="article">
    <div class="article-warp">
      <div class="article-message">
        <p class="article-title">
          {{ article.title }}
        </p>
        <div class="article-info">
          <i class="iconfont icon-calendar"></i>
          发表于 {{ article.created_time }} •
          <i class="iconfont icon-folder"></i>
          <span class="classify">python</span> •
          <i class="iconfont icon-eye"></i>
          0次围观
        </div>
      </div>
      <div v-html="article.html"></div>

      <!-- <div class="tags">
        <div
          v-for="(tag, index) in tags"
          :key="index"
          class="tag"
          @click="$router.push({name: 'articleList', query:{type: 'tag', id: tag.id}})">
          <i class="iconfont icon-tag"></i>
          {{ tag.name }}
        </div>
      </div> -->
    </div>
  </div>
</template>

<script>

import { fetchArticle } from '@/api/article'
export default {
  name: 'ArticleDetail',
  data () {
    return {
      article: {}
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
        })
        .catch(err => {
          console.log(err)
        })
    }
  }
}
</script>