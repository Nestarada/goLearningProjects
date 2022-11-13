package main

import (
	"fmt"
	"reflect"
)

type Queue struct {
	queue []interface{}
}

func (q *Queue) Add(input interface{}) {
	if len(q.queue) == 0 {
		q.queue = append(q.queue, input)
	} else {
		if reflect.TypeOf(q.queue[0]) == reflect.TypeOf(input) {
			q.queue = append(q.queue, input)
		} else {
			panic("cannot add value to queue (mismatched value type on queue and input)")
		}
	}
}

func (q *Queue) Top() interface{} {
	return q.queue[0]
}

func (q *Queue) Remove() {
	q.queue = q.queue[1:]
}

func main() {
	q := new(Queue)
	q2 := new(Queue)
	q.Add(1)
	q.Add(2)
	//q.Add(3.1)
	q2.Add(true)
	fmt.Println(q.queue, q2.queue)
	topQ := q.Top()
	topQ2 := q2.Top()
	fmt.Println(topQ, topQ2)
	q.Remove()
	fmt.Println(q.queue, q2.queue)
}
