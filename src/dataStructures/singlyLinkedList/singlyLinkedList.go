package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) printElements() {

	if l.len == 0 {
		fmt.Println("Empty Linked List")
	}

	temp := l.head

	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}

}

func (l *LinkedList) printElementsReverseRecursive() {
	//TODO: Finish printElementReverseRecursiveFunction
	if l.len == 0 {
		fmt.Println("This Linked List is Empty")
	}

	temp := l.head

	if temp.next == nil {
		return
	}

}

func (l *LinkedList) insertHeadElement(data int) {

	if l.len == 0 {
		myHeadNode := Node{
			data: data,
			next: nil,
		}

		l.head = &myHeadNode
		l.len++
		return
	}
	temp := l.head

	myHeadNode := Node{
		data: data,
		next: temp,
	}

	l.head = &myHeadNode
	l.len++

}

func (l *LinkedList) insertLastElement(data int) {

	myNode := Node{
		data: data,
		next: nil,
	}

	temp := l.head

	for temp.next != nil {
		temp = temp.next
	}

	temp.next = &myNode
	l.len++

}

func (l *LinkedList) insertAtGivenPosition(positionData int, data int) {

	temp := l.head

	if l.head.data == positionData {
		myNode := Node{
			data: data,
			next: l.head.next,
		}

		l.head.next = &myNode
		l.len++

		return
	}

	for temp != nil {
		if temp.data == positionData {
			myNode := Node{
				data: data,
				next: temp.next,
			}
			temp.next = &myNode
			l.len++

			return
		}
		temp = temp.next
	}

	fmt.Println("This Linked List has not this data : ", positionData)

}

func (l *LinkedList) length() int {
	return l.len
}

func (l *LinkedList) isAvailable(data int) bool {
	temp := l.head
	for temp != nil {
		if data == temp.data {
			return true
		}
		temp = temp.next
	}
	return false
}

func (l *LinkedList) deleteElement(data int) {

	if l.len == 0 {
		fmt.Println("This Linked List Empty")
	} else {
		temp := l.head

		if l.head.data == data {
			l.head = l.head.next
			l.len--
			return
		}

		for temp.next.data != data {
			if temp.next.next == nil && temp.next.data != data {
				fmt.Println("This Linked List has not this data:", data)
				return
			}
			temp = temp.next
		}

		temp2 := temp.next
		temp.next = temp2.next
		temp2 = nil
		l.len--
	}

}

func main() {

	//Tested Some Cases in Main Function

	//CreateLinkedList
	myLinkedList := LinkedList{head: nil, len: 0}
	mySecondLinkedList := LinkedList{head: nil, len: 0}

	//AddHead
	myLinkedList.insertHeadElement(10)
	myLinkedList.insertHeadElement(20)

	//AddLast
	myLinkedList.insertLastElement(30)

	//AddGivenPosition
	//myLinkedList.insertAtGivenPosition(240, 15)

	mySecondLinkedList.insertHeadElement(10)
	mySecondLinkedList.insertLastElement(30)

	mySecondLinkedList.insertAtGivenPosition(10, 20)

	mySecondLinkedList.printElements()

	mySecondLinkedList.deleteElement(330)
	mySecondLinkedList.printElements()

}
