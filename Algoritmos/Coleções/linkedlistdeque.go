package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedListDeque struct {
	head *Node
	tail *Node
	size int
}

type IDeque interface {
	EnqueueFront(value int)
	EnqueueRear(value int)
	DequeueFront() (int, error)
	DequeueRear() (int, error)
	Front() (int, error)
	Rear() (int, error)
	IsEmpty() bool
	Size() int
}

// EnqueueFront adds an element to the front of the deque.
func (d *LinkedListDeque) EnqueueFront(value int) {
	newNode := &Node{value: value}
	if d.head == nil {
		// If the deque is empty, both head and tail point to the new node.
		d.head = newNode
		d.tail = newNode
	} else {
		// Insert the new node at the front.
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.size++
}

// EnqueueRear adds an element to the rear of the deque.
func (d *LinkedListDeque) EnqueueRear(value int) {
	newNode := &Node{value: value}
	if d.tail == nil {
		// If the deque is empty, both head and tail point to the new node.
		d.head = newNode
		d.tail = newNode
	} else {
		// Insert the new node at the rear.
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
	}
	d.size++
}

// DequeueFront removes and returns the front element of the deque.
func (d *LinkedListDeque) DequeueFront() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque is empty")
	}
	value := d.head.value
	d.head = d.head.next
	if d.head == nil {
		// If the deque becomes empty, reset the tail as well.
		d.tail = nil
	} else {
		d.head.prev = nil
	}
	d.size--
	return value, nil
}

// DequeueRear removes and returns the rear element of the deque.
func (d *LinkedListDeque) DequeueRear() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque is empty")
	}
	value := d.tail.value
	d.tail = d.tail.prev
	if d.tail == nil {
		// If the deque becomes empty, reset the head as well.
		d.head = nil
	} else {
		d.tail.next = nil
	}
	d.size--
	return value, nil
}

// Front returns the front element of the deque without removing it.
func (d *LinkedListDeque) Front() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque is empty")
	}
	return d.head.value, nil
}

// Rear returns the rear element of the deque without removing it.
func (d *LinkedListDeque) Rear() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque is empty")
	}
	return d.tail.value, nil
}

// IsEmpty checks if the deque is empty.
func (d *LinkedListDeque) IsEmpty() bool {
	return d.size == 0
}

// Size returns the number of elements in the deque.
func (d *LinkedListDeque) Size() int {
	return d.size
}

// Example usage.
func main() {
	var deque IDeque = &LinkedListDeque{}

	deque.EnqueueFront(10)
	deque.EnqueueRear(20)
	deque.EnqueueFront(5)

	fmt.Println("Size:", deque.Size()) // Output: Size: 3

	front, _ := deque.Front()
	fmt.Println("Front:", front) // Output: Front: 5

	rear, _ := deque.Rear()
	fmt.Println("Rear:", rear) // Output: Rear: 20

	dequeuedFront, _ := deque.DequeueFront()
	fmt.Println("Dequeued Front:", dequeuedFront) // Output: Dequeued Front: 5

	dequeuedRear, _ := deque.DequeueRear()
	fmt.Println("Dequeued Rear:", dequeuedRear) // Output: Dequeued Rear: 20

	fmt.Println("Size:", deque.Size()) // Output: Size: 1
}
