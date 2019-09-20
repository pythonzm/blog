<template>
  <div>
    <div class="ca" v-for="category in list" :key="category.id">
      <router-link
        :to="{
          name: 'CTArticle',
          query: { category: category.category_name }
        }"
        >{{ category.category_name }}</router-link
      >
    </div>
  </div>
</template>

<script>
import { fetchCategoryList } from '@/api/category'
export default {
  name: "Category",
  data () {
    return {
      list: null,
      listLoading: true,
    }
  },
  created () {
    this.getList()
  },
  methods: {
    async getList () {
      this.listLoading = true
      const { data } = await fetchCategoryList()
      this.list = data
      this.listLoading = false
    },
  }
}
</script>

<style scoped>
.ca {
  display: inline-block;
  margin: 10px;
  color: #555;
  text-decoration: none;
  outline: none;
  border-bottom: 1px solid #999;
  word-wrap: break-word;
}
</style>