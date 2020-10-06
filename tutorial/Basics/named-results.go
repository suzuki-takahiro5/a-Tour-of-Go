package main

import "fmt"

func split(sum int) (x, y int) {

	x = sum * 4 / 9
	y = sum - x
	return
}

func split_float(sum float64) (x, y float64) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
	fmt.Println(split_float(17))
}