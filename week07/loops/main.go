package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	//i, _ := r.ReadString('\n')
	i, err := r.ReadString('\n')
	//fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)
}
