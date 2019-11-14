package main

import "fmt"

func main() {
	fmt.Println("First for")
	var i int = 0
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("Next for")
	for j := 3; j >= 0; j-- {
		fmt.Println(j)
	}
	fmt.Println("End")
}
