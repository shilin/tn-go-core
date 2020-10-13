package fib

func Fib(n int64) int64 {
	if n == 1 {
		return 1
	}

	return Fib(n-1) * n
}
