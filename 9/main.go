package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2.

После этого данные из второго канала должны выводиться в stdout.

То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка. Убедитесь, что чтение из второго канала корректно завершается.
*/

func main() {
	in := make(chan int)
	out := make(chan int)
	arr := [50]int{}
	for i := range arr {
		arr[i] = i
	}
	go func() {
		for _, i := range arr {
			in <- arr[i]
		}
		close(in)
	}()
	go func() {
		for j := range in {
			out <- j * 2
		}
		close(out)
	}()
	for k := range out {
		fmt.Println(k)
	}
}
