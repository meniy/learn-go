package main

import "fmt"

func main() {
	fmt.Println("###### Array in GO #########")
	v1 := [...]int{1, 2, 3}
	fmt.Println("v1:", v1)
	v2 := v1
	fmt.Println("v2:", v2)
	v1[0] = 10
	v2[0] = 0 // modificando v2, v1 non cambia
	fmt.Println("v1:", v1, "v1[0]:", v1[0])
	fmt.Println("v2:", v2)

	v3 := make([]int, 4) // 4 ints, initialized to all 0
	fmt.Println("v3:", v3)

	v4 := [][]float32{{0,0}, {4}}
	fmt.Println("v4:", v4)

	v5 := []byte("a slice")
	fmt.Println("v5:", v5)
}
