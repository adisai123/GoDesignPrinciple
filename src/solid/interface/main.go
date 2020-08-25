package main

import "fmt"

//interface segregation
//old fasion and latest (scan , print, fax seperate seprate interface instead of only one printer interface)

type basicprinter interface {
	print()
}

type basicFax interface {
	fax()
}
type basicScanner interface {
	scan()
}

type multifunctionPrinter interface {
	basicprinter
	basicScanner
	basicFax
}

type oldPrinter struct {
	msg string
}

func (oldp *oldPrinter) print() {
	fmt.Println("msg to print :=>", oldp.msg)
}

type mordenprinter struct {
	doc string
}

func (mp *mordenprinter) print() {
	fmt.Println("print document :=", mp.doc)
}

func (mp *mordenprinter) scan() {
	fmt.Println("scanning document :=", mp.doc)
}

func (mp *mordenprinter) fax() {
	fmt.Println("srnding fax of this  document :=", mp.doc)
}

func usePrinter(p multifunctionPrinter) {
	p.scan()
	p.print()
	p.fax()
}

func main() {
	oldP := oldPrinter{"printing on old printer"}
	oldP.print()
	modernP := mordenprinter{"aadhar"}
	usePrinter(&modernP)
}
