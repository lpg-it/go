**输出**就是将数据信息打印到电脑屏幕上. 本节我们就来学习一下Go语言中的三种输出方式: fmt.Print()、fmt.Println()、fmt.Printf().

## fmt.Print()

fmt.Print() 主要的一个特点就是打印数据时 **不换行**.

```go
package main

import "fmt"

func main() {
    a, b := 10, 20
    
    // 输出: Print, 打印数据时不带换行
    fmt.Print(a)
    fmt.Print(b)
}
// 结果:
1020
```

## 2. fmt.Println()

fmt.Println() 之前已经用到过, 为**换行输出**.

举个例子：

```go
package main

import "fmt"

func main() {
    a, b := 10, 20
    // 输出: Println, 打印数据时自带换行
    fmt.Println(a)
    fmt.Println(b)
}
// 结果:
10
20
```

这个时候, 你是知道这个结果`10` `20` 都代表什么意思. 但是如果换一位程序员来看, 就不知道了, 尤其是在代码量特别大的情况下. 所以, 应该采用以下输出:

```go
package main

import "fmt"

func main() {
    a, b := 10, 20
    // 双引号内的内容会原样输出. 注意与变量名之间用逗号分隔
    fmt.Println("a =", a)
    fmt.Println("b =", b)
}
// 结果:
a = 10
b = 20
```

## 3. fmt.Printf()

除了以上两种输出函数以外, Go语言中还有一个函数 `fmt.Printf()` : 格式化输出.

格式化输出也可以实现换行输出;

```go
package main

import "fmt"

func main() {
    a, b := 10, 20
    
	// %d 占位符, 表示输出一个整型数据
    // 第一个 %d 会被变量 a 的值替换, 第二个 %d 会被变量 b 替换
    // \n 表示换行
    fmt.Printf("a = %d\nb = %d", a, b)
}
// 结果:
a = 10
b = 20
```

fmt.Printf() 适合有结构的输出多个变量的值:

```go
package main

import "fmt"

func main() {
    a, b, c := 10, 20, 30
    fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}
// 结果:
a = 10, b = 20, c = 30

```

关于占位符的内容, 会在后面的文章中讲到.

## 李培冠博客

[lpgit.com](https://lpgit.com)

