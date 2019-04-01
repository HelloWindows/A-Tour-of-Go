package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walkTree(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Walk is func
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	} // end if
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same is func
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go walkTree(t1, ch1)
	go walkTree(t2, ch2)
	for v1 := range ch1 {
		if v1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go walkTree(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("Should return true:", Same(tree.New(1), tree.New(1)))
	fmt.Println("Should return false:", Same(tree.New(1), tree.New(2)))
}
