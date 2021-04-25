package main

import "fmt"

func main() {
	var a = "This is a string"
	fmt.Println(a)
	var a1 string = "This is also a string"
	fmt.Printf("This variable type is %T\n", a1)
	var b, c int = 1, 2
	fmt.Println("b = ", b, " and c = ", c)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)
	f := "This is an apple"
	fmt.Println(f)
	g := 3
	fmt.Println(g)
}
