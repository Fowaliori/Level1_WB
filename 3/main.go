package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

/*
Реализовать постоянную запись данных в канал (в главной горутине).

Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.
*/

func main() {
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

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, &wg)
	}

	for {
		randomNum := rand.Intn(250) + 1
		dataChannel <- randomNum
		time.Sleep(500 * time.Millisecond)
	}

}

func worker(id int, dataChannel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for message := range dataChannel {
		fmt.Printf("Воркер %d: %d\n", id, message)
	}

	fmt.Printf("Воркер %d завершил работу\n", id)
}
