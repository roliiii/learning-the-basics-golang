package main

import "fmt"

var g int

func main() {
	const LEGTH int = 10
	const (
		pi = 3.14
		//...
	)

	var a int
	a = 1
	fmt.Println(a)

	var b = 2
	fmt.Println(b)

	var c = 1 + 2i
	fmt.Println(c)

	var d, e, f = 1, 2, "foo"
	fmt.Printf("d: %d %T\n", d, d)
	fmt.Printf("e: %d %T\n", e, e)
	fmt.Printf("f: %s %T\n", f, f)

	//swap
	fmt.Printf("d: %d e:%d ->", d, e)
	d, e = e, d
	fmt.Printf("d: %d e:%d\n", d, e)

	//shorthand
	name, age := "Roland", 25
	fmt.Println(name, age)

	definingMultipleVariables()

	funcInVariable := func() {
		fmt.Printf("Hello from function")
	}

	funcInVariable()

}

func definingMultipleVariables() {
	var (
		name = "Roland"
		age  = 25
	)

	fmt.Println(name)
	fmt.Println(age)
}
