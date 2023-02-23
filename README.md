# 介绍
go 脚手架项目，具有以下功能：
1. 一键生成基于 go-chassis 的新项目目录结构，及相关基础代码
2. 一键生成 proto.go 文件
3. 根据 pb 文件或数据库文件一键生成接口文档（TODO）
4. 通过数据库一键生成 crud 代码（TODO）


# Quick Start 
## 方式一：编译并使用
1 编译(可跳过，目录中 go_scaffold 是已经生成好的)
   - packr Binary Installation [参考](https://github.com/gobuffalo/packr)
   - packr2 build
   - CGO_ENABLED=0  GOOS=linux  GOARCH=amd64  go build -o go_scaffold main.go
   - packr2 clean
2 使用
将 go_scaffold 放到某个路径，并加入环境变量，然后执行 go_scaffold 命令，将在当前目录生成项目
```
go_scaffold <project_name>
```


## 方式二：生成 docker 镜像使用

1. 生成镜像

`docker build -t go_scaffold:v1 .`

2. 使用镜像

功能1： 在当前目录生成项目 

```docker run --rm  -v $(PWD):$(PWD) -w $(PWD)  go_scaffold:v1  go_scaffold <project_name>```

功能2：一键生成 proto.go 文件。可以忽略操作系统及 protc 版本差异
```
#在项目目录执行 
make pb
 ```

# 自定义 template
1. 修改 template 目录地址
const TEMPLATE_DIR = "../go_chassis_template"
2. 修改 template key word
const TMEPLATE_KEY_WORD = "template"

# 原理解析
[使用 Docker 搭建项目脚手架](https://zhuanlan.zhihu.com/p/558318498)

