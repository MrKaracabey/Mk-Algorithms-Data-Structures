package main

import "fmt"

const StackSize = 5

type Stack struct {
	top  int
	data [StackSize]int
}

func (s *Stack) printElements() {
	if s.top == -1 {
		fmt.Println("This Stack is empty")
		return
	}

	index := StackSize - 1

	for i := 0; i < len(s.data); i++ {
		fmt.Println(s.data[index])
		index--
	}
}

func (s *Stack) isFull() bool {
	return s.top == StackSize-1 && s.data[s.top] != 0
}

func (s *Stack) isEmpty() bool {
	return s.top == 0
}

func (s *Stack) push(data int) {

	if s.isFull() {
		fmt.Println("This stack is Already Full")
		return
	}

	if s.top < StackSize-1 {
		s.top++
		s.data[s.top] = data
	}

}

func (s *Stack) pop() {
	if s.isEmpty() {
		fmt.Println("This stack is empty")
	}

	s.data[s.top] = 0
	s.top--

}

func (s *Stack) peek() int {
	return s.data[s.top]
}

func main() {
	//0 Value Means Free in Memory according to this implementation :)
	myStack := Stack{top: -1}

	myStack.push(5)
	myStack.push(10)
	myStack.push(15)
	myStack.push(20)
	myStack.push(25)

	myStack.pop()
	myStack.pop()
	myStack.pop()
	myStack.pop()

	myStack.push(20)
	myStack.push(25)

	fmt.Println(myStack.peek())

}
