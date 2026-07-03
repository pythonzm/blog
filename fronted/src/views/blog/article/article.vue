<template>
  <div v-if="ifDetail" class="article-route-wrapper">
    <article-detail />
  </div>
  <div v-else class="article-route-wrapper">
    <article-list />
  </div>
</template>

<script>
import ArticleList from './components/ArticleList'
import ArticleDetail from './components/ArticleDetail'

export default {
  name: 'Article',
  components: { ArticleList, ArticleDetail },
  data() {
    return {
      ifDetail: true
    }
  },
  created() {
    if (!('id' in this.$route.query)) {
      this.ifDetail = false
    }
  }
}
</script>

<style scoped>
.article-route-wrapper {
  width: 100%;
  box-sizing: border-box;
}
</style>
