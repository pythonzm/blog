<template>
  <div class="about-page-container">
    <!-- 个人简介头图卡片 -->
    <div class="profile-header-card">
      <div class="profile-avatar-wrapper">
        <div class="profile-avatar-inner">
          <img class="profile-avatar" :src="myAvatar" alt="Avatar">
        </div>
      </div>
      <div class="profile-info">
        <h1 class="profile-name">PoorOPS</h1>
        <p class="profile-title">可怜的/贫穷的运维</p>
        <p class="profile-bio">。</p>
        <div class="profile-social-links">
          <a href="https://github.com/pythonzm" target="_blank" class="social-btn github">
            <svg-icon icon-class="github" /> GitHub
          </a>
          <a href="mailto:poorops@163.com" class="social-btn email">
            <i class="el-icon-message" /> Email
          </a>
        </div>
      </div>
    </div>

    <!-- 博主详介 -->
    <div class="about-content-card">
      <div class="card-header-bar">
        <h3 class="card-header-title">
          <i class="el-icon-user icon" /> 博主生平
        </h3>
      </div>
      <div class="markdown-content" v-html="about" />
    </div>
  </div>
</template>

<script>
import myAvatar from '@/assets/avatar/author.png'
import { getAbout } from '@/api/user'
import marked from 'marked'
export default {
  name: 'About',
  data() {
    return {
      myAvatar: myAvatar,
      about: ''
    }
  },
  created() {
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
    fetchData() {
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

<style scoped lang="scss">
.about-page-container {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 28px;
}

// 个人卡片头部
.profile-header-card {
  background: #ffffff;
  border-radius: 12px;
  border: 1px solid rgba(226, 232, 240, 0.8);
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
  padding: 40px;
  display: flex;
  align-items: center;
  gap: 32px;
  text-align: left;
}

.profile-avatar-wrapper {
  flex-shrink: 0;
}

.profile-avatar-inner {
  width: 110px;
  height: 110px;
  border-radius: 50%;
  padding: 4px;
  background: linear-gradient(135deg, #1e293b, #475569);
  box-shadow: 0 10px 20px -8px rgba(30, 41, 59, 0.2);
  transition: transform 0.5s ease;

  &:hover {
    transform: rotate(5deg) scale(1.03);
  }
}

.profile-avatar {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
  background: #ffffff;
  border: 2px solid #ffffff;
}

.profile-info {
  flex: 1;
  min-width: 0;
}

.profile-name {
  font-size: 26px;
  font-weight: 800;
  color: #0f172a;
  margin: 0 0 6px 0;
}

.profile-title {
  font-size: 15px;
  font-weight: 600;
  color: #1e293b;
  margin: 0 0 12px 0;
}

.profile-bio {
  font-size: 14.5px;
  color: #64748b;
  line-height: 1.6;
  margin: 0 0 20px 0;
}

.profile-social-links {
  display: flex;
  gap: 12px;
}

.social-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 13.5px;
  font-weight: 600;
  padding: 8px 16px;
  border-radius: 20px;
  text-decoration: none;
  transition: all 0.3s ease;

  &.github {
    background-color: #0f172a;
    color: #ffffff;

    &:hover {
      background-color: #1e293b;
      transform: translateY(-1px);
      box-shadow: 0 4px 12px rgba(15, 23, 42, 0.25);
    }
  }

  &.email {
    background-color: #f1f5f9;
    color: #475569;
    border: 1px solid #e2e8f0;

    &:hover {
      background-color: #e2e8f0;
      color: #0f172a;
      transform: translateY(-1px);
    }
  }
}

// 详细简介内容
.about-content-card {
  background: #ffffff;
  border-radius: 12px;
  border: 1px solid rgba(226, 232, 240, 0.8);
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
  padding: 40px;
  text-align: left;
}

.card-header-bar {
  border-bottom: 1px solid #f1f5f9;
  padding-bottom: 16px;
  margin-bottom: 24px;
}

.card-header-title {
  font-size: 18px;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
  display: flex;
  align-items: center;
  gap: 6px;

  .icon {
    color: #1e293b;
  }
}

// 高级 Markdown 样式覆写
.markdown-content {
  font-size: 16px;
  color: #334155;
  line-height: 1.8;
  word-break: break-word;

  ::v-deep {
    p {
      margin-top: 0;
      margin-bottom: 20px;
    }

    h1, h2, h3, h4, h5, h6 {
      color: #0f172a;
      font-weight: 700;
      margin-top: 32px;
      margin-bottom: 16px;
      line-height: 1.3;
    }

    h1 {
      font-size: 24px;
      border-bottom: 1px solid #e2e8f0;
      padding-bottom: 8px;
    }

    h2 {
      font-size: 20px;
      border-bottom: 1px solid #f1f5f9;
      padding-bottom: 6px;
    }

    h3 {
      font-size: 18px;
      border-bottom: none;
    }

    a {
      color: #1e293b !important;
      text-decoration: none;
      border-bottom: 1px dashed rgba(30, 41, 59, 0.4);
      transition: all 0.2s ease;

      &:hover {
        color: #0f172a !important;
        border-bottom-style: solid;
      }
    }

    blockquote {
      background: linear-gradient(to right, #f8fafc, rgba(248, 250, 252, 0.2));
      border-left: 4px solid #1e293b;
      color: #475569;
      padding: 12px 20px;
      margin: 20px 0;
      border-radius: 0 8px 8px 0;
      font-style: italic;
    }

    code {
      background-color: #f1f5f9;
      color: #e11d48;
      padding: 3px 6px;
      border-radius: 4px;
      font-size: 0.9em;
      font-family: Menlo, Monaco, Consolas, "Courier New", monospace;
    }

    pre {
      background-color: #1e293b;
      padding: 16px 20px;
      border-radius: 8px;
      overflow-x: auto;
      margin: 20px 0;

      code {
        background-color: transparent;
        color: #f8fafc;
        padding: 0;
        border-radius: 0;
        font-size: 14px;
      }
    }

    img {
      max-width: 100%;
      border-radius: 8px;
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
      margin: 16px auto;
      display: block;
    }

    ul, ol {
      padding-left: 20px;
      margin-bottom: 20px;

      li {
        margin-bottom: 6px;
      }
    }
  }
}

@media (max-width: 768px) {
  .profile-header-card {
    flex-direction: column;
    padding: 30px 20px;
    align-items: center;
    text-align: center;
  }

  .profile-social-links {
    justify-content: center;
  }

  .about-content-card {
    padding: 30px 20px;
  }
}
</style>
