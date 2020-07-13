package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	p.first = `Miss`
	p.last = `Moneypenny`
	//or,
	//(*p).first = `Miss`
	//(*p).last = `Moneypenny`
}

func main() {
	p1 := person{
		first: `James`,
		last:  `Bond`,
	}
	fmt.Println("Before pointer:", p1)
	changeMe(&p1)
	fmt.Println("After pointer:", p1)
}
