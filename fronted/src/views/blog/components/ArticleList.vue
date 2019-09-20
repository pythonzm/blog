<template>
  <div class="dashboard-container">
    <div class="title-article list-card" v-for="item in list" :key="item.id">
      <router-link
        :to="{
          name: 'CTArticle',
          query: { id: item.id }
        }"
        ><h1>{{ item.title }}</h1></router-link
      >
      <hr />
    </div>

    <pagination
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      layout="prev, pager, next"
      @pagination="getList"
    />
  </div>
</template>

<script>
import { fetchList } from '@/api/article'
import Pagination from '@/components/Pagination'
export default {
  name: "ArticleList",
  components: { Pagination },
  props: {
    searchType: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 5,
        category: undefined,
        tag: undefined
      }
    }
  },
  created () {
    if (this.$route.query.hasOwnProperty('category')) {
      this.listQuery.category = this.$route.query.category
    }
    if (this.$route.query.hasOwnProperty('tag')) {
      this.listQuery.tag = this.$route.query.tag
    }
    this.getList()
  },
  methods: {
    getList () {
      this.listLoading = true
      fetchList(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.listLoading = false
      })
    },
  }
}
</script>

<style>
.el-pagination.is-background .el-pager li:not(.disabled).active {
  background-color: #000000 !important;
}
</style>

<style scoped>
.title-article {
  overflow: hidden;
  position: relative;
}
.list-card {
  border-radius: 6px;
}
.list-card:hover {
  background: #dcdfe6;
}
.title-article .title-msg {
  margin-bottom: 10px;
}
.title-article .title-msg span {
  color: #999;
  margin-right: 10px;
}
</style>
