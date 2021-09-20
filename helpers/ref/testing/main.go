package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100.0, 0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("result of division is", result)
}
