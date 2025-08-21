package main

import (
	"fmt"
	"sync"
)

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
