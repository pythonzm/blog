<template>
  <div class="header-container">
    <!-- 移动端汉堡菜单触发器 -->
    <div v-if="mobile" class="mobile-menu-trigger" @click="drawerVisible = true">
      <i class="el-icon-menu" />
    </div>

    <!-- Logo (桌面端靠左，移动端居中) -->
    <div class="logo" :class="{ 'mobile-logo': mobile }">
      <router-link to="/">{{ logo }}</router-link>
    </div>

    <!-- 桌面端导航菜单 -->
    <div v-if="!mobile" class="nav-wrapper">
      <el-menu
        :default-active="activeIndex"
        mode="horizontal"
        active-text-color="#ffffff"
        router
        class="nav-menu"
      >
        <el-menu-item
          v-for="(item, key) in navOptions"
          :key="key"
          :index="item.index"
          class="custom-menu-item"
        >{{ item.label }}</el-menu-item>

        <el-menu-item class="search-item">
          <svg-icon
            v-if="algoliaSearch"
            class-name="search-icon"
            icon-class="search"
            @click="dialogVisible = true"
          />
          <HeaderSearch v-else id="header-search" />
        </el-menu-item>
      </el-menu>
    </div>

    <!-- 移动端搜索触发器 (靠右) -->
    <div v-if="mobile" class="mobile-search-trigger">
      <svg-icon
        v-if="algoliaSearch"
        class-name="search-icon"
        icon-class="search"
        @click="dialogVisible = true"
      />
      <HeaderSearch v-else id="header-search" />
    </div>

    <!-- 全局搜索弹窗 (Algolia) -->
    <el-dialog
      :visible.sync="dialogVisible"
      append-to-body
      :show-close="showClose"
      :fullscreen="mobile"
      @opened="focusAlgolia"
      @closed="blurAlgolia"
    >
      <div v-if="mobile" class="mobile-close-btn" @click="dialogVisible = false">
        <i class="el-icon-close" />
      </div>
      <AlgoliaSearch class="el-dialog-div" />
    </el-dialog>

    <!-- 移动端侧边抽屉菜单 -->
    <el-drawer
      title="导航菜单"
      :visible.sync="drawerVisible"
      direction="ltr"
      size="240px"
      append-to-body
      custom-class="mobile-nav-drawer"
    >
      <div class="drawer-menu-list">
        <router-link
          v-for="(item, key) in navOptions"
          :key="key"
          :to="item.index"
          class="drawer-menu-item"
          :class="{ 'active': activeIndex === item.index }"
          @click.native="drawerVisible = false"
        >
          <i :class="getMenuIcon(item.label)" class="menu-icon" />
          {{ item.label }}
        </router-link>
      </div>
    </el-drawer>
  </div>
</template>

<script>
import HeaderSearch from '@/components/HeaderSearch'
import AlgoliaSearch from '@/components/AlgoliaSearch'
import defaultSettings from '@/settings'

const { algoliaSearch } = defaultSettings
export default {
  name: 'Header',
  components: {
    AlgoliaSearch,
    HeaderSearch
  },
  props: {
    mobile: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      logo: 'PoorOPS',
      navOptions: [
        { label: '首页', index: '/' },
        { label: '分类', index: '/category' },
        { label: '标签', index: '/tag' },
        { label: '藏宝阁', index: '/collection' },
        { label: '关于', index: '/about' }
      ],
      dialogVisible: false,
      drawerVisible: false,
      showClose: false,
      algoliaSearch
    }
  },
  computed: {
    activeIndex() {
      const path = this.$route.path
      const names = ['category', 'tag', 'collection', 'about']
      for (const name of names) {
        if (path.indexOf(name) !== -1) {
          return `/${name}`
        }
      }
      return '/'
    }
  },
  methods: {
    focusAlgolia() {
      document.getElementsByClassName('ais-SearchBox-input')[0].focus()
    },
    blurAlgolia() {
      document.getElementsByClassName('ais-SearchBox-input')[0].blur()
    },
    getMenuIcon(label) {
      const icons = {
        '首页': 'el-icon-s-home',
        '分类': 'el-icon-folder-opened',
        '标签': 'el-icon-price-tag',
        '藏宝阁': 'el-icon-star-off',
        '关于': 'el-icon-user'
      }
      return icons[label] || 'el-icon-link'
    }
  }
}
</script>

<style lang="scss">
.header-search-select .el-input__inner {
  font-size: 13px;
  height: 32px;
  line-height: 32px;
  border-radius: 20px;
  background-color: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  color: #fff;
  transition: all 0.3s ease;

  &:focus {
    border-color: #ffffff;
    background-color: rgba(255, 255, 255, 0.15);
  }
}

.el-dialog__header {
  display: none;
}

.mobile-close-btn {
  position: absolute;
  top: 16px;
  right: 16px;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: #f1f5f9;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  color: #475569;
  z-index: 10000;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
  cursor: pointer;

  &:active {
    background-color: #e2e8f0;
    color: #0f172a;
  }
}

.el-dialog.is-fullscreen {
  .el-dialog__body {
    padding-top: 64px !important; /* 预留出顶部关闭按钮的空间，防止重叠 */
  }
}

// 覆写ElementUI的水平菜单样式
.header-container {
  .el-menu.el-menu--horizontal {
    border-bottom: none !important;
    background-color: transparent !important;
  }

  .el-menu--horizontal > .el-menu-item {
    color: #94a3b8 !important;
    border-bottom: 2px solid transparent !important;
    height: 64px;
    line-height: 64px;
    font-size: 15px;
    font-weight: 500;
    transition: all 0.3s ease;

    &:hover, &:focus {
      background-color: transparent !important;
      color: #ffffff !important;
    }

    &.is-active {
      color: #ffffff !important;
      border-bottom-color: #ffffff !important;
      font-weight: 700;
    }
  }

  .search-item {
    background-color: transparent !important;
    border-bottom: none !important;
    display: flex;
    align-items: center;

    &:hover {
      background-color: transparent !important;
    }
  }
}

// 移动端侧边抽屉样式覆盖
.mobile-nav-drawer {
  background-color: #ffffff !important;

  .el-drawer__header {
    color: #1e293b !important;
    border-bottom: 1px solid #f1f5f9;
    padding: 16px 20px;
    margin-bottom: 0;
    font-weight: 700;
    font-size: 16px;
  }

  .el-drawer__close-btn {
    font-size: 18px;
    color: #94a3b8;

    &:hover {
      color: #1e293b;
    }
  }
}
</style>

<style scoped lang="scss">
.header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  height: 64px;
  padding: 0 40px;
  box-sizing: border-box;
  background: rgba(15, 23, 42, 0.9);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.logo {
  font-size: 1.6rem;
  font-weight: 800;
  letter-spacing: 0.5px;
  flex-shrink: 0;

  a {
    color: #ffffff;
    text-decoration: none;
    transition: opacity 0.3s ease;

    &:hover {
      opacity: 0.9;
    }
  }

  &.mobile-logo {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
    font-size: 1.4rem;
  }
}

.mobile-menu-trigger {
  font-size: 22px;
  color: #cbd5e1;
  cursor: pointer;
  padding: 4px;
  transition: color 0.2s ease;

  &:hover {
    color: #ffffff;
  }
}

.mobile-search-trigger {
  display: flex;
  align-items: center;
  flex-shrink: 0;
}

.nav-wrapper {
  display: flex;
  align-items: center;
  flex-shrink: 0;
}

.nav-menu {
  display: flex;
  align-items: center;
}

.search-icon {
  font-size: 18px;
  color: #94a3b8;
  transition: color 0.3s ease;
  cursor: pointer;

  &:hover {
    color: #ffffff;
  }
}

.el-dialog-div {
  max-height: 70vh;
  overflow: auto;
}

// 抽屉内菜单列表
.drawer-menu-list {
  display: flex;
  flex-direction: column;
  padding: 20px 12px;
  gap: 8px;
}

.drawer-menu-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  border-radius: 8px;
  color: #475569;
  text-decoration: none;
  font-size: 15px;
  font-weight: 500;
  transition: all 0.2s ease;

  .menu-icon {
    font-size: 16px;
    color: #94a3b8;
    transition: color 0.2s ease;
  }

  &:hover {
    background-color: #f1f5f9;
    color: #0f172a;

    .menu-icon {
      color: #1e293b;
    }
  }

  &.active {
    background-color: rgba(30, 41, 59, 0.05);
    color: #0f172a;
    font-weight: 700;

    .menu-icon {
      color: #1e293b;
    }
  }
}

@media (max-width: 768px) {
  .header-container {
    padding: 0 16px;
  }
}
</style>
