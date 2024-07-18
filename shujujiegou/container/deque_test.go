package container

import (
	"fmt"
	"testing"
)

func TestQue(t *testing.T) {
	que := newArrayDeque(10)
	que.pushLeft(12)
	que.pushLeft(11)
	que.pushRight(13)

	fmt.Println(que.popLeft())
	fmt.Println(que.popLeft())
	fmt.Println(que.popLeft())
	fmt.Println(que.popLeft())
}

func TestQue2(t *testing.T) {
	que := newArrayDeque(10)
	que.pushLeft(12)
	que.pushLeft(11)
	que.pushRight(13)

	fmt.Println(que.popRight())
	fmt.Println(que.popRight())
	fmt.Println(que.popRight())
	fmt.Println(que.popRight())
}
