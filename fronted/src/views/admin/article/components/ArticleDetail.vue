<template>
  <div class="createPost-container">
    <el-form
      ref="postForm"
      :model="postForm"
      :rules="rules"
      class="form-container"
    >
      <sticky :z-index="10" :class-name="'sub-navbar ' + postForm.status">
        <el-button
          v-loading="loading"
          style="margin-left: 10px;"
          type="success"
          @click="submitForm"
          >发布</el-button
        >
        <el-button v-loading="loading" type="warning" @click="draftForm"
          >存草稿</el-button
        >
      </sticky>

      <div class="createPost-main-container">
        <el-row>
          <el-col :span="24">
            <el-form-item style="margin-bottom: 40px;" prop="title">
              <MDinput
                v-model="postForm.title"
                :maxlength="100"
                name="name"
                required
                >标题</MDinput
              >
            </el-form-item>

            <div class="postInfo-container">
              <el-row>
                <el-col :span="8">
                  <el-form-item
                    label-width="60px"
                    label="分类"
                    class="postInfo-container-item"
                    prop="category_id"
                  >
                    <el-select
                      v-model="postForm.category_id"
                      placeholder="请选择分类"
                    >
                      <el-option
                        v-for="item in categories"
                        :key="item.id"
                        :label="item.category_name"
                        :value="item.id"
                      />
                    </el-select>
                  </el-form-item>
                </el-col>

                <el-col :span="8">
                  <el-form-item
                    label-width="60px"
                    label="标签"
                    class="postInfo-container-item"
                    prop="tag_id"
                  >
                    <el-select
                      v-model="postForm.tag_id"
                      multiple
                      placeholder="请选择标签"
                    >
                      <el-option
                        v-for="item in tags"
                        :key="item.id"
                        :label="item.tag_name"
                        :value="item.id"
                      />
                    </el-select>
                  </el-form-item>
                </el-col>
              </el-row>
            </div>
          </el-col>
        </el-row>

        <el-form-item prop="content" style="margin-bottom: 30px;">
          <markdown-editor
            ref="editor"
            v-model="postForm.content"
            :options="{ hideModeSwitch: true, previewStyle: 'tab' }"
            height="1000px"
          />
        </el-form-item>
      </div>
    </el-form>
  </div>
</template>

<script>
import MarkdownEditor from '@/components/MarkdownEditor'
import MDinput from '@/components/MDinput'
import Sticky from '@/components/Sticky' // 粘性header组件
import { fetchArticle, createArticle, editArticle } from '@/api/article'
import { fetchCategoryList } from '@/api/category'
import { fetchTagList } from '@/api/tag'

const defaultForm = {
  status: 'draft',
  title: '', // 文章题目
  category_id: '',
  tag_id: [],
  content: '', // 文章markdwon内容
  html: '', // 文章html内容
  id: undefined
}

export default {
  name: 'ArticleDetail',
  components: {
    MarkdownEditor,
    MDinput,
    Sticky
  },
  props: {
    isEdit: {
      type: Boolean,
      default: false
    }
  },
  data () {
    return {
      postForm: Object.assign({}, defaultForm),
      loading: false,
      categories: [],
      tags: [],
      rules: {
        title: [
          { required: true, message: '标题不能为空', trigger: 'blur' },
          { min: 1, max: 20, message: '长度在 1 到 20 个字符', trigger: 'blur' }
        ],
        category_id: [
          {
            type: 'number',
            required: true,
            message: '请选择分类',
            trigger: 'change'
          }
        ],
        tag_id: [
          {
            type: 'array',
            required: true,
            message: '至少选择一个标签',
            trigger: 'change'
          }
        ],
        content: [
          { required: true, message: '文章内容不能为空', trigger: 'blur' }
        ]
      }
    }
  },
  created () {
    if (this.isEdit) {
      const id = this.$route.params && this.$route.params.id
      this.fetchData(id)
    } else {
      this.postForm = Object.assign({}, defaultForm)
    }
    this.getCategories()
    this.getTags()
  },
  methods: {
    fetchData (id) {
      fetchArticle(id, { admin: true })
        .then(response => {
          const ids = []
          this.postForm = response.data.article
          for (let tag of response.data.tags) {
            ids.push(tag.id)
          }
          this.postForm.tag_id = ids
        })
        .catch(err => {
          console.log(err)
        })
    },

    submitForm () {
      this.$refs.postForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.postForm.status = 'published'
          this.getHtml()
          if (this.isEdit) {
            const id = this.$route.params && this.$route.params.id
            this.editData(id, this.postForm)
          } else {
            this.createData(this.postForm)
          }
          this.loading = false
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    draftForm () {
      this.$refs.postForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.postForm.status = 'draft'
          this.getHtml()
          if (this.isEdit) {
            const id = this.$route.params && this.$route.params.id
            this.editData(id, this.postForm)
          } else {
            this.createData(this.postForm)
          }
          this.loading = false
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    getCategories () {
      fetchCategoryList().then(response => {
        this.categories = response.data
      })
    },
    getTags () {
      fetchTagList().then(response => {
        this.tags = response.data
      })
    },
    createData (data) {
      createArticle(data).then(response => {
        this.$notify({
          title: '成功',
          message: '发布文章成功',
          type: 'success',
          duration: 2000
        })
      })
    },
    editData (id, data) {
      editArticle(id, data).then(response => {
        this.$notify({
          title: '成功',
          message: '编辑文章成功',
          type: 'success',
          duration: 2000
        })
      })
    },
    getHtml () {
      this.postForm.html = this.$refs.editor.getHtml()
    }
  }
}
</script>

<style lang="scss" scoped>
@import "~@/styles/mixin.scss";

.createPost-container {
  position: relative;

  .createPost-main-container {
    padding: 40px 45px 20px 50px;

    .postInfo-container {
      position: relative;
      @include clearfix;
      margin-bottom: 10px;

      .postInfo-container-item {
        float: left;
      }
    }
  }

  .word-counter {
    width: 40px;
    position: absolute;
    right: 10px;
    top: 0px;
  }
}

.article-textarea /deep/ {
  textarea {
    padding-right: 40px;
    resize: none;
    border: none;
    border-radius: 0px;
    border-bottom: 1px solid #bfcbd9;
  }
}
</style>
