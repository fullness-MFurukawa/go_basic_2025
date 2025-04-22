package chapter04

import (
	"fmt"
	"strconv"
)

/*
* 四則演算結果を返す
 */
func calc(x int, y int) (int, int, int, int) {
	add := x + y
	sub := x - y
	mul := x * y
	div := x / y
	return add, sub, mul, div
}
func Exe4_1() {
	var input string
	fmt.Print("値1を入力してください->")
	fmt.Scanln(&input) // 標準入力から値を取得
	// 文字列を整数に変換する
	value1, _ := strconv.Atoi(input)
	fmt.Print("値2を入力してください->")
	fmt.Scanln(&input) // 標準入力から値を取得
	// 文字列を整数に変換する
	value2, _ := strconv.Atoi(input)
	add, sub, mul, div := calc(value1, value2)
	fmt.Println("値1 + 値2 = ", add)
	fmt.Println("値1 - 値2 = ", sub)
	fmt.Println("値1 * 値2 = ", mul)
	fmt.Println("値1 / 値2 = ", div)
}
