### 后端API使用[gin](https://github.com/gin-gonic/gin)编写，前端页面使用element-ui，前端直接借鉴了这个大佬的项目：https://github.com/PanJiaChen/vue-admin-template

> **项目特点**：前端界面支持暗色高亮质感并已全面适配手机端访问，包括响应式自适应布局、移动端抽屉菜单及专属手机端翻页样式优化等。

## [线上地址：https://www.poorops.com](https://www.poorops.com)

## [Docker快速部署](https://github.com/pythonzm/blog/wiki/Docker-部署)

## 标准部署

### 准备
  1. 安装mysql（略）
  2. 安装redis（略）
  3. 安装elastic（略），主要用于全文检索功能，如果不用全文检索可以不安装
  3. 创建blog数据库
  
  ```
  mysql> create database blog DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
  mysql> grant all privileges on blog.* to YOURUSER@'%' identified by 'YOUPASSWORD';
  mysql> flush privileges;
  ```
  
### 安装
  1. go get github.com/pythonzm/blog
  2. 生成数据表
  
  ```
  cd backend
  bin/goose -dir migrations/ mysql "YOURUSER:YOURPASSWORD@tcp(YOURIP:YOURPORT)/blog?charset=utf8" up
  ```
  > 更多goose使用方法，参考：https://github.com/pressly/goose
  
### 启动后端
1. 重命名配置文件 mv backend/conf/config.yml.example backend/conf/config.yml
2. 修改 backend/conf/config.yml 中的配置信息

`go run main.go`
  
### 启动前端
  > cd fronted
  1. npm install
  2. npm run dev
  > 如果出现类似 Command failed: git clone --mirror -q git://github.com/adobe-webplatform/eve.git
  > npm ERR! fatal: read error: Invalid argument 的报错
  > 首先将本地git版本升级到最新的，如果还有问题则执行如下命令
  > git config --global url."https://".insteadOf git:// 
  
  #### 构建生产环境
  npm run build:prod

### 访问
  前台界面：$YOURIP:9528/#/
  
  后台界面: $YOURIP:9528/#/admin     默认用户名密码：admin/12346

## 2024-07-09更新
搜索功能接入Algolia，默认按照文章标题搜索
如果要开启Algolia搜索，需要同时在前端和后端开启
* 前端开启：编辑 `fronted/src/settings.js` 中相关配置，并在 `fronted/` 目录下创建 `.env.local` 配置文件存放敏感凭证（见下方配置说明）。
* 后端开启：编辑 `backend/conf/config.yml` 中相关配置。

---

## 💡 前端高级开发与配置说明

### 1. 敏感凭证防泄漏配置 (Algolia)
前端的 Algolia API 密钥已支持通过**环境变量**在本地覆写，防止本地密钥因 `git commit` 被意外提交泄露：
* 在前端根目录 `fronted/` 下新建 `.env.local` 本地配置文件，填入敏感凭证：
  ```env
  VUE_APP_ALGOLIA_APP_ID=你的APP_ID
  VUE_APP_ALGOLIA_API_KEY=你的API_KEY
  VUE_APP_ALGOLIA_INDEX_NAME=你的索引名称
  ```
* 项目中 `fronted/src/settings.js` 会优先读取本地 `.env.local` 文件中的配置。
* 该 `.env.local` 文件已加入全局忽略，永远不会被提交。

### 2. 控制台调试保护开关 (Anti-Debug)
在发布线上时，本站支持在生产环境中禁止访客通过 F12 进入控制台调试进行安全加固：
* **本地开发环境**（`npm run dev`）下，此调试阻断脚本会**自动关闭**，无需任何操作，方便正常本地开发。
* **生产环境打包**（`npm run build:prod`）下，此调试阻断脚本会**自动开启**。
* **手动全局控制**：若需要手动完全开启/关闭，可在 `fronted/src/settings.js` 中修改配置项 `antiDebug: true | false`。
