package main

import (
	"fmt"
	"log"
	calendar "week11/cmd/pkg/calender"
)

func main() {
	today := calendar.Event{}
	today.SetTitle("Final Exam D-14...............................")

	err := today.SetYear(2025)
	if err != nil {
		log.Fatal(err)
	}
	err = today.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}
	// err = today.setDay(77)
	err = today.SetDay(24)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n%d년 %d월 %d일", today.Title(), today.Year(), today.Month(), today.Day())
}
