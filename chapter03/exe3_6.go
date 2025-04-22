package chapter03

import (
	"fmt"
	"strconv"
)

type Employee struct {
	EmpNo   int
	EmpName string
}

func Exe3_6() {
	var employee Employee
	var input string
	fmt.Print("社員番号を入力してください->")
	fmt.Scanln(&input)
	employee.EmpNo, _ = strconv.Atoi(input)
	fmt.Print("社員名を入力してください->")
	fmt.Scanln(&input)
	employee.EmpName = input
	fmt.Println("社員番号=", employee.EmpNo, "社員名=", employee.EmpName)
}
