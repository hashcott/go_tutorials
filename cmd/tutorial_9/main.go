package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX_CHICKENS_PRICE float32 = 5
	MAX_TOFUS_PRICE    float32 = 2
)

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)

	var websites = []string{"walmart.com", "costco.com", "wholefoods.com", "amazon.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice < MAX_CHICKENS_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice < MAX_TOFUS_PRICE {
			tofuChannel <- website
			break
		}
	}
}
func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nFound deals on chicken at: %s", website)
	case website := <-tofuChannel:
		fmt.Printf("\nFound deals on tofu at: %s", website)
	}

}
