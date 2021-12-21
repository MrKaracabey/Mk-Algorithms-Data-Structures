package main

import "fmt"

const QueueSize = 5

type Queue struct {
	count int
	front int
	rear  int
	data  [QueueSize]int
}

func (q *Queue) isEmpty() bool {
	return q.count == 0
}

func (q *Queue) isFull() bool {
	return q.count == QueueSize
}

func (q *Queue) printElements() {

	if q.isEmpty() {
		fmt.Println("This Queue is Empty")
	}
	for _, item := range q.data {
		if item != 0 {
			fmt.Println(item)
		}
	}

}

func (q *Queue) enqueue(data int) {
	if q.isFull() {
		fmt.Println("This Queue is Full")
	}

	q.rear++
	q.count++

	if q.rear == QueueSize-1 {
		q.rear = 0
	}

	q.data[q.rear] = data
}

func (q *Queue) dequeue() {

	if q.isEmpty() {
		fmt.Println("This Queue is empty")
	}

	q.data[q.front] = 0
	q.front++
	q.count--

	if q.front == QueueSize-1 {
		q.front = 0
	}

}

func (q *Queue) len() int {
	return q.count
}

func main() {
	//0 Value Means Free in Memory according to this implementation :)
	myQueue := Queue{
		count: 0,
		front: 0,
		rear:  -1,
	}

	myQueue.enqueue(10)
	myQueue.enqueue(20)
	myQueue.enqueue(30)

	myQueue.dequeue()

	myQueue.printElements()

}
