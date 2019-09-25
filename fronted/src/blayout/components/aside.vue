<template>
  <el-card>
    <div class="user-bio">
      <div class="user-education user-bio-section">
        <div class="user-bio-section-header">
          <svg-icon icon-class="education" /><span>每日一汤</span>
        </div>
        <div class="user-bio-section-body">
          <div class="text-muted">
            {{ soup.content }}
          </div>
        </div>
      </div>
    </div>
  </el-card>
</template>

<script>

export default {
  name: 'Aside',
  props: {
    soup: {
      type: Object,
      default: () => {
        return {
          content: '',
        }
      }
    }
  },
}
</script>

<style lang="scss" scoped>
.el-card {
  float: right;
  width: 25%;
  position: fixed;
  left: 70%;
}
.box-center {
  margin: 0 auto;
  display: table;
}

.text-muted {
  color: #777;
  white-space: pre-line;
  line-height: 1.5em;
}

.user-profile {
  .user-name {
    padding: 15px;
    font-size: 20px;
    font-weight: 700;
  }

  .box-center {
    padding-top: 10px;
  }
}

.user-bio {
  margin-top: 20px;
  color: #606266;

  span {
    padding-left: 4px;
  }

  .user-bio-section {
    font-size: 14px;
    padding: 15px 0;

    .user-bio-section-header {
      border-bottom: 1px solid #dfe6ec;
      padding-bottom: 10px;
      margin-bottom: 10px;
      font-weight: bold;
    }
  }
}
</style>
