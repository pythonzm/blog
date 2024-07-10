### 后端API使用[gin](https://github.com/gin-gonic/gin)编写，前端页面使用element-ui，前端直接借鉴了这个大佬的项目：https://github.com/PanJiaChen/vue-admin-template

> 整体就是以黑白为主，不会做手机适配，所以手机浏览布局会有问题，如果哪位大佬感兴趣做一下手机适配就非常感谢了

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
* 前端开启：编辑 fronted\src\settings.js 中相关配置
* 后端开启：编辑 backend\conf\config.yml 中相关配置
