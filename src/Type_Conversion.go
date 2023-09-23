package main

func main() {
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(u, f)

	var str = "ABC"
	var bytes = []byte(str)
	var str2 = string(bytes)
	println(bytes, str2)
}
