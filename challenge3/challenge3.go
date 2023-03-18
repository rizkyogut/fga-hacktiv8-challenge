package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "selamat malam"
	split := strings.Split(text, "")
	count := make(map[string]int)

	for _, v := range split {
		fmt.Println(v)
		count[v]++
	}

	fmt.Print("map : ")
	for i, v := range count {
		fmt.Printf("[ %s: %d ]", i, v)
	}
}
