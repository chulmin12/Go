package main

// fmt 패키지는 포맷 템플릿이다.
import "fmt"

func main() {
	var a []int        //슬라이스 변수 선언
	a = []int{1, 2, 3} // 슬라이스에 리터럴값 지정
	a[1] = 10
	println(a)     // 3/3]0x14000014120 이 출력됨
	fmt.Println(a) // [1 10 3] 이 출력됨

	s := make([]int, 5, 10)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	s[3] = 4
	s[4] = 5
	println(len(s), cap(s)) // 5, 10
	fmt.Println(s)
}
