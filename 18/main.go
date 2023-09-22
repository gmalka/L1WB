package main

import "fmt"

func main() {
	str := "главрыба"

	r := []rune(str)

	end := len(r) - 1
	for i := 0; i < len(r)/2; i++ {
		r[i], r[end] = r[end], r[i]

		end--
	}

	fmt.Println(string(r))
}