Go 语言中的函数名、变量名、常量名、类型名、语句标号和包名等所有的命名，都遵循一个简单的命名规则。

> 必须以一个字母或者下划线（_）开头，后面可以跟任意数量的字母、数字或下划线。

在 Go 语言中，大小写字母是不同的。

Go 语言中有 25 个关键字，不能用于自定义名字：

```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

- 变量只能由字母、数字、下划线组成。
- 不能以数字开头。
- 不能是Go中的关键字及保留字
- 大小写区分，`a := 1`和 `A := 1`是两个变量。

以上要求是必须满足的，下面的要求要尽量做到

- 变量名要有描述性，要简洁、易读，不宜过长。
- 专有名词通常全部大写，例如：escapeHTML。
- 局部变量优先使用短名（用 i 代替 index）。
- 变量名不能使用中文以及拼音。
- 推荐使用的变量名：
    - 驼峰体：`MyName := "Conan"` 或 `myName := "Conan"`

Go语言中的 37 个保留字:

```go
Constants:    true  false  iota  nil
    Types:    int  int8  int16  int32  int64
			  uint  uint8  uint16  uint32  uint64  uintptr
			  float32  float64  complex128  complex64
			  bool  byte  rune  string  error
Functions:    make  len  cap  new  append  copy  close  delete
			  complex  real  imag
			  panic  recover
```

## 李培冠博客

[lpgit.com](https://lpgit.com)