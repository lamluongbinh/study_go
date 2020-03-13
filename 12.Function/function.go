package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func main() {
	var a int = 10
	var b int = 2
	r1 := sum(a, b)
	fmt.Println("10 + 2 = ", r1)
	r2 := multiply(a, b)
	fmt.Println("10 x 2 = ", r2)
}
