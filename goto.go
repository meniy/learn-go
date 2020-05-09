package main

import "fmt"

func main() {

	fmt.Println("Goto Example")
	goto routine
	fmt.Println("Unreachable code")


routine:
	fmt.Println("My routine!")
}
