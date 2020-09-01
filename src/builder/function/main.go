package main

import "fmt"

type person struct {
	name, position string
}

type acton func(*person)
type personBuilder struct {
	actions []acton
}

func (p *personBuilder) called(name string) *personBuilder {
	p.actions = append(p.actions, func(p *person) {
		p.name = name
	})
	return p
}
func (p *personBuilder) worksAsA(position string) *personBuilder {
	p.actions = append(p.actions, func(p *person) {
		p.position = position
	})
	return p
}
func (p *personBuilder) build() *person {
	per := &person{}
	for _, act := range p.actions {
		act(per)
	}
	return per
}
func main() {
	pbuilder := personBuilder{}
	per := pbuilder.called("aditya").worksAsA("developer").build()
	fmt.Println(per)
}
