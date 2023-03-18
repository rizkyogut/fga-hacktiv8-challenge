package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			for posisi, character := range "САШАРВО" {
				fmt.Printf("character %#U starts at byte position %d\n", character, posisi)
			}
			continue
		}
		fmt.Println("Nilai j = ", j)
	}
}
