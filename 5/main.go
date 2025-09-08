package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.

Подсказка: используйте time.After или таймер для ограничения времени работы.
*/

func main() {
	rand.Seed(time.Now().UnixNano())
	t := 10 * time.Second
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run main.go <количество_воркеров>")
		os.Exit(1)
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Ошибка: количество воркеров должно быть положительным числом")
		os.Exit(1)
	}

	dataChannel := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup
	go func() {
		<-time.After(t)
		fmt.Println("Время истекло!")
		os.Exit(1)
	}()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, done, &wg)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				randomNum := rand.Intn(250) + 1
				dataChannel <- randomNum
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	<-sigChan
	fmt.Println("\nПолучен сигнал завершения...")

	close(done)
	close(dataChannel)
	wg.Wait()
	fmt.Println("Программа завершена")
}

func worker(id int, dataChannel <-chan int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case num, ok := <-dataChannel:
			if !ok {
				fmt.Printf("Воркер %d: канал закрыт\n", id)
				return
			}
			fmt.Printf("Воркер %d: %d\n", id, num)
		case <-done:
			fmt.Printf("Воркер %d: получен сигнал завершения\n", id)
			return
		}
	}
}
