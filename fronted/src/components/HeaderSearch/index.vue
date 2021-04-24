<template>
  <div :class="{ show: show }" class="header-search">
    <svg-icon
      class-name="search-icon"
      icon-class="search"
      @click.stop="click"
    />
    <el-input
      ref="headerSearch"
      v-model="search"
      type="search"
      placeholder="输入关键字，然后回车"
      class="header-search-select"
      @keyup.enter.native="toSearch"
    />
  </div>
</template>

<script>
export default {
  name: 'HeaderSearch',
  data() {
    return {
      search: '',
      show: false
    }
  },
  watch: {
    show(value) {
      if (value) {
        document.body.addEventListener('click', this.close)
      } else {
        document.body.removeEventListener('click', this.close)
      }
    }
  },
  methods: {
    click() {
      this.show = !this.show
      if (this.show) {
        this.$refs.headerSearch && this.$refs.headerSearch.focus()
      }
    },
    close() {
      this.$refs.headerSearch && this.$refs.headerSearch.blur()
      this.show = false
    },
    toSearch() {
      if (this.search === '') {
        this.$message.error('请正确填写关键字')
        return
      } else {
        this.$router.push({ name: 'CTQArticle', query: { q: this.search }})
        this.close()
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.search-icon {
  color: #909399;
}
.search-icon:hover {
  color: #fff;
}
.header-search {
  font-size: 0 !important;

  .search-icon {
    cursor: pointer;
    font-size: 18px;
    vertical-align: middle;
  }

  .header-search-select {
    font-size: 18px;
    transition: width 0.2s;
    width: 0;
    overflow: hidden;
    background: transparent;
    border-radius: 0;
    display: inline-block;
    vertical-align: middle;

    /deep/ .el-input__inner {
      border-radius: 0;
      border: 0;
      padding-left: 0;
      padding-right: 0;
      box-shadow: none !important;
      border-bottom: 1px solid #d9d9d9;
      vertical-align: middle;
    }
  }

  &.show {
    .header-search-select {
      width: 210px;
      margin-left: 10px;
    }
  }
}
</style>
