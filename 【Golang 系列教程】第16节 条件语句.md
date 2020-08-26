在Go语言中，条件语句主要包括有`if` 、 `switch` 与 `select`。

**注意：** Go语言中没有三目运算符，不支持 `?:` 形式的条件判断。

## if 语句 

### 最简单的if语句

最简单的 `if` 语句的基本语法：

```go
if 条件判断 {
    // 在当前条件判断为true时执行
}
```

条件判断如果为真（true），那么就执行大括号中的语句；如果为假（false），就不执行大括号中的语句，继续执行`if`结构后面的代码。

**值得注意的是：**Go语言规定与 `if` 匹配的左括号 `{` 必须与 `if和条件判断` 放在同一行。

#### 示例

```go
package main

import "fmt"

func main() {
    var year int = 2020
    
    if year > 1996 {
        // 如果条件为 true，则执行以下语句
        fmt.Printf("%d大于1996\n", year)
    }
    fmt.Println("year的值为: ", year)
}
```

执行结果为：

```go
2020大于1996
year的值为:  2020
```

### if...else语句

`if...else` 语句的基本语法：

```go
if 条件判断 {
    // 在当前条件判断为true时执行
} else {
    // 在当前条件判断为false时执行
}
```

条件判断如果为真（true），那么就执行其后紧跟的语句块；如果为假（false），则执行 `else` 后面的语句块。

**值得注意的是：**`else` 必须与上一个 `if` 右边的大括号在同一行；与 `else` 匹配的左括号 `{` 也必须与 `else` 卸载同一行。

#### 示例

```go
package main

import "fmt"

func main(){
	year := 2020

	if year > 1996 {
        // 如果条件为 true，则执行以下语句
        fmt.Printf("%d大于1996\n", year)
	} else {
		// 如果条件为 false，则执行以下语句
        fmt.Printf("%d小于1996\n", year)
	}
	fmt.Println("year的值为: ", year)
}
```

执行结果为：

```go
2020大于1996
year的值为:  2020
```

### if...else if ...else语句

`if...else if ...else` 语句的基本语法：

```go
if 条件判断1 {
    // 如果条件判断1为 true，则执行这里的语句
} else if 条件判断2 {
    // 如果条件判断2为 true，则执行这里的语句
} else {
    // 如果以上条件判断都为 false，则执行这里的语句
}
```

**同样的：**`else if` 必须与上一个 `if`  或者 `else if` 右边的大括号在同一行。

#### 示例

```
package main

import "fmt"

func main(){
	year := 2020
	
	if year > 2050 {
        fmt.Printf("%d大于2050\n", year)
	} else if year > 2000 {
		fmt.Printf("%d大于2000\n", year)
	} else {
        fmt.Println("year的值为: ", year)
	}
}
```

执行结果为：

```go
2020大于2000
```

### if嵌套语句

可以在以上语句中嵌套多个同样的语句，均是合法的。

在 `if语句` 中嵌套 `if语句` 的基本语法如下：

```
if 条件判断1 {
	// 在条件判断1为 true 时，执行这里的语句
	if 条件判断2 {
		// 在条件判断2为 true 时，执行这里的语句
	}
}
```

#### 示例

```go
package main

import "fmt"

func main(){
	year := 2020
	
	if year > 2000 {
		if year > 2010 {
			fmt.Println("year 大于2010.")
		}
	}
}
```

执行结果为：

```go
year 大于2010.
```

## switch语句

switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。

**注意：虽然说 `case` 表达式不能重复，但是如果 `case` 为布尔值，则可以重复。**

```go
package main

import "fmt"

func main() {
	a := false
	switch false {
	case a:
		fmt.Println("123")
	case a:
		fmt.Println("456")
	}
}
```

执行结果：

```go
123
```

下面来看一下一般的例子：

```go
package main

import "fmt"

func main(){
	date := 3
	switch date {
	case 1:
		fmt.Println("周一")
	case 2:
		fmt.Println("周二")
	case 3:
		fmt.Println("周三")
	case 4:
		fmt.Println("周四")
	case 5:
		fmt.Println("周五")
	case 6:
		fmt.Println("周六")
	case 7:
		fmt.Println("周日")
	default:
		fmt.Println("无效的输入")
	}
}
```

执行的结果：

```go
周三
```

Go语言规定每个 `switch` 只能有一个 `default` 分支。

一个分支可以有多个值，多个 `case` 值中间使用英文逗号分隔。

```go
package main

import "fmt"

func main(){
	num := 5
	switch num {
	case 1, 3, 5, 7, 9:
		fmt.Println("num是奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("num是偶数")
	default:
		fmt.Println("num：", num)
	}
}
```

执行的结果：

```go
num是奇数
```

当 `case` 分支后面使用的是表达式时，`switch` 语句后面不需要在跟判断变量。

```go
package main

import "fmt"

func main(){
	score := 61
	switch {
	case score > 80:
		fmt.Println("考得不错")
	case score >= 60:
		fmt.Println("努力学习吧")
	default:
		fmt.Println("还不学习？")
	}
}
```

执行结果：

```go
努力学习吧
```

`fallthrough`会强制执行后面的**一条**case语句。

```go
package main

import "fmt"

func main(){
	num := 1
	switch num {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("...")
	}
}
```

执行结果：

```go
1
2
```

我们使用 `fallthrough` 来执行多个 `case`，也可以使用 `break` 来终止。

```go
package main

import "fmt"

func main(){
	num := 1
	switch num {
	case 1:
		fmt.Println(1)
		if num == 1 {
			break
		}
		fallthrough
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("...")
	}
}
```

执行结果：

```go
1
```



## select语句

`select` 语句在后面会讲解。

## 李培冠博客

欢迎访问我的个人网站：

李培冠博客：[lpgit.com](https://lpgit.com)
