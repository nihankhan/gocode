package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

func (c circle) area() float64 {
	return (2 * math.Pi * (c.radius * c.radius))
}

func (s square) area() float64 {
	return (s.side * s.side)
}

func main() {
	c := circle{10}
	s := square{5}

	fmt.Println(c)
	fmt.Println(s)

	fmt.Println(c.area())
	fmt.Println(s.area())
}
