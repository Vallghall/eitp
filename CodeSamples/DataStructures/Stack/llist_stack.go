package stack

import list "eitp/CodeSamples/DataStructures/LinkedList/go"

// listStack - стек на основе связного списка
type listStack[T comparable] struct {
	xs *list.List[T]
}

// LNew - конструктор стека
func LNew[T comparable]() Stack[T] {
	return &listStack[T]{
		xs: list.New[T](),
	}
}

// Push - добавление элемента в стек
func (l *listStack[T]) Push(value T) {
	l.xs.Prepend(value)
}

// Pop - удаление и возврат элемента с верхушки стека
func (l *listStack[T]) Pop() T {
	return l.xs.Shift()
}

// Peek - Получение верхнего элемента стека, без удаления
func (l *listStack[T]) Peek() (result T, ok bool) {
	if l.Size() > 0 {
		return l.xs.Head(), true
	}

	return result, false
}

// Size - получение размера списка
func (l *listStack[T]) Size() int {
	return l.xs.Length()
}
