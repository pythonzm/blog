<template>
  <div>
    <div class="ca" v-for="category in list" :key="category.id">
      <router-link
        class="pan-btn blue-btn"
        :style="{
          fontSize: category.category_name.length * 2 + 13 + 'px',
          background: getColor(category.category_name.length)
        }"
        :to="{
          name: 'CTQArticle',
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
    getColor (count) {
      let alpha = 0.2
      if (count < 5) {
        alpha = 0.2
      } else if (count < 10) {
        alpha = 0.4
      } else if (count < 15) {
        alpha = 0.6
      } else if (count < 25) {
        alpha = 0.8
      } else {
        alpha = 1
      }
      return 'rgba(38, 42, 48, ' + alpha + ')'
    },
  }
}
</script>

<style scoped>
.ca {
  display: inline-block;
  margin: 10px;
  padding: 5px 10px;
  color: #fff;
  border-radius: 5px;
  text-decoration: none;
  outline: none;
  word-wrap: break-word;
  transition: all 0.3s;
}
</style>
