package main

import (
	"fmt"
)

func main() {
	//arrayInt := [3]int{-9, 11, 7}
	//for i := 0; i < len(arrayInt); i++ {
	//	fmt.Println(i, arrayInt[i])
	//}
	numbers := [3]int{-9, 11, 7}
	for i, number := range numbers {
		fmt.Println(i, number)
	}
}
