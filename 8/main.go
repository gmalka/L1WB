package main

import (
	"fmt"
)

// Первая горутина, которая ситывает данные из массива и отправяет их в канал
func foo(arr []int) <-chan int {
	first := make(chan int)

	go func() {
		for _, v := range arr {
			first <- v
		}
		close(first)
	}()

	return first
}

// Вторая горутина, которая читает из первого канала, умножает данные на 2 и отправляет во второй канал
func foo2(first <-chan int) <-chan int {
	second := make(chan int)

	go func() {
		for v := range first {

			second <- 2 * v
		}
		close(second)
	}()

	return second
}

func main() {
	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := foo(m)
	ch2 := foo2(ch)

	// Выводим данные из второго канала
	for v := range ch2 {
		fmt.Println(v)
	}
}
