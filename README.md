# Frame Web 项目

## 项目简介
一个前后端分离的Web应用框架，使用Go语言作为后端，Vue.js作为前端。

## 技术栈
### 后端技术
- Web框架: Gin
- ORM: GORM
- 配置管理: Viper
- 日志系统: Zap
- 数据库: MySQL
- 缓存: Redis

### 前端技术
- 框架: Vue.js 3
- 构建工具: Vite
- UI组件库: Element Plus (可选)

## 项目结构

分支 frame-init

配置viper
日志zap
持久层gorm
web框架gin


一个前后端分离的Web应用框架，使用Go语言作为后端，Vue.js作为前端。

接口文档地址http://localhost:8888/swagger/index.html


## 快速开始
### 后端启动
```bash
cd server
go run main.go
### 前端启动
cd frontend
npm install
npm run dev

持续完善中....
后端目录结构
server/
├── api/               # API 控制器层
│   ├── file_upload.go # 文件上传API
│   └── userApi.go     # 用户相关API
├── config/            # 配置文件
│   ├── config.yaml    # 主配置文件
│   ├── db_list.go     # 数据库配置
│   ├── gorm_mysql.go  # GORM MySQL 配置
│   └── ...
├── core/              # 核心模块
│   ├── db.go         # 数据库初始化
│   ├── viper.go      # 配置读取
│   └── zap.go        # 日志模块
├── initialize/        # 初始化模块
│   └── router.go     # 路由初始化
├── middleware/       # 中间件
│   ├── auth.go       # 认证中间件
│   ├── cor.go        # 跨域中间件
│   └── ...
├── model/            # 模型层
│   ├── common.go     # 公共模型
│   ├── request/      # 请求参数模型
│   └── response/     # 响应模型
├── svc/              # 服务层
│   ├── models/       # 数据库模型
│   ├── service/      # 业务逻辑
│   └── ...
├── utils/            # 工具类
│   ├── upload/       # 文件上传工具
│   └── redis_util.go # Redis工具
└── main.go           # 程序入口
