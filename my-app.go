package main

import (
	"log"
	"os"
	"time"
)

var greeting = os.Getenv("GREETING")

func main() {
	for {
		log.Println(os.Getenv("GREETING"))
		time.Sleep(5000)
	}
}
