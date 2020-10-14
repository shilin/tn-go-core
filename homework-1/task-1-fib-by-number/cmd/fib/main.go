package main

import (
	"flag"
	"fmt"
	"go-core.course/homework-1/task-1-fib-by-number/pkg/fib"
)

func main() {
	seqNum := flag.Int64("number", 5, "a Fibonacchi sequence number")
	flag.Parse()
	if (*seqNum < 0 || *seqNum > 20) {
		fmt.Println("number must be in 1..20 range")
		return 
	}
	fmt.Println(fib.Fib(*seqNum))
}
