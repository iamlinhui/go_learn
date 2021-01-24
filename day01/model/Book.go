package model

import "github.com/shopspring/decimal"

type Book struct {
	ID        int
	Title     string          //书名
	Author    string          //作者
	Price     decimal.Decimal //价格
	Sales     int             //已销售
	Stock     int             //库存
	ImagePath string          //图片路径
}
