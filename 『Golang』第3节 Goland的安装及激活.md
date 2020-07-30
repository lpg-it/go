## 前言

前面我们已经将GO的环境安装好了，那么是否可以进行开发了呢？

可以，但是为了能够更高效率的开发，我们还需要下载一个软件，该软件的作用就是方便我们能够快速的编写GO指令，快速的运行我们编写好的GO指令。

这个软件就是 GoLand ，就像我们要处理文字安装 Word ,处理表格用 Excel 等等。

我们把这种用来能够用来快速编写某种语言（GO,Python,JAVA,C#）指令，快速运行，同时如果出错可以方便我们查找错误（排错）的软件就称为IDE.

## IDE是什么

IDE(Integrated Development,集成开发环境)，我们GO语言在Windows下用到的IDE是什么呢？GoLand是一个跨平台的IDE，使用范围包括Windows，maxOS以及linux操作系统。

## Windows下安装GoLand

1. 登录`JetBrains`官网, 下载`GoLand`安装程序, 这里以`Windows`为例.

    官网下载地址: http://www.jetbrains.com/go

2. 下载完成后双击安装程序运行, 开始安装.

![](https://i.loli.net/2019/11/25/tvBf7r2Dh6PAbXj.png)

3. 点击`Next`, 选择好自己的安装位置.

![](https://i.loli.net/2019/11/25/ZtbySnjeX5kBFPM.png)

4. 点击`Next`, 选择要创建桌面快捷方式, 并关联`.go`文件, 并添加到`PATH`.

![](https://i.loli.net/2019/11/25/uvOiDSZfLdGXhk1.png)

5. 后面一路`Next`, 安装完成后点击启动即可.

6. 重启软件, 激活`GoLand`, 点击`Activation code`, 输入激活码.

## 激活GoLand（亲测有效）

1、 下载破解补丁

[https://lanzous.com/b09r1cfyj](https://lanzous.com/b09r1cfyj)

密码：giws

2、 将补丁放到 `lib` 目录下

将刚才下载的补丁文件 `jetbrains-agent.jar` 放置在 GoLand 安装目录里面的 `lib` 目录里面

3、 将补丁 **jetbrains-agent.jar** 放置到lib目录后，我们就可以启动软件了。

## Linux下安装Goland

1）首先进入下载页面：[https://www.jetbrains.com/go/download/](https://www.jetbrains.com/go/download/)

2）点击下载：

![image](https://i.loli.net/2020/05/20/Fz9qpU3ThrEdPYw.png)

3）`cd` 到刚刚下载的文件的路径下面，默认是在 `下载` 文件夹下面，并使用 `tar` 命令将 `GoLand` 安装包解压到 `/opt/` 路径下

```linux
cd
cd 下载
sudo tar xvfz goland-2020.1.3.tar.gz -C /opt/
```

![image](https://i.loli.net/2020/05/20/DFJuVgR26naUHPv.png)

4）进入到解压目录

```linux
cd /opt/GoLand-2020.1.3/bin/
```

5）运行 `golang.sh` 文件

```linux
./goland.sh
```

6）如果一切正常的话，就会弹出GoLand启动页面


### 破解同上

## 李培冠博客

[lpgit.com](https://lpgit.com)
