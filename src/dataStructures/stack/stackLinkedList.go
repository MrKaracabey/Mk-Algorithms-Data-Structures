package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type stack struct {
	top   *Node
	count int
}

func CreateNode(data int) *Node {
	myNode := Node{
		data: data,
		next: nil,
	}

	return &myNode
}

func printElements(node *Node) {

	if node == nil {
		return
	}

	fmt.Println(node.data)

	printElements(node.next)

}

func (s *stack) isFull() bool {
	return StackSize == s.count
}

func (s *stack) isEmpty() bool {
	return s.count == 0
}

func (s *stack) pushStack(myNode *Node) {
	if s.isFull() {
		fmt.Println("This Stack is Full")
		return
	}
	temp := s.top
	s.top = myNode
	s.top.next = temp

	s.count++
}

func (s *stack) popStack() {
	if s.isEmpty() {
		fmt.Println("This stack is Empty")
	}

	s.top = s.top.next
	s.count--

}

func main() {
	myStack := stack{
		count: 0,
		top:   CreateNode(10),
	}

	myStack.pushStack(CreateNode(20))
	myStack.pushStack(CreateNode(30))

	myStack.popStack()
	myStack.popStack()

	printElements(myStack.top)

}
