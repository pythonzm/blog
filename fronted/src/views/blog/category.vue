<template>
  <div class="categories-page-container">
    <div class="page-header">
      <h2 class="page-title">
        <i class="el-icon-folder-opened page-icon" />
        分类归档
      </h2>
      <p class="page-subtitle">按主题分类索引的文章归档夹</p>
    </div>

    <div class="categories-grid">
      <div
        v-for="item in list"
        :key="item.id"
        class="category-card"
        @click="goToCategory(item.category_name)"
      >
        <div class="card-body">
          <div class="folder-icon-box">
            <i class="el-icon-folder" />
          </div>
          <div class="category-text-info">
            <h3 class="category-title">{{ item.category_name }}</h3>
            <span class="category-explore">
              探索分类 <i class="el-icon-arrow-right" />
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { fetchCategoryList } from '@/api/category'
export default {
  name: 'Category',
  data() {
    return {
      list: null,
      listLoading: true
    }
  },
  created() {
    this.getList()
  },
  methods: {
    async getList() {
      this.listLoading = true
      const { data } = await fetchCategoryList()
      this.list = data
      this.listLoading = false
    },
    goToCategory(name) {
      this.$router.push({
        name: 'CTQArticle',
        query: { category: name }
      })
    }
  }
}
</script>

<style scoped lang="scss">
.categories-page-container {
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

.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(240px, 1fr));
  gap: 20px;
}

.category-card {
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 20px;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);

  &:hover {
    background: #ffffff;
    border-color: #1e293b;
    transform: translateY(-2px);
    box-shadow: 0 10px 15px -3px rgba(30, 41, 59, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.02);

    .folder-icon-box {
      background: rgba(30, 41, 59, 0.05);
      color: #1e293b;
      transform: scale(1.05);
    }

    .category-title {
      color: #1e293b;
    }

    .category-explore {
      color: #0f172a;
    }
  }
}

.card-body {
  display: flex;
  align-items: center;
  gap: 16px;
}

.folder-icon-box {
  width: 46px;
  height: 46px;
  background: #f1f5f9;
  border-radius: 10px;
  color: #64748b;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  transition: all 0.3s ease;
  flex-shrink: 0;
}

.category-text-info {
  display: flex;
  flex-direction: column;
  min-width: 0;
}

.category-title {
  font-size: 16px;
  font-weight: 700;
  color: #1e293b;
  margin: 0 0 4px 0;
  white-space: nowrap;
  overflow: hidden;
  text-transform: capitalize;
  text-overflow: ellipsis;
  transition: color 0.3s ease;
}

.category-explore {
  font-size: 12px;
  color: #94a3b8;
  display: inline-flex;
  align-items: center;
  gap: 2px;
  font-weight: 500;
  transition: color 0.3s ease;
}

@media (max-width: 768px) {
  .categories-page-container {
    padding: 20px;
  }

  .categories-grid {
    grid-template-columns: 1fr;
  }
}
</style>
