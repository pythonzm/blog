### 该项目是用于练习go语言以及vue学习，后端API使用[gin](https://github.com/gin-gonic/gin)编写，前端页面使用element-ui，直接借鉴了这个大佬的项目：https://github.com/PanJiaChen/vue-admin-template

### [线上站点](https://www.poorops.com)

## [Docker快速部署](https://github.com/pythonzm/blog/wiki/Docker-部署)

## 标准部署

### 准备
  1. 安装mysql（略）
  2. 安装redis（略）
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
首先修改conf/config.yml中的配置信息

`go run main.go`
  
### 启动前端
  > cd fronted
  1. npm install
  2. npm run dev

### 访问
  前台界面：$YOURIP:9528/#/
  
  后台界面: $YOURIP:9528/#/admin     默认用户名密码：admin/12346
  
## TODO LIST
  美化界面

### 添加文章
![image](https://github.com/pythonzm/blog/blob/master/screenshots/add_article.png)

### 文章列表
![image](https://github.com/pythonzm/blog/blob/master/screenshots/article_list.png)

### 前台首页
![image](https://github.com/pythonzm/blog/blob/master/screenshots/home.png)

### 文章详情页
![image](https://github.com/pythonzm/blog/blob/master/screenshots/article.png)
