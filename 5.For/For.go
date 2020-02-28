package main

import "fmt"

func main() {
	fmt.Println("First for")
	var i int = 0
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("Second for")
	for j := 3; j >= 0; j-- {
		fmt.Println(j)
	}
	fmt.Println("End")
	fmt.Println("Let do multiple for loop:")
	for a := 1; a <= 10; a++ {
		for b := 1; b <= 3; b++ {
			fmt.Println("Multiply", a, " x ", b, "=", a*b)
		}
	}
}
