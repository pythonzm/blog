<template>
  <div v-if="ifDetail">
    <article-detail />
  </div>
  <div v-else>
    <article-list />
  </div>
</template>

<script>
import ArticleList from './components/ArticleList'
import ArticleDetail from './components/ArticleDetail'

export default {
  name: 'Article',
  components: { ArticleList, ArticleDetail },
  data () {
    return {
      ifDetail: true
    }
  },
  created () {
    if (!this.$route.query.hasOwnProperty('id')) {
      this.ifDetail = false
    }
  }
}
</script>