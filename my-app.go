package main

import (
	"os"
	"time"
)

func main() {
	for {
		panic(os.Getenv("GREETING"))
		time.Sleep(1*time.Second)
	}
}
