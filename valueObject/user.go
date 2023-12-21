package main

import (
	"fmt"
	"time"
)

func main() {
	up := userP{id: "1", name: "まきばおう"}
	fmt.Println(up)

	ud := userId{"1"}
	un := userName{"まきばおう"}
	uv := userV{id: ud, name: un}

	fmt.Printf("%+v\n", uv)

	n := nickName{serviceName: "Jump", registDate: time.Now().Unix()}
	nickName := n.createNickName()

	fmt.Printf(nickName)

}

// プリミティブ型
type userP struct {
	id   string
	name string
}

// 値オブジェクト
type userV struct {
	id   userId
	name userName
}

type userId struct {
	value string
}

type userName struct {
	value string
}

type nickName struct {
	serviceName string
	registDate  int64
}

func (n nickName) createNickName() string {
	return fmt.Sprintf("%s%d", n.serviceName, n.registDate)
}

//かろうじてstring型だとわかる
func createUser(nickname string) string {
	//ユーザー登録処理
}

func createUser(nickname nickName) string {
	//ユーザー登録処理
}
