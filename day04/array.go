package main

import (
	"fmt"
)

func main() {

	var array = [3]int{}

	var s = array[1:3]

	fmt.Println(s)
	s[0] = 1000
	s = append(s, 12, 32, 43, 32)

	s[0] = 20000

	fmt.Println(s)
	fmt.Println(array)

	var b []int

	b[0] = 19

	fmt.Println(b)
	b = append(b, 8, 9, 7, 6)
	fmt.Println(b)

}
