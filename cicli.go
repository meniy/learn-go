package main

import "fmt"

func main() {
	// Ciclo con posizione e valore
	for pos, value := range []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0} {
		fmt.Println("Pos:", pos)
		fmt.Println("Value:", value)
	}
	// Ciclo con solo il valore
	for _, value := range []string{"hello", " ", "world"} {
		fmt.Print(value)
	}

	//Ciclo con l uso di una map
	for key, value := range map[string]int{"one": 1, "two": 2} {
		fmt.Println(key, ":", value)
	}

	//Ciclo con l uso di una map, no chiave
	for _, value := range map[string]string{"primo": "Hello", "secondo": " ", "terzo": "World"} {
		fmt.Print(value)
	}
}
