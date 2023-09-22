package main

import "fmt"

func main() {
	fmt.Println(BinarySearch([]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29}, 13))
}

func BinarySearch(arr []int, finded int) (int, bool) {
	if len(arr) <= 1 {
		for i, v := range arr {
			if v == finded {
				return i, true
			}
		}

		return 0, false
	}
	mid := len(arr) / 2

	if arr[mid] == finded {
		return mid, true
	} else if arr[mid] < finded {
		v, t := BinarySearch(arr[mid:], finded)
		if t {
			return v + mid, t
		}
	} else {
		v, t := BinarySearch(arr[:mid], finded)
		if t {
			return v, t
		}
	}

	return 0, false
}