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
          <span>
            <el-button
              type="text"
              size="mini"
              @click="addSubNode(data)"
            >
              添加子节点
            </el-button>
            <el-button
              type="text"
              size="mini"
              @click="editNode(data)"
            >
              编辑节点
            </el-button>
            <el-button
              type="text"
              size="mini"
              @click="() => remove(node, data)"
            >
              删除节点
            </el-button>
          </span>
        </span>
      </el-tree>
      <div class="buttons">
        <el-button type="primary" class="sub-button" @click="addMainNode">添加主节点</el-button>
        <el-button type="success" class="sub-button">发布</el-button>
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
        <el-button type="primary" @click="addNode">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      filterText: '',
      data: [{
        label: 'Linux',
        children: [{
          label: 'Nginx',
          children: [{
            label: '官网',
            addr: 'https://www.baidu.com'
          }, {
            label: 'org',
            addr: 'https://cn.bing.com'
          }]
        }]
      }, {
        id: 2,
        label: '一级 2',
        children: [{
          id: 5,
          label: '二级 2-1'
        }, {
          id: 6,
          label: '二级 2-2'
        }]
      }, {
        id: 3,
        label: '一级 3',
        children: [{
          id: 7,
          label: '二级 3-1'
        }, {
          id: 8,
          label: '二级 3-2'
        }]
      }],
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
      subNodeData: ''
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
  methods: {
    filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    },
    append(data) {
      const newChild = { label: this.form.label, children: [] }
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
      this.subNodeData = data
    },
    addMainNode() {
      this.dialogFormVisible = true
      this.addMainNodeFlag = true
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
      this.dialogFormVisible = true
      this.addMainNodeFlag = !data.addr
      this.form.label = data.label
      this.form.addr = data.addr
      this.form = { ...this.form, children: data.children ? data.children : [] }
      this.$set(data, this.form)
    }
  }
}
</script>

<style>
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
