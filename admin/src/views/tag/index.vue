<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button
        class="filter-item"
        style="margin-bottom: unset;"
        type="primary"
        icon="el-icon-edit"
        @click="createtag"
      >添加标签</el-button>
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

      <el-table-column align="center" label="标签名称" width="auto">
        <template v-slot="{row}">
          <template v-if="row.edit">
            <el-input v-model="row.tag_name" class="edit-input" size="small" />
            <el-button
              class="cancel-btn"
              size="small"
              icon="el-icon-refresh"
              type="warning"
              @click="cancelEdit(row)"
            >取消</el-button>
          </template>
          <span v-else>{{ row.tag_name }}</span>
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

    <el-dialog title="添加标签名称" :visible.sync="dialogFormVisible">
      <el-form
        :model="temp"
        label-position="left"
        label-width="70px"
        style="width: 400px; margin-left:50px;"
      >
        <el-form-item label="标签名称">
          <el-input v-model="temp.tag_name" required />
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
import { fetchTagList, createTag, editTag, deleteTag } from '@/api/tag'

export default {
  name: 'Tag',
  data() {
    return {
      list: null,
      listLoading: true,
      temp: {
        tag_name: ''
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
      const { data } = await fetchTagList()
      const items = data
      this.list = items.map(v => {
        this.$set(v, 'edit', false) // https://vuejs.org/v2/guide/reactivity.html
        v.originalName = v.tag_name //  will be used when user click the cancel botton
        return v
      })
      this.listLoading = false
    },
    cancelEdit(row) {
      row.tag_name = row.originalName
      row.edit = false
    },
    confirmEdit(row) {
      row.edit = false
      this.temp.tag_name = row.tag_name
      editTag(row.id, this.temp).then(response => {
        row.originalName = row.tag_name
        this.$message({
          message: '修改成功',
          type: 'success'
        })
      })
    },
    deleteRow(index, rows, id) {
      this.$confirm('此操作将永久删除该标签, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          deleteTag(id).then(response => {
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
        tag_name: ''
      }
    },
    createtag() {
      this.resetTemp()
      this.dialogFormVisible = true
    },
    createData() {
      createTag(this.temp).then(response => {
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
