package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func (p *person) speak() {
	fmt.Println("Hello,")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: `James`,
		last:  `Bond`,
	}

	//works as `&p1` is a pointer to type person
	saySomething(&p1)

	//does not work as `p1` not a pointer to type person
	//saySomething(p1)

	p1.speak()

	fmt.Println(p1.first, p1.last)
}
