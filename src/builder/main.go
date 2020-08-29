package main

import "fmt"

type linklist struct {
	in   int
	next *linklist
}

var root *linklist

func (l *linklist) addElements(num int) *linklist {
	n := &linklist{}
	n.in = num
	if root != nil {
		l.next = n
	} else {
		root = n
		l = n
	}
	return n
}

func (l *linklist) display() {

	fmt.Print("printing current linklist:===>\n[ ")
	for l != nil {
		fmt.Printf("%d ", l.in)
		l = l.next
	}
	fmt.Printf("]\n")
}
func (l *linklist) search(i int) int {
	for l != nil {
		if l.in == i {
			fmt.Printf("element %d found\n", l.in)
			return 0
		}
		l = l.next
	}
	fmt.Println("element not found")
	return -1
}

func (l *linklist) deleteElementAt(position int) {

	for i := 0; i <= position-2 && l.next != nil; i++ {

		l = l.next
	}
	if l == nil || l.next == nil || position < 0 {
		fmt.Println("element not found at positon", position)
		return
	}
	if position == 0 && root != nil {
		root = l.next
		return
	}
	temp := l.next
	l.next = temp.next

}

func (l *linklist) addElementAt(number int, position int) {
	i := 0
	for ; position > 0 && i != position-1 && l.next != nil; i++ {
		l = l.next
	}
	if i != position {
		fmt.Printf("\n Its not possible to add element at positon %d, so addiing it at position %d \n", position, i)
	}
	newlinkList := &linklist{}
	newlinkList.in = number
	newlinkList.next = l.next
	l.next = newlinkList
}
func main() {
	root.deleteElementAt(0)
	var l *linklist
	l = l.addElements(1).addElements(2).addElements(3).addElements(23).addElements(123324).addElements(6543)
	root.display()
	l.display()
	root.search(10)
	root.search(23)
	root.search(23)
	l.addElements(345678)
	root.addElementAt(33, 3)
	root.addElementAt(34, 3)
	root.addElementAt(333, 33)
	root.display()
	root.deleteElementAt(0)
	root.deleteElementAt(3)
	root.deleteElementAt(3333)
	root.deleteElementAt(-3333)
	root.display()

}
