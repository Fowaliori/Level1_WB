package main

import (
	"fmt"
	"sync"
)

/*
Реализовать безопасную для конкуренции запись данных в структуру map.

Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).

Проверьте работу кода на гонки (util go run -race).
*/

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
