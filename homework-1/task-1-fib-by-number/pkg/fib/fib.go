package fib

func Num(n int64) int64 {
	if n == 1 {
		return 1
	}

	return Num(n-1) * n
}
