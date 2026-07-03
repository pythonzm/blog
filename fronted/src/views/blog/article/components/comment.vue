<template>
  <div id="comments" class="comments-container">
    <!-- 评论输入框 -->
    <div id="comments-input-top" v-loading="loading" class="comment-input-card">
      <div v-if="!isAuthor" class="input-header">
        <el-input
          v-model="temp.username"
          class="name-input"
          size="medium"
          placeholder="您的称呼（必填）"
          prefix-icon="el-icon-user"
        />
      </div>
      <div class="input-body">
        <el-input
          id="comments-content-area"
          v-model="temp.content"
          class="content-textarea"
          type="textarea"
          size="medium"
          :rows="4"
          resize="none"
          :placeholder="placeholder"
        />
      </div>
      <div class="input-footer">
        <span
          v-show="temp.content !== ''"
          class="cancel-link"
          @click="resetCommentInput"
        >取消</span>
        <el-button class="submit-button" type="primary" round size="small" @click="check">
          发表评论
        </el-button>
      </div>
    </div>

    <!-- 评论计数 -->
    <div class="comments-header">
      <h3 class="comments-title">
        <i class="el-icon-chat-square" style="color: #1e293b; margin-right: 6px;" />
        {{ count === 0 ? '暂无评论' : `${count} 条评论` }}
      </h3>
    </div>

    <!-- 评论列表 -->
    <ul v-if="count !== 0" class="comments-list">
      <li
        v-for="(comment, index) in comments"
        :key="index"
        class="comment-node"
      >
        <div class="comment-main">
          <img class="comment-avatar" :src="getAvatar(comment)" alt="avatar">
          <div class="comment-body">
            <div class="comment-meta">
              <span class="comment-author" :class="{ 'author-badge': comment.is_author }">
                {{ comment.username }}
                <span v-if="comment.is_author" class="author-label">作者</span>
              </span>
              <span class="comment-date">{{ comment.created_time }}</span>
            </div>
            <p class="comment-text">{{ comment.content }}</p>
            <div class="comment-actions">
              <span class="action-reply-btn" @click="replyRoot(comment)">
                <i class="el-icon-position" /> 回复
              </span>
            </div>
          </div>
        </div>

        <!-- 子级回复列表 -->
        <ul v-if="comment.children && comment.children.length > 0" class="comments-nested-list">
          <li
            v-for="(child, childIndex) in comment.children"
            :key="childIndex"
            class="comment-node nested-node"
          >
            <div class="comment-main">
              <img class="comment-avatar avatar-small" :src="getAvatar(child)" alt="avatar">
              <div class="comment-body">
                <div class="comment-meta">
                  <span class="comment-author" :class="{ 'author-badge': child.is_author }">
                    {{ child.username }}
                    <span v-if="child.is_author" class="author-label">作者</span>
                  </span>
                  <span class="comment-date">{{ child.created_time }}</span>
                </div>
                <p class="comment-text">{{ child.content }}</p>
                <div class="comment-actions">
                  <span class="action-reply-btn" @click="replyChild(comment, child)">
                    <i class="el-icon-position" /> 回复
                  </span>
                </div>
              </div>
            </div>
          </li>
        </ul>
      </li>
    </ul>
  </div>
</template>

<script>
import guest_avatar from '@/assets/avatar/guest.png'
import author_avatar from '@/assets/avatar/author.png'
import { mapState } from 'vuex'
import { fetchList, createComment } from '@/api/comment'
export default {
  name: 'Comments',
  data() {
    return {
      placeholder: '说点什么吧，支持 Markdown 内容...',
      temp: {
        username: '',
        is_author: false,
        parent_id: {
          Int64: 0,
          Valid: false
        },
        root_id: {
          Int64: 0,
          Valid: false
        },
        article_id: parseInt(this.$route.query && this.$route.query.id),
        content: ''
      },
      comments: [],
      count: 0,
      replyUserName: '',
      loading: false
    }
  },
  computed: mapState({
    isAuthor: state => state.user.isAuthor === undefined ? state.user.isAuthor : JSON.parse(state.user.isAuthor)
  }),
  created() {
    this.getList(this.temp.article_id)
  },
  methods: {
    getName(comment) {
      return comment.username + (comment.is_author ? '（作者）' : '')
    },
    getAvatar(comment) {
      return comment.is_author ? author_avatar : guest_avatar
    },
    getList(id) {
      fetchList(id)
        .then(response => {
          this.comments = response.data.total === 0 ? [] : response.data.items
          this.count = response.data.total
        })
        .catch(err => {
          console.log(err)
        })
    },
    addComment() {
      this.loading = true
      if (this.isAuthor) {
        this.temp.username = 'poorops'
        this.temp.is_author = true
      }
      createComment(this.temp).then(response => {
        if (this.replyUserName === '') {
          response.data['children'] = []
          if (this.count === 0) {
            this.count = 1
          }
          this.comments.push(response.data)
        } else {
          const comment = this.comments.find(item => item.id === response.data.root_id.Int64)
          comment.children.push(response.data)
          this.replyUserName = ''
          this.temp.parent_id.Int64 = 0
          this.temp.parent_id.Valid = false
          this.temp.root_id.Int64 = 0
          this.temp.root_id.Valid = false
        }
        this.$message({
          message: '评论成功发表',
          type: 'success'
        })
        this.temp.content = ''
      })
      this.loading = false
    },
    check() {
      if (this.temp.username === '' && !this.isAuthor) {
        this.$message.error('请填写您的称呼')
        return
      }
      if (this.temp.content === '' || this.temp.content === this.replyUserName) {
        this.$message.error('评论内容不能为空')
        return
      }
      this.addComment()
    },
    replyFocus() {
      let top = document.getElementById('comments-input-top').getBoundingClientRect().top
      top += document.body.scrollTop || document.documentElement.scrollTop
      window.scrollTo(0, top - 100)
      document.getElementById('comments-content-area').focus()
    },
    replyRoot(comment) {
      this.temp.parent_id.Int64 = comment.id
      this.temp.parent_id.Valid = true
      this.temp.root_id.Int64 = comment.id
      this.temp.root_id.Valid = true
      this.replyUserName = `@${comment.username} `
      this.temp.content = this.replyUserName
      this.replyFocus()
    },
    replyChild(comment, child) {
      this.temp.parent_id.Int64 = child.id
      this.temp.parent_id.Valid = true
      this.temp.root_id.Int64 = comment.id
      this.temp.root_id.Valid = true
      this.replyUserName = `@${child.username} `
      this.temp.content = this.replyUserName
      this.replyFocus()
    },
    resetCommentInput() {
      this.temp.content = ''
      this.replyUserName = ''
      this.temp.parent_id.Int64 = 0
      this.temp.parent_id.Valid = false
      this.temp.root_id.Int64 = 0
      this.temp.root_id.Valid = false
    }
  }
}
</script>

<style lang="scss">
// 全局覆盖输入框高亮色
.comment-input-card {
  .name-input .el-input__inner,
  .content-textarea .el-textarea__inner {
    border-color: #e2e8f0;
    background-color: #ffffff !important;
    color: #475569 !important;
    transition: all 0.3s ease;

    &:focus {
      border-color: #1e293b;
      box-shadow: 0 0 0 3px rgba(30, 41, 59, 0.08);
    }
  }

  .name-input .el-input__inner {
    border-radius: 8px;
  }

  .content-textarea .el-textarea__inner {
    border-radius: 12px;
    padding: 12px 16px;
    font-family: inherit;
  }
}
</style>

<style scoped lang="scss">
.comments-container {
  margin-top: 40px;
  background-color: transparent;
}

// 评论输入区域美化
.comment-input-card {
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 20px;
  background: #f8fafc;
  margin-bottom: 30px;
}

.input-header {
  margin-bottom: 12px;
  max-width: 260px;
}

.input-body {
  margin-bottom: 14px;
}

.input-footer {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 16px;

  .cancel-link {
    cursor: pointer;
    font-size: 14px;
    color: #64748b;
    transition: color 0.2s ease;

    &:hover {
      color: #0f172a;
    }
  }

  .submit-button {
    background-color: #1e293b;
    border-color: #1e293b;
    font-weight: 500;
    padding: 9px 22px;

    &:hover {
      background-color: #0f172a;
      border-color: #0f172a;
    }
  }
}

// 评论数标题
.comments-header {
  border-bottom: 1px solid #f1f5f9;
  padding-bottom: 12px;
  margin-bottom: 24px;

  .comments-title {
    font-size: 18px;
    font-weight: 700;
    color: #0f172a;
    margin: 0;
    display: flex;
    align-items: center;
  }
}

// 评论列表树美化
.comments-list {
  padding: 0;
  margin: 0;
}

.comment-node {
  list-style: none;
  border-bottom: 1px solid #f1f5f9;
  padding: 20px 0;

  &:last-child {
    border-bottom: none;
  }
}

.comment-main {
  display: flex;
  gap: 16px;
}

.comment-avatar {
  width: 42px;
  height: 42px;
  border-radius: 50%;
  border: 1px solid #e2e8f0;
  object-fit: cover;
  background-color: #ffffff;

  &.avatar-small {
    width: 32px;
    height: 32px;
  }
}

.comment-body {
  flex: 1;
  min-width: 0;
}

.comment-meta {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 6px;
}

.comment-author {
  font-size: 14.5px;
  font-weight: 700;
  color: #1e293b;
  display: inline-flex;
  align-items: center;
  gap: 4px;

  &.author-badge {
    color: #0f172a;
  }

  .author-label {
    background-color: rgba(30, 41, 59, 0.08);
    color: #1e293b;
    font-size: 10.5px;
    font-weight: 600;
    padding: 1px 6px;
    border-radius: 4px;
    transform: scale(0.9);
  }
}

.comment-date {
  font-size: 12px;
  color: #94a3b8;
}

.comment-text {
  font-size: 14.5px;
  line-height: 1.6;
  color: #334155;
  margin: 0 0 8px 0;
  word-break: break-all;
}

.comment-actions {
  display: flex;
  gap: 16px;
  font-size: 12.5px;

  .action-reply-btn {
    color: #64748b;
    cursor: pointer;
    display: inline-flex;
    align-items: center;
    gap: 4px;
    font-weight: 500;
    transition: color 0.2s ease;

    &:hover {
      color: #1e293b;
    }
  }
}

// 嵌套子评论列表树美化
.comments-nested-list {
  padding: 0 0 0 24px;
  margin: 12px 0 0 24px;
  border-left: 2px solid #e2e8f0;

  .nested-node {
    border-bottom: none;
    border-top: 1px dashed #f1f5f9;
    padding: 16px 0 0 0;
    margin-top: 16px;

    &:first-child {
      border-top: none;
      padding-top: 0;
      margin-top: 0;
    }
  }
}
</style>
