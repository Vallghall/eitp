package list

import (
	"fmt"
	"strings"
)

// node - узел связанного списка
type node[T comparable] struct {
	value T
	next  *node[T] // указатель на следующий элемент списка
}

// List - обертка/фиктивный узел списка
type List[T comparable] struct {
	length int
	head   *node[T] // указатель на голову - первый элемент - списка
}

// New - конструктор списка
func New[T comparable](params ...T) *List[T] {
	list := &List[T]{length: len(params)}

	if list.length == 0 {
		return list
	}

	list.head = &node[T]{value: params[0]}
	cur := list.head
	for i := 1; i < len(params); i++ {
		cur.next = &node[T]{value: params[i]}
		cur = cur.next
	}

	return list
}

// Prepend - вставка элемента в начало списка - долго
func (l *List[T]) Prepend(value T) *List[T] {
	n := &node[T]{
		value: value,
		next:  l.head,
	}
	l.head = n
	l.length++

	return l
}

// Append - вставка элемента в конец списка - медленно
func (l *List[T]) Append(value T) *List[T] {
	n := &node[T]{value: value}
	l.length++

	if l.head == nil {
		l.head = n
		return l
	}

	var cur *node[T]
	for cur = l.head; cur.next != nil; cur = cur.next {
	}

	cur.next = n
	return l
}

// AddAt - вставка элемента в середину списка - тоже медленно
func (l *List[T]) AddAt(idx int, value T) (*List[T], error) {
	if idx < 0 && idx > l.length {
		return nil, fmt.Errorf("index %d is out of range: list length is %d", idx, l.length)
	}

	if idx == 0 {
		return l.Prepend(value), nil
	}

	if idx == l.length {
		return l.Append(value), nil
	}

	cur := l.head
	for j := 0; j < idx-1; j++ {
		cur = cur.next
	}

	n := &node[T]{
		value: value,
		next:  cur.next,
	}

	cur.next = n
	l.length++
	return l, nil
}

// Shift - удаление и возвращение элемента из начала списка
func (l *List[T]) Shift() T {
	result := l.head.value
	l.head = l.head.next
	l.length--

	return result
}

// Pop - удаление и возвращение элемента из конца списка
func (l *List[T]) Pop() T {
	if l.head == nil {
		var val T
		return val
	}

	if l.head.next == nil {
		val := l.head.value
		l.head = nil
		return val
	}

	prev := l.head
	cur := prev.next
	for cur.next != nil {
		prev = cur
		cur = cur.next
	}

	prev.next = nil
	l.length--
	return cur.value
}

// Delete - удаление и возвращение элемента по индексу
func (l *List[T]) Delete(idx int) (T, error) {
	if idx < 0 && idx >= l.length {
		var val T
		return val, fmt.Errorf("index %d is out of range: list length is %d", idx, l.length)
	}

	if idx == 0 {
		return l.Shift(), nil
	}

	prev := l.head
	cur := prev.next
	for j := 1; j < idx; j++ {
		prev = cur
		cur = cur.next
	}

	prev.next = cur.next
	l.length--
	return cur.value, nil
}

// Get - получение элемента по индексу
func (l List[T]) Get(idx int) (T, error) {
	if idx < 0 && idx >= l.length {
		var val T
		return val, fmt.Errorf("index %d is out of range: list length is %d", idx, l.length)
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}

	return cur.value, nil
}

// Set - установка элемента по индексу
func (l List[T]) Set(idx int, value T) error {
	if idx < 0 && idx >= l.length {
		return fmt.Errorf("index %d is out of range: list length is %d", idx, l.length)
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}

	cur.value = value
	return nil
}

// Head - получение значения первого элемента списка
func (l *List[T]) Head() T {
	return l.head.value
}

// Tail - получение хвоста списка
func (l *List[T]) Tail() *List[T] {
	return &List[T]{
		length: l.length - 1,
		head:   l.head.next,
	}
}

// Find - поиск позиции первого совпадающего элемента
func (l *List[T]) Find(value T) int {
	cur := l.head
	for j := 0; j < l.length; j++ {
		if cur.value == value {
			return j
		}

		cur = cur.next
	}

	return -1
}

func (l List[T]) String() string {
	sb := new(strings.Builder)
	sb.WriteRune('[')

	for cur := l.head; cur != nil; cur = cur.next {
		sb.WriteString(fmt.Sprintf("%v,", cur.value))
	}

	return strings.TrimRight(sb.String(), ",") + "]"
}
