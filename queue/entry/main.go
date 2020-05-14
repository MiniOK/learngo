package main

import (
	"fmt"
	"hello.go/queue"
)

func main() {
	q := queue.Queue{}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	q.Push("abc")
	fmt.Println(q.Pop())

}
