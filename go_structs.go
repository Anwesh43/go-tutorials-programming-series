package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Engineer struct {
	*Person
	deg string
}

func (p *Person) print() {
	fmt.Println("name is", p.name, "age is", p.age)
}

func main() {
	engineer := Engineer{&Person{name: "Anwesh", age: 25}, "ece"}
	engineer.print()
}
