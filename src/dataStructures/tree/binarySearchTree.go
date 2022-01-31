package main

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}

type binarySearchTree struct {
	root *node
}

func createNode(data int) *node {
	myNode := node{
		data:  data,
		left:  nil,
		right: nil,
	}

	return &myNode
}

func showPreOrderElements(root *node) {
	if root != nil {
		fmt.Println(root.data)
		showPreOrderElements(root.left)
		showPreOrderElements(root.right)

	}

}

func showInorderElements(root *node) {
	if root != nil {
		showInorderElements(root.left)
		fmt.Println(root.data)
		showInorderElements(root.right)

	}
}

func showPostOrderElements(root *node) {
	if root != nil {
		showPostOrderElements(root.left)
		showPostOrderElements(root.right)
		fmt.Println(root.data)
	}
}

func insertElement(root *node, data int) *node {
	if root != nil {
		if data < root.data {
			root.left = insertElement(root.left, data)
		} else {
			root.right = insertElement(root.right, data)
		}

	} else {
		root = createNode(data)
	}

	return root
}

func length(root *node) int {
	if root == nil {
		return 0
	}

	return length(root.left) + length(root.right) + 1
}

//dont count rootNode (Check This Function)
func height(root *node) int {
	if root != nil {
		lHeight := height(root.left)
		rHeight := height(root.right)

		if lHeight > rHeight {
			return lHeight + 1
		} else {
			return rHeight + 1
		}
	} else {
		return -1
	}
}

func findElement(root *node, key int) int {
	if root == nil {
		return 0
	}

	if root.data == key {
		return 1
	}

	isLeft := findElement(root.left, key)
	isRight := findElement(root.right, key)

	if isLeft > isRight {
		return isLeft
	}

	return isRight
}

func minElement(root *node) int {

	if root == nil {
		return 0
	}

	if root.left == nil {
		return root.data
	}

	return minElement(root.left)
}

func maxElement(root *node) int {

	if root == nil {
		return 0
	}

	if root.right == nil {
		return root.data
	}

	return maxElement(root.right)
}

//Todo : Çöz
func isBinarySearchTree(root *node) int {
	// 1- parent may have max 2 child
	// 2- left child should be smaller than parent node
	// 3- right child should be bigger than parent node

	return 1
}

func sumTreeNodes(root *node) int {

	if root == nil {
		return 0
	}

	return sumTreeNodes(root.left) + sumTreeNodes(root.right) + root.data
}

// Todo: İncele
func showLeftChild(root *node) {

	if root != nil {
		if root.left != nil && root.right != nil {
			fmt.Println(root.data)
		}

		showLeftChild(root.left)
		showLeftChild(root.right)
	} else {
		return
	}

}

func main() {
	myNode := createNode(10)

	myBinarySearchTree := binarySearchTree{
		root: myNode,
	}

	showPreOrderElements(myBinarySearchTree.root)
	myBinarySearchTree.root = insertElement(myBinarySearchTree.root, 8)
	myBinarySearchTree.root = insertElement(myBinarySearchTree.root, 12)
	myBinarySearchTree.root = insertElement(myBinarySearchTree.root, 13)
	myBinarySearchTree.root = insertElement(myBinarySearchTree.root, 25)
	myBinarySearchTree.root = insertElement(myBinarySearchTree.root, 22)
	myBinarySearchTree.root = insertElement(myBinarySearchTree.root, 4)
	myBinarySearchTree.root = insertElement(myBinarySearchTree.root, 5)
	//Total : 99

	fmt.Println("ShowPreOrder : ")
	showPreOrderElements(myBinarySearchTree.root)

	fmt.Println("*****")
	lengthx := length(myBinarySearchTree.root)
	height := height(myBinarySearchTree.root)
	fmt.Println(lengthx)
	fmt.Println(height)
	fmt.Println("The Length is : ", length(myBinarySearchTree.root))

	isFind := findElement(myBinarySearchTree.root, 12)

	fmt.Println("İsFind : ", isFind)

	maxElement := maxElement(myBinarySearchTree.root)
	minElement := minElement(myBinarySearchTree.root)

	fmt.Println("Max Element is ", maxElement)
	fmt.Println("Min Element is ", minElement)

	isSpecialTree := sumTreeNodes(myBinarySearchTree.root)
	fmt.Println("İs Special : ", isSpecialTree)

	showLeftChild(myBinarySearchTree.root)
}
