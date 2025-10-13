package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var month string = now.Month().String() //month := now.Month()
	var day int = now.Day()
	fmt.Println(month, day)

	var univ string = "GO$ Inha$"
	changer := strings.NewReplacer("$", "!")
	changed := changer.Replace(univ)
	fmt.Println(changed)
}
