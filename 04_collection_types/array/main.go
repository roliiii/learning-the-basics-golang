package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Word"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	x := [5]int{15, 1337, 324, 412, 21}
	for i, value := range x {
		fmt.Printf("%d: %d\n", i, value)
	}

	y := [...]float32{1.1, 2.5, 3.14, 4, 5}
	fmt.Println(y[1:4])
}
