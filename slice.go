package main

import "fmt"

func main() {
	fmt.Println("###### Slice in GO #########")
	// Slice have dynamic size
	s3 := []int{4, 5, 6}
	fmt.Println("s3:", s3)
	s3 = append(s3, 1, 2, 3)
	fmt.Println("s3:", s3)
	s4 := append(s3, []int{10, 11, 12}...)
	fmt.Println("s4:", s4)
	s5 := []int{13, 14, 15}
	s5 = append(s3, s4...)
	fmt.Println("s5:", s5)
	s6 := []string{"hello", "world"}
	fmt.Println("s6:", s6)

}
