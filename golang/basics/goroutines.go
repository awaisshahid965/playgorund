package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

var advices = []string{
	// instagram Oogway advices
	"Patience you must have… especially when WiFi is slow.",
	"The universe has a plan… but it did not plan for you to skip leg day.",
	"One often meets their destiny… on the path they took to avoid doing their chores.",
	"There are no accidents… except when you forget your keys again.",
	"Your mind is like water… which explains why mine keeps leaking ideas.",
	"To make something special… you just have to believe in the magic of coffee.",
	"Sometimes the best way to solve a problem… is to take a nap first.",
}

func main() {
	fmt.Println("Practing Go Routines...\n")
	fmt.Println("Welcome to the world of Oogway's wisdom:\n")
	
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(GetAdvise())
		}()
	}
	wg.Wait()
}

func GetAdvise() string {
	index := rand.Intn(len(advices))
	return advices[index]
}
