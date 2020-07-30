前面我们所写的程序, 都是直接给变量赋值. 但是在很多情况下, 我们希望用户通过键盘输入一个数值, 然后存储到某个变量中, 接着将该变量的值取出来, 进行操作.

那么Go语言中怎么接收用户的键盘输入呢? 具体操作如下: 

## 第一种形式：fmt.Scanf()

```go
package main

import "fmt"

func main() {
	var age int
	fmt.Print("请输入你的年龄: ")
	fmt.Scanf("%d", &age)
	fmt.Printf("你的年龄为: %d", age)
}
```

在Go语言中, 我们用到了 `fmt` 这个包中的 `Scanf()` 函数来接收用户键盘输入的数据. 

当程序执行到 `Scanf()` 函数后, 会停止往下执行, 等待用户的输入 , 输入完成后程序继续往下执行.

在这里要重点注意的是 `Scanf()` 函数的书写格式:

1. 要用 "%d" 来表示输入的是一个整数, 输入完整数后存储到变量 `age` 中.
2. 这里的 `age` 变量前面一定要加上 `&` 符号, 表示获取内存单元的地址, 然后才能存储.

### 第二种形式：fmt.Scan()

还有另外一种获取用户输入数据的方式, 具体如下: 

```go
package main

import "fmt"

func main() {
	var age int
	fmt.Print("请输入你的年龄: ")
	fmt.Scan(&age)
	fmt.Printf("你的年龄为: %d", age)
}
```

通过 `Scan()` 函数接收用户输入, 这时可以省略掉 `%d`, 写法更简单.

## 李培冠博客

[lpgit.com](https://lpgit.com)

