package models

import "go-resto/internal/models/constant"



type MenuItem struct {
	Name      string
	OrderCode string
	Price     int64
	Type      constant.MenuType
}
