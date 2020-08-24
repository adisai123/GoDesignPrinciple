package main

import "fmt"

//open for extension and close for modification
type produce struct {
	name  string
	color color
	size  size
}

func (p *produce) String() string {
	return fmt.Sprintf("name : %s ;color : %s;size : %s", p.name, p.color, p.size)
}

type size int
type color int

const (
	red color = iota
	yello
	green
)
const (
	large size = iota
	small
	medium
)

type specification interface {
	isSatisfy(p *produce) bool
}
type colorSpec struct {
	color color
}

func (s *colorSpec) isSatisfy(p *produce) bool {
	return s.color == p.color
}

type sizeSpec struct {
	size size
}

func (s *sizeSpec) isSatisfy(p *produce) bool {
	return s.size == p.size
}

type basicfilter struct {
}
type andspecification struct {
	first, second specification
}

func (a *andspecification) isSatisfy(p *produce) bool {
	return a.first.isSatisfy(p) && a.second.isSatisfy(p)
}

func (f *basicfilter) filter(p []produce, s specification) []*produce {
	result := make([]*produce, 0)

	for _, v := range p {
		if s.isSatisfy(&v) {
			result = append(result, &v)
		}
	}
	return result
}

func main() {
	item1 := produce{"item1", red, large}
	item2 := produce{"item2", green, medium}
	item3 := produce{"item3", yello, small}
	proc := []produce{item1, item2, item3}
	sspec := sizeSpec{size: large}
	colrsp := colorSpec{color: green}
	bc := basicfilter{}
	for _, val := range proc {
		fmt.Println(val.String)
	}
	for _, val := range bc.filter(proc, &colrsp) {
		fmt.Println(val.String())
	}
	for _, val := range bc.filter(proc, &sspec) {
		fmt.Println(val.String())
	}
	ac := andspecification{&colrsp, &sspec}
	for _, val := range bc.filter(proc, &ac) {
		fmt.Println(val.String())
	}
}
