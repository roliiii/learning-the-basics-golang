package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["one"] = 1
	x["two"] = 2
	x["ten"] = 10
	fmt.Println(x["ten"])

	delete(x, "one")
	fmt.Println(x)

	if number, ok := x["ten"]; ok {
		fmt.Println(number, ok)
	}

	g := map[string]int{
		"ten": 10,
		"one": 1,
	}
	for key, value := range g {
		fmt.Printf("%s: %d\n", key, value)
	}

	y := map[string]map[string]string{
		"ten": map[string]string{
			"number": "10",
			"roman":  "x",
		},
	}
	fmt.Println(y["ten"]["roman"])
}
