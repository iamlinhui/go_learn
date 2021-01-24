package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"unsafe"
)

/**
数据类型
*/
func main() {

	// 默认整型位int
	i := 100
	// 32位系统4字节 或者 64位系统8个字节|-2^31～2^31-1 或者 -2^63～2^63-1
	fmt.Printf("类型:%T\n 占用字节数:%d\n", i, unsafe.Sizeof(i))

	// decimal 大数据类型运算
	result := decimal.NewFromInt(1).DivRound(decimal.NewFromInt(3), 128).Round(2)
	result = result.Mul(decimal.NewFromInt(3))
	fmt.Println(result)

	j := 100.21 //float64 占用字节数:8
	fmt.Printf("类型:%T\n 占用字节数:%d\n", j, unsafe.Sizeof(j))

	// 科学计数法
	k := 2.31E-5
	m := 3.1456e10
	n := 8.43e-21
	fmt.Println(k, m, n)

}
