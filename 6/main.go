package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	stopByCondition()
	stopByChannel()
	stopByContext()
	stopByGoexit()
	stopByDataChannel()
}

// Остановка по условию
func stopByCondition() {
	var wg sync.WaitGroup
	stop := false

	wg.Add(1)
	go func() {
		defer wg.Done()
		counter := 0
		for !stop {
			counter++
			fmt.Printf("Условие: работает (%d)\n", counter)
			time.Sleep(100 * time.Millisecond)

			if counter >= 3 {
				stop = true
			}
		}
		fmt.Println("Условие: горутина завершена")
	}()

	wg.Wait()
}

// Остановка через канал уведомления
func stopByChannel() {
	var wg sync.WaitGroup
	stopChan := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		counter := 0
		for {
			select {
			case <-stopChan:
				fmt.Println("Канал: получен сигнал остановки")
				return
			default:
				counter++
				fmt.Printf("Канал: работает (%d)\n", counter)
				time.Sleep(100 * time.Millisecond)

				if counter >= 3 {
					close(stopChan)
				}
			}
		}
	}()

	wg.Wait()
}

// Остановка через контекст
func stopByContext() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()
		counter := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Контекст: получен сигнал отмены")
				return
			default:
			}

			counter++
			fmt.Printf("Контекст: работает (%d)\n", counter)
			time.Sleep(100 * time.Millisecond)
			if counter >= 3 {
				cancel()
			}
		}
	}()

	wg.Wait()
}

// Остановка через runtime.Goexit()
func stopByGoexit() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		counter := 0
		for {
			counter++
			fmt.Printf("Goexit: работает (%d)\n", counter)
			time.Sleep(100 * time.Millisecond)

			if counter >= 3 {
				fmt.Println("Goexit: вызывается runtime.Goexit()")
				runtime.Goexit()
			}
		}
	}()

	wg.Wait()
}

// Остановка через канал для получения данных
func stopByDataChannel() {
	var wg sync.WaitGroup
	dataChan := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(dataChan)

		for i := 1; i <= 5; i++ {
			dataChan <- i
			fmt.Printf("Данные: отправлено %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range dataChan {
			fmt.Printf("Данные: получено %d\n", data)
			if data >= 3 {
				fmt.Println("Данные: достигнут лимит, продолжаем чтение до закрытия")
			}
		}
		fmt.Println("Данные: канал закрыт, горутина завершена")
	}()

	wg.Wait()
}
