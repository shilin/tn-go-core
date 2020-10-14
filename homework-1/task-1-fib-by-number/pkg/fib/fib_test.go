package fib

import "testing"

func TestFib(t *testing.T) {
	const want = 13
	got := Num(7)
	if got != want {
		t.Fatalf("пятый элемент ряда Фибоначчи %d, а ожидалось %d", got, want)
	}
}
