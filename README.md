# swift
A simple and practical API framework to handle general http requests. With this framework, you can organize your code
in modules, making your codes more reusable and portable between projects.  
Advantages:
1. More focus on business codes, not application architecture.
2. More modularization&reusability, not repeating.
3. More thinking, not typing.

## Example

```go
// user.go

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

func (uc *UserCommand)Name() string {
	return "User"
}

func (uc *UserCommand)NewReq() interface{} {
	return &UserReq{}
}

func (uc *UserCommand)NewResp() interface{} {
	return &UserResp{}
}

func (uc *UserCommand) Handle() {
	req := uc.Req.(*UserReq)
	resp := uc.Resp.(*UserResp)
	log.Printf("handle user:%v", req)

	resp.Msg = "test ok"

	log.Printf("time used:%d ms", time.Now().Sub(uc.Start).Nanoseconds()/1000000)
}

// test.go

func main() {
	swift.Init()

	swift.RegisterCommand("/user", &UserCommand{})

	http.ListenAndServe(":5600", nil)
}
```
You can send requests using commands like:  
curl http://localhost:5600/user -d '{"name":"xcww", "sex":"male", "phone":"110"}' // correct request  
curl http://localhost:5600/test -d '{"test":"0000"}' // unknown request
