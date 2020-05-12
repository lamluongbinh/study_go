package main

import "fmt"

type order struct {
	name string
	price float64
	number int
}

func (in order) info() string {
	info := "Name: " + in.name + " -  Price: " + fmt.Sprintf("%g", in.price) + " - Piece: " + fmt.Sprintf("%d", in.number)  
	return info
}

func (c order) cal() float64 {
	return c.price * float64(c.number)
}

func main() {
	var i = order{name: "Iphone 11", price: 395, number: 12 }
	fmt.Println("===============================")
	fmt.Println(i.info())
	fmt.Println("Cost: ", i.cal())
	
	fmt.Println("===============================")
	s := order{name: "Samsung Galaxy S20", price: 999, number: 2 }
	fmt.Println(s.info())
	fmt.Println("Cost: ", s.cal())
}