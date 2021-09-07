package main

import (
	"github.com/byitkc/goland-web-learning/helpers"
	"log"
)

const numPool = 1000

func CalculateValue(c chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	c <- randomNumber
}

func main() {
	// Working with channels
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
