package main

import (
	"fmt"
	"os"
)

func main() {
	// file
	s := "Hello"
	var s2 = string("Hello2")
	list := [...]string{"Ciao", "Hello", "Hi"}
	file, _ := os.Create("output.txt")
	// Sovrascirve
	//fmt.Fprint(file, "Wooow!")
	fmt.Fprintln(file, s)
	fmt.Fprintln(file, s2)
	fmt.Fprintln(file, list)
	file.Close()
	fmt.Println(func() (a int) {
		return 1 + 2
	}())
}
