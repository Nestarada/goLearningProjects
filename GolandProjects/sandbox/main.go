package main

import (
	"queue"
)

func main() {
	q := new(queue.Queue)
	q.Add(1)
	q.Print()
	q.Remove()
	q.Print()
}
