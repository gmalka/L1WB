package main

import "fmt"

func main() {
	first, second := 1, 2

	// Сначала вычисляется правая часть, затем она присваивается левой части
	first, second = second, first

	fmt.Println(first, second)
}