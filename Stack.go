package main

type Stack[T any] struct {
	top *SNode[T]
}

type SNode[T any] struct {
	val   T
	below *SNode[T]
}

func InitStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (stack Stack[T]) Push(val T) {
	temp := SNode[T]{
		val:   val,
		below: stack.top,
	}
	stack.top = &temp
}

func (stack Stack[T]) Pop() T {
	temp := stack.top
	stack.top = stack.top.below
	return temp.val
}

func (stack Stack[T]) Peek() T {
	return stack.top.val
}
