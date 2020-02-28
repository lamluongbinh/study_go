package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a [3]int
	fmt.Println("Empty array", a)
	a[2] = 10
	fmt.Println("Set Array: ", a)
	fmt.Println("Get Array position 2 value: ", a[2])
	fmt.Println("Get length of array:", len(a))

	fmt.Println("=========== 2 Dimensional Array ===========")
	var b [2][2]string
	b[0][0] = "One"
	b[0][1] = "Two"
	b[1][0] = "Three"
	b[1][1] = "Four"
	fmt.Println("Print string array b:", b)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(b[i][j])
		}
	}
	fmt.Println("=========== Set by for ===========")
	var tmp int = 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			b[i][j] = "Number " + strconv.Itoa(tmp)
			tmp++
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(b[i][j])
		}
	}
}
