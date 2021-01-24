package main

import "fmt"

/**
变量的使用方式
变量 = 变量名 + 数值 + 数据类型
*/
func main() {

	// 指定变量类型，声明后不赋值，使用默认值（int 默认值为0 string默认值为空串 小数默认值为0）
	var i int
	fmt.Println(i)

	// 根据值自行判断变量类型（类型推导）
	var j = 12.43
	fmt.Println(j)

	// 省略var,:=左侧的变量不能是已经声明过了的，否则会编译错误
	name := "Tom"
	fmt.Println(name)

	// 声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	a, b, c := 100, "a", 12.32
	fmt.Println(a, b, c)

	var (
		x = 1
		y = "v"
		z = 32.344444
	)
	fmt.Println(x, y, z)

}
