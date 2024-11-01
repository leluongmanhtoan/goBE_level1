package lession1

import "fmt"

func Div2Check() {
	var str string
	fmt.Println("----------Excercise 2----------")
	fmt.Print("Nhap vao mot chuoi: ")
	fmt.Scanln(&str)
	fmt.Println("Chieu dai chuoi chia het cho 2: ", isDivBy2(str))
}

func isDivBy2(str string) bool {
	return len(str)%2 == 0
}
