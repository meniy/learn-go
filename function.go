package main

import "fmt"

// helloWorld stampa a video Hello World!
func helloWorld() () {
	fmt.Println("Hello World!")
}

//multiple Restituisce la somma e il prodotto di due valori interi
func multiple(a, b int) (sum, prod int) {
	return a + b, a * b
}

func sum(a, b int) (s int) {
	return a + b
}
func prod(c, d int) (p int) {
	return c * d
}

func sentenceFactory(mystring string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, mystring, after) // new string
	}
}

// sumValue value1 viene fissato ad un valore inter, value2 si puo,
// richiamanda la funzione variare
func sumValue(value1 int) func(value2 int) int {
	return func(value2 int) int {
		return value1 + value2 // new string
	}
}

// Vedi struct
//func (s Salary) getSalary() int {
//	return s.basic + s.insurance
//}

// Functions can have variadic parameters.
func learnVariadicParams(myStrings ...interface{}) {
	// Iterate each value of the variadic.
	// The underbar here is ignoring the index argument of the array.
	for _, param := range myStrings {
		fmt.Println("param:", param)
	}

	// Pass variadic value as a variadic parameter.
	fmt.Println("params:", fmt.Sprintln(myStrings...))

}

func main() {
	helloWorld()
	//a:=1
	//b:=2
	sumRes, prodRes := multiple(1, 2)
	// sumRes, prodRes := multiple(a, b)
	fmt.Println("sum:", sumRes, "prod:", prodRes)

	//  anonymouse function in a print
	fmt.Println("Add + double two numbers: ",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2))

	// Function literals are closures.
	x := 10
	xBig := func() bool {
		return x > 10000 // References x declared above switch statement.
	}
	fmt.Println("xBig:", xBig())

	fmt.Println(sentenceFactory("summer")("A beautiful", "day!"))
	// altro esempio
	sumFun := sumValue(1)   // 1
	fmt.Println(sumFun(1))  // 1 + 1 = 2
	fmt.Println(sumFun(2))  // 2 +1 = 3
	fmt.Println(sumFun(10)) // 10 +1 = 11

	learnVariadicParams("great", "learning", "here!") //variadic params


}
