package main

import (
	"fmt"
	"math/rand"
)

var MAX_CHICKEN_PRICE float64 = 21.0
var MAX_TOFU_PRICE float64 = 21.0

func main() {
	fmt.Println("Practicing Go Channels...\n")

	chickenChannel := make(chan string)
	tofuChannel := make(chan string)
	websites := [3]string{"walmart.com", "costco.com", "wholefoods.com"}

	for _, website := range websites {
		go checkChickenPrices(chickenChannel, website)
		go checkTofuPrices(tofuChannel, website)
	}

	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(ch chan string, website string) {
	for {
		price := rand.Float64() * 20.0
		if price < MAX_CHICKEN_PRICE {
			ch <- website
			break
		}
	}
}

func checkTofuPrices(ch chan string, website string) {
	for {
		price := rand.Float64() * 20.0
		if price < MAX_TOFU_PRICE {
			ch <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
		case website := <-chickenChannel:
			fmt.Printf("Chicken is cheap at %s!\n", website)
		case website := <-tofuChannel:
			fmt.Printf("Tofu is cheap at %s!\n", website)
	}
}
