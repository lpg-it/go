## 前言

前面我已经讲过 GoLand 的安装，当然，你也可以使用 `VS Code` 来进行开发。

VS Code 是微软开源的一款编辑器, 本文主要介绍如何使用VS Code搭建Go语言的开发环境.

## 下载与安装VS Code
官方下载地址: https://code.visualstudio.com/Download

![](https://i.loli.net/2019/07/20/5d32efdcad62b69752.png)

双击下载好的安装文件, 安装即可

## 安装中文简体插件

点击左侧菜单栏最后一项 `管理扩展`, 在 `搜索框` 中输入 `chinese`, 选中结果列表第一项, 点击 `install` 安装.

安装完毕后右下角会提示 `重启VS Code`, 重启之后VS Code就显示中文了.

## 安装Go开发环境

点击左侧菜单栏最后一项 `管理扩展`, 在 `搜索框` 中输入 `go`, 选中结果列表第一项, 点击 `安装` 进行安装.

![](https://i.loli.net/2019/07/20/5d32f12d21b0579091.png)

## 安装Go语言开发工具包

在进行Go语言开发的时候, 安装下面这些插件可以为我们提供如代码提示、代码自动补全等功能.

Windows平台下按 `Ctrl+Shift+p`, Mac平台下按 `Command+Shift++p`, VS Code界面会弹出一个输入框, 在输入框中输入 `go:install`, 选择 `Go:install/Update Tools` 如下图:

![](https://i.loli.net/2019/07/20/5d32f2f617a5221082.png)

选中点击 `确定`, 进行安装.

![](https://i.loli.net/2019/07/20/5d32f33ff2d5411232.png)



![](https://i.loli.net/2019/07/20/5d32ee83916f830520.png)

由于国内网络环境基本上都会出现安装失败, 如下图: 

![](https://i.loli.net/2019/07/20/5d32f2261a76581735.png)

解决方案:
手动从github上下载工具. (前提需要电脑上已经安装了git)

> 1. 先在自己的 `GOPATH` 的 `src` 目录下创建 `golang.org/x` 目录
> 2. 在该目录下打开`终端` / `cmd`
> 3. 执行 `git clone https://github.com/golang/tools.git tools` 命令
> 4. 执行 `git clone https://github.com/golang/lint.git` 命令
> 5. 在VS Code按下 `Ctrl / Command+Shift+p` 再次执行 `go:install/Update Tools` 命令, 在弹出的窗口中全选并点击确定, 这次安装就可以全部安装成功了.

经过上面的步骤就可以安装成功了.

## 李培冠博客

[lpgit.com](https://lpgit.com)

