package chapter08

import (
	"fmt"
	"time"
)

func selectCalc(operation int, c chan<- string, end chan<- int) {
	for i := 0; i < 5; i++ {
		if operation == 1 {
			c <- fmt.Sprintf("加算%d回目:%d", i+1, i+i)
		} else if operation == 2 {
			c <- fmt.Sprintf("乗算%d回目:%d", i+1, i*i)
		} else {
			panic("不正な値です")
		}
		time.Sleep(1 * time.Second)
	}
	end <- 1
}

func Exe8_2() {
	// 前処理
	stringchannel := make(chan string) // 計算結果用チャネル
	endchannel := make(chan int)       // 終了通知⽤チャネルの作成
	// 後処理
	close := func() { // チャネルのクローズ
		close(stringchannel)
		close(endchannel)
		fmt.Println("チャネルのクローズ")
	}
	defer close()
	// 主処理
	go selectCalc(1, stringchannel, endchannel)
	go selectCalc(2, stringchannel, endchannel)
	endcount := 0 // 終了判定カウンタ
	for {
		select {
		case svale := <-stringchannel:
			fmt.Println(svale)
		case eval := <-endchannel:
			endcount += eval
		}
		if endcount == 2 {
			break
		}
	}
}
