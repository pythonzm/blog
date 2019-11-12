<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input
        v-model="listQuery.title"
        placeholder="评论文章"
        style="width: 400px;"
        class="filter-item"
      />
      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="getList"
        >搜索</el-button
      >
    </div>
    <el-table
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
    >
      <el-table-column label="ID" align="center" width="80">
        <template v-slot="scope">
          <span>{{ scope.row.comment.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="称呼" align="center" min-width="auto">
        <template v-slot="{ row }">
          <span class="link-type">{{ row.comment.username }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="评论内容" width="auto">
        <template v-slot="scope">
          <span style="white-space: pre-line;">{{
            scope.row.comment.content
          }}</span>
        </template>
      </el-table-column>
      <el-table-column label="评论文章" align="center" min-width="auto">
        <template v-slot="{ row }">
          <router-link
            :to="{ name: 'EditArticle', params: { id: row.article.id } }"
            class="link-type"
          >
            <span>{{ row.article.title }}</span>
          </router-link>
        </template>
      </el-table-column>
      <el-table-column label="是否是根评论" align="center" min-width="auto">
        <template v-slot="{ row }">
          <el-tag :type="row.comment.root_id.Valid | statusFilter">{{
            row.comment.root_id.Valid | isRootFilter
          }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="评论时间" width="auto" align="center">
        <template v-slot="scope">
          <span>{{ scope.row.comment.created_time }}</span>
        </template>
      </el-table-column>

      <el-table-column
        label="操作"
        align="center"
        width="auto"
        class-name="small-padding fixed-width"
      >
        <template v-slot="scope">
          <el-button
            size="mini"
            type="danger"
            @click.native.prevent="
              deleteRow(scope.$index, list, scope.row.comment.id)
            "
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <pagination
      v-show="total > 0"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="getList"
    />
  </div>
</template>

<script>
import { fetchAll, deleteComment } from '@/api/comment'
import Pagination from '@/components/Pagination'

export default {
  name: 'Comment',
  components: { Pagination },
  filters: {
    statusFilter (valid) {
      return valid ? 'info' : 'success'
    },
    isRootFilter (valid) {
      return valid ? '否' : '是'
    }
  },
  data () {
    return {
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10,
        title: undefined
      },
    }
  },
  created () {
    this.getList()
  },
  methods: {
    getList () {
      this.listLoading = true
      fetchAll(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.listLoading = false
      })
    },
    deleteRow (index, rows, id) {
      this.$confirm('此操作将永久删除该内容, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          deleteComment(id).then(response => {
            rows.splice(index, 1)
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
          })
        })
        .catch(() => {
          // pass
        })
    },
  }
}
</script>

<style scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
