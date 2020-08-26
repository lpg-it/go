## 前言

前面我们为了解决go程同步的问题我们使用了channel, 但是go也提供了传统的同步工具.

它们都在go的标准库代码包 `sync` 和 `sync/atomic` 中.

下面我们来看一下锁的应用.

什么是锁呢? 就是某个协程(线程)在访问某个资源时先锁住, 防止其他协程的访问, 等访问完毕解锁后其他协程再来加锁进行访问.

这和我们生活中加锁使用公共资源相似, 例如: 公共卫生间.

## 死锁

死锁是指两个或者两个以上的进程在执行过程中, 由于竞争资源或者由于彼此通信而造成的一种阻塞的现象, 若无外力作用, 它们都将无法推进下去. 此时称系统处于**死锁状态**或**系统产生了死锁**.

**死锁不是锁的一种! 它是一种错误使用锁导致的现象.**

### 产生死锁的几种情况

- 单go程自己死锁
- go程间channel访问顺序导致死锁
- 多go程, 多channel交叉死锁
- **将 互斥锁、读写锁与channel混用 -- 隐性死锁**(在 `读写锁` 讲到)



**单go程自己死锁** 示例代码:

```go
package main

import "fmt"

// 单go程自己死锁
func main() {
	ch := make(chan int)
	ch <- 789
	num := <- ch
	fmt.Println(num)
}
```

上面这段乍一看有可能会觉得没有什么问题, 可是仔细一看就会发现这个 `ch` 是一个无缓冲的channel, 当789写入缓冲区时, 这时**读端**还没有准备好. 所以, **写端** 会发生阻塞, 后面的代码不再运行.

所以可以得出一个结论: **channel应该在至少2个及以上的go程进行通信, 否则会造成死锁.**

我们继续看 **go程间channel访问顺序导致死锁** 的例子:

```go
package main

import "fmt"

// go程间channel访问顺序导致死锁
func main(){
	ch := make(chan int)
	num := <- ch
	fmt.Println("num = ", num)
	go func() {
		ch <- 789
	}()
}
```

在代码运行到 `num := <- ch` 时, 发生阻塞, 并且下面的代码不会执行, 所以发生死锁.

正确应该这样写:

```go
package main

import "fmt"

func main(){
	ch := make(chan int)
	go func() {
		ch <- 789
	}()
	num := <- ch
	fmt.Println("num = ", num)
}
```

所以, **在使用channel一端读(写)时, 要保证另一端写(读)操作有机会执行.**

我们再来看下 **多go程, 多channel交叉死锁** 的示例代码:

```go
package main

import "fmt"

// 多go程, 多channel交叉死锁
func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case num := <- ch1:
				ch2 <- num
			}
		}
	}()

	for {
		select {
		case num := <- ch2:
			ch1 <- num
		}
	}
}
```

## 互斥锁(互斥量)

每个资源都对应于一个可称为"互斥锁"的标记, 这个标记用来保证在任意时刻, 只能有一个协程(线程)访问该资源, 其它的协程只能等待.

互斥锁是传统并发编程对共享资源进行访问控制的主要手段, 它由标准库 `sync` 中的 `Mutex` 结构体类型表示.

`sync.Mutex` 类型只有两个公开的指针方法, **Lock** 和 **Unlock**. 

Lock锁定当前的共享资源, Unlock进行解锁.

在使用互斥锁时, 一定要注意, 对资源操作完成后, 一定要解锁, 否则会出现流程执行异常, 死锁等问题, 通常借助defer. 锁定后, 立即使用 `defer` 语句保证互斥锁及时解锁. 如下所示:

```go
var mutex sync.Mutex  // 定义互斥锁变量: mutex

func write() {
    mutex.Lock()
    defer mutex.Unlock()
}
```

我们先来回顾一下channel是怎么样完成数据同步的.

```go
package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func printer(str string) {
	for _, s := range str {
		fmt.Printf("%c ", s)
		time.Sleep(time.Millisecond * 300)
	}
}

func person1() {        // 先
	printer("hello")
	ch <- 666
}

func person2() {        // 后
	<-ch
	printer("world")
}

func main() {
	go person1()
	go person2()
	time.Sleep(5 * time.Second)
}
```

同样可以使用互斥锁来解决, 如下所示:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用传统的 "锁" 完成同步  -- 互斥锁
var mutex sync.Mutex  // 创建一个互斥锁(互斥量), 新建的互斥锁状态为0 -> 未加锁状态. 锁只有一把.
func printer(str string) {
	mutex.Lock()        // 访问共享数据之前, 加锁
	for _, s := range str {
		fmt.Printf("%c ", s)
		time.Sleep(time.Millisecond * 300)
	}
	mutex.Unlock()  // 共享数据访问结束, 解锁
}

func person1() {
	printer("hello")
}

func person2() {
	printer("world")
}

func main() {
	go person1()
	go person2()
	time.Sleep(5 * time.Second)
}
```

这种锁为**建议锁**: 操作系统提供, 建议你在编程时使用.

**强制锁**只会在底层操作系统自己用到, 我们在写代码时用不到.

person1与person2两个go程共同访问共享数据, 由于CPU调度随机, 需要对 **共享数据访问顺序加以限定(同步).**

创建mutex(互斥锁), 访问共享数据之前, 加锁; 访问结束, 解锁.

在person1的go程加锁期间, person2的go程加锁会失败 --> 阻塞.

直至person1的go程解锁mutext, person2从阻塞处, 恢复执行.

## 读写锁

互斥锁的本质是当一个goroutine访问的时候, 其它goroutine都不能访问. 这样在资源同步, 避免竞争的同时, 也降低了程序的并发性能, 程序由原来的并行执行变成了串行执行.

其实, 当我们对一个不会变化的数据只做**读**操作的话, 是不存在资源竞争的问题的. 因为数据是不变的, 不管怎么读取, 多少goroutine同时读取, 都是可以的.

所以问题不是出在**读**上, 主要是修改, 也就是**写**. 修改的数据要同步, 这样其它goroutine才可以感知到. 所以真正的互斥应该是读取和修改、修改和修改之间, **读和读是没有互斥操作的必要的.**

因此, 衍生出另外一种锁, 叫做**读写锁.**

读写锁可以让多个读操作并发, 同时读取, 但是对于写操作是完全互斥的. 也就是说, 当一个goroutine进行写操作的时候, 其它goroutine既不能进行读操作, 也不能进行写操作.

Go中的读写锁由结构体类型 `sync.RWMutex` 表示. 此类型的方法集合中包含两对方法:

一组是对写操作的锁定和解锁, 简称为: **写锁定** 和 **写解锁.**

```go
func (*RWMutex) Lock()
func (*RWMutex) Unlock()
```

另一组表示对读操作的锁定和解锁, 简称为: **读锁定** 和 **读解锁.**

```go
func (*RWMutex) RLock()
func (*RWMutex) RUnlock()
```

我们先来看一下没有使用读写锁的情况下会发生什么:

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readGo(in <-chan int, idx int){
	for {
		num := <- in
		fmt.Printf("----第%d个读go程, 读入: %d\n", idx, num)
	}

}

func writeGo(out chan<- int, idx int){
	for {
		// 生成随机数
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("第%d个写go程, 写入: %d\n", idx, num)
		time.Sleep(time.Millisecond * 300)
	}

}

func main() {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)

	for i:=0; i<5; i++ {
		go readGo(ch, i+1)
	}

	for i:=0; i<5; i++ {
		go writeGo(ch, i+1)
	}
	time.Sleep(time.Second * 3)
}
```

结果(截取部分):

```go
......
第4个写go程, 写入: 763
----第1个读go程, 读入: 998
第1个写go程, 写入: 238
第3个写go程, 写入: 998
......
第5个写go程, 写入: 607
第4个写go程, 写入: 151
----第1个读go程, 读入: 992
----第2个读go程, 读入: 151
......
```

通过结果我们可以知道, 当写入 `763` 时, 由于创建的是无缓冲的channel, 应该先把这个数读出来, 然后才可以继续写数据, 但是结果显示, 读到的是 `998`, `998` 在下面才显示写入啊, 怎么会先读出来呢? 出现这个情况的问题在于, 当运行到 `num := <- in` 时, 已经把 `998` 写进去了, 但是这个时候还没有来得及打印, 就失去了CPU, 失去CPU之后, 缓冲区中的数据就会被覆盖掉, 这时被 `763` 所覆盖.

这是第一个错误现象, 我们再来看一下第二个错误现象.

既然都是对数据进行读操作, 相邻的读入应该都是相同的数, 比如说`----第1个读go程, 读入: 992 ----第2个读go程, 读入: 151`, 这两个应该读到的数都是一样的, 但是结果显示却是不同的. 

那么加了读写锁之后, 先来看一下错误代码, 大家可以想一下为什么会出现这种错误.

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex sync.RWMutex

func readGo(in <-chan int, idx int){
	for {
		rwMutex.RLock()    // 以读模式加锁
		num := <- in
		fmt.Printf("----第%d个读go程, 读入: %d\n", idx, num)
		rwMutex.RUnlock()    // 以读模式解锁
	}
}

func writeGo(out chan<- int, idx int){

	for {
		// 生成随机数
		num := rand.Intn(1000)
		rwMutex.Lock()    // 以写模式加锁
		out <- num
		fmt.Printf("第%d个写go程, 写入: %d\n", idx, num)
		time.Sleep(time.Millisecond * 300)
		rwMutex.Unlock()    // 以写模式解锁
	}
}

func main() {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)

	for i:=0; i<5; i++ {
		go readGo(ch, i+1)
	}

	for i:=0; i<5; i++ {
		go writeGo(ch, i+1)
	}
	time.Sleep(time.Second * 3)
}
```

上面代码的结果会一直阻塞, 没有输出, 大家可以简单想一下出现这种情况的原因是什么?

代码看得仔细的应该都可以看出来, 这上面的代码中, 比如说读操作先抢到了CPU, 运行代码 `rwMutex.RLock()` 读加锁, 然后运行到 `num := <- in` 时, 会要求写端同时在线, 否则就会发生阻塞, 但是这时写端不可能在线, 因为读加锁了. 所以就会一直在这发生阻塞.

这也就是我们之前在死锁部分中提到的 **隐性死锁** (不报错).

那么解决办法有两种: 一种是不混用, 另一种是使用条件变量(之后会讲到)

我们先看一下不混用读写锁与channel的解决办法(**只使用读写锁, 如果只使用channel达不到想要的效果**):

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex2 sync.RWMutex    // 锁只有一把, 两个属性: r w

var value int    // 定义全局变量, 模拟共享数据

func readGo2(in <-chan int, idx int){
	for {
		rwMutex2.RLock()    // 以读模式加锁
		num := value
		fmt.Printf("----第%d个读go程, 读入: %d\n", idx, num)
		rwMutex2.RUnlock()    // 以读模式解锁
	}
}

func writeGo2(out chan<- int, idx int){
	for {
		// 生成随机数
		num := rand.Intn(1000)
		rwMutex2.Lock()    // 以写模式加锁
		value = num
		fmt.Printf("第%d个写go程, 写入: %d\n", idx, num)
		time.Sleep(time.Millisecond * 300)
		rwMutex2.Unlock()    // 以写模式解锁
	}
}

func main() {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)

	for i:=0; i<5; i++ {
		go readGo2(ch, i+1)
	}

	for i:=0; i<5; i++ {
		go writeGo2(ch, i+1)
	}
	time.Sleep(time.Second * 3)
}
```

结果: 

```go
......
第5个写go程, 写入: 363
----第4个读go程, 读入: 363
----第4个读go程, 读入: 363
----第4个读go程, 读入: 363
----第4个读go程, 读入: 363
----第2个读go程, 读入: 363
第5个写go程, 写入: 726
----第5个读go程, 读入: 726
----第4个读go程, 读入: 726
----第2个读go程, 读入: 726
----第1个读go程, 读入: 726
----第3个读go程, 读入: 726
第1个写go程, 写入: 764
----第5个读go程, 读入: 764
----第2个读go程, 读入: 764
----第5个读go程, 读入: 764
----第1个读go程, 读入: 764
----第3个读go程, 读入: 764
......
```

处于读锁定状态, 那么针对它的写锁定操作将永远不会成功, 且相应的goroutine也会被一直阻塞, 因为它们是互斥的.

**总结:** 读写锁控制下的多个写操作之间都是互斥的, 并且写操作与读操作之间也都是互斥的. 但是多个读操作之间不存在互斥关系.

从互斥锁和读写锁的源码可以看出, 它们是同源的. 读写锁的内部用互斥锁来实现写锁定操作之间的互斥. 可以把读写锁看作是互斥锁的一种扩展.

## 条件变量

在讲条件变量之前, 我们先来回顾一下之前的生产者消费者模型:

```go
package main

import (
	"fmt"
	"time"
)

func producer(out chan <- int) {
	for i:=0; i<5; i++ {
		fmt.Println("生产者, 生产: ", i)
		out <- i
	}
	close(out)
}

func consumer(in <- chan int) {
	for num := range in {
		fmt.Println("---消费者, 消费: ", num)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)
	time.Sleep(5 * time.Second)
}
```

之前都是一个生产者与一个消费者, 那么如果是多个生产者与多个消费者的情况呢?

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(out chan <- int, idx int) {
	for i:=0; i<10; i++ {
		num := rand.Intn(800)
		fmt.Printf("第%d个生产者, 生产: %d\n", idx, num)
		out <- num
	}
}

func consumer(in <- chan int, idx int) {
	for num := range in {
		fmt.Printf("---第%d个消费者, 消费: %d\n", idx, num)
	}
}

func main() {
	ch := make(chan int)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		go producer(ch, i + 1)
	}
	for i := 0; i < 5; i++ {
		go consumer(ch, i + 1)
	}
	time.Sleep(5 * time.Second)
}
```

如果是按照上面的代码写的话, 就又会出现之前的错误.

上面已经说过了, 解决这种错误有两种方法: 用锁或者用条件变量. 

这次就用条件变量来解决一下.

首先, 强调一下. **条件变量本身不是锁!! 但是经常与锁结合使用!!**

还有另外一个问题, 如果消费者比生产者多, 仓库中就会出现没有数据的情况. 我们需要不断的通过循环来判断仓库队列中是否有数据, 这样会造成cpu的浪费. 反之, 如果生产者比较多, 仓库很容易满, 满了就不能继续添加数据, 也需要循环判断仓库满这一事件, 同样也会造成cpu的浪费.

我们希望当仓库满时, 生产者停止生产, 等待消费者消费; 同理, 如果仓库空了, 我们希望消费者停下来等待生产者生产. 为了达到这个目的, 这里就引入了**条件变量**. (需要注意, 如果仓库队列用channel, 是不存在以上情况的, 因为channel被填满后就阻塞了, 或者channel中没有数据也会阻塞).

**条件变量:** 条件变量的作用并不保证在同一时刻仅有一个协程(线程)访问某个共享的数据资源, 而是在对应的共享数据的状态发生变化时, 通知阻塞在某个条件上的协程(线程). 条件变量不是锁, 在并发中不能达到同步的目的, 因此**条件变量总是与锁一块使用.**

例如, 我们上面说的, 如果仓库队列满了, 我们可以使用条件变量让生产者对应的goroutine暂停(阻塞), 但是当消费者消费了某个产品后, 仓库就不再满了, 应该唤醒(发送通知给)阻塞的生产者goroutine继续生产产品.

Go标准库中的 `sync.Cond` 类型代表了条件变量. 条件变量要与锁(互斥锁或者读写锁)一起使用. 成员变量L代表与条件变量搭配使用的锁.

```go
type Cond struct {
    noCopy noCopy
    L Locker
    notify notifyList
    checker copyChecker
}
```

对应的有3个常用的方法, `Wait`, `Signal`, `Broadcast`

1) func (c *Cond) Wait()

该函数的作用可归纳为如下三点:

- 阻塞等待条件变量满足
- 释放已掌握的互斥锁相当于cond.L.Unlock()。注意: **两步为一个原子操作(第一步与第二步操作不可再分).**
- 当被唤醒时, `Wait()` 函数返回时, 解除阻塞并**重新获取互斥锁**. 相当于cond.L.Lock()

2) func (c *Cond) Signal()

单发通知, 给一个正等待(阻塞)在该条件变量上的goroutine(线程)发送通知.

3) func (c *Cond) Broadcast()

广播通知, 给正在等待(阻塞)在该条件变量上的所有goroutine(线程)发送通知


下面, 我们就用条件变量来写一个**生产者消费者模型.**

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond  // 定义全局变量

func producer2(out chan<- int, idx int) {
	for {
		// 先加锁
		cond.L.Lock()
		// 判断缓冲区是否满
		for len(out) == 3 {
			cond.Wait()
		}
		num := rand.Intn(800)
		out <- num
		fmt.Printf("第%d个生产者, 生产: %d\n", idx, num)
		// 访问公共区结束, 并且打印结束, 解锁
		cond.L.Unlock()
		// 唤醒阻塞在条件变量上的 消费者
		cond.Signal()
	}
}

func consumer2(in <- chan int, idx int) {
	for {
		// 先加锁
		cond.L.Lock()
		// 判断缓冲区是否为 空
		for len(in) == 0 {
			cond.Wait()
		}
		num := <- in
		fmt.Printf("---第%d个消费者, 消费: %d\n", idx, num)
		// 访问公共区结束后, 解锁
		cond.L.Unlock()
		// 唤醒阻塞在条件变量上的生产者
		cond.Signal()
	}
}

func main() {
	// 设置随机种子数
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int, 3)

	cond.L = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go producer2(ch, i + 1)
	}
	for i := 0; i < 5; i++ {
		go consumer2(ch, i + 1)
	}
	time.Sleep(time.Second * 1)
}
```

1）定义 `ch` 作为队列, 生产者产生数据保存至队列中, 最多存储3个数据, 消费者从中取出数据模拟消费

2）条件变量要与**锁**一起使用, 这里定义全局条件变量 `cond`, 它有一个属性: `L Locker`, 是一个互斥锁.

3）开启5个消费者go程, 开启5个生产者go程.

4）`producer2` 生产者, 在该方法中开启互斥锁, 保证数据完整性. 并且判断队列是否满, 如果已满, 调用 `cond.Wait()` 让该goroutine阻塞. 当消费者取出数据后执行 `cond.Signal()`, 会唤醒该goroutine, 继续产生数据.

5）`consumer2` 消费者, 同样开启互斥锁, 保证数据完整性. 判断队列是否为空, 如果为空, 调用 `cond.Wait()` 使得当前goroutine阻塞. 当生产者产生数据并添加到队列, 执行 `cond.Signal()` 唤醒该goroutine.


**条件变量使用流程:**

1. 创建**条件变量**: var cond sync.Cond
2. 指定条件变量用的**锁**: cond.L = new(sync.Mutex)
3. 给公共区加锁(互斥锁): cond.L.Lock()
4. 判断是否到达阻塞条件(缓冲区**满/空**) --> for循环判断
    ```go
    for len(ch) == cap(ch) { cond.Wait() }
    或者 for len(ch) == 0 { cond.Wait() }
    1) 阻塞 2)解锁 3)加锁
    ```
5. 访问公共区 --> 读、写数据、打印
6. 解锁条件变量用的**锁**: cond.L.Unlock()
7. 唤醒阻塞在条件变量上的**对端**: cond.Signal()  cond.Broadcast()

## 李培冠博客

欢迎访问我的个人网站：

李培冠博客：[lpgit.com](https://lpgit.com)
