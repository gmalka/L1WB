package main

import "fmt"

func main() {
	var t interface{} = "twgweg"

	switch t.(type) {
	case int:
		fmt.Println("It is Int")
	case string:
		fmt.Println("It is String")
	case bool:
		fmt.Println("It is Boolean")
	case chan int:
		fmt.Println("It is Chan int")
	default:
		fmt.Println("Unknown type")
	}
}