package main

import "fmt"

// ok
var t = 4
// ng
// t := 4

func main() {
	var i, j int = 1, 2
	k := 3
	c, python ,java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java, t)
}