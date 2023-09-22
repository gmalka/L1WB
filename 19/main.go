package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog  sun"

	strs := strings.Split(str, " ")

	end := len(strs) - 1

	for i := 0; i < len(strs) / 2; i++ {
		strs[i], strs[end] = strs[end], strs[i]
	}

	str = strings.Join(strs, " ")

	fmt.Println(str)
}