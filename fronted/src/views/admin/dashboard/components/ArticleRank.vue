<template>
  <el-table :data="list" style="width: 100%;padding-top: 15px;">
    <el-table-column label="文章名称" min-width="200">
      <template slot-scope="scope">
        {{ scope.row.title }}
      </template>
    </el-table-column>
    <el-table-column label="访问量" align="center">
      <template slot-scope="scope">
        {{ scope.row.views }}
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import { fetchArticleRank } from "@/api/article";

export default {
  data() {
    return {
      list: null
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      fetchArticleRank().then(response => {
        this.list = response.data;
      });
    }
  }
};
</script>
