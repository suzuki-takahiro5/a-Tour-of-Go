package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = 1.0
	var before_z float64
	for i  := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)

		if before_z == z { 
			break
		}
		
		before_z = z
	}
	
	return
}

func main() {
	var n float64 = 2
	fmt.Println(Sqrt(n))
	fmt.Println(math.Sqrt(n))
}
