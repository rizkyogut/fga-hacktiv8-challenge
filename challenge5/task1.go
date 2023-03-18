package main

import (
	"fmt"
	"sync"
)

func main() {
	coba := []string{"coba1", "coba2", "coba3"}
	bisa := []string{"bisa1", "bisa2", "bisa3"}
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {

		wg.Add(1)
		go printCoba(i, coba, &wg)

		wg.Add(1)
		go printBisa(i, bisa, &wg)
	}

	wg.Wait()
}

func printCoba(index int, c []string, wg *sync.WaitGroup) {
	fmt.Println(c, index)
	wg.Done()
}

func printBisa(index int, b []string, wg *sync.WaitGroup) {
	fmt.Println(b, index)
	wg.Done()
}
