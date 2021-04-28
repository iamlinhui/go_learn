package main

import (
	"fmt"
	"reflect"
)

func main() {

	stu := student{
		Name: "ABC",
		Age:  10,
	}

	stuType := reflect.TypeOf(stu)

	fmt.Println(stuType)

	stuValue := reflect.ValueOf(stu)
	fmt.Println(stuValue)

	for i := 0; i < stuType.NumField(); i++ {
		field := stuType.Field(i)
		fmt.Println(field.Tag.Get("json"))
	}

}

type student struct {
	Name string `json:"userName" `
	Age  int
}
