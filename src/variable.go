package main

func main() {
	var a int
	println(a)
	// 실수형은 이상헤 나옴,,,,
	var f float64 = 11.
	println(f)
	a = 10
	f = 12.0
	println(a)
	println(f)

	// 복수 변수들이 선언된 상황에서 초기값을 지정할 수 있음
	var i, j, k int = 1, 2, 3
	println(i, j, k)

	// 할당되는 값을 보고 그 타입을 추론하는 기능이 있음
	var q = 1
	var s = "Hi"
	println(q)
	println(s)
}
