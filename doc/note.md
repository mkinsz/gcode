# Golang Note

## Gin

## Graphql

## Docker

## Record

### Golang依赖

1. golang未安装依赖的情况下，通过下面这个命令即可安装所有依赖

```
go get -d -v ./...
```
 - -d标志只下载代码包，不执行安装命令；
 - -v打印详细日志和调试日志。这里加上这个标志会把每个下载的包都打印出来；
 - ./...这个表示路径，代表当前目录下所有的文件

```
go get -t -u -v ./...
```

### Net CMD
1. windows查看端口被占用
```
netstat -ano|findstr "8080"
```
记住LISTENING状态的PID, 查看具体的占用进程
```
tasklist|findstr "1234"
```
进入任务管理器
```
taskmgr
```

### Git
1. 以简短的形式显示版本记录：
git log --pretty=oneline

2. 版本使用的介绍

前多少个版本：
HEAD^  版本一
HEAD^^  版本二
windows一般用下面的
HEAD~1  版本一
HEAD~100 版本一百

3. 查看版本号：
git reflog 可以看到操作记录

git log --graph --pretty=oneline