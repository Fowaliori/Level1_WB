package main

import "fmt"

/*
Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.

Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел. Для выбора опорного элемента можно взять середину или первый элемент.
*/

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var low, mid, high []int

	for _, v := range arr {
		if v < pivot {
			low = append(low, v)
		} else if v > pivot {
			high = append(high, v)
		} else {
			mid = append(mid, v)
		}
	}

	low = quickSort(low)
	high = quickSort(high)

	return append(append(low, mid...), high...)
}

func main() {
	arr := []int{3, -5, 3, 4, 9, 9, 9, 7, 2, 8}
	fmt.Println(quickSort(arr))
}
