package main

import (
	"fmt"
	"slices"
)

/*
Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е. вывести элементы, присутствующие и в первом, и во втором.

Пример:
A = {1,2,3}
B = {2,3,4}
Пересечение = {2,3}
*/

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}
	C := []int{}
	for _, i := range A {
		if slices.Contains(B, i) {
			C = append(C, i)
		}
	}
	for _, i := range C {
		fmt.Println(i)
	}
}
