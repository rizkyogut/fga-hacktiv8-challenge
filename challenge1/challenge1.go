package main

import (
	"fmt"
)

func main(){
	var i int = 21
	const j bool = true
	var k float64 = 123.456
	l := '\u042F' 
	m := 15

	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	fmt.Printf("%t \n", j)
	fmt.Printf("%t \n", !j)
	fmt.Printf("%c \n", l)
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%x \n", m)
	fmt.Printf("%X \n", m)
	fmt.Printf("%U \n", l)
	fmt.Printf("%.6f \n" , k)
	fmt.Printf("%E \n" , k)
}