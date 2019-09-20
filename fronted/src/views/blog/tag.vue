<template>
  <div>
    <div class="ca" v-for="tag in list" :key="tag.id">
      <router-link
        :to="{
          name: 'CTArticle',
          query: { tag: tag.tag_name }
        }"
        >{{ tag.tag_name }}</router-link
      >
    </div>
  </div>
</template>

<script>
import { fetchTagList } from '@/api/tag'
export default {
  name: "Tag",
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
      const { data } = await fetchTagList()
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