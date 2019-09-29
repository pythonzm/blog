<template>
  <div id="app">
    <router-view />

    <back-to-top
      :custom-style="myBackToTopStyle"
      :visibility-height="300"
      :back-position="20"
      transition-name="fade"
    />
  </div>
</template>

<script>
import BackToTop from '@/components/BackToTop'
export default {
  name: 'App',
  components: {
    BackToTop
  },
  data () {
    return {
      myBackToTopStyle: {
        'z-index': 10,
        right: '50px',
        bottom: '80px',
        width: '40px',
        height: '40px',
        'border-radius': '4px',
        'line-height': '45px', // 请保持与高度一致以垂直居中 Please keep consistent with height to center vertically
        background: '#000'// 按钮的背景颜色 The background color of the button
      }
    }
  },
}
</script>
