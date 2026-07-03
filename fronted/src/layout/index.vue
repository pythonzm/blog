<template>
  <div :class="classObj" class="app-wrapper">
    <div v-if="device==='mobile'&&sidebar.opened" class="drawer-bg" @click="handleClickOutside" />
    <sidebar class="sidebar-container" />
    <div class="main-container">
      <div :class="{'fixed-header':fixedHeader}">
        <navbar />
      </div>
      <app-main />
    </div>
  </div>
</template>

<script>
import { Navbar, Sidebar, AppMain } from './components'
import ResizeMixin from './mixin/ResizeHandler'

export default {
  name: 'Layout',
  components: {
    Navbar,
    Sidebar,
    AppMain
  },
  mixins: [ResizeMixin],
  computed: {
    sidebar() {
      return this.$store.state.app.sidebar
    },
    device() {
      return this.$store.state.app.device
    },
    fixedHeader() {
      return this.$store.state.settings.fixedHeader
    },
    classObj() {
      return {
        hideSidebar: !this.sidebar.opened,
        openSidebar: this.sidebar.opened,
        withoutAnimation: this.sidebar.withoutAnimation,
        mobile: this.device === 'mobile'
      }
    }
  },
  methods: {
    handleClickOutside() {
      this.$store.dispatch('app/closeSideBar', { withoutAnimation: false })
    }
  }
}
</script>

<style lang="scss" scoped>
  @use "~@/styles/mixin";
  @use "~@/styles/variables";

  .app-wrapper {
    @include mixin.clearfix;
    position: relative;
    height: 100%;
    width: 100%;
    &.mobile.openSidebar{
      position: fixed;
      top: 0;
    }
  }
  .drawer-bg {
    background: #000;
    opacity: 0.3;
    width: 100%;
    top: 0;
    height: 100%;
    position: absolute;
    z-index: 999;
  }

  .fixed-header {
    position: fixed;
    top: 0;
    right: 0;
    z-index: 9;
    width: calc(100% - #{variables.$sideBarWidth});
    transition: width 0.28s;
  }

  .hideSidebar .fixed-header {
    width: calc(100% - 54px)
  }

  .mobile .fixed-header {
    width: 100%;
  }
</style>

<style lang="scss">
/* Admin UI Charcoal/Slate Highlight Overrides */
.app-wrapper {
  background-color: transparent !important;
  color: inherit;

  .main-container {
    background-color: #f0f2f5 !important;
    min-height: 100vh;
  }

  .app-main {
    background-color: #f0f2f5 !important;
    padding: 20px;
  }

  .navbar {
    background: #ffffff !important;
    border-bottom: 1px solid #e6e6e6 !important;

    .hamburger-container {
      svg {
        fill: #5a5e66 !important;
      }
    }
  }

  // Cards
  .el-card {
    background-color: #ffffff !important;
    border: 1px solid #ebeef5 !important;
    color: #303133 !important;

    .el-card__header {
      border-bottom: 1px solid #ebeef5 !important;
      color: #303133 !important;
    }
  }

  // Tables
  .el-table {
    background-color: #ffffff !important;
    color: #606266 !important;
    border: 1px solid #ebeef5 !important;

    th, tr {
      background-color: #ffffff !important;
      color: #909399 !important;
    }

    td {
      border-bottom: 1px solid #ebeef5 !important;
    }

    th.is-leaf {
      border-bottom: 1px solid #ebeef5 !important;
    }

    &::before {
      background-color: #ebeef5 !important;
    }

    .el-table__body tr:hover > td {
      background-color: #f5f7fa !important;
    }
  }

  // Primary Buttons
  .el-button--primary {
    background-color: #1e293b !important;
    border-color: #1e293b !important;
    color: #ffffff !important;

    &:hover, &:focus {
      background-color: #0f172a !important;
      border-color: #0f172a !important;
      color: #ffffff !important;
    }

    &.is-active, &:active {
      background-color: #0f172a !important;
      border-color: #0f172a !important;
    }
  }

  // Text Buttons
  .el-button--text {
    color: #1e293b !important;

    &:hover, &:focus {
      color: #0f172a !important;
    }
  }

  // Form input focus states
  .el-input__inner:focus,
  .el-textarea__inner:focus {
    border-color: #1e293b !important;
  }

  // Checkboxes
  .el-checkbox__input.is-checked {
    .el-checkbox__inner {
      background-color: #1e293b !important;
      border-color: #1e293b !important;
    }

    + .el-checkbox__label {
      color: #1e293b !important;
    }
  }

  .el-checkbox__input.is-indeterminate .el-checkbox__inner {
    background-color: #1e293b !important;
    border-color: #1e293b !important;
  }

  // Radio buttons
  .el-radio__input.is-checked {
    .el-radio__inner {
      border-color: #1e293b !important;
      background: #1e293b !important;
    }

    + .el-radio__label {
      color: #1e293b !important;
    }
  }

  // Switch component
  .el-switch.is-checked .el-switch__core {
    border-color: #1e293b !important;
    background-color: #1e293b !important;
  }

  // Dialogs
  .el-dialog {
    background-color: #ffffff !important;
    border: 1px solid #ebeef5 !important;

    .el-dialog__title {
      color: #303133 !important;
    }

    .el-dialog__body {
      color: #606266 !important;
    }
  }

  // Breadcrumb
  .el-breadcrumb__inner a,
  .el-breadcrumb__inner.is-link {
    color: #303133 !important;

    &:hover {
      color: #1e293b !important;
    }
  }

  // Pagination (Active page item)
  .pagination-container {
    background: #ffffff !important;
  }

  .el-pagination.is-background .el-pager li:not(.disabled).active {
    background-color: #1e293b !important;
    border-color: #1e293b !important;
    color: #ffffff !important;
  }

  .el-pagination.is-background .el-pager li:not(.disabled):hover {
    color: #1e293b !important;
  }

  // Tabs active states
  .el-tabs__item {
    &:hover {
      color: #1e293b !important;
    }

    &.is-active {
      color: #1e293b !important;
    }
  }

  .el-tabs__active-bar {
    background-color: #1e293b !important;
  }

  // Tag badges
  .el-tag {
    background-color: #f1f5f9 !important;
    border-color: #e2e8f0 !important;
    color: #475569 !important;

    &.el-tag--success {
      background-color: rgba(16, 185, 129, 0.1) !important;
      border-color: rgba(16, 185, 129, 0.2) !important;
      color: #10b981 !important;
    }

    &.el-tag--danger {
      background-color: rgba(239, 68, 68, 0.1) !important;
      border-color: rgba(239, 68, 68, 0.2) !important;
      color: #ef4444 !important;
    }

    &.el-tag--warning {
      background-color: rgba(245, 158, 11, 0.1) !important;
      border-color: rgba(245, 158, 11, 0.2) !important;
      color: #f59e0b !important;
    }
  }

  // Dashboard Widget elements
  .panel-group {
    .card-panel {
      background: #ffffff !important;
      border: 1px solid #ebeef5 !important;
      color: #666 !important;

      &:hover {
        border-color: #1e293b !important;

        .card-panel-icon-wrapper {
          color: #fff !important;

          &.icon-people { background: #475569 !important; }
          &.icon-message { background: #475569 !important; }
          &.icon-money { background: #475569 !important; }
          &.icon-shopping { background: #475569 !important; }
        }
      }

      .card-panel-icon-wrapper {
        transition: all 0.38s ease-out;

        &.icon-people { color: #475569 !important; }
        &.icon-message { color: #475569 !important; }
        &.icon-money { color: #475569 !important; }
        &.icon-shopping { color: #475569 !important; }
      }
    }
  }
}

// Select dropdown items active & Dropdown menus
.el-select-dropdown,
.el-dropdown-menu {
  background-color: #ffffff !important;
  border: 1px solid #e4e7ed !important;

  .el-select-dropdown__item,
  .el-dropdown-menu__item {
    color: #606266 !important;

    &.hover, &:hover {
      background-color: #f5f7fa !important;
      color: #1e293b !important;
    }

    &.selected {
      color: #1e293b !important;
      font-weight: 700;
    }
  }
}
</style>
