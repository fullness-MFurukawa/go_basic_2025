package chapter05

import (
	"fmt"
	"strconv"
)

type Account struct {
	no      string // 口座番号
	name    string // 口座名義
	balance int    // 残高
}

// Print()メソッド
func (value Account) Print() {
	fmt.Println("口座番号 = ", value.no)
	fmt.Println("口座名義 = ", value.name)
	fmt.Println("残高 	 = ", value.balance)
}

func Exe5_1() {
	var input string
	var account Account
	fmt.Print("口座番号を入力してください->")
	fmt.Scanln(&account.no)
	fmt.Print("口座名義を入力してください->")
	fmt.Scanln(&account.name)
	fmt.Print("残高を入力してください->")
	fmt.Scanln(&input)
	account.balance, _ = strconv.Atoi(input)
	account.Print()
}
