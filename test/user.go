package main

import (
	"log"
	"time"

	"github.com/willkk/swift"
)

type UserReq struct {
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Phone string `json:"phone"`
}

type UserResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// It should implement Command interface
type UserCommand struct {
	swift.BaseCommand
}

func (uc *UserCommand) New(base *swift.BaseCommand) swift.Command {
	return &UserCommand{
		*base,
	}
}

func (uc *UserCommand) Name() string {
	return "User"
}

func (uc *UserCommand) NewReq() interface{} {
	return &UserReq{}
}

func (uc *UserCommand) NewResp() interface{} {
	return &UserResp{}
}

func (uc *UserCommand) Handle() {
	req := uc.Req.(*UserReq)
	resp := uc.Resp.(*UserResp)
	log.Printf("handle user:%v", req)

	resp.Msg = "test ok"

	log.Printf("time used:%d ms", time.Now().Sub(uc.Start).Nanoseconds()/1000000)
}
