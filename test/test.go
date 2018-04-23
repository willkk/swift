package main

import (
	"github.com/willkk/swift"
	"net/http"
	"log"
)

type UserReq struct {
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Phone string `json:"phone"`
}

func (ur *UserReq)New() interface{} {
	return &UserReq{}
}

type UserResp struct {
	Code int 	`json:"code"`
	Msg  string `json:"msg"`
}

func (urs *UserResp)New() interface{} {
	return &UserResp{}
}

func HandleUser(command *swift.BaseCommand) {
	req := command.Req.(*UserReq)
	resp := command.Resp.(*UserResp)
	log.Printf("handle user:%v", req)

	resp.Msg = "test ok"
}

func main() {
	swift.Init()

	swift.RegisterCmd("/user", "User",&UserReq{} , &UserResp{}, HandleUser)

	http.ListenAndServe(":5600", nil)
}