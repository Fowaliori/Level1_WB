package main

import (
	"fmt"
)

/*
Разработать программу, которая в runtime способна определить тип переменной, переданной в неё (на вход подаётся interface{}).

Типы, которые нужно распознавать: int, string, bool, chan (канал).

Подсказка: оператор типа switch v.(type) поможет в решении.
*/

func getType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan any:
		fmt.Println("chan")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	getType(2)
	getType("Home")
	getType(false)
	getType(make(chan any))
}
