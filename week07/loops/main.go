package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a score")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal((err))
	}

	i = strings.TrimSpace(i)
	score, err := strconv.ParseFloat(i, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if score >= 60 {
		status = "Pass"
	} else {
		status = "Fail"
	}
	fmt.Println(score, status)
	/*
		var float64 float64 = 2.71
		var f float64 = 3.991
		fmt.Println(float64)
		fmt.Println(f)
		var fmt float64 = 2.71
		fmt.Println(fmt)
	*/
}
