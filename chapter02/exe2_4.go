package chapter02

import (
	"fmt"
	"strconv"
)

func Exe2_4() {
	fmt.Print("値を入力してください->")
	var input string
	fmt.Scanln(&input)
	value, _ := strconv.Atoi(input)
	switch {
	case value > 0:
		fmt.Println("正の値です")
	case value < 0:
		fmt.Println("負の値です")
	default:
		fmt.Println("0です")
	}
}
