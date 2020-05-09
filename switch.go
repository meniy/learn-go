package main

import "fmt"

func main() {
	val := 3
	switch val {
	case 2:
		fmt.Println("2")
	case 1:
		fmt.Println("1")
	case 3,4:
		fmt.Println("3 o 4")
	default:
		fmt.Println(val)
	}
}
