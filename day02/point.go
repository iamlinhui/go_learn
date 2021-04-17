package main

import "fmt"

func main() {

	var i int = 10

	var name = 9899

	var ptr *int = &i

	fmt.Println(*ptr)
	fmt.Println(ptr)

	//*ptr = 100

	ptr = &name

	fmt.Println(i)

	fmt.Println(*ptr)
	fmt.Println(ptr)

}
