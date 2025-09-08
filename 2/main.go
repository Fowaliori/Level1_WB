package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значения квадратов чисел, взятых из массива [2,4,6,8,10], и выведет результаты в stdout.

Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.
*/

var wg sync.WaitGroup

func square(a int) {
	fmt.Println(a * a)
	wg.Done()
}

func main() {

	arr := [5]int{2, 4, 6, 8, 10}

	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go square(arr[i])
	}
	wg.Wait()
}
