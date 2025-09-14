package main

import (
	"fmt"
)

/*
Разработать программу, которая переворачивает подаваемую на вход строку.

Например: при вводе строки «главрыба» вывод должен быть «абырвалг».

Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.), то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).
*/

func reverseString(str string) string {
	copyrune := []rune(str)
	left := 0
	right := len(copyrune) - 1
	for left <= right {
		copyrune[left], copyrune[right] = copyrune[right], copyrune[left]
		left++
		right--
	}
	return string(copyrune)
}

func main() {
	str := "главрыба"
	fmt.Println(reverseString(str))
}
