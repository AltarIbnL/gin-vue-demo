### 1. 运行后端程序

> 先确保你电脑上正确安装了 golang 环境

从master分支拉取后端golang代码

```
# 拉取代码
git clone -b master https://github.com/kotlindev/gin-demo.git backend
# 进入项目目录
cd  backend
# 安装项目依赖
go get
```

打开 `config/application.yaml` 文件，修改数据库链接配置，修改项目运行端口，确保端口不被占用，参考如下

```
server:
  port: 9090
datasource:
  driverName: mysql
  host: 127.0.0.1
  port: 3306
  database: api
  username: root
  password: 12345
  charset: utf8
  loc: Asia/Shanghai
```

启动项目

```
go run routes.go
```



### 运行前端程序

> 先确保你电脑上正确安装了 npm 环境，并安装了 vue、yarn

从vue分支拉取前端vue代码

```
# 拉取代码
git clone -b vue https://github.com/kotlindev/gin-demo.git vue
# 进入项目目录
cd  vue
# 安装项目依赖
yarn install
```



根据1中的 后端代码的运行端口，修改 `.env.development.local` 和 `.env.development` 两个配置文件，修改配置如下为

```
VUE_APP_BASE_URL= http://localhost:9090/api/
```

在运行项目

```
yarn serve
```

