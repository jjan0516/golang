package main

import "fmt"

// global 변수는 어디든 참조가 가능
var g int = 10 // 패키지 전역 변수

func main(){
	// main 함수 전체 스코프
	var m int = 20 // 지역 변수
	
	// {} 스코프
	{
		var s int = 50 // 지역 변수
		fmt.Println(m, s, g)
	} // s 지역 변수는 사라짐
		// m = s + 20 // Error
}