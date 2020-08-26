## 什么是Socket

Socket，英文含义是【插座、插孔】，一般称之为**套接字**，用于描述IP地址和端口。可以实现不同程序间的数据通信。

Socket起源于Unix，而Unix基本哲学之一就是“一切皆文件”，都可以用“打开open –> 读写write/read –> 关闭close”模式来操作。Socket就是该模式的一个实现，网络的Socket数据传输是一种特殊的I/O，Socket也是一种文件描述符。Socket也具有一个类似于打开文件的函数调用：Socket()，该函数返回一个整型的Socket描述符，随后的连接建立、数据传输等操作都是通过该Socket实现的。

套接字的内核实现较为复杂，不宜在学习初期深入学习，了解到如下结构足矣。

![套接字通讯原理示意](https://i.loli.net/2020/04/26/wLBFEm41iSGu2Up.png)

在TCP/IP协议中，“IP地址+TCP或UDP端口号”唯一标识网络通讯中的一个进程。“IP地址+端口号”就对应一个socket。欲建立连接的两个进程各自有一个socket来标识，那么这两个socket组成的socket pair就唯一标识一个连接。因此可以用Socket来描述网络连接的一对一关系。

常用的Socket类型有两种：流式Socket（SOCK_STREAM）和数据报式Socket（SOCK_DGRAM）。流式是一种面向连接的Socket，针对于面向连接的TCP服务应用；数据报式Socket是一种无连接的Socket，对应于无连接的UDP服务应用。

## 网络应用程序设计模式

### C/S模式

传统的网络应用设计模式，客户机(client)/服务器(server)模式。需要在通讯两端各自部署客户机和服务器来完成数据通信。

### B/S模式

浏览器(Browser)/服务器(Server)模式。只需在一端部署服务器，而另外一端使用每台PC都默认配置的浏览器即可完成数据的传输。

### 优缺点

对于C/S模式来说，其优点明显。客户端位于目标主机上可以保证性能，将数据缓存至客户端本地，从而**提高数据传输效率**。而且一般来说客户端和服务器程序由一个开发团队创作，所以他们之间**所采用的协议相对灵活**。可以在标准协议的基础上根据需求裁剪及定制。例如，腾讯所采用的通信协议，即为ftp协议的修改剪裁版。

因此，传统的网络应用程序及较大型的网络应用程序都首选C/S模式进行开发。如，知名的网络游戏魔兽世界。3D画面，数据量庞大，使用C/S模式可以提前在本地进行大量数据的缓存处理，从而提高观感。

C/S模式的缺点也较突出。由于客户端和服务器都需要有一个开发团队来完成开发。**工作量将成倍提升**，开发周期较长。另外，从用户角度出发，需要将客户端安插至用户主机上，**对用户主机的安全性构成威胁**。这也是很多用户不愿使用C/S模式应用程序的重要原因。

B/S模式相比C/S模式而言，由于它没有独立的客户端，使用标准浏览器作为客户端，其**工作开发量较小**。只需开发服务器端即可。另外由于其采用浏览器显示数据，因此移植性非常好，**不受平台限制**。如早期的偷菜游戏，在各个平台上都可以完美运行。

B/S模式的缺点也较明显。由于使用第三方浏览器，因此**网络应用支持受限**。另外，没有客户端放到对方主机上，**缓存数据不尽如人意**，从而传输数据量受到限制。应用的观感大打折扣。第三，必须与浏览器一样，采用**标准http协议**进行通信，**协议选择不灵活。**

因此在开发过程中，模式的选择由上述各自的特点决定。根据实际需求选择应用程序设计模式。

## TCP的C/S架构


## 简单的C/S模型通信

### Server端

**Listen函数**

```go
func Listen(network, address string) (Listener, error)

network: 选用的协议: TCP、UDP, 如: "tcp" 或 "udp"
address: IP地址+端口号, 如: "127.0.0.1:8000" 或 ":8000"
```

**Listener接口**

```go
type Listener interface {
    Accept() (Conn, error)
    Close() error
    Addr() Addr
}
```

**Conn接口**

```go
type Conn interface {
    Read(b []byte) (n int, err error)
    Write(b []byte) (n int, err error)
    Close() error
    LocalAddr() Addr
    RemoteAddr() Addr
    SetDeadline(t time.Time) error
    SetReadDeadline(t time.Time) error
    SetWriteDeadline(t time.Time) error
}
```

示例代码:

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建监听: 指定 服务器通信协议, IP地址, 端口号
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.listen error: ", err)
		return
	}
	defer listener.Close()  // 主go程结束时, 关闭listener

	fmt.Println("服务器等待与客户端建立连接...")
	// 阻塞监听客户端连接请求
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept error: ", err)
		return
	}
	defer conn.Close()  // 使用结束, 断开与客户端连接
	fmt.Println("服务器与客户端连接成功!")

	// 读取客户端发送的数据
	buf := make([]byte, 4096)  // 创建4096大小的缓冲区, 用于read
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read error: ", err)
		return
	}

	// 处理数据 --> 打印
	fmt.Println("服务器读到数据: ", string(buf[:n]))  // 读多少, 打印多少
}
```

**大体步骤**

1. 创建监听socket: listener := net.Listen("tcp", "IP:port")  -- 服务器自己的IP和port
2. 启动监听: conn := listener.Accept(), conn 用于通信的socket
3. conn.Read()
4. 处理使用数据
5. conn.Write()
6. 关闭 listener、conn


如图，在整个通信过程中，服务器端有两个socket参与进来，**但用于通信的只有 conn 这个socket**。它是由 `listener` 创建的，隶属于服务器端。

![image.png](https://i.loli.net/2020/04/26/1LPaA9DyuVZEYhd.png)

我们用 `netcat` 来测试一下效果，具体的 `netcat` 配置请到这里查看：

结果：

![image](https://i.loli.net/2020/04/26/Bdrpi75hLMCKqmR.png)

### Client端

**Dial函数**

```go
func Dial(network, address string) (Conn, error)

network: 选用的协议: TCP、UDP. 如: "tcp"或者"udp" (不要大写)
address: 服务器的IP地址+端口号. 如: "104.28.0.106:8000" 或者 "www.lpgit.com:8000"
```

**Conn接口**

```go
type Conn interface {
    Read(b []byte) (n int, err error)
    Write(b []byte) (n int, err error)
    Close() error
    LocalAddr() Addr
    RemoteAddr() Addr
    SetDeadline(t time.Time) error
    SetReadDeadline(t time.Time) error
    SetWriteDeadline(t time.Time) error
}
```

示例代码:

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	// 指定服务器的 IP+port 创建通信socket(套接字)
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return
	}
	defer conn.Close()  // 结束时, 关闭连接

	// 主动写给服务器数据
	conn.Write([]byte("Hello Socket"))

	// 接收服务器回发的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read error: ", err)
		return
	}
	fmt.Println("服务器回发: ", string(buf[:n]))
}
```

**大体步骤:**

1. conn, err := net.Dial("tcp", "服务器的IP:port")
2. 写数据给服务器: conn.Write()
3. 读取服务器回发的数据: conn.Read()
4. conn.Close()

## 并发的C/S模型通信

### 并发Server

现在已经完成了客户端与服务端的通信，但是服务端只能接收一个用户发送过来的数据，怎样接收多个客户端发送过来的数据，实现一个高效的并发服务器呢？

Accept()函数的作用是**等待客户端的连接**，如果客户端没有连接，该方法会阻塞。如果有客户端连接，那么该方法返回一个Socket负责与客户端进行通信。所以，每来一个客户端，该方法就应该返回一个Socket与其通信，因此，可以使用一个死循环，将Accept()调用过程包裹起来。

需要注意的是，实现并发处理多个客户端数据的服务器，就需要**针对每一个客户端连接，单独产生一个Socket，并创建一个单独的goroutine与之完成通信。**

```go
package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("net.Listener error: ", err)
		return
	}
	defer listener.Close()

	// 监听客户端连接请求, 接收多个用户
	for {
		fmt.Println("服务器等待客户端连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error: ", err)
			return
		}
		// 具体完成服务器和客户端的数据通信
		go HandleConnect(conn)
	}
}
```

将客户端的数据处理工作封装到 `HandleConnect` 方法中，需将Accept()返回的Socket传递给该方法，变量conn的类型为：net.Conn。可以使用 `conn.RemoteAddr()` 来获取成功与服务器建立连接的客户端IP地址和端口号：

```go
// 获取客户端的网络地址信息
addr := conn.RemoteAddr().String()
```

客户端可能持续不断的发送数据, 因此接收数据的过程可以放在for循环中, 服务端也持续不断的向客户端返回处理后的数据.

添加一个限定，如果客户端发送一个“exit”字符串，表示客户端通知服务器不再向服务端发送数据，此时应该结束HandleConnect方法，同时关闭与该客户端关联的Socket。

```go
// 循环读取客户端发送数据
buf := make([]byte, 4096)
for {
	n, err := conn.Read(buf)
	fmt.Println(buf[:n])
	if string(buf[:n]) == "exit\n" || string(buf[:n]) == "exit\r\n" {
		fmt.Println("服务器接收到客户端", addr, "退出请求, 关闭.")
		return
	}
	if n == 0 {
		fmt.Println("服务器检测到客户端", addr, "已关闭, 断开连接.")
		return
	}
	if err != nil {
		fmt.Println("conn.Read error: ", err)
		return
	}
	fmt.Println("服务器读到", addr, "的数据: ", string(buf[:n]))

	// 小写转大写, 回发给客户端
	conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
}
```

在上面的代码中，Read()方法获取客户端发送过来的数据，填充到切片buf中，返回的是实际填充的数据的长度，所以将客户端发送过来的数据进行打印。

在判断客户端数据是否为“exit”字符串时，要注意，在Windows环境下客户端会自动的多发送2个字符："\r\n"（这在windows系统下代表回车、换行）; 在Linux环境下会自动多发送一个字符: "\n".

Server使用Write方法将数据写回给客户端，参数类型是 []byte，需使用strings包下的ToUpper函数来完成大小写转换。转换的对象即为string(buf[:n])

综上，HandleConnect方法完整定义如下：

```go
func HandleConnect(conn net.Conn) {
	// 函数调用完毕, 自动关闭conn
	defer conn.Close()
	
	// 获取连接的客户端 Addr
	addr := conn.RemoteAddr().String()
	fmt.Println("客户端连接成功", addr)

	// 循环读取客户端发送数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		//fmt.Println(buf[:n])
		if string(buf[:n]) == "exit\n" || string(buf[:n]) == "exit\r\n" {
			fmt.Println("服务器接收到客户端", addr, "退出请求, 关闭.")
			return
		}
		if n == 0 {
			fmt.Println("服务器检测到客户端", addr, "已关闭, 断开连接.")
			return
		}
		if err != nil {
			fmt.Println("conn.Read error: ", err)
			return
		}
		fmt.Println("服务器读到", addr, "的数据: ", string(buf[:n]))

		// 小写转大写, 回发给客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
```

**大体步骤:**

1. 创建监听套接字: listener, err := net.Listen("tcp", "IP:port")  // IP不能大写
2. defer listener.Close()
3. 



### 并发Client





## TCP通信











## UDP通信









### UDP服务器









### UDP客户端







### 并发UDP服务器







### 并发UDP客户端







## TCP与UDP的差异









## 文件传输

**命令行参数**：在main函数启动时，向整个程序传参。

例如：

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args  // 获取命令行参数
	fmt.Println(list)
}
```

```go
go run xxx.go argv1 argv2 argv3
	xxx.go：第0个参数
	argv1：第1个参数
	argv2：第2个参数
	argv3：第3个参数
```

**获取文件属性**: 

```go
fileInfo, err := os.Stat(文件访问绝对路径)
fileInfo接口: 
    Name()：获取文件名
    Size()：获取文件大小
```

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args  // 获取命令行参数，存入 list 中
	if len(list) != 2 {  // 确保用户输入一个命令行参数
		fmt.Println("格式为: go run xxx.go fileName")
		return
	}

	path := list[1]  // 从命令行保存文件名（含路径）
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat error: ", err.Error())
		return
	}
	fmt.Println("文件名", fileInfo.Name())  // 得到文件名（不含路径）
	fmt.Println("文件大小", fileInfo.Size())
}
```

### 文件传输流程简析

借助TCP完成文件的传输，基本思路如下：

1：发送方（客户端）获取文件名（不包含路径）

2：发送方（客户端）与服务器建立连接（Dial）

3：发送方（客户端）向服务端发送文件名，服务端保存该文件名。

2：接收方（服务端）向客户端返回一个消息ok，确认文件名保存成功。

3：发送方（客户端）收到消息后，开始向服务端发送文件数据。

4：接收方（服务端）读取文件内容，写入到之前保存好的文件中。

![image.png](https://i.loli.net/2020/05/01/OY3aGLhTxFPlItb.png)

首先获取文件名，刚刚已经讲过，借助 os 包中的 `Stat()` 函数来获取文件属性信息。在函数返回的文件属性中包含文件名和文件大小。Stat 参数 name 传入的是文件访问的绝对路径。FileInfo 中的 `Name()` 函数可以将文件名单独提取出来。

### 发送端实现

发送端(客户端)实现流程大致如下：

1：提示用户输入文件名，接收文件名path（含访问路径）

2：使用 os.Stat() 获取文件属性，得到纯文件名（不含访问路径）

3：主动连接服务器，结束时关闭连接

4：给接收端（服务器）发送文件名：conn.Write()

5：读取接收端回发的确认数据：conn.Read()

6：判断是否是 `ok`。如果是，封装函数 `SendFile()` 发送文件内容，传参 `path` 和 `conn`

7：只读 Open 文件，结束时 Close 文件

8：循环读文件，读到 **EOF** 终止文件读取

9：将读到的内容原封不动 Write 给接收端（服务器）

**代码实现**

```go
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(conn net.Conn, filePath string) {
	// 只读打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open error: ", err.Error())
		return
	}
	defer f.Close()

	// 从本地文件中读数据, 写给接收端. 读多少写多少
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取文件完毕")
			} else {
				fmt.Println("f.Read error: ", err.Error())
			}
			return
		}

		// 写到接收端
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("conn.Write error: ", err.Error())
			return
		}
	}
}

func main() {
	list := os.Args  // 获取命令行参数
	if len(list) != 2 {
		fmt.Println("格式为: go run xxx.go 文件绝对路径")
		return
	}
	// 提取 文件的绝对路径
	filePath := list[1]
	// 提取 文件名
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("os.Stat error: ", err.Error())
		return
	}
	fileName := fileInfo.Name()

	// 发送方(客户端) 主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8006")
	if err != nil {
		fmt.Println("net Dial error: ", err.Error())
		return
	}
	defer conn.Close()

	// 发送文件名给接收方(服务器)
	_, err = conn.Write([]byte(fileName))
	if err != nil {
		fmt.Println("conn.Write error: ", err.Error())
		return
	}

	// 读取服务器回发的 ok
	buf := make([]byte, 16)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read error: ", err.Error())
		return
	}
	//如果是 ok, 写文件内容给服务器
	if "ok" == string(buf[:n]) {
		sendFile(conn, filePath)
	}
}
```

### 接收端实现

接收端(服务器)实现流程大致如下：

1：创建监听 listener，程序结束时关闭

2：阻塞等待发送端连接 Accept，程序结束时关闭 conn

3：读取发送端发送文件名，保存 fileName

4：回发 `ok` 给发送端做应答

5：封装函数 `RecvFile` 接收发送端发送的文件内容，传参 `fileName`  和 `conn`

6：按照文件名 Create 文件，结束时  Close

7：循环 Read 发送端发送的文件内容，当读到 **0** 时说明文件读取完毕

8：将读到的内容原封不动 Write 到创建的文件中

**代码实现**

```go
package main

import (
	"fmt"
	"net"
	"os"
)

func recvFile(conn net.Conn, fileName string){
	// 按照文件名创建新文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create error: ", err.Error())
		return
	}
	defer f.Close()

	// 从 网络中读数据, 写入本地文件
	buf := make([]byte, 4096)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("文件接收完成")
			return
		}
		// 写入本地文件, 读多少, 写多少
		f.Write(buf[:n])
	}
}

func main() {
	// 创建用于监听的Socket
	listener, err := net.Listen("tcp", "127.0.0.1:8006")
	if err != nil {
		fmt.Println("net.Listen error: ", err.Error())
		return
	}
	defer listener.Close()

	// 阻塞监听
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept error: ", err.Error())
		return
	}
	defer conn.Close()

	// 获取文件名, 保存
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read error: ", err.Error())
		return
	}
	fileName := string(buf[:n])

	// 回写 ok 给发送端
	_, err = conn.Write([]byte("ok"))
	if err != nil {
		fmt.Println("conn.Write error: ", err.Error())
		return
	}

	// 获取文件内容
	recvFile(conn, fileName)
}
```




## 李培冠博客

欢迎访问我的个人网站：

李培冠博客：[lpgit.com](https://lpgit.com)
