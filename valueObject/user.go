package valueObject

import "fmt"

func main() {
	up := userP{id: "1", name: "まきばおう"}
	fmt.Println(up)

	ud := userId{"1"}
	un := userName{"まきばおう"}
	uv := userV{id: ud, name: un}
	
	fmt.Printf("%+v\n", uv)
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
