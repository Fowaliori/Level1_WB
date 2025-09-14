package main

import (
	"fmt"
)

/*
Разработать программу, которая переворачивает порядок слов в строке.

Пример: входная строка:

«snow dog sun», выход: «sun dog snow».

Считайте, что слова разделяются одиночным пробелом. Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».
*/

func reverseSent(s string) string {
	runes := []rune(s)

	reverseString(runes, 0, len(runes)-1)

	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			reverseString(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func reverseString(copyrune []rune, left, right int) {
	for left <= right {
		copyrune[left], copyrune[right] = copyrune[right], copyrune[left]
		left++
		right--
	}
}

func main() {
	str := "snow dog sun"
	fmt.Println(reverseSent(str))
}
