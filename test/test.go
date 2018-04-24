package main

import (
	"github.com/willkk/swift"
	"net/http"
)

func main() {
	swift.Init()

	swift.RegisterCommand("/user", &UserCommand{})

	http.ListenAndServe(":5600", nil)
}