package main

import (
	"fmt"
	"reflect"
)

/**
变量 --传参--> Interface
变量 <--类型断言-- Interface

Interface --reflect.Value()函数--> reflect.Value
Interface <--stuValue.Interface()函数-- reflect.Value
*/
func main() {

	stu := student{
		Name: "ABC",
		Age:  10,
	}

	stuType := reflect.TypeOf(stu)

	fmt.Println(stuType)

	// struct --> reflect.Value
	stuValue := reflect.ValueOf(stu)
	fmt.Println(stuValue)

	// reflect.Value -->  Interface
	iVal := stuValue.Interface()
	fmt.Println(iVal)

	// Interface --> struct
	v := iVal.(student)
	fmt.Println(v)

	for i := 0; i < stuType.NumField(); i++ {
		field := stuType.Field(i)
		fmt.Println(field.Tag.Get("json"))
	}

}

type student struct {
	Name string `json:"userName" `
	Age  int
}
