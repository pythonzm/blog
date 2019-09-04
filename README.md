### 该项目是用于练习go语言以及vue学习，目前完成了后端api的编写，采用的是[gin](https://github.com/gin-gonic/gin)框架

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
  bin/goose goose mysql "YOURUSER:YOURPASSWORD@tcp(YOURIP:YOURPORT)/blog?charset=utf8" up
  ```
   
### 启动后端
首先修改conf/config.yml中的配置信息

`go run main.go`
  
### 启动管理
  > cd admin
  1. npm install
  2. npm run dev
## TODO LIST
  使用vue编写页面展示

### 添加文章
![image](https://github.com/pythonzm/blog/blob/master/screenshots/add_article.png)

### 文章列表
![image](https://github.com/pythonzm/blog/blob/master/screenshots/article_list.png)
