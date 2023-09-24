package main

func main() {
	// 3개의 배열 생성
	var a [3]int
	a[0] = 1
	a[1] = 2

	//배열크기 자동으로 할당
	var b = [...]int{1, 2, 3}
	println(b[2])

	// 다차원 배열
	var arr = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, // 끝에 콤마 추가
	}
	println(arr[1][2])

	/* [...] 다차원 배열은 불가능
	var arr1 = [...][...]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	println(arr1[1][2])
	*/
}
