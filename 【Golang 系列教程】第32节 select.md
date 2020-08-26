## select的作用

Go里面提供了一个关键字 `select`, 通过 `select` 可以监听channel上的数据流动.

`select` 的用法与 `switch` 语言非常类似, 由 `select` 开始一个新的选择块, 每个选择条件由 `case` 语句来描述.

与 `switch` 语句相比, `select` 有比较多的限制, 其中最大的一条限制就是**每个case语句里必须是一个IO操作.**

大致的结构如下:

```go
select {
case <- chan1:
	// 如果chan1成功读到数据, 则进行该case处理语句 
case chan2 <- -1:
	// 如果成功向chan2写入数据, 则进行该case处理语句
default:
	// 如果上面都没有成功, 则进入default处理流程
}
```

在一个 `select` 语句中, Go语言会按照顺序从头至尾评估每一个发送和接收的语句.

如果其中的**任意**一条语句可以继续执行(即没有阻塞), 那么就从那些可以执行的语句中**任意**选择一条来使用.

如果没有任意一条语句可以执行(即所有的通道都被阻塞), 那么有两种可能的情况: 

- 如果给出了default语句, 那么就会执行default语句, 同时程序的执行会从select语句后的语句中恢复.
- 如果没有default语句, 那么select语句将被阻塞, 直到至少有一个通信可以进行下去.

## select的基本使用

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)  // 用来进行数据通信的channel
	quit := make(chan bool)  // 用来判断是否退出的channel

	go func() {  // 写数据
		for i:=0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true  // 通知主go程 退出
		runtime.Goexit()
	}()

	for {
		select {
		case num := <- ch:
			fmt.Println("读到: ", num)
		case <- quit:
			return
			//break  // break 跳出select循环
		}
		fmt.Println("============")
	}
}
```

结果:

```go
读到:  0
============
读到:  1
============
读到:  2
============
读到:  3
============
读到:  4
============
读到:  0
============
读到:  0
============
```

**注意, 因为是任意挑选一个case执行, 所以最后的 读到:0 的数量相当于是个随机数.**

所以, 总结下select的注意事项:

- case后面必须是IO操作, 不可以是判别表达式.
- 监听的case中, 没有满足监听条件, 阻塞.
- 监听的case中, 有多个满足监听条件, **任选**一个执行.
- 可以使用default来处理所有case都不满足监听条件的状况.(通常不用, 会产生 忙轮询)
- select自身不带有循环机制, 需借助外层for循环来进行循环监听
- break只能跳出select. 类似于switch中的用法.

## select实现斐波那契数列

```go
package main

import (
	"fmt"
	"runtime"
)

func fibonacci(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Println(num)
		case <-quit:
			//return
			runtime.Goexit()
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go fibonacci(ch, quit)

	x, y := 1, 1
	for i := 0; i < 50; i++ {
		ch <- x
		x, y = y, x+y
	}
	quit <- true
}

```

## 为什么要用到select?

如果不用select的话, 每一个case都要创建一个go程去处理, 这样的话太浪费了, 而用select的话, 只需要一个go程就可以了.

## 超时

有时候会出现goroutine阻塞的情况, 那么我们如何避免整个程序进入阻塞的情况呢?我们可以利用select来设置超时, 通过如下的方式来实现:

示例代码:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timeOut := make(chan bool)

	go func() {
		for {
			select {
			case num := <- ch:
				fmt.Println("num: ", num)
			case <- time.After(5 * time.Second):
				fmt.Println("timeout")
				timeOut <- true
				return
			}
		}
	}()
	ch <- 666
	<- timeOut  // 主go程, 阻塞等待子go程通知, 退出
	fmt.Println("finish.")
}
```

select监听time.After() 中channel的读事件, 如果定时时间到, 系统会向该channel中写入系统当前时间.

## 李培冠博客

欢迎访问我的个人网站：

李培冠博客：[lpgit.com](https://lpgit.com)
