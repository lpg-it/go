比如想要把`int`类型的`123`转为`string`类型的`"123"`, 应该如何操作呢？

如果按照下面的写法. 那么我会得到`"{"`, 而不是`"123"`。

```go
package main

import "fmt"

func main() {
	i := 123
	s := string(i)
	fmt.Println(s)
}
```

可以使用`strconv`包中的`Itoa`功能。

例如：

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 123
	s := strconv.Itoa(i)
	fmt.Println(s)
}
```

有人认为`Itoa`这个名字很难记，为什么不用一个更具有描述性的名字呢？

我们可以这样记，就比较容易记住了。

`Itoa`：**ItoA** - **整数** 转 **A**SCII码

这样就比较容易记住了。