package main

import "fmt"

// HL module should not depend on LL modules
//Both should be depend on abstraction
type relationship int

const (
	parent relationship = iota
	child
)

type person struct {
	name string
}

type info struct {
	from     *person
	relation relationship
	to       *person
}

type relationships struct {
	relations []info
}

type relationbrowse interface {
	findChilds(string) []*person
}

func (rln *relationships) findChilds(name string) []*person {
	result := make([]*person, 0)
	for _, pers := range rln.relations {
		if pers.relation == parent && name == pers.from.name {
			result = append(result, pers.to)
		}
	}
	return result
}

func (rtl *relationships) addParentChild(parents *person, ch *person) {
	rtl.relations = append(rtl.relations, info{parents, parent, ch})
	rtl.relations = append(rtl.relations, info{ch, child, parents})
}

type reserch struct {
	br relationbrowse // low-level
}

func (r *reserch) Investigate(name string) {
	for _, v := range r.br.findChilds(name) {
		fmt.Println("name", v.name)
	}
}

func main() {
	jonh := person{"john"}
	rom := person{"rom"}
	rom1 := person{"rom1"}

	// low-level module
	rs := &relationships{}
	rs.addParentChild(&jonh, &rom)
	rs.addParentChild(&jonh, &rom1)
	r := &reserch{rs}
	r.Investigate("john")
}
