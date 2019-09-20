<template>
  <div class="app-container">
    <el-form ref="user" :model="user" label-width="120px">
      <el-form-item>
        <pan-thumb :image="user.avatar" />

        <el-button
          type="primary"
          icon="upload"
          style="position: absolute;bottom: 15px;margin-left: 40px;"
          @click="imagecropperShow = true"
          >选择头像</el-button
        >

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
        <el-input v-model="user.nickname" />
      </el-form-item>
      <el-form-item label="个人说明">
        <el-input v-model="user.introduction" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">确认修改</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import ImageCropper from '@/components/ImageCropper'
import PanThumb from '@/components/PanThumb'
import { editUser } from '@/api/user'

export default {
  components: { ImageCropper, PanThumb },
  data () {
    return {
      user: {},
      imagecropperShow: false,
      imagecropperKey: 0,
      params: {
        t: 'avatar'
      }
    }
  },

  computed: {
    ...mapGetters([
      'nickname',
      'avatar',
      'introduction'
    ])
  },
  created () {
    this.getUser()
  },

  methods: {
    getUser () {
      this.user = {
        nickname: this.nickname,
        avatar: this.avatar,
        introduction: this.introduction
      }
    },
    onSubmit () {
      editUser(this.user).then(response => {
        this.$message({
          message: response.msg,
          type: 'success'
        })
      })
    },
    cropSuccess (resData) {
      this.imagecropperShow = false
      this.imagecropperKey = this.imagecropperKey + 1
      this.user.avatar = resData.AvatarFullUrl
    },
    close () {
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

