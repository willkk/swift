package main

import (
	"github.com/willkk/swift"
	"log"
)

type UserReq struct {
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Phone string `json:"phone"`
}

type UserResp struct {
	Code int 	`json:"code"`
	Msg  string `json:"msg"`
}

// It should implement Command interface
type UserCommand struct {
}

func (uc *UserCommand)Name() string {
	return "User"
}

func (uc *UserCommand)NewReq() interface{} {
	return &UserReq{}
}

func (uc *UserCommand)NewResp() interface{} {
	return &UserResp{}
}

func (uc *UserCommand)Handle(bCmd *swift.BaseCommand) {
	req := bCmd.Req.(*UserReq)
	resp := bCmd.Resp.(*UserResp)
	log.Printf("handle user:%v", req)

	resp.Msg = "test ok"
}
