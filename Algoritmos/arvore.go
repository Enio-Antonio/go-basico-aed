package main

import "fmt"

type BTSNode struct {
	left *BTSNode
	val int
	right *BTSNode
}

func createNode(val int) *BTSNode {
	return &BTSNode{value: val}
}

func (node *BTSNode) add(val int) {
	if node != nil {
		if val <= node.value {
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

func (node *BTSNode) search(val int) bool {
	if node == nil {
		return false
	} else {
		if val == node.value {
			return true
		} else if val < node.value {
			return node.left.search(val)
		} else if val > node.value {
			return node.right.search(val)
		}
	}
}

func (node *BTSNode) min() int {
	if node == nil { return -1 }
	if node.left == nil {
		return node.val
	} else {
		return node.left.min()
	}
}

func (node *BTSNode) max() int {
	if node == nil { return -1 }
	if node.right == nil {
		return node.val
	} else {
		return node.right.max()
	}
}



func main() {
	var bt BTSNode
	fmt.Println(bt.search(1))
}
