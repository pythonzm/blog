<template>
  <div class="dashboard-container">
    <div class="md-body" style="text-align:left" v-html="about" />
  </div>
</template>

<script>
import { getAbout } from '@/api/user'
import marked from 'marked'
import '@/assets/md.css'
export default {
  name: 'About',
  data () {
    return {
      about: ''
    }
  },
  created () {
    marked.setOptions({
      renderer: new marked.Renderer(),
      gfm: true,
      tables: true,
      breaks: false,
      pedantic: false,
      sanitize: false,
      smartLists: true,
      smartypants: false
    })
    this.fetchData()
  },
  methods: {
    fetchData () {
      getAbout()
        .then(response => {
          this.about = marked(response.data)
        })
        .catch(err => {
          console.log(err)
        })
    }
  }
}
</script>
