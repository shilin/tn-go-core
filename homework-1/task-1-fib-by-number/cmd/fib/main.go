package main

import (
	"fmt"
	"go-core.course/homework-1/task-1-fib-by-number/pkg/fib"
)

func main() {
	fmt.Println(fib.Fib(5))
	fmt.Println(fib.Fib(15))
	fmt.Println(fib.Fib(20))
}
