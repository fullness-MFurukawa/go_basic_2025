package chapter03

import "fmt"

func Exe3_3() {
	months := []string{"April", "May", "June", "July", "August", "September"}
	months1 := append(months, "October")
	for _, value := range months1 {
		fmt.Println(value)
	}
}
