package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", n)
			m.Store(key, n*100)
		}(i)
	}

	wg.Wait()

	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%s: %d\n", key, value)
		return true
	})
}
