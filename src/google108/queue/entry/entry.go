package main

import (
	"fmt"
	"google108/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	fmt.Println(q)
	fmt.Println(q.IsEmpty())
	q.Pop()
	q.Pop()
	fmt.Println(q.IsEmpty())
	q.Pop()
}
