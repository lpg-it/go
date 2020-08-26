Go语言中除了可以使用通道（channel）和互斥锁进行两个并发程序间的同步外，还可以使用等待组进行多个任务的**同步**，**等待组可以保证在并发环境中完成指定数量的任务**

在 sync.WaitGroup（等待组）类型中，每个 sync.WaitGroup 值在内部维护着一个计数，此计数的初始默认值为零。

等待组有下面几个方法可用，如下所示。

- func (wg *WaitGroup) Add(delta int)： 等待组的计数器 +1
- func (wg *WaitGroup) Done()： 等待组的计数器 -1
- func (wg *WaitGroup) Wait()： 当等待组计数器不等于 0 时阻塞直到变 0。

对于一个可寻址的 sync.WaitGroup 值 wg：

- 我们可以使用方法调用 wg.Add(delta) 来改变值 wg 维护的计数。
- 方法调用 wg.Done() 和 wg.Add(-1) 是完全等价的。
- 如果一个 wg.Add(delta) 或者 wg.Done() 调用将 wg 维护的计数更改成一个负数，将会产生 panic 异常。
- 当一个协程调用了 wg.Wait() 时，
    - 如果此时 wg 维护的计数为零，则此 wg.Wait() 此操作为一个空操作（noop）；
    - 否则（计数为一个正整数），此协程将进入阻塞状态。当以后其它某个协程将此计数更改至 0 时（一般通过调用 wg.Done()），此协程将重新进入运行状态（即 wg.Wait() 将返回）。

等待组内部拥有一个计数器，计数器的值可以通过方法调用实现计数器的增加和减少。当我们添加了 N 个并发任务进行工作时，就将等待组的计数器值增加 N。每个任务完成时，这个值减 1。同时，在另外一个 goroutine 中等待这个等待组的计数器值为 0 时，表示所有任务已经完成。

什么意思？我们先来回忆一下之前我们为了保证子 go 程运行完毕，主 go 程是怎么做的：

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    go func() {
        fmt.Println("Goroutine 1")
    }()

    go func() {
        fmt.Println("Goroutine 2")
    }()

    time.Sleep(time.Second) // 睡眠 1 秒，等待上面两个子 go 程结束
}
```

我们为了让子 go 程可以顺序的执行完，在主 go 程中加入了等待。我们知道，这不是一个很好的解决方案，可以用 channel 来实现同步：

```go
package main

import (
    "fmt"
)

func main() {

    ch := make(chan struct{})
    count := 2 // count 表示活动的 go 程个数

    go func() {
        fmt.Println("Goroutine 1")
        ch <- struct{}{} // go 程结束，发出信号
    }()

    go func() {
        fmt.Println("Goroutine 2")
        ch <- struct{}{} // go 程结束，发出信号
    }()

    for range ch {
        // 每次从 ch 中接收数据，表明一个活动的 go 程结束
        count--
        // 当所有活动的 go 程都结束时，关闭 channel
        if count == 0 {
            close(ch)
        }
    }
}
```

上面的解决方案是虽然已经比较好了，但是 Go 提供了更简单的方法：使用 `sync.WaitGroup`。

```go
package main

import (
    "fmt"
    "sync"
)

func main() {

    var wg sync.WaitGroup

    wg.Add(2) // 因为有两个动作，所以增加 2 个计数
    go func() {
        fmt.Println("Goroutine 1")
        wg.Done() // 操作完成，减少一个计数
    }()

    go func() {
        fmt.Println("Goroutine 2")
        wg.Done() // 操作完成，减少一个计数
    }()

    wg.Wait() // 等待，直到计数为0
}
```

可见用 `sync.WaitGroup` 是最简单的方式。

**强调一下：**

- 计数器不能为负值：不能使用 `Add()` 或者 `Done()` 给 wg 设置一个负值，否则代码将会报错。
- WaitGroup 对象不是一个引用类型：在通过函数传值的时候需要使用地址。

官方文档看这里：[https://golang.org/pkg/sync/#WaitGroup](https://golang.org/pkg/sync/#WaitGroup)

## 练习题

1、写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- rand.Int()
		}
		close(ch)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println("num = ", num)
		}
	}()
	wg.Wait()
}
```

## 李培冠博客

欢迎访问我的个人网站：

李培冠博客：[lpgit.com](https://lpgit.com)
