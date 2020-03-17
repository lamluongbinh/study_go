package main

import "fmt"

func recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursion(n-1)
}

func main() {
	var a int = 10
	result := recursion(a)
	fmt.Println("Recursion of n = 10, n x (n-1) x .... x 1 = ", result)

}
