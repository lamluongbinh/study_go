package main

import "fmt"

func main() {
	var i int = 10
	b := &i
	fmt.Println("b: ", b, "and Address of i: ", &i, "And currently value of i:", i)
	fmt.Println("If we change the value of i to 20")
	i = 20
	fmt.Println("As the result, if we get the value of i from b:", *b)
}
