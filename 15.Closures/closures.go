package main

import "fmt"

// This function return another function that define anonymously in its
func intSeq() func() int {
	i := 1
	return func() int {
		i = i * 2
		return i
	}
}
func main() {
	fmt.Println("The variable i is close over the variable, we can use it again if we call the func")
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println("The function is unique to that particular function")
	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())

	fmt.Println("Call nextInt again")
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
