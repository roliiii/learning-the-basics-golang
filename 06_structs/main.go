package main

import (
	"fmt"
	"math"
)

func main() {

	var circle1 Circle
	circle2 := new(Circle) // returns a pointer. (*Circle)
	circle3 := Circle{x: 0, y: 0, r: 0}
	circle4 := Circle{0, 0, 0}

	fmt.Println(circle1) //{0 0 0}
	fmt.Println(circle2) //&{0 0 0}
	fmt.Println(circle3) //{0 0 0}
	fmt.Println(circle4) //{0 0 0}

	var c Circle
	c.x, c.y, c.r = 1, 13, 7
	fmt.Println(c) //{1 13 7}

	circleCopied(c)
	fmt.Println(c) //{1 13 7}

	circleByPointer(&c)
	fmt.Println(c) //{1 13 6}

	fmt.Println(c.print())

	var circle5 Area = &Circle{r: 7}
	fmt.Println(circle5.getArea())

}

type Area interface {
	getArea() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) print() string {
	return fmt.Sprintf("x: %f, y: %f, r: %f", c.x, c.y, c.r)
}

func (c *Circle) getArea() float64 {
	return math.Pi * c.r * c.r
}

func circleCopied(c Circle) {
	c.r = 6
}

func circleByPointer(c *Circle) {
	c.r = 6
}
