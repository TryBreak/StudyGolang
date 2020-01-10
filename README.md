# golang 学习笔记

这么火了都,没理由不学啊

标准库文档 <https://studygolang.com/pkgdoc>

## windows 下环境安装

下载安装包 , 无脑下一步

<https://studygolang.com/dl>

安装完成后命令行敲入检测环境变量和安装是否成功

```shell

go version

```

![go-version](/image/go-version.png)

## Linux 下环境安装

<https://studygolang.com/dl>

1. 下载 linux 包

2. 然后解压

3. 配置环境变量等：

```bash

sudo vim /etc/profile

## 在最后添加以下内容

export GOROOT=/opt/go  #指向你的安装目录
export GOPATH=~/golib:~/goProject
export GOBIN=~/gobin
export PATH=$PATH:$GOROOT/bin:$GOBIN

```

## 设置依赖代理

```

go env -w GOPROXY=https://goproxy.cn,direct

```

依赖代理的说明文档:

<https://github.com/goproxy/goproxy.cn/blob/master/README.zh-CN.md>

## 编辑器配置

宇宙无敌 IDE 的亲儿子
vscode 下载地址:

https://code.visualstudio.com/

插件安装(认准 Microsoft)

![go](/image/vscode-go-plug.png)

然后打开 vscode 新建一个文件夹, 写一个 hello.go ,然后 vscode 右下角会提示你安装扩展 , 点击 ALL Install 然后去泡杯咖啡等待提示全部安装成功就好了

> 如果提示失败, 就重新设置依赖代理 , 再来一遍

## 简单的指令

仓库目录下执行 `go run ./hello.go` 就完了

执行 `go build ./hello.go` 则可以编译成二进制文件

语法:

```bash

  go run <path>

  go build <path>

  # `<path>` 为文件路径, 关于相对路径和绝对路径的概念请自行使用搜索引擎

```

> go 支持交叉编译 , 即:windows 下编译 linux 运行的程序,或 linux 编译成 windows 运行的程序

项管详细说明请搜索 `go 交叉编译`相关的文档即可

## 相关学习资料

https://studygolang.com/

## 作者

QQ : 670188307 , 欢迎交流
