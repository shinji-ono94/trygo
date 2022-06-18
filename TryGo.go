package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) hello() {
	fmt.Printf("%s (%d)\n", p.name, p.age)
}

func main() {
	p := person{
		name: "John",
		age:  30,
	}
	p.hello() // John (30)
}
