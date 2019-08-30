<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button
        class="filter-item"
        style="margin-bottom: unset;"
        type="primary"
        icon="el-icon-edit"
        @click="createCategory"
      >添加分类</el-button>
    </div>

    <el-table
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%"
    >
      <el-table-column align="center" label="ID" width="80">
        <template v-slot="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="分类名称" width="auto">
        <template v-slot="{row}">
          <template v-if="row.edit">
            <el-input v-model="row.category_name" class="edit-input" size="small" />
            <el-button
              class="cancel-btn"
              size="small"
              icon="el-icon-refresh"
              type="warning"
              @click="cancelEdit(row)"
            >取消</el-button>
          </template>
          <span v-else>{{ row.category_name }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="操作" width="auto">
        <template v-slot="scope">
          <el-button
            v-if="scope.row.edit"
            type="success"
            size="small"
            icon="el-icon-circle-check-outline"
            @click="confirmEdit(scope.row)"
          >确认</el-button>
          <el-button
            v-else
            type="primary"
            size="small"
            icon="el-icon-edit"
            @click="scope.row.edit=!scope.row.edit"
          >编辑</el-button>
          <el-button
            size="small"
            type="danger"
            @click.native.prevent="deleteRow(scope.$index, list, scope.row.id)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog title="添加分类名称" :visible.sync="dialogFormVisible">
      <el-form
        :model="temp"
        label-position="left"
        label-width="70px"
        style="width: 400px; margin-left:50px;"
      >
        <el-form-item label="分类名称">
          <el-input v-model="temp.category_name" required />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="createData">确认</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  fetchCategoryList,
  createCategory,
  editCategory,
  deleteCategory
} from '@/api/category'

export default {
  name: 'Category',
  data() {
    return {
      list: null,
      listLoading: true,
      temp: {
        category_name: ''
      },
      dialogFormVisible: false
    }
  },
  created() {
    this.getList()
  },
  methods: {
    async getList() {
      this.listLoading = true
      const { data } = await fetchCategoryList()
      const items = data
      this.list = items.map(v => {
        this.$set(v, 'edit', false) // https://vuejs.org/v2/guide/reactivity.html
        v.originalName = v.category_name //  will be used when user click the cancel botton
        return v
      })
      this.listLoading = false
    },
    cancelEdit(row) {
      row.category_name = row.originalName
      row.edit = false
    },
    confirmEdit(row) {
      row.edit = false
      this.temp.category_name = row.category_name
      editCategory(row.id, this.temp).then(response => {
        row.originalName = row.category_name
        this.$message({
          message: '修改成功',
          type: 'success'
        })
      })
    },
    deleteRow(index, rows, id) {
      this.$confirm('此操作将永久删除该分类, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          deleteCategory(id).then(response => {
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
    resetTemp() {
      this.temp = {
        category_name: ''
      }
    },
    createCategory() {
      this.resetTemp()
      this.dialogFormVisible = true
    },
    createData() {
      createCategory(this.temp).then(response => {
        this.list.unshift(response.data)
        this.dialogFormVisible = false
        this.$notify({
          title: 'Success',
          message: '操作成功',
          type: 'success',
          duration: 2000
        })
      })
    }
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
