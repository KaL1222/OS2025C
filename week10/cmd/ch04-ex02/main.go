package main

import (
	"fmt"
	"log"
	"week10/pkg/keyboard"
) // custom package

func main() {
	fmt.Print("정수입력 : ")
	score, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	passFail := ""
	if score >= 60 {
		passFail = "합격!"
	} else {
		passFail = "불합격"
	}
	fmt.Printf("%.2f점은 %v\n", score, passFail)
}
