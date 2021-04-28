package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// 将NewInt定义为int类型
type NewInt int

// 将int取一个别名叫IntAlias
type IntAlias = int

func main() {

	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名
	fmt.Printf("a type: %T\n", a) // a type: main.NewInt
	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2) // a2 type: int

	// var now LocalTime = LocalTime(time.Now())
	var now = LocalTime(time.Now())
	data, _ := json.Marshal(now)
	fmt.Println(string(data)) //"2021-04-28 01:13:55"

	var createTime LocalTime
	_ = json.Unmarshal([]byte("\"2021-04-27 23:58:50\""), &createTime)
	fmt.Println(createTime) //2021-04-27 23:58:50
}
