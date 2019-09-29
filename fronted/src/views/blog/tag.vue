<template>
  <div>
    <div class="ca" v-for="tag in list" :key="tag.id">
      <mallki class-name="mallki-text" :text="tag.tag_name" />
    </div>
  </div>
</template>

<script>
import { fetchTagList } from '@/api/tag'
import Mallki from '@/components/TextHoverEffect/Mallki'
export default {
  name: "Tag",
  components: {
    Mallki
  },
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
  border-bottom: 1px solid #999;
  padding-bottom: 10px;
}
</style>