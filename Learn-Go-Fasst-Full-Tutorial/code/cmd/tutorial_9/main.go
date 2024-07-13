package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5

func main() {

	var chickenChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
	}
	sendMessage(chickenChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(1 * time.Second)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Found a deal on chicken at %s\n", <-chickenChannel)
	}
}
