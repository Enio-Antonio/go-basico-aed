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

func (bs *BinaryTree) min() int {
	aux := bs.root
	var p *BSNode
	for aux != nil {
		p = aux
		//fmt.Println(aux.val)
		aux = aux.esq
	}
	return p.val
}

func (bs *BinaryTree) max() int {
	aux := bs.root
	var p *BSNode
	for aux != nil {
		p = aux
		aux = aux.dir
	}
	return p.val
}

//func (bs *BinaryTree) search(e int) bool {

func main() {
	var bs BinaryTree
	bs.add(5)
	bs.add(4)
	bs.add(1)
	a := bs.min()
	m := bs.max()
	fmt.Println(a)
	fmt.Println(m)
	fmt.Println("Acabou")
}
