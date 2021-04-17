package main

import "fmt"

func main() {

	i := 100

	//0xc00001a0b8
	fmt.Println(&i)

	point := &i

	//*int
	fmt.Printf("%T\n", &i)

	// 100 指针指向的值
	fmt.Printf("%v\n", *point)

}
