package stack

import "fmt"

// sliceStack - стек, основанный на слайсе
type sliceStack[T comparable] struct {
	xs []T // слайс, на котором будет основан стек
}

// SNew - конструктор стека на основе слайса
func SNew[T comparable]() Stack[T] {
	return &sliceStack[T]{
		xs: make([]T, 0),
	}
}

// Push - добавление элемента в стек
func (s *sliceStack[T]) Push(value T) {
	s.xs = append(s.xs, value)
}

// Pop - удаление и возврат элемента с верхушки стека
func (s *sliceStack[T]) Pop() (result T) {
	if s.Size() == 0 {
		return
	}

	if s.Size() == 1 {
		result = s.xs[0]
		s.xs = make([]T, 0)
		return
	}

	result = s.xs[s.Size()-1]
	s.xs = s.xs[:s.Size()-1]
	return
}

// Peek - Получение верхнего элемента стека, без удаления
func (s *sliceStack[T]) Peek() (result T, ok bool) {
	if s.Size() == 0 {
		return
	}

	return s.xs[0], true
}

// Size - получение размера списка
func (s *sliceStack[T]) Size() int {
	return len(s.xs)
}

// String - реализация интерфейса fmt.Stringer
// для печати раелизуемого типа
func (s *sliceStack[T]) String() string {
	return fmt.Sprintf("%v", s.xs)
}
