<template>
  <el-table :data="list" style="width: 100%;padding-top: 15px;">
    <el-table-column label="文章名称" align="center">
      <template slot-scope="scope">
        {{ scope.row.title }}
      </template>
    </el-table-column>
    <el-table-column label="文章评论" align="center">
      <template slot-scope="scope">
        {{ scope.row.content }}
      </template>
    </el-table-column>
    <el-table-column label="评论时间" align="center">
      <template slot-scope="scope">
        {{ scope.row.create_time }}
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import { fetchRecentComments } from "@/api/comment";

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
      fetchRecentComments().then(response => {
        this.list = response.data;
      });
    }
  }
};
</script>
