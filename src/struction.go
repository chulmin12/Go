package main

import "fmt"

// struct 정의
type person struct {
	name string
	age  int
}

func main() {
	// 객체 생성
	p := person{}
	// 필드값 설정
	p.name = "Lee"
	p.age = 10

	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}
	fmt.Println(p)
	fmt.Println(p1, p2)
}
