package main

import "fmt"

func main() {
	defer lastMethodCall()
	fmt.Println(average([]float64{1.2, 14}))
	fmt.Println(namedReturn(1))

	a, b := returnMultipleValues(1)
	fmt.Printf("a:%d b:%t \n", a, b)

	fmt.Println(variadicAdd(1, 2, 4, 8))
	fmt.Println(funcInFunc())

	funcReturnedByFunc := funcReturnFunc()
	fmt.Println(funcReturnedByFunc())
	fmt.Println(funcReturnedByFunc())
	fmt.Println(funcReturnedByFunc())

	defer catchThePanic()
	makePanic()
	fmt.Println("No-No")

}

func catchThePanic() {
	str := recover()
	fmt.Println(str, "asd")
}

func lastMethodCall() {
	fmt.Println("----")
}

func average(nums []float64) float64 {
	total := 0.0
	for _, actNumber := range nums {
		total += actNumber
	}
	return total / float64(len(nums))
}

func namedReturn(a int64) (b int64) {
	b = a + 1
	return
}

func returnMultipleValues(a int64) (b int64, ok bool) {
	b = a + 1
	ok = true
	return
}

func variadicAdd(nums ...int64) int64 {
	var total int64
	for _, actNumber := range nums {
		total += actNumber
	}
	return total
}

func funcInFunc() int {
	x := 0
	increment := func() {
		x++
	}

	increment()
	increment()
	return x
}

func funcReturnFunc() func() int {
	i := int(1)
	return func() int {
		j := i
		i = i * 2
		return j
	}
}

func makePanic() {
	panic("panic")
}
