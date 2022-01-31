package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type queue struct {
	count int
	front *Node
	rear  *Node
}

func createNode(data int) *Node {
	myNode := Node{
		data: data,
		next: nil,
	}

	return &myNode
}

func (q *queue) isEmpty() bool {
	return q.count == 0
}

func (q *queue) isFull() bool {
	return q.count == QueueSize
}

//Todo : Update This
func (q *queue) printElements() {
	if q.isEmpty() {
		fmt.Println("This Queue is Empty")
		return
	}

	for q.front != nil {
		fmt.Println(q.front.data)
	}

}

func (q *queue) len() int {
	return q.count
}

func (q *queue) enqueue(myNode *Node) {
	if q.isFull() {
		fmt.Println("Queue is full")
		return
	}

	if q.isEmpty() {
		q.rear = myNode
		q.front = myNode
	} else {
		q.rear.next = myNode
		q.rear = myNode
	}
	q.count++

}

func main() {
	myQueue := queue{
		count: 0,
		rear:  nil,
		front: nil,
	}

	myQueue.enqueue(createNode(10))
}
