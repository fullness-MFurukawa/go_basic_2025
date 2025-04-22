package chapter08

import (
	"fmt"
	"sync"
	"time"
)

func syncCalc(wait *sync.WaitGroup, operation int) {
	for i := 0; i < 5; i++ {
		if operation == 1 {
			fmt.Printf("加算%d回目:%d\n", i+1, i+i)
		} else if operation == 2 {
			fmt.Printf("乗算%d回目:%d\n", i+1, i*i)
		} else {
			panic("不正な値です")
		}
		time.Sleep(1 * time.Second)
	}
	wait.Done() // カウンタを減らす
}
func Exe8_1() {
	// GroupWaitの作成
	var wait *sync.WaitGroup = new(sync.WaitGroup)
	fmt.Println("処理の開始!!")
	wait.Add(2)          // カウンタを設定する
	go syncCalc(wait, 1) // ゴルーチンを呼び出す
	go syncCalc(wait, 2) // ゴルーチンを呼び出す
	wait.Wait()          // ゴルーチン終了待ち状態にする
	fmt.Println("処理の終了!!")
}
