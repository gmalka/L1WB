package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	i := 3
	if i >= 0 && i < len(arr) {
		arr = append(arr[:i], arr[i+1:]...)
	}
	fmt.Println(arr)
}