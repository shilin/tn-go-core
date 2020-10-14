package fib

func Num(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	return helper([2]int{0, 1}, n)
}

func helper(accumulator [2]int, k int) int {
	if k > 1 {
		nextAccumulator := [2]int{accumulator[1], accumulator[0] + accumulator[1]}
		return helper(nextAccumulator, k-1)
	}
	return accumulator[1]
}
