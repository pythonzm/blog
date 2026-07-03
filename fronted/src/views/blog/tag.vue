<template>
  <div v-loading="listLoading" class="tags-page-container">
    <div class="page-header">
      <h2 class="page-title">
        <i class="el-icon-price-tag page-icon" />
        标签云
      </h2>
      <p class="page-subtitle">点击下方任意标签，快速筛选并阅读相关技术文章</p>
    </div>

    <div class="tags-cloud">
      <div
        v-for="(tag, index) in list"
        :key="tag.id"
        class="tag-chip"
        :style="getTagStyle(index)"
        @click="goToTag(tag.tag_name)"
      >
        <i class="el-icon-price-tag tag-icon-mini" />
        <span class="tag-name">{{ tag.tag_name }}</span>
      </div>
    </div>
  </div>
</template>

<script>
import { fetchTagList } from '@/api/tag'
export default {
  name: 'Tag',
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
      try {
        const { data } = await fetchTagList()
        this.list = data
      } catch (err) {
        console.error(err)
      }
      this.listLoading = false
    },
    goToTag(name) {
      this.$router.push({
        name: 'CTQArticle',
        query: { tag: name }
      })
    },
    getTagStyle(index) {
      // 预设六种高档暗灰/石板/碳黑色调的配色方案，符合“极简暗色特效”设计
      const schemes = [
        { color: '#1e293b', bg: '#f8fafc', border: '#e2e8f0', shadow: 'rgba(30, 41, 59, 0.05)' }, // 经典暗灰
        { color: '#334155', bg: '#f1f5f9', border: '#cbd5e1', shadow: 'rgba(51, 65, 85, 0.05)' }, // 石板灰
        { color: '#475569', bg: '#f8fafc', border: '#e2e8f0', shadow: 'rgba(71, 85, 105, 0.05)' }, // 钢灰色
        { color: '#0f172a', bg: '#f1f5f9', border: '#cbd5e1', shadow: 'rgba(15, 23, 42, 0.05)' }, // 深碳黑
        { color: '#374151', bg: '#f3f4f6', border: '#e5e7eb', shadow: 'rgba(55, 65, 81, 0.05)' }, // 炭灰
        { color: '#1f2937', bg: '#fafafa', border: '#e2e8f0', shadow: 'rgba(31, 41, 55, 0.05)' } // 雾夜黑
      ]
      const scheme = schemes[index % schemes.length]
      return {
        '--tag-color': scheme.color,
        '--tag-bg': scheme.bg,
        '--tag-border-color': scheme.border,
        '--tag-shadow-color': scheme.shadow
      }
    }
  }
}
</script>

<style scoped lang="scss">
.tags-page-container {
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
  margin-bottom: 35px;
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

.tags-cloud {
  display: flex;
  flex-wrap: wrap;
  gap: 14px;
  padding: 10px 0;
}

.tag-chip {
  display: inline-flex;
  align-items: center;
  color: var(--tag-color);
  background-color: var(--tag-bg);
  border: 1px solid var(--tag-border-color);
  border-radius: 24px;
  padding: 10px 22px;
  cursor: pointer;
  box-shadow: 0 2px 4px var(--tag-shadow-color);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  font-size: 14.5px;
  font-weight: 600;

  .tag-icon-mini {
    margin-right: 6px;
    font-size: 13px;
    opacity: 0.85;
    transition: transform 0.3s ease;
  }

  .tag-name {
    letter-spacing: 0.2px;
  }

  &:hover {
    transform: translateY(-2px) scale(1.02);
    box-shadow: 0 8px 16px rgba(30, 41, 59, 0.12);
    border-color: var(--tag-color);
    background-color: #ffffff;

    .tag-icon-mini {
      transform: rotate(15deg) scale(1.1);
    }
  }
}

@media (max-width: 768px) {
  .tags-page-container {
    padding: 25px 20px;
  }

  .tag-chip {
    padding: 8px 16px;
    font-size: 13.5px;
  }
}
</style>
