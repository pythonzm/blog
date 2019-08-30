<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item>
        <pan-thumb :image="form.avatar" />

        <el-button
          type="primary"
          icon="upload"
          style="position: absolute;bottom: 15px;margin-left: 40px;"
          @click="imagecropperShow=true"
        >更换头像</el-button>

        <image-cropper
          v-show="imagecropperShow"
          :key="imagecropperKey"
          field="file"
          :width="300"
          :height="300"
          :params="params"
          url="/upload"
          lang-type="zh"
          @close="close"
          @crop-upload-success="cropSuccess"
        />
      </el-form-item>
      <el-form-item label="昵称">
        <el-input v-model="form.nickname" />
      </el-form-item>
      <el-form-item label="个人说明">
        <el-input v-model="form.introduction" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">确认修改</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import ImageCropper from '@/components/ImageCropper'
import PanThumb from '@/components/PanThumb'
import { editUser } from '@/api/user'

export default {
  components: { ImageCropper, PanThumb },
  data() {
    return {
      form: {
        nickname: this.$store.state.user.nickname,
        avatar: this.$store.state.user.avatar,
        introduction: this.$store.state.user.introduction
      },
      imagecropperShow: false,
      imagecropperKey: 0,
      params: {
        t: 'avatar'
      }
    }
  },

  methods: {
    onSubmit() {
      editUser(this.form).then(response => {
        this.form.nickname = response.data.nickname
        this.form.introduction = response.data.introduction
        this.form.avatar = response.data.avatar

        this.$message({
          message: response.msg,
          type: 'success'
        })
      })
    },
    cropSuccess(resData) {
      this.imagecropperShow = false
      this.imagecropperKey = this.imagecropperKey + 1
      this.form.avatar = resData.AvatarFullUrl
    },
    close() {
      this.imagecropperShow = false
    }
  }
}
</script>

<style scoped>
.line {
  text-align: center;
}
.el-input,
.el-textarea {
  width: 50%;
}
.avatar {
  width: 200px;
  height: 200px;
  border-radius: 50%;
}
</style>

