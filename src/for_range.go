package main

func main() {
	var names = []string{"홍길동", "이순신", "강감찬"}

	for index, name := range names {
		println(index, name)
		// 0 홍길동
		// 1 이순신
		// 2 강감찬 
	}
}
