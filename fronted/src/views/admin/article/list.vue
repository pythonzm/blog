<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input
        v-model="listQuery.key"
        placeholder="标题"
        style="width: 200px;"
        class="filter-item"
      />
      <el-select
        v-model="listQuery.status"
        placeholder="状态"
        clearable
        class="filter-item"
        style="width: 130px"
      >
        <el-option
          v-for="item in statusOptions"
          :key="item.key"
          :label="item.display_name"
          :value="item.key"
        />
      </el-select>
      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="getList"
        >搜索</el-button
      >
      <router-link :to="{ name: 'CreateArticle' }">
        <el-button
          class="filter-item"
          style="margin-left: 10px;"
          type="primary"
          icon="el-icon-edit"
          >添加文章</el-button
        >
      </router-link>
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
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" width="auto" align="center">
        <template v-slot="scope">
          <span>{{ scope.row.created_time }}</span>
        </template>
      </el-table-column>
      <el-table-column label="更新时间" width="auto" align="center">
        <template v-slot="scope">
          <span>{{ scope.row.updated_time }}</span>
        </template>
      </el-table-column>
      <el-table-column label="标题" min-width="auto">
        <template v-slot="{ row }">
          <span class="link-type">{{ row.title }}</span>
        </template>
      </el-table-column>
      <el-table-column label="状态" class-name="status-col" width="auto">
        <template v-slot="{ row }">
          <el-tag :type="row.status | statusFilter">{{
            row.status | statusShowFilter
          }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        label="操作"
        align="center"
        width="auto"
        class-name="small-padding fixed-width"
      >
        <template v-slot="scope">
          <router-link
            :to="{ name: 'EditArticle', params: { id: scope.row.id } }"
          >
            <el-button type="primary" size="mini">编辑</el-button>
          </router-link>
          <el-button
            size="mini"
            type="danger"
            @click.native.prevent="deleteRow(scope.$index, list, scope.row.id)"
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
import { fetchList, deleteArticle } from '@/api/article'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination

const statusOptions = [
  { key: 'published', display_name: '已发布' },
  { key: 'draft', display_name: '草稿' }
]

export default {
  name: 'ArticleList',
  components: { Pagination },
  filters: {
    statusFilter (status) {
      const statusMap = {
        published: 'success',
        draft: 'info'
      }
      return statusMap[status]
    },
    statusShowFilter (status) {
      const statusMap = {
        published: '已发布',
        draft: '草稿'
      }
      return statusMap[status]
    }
  },
  data () {
    return {
      list: null,
      total: 0,
      listLoading: true,
      statusOptions,
      listQuery: {
        page: 1,
        limit: 10,
        key: undefined,
        status: undefined,
        admin: true
      }
    }
  },
  created () {
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
    deleteRow (index, rows, id) {
      this.$confirm('此操作将永久删除该文章, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteArticle(id).then(response => {
          this.$message({
            type: 'success',
            message: '删除成功!'
          })
        })
        rows.splice(index, 1)
      })
    }
  }
}
</script>
