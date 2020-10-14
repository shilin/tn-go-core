package fib

import "testing"

func TestFib(t *testing.T) {
	const want int64 = 120
	got := Num(5)
	if got != want {
		t.Fatalf("пятый элемент ряда Фибоначчи %d, а ожидалось %d", got, want)
	}
}
