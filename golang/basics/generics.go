// generics.go
package main

import (
	"fmt"
)

func traverse[T int | string](slice []T) {
	for _, v := range slice {
		fmt.Println(v)
	}
}

func main() {
	fmt.Println("Practicing Go Generics...\n")

	ints := []int{1, 2, 3, 4, 5}
	strs := []string{"Hello", "Generics", "in", "Go"}

	fmt.Printf("Traversing int slice:\n")
	traverse(ints)

	fmt.Printf("\nTraversing string slice:\n")
	traverse(strs)
}
