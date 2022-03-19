//btree

package main

import "fmt"

type Tree struct {
	node *Node
}

func (t *Tree) Insert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.Insert(value)
	}
	return t
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.Insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.Insert(value)
		}
	}
}

func (n *Node) exist(value int) bool {
	if n == nil {
		return false
	}
	if n.value == value {
		return true
	}
	fmt.Println("value: ", n.value)
	if value <= n.value {
		return n.left.exist(value)
	} else {
		return n.right.exist(value)
	}
}

func printNode(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.value)
	fmt.Println(n.left)
	fmt.Println(n.right)
}

func main() {
	t := &Tree{}
	t.Insert(10).
		Insert(8).
		Insert(20).
		Insert(9).
		Insert(0).
		Insert(15).
		Insert(25)
	fmt.Println(t.node.exist(26))
}
