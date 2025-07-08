type BSTNode struct {
	left *BSTNode
	val int
	right *BSTNode
	height int
	bf int
}
func createNode(val int) *BSTNode {
	return &BSTNode{val: val}
}
func (node *BSTNode) add(val int) *BSTNode {
	if node != nil {
		if val <= node.val {
			if node.left == nil {
				node.left = createNode(val)
			} else {
				node.left = node.left.add(val)
			}
		} else {
			if node.right == nil {
				node.right = createNode(val)
			} else {
				node.left = node.right.add(val)
			}
		}
	}
	node.updateProperties()
	return node.rebalance()
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

func (node *BSTNode) remove(val int) *BSTNode {
	if node == nil {return nil}
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

func (node *BSTNode) updateProperties() {
}

func (node *BSTNode) rotRight() *BSTNode {
}

func (node *BSTNode) rotLeft() *BSTNode {
}

func (node *BSTNode) rebLeftLeft() *BSTNode {}
func (node *BSTNode) rebLeftRight() *BSTNode {}
func (node *BSTNode) rebRightRight() *BSTNode {}
func (node *BSTNode) rebRightLeft() *BSTNode {}

func (node *BSTNode) rebalance() *BSTNode {}
