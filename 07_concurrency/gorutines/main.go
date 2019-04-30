package main

import (
	"fmt"
	"time"
)

func main() {
	go say("word")
	say("Hello ")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
