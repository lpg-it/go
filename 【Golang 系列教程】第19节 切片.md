## 前言

数组的长度在定义之后无法再次修改, 数组是值类型, 每次传递都将产生一份副本.

显然这种数据结构无法完全满足我们的真实需求. 所以, Go语言提供了**切片(Slice)**来弥补数组的不足.

Slice代表变长的序列, 序列中的每个元素都有**相同的类型**, 一个Slice类型一般写作`[]T`, 其中 `T` 代表Slice中元素的**类型**; Slice的语法和数组很像, 只是没有固定长度而已.

数组和Slice之间有着紧密的联系. 一个Slice是一个轻量级的数据结构, 提供了访问数组子序列(或者全部)元素的功能, 而且**Slice的底层确实引用一个数组对象**.

一个Slice由三个部分构成: **指针**, **长度**和**容量**.

**指针指向第一个Slice元素对应的底层数组元素的地址**, 要注意的是Slice的第一个元素并不一定就是数组的第一个元素.

切片并不是数组或数组指针, **它通过内部指针和相关属性引用数组片段, 以实现变长方案.**

多个Slice之间可以**共享底层的数据**, 并且引用的数组部分区间可能重叠.

Slice并不是真正意义上的动态数组, 而是一个引用类型, Slice总是指向一个底层数组.

所以, **为什么要有切片?**

- **数组的容量固定**, 不能自动扩展.
- **数组为值传递**. 数组作为函数参数时, 会将整个数组值拷贝一份给形参.

在Go语言中, 我们几乎可以在所有的场景中, 使用切片来替换数组使用.

**切片的本质**

- 不是一个数组的指针, 是一种数据结构, 用来操作数组内部元素.

## 创建切片

Slice和数组的区别: 声明**数组**时, `[]` 内写明了数组的**长度**, 而声明**Slice**时, `[]` 内**没有任何字符**或使用`...`自动计算长度.

可以使用创建数组的方式对切片进行初始化:

```go
arr := [6]int{1, 2, 3, 4, 5, 6}
s := arr[1: 3: 5]
```

```go
切片名称[low: high: max]
low: 起始下标位置
high: 结束下标位置    len = high - low
max: 容量    cap = max - low
```

我们可以看一下这个切片的长度以及容量:

```go
package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	s := arr[1: 3: 5]
	fmt.Println(s)
	fmt.Println("len(s) = ", len(s))
	fmt.Println("cap(s) = ", cap(s))
}
```

结果:

```go
[2 3]
len(s) =  2
cap(s) =  4
```

当然, 在**截取原数组**时, 我们也可以忽略容量, 此时**容量 = 原数组长度 - low**.

```go
package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}

	s := arr[1:3]
	fmt.Println(s)
	fmt.Println("len(s) = ", len(s))
	fmt.Println("cap(s) = ", cap(s))
}
```

结果: 

```go
[2 3]
len(s) =  2
cap(s) =  5
```

再来看下面这段代码, `s2` 的长度与容量是多少呢? 

```go
package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := arr[1: 5: 8]
	fmt.Println(s)
	fmt.Println("len(s) = ", len(s))
	fmt.Println("cap(s) = ", cap(s))

	s2 := s[: 6]
	fmt.Println(s)
	fmt.Println("len(s2) = ", len(s2))
	fmt.Println("cap(s2) = ", cap(s2))
}
```

结果: 

```go
[2 3 4 5]
len(s) =  4
cap(s) =  7
[2 3 4 5 6 7]
len(s2) =  6
cap(s2) =  7
```

这时, `s2` 没有使用**容量**, 所以它的**容量**应该等于**s的容量 - low**.

我们继续来研究一下下面的代码, 猜猜`s3`的结果是多少?

```go
package main

import "fmt"

func main(){
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := arr[2: 5]  // [3, 4, 5]
	fmt.Println("s = ", s)

	s2 := arr[2: 7]  // [3, 4, 5, 6, 7]
	fmt.Println("s2 = ", s2)

	s3 := s[2: 7]
	fmt.Println("s3 = ", s3)
}
```

结果:

```go
s =  [3 4 5]
s2 =  [3 4 5 6 7]
s3 =  [5 6 7 8 9]
```

有数组时我们可以截取数组来创建切片, 下面我们来看一下没有数组时创建切片的方法.

**常用的切片创建方法:**

1. 自动推导类型创建Slice

```go
s := []int{1, 2, 3, 4}  // 创建有 4 个元素的切片
```

2. 借助**make**创建Slice, 格式: `make(切片类型, 长度, 容量)`

```go
s := make([]int, 5, 10)  // len(s) = 5, cap(s) = 10
```

3. **make**时, 没有指定容量, 那么**容量 = 长度**.

```go
s := make([]int, 5)  // len(s) = 5, cap(s) = 5
```

```go
package main

import "fmt"

func main(){
	// 自动推导赋初始值
	s1 := []int{1, 2, 4, 8}
	fmt.Println("s1 = ", s1, "len(s1) = ", len(s1), "cap(s1) = ", cap(s1))

	// make创建切片, 指定容量
	s2 := make([]int, 5, 10)
	fmt.Println("s2 = ", s2, "len(s2) = ", len(s2), "cap(s2) = ", cap(s2))

	// make创建切片, 不指定容量
	s3 := make([]int, 5)
	fmt.Println("s3 = ", s3, "len(s3) = ", len(s3), "cap(s3) = ", cap(s3))

}
```

**注意: make只能创建`Slice`, `map`, `channel`, 并且返回一个有初始值的对象.**

## 切片做函数参数

切片在做函数参数时, **传引用(地址).**

```go
package main

import "fmt"

func foo(s []int){  // 切片做函数参数
	s[0] = -1    // 直接修改 main 中的 slice
}

func main(){
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)

	foo(slice)  // 传引用

	fmt.Println(slice)
}
```

结果:

```go
[1 2 3 4]
[-1 2 3 4]
```

## 常用操作函数

### append函数

`append()` 函数可以向Slice尾部添加数据, 可以自动为切片扩容, 常常会返回**新**的Slice对象.




**append**函数会智能的将底层数组的容量增长, 一旦超过原底层数组容量, 通常以2倍(1024以下)容量重新分配底层数组, 并复制原来的数据. 因此, 使用**append**给切片做扩充时, **切片的地址可能发生变化**. 但是数据都被重新保存了, 不影响使用.













## 李培冠博客

欢迎访问我的个人网站：

李培冠博客：[lpgit.com](https://lpgit.com)












