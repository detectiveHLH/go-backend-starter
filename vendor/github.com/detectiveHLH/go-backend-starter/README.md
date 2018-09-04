# backend
基于gin、gorm的RESTFUL风格的后端框架

### 包管理器
包管理器，使用的 [govendor](https://github.com/kardianos/govendor)，安装命令如下。
```go
go get -u github.com/kardianos/govendor
```
- 初始化
```go
govendor init
```
- 添加所有依赖包
```go
govendor add +external
```
- 获取更新
```go
govendor fetch 具体包名
```