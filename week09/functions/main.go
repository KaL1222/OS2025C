package main

import (
	"fmt"
)

func swap(first *int, second *int) {
	var temp int = 0
	temp = *first
	*first = *second
	*second = temp
	fmt.Println(first, second)
}

func main() {
	a, b := 10, 20
	fmt.Println(a, b)
	swap(&a, &b) //call me pointer
	fmt.Println(a, b)
}
