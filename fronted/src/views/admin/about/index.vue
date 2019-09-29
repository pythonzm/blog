<template>
  <div class="app-container">
    <el-form ref="postForm" :model="postForm" class="form-container">
      <sticky :z-index="10" class-name="sub-navbar draft">
        <el-button
          v-loading="loading"
          style="margin-left: 10px;"
          type="success"
          @click="submitForm"
          >发布</el-button
        >
      </sticky>
      <el-form-item style="margin-bottom: 30px;">
        <markdown-editor
          ref="editor"
          v-model="postForm.about"
          :options="{ hideModeSwitch: true, previewStyle: 'tab' }"
          height="600px"
        />
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import Sticky from '@/components/Sticky'
import MarkdownEditor from '@/components/MarkdownEditor'
import { editUser, getAbout } from '@/api/user'
export default {
  name: 'AdminAbout',
  components: {
    MarkdownEditor,
    Sticky
  },
  data () {
    return {
      loading: false,
      postForm: {
        about: ''
      },
    }
  },
  created () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      getAbout()
        .then(response => {
          this.postForm.about = response.data
        })
        .catch(err => {
          console.log(err)
        })
    },
    submitForm () {
      this.loading = true
      editUser(this.postForm).then(response => {
        this.$notify({
          title: '成功',
          message: '编辑关于页成功',
          type: 'success',
          duration: 2000
        })
      })
      this.loading = false
    },
  },
}
</script>