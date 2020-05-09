package main

import "fmt"

func main() {
	m := map[int]string{3: "three", 4: "four"}
	if x, ok := m[3]; !ok { // ok will be false because 1 is not in the map.
		fmt.Println("no one there")
	} else {
		fmt.Println("find it! ", x) // x would be the value, if it were in the map.
	}
	if x, ok := m[1]; !ok { // ok will be false because 1 is not in the map.
		fmt.Println("no one there")
	} else {
		fmt.Println("find it! ", x) // x would be the value, if it were in the map.
	}

	if ok := 2; ok <= 5 {
		fmt.Println(("ok"))
	} else {
		fmt.Println("no!")
	}
	v := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, value := range v {
		if value < 4 {
			fmt.Println(value, "minore di 4")
		}
	}

	localFun := func(a int) func(b int) (r int) {
		return func(b int) (r int) {
			return a + b
		}
	}

	lista := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for pos, value := range lista {
		fmt.Println("Value:",value)
		lf := localFun(value)
		if pos+1 < len(lista) {
			fmt.Println(lf(lista[pos+1]))
		}
	}

}
