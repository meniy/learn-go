package main

import "fmt"

// pair struct di due elementi di tipo string
type pair struct {
	x, y string
}

type Vertex struct {
	X int
	Y string
}

// dynamicPair le due variabili che compongono la struct possono essere di diverso tipo
type dynamicPair struct {
	x, y interface{}
}

// anonymouse fields
type anonyPair struct {
	int
	string
}
type pair2 struct {
	pair
	int
	string
}

type Salaried interface {
	getSalary() int
}
type Salary struct {
	basic     int
	insurance int
}

type Employee struct {
	string
	Salaried
}

func (s Salary) getSalary() int {
	return s.basic + s.insurance
}

func main() {
	p := pair{"a", "b"}
	fmt.Println("P:", p)
	fmt.Println("x:", p.x)
	fmt.Println("y:", p.y)
	p.x = "c"
	// p.x = 1  non si puo cambiare il tipo
	fmt.Println("x:", p.x)

	d := dynamicPair{1, 2}
	fmt.Println("x:", d.x)
	d.x = "a"
	fmt.Println("x:", d.x)
	fmt.Println(d)

	v := Vertex{10, "x"}
	fmt.Println("v:", v)

	v2 := Vertex{X: 10, Y: "ciao"}
	fmt.Println("v2:", v2)

	// ANONYMOUSE STRUCT
	user := struct {
		email, psw string
	}{email: "l@c.com", psw: "asdasd"}

	fmt.Println(user)

	userAnon := anonyPair{1, "c"}
	fmt.Println(userAnon)

	pp := pair2{
		pair{x: string("ciao"), y: string("ciao")},
		1, string("riciao"),
	}
	fmt.Println(pp)
	e := Employee{"a", Salary{1, 2}}
	fmt.Println(e.getSalary())

}
