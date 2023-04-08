package main

import (
	"time"
)

func main() {
	ticker := time.NewTicker(15 * time.Second)

	for ; true; <-ticker.C {
		PostRequest()
	}
}
