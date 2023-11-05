package list

import (
	"fmt"
	"io"
)

type Node struct {
	Value string
	Prev  *Node
	Next  *Node
}

type NodeList struct {
	Head *Node
	Tail *Node
}

func (l *NodeList) Revert() {
	current := l.Head
	l.Tail = l.Head
	var temp *Node

	for current != nil {
		// Переставляем указатели на следующий и предыдущий узлы
		temp = current.Prev
		current.Prev = current.Next
		current.Next = temp

		// Переходим к следующему узлу
		current = current.Prev
	}

	l.Head = temp.Prev

}

func (l *NodeList) Prepend(value string) {
	n := &Node{Value: value, Prev: nil, Next: l.Head}
	l.Head.Prev = n
	l.Head = n
}

func (l *NodeList) Show(w io.Writer) {
	c := l.Head
	for c.Next != nil {
		fmt.Fprintf(w, "%#v\n", c)
		c = c.Next
	}
	fmt.Fprintf(w, "%#v\n", c)
	fmt.Fprintf(w, "---------\n")
}

func (l *NodeList) Append(value string) {
	n := &Node{Value: value, Prev: l.Tail, Next: nil}
	if l.Tail != nil {
		l.Tail.Next = n
	}
	l.Tail = n
	if l.Head == nil {
		l.Head = n
	}
}

func (l *NodeList) IsEmpty() bool {
	return l.Head == nil
}

func (l *NodeList) GetLast() *Node {
	return l.Tail
}

func (l *NodeList) GetFirst() *Node {
	return l.Head
}

func (l *NodeList) Remove(node *Node) {
	n := l.Head
	for n != nil {
		if n.Value == node.Value {
			if n.Prev != nil {
				n.Prev.Next = n.Next
			} else {
				l.Head = n.Next
			}

			if n.Next != nil {
				n.Prev.Next = n.Prev
			} else {
				l.Tail = n.Prev
			}

			return
		}
		n = n.Next
	}
}

func (l *NodeList) Search(value string) (*Node, bool) {
	n := l.Head
	for n != nil {
		if n.Value == value {
			return n, true
		}
		n = n.Next
	}

	return nil, false
}
