package list

import (
	"fmt"
)

// Пакет реализует двусвязный список вместе с базовыми операциями над ним.

// List - двусвязный список.
type List struct {
	root Elem
}

// Elem - элемент списка.
type Elem struct {
	Val        interface{}
	next, prev *Elem
}

// New - конструктор нового списка.
func New() *List {
	var l List
	l.root.next = &l.root
	l.root.prev = &l.root
	return &l
}

// Push вставляет элемент в начало списка.
func (l *List) Push(e Elem) *Elem {
	e.prev = &l.root
	e.next = l.root.next
	l.root.next = &e
	e.next.prev = &e

	return &e
}

// String реализует интерфейс fmt.Stringer представляя список в виде строки.
func (l *List) String() string {
	el := l.root.next
	var s string
	for el != &l.root {
		s += fmt.Sprintf("%v ", el.Val)
		el = el.next
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

// Pop удаляет первый элемент списка.
func (l *List) Pop() *List {
	// no elements in the list except for root
	if l.root.next == &l.root {
		return l
	}

	// exactly one element in the list
	if l.root.next.next == &l.root {
		l.root.next = &l.root
		return l
	}

	// two or more elements in the list
	l.root.next = l.root.next.next
	l.root.next.next.prev = &l.root
	return l
}

// Reverse разворачивает список.
func (l *List) Reverse() *List {

	current := &l.root
	notEnd := true
	for notEnd {
		if current.next == &l.root {
			notEnd = false
		}
		temp := current.next
		current.next = current.prev
		current.prev = temp
		current = temp
	}
	return l
}
