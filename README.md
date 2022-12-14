# 介绍
go 脚手架项目，目前项目使用的是 go-chassis-rpc template

# Quick Start 
##  1 生成 go_scaffold 可运行文件(可跳过，目录中 go_scaffold 是已经生成好的)
   - packr Binary Installation [参考](https://github.com/gobuffalo/packr)
   - packr2 build
   - CGO_ENABLED=0  GOOS=linux  GOARCH=amd64  go build -o go_scaffold main.go
   - packr2 clean
## 2.1 直接使用

将 go_scaffold 放到某个路径，并加入环境变量，然后执行 go_scaffold 命令，将在当前目录生成项目
```
go_scaffold <project_name>
```
## 2.2 生成 docker 镜像，同时集成 protoc 

生成镜像

`docker build -t go_scaffold:v1 .`

在当前目录生成项目 

```docker run --rm  -v $(PWD):$(PWD) -w $(PWD)  go_scaffold:v1  go_scaffold <project_name>```

在项目中使用 protoc 插件：
```
#在项目目录执行 
make pb
 ```
# 自定义 template
1. 修改 template 目录地址
const TEMPLATE_DIR = "../go_chassis_template"
2. 修改 template key word
const TMEPLATE_KEY_WORD = "template"

# 相关文章
[使用 Docker 搭建项目脚手架](https://zhuanlan.zhihu.com/p/558318498)

