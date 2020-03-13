package main

import "fmt"

func cal(a int, b int) (int, int) {
	circum := (a + b) * 2
	area := a * b
	return circum, area
}

func main() {
	fmt.Println("Calculate Circumference and Area of a rectangle 10 x 20")
	var a int = 10
	var b int = 20
	circum, area := cal(a, b)
	fmt.Println("Circumference:", circum)
	fmt.Println("Area:", area)
}
