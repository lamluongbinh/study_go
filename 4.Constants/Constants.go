package main

import "fmt"

const s string = "Truth"

func main() {
	fmt.Println(s, "is never changed")
	fmt.Println("If we test to change", s, "to a Lie")
	fmt.Println("We will get error, @@, just test by add s = \"lie\" to the code")
	const a = 10
	fmt.Println("Remember that number is also can be a const, just like the code above")
	var b int = 2
	fmt.Println("Let multiply const and var: 10 x 2 = ", a*b)
}
