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
	hl := node.left.height()
	hr := node.right.height()
	return int(math.Max(float64(hl), float64(hr)))
}

func (node *BSTNode) preOrder() {
	if node != nil {
		fmt.Print(node.val, ", ")
		node.left.preOrder()
		node.right.preOrder()
	}
}

func main() {
	root := &BSTNode{val:50}
    root.add(30)
    root.add(45)
    root.add(32)
    root.add(10)
    root.add(90)
    root.add(55)
    root.add(49)
    root.add(-5)
    root.add(88)
    root.add(110)
    root.add(40)
	fmt.Println(root.height())
}
