### 后端 API 使用[gin](https://github.com/gin-gonic/gin)编写，前端页面使用 element-ui，前端直接借鉴了这个大佬的项目：https://github.com/PanJiaChen/vue-admin-template

> 整体就是以黑白为主，不会做手机适配，所以手机浏览布局会有问题，如果哪位大佬感兴趣做一下手机适配就非常感谢了

## [线上地址](https://www.poorops.com)

## 搜索功能使用 elastic 做的全文检索，由于租的机器内存比较小，elastic 装上跑着有点儿费劲，线上就没有使用搜索功能

## [Docker 快速部署(生产环境)](https://github.com/pythonzm/blog/wiki/Docker-部署)

## 标准部署

### 准备

1. 安装 mysql（略）
2. 安装 redis（略）
3. 安装 elastic（略），主要用于全文检索功能，如果不用全文检索可以不安装
4. 创建 blog 数据库

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

> 更多 goose 使用方法，参考：https://github.com/pressly/goose

### 启动后端

首先修改 conf/config.yml 中的配置信息

`go run main.go`

### 启动前端

> cd fronted

1. npm install
2. npm run dev

#### 构建生产环境

npm run build:prod

### 访问

前台界面：\$YOURIP:9528/#/

后台界面: \$YOURIP:9528/#/admin 默认用户名密码：admin/12346

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

### 手机端

![image](https://github.com/pythonzm/blog/blob/master/screenshots/mobile.png)
