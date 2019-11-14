package main

import (
	"fmt"
)

func main() {
	a := 9
	if a%2 == 0 {
		fmt.Println(a, "is even")
	} else {
		fmt.Println(a, "is odd")
	}

	if a < 0 {
		fmt.Println(a, "is negative")
	} else if a < 10 {
		fmt.Println(a, "has 1 digit")
	} else {
		fmt.Println(a, "has multi digit")
	}
}
