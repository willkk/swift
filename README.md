# swift
A simple and fast API framework to handle general http requests. With this framework, you can organize your code
in modules, making your codes more reusable and portable between projects.
Advantages:
1. More business codes, no more application architecture.
2. More modularization&reusability, no more repeating.
3. More thinking, no more typing.

## Example

```go
// user.go

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

// test.go

func main() {
	swift.Init()

	swift.RegisterCommand("/user", &UserCommand{})

	http.ListenAndServe(":5600", nil)
}
```
You can send request using like:  
curl http://localhost:5600/user -d '{"name":"xcww", "sex":"male", "phone":"110"}' // correct request  
curl http://localhost:5600/test -d '{"test":"0000"}' // unknown request
