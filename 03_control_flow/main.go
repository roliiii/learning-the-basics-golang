package main

import (
	"fmt"
	"time"
)

func main() {
	ifStatement()
	forLoop()
	switchStatement()
}

func ifStatement() {
	if true {
		fmt.Println("normal if")
	}
	if x := 1; x == 1 {
		fmt.Println("Short statement before condition")
	}
}

func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//while
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	for {
		if sum > 1000 {
			break
		}
	}

	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("Even: %d\n", i)
	}

}

func switchStatement() {
	mins := time.Now().Minute()
	switch mins {
	case 0:
		fmt.Println("0")
	case 1, 2, 3:
		fmt.Println("1,2,3")
	case 2 + 2:
		fmt.Println("4")
	default:
		fmt.Println("4-59")
	}

	i := 0
	switch i {
	case 0:
		fmt.Print("fall")
		fallthrough //<--fallthrough
	case 1:
		fmt.Println("trough")
	case 2:
		fmt.Println("stop")
	}
}
