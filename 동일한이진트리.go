package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int)

func Same(t1, t2 *tree.Tree) bool

func main() {
	c := make(chan int)

	go Walk(tree.Tree(1), c)
	go Same(tree.Tree(1), c)
	fmt.Println("test")
}
