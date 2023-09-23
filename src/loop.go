package main

func main() {
	// 기본 for 루프
	var sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	// 조건식만 쓰는 for 루프
	var n = 1
	for n < 100 {
		n *= 2
	}
	println(n)

	// 무한루프
	for{
		println("Infinite loop")
	}
}
