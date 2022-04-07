package main

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

type Node[T any] struct {
	Val     T
	Next   *Node[T]
	Before *Node[T]
}

func InitList[T any](val T) *LinkedList[T] {
	node := &Node[T]{
		Val: val,
	}
	return &LinkedList[T]{
		Head: node,
		Tail: node,
	}

}
