package main

import "fmt"

/*
Поменять местами два числа без использования временной переменной.

Подсказка: примените сложение/вычитание или XOR-обмен.
*/

func main() {
	a := 25
	b := 15
	a += b
	b = a - b
	a = a - b
	fmt.Println(a, b)
	a = 25
	b = 15
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}
