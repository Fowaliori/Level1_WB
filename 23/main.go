package main

import "fmt"

/*
Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.

Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента (copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.
*/

func del(arr []int, i int) []int {
	arr[i] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	return arr
}

func main() {
	i := 2
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(len(arr))
	arr = del(arr, i)
	fmt.Println(arr)
	fmt.Println(len(arr))
}
