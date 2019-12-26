# golang 学习笔记

这么火了都,没理由不学啊,

## 环境安装

下载安装包 , 无脑下一步

<https://studygolang.com/dl>

安装完成后命令行敲入

```shell

go version

```

![go-version](/image/go-version.png)

## 设置依赖代理

```

go env -w GOPROXY=https://goproxy.cn,direct

```

说明文档:

<https://github.com/goproxy/goproxy.cn/blob/master/README.zh-CN.md>

## vscode 安装 golang 插件

https://code.visualstudio.com/

![go](/image/vscode-go-plug.png)

然后打开 vscode 写一个 hello.go ,然后 vscode 右下角会提示你安装扩展 , 点击 ALL Install 然后登上一会儿就好了

## 简单的指令

仓库目录下执行 `go run hello.go` 就完了

执行 `go build` 则可以编译成二进制文件

## 相关学习资料

https://studygolang.com/


