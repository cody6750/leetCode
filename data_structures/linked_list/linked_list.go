package main

import (
	"fmt"
)

type LinkedList struct {
	Root *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (l *LinkedList) insert(v int) {
	if l.Root == nil {
		l.Root = &Node{Value: v}
		return
	}
	current := l.Root
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{Value: v}
}

func (l *LinkedList) Print() {
	if l.Root == nil {
		return
	}
	current := l.Root
	for current != nil {
		fmt.Printf("%v ", current.Value)
		// log.Print(current.Value)
		current = current.Next
	}
}

func (l *LinkedList) Delete(v int) {
	if l.Root == nil {
		return
	}

	if l.Root.Value == v {
		l.Root = nil
		return
	}

	current := l.Root.Next
	prev := l.Root
	for current != nil {
		if current.Value == v {
			if current.Next == nil {
				prev.Next = nil
			} else {
				prev.Next = current.Next
			}
			return
		}

		prev, current = current, current.Next
	}
	return
}

// 1 2 3 4 5
func (l *LinkedList) Reverse() {
	if l.Root == nil || l.Root.Next == nil {
		return
	}
	current := l.Root
	var prev *Node
	for current != nil {
		tmp := current.Next // 2
		current.Next = prev // nil
		prev = current      // 1
		current = tmp       // 2
	}
	l.Root = prev

}

func initLinkedList() {
	linkedList := LinkedList{}
	linkedList.insert(0)
	linkedList.insert(1)
	linkedList.insert(2)
	linkedList.insert(3)
	linkedList.insert(4)
	linkedList.insert(5)
	linkedList.insert(6)
	linkedList.Delete(1)
	linkedList.Print()
	linkedList.Reverse()
	fmt.Println()
	linkedList.Print()
}
