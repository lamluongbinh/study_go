package main

import "fmt"

func sum(nums ...int) int {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func main() {
	a := sum(1, 2)
	fmt.Println("1 + 2 = ", a)
	b := sum(1, 2, 3)
	fmt.Println("1 + 2 + 3= ", b)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Sum multiple int variabl:", sum(nums...))

}
