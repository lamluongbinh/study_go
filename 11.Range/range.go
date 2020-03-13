package main

import "fmt"

func main() {
	fmt.Println("Range is a way to loop through array or slice")
	fmt.Println("Number Range: {2,3,5}")
	nums := []int{2, 3, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum of range: ", sum)

	strings := map[string]string{"key1": "value1", "key2": "value2"}
	for k, v := range strings {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
