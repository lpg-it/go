关于字符与字符串的区别: 

**字符:** 

- 单引号
- 往往只包含一个字符, 转义字符除外: `\n`

**字符串:** 

- 双引号
- 字符串有一个或者多个字符组成
- 字符串都是隐藏了一个结束符: `\0`

下面通过代码来看一下两者的区别: 

```go
package main

import "fmt"

func main() {
	var a byte = 'a'
	var b string = "a"  // 'a' and '\0' 两个字符组成
	fmt.Println(a)  // 97
	fmt.Println(b)  // a
    
	// \n 换行, \\表示一个\, 一般用于文件操作
	fmt.Printf("%c\n", a)  // a
	var c string = "helloworld"
	fmt.Printf("%s", c)
	// fmt.Println(a == b)  // 不同类型不能比较
}
```

计算字符串的个数: 

```go
package main

import "fmt"

func main() {
	var s1 string = "hello world"
	// 计算字符串个数
	num := len(s1)
	fmt.Println(num)  // 11

	fmt.Printf("s1[0] = %c, s1[1] = %c\n", s1[0], s1[1])  // s1[0] = h, s1[1] = e

	// 在go语言中, 采用的是utf-8编码, 一个中文对应3个字符, 为了和linux统一处理
	var s2 string = "李培冠"
	num = len(s2)
	fmt.Println(num)  // 9

	var s3 string = "李培冠it"
	num = len(s3)
	fmt.Println(num)  // 11
}
```

## 李培冠博客

[lpgit.com](https://lpgit.com)