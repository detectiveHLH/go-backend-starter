<h1 align="center">go-backend-starter</h1>

<div align="center"> 
基于gin、gorm的RESTFUL风格的后端框架，使用dep作为项目的包管理器
</div>

简体中文 | [English](./README-en.md)

## 安装golang
Mac下，安装golang，使用以下命令便可以安装
```
brew install go
```
## 配置环境变量
执行以下命令打开配置文件
```
vim ~/.bash_profile
```
添加如下代码
```
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```
然后执行以下代码使改动生效
```
source ~/.bash_profile
```
## 构建本地工作区
```
mkdir $HOME/go
```
在上面的go目录下，新建三个目录，分别是bin、pkg、src。bin是存放编译之后的可执行文件的，pkg存放第三方包，src存放项目源代码。

你的项目应该放在$HOME/go/src下。

## IDE
建议使用JetBrains的 [GoLand](https://www.jetbrains.com/go/?fromMenu)，在Preferences中设置项目的GOPATH为$HOME/go。

## 安装包管理器
```
brew install dep
```

## clone项目
```
git clone git@github.com:detectiveHLH/go-backend-starter.git
```

## 安装依赖包
```
dep ensure
```

## 启动项目
```
cd $HOME/go/src/go-backend-starter && go run main.go
```