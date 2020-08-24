package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Journal struct {
	entities []string
}

var entryCount int

func (j *Journal) String() string {
	return strings.Join(j.entities, "\n")
}

func (j *Journal) Add(text string) int {
	entryCount++
	str := fmt.Sprintf("%d, %s", entryCount, text)
	j.entities = append(j.entities, str)
	return entryCount
}

func (j *Journal) Remove(index int) {

}

type persistent struct {
	filename string
}

func (p *persistent) load() {

}
func (p *persistent) save(data string) {
	ioutil.WriteFile(p.filename, []byte(data), 0777)
}
func main() {
	j := Journal{}
	j.Add("aditya2q3")
	j.Add("nupur")
	p := persistent{"aditya.txt"}
	p.save(j.String())
}
