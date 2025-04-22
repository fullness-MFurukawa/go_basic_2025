package chapter11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase01(t *testing.T) {
	// 0歳を指定
	result, _ := AgeCheck(0)
	// 結果を評価する resultが1であればテスト成功
	assert.Equal(t, 1, result)
	// 19歳を指定
	result, _ = AgeCheck(19)
	// 結果を評価する resultが1であればテスト成功
	assert.Equal(t, 1, result)
}

func TestCase02(t *testing.T) {
	// 20歳を指定
	result, _ := AgeCheck(0)
	// 結果を評価する resultが2であればテスト成功
	assert.Equal(t, 2, result)
	// 65際を指定
	result, _ = AgeCheck(65)
	// 結果を評価する resultが2であればテスト成功
	assert.Equal(t, 2, result)
}

func TestCase04(t *testing.T) {
	_, err := AgeCheck(-1) // -1歳の場合
	assert.Equal(t, "不正な年齢です", err.Error(),
		"AgeCheck(-1)の結果はエラーであるべき")
	_, err = AgeCheck(121) // 121歳の場合
	assert.Equal(t, "不正な年齢です", err.Error(),
		"AgeCheck(121)の結果はエラーであるべき")
}


dayofweek := map[string]string{
	"Sunday": "日曜日", "Monday": "月曜日", 
	"Tuesday": "火曜日", "Wednesday": "水曜日",
	"Thursday": "木曜日", "Friday": "金曜日", 
	"Saturday": "土曜日"}