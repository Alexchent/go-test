package sunfa

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func main() {
	//fmt.Println(tree.New(1))
	ch := make(chan int)
	defer close(ch)
	go Walk(tree.New(1), ch)

	for {
		fmt.Println(<-ch)
	}
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	//close(ch)
}
