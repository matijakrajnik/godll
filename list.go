// Doubly linked list.

package godll

import (
	"errors"
	"fmt"
	"io"
)

var (
	IndexOutOfRangeError = errors.New("Index is out of range!")
)

type List[T any] struct {
	head   *Node[T] // Pointer to head (first node in list).
	tail   *Node[T] // Pointer to tail (last node in list).
	length int      // Number of nodes in list.
}

// Head returns first node in list.
func (l *List[T]) Head() *Node[T] {
	return l.head
}

// Tail returns last node in list.
func (l *List[T]) Tail() *Node[T] {
	return l.tail
}

// Length returns number of List[T] length field.
func (l *List[T]) Length() int {
	return l.length
}

// Print prints all elements in a List using passed io.Writer interface.
func (l *List[T]) Print(w io.Writer) {
	if l.length == 0 {
		return
	}
	current := l.head
	for i := 1; i <= l.length; i++ {
		fmt.Fprintf(w, "%v ", current.Value)
		current = current.next
	}
	fmt.Fprintf(w, "\n")
}

// Append adds node to the end of the List.
func (l *List[T]) Append(node *Node[T]) {
	// Set node as a head and tail if list is empty.
	if l.length == 0 {
		l.head = node
		l.tail = node
		l.length++
		return
	}

	// Add current tail to new variable oldTail. Set passed node as a new tail and connect it to old tail with new next and previous links.
	oldTail := l.tail
	l.tail = node
	node.previous = oldTail
	oldTail.next = node
	l.length++
}

// Prepend adds node to the beggining of the List.
func (l *List[T]) Prepend(node *Node[T]) {
	// Set node as a head and tail if list is empty.
	if l.length == 0 {
		l.head = node
		l.tail = node
		l.length++
		return
	}

	// Add current head to new variable oldHead. Set passed node as a new head and connect it to old head with new next and previous links.
	oldHead := l.head
	l.head = node
	node.next = oldHead
	oldHead.previous = node
	l.length++
}

func (l *List[T]) InsertAt(index int, node *Node[T]) error {
	// Return error if index is larger than legth of list.
	if index > l.length {
		return IndexOutOfRangeError
	}

	// If index is 0 set node as new head of list and connect it to neighbours with new next and previous links.
	// If length is 0, set node as tail also.
	if index == 0 {
		node.next = l.head
		if l.length != 0 {
			l.head.previous = node
		}
		l.head = node
		if l.length == 0 {
			l.tail = node
		}
		l.length++
		return nil
	}

	// If index is equal to length of list set node as tail of a list and connect it to neighbours with new next and previous links.
	if index == l.length {
		node.previous = l.tail
		l.tail.next = node
		l.tail = node
		l.length++
		return nil
	}

	// Traverse through list to find index, insert new node and connect it to neighbours with new next and previous links.
	current := l.head
	for i := 1; i < index; i++ {
		current = current.next
	}
	next := current.next
	current.next = node
	node.next = next
	node.previous = current
	next.previous = node
	l.length++
	return nil
}

// GetByIndex retrieves node by index. Return error if index is out of range. Index of first node is 0.
func (l *List[T]) GetByIndex(index int) (*Node[T], error) {
	if index >= l.length {
		return nil, IndexOutOfRangeError
	}

	// Calculate index in the middle of the list.
	m := l.length / 2

	// If index is closer to head, start iterating through nodes from head.
	if index < m {
		current := l.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		return current, nil
	}

	// If index is closer to tail, start iterating through nodes from tail.
	current := l.tail
	for i := 0; i < (l.length - index - 1); i++ {
		current = current.previous
	}
	return current, nil
}
