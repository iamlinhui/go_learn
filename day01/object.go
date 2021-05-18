package main

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"go-learn/day01/model"
)

func main() {

	var books []*model.Book

	book := &model.Book{
		ID:        12,
		Title:     "story",
		Author:    "AA",
		Price:     decimal.NewFromFloat(323.43),
		Sales:     434,
		Stock:     343,
		ImagePath: "",
	}
	books = append(books, book)

	books = append(books, &model.Book{
		ID:        21,
		Title:     "class",
		Author:    "BB",
		Price:     decimal.NewFromFloat(2.43),
		Sales:     32435,
		Stock:     12,
		ImagePath: "",
	})

	for _, item := range books {
		// Json序列号
		b, _ := json.Marshal(item)
		str := string(b)
		fmt.Println(str)

		var stu2 *model.Book
		// Json反序列化
		_ = json.Unmarshal(b, &stu2)
		fmt.Println(stu2)
	}
	booksByte, _ := json.Marshal(books)
	fmt.Println(string(booksByte))

}
