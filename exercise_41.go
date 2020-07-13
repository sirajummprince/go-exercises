package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	sq := square{
		side: 5,
	}
	ci := circle{
		radius: 3,
	}
	aos := info(sq)
	aoc := info(ci)

	fmt.Println("Area of square:", aos)
	fmt.Println("Area of circle:", aoc)
}
