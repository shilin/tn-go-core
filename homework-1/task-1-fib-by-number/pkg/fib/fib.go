package fib

func helper(lastTwoFibNums [2]int64, k int64) int64 {
	if k > 1 {
		nextLastTwo := [2]int64{lastTwoFibNums[1], lastTwoFibNums[0] + lastTwoFibNums[1]}
		return helper(nextLastTwo, k-1)
	}
	return lastTwoFibNums[1]
}

func Num(n int64) int64 {
	if n == 1 || n == 0 {
		return 1
	}

	return helper([2]int64{0, 1}, n)
}
