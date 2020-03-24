package main

import "fmt"

type product struct {
	name   string
	number int
	price  float64
}

func main() {
	product := product{name: "Noodle", number: 30, price: 1.2}
	fmt.Println("Product Name:", product.name)
	fmt.Println("Number of product:", product.number)
	fmt.Println("Price:", product.price)
}
