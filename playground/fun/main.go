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

type node struct {
	// isKey use to mark whether this node is the end of a word.
	// isKey bool
	// count use to do word-frequency count, it is a substitute of isKey.
	count    int
	children [26]*node
}

func main() {
	m := &man{
		name: "yoxn",
	}
	m.say()
	fmt.Println(m.name)
	n := new(node)
	fmt.Println(n.children[1])
}
