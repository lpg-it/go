## 前言

哈希表是一种巧妙并且实用的数据结构。它是一个**无序的** key/value对 的集合，其中所有的 key 都是不同的，然后通过给定的 key 可以在常数时间复杂度内检索、更新或删除对应的 value。

在 Go 语言中，一个 map 就是一个哈希表的引用，map 类型可以写为 map[K]V，其中 K 和 V 分别对应 key 和 value。map 中所有的 key 都有相同的类型，所有的 value 也有着相同的类型，但是 key 和 value 之间可以是不同的数据类型。其中 K 对应的 key 必须是支持 == 比较运算符的数据类型（切片、函数等不支持），**所以 map 可以通过测试 key 是否相等来判断是否已经存在**。虽然浮点数类型也是支持相等运算符比较的，但是将浮点数用做 key 类型则是一个坏的想法。对于 V 对应的 value 数据类型则没有任何的限制。

- map 是无序的
- 在 Go 语言中的 map 是引用类型，**必须初始化才能使用**。

Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。由于 map 是无序的，我们无法决定它的返回顺序。

## map 的定义

可以使用内建函数 make 也可以使用 map 关键字来定义 map:

```go
// 使用 make 函数
m := make(map[keyType]valueType)
// 长度为 0 的 map
m := make(map[keyType]valueType, 0)

// 声明变量，默认 map 是 nil
var m map[keyType]valueType
// 长度为 0 的 map
var m map[keyType]valueType{}
```

其中：

- m 为 map 的变量名。
- keyType 为键类型。
- valueType 是键对应的值类型。

在声明的时候不需要知道 map 的长度，因为 map 是可以动态增长的。但是如果我们提前知道 map 需要的长度，最好指定一下。

我们可以用 `len(m)` 来查看 map 的长度。**注意，使用 `cap(m)` 会报错（cap 支持 数组、指向数组的指针、切片、channel）：**

```go
invalid argument m (type map[string]int) for cap
```

**如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对。如果向一个 nil 值的 map 存入元素将导致一个 panic 异常：**

下面我们用 make 函数创建一个 map：

```go
ages := make(map[string]int)
```

当然，我们也可以直接创建一个 map 并且指定一些最初的值：

```go
ages := map[string]int{
    "Conan": 18,
    "Kidd": 23,
}
```

这种就相当于：

```go
ages := make(map[string]int)
ages["Conan"] = 18
ages["Kidd"] = 23
```

所以，另一种创建空（**不是 nil**）的 map 方法是：

```go
ages := map[string]int{}
```

map 在**定义**时，key 是唯一的，不允许重复（value 可以重复）。下面的程序会**报错**：

```go
ages := map[string]int{
    "Conan": 18,
    "Conan": 23,
}
```

但是之后在对 map 赋值时，则会覆盖原来的 value

```go
ages["Conan"] = 18
ages["Conan"] = 23
fmt.Println(ages["Conan"])  // 23
```

map 类型的零值是 nil，也就是没有引用任何哈希表，其长度也为 0.

```go
var ages map[string]int
fmt.Println(ages == nil)  // true
fmt.Println(len(ages))  // 0
```

## map 的基本使用

### 增

增加 map 的值很简单，只需要 `m[key] = value` 即可，比如：

```go
ages := make(map[string]int)
ages["Conan"] = 18
ages["Kidd"] = 23
```

### 删

使用内置的 delete 函数可以删除元素，参数为 map 和其对应的 key，没有返回值：

```go
delete(ages, "Conan")
```

注意：即使这些 key 不在 map 中也不会报错。

### 改

修改 map 的内容和 增 的写法类似，只不过 key 是已存在的，如果不存在，则为增加，例如：

```go
ages := map[string]int{
    "Conan": 18,
    "Kidd": 23,
}
ages["Conan"] = 21
```

### 查

map 中的元素通过 key 对应的下标语法访问：

```go
ages["Conan"] = 18
fmt.Println(ages["Conan"])  // 18
```

要想遍历 map 中全部的键值对的话，可以使用 range 风格的 for 循环实现，和之前的 slice 遍历语法类似。例如：

```go
for key, value := range ages {
    fmt.Println(key, value)
}
```

如果用不到 value，无需使用匿名变量  `_`，直接不写即可：

```go
for key := range ages {
    fmt.Println(key)
}
```

如果查找失败也没有关系，程序也不会报错，而是返回 value 类型对应的零值。例如：

```go
ages := map[string]int{
    "Conan": 18,
    "Kidd": 23,
}
fmt.Println(ages["Lan"])  // 0
```

通过 key 作为索引下标来访问 map 将产生一个 value。如果 key 在 map 中是存在的，那么将得到与 key 对应的 value；如果 key 不存在，那么将得到 value 对应类型的零值。

但是有时候我们需要知道对应的元素是否真的是在 map 之中。比如，如果元素类型是一个数字，你需要区分一个已经存在的 0，和不存在而返回零值的 0。例如：

```go
ages := map[string]int{
    "Conan": 18,
    "Kidd": 23,
}
// 如果 key 存在，则 ok = true；不存在，ok = false
if value, ok := ages["Conan"]; ok {
    fmt.Println(value)
} else {
    fmt.Println("key 不存在")
}
```

在这种场景下，map 的下标语法将产生两个值；第二个是一个布尔值，用于报告元素是否真的存在。布尔变量一般命名为 ok，特别适合马上用于 if 条件判断部分。

### 排

**map 的迭代顺序是不确定的**。有没有什么办法可以顺序的打印出 map 呢？我们可以借助切片来完成。先将 key（或者 value）添加到一个切片中，再对切片排序，然后使用 for-range 方法打印出所有的 key 和 value。如下所示：

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 创建一个 ages map，并给三个值
	ages := make(map[string]int)
	ages["Conan"] = 18
	ages["Kidd"] = 23
	ages["Lan"] = 19

	// 创建一个切片用于给 key 进行排序
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)

	// 循环打印出 map 中的值
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
```

因为我们一开始就知道 names 的最终大小，因此给切片分配一个合适的容量大小将会更有效。下面的代码创建了一个空的切片，但是切片的容量刚好可以放下 map 中全部的 key：

```go
names := make([]string, 0, len(ages))
```

当然，如果使用结构体切片，这样就会更有效：

```go
type name struct {
    key string
    value int
}
```

### 比

map 之间不能进行相等比较；**唯一的例外是和 nil 进行比较**。要判断两个 map 是否包含相同的 key 和 value，我们必须通过一个循环实现：

```go
func equalMap(x, y map[string]int) bool {
    // 长度不一样，肯定不相等
    if len(x) != len(y) {
        return false
    }
    for k, xv := range x {
        if yv, ok := y[k]; !ok || xv != yv {
            return false
        }
    }
    return true
}
```

## map 作为函数参数

map 作为函数参数是**地址传递**（引用传递），作返回值时也一样。

在函数内部对 map 进行操作，会影响主调函数中实参的值。例如：

```go
func foo(m map[string]int) {
    m["Conan"] = 22
    m["Lan"] = 21
}

func main() {
    m := make(map[string]int, 2)
    m["Conan"] = 18
	fmt.Println(m)  // map[Conan:18]
	foo(m)
	fmt.Println(m)  // map[Conan:22 Lan:21]
}
```

## 并发环境中使用的 map：sync.Map

Go 语言中的 map 在并发情况下，**只读**是线程安全的，**同时读写是线程不安全的**。

下面我们来看一下在并发情况下**读写 map 时会出现的问题**，代码如下：

```go
// 创建一个 map
m := make(map[int]int)

// 开启一个 go 程
go func () {
    // 不停地对 map 进行写入
    for true {
        m[1] = 1
    }
}()

// 开启一个 go 程
go func() {
    // 不停的对 map 进行读取
    for true {
        _ = m[1]
    }
}()

// 运行 10 秒停止
time.Sleep(time.Second * 10)
```

运行代码会报错，错误如下：

```go
fatal error: concurrent map read and map write
```

当两个并发函数不断地对 map 进行读和写时，map 内部会对这种并发操作进行检查并提前发现。

当我们需要并发读写时，一般的做法是加锁，但是这样性能不高。

Go 语言在 1.9 版本中提供了一种效率较高的并发安全的 sync.Map。

sync.Map 有以下特性：

- 无须初始化，直接声明即可
- sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用：Store 表示存储，Load 表示获取，Delete 表示删除。
- 使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时返回 true，终止迭代遍历时，返回 false。

并发安全的 sync.Map 示例代码如下：

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var ages sync.Map

	// 将键值对保存到 sync.Map
	ages.Store("Conan", 18)
	ages.Store("Kidd", 23)
	ages.Store("Lan", 18)

	// 从 sync.Map 中根据键取值
	age, ok := ages.Load("Conan")
	fmt.Println(age, ok)

	// 根据键删除对应的键值对
	ages.Delete("Kidd")
	fmt.Println("删除后的 sync.Map： ", ages)

	// 遍历所有 sync.Map 中的键值对
	ages.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
```

sync.Map 没有提供获取 map 数量的方法，替代方法是**在获取 sync.Map 时遍历自行计算数量**，sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。

所以，我们用 sync.Map 时进行同时读写是没问题的，示例代码如下：

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Map

	// 开启一个 go 程
	go func() {
		// 不停地对 map 进行写入
		for true {
			m.Store(1, 1)
		}
	}()

	// 开启一个 go 程
	go func() {
		// 不停的对 map 进行读取并打印读取结果
		for true {
			value, _ := m.Load(1)
			fmt.Println(value)
		}
	}()
	time.Sleep(time.Second * 10)
}
```

这时的结果就会一直输出 1。

## 练习

1、封装 wordCountFunc() 函数。接收一段英文字符串 str。返回一个 map，记录 str 中每个“单词”出现的次数。

**示例：**

```go
输入："I love my work and I love my family too"
输出：
    family：1
    too：1
    I：2
    love：2
    my：2
    work：1
    and：1
```

**提示：使用 strings.Fields() 函数可提高效率**

**实现：**

```go
package main

import (
	"fmt"
	"strings"
)

func wordCountFunc(str string) map[string]int {
	// 使用 strings.Fields 进行拆分, 自动按照空格对字符串进行拆分成切片
	wordSlice := strings.Fields(str)
	// 创建一个用于存储 word 次数的 map
	m := make(map[string]int)

	// 遍历拆分后的字符串切片
	for _, value := range wordSlice {
		if _, ok := m[value]; !ok {
			// key 不存在
			m[value] = 1
		} else {
			// key 值已存在
			m[value]++
		}
	}
	return m
}

func main() {
	str := "I love my work and I love my family too"
	res := wordCountFunc(str)

	// 遍历 map, 展示每个 word 出现的次数
	for key, value := range res {
		fmt.Println(key, ": ", value)
	}
}
```

如需更深入的了解 map 的原理，推荐阅读这篇文章：[深度解密Go语言之map](https://www.cnblogs.com/qcrao-2018/p/10903807.html)

## 李培冠博客

欢迎访问我的个人网站：

李培冠博客：[lpgit.com](https://lpgit.com)
