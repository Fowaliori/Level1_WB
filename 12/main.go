package main

import (
	"fmt"
	"slices"
)

/*
Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.

Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.
*/

func main() {
	A := []string{"cat", "cat", "dog", "cat", "tree"}
	C := []string{}
	for _, i := range A {
		if slices.Contains(C, i) {
			continue
		} else {
			C = append(C, i)
		}
	}
	for _, i := range C {
		fmt.Println(i)
	}
}
