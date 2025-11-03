package korean

import "fmt"

func Anyunghaseyo() {
	fmt.Println("안녕하세요!")
	// anyung() 같은 함수 내에선 접근 가능
}

// func anyung() { // 함수명이 소문자로 시작하는 경우 외부 파일에서 접근이 불가
func Anyung() {
	fmt.Println("안녕!")
}
