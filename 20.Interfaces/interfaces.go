package main

import (
	"fmt"
	"math"
)

type geo interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perim() float64 {
	return (r.width + r.height) * 2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r1 := rectangle{width: 10, height: 2}
	fmt.Printf("Area1: %.2f - Perim1: %.2f \n", r1.area(), r1.perim())
	r2 := rectangle{width: 14, height: 3.3}
	fmt.Printf("Area2: %.2f - Perim2: %.2f \n", r2.area(), r2.perim())
	c := circle{radius: 10}
	fmt.Printf("Circle Area: %.2f - Circle Perim: %.2f \n", c.area(), c.perim())
}
