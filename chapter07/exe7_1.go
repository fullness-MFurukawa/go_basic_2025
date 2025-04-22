package chapter07

import (
	"fmt"
	"strconv"
)

func Exe7_1() {

	// リカバリ⽤匿名関数
	result := func() {
		err := recover() // リカバリ
		if err != nil {  // パニック発⽣
			fmt.Printf("パニック発⽣:%s \n", err) // エラーを出⼒
		}
	}
	defer result() // リカバリする匿名関数を遅延実⾏

	var input string
	fmt.Print("値1を入力してください->")
	fmt.Scanln(&input)
	value1, err := strconv.Atoi(input)
	if err != nil {
		panic("値1は整数で入力してください")
	}
	fmt.Print("値2を入力してください->")
	fmt.Scanln(&input)
	value2, err := strconv.Atoi(input)
	if err != nil {
		panic("値2は整数で入力してください")
	}
	fmt.Println("値1 + 値2 = ", value1+value2)
}
