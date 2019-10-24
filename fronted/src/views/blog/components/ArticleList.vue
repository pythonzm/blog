<template>
  <div class="dashboard-container">
    <div class="title-article">
      <div class="list-card" v-for="item in list" :key="item.id">
        <router-link
          class="pan-btn blue-btn"
          :to="{
            name: 'CTArticle',
            query: { id: item.id }
          }"
        >
          {{ item.title }}<span>{{ item.created_time | formatDate }}</span>
        </router-link>
      </div>
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
import { parseTime } from '@/utils'
export default {
  name: "ArticleList",
  components: { Pagination },
  filters: {
    formatDate (time) {
      return time.slice(0, time.indexOf(" "))
    },
  },
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
  border: 1px solid #ebeef5;
  background-color: #fff;
  color: #303133;
  transition: 0.3s;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}
a {
  width: 100%;
}
span {
  position: absolute;
  right: 36px;
}
.pan-btn {
  font-size: 25px;
  font-weight: 900;
  color: #000000;
}
.blue-btn {
  background: unset;
}

.list-card {
  border-radius: 6px;
  text-align: left;
  padding: 1rem;
  border-bottom: 1px solid rgba(0, 0, 0, 0.1);
}
</style>
