package stack

// Stack - список методов, которыми должен обладать стек
type Stack[T comparable] interface {
	// Push - добавление элемента в стек
	Push(value T)
	// Pop - удаление элемента из стека
	Pop() T
	// Peek - возврат верхнего элемента стека без удаления
	Peek() (T, bool)
	// Size - возвращает размер стека
	Size() int
}
