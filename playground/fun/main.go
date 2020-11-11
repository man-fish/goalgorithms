package main

import "fmt"

type man struct {
	name string
}

func (m *man) say() {
	(*m) = man{}
	fmt.Println(m.name)
}

func mansay(m *man) {
	if m == nil {
		fmt.Println("nil man")
		return
	}
	m.say()
}

func main() {
	m := &man{
		name: "yoxn",
	}
	m.say()
	fmt.Println(m.name)
}
