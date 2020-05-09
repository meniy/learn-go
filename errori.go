package main

import (
	"fmt"
	"strconv"
)


func main() {
	// An error value communicates not just "ok" but more about the problem.
	if _, err := strconv.Atoi("non-int"); err != nil { // _ discards value
		// prints 'strconv.ParseInt: parsing "non-int": invalid syntax'
		fmt.Println(err)
	}

	if val, err := strconv.Atoi("10"); err != nil { // _ discards value
		// prints 'strconv.ParseInt: parsing "non-int": invalid syntax'
		fmt.Println(err)
	} else{
		fmt.Println("else:")
		fmt.Println(val)
	}
}
