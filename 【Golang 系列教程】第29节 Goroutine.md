## 什么是 Goroutine

goroutine 是 Go 并行设计的核心。goroutine 说到底其实就是协程，它比线程更小，十几个 goroutine 可能体现在底层就是五六个线程，Go 语言内部帮你实现了这些 goroutine 之间的内存共享。

执行 goroutine 只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务。goroutine 比 thread 更易用、更高效、更轻便。

一般情况下，一个普通计算机跑几十个线程就有点负载过大了，但是同样的机器却可以轻松地让成百上千个 goroutine 进行资源竞争。

## Goroutine 的创建

**只需在函数调⽤语句前添加 go 关键字，就可创建并发执⾏单元。**

开发⼈员无需了解任何执⾏细节，调度器会自动将其安排到合适的系统线程上执行。

在并发编程中，我们通常想将一个过程切分成几块，然后让每个 goroutine 各自负责一块工作，当一个程序启动时，主函数在一个单独的 goroutine 中运行，我们叫它 main goroutine。新的 goroutine 会用 go 语句来创建。而 go 语言的并发设计，让我们很轻松就可以达成这一目的。

例如：

```go
package main

import (
	"fmt"
	"time"
)

func foo() {
	i := 0
	for true {
		i++
		fmt.Println("new goroutine: i = ", i)
		time.Sleep(time.Second)
	}
}

func main() {
	// 创建一个 goroutine, 启动另外一个任务
	go foo()

	i := 0
	for true {
		i++
		fmt.Println("main goroutine: i = ", i)
		time.Sleep(time.Second)
	}
}
```

结果：

```go
main goroutine: i =  1
new goroutine: i =  1
new goroutine: i =  2
main goroutine: i =  2
main goroutine: i =  3
new goroutine: i =  3
...
```

## Goroutine 特性

**主go程 退出后，其它的 子go程 也会自动退出：**

```go
package main

import (
	"fmt"
	"time"
)

func foo() {
	i := 0
	for true {
		i++
		fmt.Println("new goroutine: i = ", i)
		time.Sleep(time.Second)
	}
}

func main() {
	// 创建一个 goroutine, 启动另外一个任务
	go foo()

	time.Sleep(time.Second * 3)

	fmt.Println("main goroutine exit")
}
```

运行结果：

```go
new goroutine: i =  1
new goroutine: i =  2
new goroutine: i =  3
main goroutine exit
```

## runtime 包

### Gosched

`runtime.Gosched()` 用于出让当前 go程 所占用的 CPU 时间片，让出当前 goroutine 的执行权限，调度器安排其他等待的任务运行，并在下次再获得 cpu 时间轮片的时候，从该出让 cpu 的位置恢复执行。

有点像跑接力赛，A 跑了一会碰到代码 runtime.Gosched() 就把接力棒交给 B 了，A 歇着了，B 继续跑。

例如：

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 创建一个 goroutine
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("hello")
	}
	time.Sleep(time.Second * 3)
}
```

运行结果：

```go
world
world
hello
hello
```

如果没有 `runtime.Gosched()` 则运行结果如下：

```go
hello
hello
world
world
```

注意： `runtime.Gosched()` 只是出让一次机会，看下面的代码，注意运行结果：

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 创建一个 goroutine
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
			time.Sleep(time.Second)
		}
	}("world")

	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("hello")
	}
}
```

运行结果：

```go
world
hello
hello
```

为什么 `world` 只有一次呢？因为之前我们说过，**主 goroutine 退出后，其它的工作 goroutine 也会自动退出。**

### Goexit

调用 `runtime.Goexit()` 将立即终止**当前 goroutine** 执⾏，调度器确保所有已注册 defer 延迟调用被执行。

注意与 `return` 的区别，`return` 是返回**当前函数**调用给调用者。

例如：

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit() // 终止当前 goroutine
			fmt.Println("B") // 不会执行
		}()
		fmt.Println("A") // 不会执行
	}() // 不要忘记 ()

	time.Sleep(time.Second * 3)
}
```

运行结果：

```go
B.defer
A.defer
```

### GOMAXPROCS

调用 `runtime.GOMAXPROCS()` 用来设置可以并行计算的 CPU 核数的**最大值**，并返回 **上一次（没有则是电脑默认的）** 设置的值。

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)  // 将 cpu 设置为单核

	for true {
		go fmt.Print(0)  // 子 go 程
		fmt.Print(1)  // 主 go 程
	}
}
```

运行结果：

```go
111111 ... 1000000 ... 0111 ...
```

在执行 `runtime.GOMAXPROCS(1)` 时，最多同时只能有一个 goroutine 被执行。所以会打印很多 1。过了一段时间后，GO 调度器会将其置为休眠，并唤醒另一个 goroutine，这时候就开始打印很多 0 了，在打印的时候，goroutine 是被调度到操作系统线程上的。

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	for true {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
```

运行结果：

```go
111111111111111000000000000000111111111111111110000000000000000011111111100000...
```

在执行 `runtime.GOMAXPROCS(2)` 时， 我们使用了两个 CPU，所以两个 goroutine 可以一起被执行，以同样的频率交替打印 0 和 1。

### runtime 包中的其它函数

中文文档在这里：[https://studygolang.com/pkgdoc](https://studygolang.com/pkgdoc)

这里就简单列举一下一些函数以及功能。

```go
func GOROOT() string
```

GOROOT 返回 Go 的根目录。如果存在 GOROOT 环境变量，返回该变量的值；否则，返回创建 Go 时的根目录。

----

```go
func Version() string
```

返回 Go 的版本字符串。它要么是递交的 hash 和创建时的日期；要么是发行标签如 "go1.3"。

----

```go
func NumCPU() int
```

NumCPU返回本地机器的逻辑CPU个数（真 **·** 八核）。

----

```go
func GC()
```

GC执行一次垃圾回收。(如果你迫切的希望做一次垃圾回收，可以调用此函数)

----

其它的大家自行去文档查看吧~

## 李培冠博客

欢迎访问我的个人网站：

李培冠博客：[lpgit.com](https://lpgit.com)
