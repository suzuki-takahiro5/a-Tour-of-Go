package main

import "fmt"

func main() {
	// defer fmt.Println("world")

	// fmt.Println("hello")

	n := 1
	defer fmt.Println(n)
	n += 1
	fmt.Println(n)
}