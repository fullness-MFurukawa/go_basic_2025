package chapter06

import (
	"fmt"
	"strconv"

	"gopkg.in/yaml.v3"
)

// シリアライズ機能を表すインターフェイス
type Serializer interface {
	Serialize() ([]byte, error)
}

type Account struct {
	No      string `yaml:"no"`
	Name    string `yaml:"name"`
	Balance int    `yaml:"balance"`
}

// インターフェイスの実装
// YAML形式に変換した結果を返す
func (ins *Account) Serialize() ([]byte, error) {
	if yamlData, err := yaml.Marshal(ins); err != nil {
		return nil, err
	} else {
		return yamlData, nil
	}
}
func Exe6_1() {
	var account Account
	fmt.Print("口座番号を入力してください->")
	fmt.Scanln(&account.No)
	fmt.Print("口座名義を入力してください->")
	fmt.Scanln(&account.Name)
	var input string
	fmt.Print("残高を入力してください->")
	fmt.Scanln(&input)
	account.Balance, _ = strconv.Atoi(input)
	result, err := account.Serialize()
	if err == nil {
		fmt.Println(string(result))
	}
}
