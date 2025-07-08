package main

import (
	"fmt"
	"math"
)

type BSTNode struct {
	left  *BSTNode
	val   int
	right *BSTNode
}

func createNode(val int) *BSTNode {
	return &BSTNode{val: val}
}

func (node *BSTNode) add(val int) {
	if node != nil {
		if val <= node.val {
			if node.left == nil {
				node.left = createNode(val)
			} else {
				node.left.add(val)
			}
		} else {
			if node.right == nil {
				node.right = createNode(val)
			} else {
				node.right.add(val)
			}
		}
	}
}

func (node *BSTNode) search(val int) bool {
	if node == nil {
		return false
	} else {
		if val < node.val {
			return node.left.search(val)
		} else if val > node.val {
			return node.right.search(val)
		} else {
			return true
		}
	}
}

func (node *BSTNode) min() int {
	if node == nil {
		return -1
	}
	if node.left == nil {
		return node.val
	} else {
		return node.left.min()
	}
}
func (node *BSTNode) max() int {
	if node == nil {
		return -1
	}
	if node.right == nil {
		return node.val
	} else {
		return node.right.max()
	}
}
func (node *BSTNode) height() int {
	if node == nil {
		return -1
	}
	hl := node.left.height()
	hr := node.right.height()
	return 1 + int(math.Max(float64(hl), float64(hr)))
}
func (node *BSTNode) printPre() {
	if node != nil {
		fmt.Print(node.val, ", ")
		node.left.printPre()
		node.right.printPre()
	}
}
func (node *BSTNode) PrintIn() {
	if node == nil {
		return
	}
	node.left.PrintIn()
	fmt.Print(node.val, " , ")
	node.right.PrintIn()
}
func (node *BSTNode) PrintPost() {
	if node == nil {
		return
	}
	node.left.PrintPost()
	node.right.PrintPost()
	fmt.Print(node.val, ", ")
}

func (node *BSTNode) PrintLevels() {
	if node == nil {
		return
	}
	// Supondo uma fila do tipo BSTNode implementada
	queue := Queue{}
	queue.Enqueue(node)

	for !queue.isEmpty() {
		current := queue.Dequeue()

		fmt.Print(current.val, ", ")

		if current.left != nil {
			queue.Enqueue(current.left)
		}
		if current.right != nil {
			queue.Enqueue(current.right)
		}
	}
}

func (node *BSTNode) remove(val int) *BSTNode {
	if node == nil {
		return nil
	}
	if val < node.val {
		node.left = node.left.remove(val)
	} else if val > node.val {
		node.right = node.right.remove(val)
	} else {
		if node.left == nil && node.right == nil {
			return nil
		} else if node.left != nil && node.right == nil {
			return node.left
		} else if node.left == nil && node.right != nil {
			return node.right
		} else {
			max := node.left.max()
			node.val = max
			node.left = node.left.remove(max)
		}
	}
	return node
}

func main() {
	root := BSTNode{val: 5}
	root.add(7)
	root.add(3)
	root.printPre()
}
