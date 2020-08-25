package main

import "fmt"

// Liskov deals with interheritance
// it states that if you have a api which takes base class , it should work correctly with derived class as well

type sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type rectangle struct {
	width, height int
}

func (r *rectangle) GetWidth() int {
	return r.width
}

func (r *rectangle) SetWidth(width int) {
	r.width = width
}
func (r *rectangle) GetHeight() int {
	return r.height
}

func (r *rectangle) SetHeight(height int) {
	r.height = height
}

func calculateArea(s sized) {
	area := s.GetHeight() * s.GetWidth()
	fmt.Println("area is ", area)
}

type Square struct {
	rectangle
}

func NewSquare(size int) *Square {
	squr := &Square{}
	//breaking list of substitution principle ( rectangle is different from square , square has same witdh height)
	squr.height, squr.width = size, size
	return squr
}

func main() {
	rect := rectangle{}
	rect.SetHeight(10)
	rect.SetWidth(12)
	calculateArea(&rect)
	squr := NewSquare(10)
	calculateArea(squr)
}
