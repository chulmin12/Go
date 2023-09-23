package main

// pass by value
func main() {
	var msg = "Hello"
	say(msg)
}

// 함수 생성 
func say(msg string) {
	println(msg)
}

// pass by reference
func main() {
	var msg = "Hello"
	say(&msg)
	println(msg) // 변경된 메시지 출력
}

func say(msg *string) {
	println(*msg) // 요청 받은 메시지 출력
	*msg = "Changed"
}

// variadic function(가변인자함수)
func main(){
	say("This", "is", "a", "book")
	say("Hi")
}

func say(msg ...string){
	for _, s := range msg{
		println(s)
	}
}