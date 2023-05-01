<template>
  <div class="app-container">
    <el-input
      v-model="filterText"
      placeholder="输入关键字进行过滤"
      style="margin-bottom: 20px;"
    />

    <div class="main">
      <el-tree
        ref="tree"
        class="filter-tree"
        node-key="label"
        :data="data"
        :props="defaultProps"
        accordion
        draggable
        :expand-on-click-node="false"
        :filter-node-method="filterNode"
      >
        <span slot-scope="{ node, data }" class="custom-tree-node">
          <span>{{ node.label }}</span>
          <span style="padding-left: 15px;">
            <el-button
              type="text"
              size="mini"
              icon="el-icon-plus"
              @click="addSubNode(data)"
            />
            <el-button
              type="text"
              size="mini"
              icon="el-icon-edit"
              @click="editNode(data)"
            />
            <el-button
              type="text"
              size="mini"
              icon="el-icon-delete"
              @click="() => remove(node, data)"
            />
          </span>
        </span>
      </el-tree>
      <div class="buttons">
        <el-button type="primary" class="sub-button" @click="addMainNode">添加主节点</el-button>
        <el-button type="success" class="sub-button" @click="submit">发布</el-button>
      </div>
    </div>

    <el-dialog :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :model="form">
        <el-form-item label="节点名称" label-width="100px">
          <el-input v-model="form.label" autocomplete="off" />
        </el-form-item>
        <el-form-item v-if="!addMainNodeFlag" label="访问地址" label-width="100px">
          <el-input v-model="form.addr" autocomplete="off" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="handleClick">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { fetchCollection, cudCollection } from '@/api/collection'
export default {
  data() {
    return {
      filterText: '',
      data: [],
      defaultProps: {
        children: 'children',
        label: 'label'
      },
      dialogFormVisible: false,
      form: {
        label: '',
        addr: ''
      },
      addMainNodeFlag: false,
      subNodeData: '',
      isEditNode: false,
      editData: {}
    }
  },
  computed: {
    dialogTitle() {
      return this.addMainNodeFlag ? '添加主节点' : '添加子节点'
    }
  },
  watch: {
    filterText(val) {
      this.$refs.tree.filter(val)
    }
  },
  created() {
    this.getCollection()
  },
  methods: {
    filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    },
    append(data) {
      const newChild = { label: this.form.label, addr: this.form.addr, children: [] }
      if (!data.children) {
        this.$set(data, 'children', [])
      }
      data.children.push(newChild)
    },

    remove(node, data) {
      this.$confirm(`确定要删除节点 ${data.label} 吗？`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        const parent = node.parent
        const children = parent.data.children || parent.data
        const index = children.findIndex(d => d.label === data.label)
        children.splice(index, 1)
        this.$message({
          type: 'success',
          message: '删除成功!'
        })
      }).catch(() => {
        //
      })
    },
    addSubNode(data) {
      this.dialogFormVisible = true
      this.addMainNodeFlag = false
      this.isEditNode = false
      this.subNodeData = data
    },
    addMainNode() {
      this.dialogFormVisible = true
      this.addMainNodeFlag = true
      this.isEditNode = false
    },
    addNode() {
      if (this.addMainNodeFlag) {
        this.data.push({
          label: this.form.label,
          children: []
        })
      } else {
        this.append(this.subNodeData)
      }
      this.dialogFormVisible = false
    },
    editNode(data) {
      this.isEditNode = true
      this.dialogFormVisible = true
      this.addMainNodeFlag = !data.addr
      this.editData = data
      this.form.label = data.label
      this.form.addr = data.addr
    },
    submitEdit() {
      this.$set(this.editData, 'label', this.form.label)
      if (!this.addMainNodeFlag) {
        this.$set(this.editData, 'addr', this.form.addr)
      }

      this.dialogFormVisible = false
    },
    handleClick() {
      if (this.isEditNode) {
        this.submitEdit()
      } else {
        this.addNode()
      }
    },
    getCollection() {
      fetchCollection().then(res => {
        this.data = res.data.collection
      })
    },
    submit() {
      cudCollection({ collection: this.data }).then(res => {
        this.$message({
          type: 'success',
          message: '发布成功!'
        })
      })
    }
  }
}
</script>

<style scoped>
  .custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 14px;
    padding-right: 8px;
  }
  .main {
    display: flex;
  }
  .el-tree {
    flex: 1;
  }
  .buttons {
    flex: 2;
    display: flex;
    justify-content: center;
  }
  .sub-button {
    height: 50px;
  }
</style>
