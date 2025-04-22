package chapter03

import (
	"fmt"
	"reflect"
	"strconv"
)

func Exe3_8() {
	var employee *Employee = new(Employee)
	var input string
	fmt.Print("社員番号を入力してください->")
	fmt.Scanln(&input)
	employee.EmpNo, _ = strconv.Atoi(input)
	fmt.Print("社員名を入力してください->")
	fmt.Scanln(&input)
	employee.EmpName = input
	fmt.Println("社員番号=", employee.EmpNo, "社員名=", employee.EmpName)
	fmt.Println("構造体のアドレス:", reflect.ValueOf(employee).Pointer())
}
