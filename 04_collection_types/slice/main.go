package main

import "fmt"

func main() {
	//var x []string
	x := make([]string, 3, 5) //length, capacity
	x[0] = "H"
	x[1] = "el"
	x[2] = "lo"

	fmt.Println(x[:1])
	fmt.Println(x[1:])
	fmt.Printf("len: %d\n", len(x)) //3
	fmt.Printf("cap: %d\n", cap(x)) //5

	y := append(x, "world")
	fmt.Println(y)
	fmt.Printf("len: %d\n", len(y)) //4
	fmt.Printf("cap: %d\n", cap(y)) //5

	z := make([]string, 1)
	copy(z, y)
	fmt.Println(z)
	fmt.Printf("len: %d\n", len(z)) //1
	fmt.Printf("cap: %d\n", cap(z)) //1

	u := []string{"apple", "grape"}
	u = append(u, x...)
	fmt.Println(u)
	fmt.Printf("len: %d\n", len(u)) //5
	fmt.Printf("cap: %d\n", cap(u)) //5
}
