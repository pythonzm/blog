<template>
  <div>
    <el-input
      v-model="filterText"
      placeholder="输入关键字进行过滤"
      style="margin-bottom: 20px;"
    />

    <el-tree
      ref="tree"
      class="filter-tree"
      :data="data"
      :props="defaultProps"
      accordion
      :filter-node-method="filterNode"
      icon-class=" "
    >
      <span slot-scope="{ node,data }" class="custom-tree-node">
        <svg-icon v-if="!data.children || Object.keys(data.children).length === 0" icon-class="internet" />
        <svg-icon v-else-if="node.expanded" icon-class="folder-open" />
        <svg-icon v-else icon-class="folder" />
        <el-button v-if="!data.children || Object.keys(data.children).length === 0" type="text" @click="jumpAddr(data)">{{ node.label }}</el-button>
        <span v-else>{{ node.label }}</span>
      </span>
    </el-tree>
  </div>
</template>

<script>
import { fetchCollection } from '@/api/collection'
export default {
  data() {
    return {
      filterText: '',
      data: [],
      defaultProps: {
        children: 'children',
        label: 'label'
      }
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
    jumpAddr(data) {
      window.open(data.addr, '_blank')
      console.log(this.data)
    },
    getCollection() {
      fetchCollection().then(res => {
        this.data = res.data.collection
      })
    }
  }
}
</script>

<style>
.el-input input.el-input__inner:focus {
  border-color: #000;
}
.el-button--text {
  color: #000;
}
.el-button--text:hover {
  color: #5c5a5a;
}

.el-button--text:focus {
  color: #5c5a5a;
}
</style>
