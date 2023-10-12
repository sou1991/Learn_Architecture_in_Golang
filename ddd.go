package main

import "fmt"

func main() {
	var ur IUserRepo = NewUserRepo()
	var uc IUserCntl = NewUserCntl(ur)
	var r IUserRouter = NewUserRouter(uc)
	r.Create()

}

// -------Repository------------------------
type IUserRepo interface {
	Create()
}

type UserRepo struct {
	Conn string
}

func NewUserRepo() IUserRepo {
	return UserRepo{Conn: "Postgres"}
}

func (UserRepo) Create() {
	fmt.Println("Call")
}

//-------Repository------------------------

//-------Controller------------------------

type UserCntl struct {
	Ur IUserRepo
}

type IUserCntl interface {
	Create()
}

func NewUserCntl(ur IUserRepo) IUserCntl {
	return UserCntl{Ur: ur}
}

func (uc UserCntl) Create() {
	uc.Ur.Create()
}

//-------Controller------------------------

//-------Router----------------------------

type UserRouter struct {
	Uc IUserCntl
}

type IUserRouter interface {
	Create()
}

func NewUserRouter(uc IUserCntl) IUserRouter {
	return UserRouter{Uc: uc}
}

func (ur UserRouter) Create() {
	ur.Uc.Create()
}

//-------Router----------------------------
