# 介绍
go 脚手架项目，目前项目使用的是 go-chassis-rpc template



# Quick Start 
1. 生成 go_scaffold 可运行文件(目录中 go_scaffold 是已经生成好的)
   - packr Binary Installation [参考](https://github.com/gobuffalo/packr)
   - packr2 build
   - go build
   - packr2 clean

2. 将 go_scaffold 放到某个路径，并加入环境变量
3. 执行 go_scaffold 命令，将在当前目录生成项目。
```
go_scaffold <project_name>
```
 
# 自定义 template
1. 修改 template 目录地址
const TEMPLATE_DIR = "../go_chassis_template"
2. 修改 template key word
const TMEPLATE_KEY_WORD = "template"

# 集成 go 脚手架到 docker
[使用 Docker 搭建项目脚手架] ()

