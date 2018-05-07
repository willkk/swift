package swift

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Init() {
	http.HandleFunc("/", HandleRequest)
	cmds = make(map[string]Command)
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	cmd := NewCommand(w, r)

	if cmd == nil {
		return
	}

	defer cmd.WriteResponse()

	err := cmd.ReadRequest()
	if err != nil {
		return
	}

	cmds[cmd.Name()].Handle()
}

// Store all cmds info
// r.URL.Path => Name. For example, "/user" => "User"
var cmds map[string]Command

func RegisterCommand(path string, cmd Command) {
	cmds[path] = cmd
}

// Every command should implement this interface
type Command interface {
	Name() string
	NewReq() interface{}
	NewResp() interface{}
	New(base *BaseCommand) Command

	ReadRequest() error
	Handle()
	WriteResponse()
}

type BaseCommand struct {
	r    *http.Request
	w    http.ResponseWriter
	Req  interface{}
	Resp interface{}
	Cmd  string // interface name, that is, last path in r.URL.Path.

	// For time tracking
	Start time.Time
}

type UnknownCommand struct {
	BaseCommand
}

func (this *UnknownCommand) Handle() {
	this.Resp = &struct{
		Code int
		Msg  string
	}{1000, "invalid interface"}
}

func (this *UnknownCommand) New(base *BaseCommand) Command {
	return &UnknownCommand{
		*base,
	}
}

func NewCommand(w http.ResponseWriter, r *http.Request) Command {
	// Both of path "/test" and "/test/" correspond to "/test" command
	path := r.URL.Path
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	if cmds[path] == nil {
		return nil
	}

	baseCmd := BaseCommand{
		r:     r,
		w:     w,
		Cmd:   path,
		Start: time.Now(),
		Req:   cmds[path].NewReq(),
		Resp:  cmds[path].NewResp(),
	}

	var cmd Command
	cmdTemplate, ok :=  cmds[path]
	if !ok {
		cmd = &UnknownCommand{
			BaseCommand{
				r: r,
				w: w,
			},
		}
	} else {
		cmd = cmdTemplate.New(&baseCmd)
	}

	return cmd
}

func (this *BaseCommand) Name() string {
	return ""
}

func (this *BaseCommand) NewReq() interface{} {
	return nil
}

func (this *BaseCommand) NewResp() interface{} {
	return nil
}

func (this *BaseCommand) ReadRequest() error {
	cmd := cmds[this.Cmd].Name()
	if cmd == "" {
		return errors.New("invalid request")
	}
	data, err := ioutil.ReadAll(this.r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, this.Req)
	if err != nil {
		return err
	}

	return nil
}

func (this *BaseCommand) Handle() {

}

func (this *BaseCommand) WriteResponse() {
	if this.Resp == nil {
		return
	}
	data, err := json.Marshal(this.Resp)
	if err != nil {
		return
	}
	this.w.Write(data)
}
