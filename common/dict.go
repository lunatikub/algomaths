package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// The first aim of this module is to build a tree
// based on a dictionnary, a file with 1 word by line,
// to have a dichotomic search which returns if a given word exist.

type node struct {
	childs []*node
	id     int
	r      rune
	leaf   bool
}

// Dict dictionnary
type Dict struct {
	root   node
	nextID int
}

func (n *node) addChild(id int, r rune) *node {
	new := node{[]*node{}, id, r, false}
	n.childs = append(n.childs, &new)
	return &new
}

func (n *node) findChild(r rune) *node {
	for _, iter := range n.childs {
		if iter.r == r {
			return iter
		}
	}
	return nil
}

func (D *Dict) addWord(word string) {
	currNode := &D.root
	var nextNode *node

	for _, r := range word {
		if nextNode = currNode.findChild(r); nextNode == nil {
			D.nextID++
			nextNode = currNode.addChild(D.nextID, r)
		}
		currNode = nextNode
	}

	currNode.leaf = true
}

// NewDict Build a new dictionnary
// from an input file (1 word by line)
func NewDict(inputFile string) *Dict {
	D := Dict{}

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		D.addWord(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &D
}

func (n *node) toDot(depth int) {
	label := string(n.r)
	color := "\"#93ffae\""
	if n.r == 0 {
		label = "root"
		color = "\"#ffae00\""
	} else if n.leaf {
		color = "skyblue"
	}
	fmt.Printf("%v [label=%v, style=filled, color=%v]\n",
		n.id, label, color)
	for _, iter := range n.childs {
		fmt.Printf("%v -> %v\n", n.id, iter.id)
		iter.toDot(depth + 1)
	}
}

// ToDot dictionnary to dot file
func (D *Dict) ToDot() {
	fmt.Printf("digraph {\n")
	D.root.toDot(0)
	fmt.Printf("}\n")
}
