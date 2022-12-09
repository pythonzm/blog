<template>
  <div class="footer">
    <p>
      <span>© 2019-{{ thisYear }} <a href="/">PoorOPS</a></span>
    </p>
  </div>
</template>

<script>
export default {
  name: 'Footer',
  data() {
    return {
      thisYear: new Date().getFullYear()
    }
  }
}
</script>

<style scoped>
.footer {
  font-size: 14px;
  color: #999;
}
</style>
