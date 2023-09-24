package main

import "fmt"

func main() {
	// map 선언
	var idMap map[int]string
	// 초기화
	idMap = make(map[int]string)
	idMap[901] = "Apple"
	idMap[801] = " Grape"

	// 키에 대한 값 읽기
	str := idMap[901]
	println(str)
	// 삭제
	delete(idMap, 901)

	noData := idMap[901]
	println(noData) // 값이 없으면 nill 혹은 zero리턴

	// 리터럴을 사용한 초기화
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
	}
	str1 := tickers["GOOG"]
	println(str1)

	// map 키 체크
	val, exists := tickers["MSFT"]
	if !exists {
		println("No MSFT ticker")
	}

	// for 루프를 사용한 Map 열거
	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}
