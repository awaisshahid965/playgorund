package main

import "fmt"


func main() {
	fmt.Println("Non-primitive data types in Go:\n")
	fmt.Println("Working with Arrays and Slices in Go:\n")

	// Arrays (size defined at compile time)
	var array [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", array)

	// Array with inferred size
	var arrayInferredSize = [...]int{4, 5, 6, 7}
	fmt.Println("Array with inferred size:", arrayInferredSize)

	fmt.Println("\n")

	// Slices (dynamic size)
	var slice []int = []int{8, 9, 10}
	slice = append(slice, 11) // Appending an element
	slice = append(slice, 12, 13, 14) // Appending multiple elements
	slice = append(slice, []int{15, 16, 17}...) // Appending another slice
	slice = append(slice, array[:]...) // Appending an array
	fmt.Println("Slice after appends:", slice)
	fmt.Println("Slice Length:", len(slice))
	fmt.Println("Slice Capacity:", cap(slice))

	fmt.Println("\n")

	// slices using make
	sliceMake := make([]int, 5, 10) // length 5, capacity 10
	fmt.Println("Slice created with make:", sliceMake)
	fmt.Println("Slice Length (make):", len(sliceMake))
	fmt.Println("Slice Capacity (make):", cap(sliceMake))

	fmt.Println("\n")

	// traversing an array (slices will be the same)
	fmt.Println("Traversing the slice:")
	for index, value := range array {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	fmt.Println("\n\n\n\n")

	fmt.Println("Working with Maps in Go:\n")

	// Maps (key-value pairs)
	var mapData map[string]uint8 = map[string]uint8{
		"apple":  5,
		"banana": 10,
	}
	fmt.Println("Map:", mapData)
	fmt.Println("Value for 'apple':", mapData["apple"])

	// Adding a new key-value pair
	mapData["cherry"] = 15
	fmt.Println("Map after adding 'cherry':", mapData)

	// Deleting a key-value pair
	delete(mapData, "banana")
	fmt.Println("Map after deleting 'banana':", mapData)

	// Checking if a key exists
	value, exists := mapData["banana"]
	if exists {
		fmt.Println("Value for 'banana':", value)
	} else {
		fmt.Println("'banana' key does not exist in the map")
	}

	// traversing a map
	fmt.Println("Traversing the map:")
	for key, value := range mapData {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
}
