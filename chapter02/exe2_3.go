package chapter02

import (
	"fmt"
	"strconv"
)

func Exe2_3() {
	var input string
	fmt.Print("値を入力してください->")
	fmt.Scanln(&input)
	value, _ := strconv.Atoi(input)
	if value%4 == 0 {
		fmt.Println("入力された値は4の倍数です")
	} else {
		fmt.Println("入力された値は4の倍数ではありません")
	}
}
