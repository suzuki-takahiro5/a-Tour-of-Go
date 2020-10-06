package main

import "fmt"

func main(){
	v1 := 42
	fmt.Printf("v is of type %T\n", v1)
	v2 := 3.142
	fmt.Printf("v is of type %T\n", v2)
	v3 := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", v3)
	v4 := "aaaaa"
	fmt.Printf("v is of type %T\n", v4)
	v5 := true
	fmt.Printf("v is of type %T\n", v5)
}

