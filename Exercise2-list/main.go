package main

import "fmt"

func main() {
	l := &List[int]{}

	l.Add(1)
	l.Add(2)
	l.Add(3)

	fmt.Println(l)

	l2 := &List[string]{}

	l2.Add("hello kitty")
	l2.Add("hasta la vista")
	l2.Add("no way jose")

	fmt.Println(l2)
}

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (l *List[T]) Add(value T) {
	n := &Node[T]{Value: value}

	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.Next = n
		l.tail = n
	}
}

func (l *List[T]) String() string {
	str := "["

	for n := l.head; n != nil; n = n.Next {
		str += fmt.Sprintf("%v", n.Value)
		if n.Next != nil {
			str += " => "
		}
	}

	str += "]"

	return str
}
