package main

import "fmt"

/*
Реализовать алгоритм бинарного поиска встроенными методами языка.
Функция должна принимать отсортированный слайс и искомый элемент, возвращать индекс элемента или -1, если элемент не найден.

Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.
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

func binarySearch(arr []int, num int) int {
	if len(arr) == 0 {
		return -1
	}
	mid := len(arr) / 2
	if arr[mid] == num {
		return mid
	} else if num < arr[mid] {
		return binarySearch(arr[:mid], num)
	} else {
		res := binarySearch(arr[mid+1:], num)
		if res == -1 {
			return -1
		}
		return mid + 1 + res
	}
}

func main() {
	arr := []int{3, 1, 3, 4, 9, 9, 9, 7, 2, 8}
	sortArr := quickSort(arr)
	fmt.Println(binarySearch(sortArr, 2))
}
