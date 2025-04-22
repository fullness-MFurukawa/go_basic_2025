package chapter02

import (
	"fmt"
	"strconv"
)

// 演習 2-1
func Exe2_1() {
	var input string
	fmt.Print("値1を入力してください->")
	fmt.Scanln(&input) // 標準入力から値を取得
	// 文字列を整数に変換する
	value1, _ := strconv.Atoi(input)

	fmt.Print("値2を入力してください->")
	fmt.Scanln(&input) // 標準入力から値を取得
	// 文字列を整数に変換する
	value2, _ := strconv.Atoi(input)

	fmt.Println("値1 + 値2 = ", value1+value2)
	fmt.Println("値1 - 値2 = ", value1-value2)
}
