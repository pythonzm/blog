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
          >发布
        </el-button>
        <el-button v-loading="loading" type="warning" @click="draftForm"
          >存草稿
        </el-button>
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
                >标题
              </MDinput>
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
                      filterable
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
                  <el-button
                    type="info"
                    plain
                    icon="el-icon-plus"
                    style="margin-left: 5px;"
                    @click="createCT('category')"
                  ></el-button>
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
                      filterable
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
                  <el-button
                    type="info"
                    plain
                    icon="el-icon-plus"
                    style="margin-left: 5px;"
                    @click="createCT('tag')"
                  ></el-button>
                </el-col>
              </el-row>
            </div>
          </el-col>
        </el-row>

        <el-form-item prop="content" style="margin-bottom: 30px;">
          <markdown-editor
            ref="editor"
            v-model="postForm.content"
            :options="{
              hideModeSwitch: true,
              previewStyle: 'tab',
              hooks: { addImageBlobHook: onAddImageBlob }
            }"
            height="1000px"
          />
        </el-form-item>
      </div>
    </el-form>
    <el-dialog :title="ctTile[ct]" :visible.sync="dialogFormVisible">
      <el-form
        :model="ct === 'category' ? category : tag"
        label-position="left"
        label-width="70px"
        style="width: 400px; margin-left:50px;"
      >
        <el-form-item :label="ctLable[ct]">
          <el-input
            v-if="ct === 'category'"
            v-model="category.category_name"
            required
          />
          <el-input v-else-if="ct === 'tag'" v-model="tag.tag_name" required />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button
          type="primary"
          @click="ct === 'category' ? newCategory() : newTag()"
          >确认</el-button
        >
      </div>
    </el-dialog>
  </div>
</template>

<script>
import MarkdownEditor from '@/components/MarkdownEditor'
import MDinput from '@/components/MDinput'
import Sticky from '@/components/Sticky' // 粘性header组件
import { fetchArticle, createArticle, editArticle, uploadImage } from '@/api/article'
import { fetchCategoryList, createCategory } from '@/api/category'
import { fetchTagList, createTag } from '@/api/tag'
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
      },
      category: {
        category_name: ''
      },
      tag: {
        tag_name: ''
      },
      ct: '',
      dialogFormVisible: false,
      ctTile: {
        category: '添加分类',
        tag: '添加标签'
      },
      ctLable: {
        category: '分类名称',
        tag: '标签名称'
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
          for (const tag of response.data.tags) {
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
    },
    onAddImageBlob (blob, callback) {
      let formData = new FormData();
      formData.append('file', blob);
      formData.set('t', 'image')
      uploadImage(formData).then(response => {
        callback(window.location.protocol + '//' + window.location.host + response.data.ImageFullUrl, blob.name);
      }).catch(error => {
        console.log(error);
      });
    },
    resetTemp () {
      this.category = {
        category_name: ''
      }
      this.tag = {
        tag_name: ''
      }
    },
    createCT (ct) {
      this.resetTemp()
      this.dialogFormVisible = true
      this.ct = ct
    },
    newCategory () {
      createCategory(this.category).then(response => {
        this.categories.push(response.data)
        this.dialogFormVisible = false
        this.$notify({
          title: 'Success',
          message: '添加分类成功',
          type: 'success',
          duration: 2000
        })
      })
    },
    newTag () {
      createTag(this.tag).then(response => {
        this.tags.push(response.data)
        this.dialogFormVisible = false
        this.$notify({
          title: 'Success',
          message: '添加标签成功',
          type: 'success',
          duration: 2000
        })
      })
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