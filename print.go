package main

import "fmt"

func main() {
	myStrings := []string{"Hello", "Dynamic", "Print!", "(con Sprintln)"}
	myStrings2 := []string{"Hello", "Dynamic", "Print!", "(senza Sprintln)"}

	fmt.Print("Hello Print!") // Non va a capo
	fmt.Print("Hello Print!\n") // Va a capo per il \n
	fmt.Println("Hello Println!") // Va a capo di default

	fmt.Println(fmt.Sprintln(myStrings)) // Va a capo due volte
	fmt.Println("testo dopo una riga vuota")
	fmt.Println(myStrings2)
	fmt.Println("testo a capo senza una riga vuota")

}
