package main

import (
	"fmt"
	"sync"
)

func main() {
	coba := []string{"coba1", "coba2", "coba3"}
	bisa := []string{"bisa1", "bisa2", "bisa3"}
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go print(i, coba, bisa, &wg, &mutex)
	}

	wg.Wait()
}

func print(index int, c []string, b []string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	fmt.Println(c, index)
	fmt.Println(b, index)
	mutex.Unlock()

	wg.Done()
}
