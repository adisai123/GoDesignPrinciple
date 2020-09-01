package main

import "fmt"

type person struct {
	age  int
	name string
}
type Person interface {
	sayHello()
}

type oldPerson struct {
	age  int
	name string
}

func (p *person) sayHello() {
	fmt.Printf("\nname = %s, age := %d\n", p.name, p.age)
}
func (p *oldPerson) sayHello() {
	fmt.Printf("old person name = %s, age := %d\n", p.name, p.age)
}

func NewPerson(name string, age int) Person {
	if age > 50 {
		return &oldPerson{age, name}
	} else {
		return &person{age, name}
	}
}
func main() {
	p := NewPerson("adita", 222)
	p.sayHello()
	p2 := NewPerson("si", 12)
	p2.sayHello()
}
