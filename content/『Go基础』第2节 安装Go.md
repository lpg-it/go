## 前言

前面已经对GO语言做了一个简单的介绍，相信大家也已经基本有所了解了，那么我相信大家迫不及待的想学习GO语言了，但是我们在具体学习GO语言前，需要先安装GO的开发环境，只有把GO的开发环境安装好了，才能有GO语言的指令来开发网站，飞秋等等软件。

所谓的环境安装，就是将GO语言的指令都安装在你的电脑上，那么你在开发软件的时候就可以直接用了。GO环境的安装分为windows系统的安装和linux系统的安装。

## 下载地址

Go语言官网下载地址: https://golang.org/dl/

Go语言镜像站: https://golang.google.cn/dl/

Windows平台和Mac平台推荐下载可执行文件版, Linux平台下载压缩文件版.

大家根据自己的操作系统来选择对应的版本.

![](https://i.loli.net/2019/11/22/5sBZMO4wAqvbFN8.png)

## Windows系统下安装Go

此次安装以`64位Windows10`系统安装`Go 1.13.4可执行文件版本`为例。

1. 打开下载好的安装包.

![](https://i.loli.net/2019/11/22/Lwc1JDyOkrIvdRm.png)

2. 点击 `next`, 继续点击`next`

![](https://i.loli.net/2019/11/22/FaRbC8kyPBYOZj4.png)

3. 选择Go语言的安装目录, 尽量选择比较容易记的（**安装路径中不要出现中文**）.

![](https://i.loli.net/2019/11/22/bD7udMQY5P8yNgC.png)

4. 安装.

![](https://i.loli.net/2019/11/22/t6IA1LCSKJX9eca.png)

5. 显示这个界面就表明安装成功.

![](https://i.loli.net/2019/11/22/HJnf2ctT5rEX61q.png)

6. 安装完成后, 可以打开终端窗口, 输入`go version`命令, 查看安装的Go版本.

![](https://i.loli.net/2019/11/22/Xwtnrf4y8qvclmz.png)

### 配置GOPATH

`GOPATH` 是一个环境变量, 用来表明你写的Go项目的存放路径. 

**注意**: 不是安装目录, 是工作目录, 写代码的目录.

`GOPATH` 路径最好只设置一个, 所有的项目代码都放到`GOPATH`的`src`目录下.

**注意**: 在`Go 1.11` 版本之后, 开启 `go mod` 模式之后就不再强制需要配置`GOPATH`了.

Windows平台下按照下面的步骤将 `E:\code\go` 添加到环境变量.

1. `我的电脑` --> `属性` --> `高级系统设置` --> `环境变量`

![](https://i.loli.net/2019/11/22/dMSlwrEJv6ynmN5.png)

![](https://i.loli.net/2019/11/22/7dXM1WZBpPHoRzl.png)

![](https://i.loli.net/2019/11/22/kdILnl19JtgWsDC.png)

2. 点击`系统变量`下的`新建`, 变量名写 `GOPATH`, 变量值写`保存Go代码的目录`, 我这里是`E:\code\go`, 点击确定.

![](https://i.loli.net/2019/11/22/ioEGy7thklDudwF.png)

![](https://i.loli.net/2019/11/22/xDoVfecFsh2duJz.png)

3. 点击`用户变量`下的`新建`, 变量名写 `GOPATH`, 变量值写`保存Go代码的目录`, 我这里是`E:\code\go`, 点击确定.

![](https://i.loli.net/2019/11/22/O8UQFxNTgvbWecn.png)

![](https://i.loli.net/2019/11/22/27ploQDRbiLPWAm.png)

4. 在`GOPATH`目录下新建三个文件夹. `bin`: 用来存放编译后生成的可执行文件. `pkg`: 用来存放编译后生成的归档文件. `src`: 用来存放源码文件.

![](https://i.loli.net/2019/11/22/p8OeyIs4BXTMhvK.png)

## Linux系统下安装Go

我这里以Ubuntu2004为例。

### 下载Go

```bash
wget https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz
```

**如果想要安装其它版本，只需要更改版本号就可以了。**

下载之后解压到 `/usr/local/`

```bash
sudo tar -C /usr/local -xzf go1.14.2.linux-amd64.tar.gz
```

### 配置环境变量

打开 `profile` 文件

```bash
sudo vi /etc/profile
```

打开之后添加如下三行：

```bash
export GOROOT=/usr/local/go
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
```

添加之后效果如下：

![image](https://i.loli.net/2020/05/11/4uzibFIHqV9D3Za.png)

**说明：**

- GOROOT就是刚才下载并解压之后得到的那个文件夹；
- PATH后面加上go命令所在路径，这样，在任意地方输入go都可以被识别；
- GOPATH这个根据个人情况设置，我推荐这个就设置为主文件夹下面的go文件夹（自己创建的），后面你从github上下的一些包，都要放在$GOPATH的src文件夹内，这样程序里才能找到这些包。

改完 `profile` 文件之后执行下面命令使之生效：

```bash
source /etc/profile
```

然后，输入 `go env` ，看到如下输出，解说明安装成功且配置生效。

![image](https://i.loli.net/2020/05/11/S7a4GoRxtNB8yuw.png)

## 李培冠博客

[lpgit.com](https://lpgit.com)
