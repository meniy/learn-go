package main

import "fmt"

func learnMemory() (p, q *int) {
	s := make([]int, 20) // 20ints as single block of memory
	fmt.Println("20ints:",s)
	s[3] = 7
	r := -2
	fmt.Println("Indirizzo di memoria di s[3]:",&s[3])
	fmt.Println("Indirizzo di memoria di r:",&r)
	return &s[3], &r
}

func main() {
	// learn memory
	p, q := learnMemory() // Declares p, q to be type pointer to int.
	fmt.Println("Indirizzo della cella:",&p,"punta all indirizzo:",p, "che contiene il valore:",*p)
	fmt.Println("Indirizzo della cella:",&q,"punta all indirizzo:",q,"che contiene il valore:",*q)

}
