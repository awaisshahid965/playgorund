package main

import "fmt"

func fibonacci(n int) []int {
	fmt.Println("Calculating Fibonacci sequence...")

	if n <= 0 {
		return []int{}
	}

	if n == 1 {
		return []int{0}
	}

	fibSeq := make([]int, n)
	fibSeq[0] = 0
	fibSeq[1] = 1

	for i := 2; i < n; i++ {
		fibSeq[i] = fibSeq[i-1] + fibSeq[i-2]
	}

	return fibSeq
}

func main() {
	n := 10
	fmt.Printf("First %d Fibonacci numbers: %v\n", n, fibonacci(n))
}

// Output:
// Calculating Fibonacci sequence...
// First 10 Fibonacci numbers: [0 1 1 2 3 5 8 13 21 34]
