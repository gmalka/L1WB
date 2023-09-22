package main

import "fmt"

func main() {
	fmt.Println(QuickSort([]int{1, 1, 2, 3, 5, 1, 2, -1}))
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var (
		l []int
		r []int
		cur []int
	)

	mid := len(arr) / 2

	for i := 0; i < len(arr); i++ {
		if arr[i] < arr[mid] {
			l = append(l, arr[i])
		} else if arr[i] > arr[mid] {
			r = append(r, arr[i])
		} else {
			cur = append(cur, arr[i])
		}
	}

	result := QuickSort(l)
	result = append(result, cur...)
	result =append(result, QuickSort(r)...)

	return(result)
}