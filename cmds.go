package swift

import (
	"net/http"
	"strings"
	"errors"
	"io/ioutil"
	"encoding/json"
)

func Init() {
	http.HandleFunc("/", HandleRequest)
	Cmds = make(map[string]*CmdCore)
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	cmd := newCommand(w, r)

	if cmd == nil {
		return 
	}

	defer cmd.WriteResponse()

	err := cmd.ReadRequest()
	if err != nil {

		return
	}

	Cmds[cmd.Cmd].handler(cmd)
}

// Store all cmds info
// r.URL.Path => Name. For example, "/user" => "User"
var Cmds map[string]*CmdCore

func RegisterCmd(path, name string, req, resp Newable, handler func(*BaseCommand)) {
	Cmds[path] = &CmdCore{req, resp, name, handler}
}

type CmdCore struct {
	req  Newable
	resp Newable
	name string
	handler func(*BaseCommand)
}

type Newable interface {
	New() interface{}
}

type BaseCommand struct {
	r    *http.Request
	w    http.ResponseWriter
	Req  interface{}
	Resp interface{}
	Cmd  string // interface name, that is, last path in r.URL.Path.
}

func newCommand(w http.ResponseWriter, r *http.Request) *BaseCommand {
	// Both of path "/test" and "/test/" correspond to "/test" command
	path := r.URL.Path
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	if Cmds[path] == nil {
		return nil
	}

	baseCmd := &BaseCommand{
		r:   r,
		w:   w,
		Cmd: path,
	}

	baseCmd.Req = Cmds[baseCmd.Cmd].req.New()
	baseCmd.Resp = Cmds[baseCmd.Cmd].resp.New()

	return baseCmd
}

func (this *BaseCommand) ReadRequest() error {
	cmd := Cmds[this.Cmd].name
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
