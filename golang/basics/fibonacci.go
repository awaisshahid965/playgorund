package main

import "fmt"

func fibonacci(n int) []int {
	fmt.Println("Calculating Fibonacci sequence...")
	return []int{}
}

func main() {
	n := 10
	fmt.Printf("First %d Fibonacci numbers: %v\n", n, fibonacci(n))
}
