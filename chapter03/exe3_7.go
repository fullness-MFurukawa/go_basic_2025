package chapter03

import (
	"fmt"
	"strconv"
)

type Position struct {
	Employee
	PositionName string // 役職名
}

func Exe3_7() {
	var position Position
	var input string
	fmt.Print("社員番号を入力してください->")
	fmt.Scanln(&input)
	position.EmpNo, _ = strconv.Atoi(input)
	fmt.Print("社員名を入力してください->")
	fmt.Scanln(&input)
	position.EmpName = input
	fmt.Print("各職を入力してください->")
	fmt.Scanln(&input)
	position.PositionName = input
	fmt.Println(
		"社員番号=", position.EmpNo,
		"社員名=", position.EmpName,
		"役職=", position.PositionName)
}
