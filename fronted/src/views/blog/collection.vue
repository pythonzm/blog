<template>
  <div class="collection-page-container">
    <div class="page-header">
      <h2 class="page-title">
        <i class="el-icon-star-off page-icon" />
        藏宝阁
      </h2>
      <p class="page-subtitle">个人珍藏与高质量技术文章资源导航</p>
    </div>

    <div class="filter-wrapper">
      <el-input
        v-model="filterText"
        class="search-input"
        placeholder="搜索收藏夹或链接名称..."
        prefix-icon="el-icon-search"
        clearable
      />
    </div>

    <div class="tree-wrapper">
      <el-tree
        ref="tree"
        class="filter-tree custom-el-tree"
        :data="treeData"
        :props="defaultProps"
        :filter-node-method="filterNode"
      >
        <span slot-scope="{ node, data: nodeData }" class="custom-tree-node">
          <span class="node-content">
            <template v-if="nodeData.children && Object.keys(nodeData.children).length > 0">
              <i class="el-icon-folder folder-icon" :class="{ 'active': node.expanded }" />
              <span class="folder-label">{{ node.label }}</span>
            </template>
            <template v-else>
              <i class="el-icon-link link-icon" />
              <el-button class="link-button" type="text" @click="goToUrl(nodeData.addr)">
                {{ node.label }}
              </el-button>
              <span v-if="device !== 'mobile'" class="node-address">{{ formatUrl(nodeData.addr) }}</span>
            </template>
          </span>
        </span>
      </el-tree>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { fetchCollection } from '@/api/collection'
import ResizeMixin from '@/blayout/mixin/ResizeHandler'
export default {
  name: 'Collection',
  mixins: [ResizeMixin],
  data() {
    return {
      filterText: '',
      treeData: [],
      defaultProps: {
        children: 'children',
        label: 'label'
      }
    }
  },
  computed: {
    ...mapGetters([
      'device'
    ])
  },
  watch: {
    filterText(val) {
      this.$refs.tree.filter(val)
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      fetchCollection().then(response => {
        this.treeData = response.data.collection || []
      })
    },
    filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    },
    goToUrl(address) {
      if (address) {
        window.open(address, '_blank')
      }
    },
    formatUrl(url) {
      if (!url) return ''
      try {
        const domain = new URL(url).hostname
        return domain
      } catch (e) {
        return url
      }
    }
  }
}
</script>

<style lang="scss">
// 覆盖Element UI Tree样式
.custom-el-tree {
  background: transparent !important;

  .el-tree-node {
    margin-bottom: 6px;

    &:focus > .el-tree-node__content {
      background-color: transparent !important;
    }
  }

  .el-tree-node__content {
    height: 40px !important;
    border-radius: 8px;
    padding: 0 8px;
    transition: all 0.2s ease;
    border: 1px solid transparent;

    &:hover {
      background-color: #f1f5f9 !important;
      border-color: #e2e8f0;
    }
  }

  .el-tree-node__expand-icon {
    font-size: 14px;
    color: #94a3b8;

    &.expanded {
      transform: rotate(90deg);
      color: #1e293b;
    }

    &.is-leaf {
      color: transparent;
      cursor: default;
    }
  }

  // 过滤输入框样式微调
  .search-input .el-input__inner {
    height: 42px;
    line-height: 42px;
    border-radius: 20px;
    border-color: #e2e8f0;
    background-color: #ffffff !important;
    color: #475569 !important;
    padding-left: 40px;
    transition: all 0.3s ease;

    &:focus {
      border-color: #1e293b;
      box-shadow: 0 0 0 3px rgba(30, 41, 59, 0.08);
    }
  }
}
</style>

<style scoped lang="scss">
.collection-page-container {
  width: 100%;
  background: #ffffff;
  border-radius: 12px;
  border: 1px solid rgba(226, 232, 240, 0.8);
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
  padding: 40px;
  box-sizing: border-box;
  text-align: left;
}

.page-header {
  border-bottom: 1px solid #f1f5f9;
  padding-bottom: 20px;
  margin-bottom: 30px;
}

.page-title {
  font-size: 24px;
  font-weight: 800;
  color: #0f172a;
  margin: 0 0 6px 0;
  display: flex;
  align-items: center;
  gap: 8px;

  .page-icon {
    color: #1e293b;
  }
}

.page-subtitle {
  font-size: 14px;
  color: #64748b;
  margin: 0;
}

.filter-wrapper {
  margin-bottom: 24px;
  max-width: 480px;
}

.tree-wrapper {
  background-color: #ffffff;
  border: 1px solid #f1f5f9;
  border-radius: 12px;
  padding: 16px;
}

.custom-tree-node {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  padding-right: 12px;
  font-size: 14px;
  overflow: hidden;
}

.node-content {
  display: flex;
  align-items: center;
  gap: 8px;
  min-width: 0;
}

.folder-icon {
  font-size: 16px;
  color: #64748b;
  transition: color 0.2s ease;

  &.active {
    color: #1e293b;
  }
}

.link-icon {
  font-size: 14px;
  color: #94a3b8;
}

.folder-label {
  font-weight: 600;
  color: #334155;
}

.link-button {
  color: #475569;
  font-weight: 500;
  padding: 0;
  transition: color 0.2s ease;
  text-align: left;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;

  &:hover {
    color: #1e293b;
  }

  &:focus {
    color: #1e293b;
  }
}

.node-address {
  font-size: 12px;
  color: #94a3b8;
  white-space: nowrap;
  margin-left: 12px;
  background-color: #f8fafc;
  padding: 2px 8px;
  border-radius: 4px;
  border: 1px solid #f1f5f9;
}

[v-cloak] {
  display: none;
}

@media (max-width: 768px) {
  .collection-page-container {
    padding: 20px;
  }

  .node-address {
    display: none;
  }
}
</style>
