package main

import "fmt"

func fibonacci(n int) []int {
	fmt.Println("Calculating Fibonacci sequence...")

	if n <= 0 {
		return []int{}
	}

	fibSeq := make([]int, n+1)
	fibSeq[0] = 0
	fibSeq[1] = 1

	for i := 2; i < n; i++ {
		fibSeq[i] = fibSeq[i-1] + fibSeq[i-2]
	}

	return fibSeq
}

func main() {
	n := 1
	fmt.Printf("First %d Fibonacci numbers: %v\n", n, fibonacci(n))
}
