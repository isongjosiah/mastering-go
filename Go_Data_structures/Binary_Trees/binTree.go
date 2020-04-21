package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Tree structure for the implementation of a binary tree
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

func create(n int) *Tree {
	var t *Tree
	var random []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		random = append(random, temp)
		t = insert(t, temp)
	}
	fmt.Println(random)
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func main() {
	fmt.Println("this is create")
	tree := create(10)
	fmt.Println(tree)
	fmt.Println("Left", tree.Left)
	fmt.Println("Left", tree.Left.Value)
	fmt.Println("Right", tree.Right)
	fmt.Println("Right", tree.Right.Value)
	fmt.Println("The value of the root of the tree is", tree.Value)
	fmt.Println("This is traverse")
	traverse(tree)
	fmt.Println()
	fmt.Println("This is insert")
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	fmt.Println("Traverse again")
	traverse(tree)
	fmt.Println()
	fmt.Println("Thevalue of the root of the tree is", tree.Value)

}
