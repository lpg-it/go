相对于变量, 常量是恒定不变的值, 经常用于定义程序运行期间不会改变的那些值.

## 常量的定义使用

常量的声明与变量的声明很相似, 只是把 `var` 换成了 `const`, 常量在定义的时候**必须赋值**.

在程序开发中, 我们用常量存储一直不会发生变化的数据. 例如: `Π`， 身份证号码等. 像这类数据, 在整个程序运行中都是不允许发生改变的.

```go
package main

import "fmt"

func main(){
    const pi float64 = 3.14159
    // pi = 4.56  // 报错, 常量不允许修改
    fmt.Println(pi)

    // 自动推导类型
    const e = 2.7182  // 注意: 不是使用 :=
    fmt.Println("e =", e)
}
```

在声明了 `pi` 和 `e` 这两个变量之后, 在整个程序运行期间它们的值就都不能发生变化了.

**多个常量同时声明**

```go
const (
	pi = 3.14159
    e = 2.7182
)
```

`const` 同时声明多个常量时, 如果省略了值则表示**和上面一行的值是相同的.**

```go
const (
	n1 = 99
    n2  // n2 = 99
    n3  // n3 = 99
)
```

上面的示例中, 常量 `n1`、`n2`、`n3` 都是`99`.

## 字面常量

所谓字面常量, 是指程序中硬编码的常量.

```go
123  // 整数类型的常量
3.14159  // 浮点类型的常量
3.2+12i  // 复数类型的常量
true  // 布尔类型的常量
"foo"  // 字符串类型的常量
```

## iota 枚举

`iota` 是go语言的常量计数器, **只能在常量的表达式中使用**. 它用于生成一组以相似规则初始化的常量, 但是不用每一行都写一遍初始化表达式.

**注意**：在一个`const`声明语句中, 在第一个声明的常量所在的行, `iota` 将会被置为0, 然后在每一个有常量声明的行**加一**.

`iota` 可以理解为`const`语句块中的行索引, 使用`iota`能简化定义, 在定义枚举时很有用.

**看几个例子:**

可以只写一个`iota`

```go
package main

import "fmt"

func main(){
	const (
		a = iota  // 0
		b  // 1
		c  // 2
		d  // 3
	)
	fmt.Println(a, b, c, d)
}
```

`iota` 遇到 `const`, 会重置为 0

```go
package main

import "fmt"

func main(){
	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
    // iota遇到const, 会重置为0
	const e = iota  // 0
	fmt.Println(e)
}
```

使用 `_` 跳过某些值

```go
package main

import "fmt"

func main(){
	const (
		a = iota  // 0
		_
		c  // 2
		d  // 3
	)
	fmt.Println(a, c, d)
}
```

`iota` 声明中间插队

```go
package main

import "fmt"

func main(){
	const (
		a = iota  // 0
		b = 100  // 100
		c = iota  // 2
		d  // 3
	)
	fmt.Println(a, b, c, d)
}
```

常量写在同一行, 其值相同, 换一行值`+1`

```go
package main

import "fmt"

func main() {
	// 常量写在同一行, 其值相同, 换一行值+1
	const(
		a = iota  // 0
		b, c = iota, iota  // 1, 1
		d, e  // 2, 2
		f, g, h = iota, iota, iota  // 3, 3, 3
		i, j, k  // 4, 4, 4
	)
	fmt.Println(a)
	fmt.Println(b, c)
	fmt.Println(d, e)
	fmt.Println(f, g, h)
	fmt.Println(i, j, k)
}
```

可以为其赋初始值, 但是换行后不会根据值`+1`, 而是根据 **行** `+1`.

```go
package main

import "fmt"

func main(){
    const (
    	a = 6  // 6
        b, c = iota, iota  // 1 1
        d, e  // 2 2
        f, g, h = iota, iota, iota  // 3 3 3
        i, j, k  // 4 4 4
    )
	fmt.Println(a)
	fmt.Println(b, c)
	fmt.Println(d, e)
	fmt.Println(f, g, h)
	fmt.Println(i, j, k)
}
```

如果一行中赋值的初始值不一样, 则下一行的值与上一行相等.

```go
package main

import "fmt"

func main(){
    const (
    	a, b = 1, 6  // 1 6
        c, d  // 1 6
        e, f, g = 2, 8, 10  // 2 8 10
        h, i, j  // 2 8 10
    )
    fmt.Println(a, b)
    fmt.Println(c, d)
    fmt.Println(e, f, g)
    fmt.Println(h, i, j)
}
```

如果一行中既有赋初始值, 又有`iota`, 则下一行中对应初始值的位置的值不变, 对应 `iota` 位置的值`+1`.

```go
package main

import "fmt"

func main(){
    const (
    	a, b, c = 3, iota, iota  // 3 0 0
        d, e, f  // 3 1 1
        g, h, i = iota, 16, iota  // 2 16 2
        j, k, l  // 3 16 3
    )
    fmt.Println(a, b, c)
    fmt.Println(d, e, f)
    fmt.Println(g, h, i)
    fmt.Println(j, k, l)
}
```

当对 `iota` 进行加减操作时, 下一行也进行同样操作

```go
package main

import "fmt"

func main(){
    const (
		a, b = iota+5, iota-2  // 5 -2
		c, d  // 6 -1
	)
	fmt.Println(a, b)
	fmt.Println(c, d)
}
```

定义数量级

```go
package main

import "fmt"

func main(){
    const (
    	_ = iota
        KB = 1 << (10 * iota)  // 1024
        MB = 1 << (10 * iota)
        GB = 1 << (10 * iota)
        TB = 1 << (10 * iota)
        PB = 1 << (10 * iota)
    )
    fmt.Println(KB, MB, GB, TB, PB)
}

```

这里的 `<<` 表示左移操作, `1<<10` 表示将`1`的`二进制表示`向左移`10`位, 也就是由`1`变成了`10000000000`, 也就是`十进制的1024`.

同理, `2<<3` 表示将`2`的`二进制表示`向左移`3`位, 也就是由`10`变成了`10000`, 也就是`十进制的16`

## 李培冠博客

欢迎访问我的个人网站：

李培冠博客：[lpgit.com](https://lpgit.com)
