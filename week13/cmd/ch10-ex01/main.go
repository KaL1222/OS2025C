package main

import (
	"fmt"
	calendar "week11/cmd/pkg/calender"
)

func main() {
	today := calendar.Date{}
	today.SetYear(2025)
	today.SetMonth(11)
	today.SetDay(24)
	// fmt.Println(today.Year(), "년 ", today.Month(), "월 ", today.Day(), "일")
	fmt.Println("%d년 %d월 %d일", today.Year(), today.Month(), today.Day())
}
