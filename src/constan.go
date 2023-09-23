package main

func main()
{
	const c int = 10
	const s string = "Hi"

	// 할당되는 값을 보고 그 타입을 추론하는 기능
	const a = 20
	const st = "const Hi"

	// 여러 개의 상수들을 묶어서 지정할 수 있음
	const(
		Visa = "Visa"
		Master = "Master"
		Amex = "American Express"
	)

	// iota을 지정하는 경우 iota이 지정된 변수는 0이 할당되고,
	// 나머지는 순서대로 1씩 증가된 값을 부여 받음
	const(
		Apple = iota    // 0
		Grape			// 1
		Orange			// 2
	)

}