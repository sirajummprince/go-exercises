package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("First:", p.first, "Last:", p.last, "Age:", p.age)
}

func main() {
	p1 := person{
		first: `James`,
		last:  `Bond`,
		age:   32,
	}
	p1.speak()
}
