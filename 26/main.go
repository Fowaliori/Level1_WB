package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке встречаются один раз (т.е. строка состоит из уникальных символов).

Вывод: true, если все символы уникальны, false, если есть повторения. Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.

Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.

Подумайте, какой структурой данных удобно воспользоваться для проверки условия.
*/

func hasAllUniqueChars(str string) bool {
	seen := make(map[rune]bool)

	for _, char := range strings.ToLower(str) {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func main() {
	testCases := []string{"abcd", "abCdefAaf", "aabcd", "12345", "11234", "Test"}

	for _, test := range testCases {
		fmt.Printf("'%s' -> %t\n", test, hasAllUniqueChars(test))
	}
}
