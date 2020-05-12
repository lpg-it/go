// 导入主函数的包
package main

// Goland 会自动导入所需要的包（一系列功能和函数的集合）
// format：标准输入输出格式包
import "fmt"

func main0201() {
	// 单行注释
	fmt.Println("Hello Go")  // 将信息输出到屏幕上

	/*
	多行注释
	像这样，
	注释多行
	 */
}

// 注释不参与程序编译，可以帮助理解程序
// main 主函数，是程序的主入口，程序有且只有一个主函数
/*
笔试题知识点：
main 函数不能带参数
main 函数所在的包必须为 main 包
main 函数不能定义返回值
 */

func main() {
	// 在屏幕中打印 Hello Go，ln 表示换行
	fmt.Println("Hello Go")
}
