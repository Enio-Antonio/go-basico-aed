package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedListQueue struct {
	head *Node
	tail *Node
	size int
}

type IQueue interface {
	Enqueue(value int)
	Dequeue() (int, error)
	Front() (int, error)
	IsEmpty() bool
	Size() int
}

// Enqueue adds an element to the end of the queue.
func (q *LinkedListQueue) Enqueue(value int) {
	newNode := &Node{value: value}
	if q.tail == nil {
		// If the queue is empty, head and tail point to the new node.
		q.head = newNode
	} else {
		// Add the new node to the end of the queue.
		q.tail.next = newNode
	}
	q.tail = newNode
	q.size++
}

// Dequeue removes and returns the front element of the queue.
func (q *LinkedListQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	value := q.head.value
	q.head = q.head.next
	if q.head == nil {
		// If the queue becomes empty, reset the tail as well.
		q.tail = nil
	}
	q.size--
	return value, nil
}

// Front returns the front element of the queue without removing it.
func (q *LinkedListQueue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return q.head.value, nil
}

// IsEmpty checks if the queue is empty.
func (q *LinkedListQueue) IsEmpty() bool {
	return q.size == 0
}

// Size returns the number of elements in the queue.
func (q *LinkedListQueue) Size() int {
	return q.size
}

// Example usage.
func main() {
	var queue IQueue = &LinkedListQueue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println("Size:", queue.Size()) // Output: Size: 3

	front, _ := queue.Front()
	fmt.Println("Front:", front) // Output: Front: 10

	dequeued, _ := queue.Dequeue()
	fmt.Println("Dequeued:", dequeued) // Output: Dequeued: 10

	fmt.Println("Size:", queue.Size()) // Output: Size: 2

	fmt.Println("IsEmpty:", queue.IsEmpty()) // Output: IsEmpty: false
}

