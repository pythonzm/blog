<template>
  <div id="comments">
    <div id="comments-input-top" class="input-wrap" v-loading="loading">
      <div class="input-top">
        <el-input
          v-if="!isAuthor"
          class="top-item"
          size="mini"
          v-model="temp.username"
          placeholder="称呼（必填）"
        >
        </el-input>
      </div>
      <el-input
        class="input-area"
        id="comments-content-area"
        type="textarea"
        size="mini"
        :rows="5"
        resize="none"
        v-model="temp.content"
        :placeholder="placeholder"
      >
      </el-input>
      <div class="btn-wrap">
        <div class="action-btn">
          <span
            class="cancel-btn"
            @click="temp.content = ''"
            v-show="temp.content !== ''"
            >取消</span
          >
          <span class="send-btn" @click="check">
            <el-button type="info" round>发送</el-button>
          </span>
        </div>
      </div>
    </div>
    <p class="count" v-if="count === 0">暂无评论</p>
    <p class="count" v-else>{{ count }}条评论</p>
    <ul class="comments-wrap" v-if="count !== 0">
      <li
        class="comments-item"
        v-for="(comment, index) in comments"
        :key="index"
      >
        <div class="comments-info">
          <div><img class="avatar" :src="getAvatar(comment)" /></div>
          <div>
            <div class="username">{{ getName(comment) }}</div>
            <div class="time">{{ comment.created_time }}</div>
          </div>
        </div>
        <p class="content">{{ comment.content }}</p>
        <div class="reply">
          <el-tag type="info" @click="replyRoot(comment)">回复</el-tag>
        </div>
        <ul class="comments-children" v-if="comment.children.length > 0">
          <li
            class="comments-child"
            v-for="(child, index) in comment.children"
            :key="index"
          >
            <div class="comments-info">
              <div><img class="avatar" :src="getAvatar(child)" /></div>
              <div>
                <div class="username">{{ getName(child) }}</div>
                <div class="time">{{ child.created_time }}</div>
              </div>
            </div>
            <p class="content">{{ child.content }}</p>
            <div class="reply">
              <el-tag type="info" @click="replyChild(comment, child)"
                >回复</el-tag
              >
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
  data () {
    return {
      placeholder: '写下您的评论',
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
        content: '',
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
  created () {
    this.getList(this.temp.article_id)
  },
  methods: {
    getName (comment) {
      return comment.username + (comment.is_author ? '（作者）' : '')
    },
    getAvatar (comment) {
      return comment.is_author ? author_avatar : guest_avatar
    },
    getList (id) {
      fetchList(id)
        .then(response => {
          this.comments = response.data.total === 0 ? [] : response.data.items
          this.count = response.data.total
        })
        .catch(err => {
          console.log(err)
        })
    },
    addComment () {
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
          const comment = this.comments.find(item => item.id === response.data.root_id.Int64);
          comment.children.push(response.data);
          this.replyUserName = ''
          this.temp.parent_id.Int64 = 0
          this.temp.parent_id.Valid = false
          this.temp.root_id.Int64 = 0
          this.temp.root_id.Valid = false
        }
        this.$message({
          message: '评论添加成功',
          type: 'success'
        });
        this.temp.content = ''
      })
      this.loading = false
    },
    check () {
      if (this.temp.username === '' && !this.isAuthor) {
        this.$message.error('请填写您的称呼');
        return
      }
      if (this.temp.content === '' || this.temp.content === this.replyUserName) {
        this.$message.error('评论内容不能为空');
        return
      }
      this.addComment()
    },
    replyFocus () {
      let top = document.getElementById('comments-input-top').getBoundingClientRect().top
      top += document.body.scrollTop || document.documentElement.scrollTop
      window.scrollTo(0, top + 80)
      document.getElementById('comments-content-area').focus()
    },
    replyRoot (comment) {
      this.temp.parent_id.Int64 = comment.id
      this.temp.parent_id.Valid = true
      this.temp.root_id.Int64 = comment.id
      this.temp.root_id.Valid = true
      this.replyUserName = `@${comment.username} `
      this.temp.content = this.replyUserName
      this.replyFocus()
    },
    replyChild (comment, child) {
      this.temp.parent_id.Int64 = child.id
      this.temp.parent_id.Valid = true
      this.temp.root_id.Int64 = comment.id
      this.temp.root_id.Valid = true
      this.replyUserName = `@${child.username} `
      this.temp.content = this.replyUserName
      this.replyFocus()
    },
  }
}
</script>

<style>
.input-top input.el-input__inner:focus,
.input-area textarea.el-textarea__inner:focus {
  border-color: #000;
}
</style>

<style scoped>
#comments {
  margin-top: 20px;
  position: relative;
  background-color: #fff;
  animation: show-data-v-a600e01a 0.8s;
  text-align: left;
}
.input-top {
  margin-bottom: 5px;
  border-color: black;
}
.top-item {
  min-width: 194px;
  margin-top: 5px;
}
.btn-wrap {
  position: relative;
  margin-top: 10px;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: flex-end;
  transition: padding-bottom 0.3s;
}
.cancel-btn {
  cursor: pointer;
  color: #5e6877;
  transition: color 0.3s;
  margin-right: 10px;
}
.count {
  margin-top: 10px;
  color: #262a30;
  font-weight: 700;
  border-bottom: 1px solid #eee;
  padding: 5px 0;
}
ul {
  margin: 0;
  padding: 0;
}
li {
  list-style: none;
}
.comments-item {
  padding: 10px 5px;
  transition: background-color 0.3s;
}
.comments-info {
  display: flex;
  flex-direction: row;
  margin-bottom: 5px;
}
.avatar {
  width: 32px;
  height: 32px;
  margin-right: 5px;
  border-radius: 50%;
}
.username {
  margin-bottom: 2px;
  font-size: 14px;
}
.reply {
  cursor: pointer;
  color: #5e6877;
  padding-left: 40px;
  width: fit-content;
}
.time {
  font-size: 12px;
  color: #999;
}
.comments-item .content {
  padding-left: 40px;
  font-size: 14px;
  line-height: 16px;
  word-break: break-all;
  margin: 0;
}
.comments-children {
  margin-left: 12px;
  border-left: 2px solid #999;
  margin-top: 5px;
}
.comments-children .comments-child {
  padding: 10px 5px;
  padding-left: 17px;
}
</style>