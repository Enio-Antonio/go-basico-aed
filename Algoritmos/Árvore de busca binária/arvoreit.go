// func levelOrderNav()
// func Remove(val int)
// func size() int
// func height() int
// func preOrderNav()
// func inOrderNav()
// func postOrderNav()

package main

import "fmt"

type BSNode struct {
	val int
	dir *BSNode
	esq *BSNode
}

type BinaryTree struct {
	root *BSNode
}

func (bs *BinaryTree) add(e int) {
	newn := BSNode{val: e, dir: nil, esq: nil}
	if bs.root == nil {
		bs.root = &newn

	} else {
		aux := bs.root
		var p *BSNode
		for aux != nil {
			p = aux
			if e > aux.val {
				aux = aux.dir
			} else if e < aux.val {
				aux = aux.esq
			}
		}

		if e > p.val {
			p.dir = &newn
		} else if e < p.val {
			p.esq = &newn
		}
 
	}
}

func (bs *BinaryTree) printMin() {
	aux := bs.root
	for aux != nil {
		fmt.Println(aux.val)
		aux = aux.esq
	}
}

func (bs *BinaryTree) search(e int) bool {
	
}

func main() {
	var bs BinaryTree
	bs.add(5)
	bs.add(4)
	fmt.Println("Acabou")
}
