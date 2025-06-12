// func levelOrderNav()
// func Remove(val int)
// func size() int
// func height() int
// func preOrderNav()
// func inOrderNav()
// func postOrderNav()

// pseudocódigo para navegação
// a fila é criada com o nó raiz
//enquanto fila != vazia
//	frente = fila.dequeue()
//	print(frent.val)
//	se frente.esq != nil {fila.enqueue(frente.esq)}
//	se frente.dir != nil {fila.enqueue(frente.dir)}


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

func (node *BSTNode) height() int {
	hl := node.left.height()
	hr := node.right.height()
	return Math.Max(hl, hr)
}
 
func (node *BSTNode) preOrder() {
	if node != nil {
		fmt.Print(node.val, ", ")
		node.left.preOrder()
		node.right.preOrder()
	}
}

func (node *BTSNode) remove(val int) *BTSNode {
	if node == nil {return nil}
	if val < node.val { node.left = node.left.remove(val)}
	else if val > nodel.val { node.right = node.right.remove(val)}
	else { 
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
		return node
	}
}

func main() {
	var bt BTSNode
	bt.add(1)
	fmt.Println(bt.search(1))
}
