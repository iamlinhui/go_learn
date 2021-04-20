package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {

	var num int64 = 190

	fmt.Println(num)

	num2 := 21

	fmt.Println(num2)

	type myInt int

	var num3 myInt = 21

	fmt.Println(num3)

	result := sum(decimal.NewFromInt(11), decimal.NewFromFloat32(324.4353))

	fmt.Println(result)

}

func sum(values ...decimal.Decimal) decimal.Decimal {

	var result decimal.Decimal

	for _, d := range values {
		result = result.Add(d)
	}

	return result
}
