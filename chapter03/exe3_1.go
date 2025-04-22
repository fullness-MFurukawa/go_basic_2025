package chapter03

import (
	"fmt"
	"strconv"
)

func Exe3_1() {
	fmt.Print("値を入力してください->")
	var input string
	fmt.Scanln(&input)
	_, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("整数変換できない")
	} else {
		fmt.Println("整数変換できる")
	}
}
