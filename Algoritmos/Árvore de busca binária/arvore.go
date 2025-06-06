// func levelOrderNav()
// func Remove(val int)
// func size() int
// func height() int
// func preOrderNav()
// func inOrderNav()
// func postOrderNav()

package main

import "fmt"

type BTSNode struct {
	left *BTSNode
	val int
	right *BTSNode
}

func createNode(val int) *BTSNode {
	return &BTSNode{val: val}
}

func (node *BTSNode) add(val int) {
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

func (node *BTSNode) search(val int) bool {
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
	bt.add(1)
	fmt.Println(bt.search(1))
}
