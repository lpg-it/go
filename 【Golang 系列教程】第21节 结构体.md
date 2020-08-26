## 前言

结构体是一种聚合的**数据类型**，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。

用结构体的经典案例：学校的学生信息，每个学生信息包含一个唯一的学生学号、学生的名字、学生的性别、家庭住址等等。所有的这些信息都需要绑定到一个实体中，可以作为一个整体单元被复制，作为函数的参数或返回值，或者是被存储到数组中，等等。

结构体也是值类型，因此可以通过 new 函数来创建。

组成结构体类型的那些数据称为**字段**（fields）。字段有以下特性：

- 字段拥有自己的类型和值。
- 字段名必须唯一。
- 字段的类型也可以是结构体，甚至是字段所在结构体的类型。

关于 Go 语言的类（class）

Go 语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。Go 语言的结构体与“类”都是复合结构体，但 Go 语言中结构体的内嵌配合接口比面向对象具有更高的扩展性和灵活性。

Go 语言不仅认为结构体能拥有方法，且每种自定义类型也可以拥有自己的方法。

## 结构体的定义

使用关键字 type 可以将各种基本类型定义为自定义类型，基本类型包括整型、字符串、布尔等。**结构体是一种复合的基本类型**，通过 type 定义为自定义类型后，使结构体更便于使用。

结构体的定义格式如下：

```go
type 结构体类型名 struct {
    字段1 字段1类型
    字段2 字段2类型
    …
}
```

其中：

1、结构体类型名：标识自定义结构体的名称，在同一个包内不能重复。

2、字段1：表示结构体字段名。结构体中的字段名必须唯一。

3、字段1类型：表示结构体字段的具体类型。

举个例子，我们定义一个 Student（学生）结构体，代码如下：

```go
type Student struct{
    id      int
    name    string
    age     int
    gender  int // 0 表示女生，1 表示男生
    addr    string
}
```

**在这里，Student 的地位等价于 int、byte、bool、string 等类型**

通常一行对应一个结构体成员，成员的名字在前，类型在后

不过如果相邻的成员类型如果相同的话可以被合并到一行：

```go
type Student struct{
    id          int
    name        string
    age, gender int
    addr        string
}
```

这样我们就拥有了一个 Student 的自定义类型，它有 id、name、age等字段。

这样我们使用这个 Student 结构体就能够很方便的在程序中表示和存储学生信息了。

## 递归结构体

结构体类型可以通过引用自身来定义。这在定义链表或二叉树的元素（通常叫节点）时特别有用，此时节点包含指向临近节点的链接（地址）。

如下所示，链表中的 `su`，树中的 `ri` 和 `le` 分别是指向别的节点的指针。

### 链表

```go
type Node struct {
    Data float64
    Next *Node
}
```

### 双向链表

```go
type Node struct {
    Per *Node
    Data float64
    Next *Node
}
```

### 二叉树

```go
type Tree struct {
    le *Tree
    data float64
    ri *Tree
}
```

## 结构体的实例化

结构体的定义只是一种内存布局的描述，只有当结构体实例化时，才会真正地分配内存，因此必须在定义结构体并实例化后才能使用结构体的字段。

实例化就是根据结构体定义的格式创建一份与格式一致的内存区域，结构体实例与实例间的内存是完全独立的。

Go语言可以通过多种方式实例化结构体，根据实际需要可以选用不同的写法。

### 基本的实例化形式

**结构体本身也是一种类型**，我们可以像声明内置类型一样使用 var 关键字声明结构体类型。

基本实例化格式如下：

```go
var 结构体实例 结构体类型
```

对 Student 进行实例化，代码如下：

```go
type Student struct{
    id      int
    name    string
    age     int
    gender  int // 0 表示女生，1 表示男生
    addr    string
}

func main() {
    var stu1 Student
    stu1.id = 120100
    stu1.name = "Conan"
    stu1.age = 18
    stu1.gender = 1
    fmt.Println("stu1 = ", stu1)  // stu1 =  {120100 Conan 18 1 }
}
```

注意：没有赋值的字段默认为该字段类型的零值，此时 `addr = ""`

我们可以通过 `点 "."` 的方式来访问结构体的成员变量，如 `stu1.name`，结构体成员变量的赋值方法与普通变量一致。

### 创建指针类型的结构体

Go 语言中，还可以使用 new 关键字对类型（包括结构体、整型、浮点数、字符串等）进行实例化，结构体在实例化后会形成**指针类型**的结构体。

使用 new 的格式如下：

```go
变量名 := new(类型)
```

Go 语言让我们可以像访问普通结构体一样使用 `点"."` 来访问结构体指针的成员，例如：

```go
type Student struct{
    id      int
    name    string
    age     int
    gender  int // 0 表示女生，1 表示男生
    addr    string
}

func main() {
    stu2 := new(Student)
    stu2.id = 120101
    stu2.name = "Kidd"
    stu2.age = 23
    stu2.gender = 1
    fmt.Println("stu2 = ", stu2)  // stu2 =  &{120101 Kidd 23 1 }
}
```

经过 new 实例化的结构体实例在成员赋值上与基本实例化的写法一致。

注意：在 Go 语言中，访问**结构体指针**的成员变量时可以继续使用 `点"."`，这是因为 Go 语言为了方便开发者访问结构体指针的成员变量，使用了语法糖（Syntactic sugar）技术，将 stu2.name 形式转换为 (*stu2).name。

### 取结构体的地址实例化

在 Go 语言中，对结构体进行 `&` 取地址操作时，视为对该类型进行一次 new 的实例化操作，取地址格式如下：

```go
变量名 := &结构体类型{}
```

取地址实例化是最广泛的一种结构体实例化方式，具体代码如下：

```go
type Student struct{
    id      int
    name    string
    age     int
    gender  int // 0 表示女生，1 表示男生
    addr    string
}

func main() {
	stu3 := &Student{}
	stu3.id = 120102
	stu3.name = "Lan"
	stu3.age = 18
	stu3.gender = 0
	fmt.Println("stu3 = ", stu3)  // stu3 =  &{120102 Lan 18 0 }
}
```

## 结构体的初始化

结构体在实例化时可以直接对成员变量进行初始化，初始化有两种形式分别是以字段“键值对”形式和多个值的列表形式。

键值对形式的初始化适合选择性填充字段较多的结构体，多个值的列表形式适合填充字段较少的结构体。

特别地，还有一种初始化匿名结构体。

### 使用“键值对”初始化结构体

结构体可以使用“键值对”（Key value pair）初始化字段，每个“键”（Key）对应结构体中的一个字段，键的“值”（Value）对应字段需要初始化的值。

键值对的填充是可选的，不需要初始化的字段可以不填入初始化列表中。

**结构体实例化后字段的默认值是字段类型的零值**，例如 ，数值为 0、字符串为 ""（空字符串）、布尔为 false、指针为 nil 等。

键值对初始化的格式如下：

```go
变量名 := 结构体类型名{
    字段1: 字段1的值,
    字段2: 字段2的值,
    ...
}
```

注意：

1、字段名只能出现一次。

2、键值之间以 : 分隔，键值对之间以 , 分隔。

使用键值对形式初始化结构体的代码如下：

```go
stu4 := Student{
	id:     120103,
	name:   "Gin",
	age:    25,
	gender: 1,
	addr:   "unknown",
}
fmt.Println("stu4 = ", stu4) // stu4 =  {120103 Gin 25 1 unknown}
```

### 使用多个值的列表初始化结构体

Go语言可以在“键值对”初始化的基础上忽略“键”，也就是说，可以使用多个值的列表初始化结构体的字段。

多个值使用逗号分隔初始化结构体，例如：

```go
变量名 := 结构体类型名{
    字段1的值,
    字段2的值,
    ...
}
```

注意：

1、必须初始化结构体的**所有字段**。

2、每一个初始值的填充顺序必须与字段在结构体中的声明顺序一致。

3、键值对与值列表的初始化形式不能混用。

使用多个值列表初始化结构体的代码如下：

```go
stu5 := Student{
	120104,
	"Kogorou",
	38,
	1,
	"毛利侦探事务所",
}
fmt.Println("stu5 = ", stu5) // stu5 =  {120104 Kogorou 38 1 毛利侦探事务所}
```

### 初始化匿名结构体

匿名结构体没有类型名称，无须通过 type 关键字定义就可以直接使用。

例如：

```go
package main

import (
    "fmt"
)

func main() {
    var user struct{name string; age int}
    user.name = "Conan"
    user.age = 18
    fmt.Println("user = ", user)  // user =  {Conan 18}
}
```

## 结构体的赋值与比较

### 结构体的赋值

当使用 = 对结构体赋值时，更改其中一个结构体的值不会影响另外的值：

```go
package main

import (
	"fmt"
)

type Student struct {
	id     int
	name   string
	age    int
	gender int // 0 表示女生，1 表示男生
	addr   string
}

func main() {
	var stu1 Student
	stu1.id = 120100
	stu1.name = "Conan"
	stu1.age = 18
	stu1.gender = 1

	stu6 := stu1
	fmt.Println("stu1 = ", stu1) // stu1 =  {120100 Conan 18 1 }
	fmt.Println("stu6 = ", stu6) // stu6 =  {120100 Conan 18 1 }

	stu6.name = "柯南"
	fmt.Println("stu1 = ", stu1) // stu1 =  {120100 Conan 18 1 }
	fmt.Println("stu6 = ", stu6) // stu6 =  {120100 柯南 18 1 }
}
```

### 结构体的比较

如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的；如果结构体中存在不可比较的成员变量，比如说切片、map等，那么结构体就不能比较。这时如果强行用 ==、!= 来进行判断的话，程序会直接报错，我们可以用 DeepEqual 来进行深度比较。

如果结构体的全部成员都是可以比较的，那么两个结构体将可以使用 == 或 != 运算符进行比较。

相等比较运算符 == 将比较两个结构体的每个成员，因此下面两个比较的表达式是等价的：

```go
type Student struct {
	id   int
	name string
}

func main() {
	var stu1 Student
	stu1.id = 120100
	stu1.name = "Conan"

	stu6 := stu1
	stu6.name = "柯南"

	fmt.Println(stu1.id == stu6.id && stu1.name == stu6.name) // "false"
	fmt.Println(stu1 == stu6)                                 // "false"
}
```

可比较的结构体类型和其他可比较的类型一样，可以用于 map 的 key 类型。

## 结构体数组和切片

现在我们有一个需求：用结构体存储多个学生的信息。

我们就可以定义结构体数组来存储，然后通过循环的方式，将结构体数组中的每一项进行输出：

```go
package main

import "fmt"

type student struct {
	id   int
	name string
	score  int
}

func main() {
	// 结构体数组
	students := [3]student{
		{101, "conan", 88},
		{102, "kidd", 78},
		{103, "lan", 98},
	}
	// 打印结构体数组的每一项
	for index, stu := range students {
		fmt.Println(index, stu.name)
	}
}
```

用结构体切片存储同理。

**练习1：计算以上学生成绩的总分。**

```go
package main

import "fmt"

type student struct {
	id    int
	name  string
	score int
}

func main() {
	// 结构体数组
	students := [3]student{
		{101, "conan", 88},
		{102, "kidd", 78},
		{103, "lan", 98},
	}
    // 计算以上学生成绩的总分
	sum := students[0].score
	for i, stuLen := 1, len(students); i < stuLen; i++ {
		sum += students[i].score
	}
	fmt.Println("总分是：", sum)
}
```

**练习2：输出以上学生成绩中最高分。**

```go
package main

import "fmt"

type student struct {
	id    int
	name  string
	score int
}

func main() {
	// 结构体数组
	students := [3]student{
		{101, "conan", 88},
		{102, "kidd", 78},
		{103, "lan", 98},
	}

	// 输出以上学生成绩中最高分
	maxScore := students[0].score
	for i, stuLen := 1, len(students); i < stuLen; i++ {
		if maxScore < students[i].score {
			maxScore = students[i].score
		}
	}
	fmt.Println("最高分是：", maxScore)
}
```

## 结构体作为 map 的 value

结构体作为 map 的 value 示例如下：

```go
package main

import "fmt"

type student struct {
	id    int
	name  string
	score int
}

func main0801() {
	// 结构体数组
	students := [3]student{
		{101, "conan", 88},
		{102, "kidd", 78},
		{103, "lan", 98},
	}
	// 打印结构体数组的每一项
	for index, stu := range students {
		fmt.Println(index, stu.name)
	}
	fmt.Println(students)
	// 计算以上学生成绩的总分
	sum := students[0].score
	for i, stuLen := 1, len(students); i < stuLen; i++ {
		sum += students[i].score
	}
	fmt.Println("总分是：", sum)

	// 输出以上学生成绩中最高分
	maxScore := students[0].score
	for i, stuLen := 1, len(students); i < stuLen; i++ {
		if maxScore < students[i].score {
			maxScore = students[i].score
		}
	}
	fmt.Println("最高分是：", maxScore)
}

func main() {
	// 定义 map
	m := make(map[int]student)
	m[101] = student{101, "conan", 88}
	m[102] = student{102, "kidd", 78}
	m[103] = student{103, "lan", 98}
	fmt.Println(m) // map[101:{101 conan 88} 102:{102 kidd 78} 103:{103 lan 98}]

	for k, v := range m {
		fmt.Println(k, v)
	}
}
```

## 结构体切片作为 map 的 value

结构体切片（本质上是切片）作为 map 的 value 示例如下：

```go
package main

import "fmt"

type student struct {
	id    int
	name  string
	score int
}

func main() {
	m := make(map[int][]student)
    
	m[101] = append(m[101], student{1, "conan", 88}, student{2, "kidd", 78})
	m[102] = append(m[101], student{1, "lan", 98}, student{2, "blame", 66})

	// 101 [{1 conan 88} {2 kidd 78}]
	// 102 [{1 conan 88} {2 kidd 78} {1 lan 98} {2 blame 66}]
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k, v := range m {
		for i, data := range v {
			fmt.Println(k, i, data)
		}
	}
}
```

## 结构体作为函数参数

你可以像其它数据类型一样将结构体类型作为参数传递给函数：

**结构体传递为 值传递**（形参单元和实参单元是不同的存储区域，修改不会影响其它的值）

```go
package main

import "fmt"

type student struct {
	id    int
	name  string
	score int
}

func foo(stu student) {
	stu.name = "lan"
}

func main() {
	stu := student{101, "conan", 88}
	fmt.Println(stu)  // {101 conan 88}
	foo(stu)
	fmt.Println(stu)  // {101 conan 88}
}
```

通过以上程序，我们知道：Go 函数给参数传递值的时候是以复制的方式进行的。复制传值时，如果函数的参数是一个 struct 对象，将直接复制整个数据结构的副本传递给函数。

这有两个问题：

函数内部无法修改传递给函数的原始数据结构，它修改的只是原始数据结构拷贝后的副本；

如果传递的原始数据结构很大，完整地复制出一个副本开销并不小。

所以，**如果条件允许，应当给需要 struct 实例作为参数的函数传 struct 的指针**。

> PS：
>
> 1. 结构体切片作为函数参数是**地址传递**
> 2. 结构体数组作为函数参数是**值传递**



## 练习

定义结构体，存储5名学生，三门成绩，求出每名学生的总成绩和平均成绩。

结构体定义示例：

```go
type student struct {
    id int
    name string
    score []int
}
```

```go
package main

import "fmt"

type student struct {
	id    int
	name  string
	score []int
}

func main() {
	stus := []student{
		{101, "小明", []int{100, 99, 94}},
		{102, "小红", []int{60, 123, 98}},
		{103, "小刚", []int{90, 109, 81}},
		{104, "小强", []int{55, 66, 99}},
		{105, "小花", []int{123, 65, 89}},
	}
	for _, stu := range stus {
		// 三门总成绩
		sum := 0
		for _, value := range stu.score {
			sum += value
		}
		fmt.Printf("%s 的总成绩为: %d, 平均成绩为: %d\n", stu.name, sum, sum/len(stu.score))
	}
}
```

## 李培冠博客

欢迎访问我的个人网站：

李培冠博客：[lpgit.com](https://lpgit.com)
